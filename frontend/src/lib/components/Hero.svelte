<script lang="ts">
	import type { HeroData } from '$lib/types';
	import { trackClick } from '$lib/analytics';

	let { hero }: { hero: HeroData } = $props();
</script>

<section class="hero">
	<div class="card">
		<img src="/2026.jpg" alt="Musab Bin Jamil" class="photo" />
		<p class="greeting">Hi, my name is</p>
		<h1 class="name">{hero.name}</h1>
		<p class="title">{hero.title}</p>
		<p class="tagline">{hero.tagline}</p>
		<a href="#projects" class="cta" onclick={() => trackClick('hero:cta', 'View My Work button')}>View My Work</a>
	</div>
</section>

<style>
	.hero {
		min-height: 100vh;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		text-align: center;
		padding: 120px 5vw 80px;
		max-width: 100%;
	}
	.card {
		background: var(--bg-card);
		border: 1px solid var(--border);
		border-radius: 16px;
		padding: 48px 40px;
		box-shadow: 0 4px 24px var(--shadow);
		transition: background 0.3s, border-color 0.3s;
		width: 100%;
	}
	.photo {
		width: 180px;
		height: 180px;
		border-radius: 50%;
		object-fit: cover;
		border: 3px solid var(--accent-soft);
		margin-bottom: 24px;
	}
	.greeting {
		font-size: clamp(1rem, 1.5vw, 1.725rem);
		color: var(--accent-soft);
		margin-bottom: 12px;
		font-weight: 500;
	}
	.name {
		font-family: Papyrus, fantasy, cursive;
		font-size: clamp(2.25rem, 5vw, 4.125rem);
		font-weight: 700;
		color: var(--accent);
		margin-bottom: 8px;
		line-height: 1.2;
	}
	.title {
		font-size: clamp(1.35rem, 2.5vw, 2.125rem);
		color: var(--text-primary);
		margin-bottom: 16px;
		font-weight: 600;
	}
	.tagline {
		font-size: clamp(1.1rem, 2vw, 1.875rem);
		color: var(--text-muted);
		max-width: 600px;
		margin: 0 auto 32px;
	}
	.cta {
		display: inline-block;
		padding: 14px 36px;
		background: var(--accent-soft);
		color: #ffffff;
		border-radius: 8px;
		font-weight: 600;
		font-size: clamp(1rem, 1.5vw, 1.625rem);
		text-decoration: none;
		transition: background 0.2s, box-shadow 0.2s;
	}
	.cta:hover {
		background: var(--accent-hover);
		box-shadow: 0 4px 16px var(--shadow-hover);
		text-decoration: none;
	}

	@media (max-width: 768px) {
		.hero {
			padding: 100px 16px 48px;
		}
		.card {
			padding: 28px 20px;
		}
		.photo {
			width: 130px;
			height: 130px;
		}
	}
</style>
