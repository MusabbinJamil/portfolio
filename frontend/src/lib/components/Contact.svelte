<script lang="ts">
	import type { ContactInfo } from '$lib/types';
	import { trackClick } from '$lib/analytics';

	let { contact }: { contact: ContactInfo } = $props();

	let name = $state('');
	let email = $state('');
	let message = $state('');
	let status = $state<'idle' | 'sending' | 'sent' | 'error'>('idle');

	async function handleSubmit(e: SubmitEvent) {
		e.preventDefault();
		trackClick('contact:submit', 'Contact form submit');
		status = 'sending';

		try {
			const res = await fetch('/api/contact', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ name, email, message })
			});

			if (!res.ok) throw new Error('Failed to send');

			status = 'sent';
			name = '';
			email = '';
			message = '';
		} catch {
			status = 'error';
		}
	}
</script>

<section id="contact">
	<h2>Contact</h2>
	<div class="card">
		<div class="layout">
			<div class="info">
				<p>Get in touch or find me on:</p>
				<ul>
					<li><a href="mailto:{contact.email}" onclick={() => trackClick('contact:email', 'Email link')}>{contact.email}</a></li>
					<li><a href={contact.github} target="_blank" rel="noopener" onclick={() => trackClick('contact:github', 'GitHub link')}>GitHub</a></li>
					<li><a href={contact.linkedin} target="_blank" rel="noopener" onclick={() => trackClick('contact:linkedin', 'LinkedIn link')}>LinkedIn</a></li>
				</ul>
			</div>
			<form onsubmit={handleSubmit}>
				<input bind:value={name} placeholder="Name" required />
				<input bind:value={email} type="email" placeholder="Email" required />
				<textarea bind:value={message} placeholder="Message" rows="5" required></textarea>
				<button type="submit" disabled={status === 'sending'}>
					{status === 'sending' ? 'Sending...' : 'Send Message'}
				</button>
				{#if status === 'sent'}
					<p class="success">Message sent! I'll get back to you soon.</p>
				{/if}
				{#if status === 'error'}
					<p class="error-msg">Something went wrong. Please try again.</p>
				{/if}
			</form>
		</div>
	</div>
</section>

<style>
	.card {
		background: var(--bg-card);
		border: 1px solid var(--border);
		border-radius: 16px;
		padding: 40px;
		box-shadow: 0 4px 24px var(--shadow);
		transition: background 0.3s, border-color 0.3s;
	}
	.layout {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 40px;
		align-items: start;
	}
	.info p {
		color: var(--text-secondary);
		margin-bottom: 16px;
		font-size: clamp(1rem, 1.5vw, 1.625rem);
	}
	.info ul {
		list-style: none;
		padding: 0;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}
	.info ul a {
		font-size: clamp(0.95rem, 1.4vw, 1.575rem);
	}
	form {
		display: flex;
		flex-direction: column;
		gap: 12px;
	}
	input,
	textarea {
		background: var(--bg-input);
		border: 1px solid var(--border);
		border-radius: 8px;
		padding: 14px 16px;
		color: var(--text-primary);
		font-family: inherit;
		font-size: clamp(1rem, 1.5vw, 1.625rem);
		transition: background 0.3s, border-color 0.3s, color 0.3s;
	}
	input:focus,
	textarea:focus {
		outline: none;
		border-color: var(--accent-soft);
		box-shadow: 0 0 8px var(--shadow-hover);
	}
	input::placeholder,
	textarea::placeholder {
		color: var(--text-muted);
	}
	button {
		background: var(--accent-soft);
		color: #ffffff;
		border: none;
		border-radius: 8px;
		padding: 14px 24px;
		font-weight: 600;
		font-size: clamp(1rem, 1.5vw, 1.625rem);
		cursor: pointer;
		transition: background 0.2s, box-shadow 0.2s;
	}
	button:hover:not(:disabled) {
		background: var(--accent-hover);
		box-shadow: 0 4px 16px var(--shadow-hover);
	}
	button:disabled {
		opacity: 0.6;
		cursor: not-allowed;
	}
	.success {
		color: var(--success);
		font-weight: 500;
		font-size: clamp(0.95rem, 1.4vw, 1.575rem);
	}
	.error-msg {
		color: var(--error);
		font-weight: 500;
		font-size: clamp(0.95rem, 1.4vw, 1.575rem);
	}

	@media (max-width: 768px) {
		.card {
			padding: 24px 20px;
		}
		.layout {
			grid-template-columns: 1fr;
			gap: 24px;
		}
	}
</style>
