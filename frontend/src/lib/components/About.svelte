<script lang="ts">
	import type { AboutData } from '$lib/types';

	let { about }: { about: AboutData } = $props();
</script>

<section id="about">
	<h2>About Me</h2>
	<div class="card">
		<div class="content">
			<div>
				<p class="bio">{about.bio}</p>
			</div>
			<div>
				<p class="skills-title">Technologies I work with:</p>
				<ul class="skills-list">
					{#each about.skills as skill}
						<li class="skill">
							<span class="bullet">&#9654;</span>
							{skill}
						</li>
					{/each}
				</ul>
			</div>
		</div>
	</div>
</section>

<style>
	.card {
		background: #0a0a0a;
		border: 1px solid #003300;
		border-radius: 16px;
		padding: 40px;
	}
	.content {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 40px;
		align-items: start;
	}
	.bio {
		color: #008f11;
		font-size: 1.675rem;
		line-height: 1.8;
	}
	.skills-title {
		color: #00ff41;
		font-size: 1.725rem;
		font-weight: 600;
		margin-bottom: 16px;
	}
	.skills-list {
		display: grid;
		grid-template-columns: 1fr 1fr 1fr;
		gap: 8px;
		list-style: none;
		padding: 0;
	}
	.skill {
		color: #008f11;
		font-size: 1.05rem;
		padding: 6px 0;
	}
	.bullet {
		color: #00ff41;
		margin-right: 8px;
	}
</style>
