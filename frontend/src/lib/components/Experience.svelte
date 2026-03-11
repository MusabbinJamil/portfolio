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
		background: #ffffff;
		border-radius: 12px;
		padding: 28px;
		border: 1px solid #e8dff5;
		box-shadow: 0 4px 24px rgba(196, 167, 231, 0.12);
		transition: border-color 0.2s;
	}
	.job:hover {
		border-color: #c4a7e7;
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
		color: #907aa9;
		font-size: 1.925rem;
	}
	.period {
		color: #c4a7e7;
		font-size: 1.525rem;
		font-weight: 500;
	}
	.company {
		color: #575279;
		font-size: 1.675rem;
		font-weight: 500;
		margin-bottom: 12px;
	}
	.description {
		color: #6e6a86;
		line-height: 1.7;
		font-size: 1.575rem;
	}

	@media (max-width: 768px) {
		.job {
			padding: 20px;
		}
		h3 {
			font-size: 1.3rem;
		}
		.period {
			font-size: 0.9rem;
		}
		.company {
			font-size: 1.05rem;
		}
		.description {
			font-size: 0.95rem;
		}
	}
</style>
