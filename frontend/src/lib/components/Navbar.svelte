<nav>
	<div class="logo">
		<img src="/logo.jpg" alt="Logo" />
		Musab
	</div>
	<ul class="links">
		<li><a href="#about">About</a></li>
		<li><a href="#experience">Experience</a></li>
		<li><a href="#projects">Projects</a></li>
		<li><a href="#education">Education</a></li>
		<li><a href="#contact">Contact</a></li>
	</ul>
</nav>

<style>
	nav {
		position: fixed;
		top: 0;
		width: 100%;
		background: rgba(15, 23, 42, 0.95);
		backdrop-filter: blur(8px);
		border-bottom: 1px solid #1e293b;
		z-index: 1000;
		padding: 16px 24px;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}
	.logo {
		font-size: 1.4rem;
		font-weight: 700;
		color: #38bdf8;
		display: flex;
		align-items: center;
		gap: 10px;
	}
	.logo img {
		height: 36px;
		width: auto;
	}
	.links {
		display: flex;
		gap: 24px;
		list-style: none;
	}
	.links a {
		color: #94a3b8;
		text-decoration: none;
		font-weight: 500;
		transition: color 0.2s;
	}
	.links a:hover {
		color: #f8fafc;
	}
</style>
