# Deploying Portfolio to AWS + Namecheap Domain

## Prerequisites

- AWS EC2 Ubuntu instance with SSH access
- Namecheap domain (musab.me) with GitHub Student free SSL
- Go 1.22+ installed on the server
- Node.js 18+ and npm installed on the server
- Nginx installed on the server (`sudo apt install nginx`)

---

## Step 1: DNS on Namecheap

1. Log in to **Namecheap** > **Domain List** > click **Manage** on `musab.me`
2. Go to **Advanced DNS**
3. Add a new record:

   | Type     | Host      | Value              | TTL       |
   |----------|-----------|--------------------|-----------|
   | A Record | portfolio | YOUR_AWS_SERVER_IP | Automatic |

4. Replace `YOUR_AWS_SERVER_IP` with your EC2 public IP address
5. Wait a few minutes for DNS to propagate
6. Verify with: `ping portfolio.musab.me`

---

## Step 2: Get Namecheap SSL (GitHub Student)

### 2a: Activate the SSL certificate

1. Go to **GitHub Student Developer Pack**: https://education.github.com/pack
2. Find the **Namecheap** offer and claim your free SSL certificate
3. In Namecheap, go to **SSL Certificates** > **Activate** on your certificate
4. For **Primary Domain**, enter: `portfolio.musab.me`
5. Choose **HTTP-based validation** (DCV method)
6. Namecheap will give you a validation file (a `.txt` file with a specific name and content)

### 2b: Validate domain ownership

On your AWS server, create the validation file:

```bash
sudo mkdir -p /var/www/html/.well-known/pki-validation/
sudo nano /var/www/html/.well-known/pki-validation/FILENAME.txt
```

Replace `FILENAME.txt` with the exact filename Namecheap gives you, and paste the content they provide.

Set up a temporary nginx config to serve the validation:

```bash
sudo tee /etc/nginx/sites-available/default << 'NGINX'
server {
    listen 80;
    server_name portfolio.musab.me;
    root /var/www/html;

    location / {
        try_files $uri $uri/ =404;
    }
}
NGINX

sudo nginx -t && sudo systemctl reload nginx
```

Go back to Namecheap and click **Verify**. Once validated, Namecheap will email you the SSL certificate files.

### 2c: Install the SSL certificate on your server

You will receive a `.zip` file containing:
- `portfolio_musab_me.crt` (your certificate)
- `portfolio_musab_me.ca-bundle` (CA bundle)

Upload them to your server and combine into a single file:

```bash
sudo mkdir -p /etc/ssl/portfolio
# Upload the files to /etc/ssl/portfolio/ via scp:
# scp portfolio_musab_me.crt portfolio_musab_me.ca-bundle musab@YOUR_SERVER_IP:/etc/ssl/portfolio/

# Combine certificate and CA bundle
sudo bash -c 'cat /etc/ssl/portfolio/portfolio_musab_me.crt /etc/ssl/portfolio/portfolio_musab_me.ca-bundle > /etc/ssl/portfolio/portfolio_musab_me_combined.crt'
```

### 2d: Generate a private key (if you haven't already)

If you generated the CSR yourself:

```bash
# You should already have the private key from when you generated the CSR
# If not, generate a new CSR + private key:
sudo openssl req -new -newkey rsa:2048 -nodes \
  -keyout /etc/ssl/portfolio/portfolio_musab_me.key \
  -out /etc/ssl/portfolio/portfolio_musab_me.csr \
  -subj "/CN=portfolio.musab.me"
```

If you generated a new CSR, you'll need to re-activate the SSL certificate on Namecheap with this new CSR.

---

## Step 3: Clone and Build on the Server

SSH into your AWS server:

```bash
ssh musab@YOUR_AWS_SERVER_IP
```

Clone and build:

```bash
# Clone the repo
git clone https://github.com/MusabbinJamil/portfolio.git ~/portfolio
cd ~/portfolio

# Build Go backend
cd backend
go build -o portfolio-backend .
cd ..

# Build SvelteKit static files
cd frontend
npm install
npm run build
cd ..
```

---

## Step 4: Nginx Configuration

Create the nginx config:

```bash
sudo tee /etc/nginx/sites-available/portfolio << 'NGINX'
# Redirect HTTP to HTTPS
server {
    listen 80;
    server_name portfolio.musab.me;
    return 301 https://$host$request_uri;
}

# HTTPS server
server {
    listen 443 ssl;
    server_name portfolio.musab.me;

    # SSL certificate files
    ssl_certificate     /etc/ssl/portfolio/portfolio_musab_me_combined.crt;
    ssl_certificate_key /etc/ssl/portfolio/portfolio_musab_me.key;

    # SSL settings
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers HIGH:!aNULL:!MD5;
    ssl_prefer_server_ciphers on;

    # Serve static SvelteKit build
    root /home/musab/portfolio/frontend/build;
    index index.html;

    # Proxy API requests to Go backend
    location /api/ {
        proxy_pass http://127.0.0.1:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    # SPA fallback
    location / {
        try_files $uri $uri/ /index.html;
    }
}
NGINX
```

Enable the site:

```bash
# Remove default site if it exists
sudo rm -f /etc/nginx/sites-enabled/default

# Enable portfolio site
sudo ln -sf /etc/nginx/sites-available/portfolio /etc/nginx/sites-enabled/portfolio

# Test and reload
sudo nginx -t
sudo systemctl reload nginx
```

---

## Step 5: Run Go Backend as a Service

Create a systemd service so Go runs automatically:

```bash
sudo tee /etc/systemd/system/portfolio-api.service << 'EOF'
[Unit]
Description=Portfolio Go API
After=network.target

[Service]
ExecStart=/home/musab/portfolio/backend/portfolio-backend
WorkingDirectory=/home/musab/portfolio/backend
Restart=always
User=musab
Environment=PORT=3000

[Install]
WantedBy=multi-user.target
EOF
```

Enable and start:

```bash
sudo systemctl daemon-reload
sudo systemctl enable portfolio-api
sudo systemctl start portfolio-api

# Check it's running
sudo systemctl status portfolio-api
```

---

## Step 6: AWS Security Group

In the **AWS Console** > **EC2** > **Security Groups** > your instance's security group:

Add these **inbound rules**:

| Type  | Port | Source    |
|-------|------|-----------|
| HTTP  | 80   | 0.0.0.0/0 |
| HTTPS | 443  | 0.0.0.0/0 |

---

## Step 7: Verify

1. Open `https://portfolio.musab.me` in your browser
2. Check the padlock icon for valid SSL
3. Test the contact form submission
4. Verify API works: `curl https://portfolio.musab.me/api/health`

---

## Updating the Site

When you push new code:

```bash
ssh musab@YOUR_AWS_SERVER_IP
cd ~/portfolio
git pull

# Rebuild backend (if Go code changed)
cd backend && go build -o portfolio-backend . && cd ..
sudo systemctl restart portfolio-api

# Rebuild frontend (if Svelte code changed)
cd frontend && npm run build
```

No nginx reload needed for frontend changes — the static files are served directly.

---

## Troubleshooting

| Problem | Fix |
|---------|-----|
| Site not loading | Check: `sudo systemctl status nginx` and `sudo systemctl status portfolio-api` |
| 502 Bad Gateway on /api | Go backend not running: `sudo systemctl restart portfolio-api` |
| SSL certificate error | Check cert paths: `sudo ls /etc/ssl/portfolio/` |
| DNS not resolving | Wait 10-30 min, verify A record on Namecheap, try `dig portfolio.musab.me` |
| Permission denied on build dir | Run: `chmod -R 755 /home/musab/portfolio/frontend/build` |
