<footer>
	<p>&copy; {new Date().getFullYear()} Musab. Built with SvelteKit &amp; Go.</p>
</footer>

<style>
	footer {
		text-align: center;
		padding: 40px 20px;
		color: var(--text-muted);
		border-top: 1px solid var(--border);
		font-size: 1.525rem;
		transition: color 0.3s, border-color 0.3s;
	}

	@media (max-width: 768px) {
		footer {
			font-size: 0.9rem;
			padding: 24px 16px;
		}
	}
</style>
