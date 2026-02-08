<script lang="ts">
	let menuOpen = $state(false);

	function closeMenu() {
		menuOpen = false;
	}
</script>

<nav>
	<div class="logo">Musab</div>
	<button class="hamburger" onclick={() => (menuOpen = !menuOpen)} aria-label="Toggle menu">
		<span class="bar"></span>
		<span class="bar"></span>
		<span class="bar"></span>
	</button>
	<ul class="links" class:open={menuOpen}>
		<li><a href="#about" onclick={closeMenu}>About</a></li>
		<li><a href="#experience" onclick={closeMenu}>Experience</a></li>
		<li><a href="#projects" onclick={closeMenu}>Projects</a></li>
		<li><a href="#education" onclick={closeMenu}>Education</a></li>
		<li><a href="#contact" onclick={closeMenu}>Contact</a></li>
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
		font-size: 2.025rem;
		font-weight: 700;
		color: #00ff41;
	}
	.hamburger {
		display: none;
		flex-direction: column;
		gap: 5px;
		background: none;
		border: none;
		cursor: pointer;
		padding: 4px;
	}
	.bar {
		display: block;
		width: 28px;
		height: 3px;
		background: #00ff41;
		border-radius: 2px;
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
		font-size: 1.575rem;
		transition: color 0.2s;
	}
	.links a:hover {
		color: #00ff41;
	}

	@media (max-width: 768px) {
		.logo {
			font-size: 1.4rem;
		}
		.hamburger {
			display: flex;
		}
		.links {
			display: none;
			position: absolute;
			top: 100%;
			left: 0;
			width: 100%;
			background: rgba(0, 0, 0, 0.95);
			border-bottom: 1px solid #003300;
			flex-direction: column;
			gap: 0;
			padding: 8px 0;
		}
		.links.open {
			display: flex;
		}
		.links a {
			font-size: 1.1rem;
			padding: 12px 24px;
			display: block;
		}
	}
</style>
