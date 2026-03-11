<script lang="ts">
	import type { HeroData } from '$lib/types';
	import { trackClick } from '$lib/analytics';

	let { hero }: { hero: HeroData } = $props();
</script>

<section class="hero">
	<div class="card">
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
		padding: 120px 20px 80px;
		max-width: 900px;
		margin: 0 auto;
	}
	.card {
		background: var(--bg-card);
		border: 1px solid var(--border);
		border-radius: 16px;
		padding: 48px 40px;
		box-shadow: 0 4px 24px var(--shadow);
		transition: background 0.3s, border-color 0.3s;
	}
	.greeting {
		font-size: 1.725rem;
		color: var(--accent-soft);
		margin-bottom: 12px;
		font-weight: 500;
	}
	.name {
		font-size: 4.125rem;
		font-weight: 700;
		color: var(--accent);
		margin-bottom: 8px;
		line-height: 1.2;
	}
	.title {
		font-size: 2.125rem;
		color: var(--text-primary);
		margin-bottom: 16px;
		font-weight: 600;
	}
	.tagline {
		font-size: 1.875rem;
		color: var(--text-muted);
		max-width: 600px;
		margin-bottom: 32px;
	}
	.cta {
		display: inline-block;
		padding: 14px 36px;
		background: var(--accent-soft);
		color: #ffffff;
		border-radius: 8px;
		font-weight: 600;
		font-size: 1.625rem;
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
		.greeting {
			font-size: 1rem;
		}
		.name {
			font-size: 2.25rem;
		}
		.title {
			font-size: 1.35rem;
		}
		.tagline {
			font-size: 1.1rem;
		}
		.cta {
			font-size: 1rem;
			padding: 12px 28px;
		}
	}
</style>
