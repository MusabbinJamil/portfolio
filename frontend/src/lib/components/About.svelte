<script lang="ts">
	import type { AboutData } from '$lib/types';

	let { about }: { about: AboutData } = $props();
</script>

<section id="about">
	<h2>About Me</h2>
	<div class="card">
		<p class="bio">{about.bio}</p>
	</div>
</section>

<style>
	.card {
		background: var(--bg-card);
		border: 1px solid var(--border);
		border-radius: 16px;
		padding: 40px;
		box-shadow: 0 4px 24px var(--shadow);
		transition: background 0.3s, border-color 0.3s;
	}
	.bio {
		color: var(--text-secondary);
		font-size: clamp(1.05rem, 1.5vw, 1.675rem);
		line-height: 1.8;
	}

	@media (max-width: 768px) {
		.card {
			padding: 24px 20px;
		}
	}
</style>
