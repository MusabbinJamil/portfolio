<script lang="ts">
	import type { ContactInfo } from '$lib/types';

	let { contact }: { contact: ContactInfo } = $props();

	let name = $state('');
	let email = $state('');
	let message = $state('');
	let status = $state<'idle' | 'sending' | 'sent' | 'error'>('idle');

	async function handleSubmit(e: SubmitEvent) {
		e.preventDefault();
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
	<div class="layout">
		<div class="info">
			<p>Get in touch or find me on:</p>
			<ul>
				<li><a href="mailto:{contact.email}">{contact.email}</a></li>
				<li><a href={contact.github} target="_blank" rel="noopener">GitHub</a></li>
				<li><a href={contact.linkedin} target="_blank" rel="noopener">LinkedIn</a></li>
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
</section>

<style>
	.layout {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 40px;
		align-items: start;
	}
	.info p {
		color: #94a3b8;
		margin-bottom: 16px;
	}
	.info ul {
		list-style: none;
		padding: 0;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}
	form {
		display: flex;
		flex-direction: column;
		gap: 12px;
	}
	input,
	textarea {
		background: #1e293b;
		border: 1px solid #334155;
		border-radius: 8px;
		padding: 12px 16px;
		color: #e2e8f0;
		font-family: inherit;
		font-size: 1rem;
	}
	input:focus,
	textarea:focus {
		outline: none;
		border-color: #38bdf8;
	}
	button {
		background: #38bdf8;
		color: #0f172a;
		border: none;
		border-radius: 8px;
		padding: 12px 24px;
		font-weight: 600;
		font-size: 1rem;
		cursor: pointer;
		transition: background 0.2s;
	}
	button:hover:not(:disabled) {
		background: #7dd3fc;
	}
	button:disabled {
		opacity: 0.6;
		cursor: not-allowed;
	}
	.success {
		color: #4ade80;
		font-weight: 500;
	}
	.error-msg {
		color: #f87171;
		font-weight: 500;
	}
</style>
