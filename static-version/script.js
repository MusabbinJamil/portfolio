// Theme toggle
function toggleTheme() {
  const current = document.documentElement.getAttribute('data-theme');
  const next = current === 'dark' ? 'light' : 'dark';
  document.documentElement.setAttribute('data-theme', next);
  localStorage.setItem('theme', next);
  document.querySelector('.theme-toggle').innerHTML = next === 'dark' ? '&#9728;' : '&#9790;';
}

(function () {
  const saved = localStorage.getItem('theme');
  if (saved === 'dark') {
    document.documentElement.setAttribute('data-theme', 'dark');
    document.querySelector('.theme-toggle').innerHTML = '&#9728;';
  }
})();

// Close mobile menu on link click
document.querySelectorAll('.nav-links a').forEach(function (a) {
  a.addEventListener('click', function () {
    document.querySelector('.nav-links').classList.remove('open');
  });
});

// Contact form — mailto fallback (no backend)
document.getElementById('contactForm').addEventListener('submit', function (e) {
  e.preventDefault();
  var fd = new FormData(this);
  var name = fd.get('name');
  var email = fd.get('email');
  var message = fd.get('message');
  var subject = encodeURIComponent('Portfolio Contact from ' + name);
  var body = encodeURIComponent('Name: ' + name + '\nEmail: ' + email + '\n\n' + message);
  window.location.href = 'mailto:musabjamil3@gmail.com?subject=' + subject + '&body=' + body;
  document.getElementById('formMsg').textContent = 'Opening your email client...';
  document.getElementById('formMsg').className = 'form-msg success';
});

// Matrix rain
(function () {
  var canvas = document.getElementById('matrix');
  if (!canvas) return;
  var ctx = canvas.getContext('2d');
  var chars = '~*.+o:^';
  var colors = ['#7bc89b', '#a8ddb8', '#5db87e', '#b8e6c8', '#9ccfd8'];
  var particles = [];
  var fontSize = 14;

  function resize() {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;
    var count = Math.floor(canvas.width / fontSize);
    particles = [];
    for (var i = 0; i < count; i++) {
      particles.push({
        x: i * fontSize,
        y: Math.random() * canvas.height,
        speed: 0.5 + Math.random() * 1.5,
        opacity: 0.15 + Math.random() * 0.25,
        color: colors[Math.floor(Math.random() * colors.length)]
      });
    }
  }

  function draw() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.font = fontSize + 'px monospace';
    for (var i = 0; i < particles.length; i++) {
      var p = particles[i];
      ctx.globalAlpha = p.opacity;
      ctx.fillStyle = p.color;
      ctx.fillText(chars[Math.floor(Math.random() * chars.length)], p.x, p.y);
      p.y += p.speed;
      if (p.y > canvas.height) {
        p.y = -fontSize;
      }
    }
    requestAnimationFrame(draw);
  }

  window.addEventListener('resize', resize);
  resize();
  draw();
})();
