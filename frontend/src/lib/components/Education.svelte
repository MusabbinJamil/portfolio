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
		background: #0a0a0a;
		border-radius: 12px;
		padding: 28px;
		border: 1px solid #003300;
		transition: border-color 0.2s;
	}
	.edu-card:hover {
		border-color: #00ff41;
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
		color: #00ff41;
		font-size: 1.2rem;
	}
	.year {
		color: #00ff41;
		font-size: 0.9rem;
		font-weight: 500;
	}
	.institution {
		color: #00cc33;
		font-size: 1.05rem;
		font-weight: 500;
		margin-bottom: 8px;
	}
	.details {
		color: #008f11;
		line-height: 1.7;
		font-size: 0.9rem;
	}
</style>
