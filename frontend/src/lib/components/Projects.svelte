<script lang="ts">
	import type { Project } from '$lib/types';

	let { projects }: { projects: Project[] } = $props();
</script>

<section id="projects">
	<h2>Projects</h2>
	<div class="grid">
		{#each projects as project}
			<article class="card">
				<h3>{project.title}</h3>
				<p class="description">{project.description}</p>
				<ul class="tech">
					{#each project.techStack as tech}
						<li>{tech}</li>
					{/each}
				</ul>
				<div class="links">
					{#if project.githubUrl}
						<a href={project.githubUrl} target="_blank" rel="noopener">GitHub</a>
					{/if}
					{#if project.liveUrl}
						<a href={project.liveUrl} target="_blank" rel="noopener">Live Demo</a>
					{/if}
				</div>
			</article>
		{/each}
	</div>
</section>

<style>
	.grid {
		display: grid;
		gap: 24px;
	}
	.card {
		background: #1e293b;
		border-radius: 12px;
		padding: 28px;
		border: 1px solid #334155;
		transition: border-color 0.2s;
	}
	.card:hover {
		border-color: #38bdf8;
	}
	h3 {
		color: #f8fafc;
		font-size: 1.3rem;
		margin-bottom: 12px;
	}
	.description {
		color: #94a3b8;
		margin-bottom: 16px;
		line-height: 1.6;
	}
	.tech {
		display: flex;
		flex-wrap: wrap;
		gap: 8px;
		list-style: none;
		padding: 0;
		margin-bottom: 16px;
	}
	.tech li {
		background: #0f172a;
		color: #38bdf8;
		padding: 4px 12px;
		border-radius: 4px;
		font-size: 0.85rem;
		font-weight: 500;
	}
	.links {
		display: flex;
		gap: 16px;
	}
	.links a {
		color: #38bdf8;
		font-weight: 500;
		font-size: 0.95rem;
	}
</style>
