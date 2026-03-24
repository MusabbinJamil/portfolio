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
				{#if job.techStack && job.techStack.length > 0}
					<ul class="tech">
						{#each job.techStack as tech}
							<li>{tech}</li>
						{/each}
					</ul>
				{/if}
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
		background: var(--bg-card);
		border-radius: 12px;
		padding: 28px;
		border: 1px solid var(--border);
		box-shadow: 0 4px 24px var(--shadow);
		transition: border-color 0.2s, background 0.3s;
	}
	.job:hover {
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
		font-size: clamp(1.3rem, 2vw, 1.925rem);
	}
	.period {
		color: var(--accent-soft);
		font-size: clamp(0.9rem, 1.3vw, 1.525rem);
		font-weight: 500;
	}
	.company {
		color: var(--text-primary);
		font-size: clamp(1.05rem, 1.5vw, 1.675rem);
		font-weight: 500;
		margin-bottom: 12px;
	}
	.description {
		color: var(--text-secondary);
		line-height: 1.7;
		font-size: clamp(0.95rem, 1.4vw, 1.575rem);
		margin-bottom: 12px;
	}
	.tech {
		display: flex;
		flex-wrap: wrap;
		gap: 8px;
		list-style: none;
		padding: 0;
	}
	.tech li {
		background: var(--bg-badge);
		color: var(--accent);
		padding: 4px 12px;
		border-radius: 4px;
		font-size: clamp(0.85rem, 1.2vw, 1.475rem);
		font-weight: 500;
		border: 1px solid var(--border);
		transition: background 0.3s, border-color 0.3s;
	}

	@media (max-width: 768px) {
		.job {
			padding: 20px;
		}
	}
</style>
