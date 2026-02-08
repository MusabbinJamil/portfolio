<script lang="ts">
	import type { Experience } from '$lib/types';

	let { experience }: { experience: Experience[] } = $props();
</script>

<section id="experience">
	<h2>Experience</h2>
	<div class="timeline">
		{#each experience as job}
			<article class="job">
				<div class="header">
					<h3>{job.role}</h3>
					<span class="period">{job.period}</span>
				</div>
				<p class="company">{job.company}</p>
				<p class="description">{job.description}</p>
			</article>
		{/each}
	</div>
</section>

<style>
	.timeline {
		display: flex;
		flex-direction: column;
		gap: 24px;
	}
	.job {
		background: #1e293b;
		border-radius: 12px;
		padding: 28px;
		border: 1px solid #334155;
		transition: border-color 0.2s;
	}
	.job:hover {
		border-color: #38bdf8;
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
		color: #f8fafc;
		font-size: 1.3rem;
	}
	.period {
		color: #38bdf8;
		font-size: 0.9rem;
		font-weight: 500;
	}
	.company {
		color: #cbd5e1;
		font-size: 1.05rem;
		font-weight: 500;
		margin-bottom: 12px;
	}
	.description {
		color: #94a3b8;
		line-height: 1.7;
		font-size: 0.95rem;
	}
</style>
