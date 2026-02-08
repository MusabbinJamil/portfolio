<nav>
	<div class="logo">Musab</div>
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
		background: rgba(0, 0, 0, 0.95);
		backdrop-filter: blur(8px);
		border-bottom: 1px solid #003300;
		z-index: 1000;
		padding: 16px 24px;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}
	.logo {
		font-size: 1.4rem;
		font-weight: 700;
		color: #00ff41;
	}
	.links {
		display: flex;
		gap: 24px;
		list-style: none;
	}
	.links a {
		color: #008f11;
		text-decoration: none;
		font-weight: 500;
		transition: color 0.2s;
	}
	.links a:hover {
		color: #00ff41;
	}
</style>
