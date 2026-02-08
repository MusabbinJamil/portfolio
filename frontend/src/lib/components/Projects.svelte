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
						<a class="explore-btn" href={project.liveUrl} target="_blank" rel="noopener">Explore</a>
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
		background: #0a0a0a;
		border-radius: 12px;
		padding: 28px;
		border: 1px solid #003300;
		transition: border-color 0.2s;
	}
	.card:hover {
		border-color: #00ff41;
	}
	h3 {
		color: #00ff41;
		font-size: 1.3rem;
		margin-bottom: 12px;
	}
	.description {
		color: #008f11;
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
		background: #000000;
		color: #00ff41;
		padding: 4px 12px;
		border-radius: 4px;
		font-size: 0.85rem;
		font-weight: 500;
		border: 1px solid #003300;
	}
	.links {
		display: flex;
		justify-content: center;
		gap: 16px;
	}
	.links a {
		color: #00ff41;
		font-weight: 500;
		font-size: 0.95rem;
	}
	.explore-btn {
		background: #00ff41;
		color: #000000 !important;
		padding: 10px 32px;
		border-radius: 8px;
		text-decoration: none;
		font-weight: 600;
		font-size: 0.95rem;
		transition: background 0.2s, box-shadow 0.2s;
	}
	.explore-btn:hover {
		background: #33ff66;
		box-shadow: 0 0 15px rgba(0, 255, 65, 0.4);
	}
</style>
