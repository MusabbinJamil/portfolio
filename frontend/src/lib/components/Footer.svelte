<footer>
	<p>&copy; {new Date().getFullYear()} Musab. Built with SvelteKit &amp; Go.</p>
</footer>

<style>
	footer {
		text-align: center;
		padding: 40px 20px;
		color: var(--text-muted);
		border-top: 1px solid var(--border);
		font-size: clamp(0.9rem, 1.3vw, 1.525rem);
		transition: color 0.3s, border-color 0.3s;
	}

	@media (max-width: 768px) {
		footer {
			padding: 24px 16px;
		}
	}
</style>
