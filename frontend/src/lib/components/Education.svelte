<script lang="ts">
	import type { Education } from '$lib/types';

	let { education }: { education: Education[] } = $props();
</script>

<section id="education">
	<h2>Education</h2>
	<div class="list">
		{#each education as edu}
			<article class="edu-card">
				<div class="header">
					<h3>{edu.degree}</h3>
					<span class="year">{edu.year}</span>
				</div>
				<p class="institution">{edu.institution}</p>
				{#if edu.details}
					<p class="details">{edu.details}</p>
				{/if}
			</article>
		{/each}
	</div>
</section>

<style>
	.list {
		display: flex;
		flex-direction: column;
		gap: 24px;
	}
	.edu-card {
		background: var(--bg-card);
		border-radius: 12px;
		padding: 28px;
		border: 1px solid var(--border);
		box-shadow: 0 4px 24px var(--shadow);
		transition: border-color 0.2s, background 0.3s;
	}
	.edu-card:hover {
		border-color: var(--accent-soft);
	}
	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 4px;
		flex-wrap: wrap;
		gap: 8px;
	}
	h3 {
		color: var(--accent);
		font-size: clamp(1.2rem, 2vw, 1.825rem);
	}
	.year {
		color: var(--accent-soft);
		font-size: clamp(0.9rem, 1.3vw, 1.525rem);
		font-weight: 500;
	}
	.institution {
		color: var(--text-primary);
		font-size: clamp(1.05rem, 1.5vw, 1.675rem);
		font-weight: 500;
		margin-bottom: 8px;
	}
	.details {
		color: var(--text-secondary);
		line-height: 1.7;
		font-size: clamp(0.9rem, 1.3vw, 1.525rem);
	}

	@media (max-width: 768px) {
		.edu-card {
			padding: 20px;
		}
	}
</style>
