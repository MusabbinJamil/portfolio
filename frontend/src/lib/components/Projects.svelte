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
					<a class="explore-btn" href={project.liveUrl || project.githubUrl} target="_blank" rel="noopener"
						onclick={() => trackClick(`project:${project.title}:explore`, `${project.title} Explore button`)}>Explore</a>
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
		background: var(--bg-card);
		border-radius: 12px;
		padding: 28px;
		border: 1px solid var(--border);
		box-shadow: 0 4px 24px var(--shadow);
		transition: border-color 0.2s, background 0.3s;
	}
	.card:hover {
		border-color: var(--accent-soft);
	}
	h3 {
		color: var(--accent);
		font-size: clamp(1.3rem, 2vw, 1.925rem);
		margin-bottom: 12px;
	}
	.description {
		color: var(--text-secondary);
		margin-bottom: 16px;
		line-height: 1.6;
		font-size: clamp(0.95rem, 1.4vw, 1.575rem);
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
		background: var(--bg-badge);
		color: var(--accent);
		padding: 4px 12px;
		border-radius: 4px;
		font-size: clamp(0.85rem, 1.2vw, 1.475rem);
		font-weight: 500;
		border: 1px solid var(--border);
		transition: background 0.3s, border-color 0.3s;
	}
	.links {
		display: flex;
		justify-content: center;
		gap: 16px;
	}
	.links a {
		color: var(--accent);
		font-weight: 500;
		font-size: clamp(0.95rem, 1.4vw, 1.575rem);
	}
	.explore-btn {
		background: var(--accent-soft);
		color: #ffffff !important;
		padding: 12px 36px;
		border-radius: 8px;
		text-decoration: none;
		font-weight: 600;
		font-size: clamp(0.95rem, 1.4vw, 1.575rem);
		transition: background 0.2s, box-shadow 0.2s;
	}
	.explore-btn:hover {
		background: var(--accent-hover);
		box-shadow: 0 4px 16px var(--shadow-hover);
	}

	@media (max-width: 768px) {
		.card {
			padding: 20px;
		}
	}
</style>
