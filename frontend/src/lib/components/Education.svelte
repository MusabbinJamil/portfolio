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
		background: #ffffff;
		border-radius: 12px;
		padding: 28px;
		border: 1px solid #d4edda;
		box-shadow: 0 4px 24px rgba(123, 200, 155, 0.12);
		transition: border-color 0.2s;
	}
	.edu-card:hover {
		border-color: #7bc89b;
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
		color: #4a8c64;
		font-size: 1.825rem;
	}
	.year {
		color: #7bc89b;
		font-size: 1.525rem;
		font-weight: 500;
	}
	.institution {
		color: #3d5a47;
		font-size: 1.675rem;
		font-weight: 500;
		margin-bottom: 8px;
	}
	.details {
		color: #5a7a64;
		line-height: 1.7;
		font-size: 1.525rem;
	}

	@media (max-width: 768px) {
		.edu-card {
			padding: 20px;
		}
		h3 {
			font-size: 1.2rem;
		}
		.year {
			font-size: 0.9rem;
		}
		.institution {
			font-size: 1.05rem;
		}
		.details {
			font-size: 0.9rem;
		}
	}
</style>
