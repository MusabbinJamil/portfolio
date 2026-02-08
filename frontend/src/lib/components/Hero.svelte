<script lang="ts">
	import type { HeroData } from '$lib/types';

	let { hero }: { hero: HeroData } = $props();
</script>

<section class="hero">
	<p class="greeting">Hi, my name is</p>
	<h1 class="name">{hero.name}</h1>
	<p class="title">{hero.title}</p>
	<p class="tagline">{hero.tagline}</p>
	<a href="#projects" class="cta">View My Work</a>
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
	.greeting {
		font-size: 1.1rem;
		color: #00ff41;
		margin-bottom: 12px;
		font-weight: 500;
	}
	.name {
		font-size: 3.5rem;
		font-weight: 700;
		color: #00ff41;
		margin-bottom: 8px;
		line-height: 1.2;
		text-shadow: 0 0 10px rgba(0, 255, 65, 0.5);
	}
	.title {
		font-size: 1.5rem;
		color: #00cc33;
		margin-bottom: 16px;
		font-weight: 600;
	}
	.tagline {
		font-size: 1.25rem;
		color: #008f11;
		max-width: 600px;
		margin-bottom: 32px;
	}
	.cta {
		display: inline-block;
		padding: 12px 32px;
		background: #00ff41;
		color: #000000;
		border-radius: 8px;
		font-weight: 600;
		font-size: 1rem;
		text-decoration: none;
		transition: background 0.2s, box-shadow 0.2s;
	}
	.cta:hover {
		background: #33ff66;
		box-shadow: 0 0 15px rgba(0, 255, 65, 0.4);
		text-decoration: none;
	}
</style>
