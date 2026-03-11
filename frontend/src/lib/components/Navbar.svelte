<script lang="ts">
	import { onMount } from 'svelte';
	import { trackClick } from '$lib/analytics';

	let menuOpen = $state(false);
	let dark = $state(false);

	onMount(() => {
		dark = localStorage.getItem('theme') === 'dark';
		applyTheme();
	});

	function applyTheme() {
		document.documentElement.setAttribute('data-theme', dark ? 'dark' : 'light');
	}

	function toggleTheme() {
		dark = !dark;
		localStorage.setItem('theme', dark ? 'dark' : 'light');
		applyTheme();
	}

	function handleNavClick(section: string) {
		trackClick(`nav:${section}`, `${section} nav link`);
		menuOpen = false;
	}
</script>

<nav>
	<div class="logo">Musab</div>
	<div class="nav-links-area">
		<ul class="links" class:open={menuOpen}>
			<li><a href="#about" onclick={() => handleNavClick('about')}>About</a></li>
			<li><a href="#experience" onclick={() => handleNavClick('experience')}>Experience</a></li>
			<li><a href="#projects" onclick={() => handleNavClick('projects')}>Projects</a></li>
			<li><a href="#education" onclick={() => handleNavClick('education')}>Education</a></li>
			<li><a href="#contact" onclick={() => handleNavClick('contact')}>Contact</a></li>
		</ul>
		<button class="theme-toggle" onclick={toggleTheme} aria-label="Toggle dark mode">
			{#if dark}
				<svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
					<circle cx="12" cy="12" r="5"/>
					<line x1="12" y1="1" x2="12" y2="3"/>
					<line x1="12" y1="21" x2="12" y2="23"/>
					<line x1="4.22" y1="4.22" x2="5.64" y2="5.64"/>
					<line x1="18.36" y1="18.36" x2="19.78" y2="19.78"/>
					<line x1="1" y1="12" x2="3" y2="12"/>
					<line x1="21" y1="12" x2="23" y2="12"/>
					<line x1="4.22" y1="19.78" x2="5.64" y2="18.36"/>
					<line x1="18.36" y1="5.64" x2="19.78" y2="4.22"/>
				</svg>
			{:else}
				<svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
					<path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"/>
				</svg>
			{/if}
		</button>
		<button class="hamburger" onclick={() => (menuOpen = !menuOpen)} aria-label="Toggle menu">
			<span class="bar"></span>
			<span class="bar"></span>
			<span class="bar"></span>
		</button>
	</div>
</nav>

<style>
	nav {
		position: fixed;
		top: 0;
		width: 100%;
		background: var(--nav-bg);
		backdrop-filter: blur(12px);
		border-bottom: 1px solid var(--border);
		z-index: 1000;
		padding: 16px 24px;
		display: flex;
		justify-content: space-between;
		align-items: center;
		transition: background 0.3s, border-color 0.3s;
	}
	.logo {
		font-size: 2.025rem;
		font-weight: 700;
		color: var(--accent);
	}
	.nav-links-area {
		display: flex;
		align-items: center;
		gap: 20px;
	}
	.theme-toggle {
		background: none;
		border: 1px solid var(--border);
		border-radius: 8px;
		padding: 8px 10px;
		cursor: pointer;
		color: var(--accent-soft);
		display: flex;
		align-items: center;
		justify-content: center;
		transition: border-color 0.2s, color 0.2s;
	}
	.theme-toggle:hover {
		border-color: var(--accent-soft);
		color: var(--accent);
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
		background: var(--accent-soft);
		border-radius: 2px;
	}
	.links {
		display: flex;
		gap: 24px;
		list-style: none;
	}
	.links a {
		color: var(--text-muted);
		text-decoration: none;
		font-weight: 500;
		font-size: 1.575rem;
		transition: color 0.2s;
	}
	.links a:hover {
		color: var(--accent);
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
			right: 0;
			width: 100%;
			background: var(--nav-bg-mobile);
			border-bottom: 1px solid var(--border);
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
