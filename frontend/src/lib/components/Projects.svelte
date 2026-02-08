<script lang="ts">
	import type { Project } from '$lib/types';
	import { trackClick } from '$lib/analytics';

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
						<a href={project.githubUrl} target="_blank" rel="noopener"
							onclick={() => trackClick(`project:${project.title}:github`, `${project.title} GitHub link`)}>GitHub</a>
					{/if}
					{#if project.liveUrl}
						<a class="explore-btn" href={project.liveUrl} target="_blank" rel="noopener"
							onclick={() => trackClick(`project:${project.title}:live`, `${project.title} Explore button`)}>Explore</a>
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
		font-size: 1.925rem;
		margin-bottom: 12px;
	}
	.description {
		color: #008f11;
		margin-bottom: 16px;
		line-height: 1.6;
		font-size: 1.575rem;
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
		font-size: 1.475rem;
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
		font-size: 1.575rem;
	}
	.explore-btn {
		background: #00ff41;
		color: #000000 !important;
		padding: 12px 36px;
		border-radius: 8px;
		text-decoration: none;
		font-weight: 600;
		font-size: 1.575rem;
		transition: background 0.2s, box-shadow 0.2s;
	}
	.explore-btn:hover {
		background: #33ff66;
		box-shadow: 0 0 15px rgba(0, 255, 65, 0.4);
	}

	@media (max-width: 768px) {
		.card {
			padding: 20px;
		}
		h3 {
			font-size: 1.3rem;
		}
		.description {
			font-size: 0.95rem;
		}
		.tech li {
			font-size: 0.85rem;
		}
		.links a {
			font-size: 0.95rem;
		}
		.explore-btn {
			font-size: 0.95rem;
			padding: 10px 28px;
		}
	}
</style>
