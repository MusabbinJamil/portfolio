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
	<div class="card">
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
	</div>
</section>

<style>
	.card {
		background: #0a0a0a;
		border: 1px solid #003300;
		border-radius: 16px;
		padding: 40px;
	}
	.layout {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 40px;
		align-items: start;
	}
	.info p {
		color: #008f11;
		margin-bottom: 16px;
		font-size: 1.625rem;
	}
	.info ul {
		list-style: none;
		padding: 0;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}
	.info ul a {
		font-size: 1.575rem;
	}
	form {
		display: flex;
		flex-direction: column;
		gap: 12px;
	}
	input,
	textarea {
		background: #0a0a0a;
		border: 1px solid #003300;
		border-radius: 8px;
		padding: 14px 16px;
		color: #00cc33;
		font-family: inherit;
		font-size: 1.625rem;
	}
	input:focus,
	textarea:focus {
		outline: none;
		border-color: #00ff41;
		box-shadow: 0 0 5px rgba(0, 255, 65, 0.3);
	}
	button {
		background: #00ff41;
		color: #000000;
		border: none;
		border-radius: 8px;
		padding: 14px 24px;
		font-weight: 600;
		font-size: 1.625rem;
		cursor: pointer;
		transition: background 0.2s, box-shadow 0.2s;
	}
	button:hover:not(:disabled) {
		background: #33ff66;
		box-shadow: 0 0 15px rgba(0, 255, 65, 0.4);
	}
	button:disabled {
		opacity: 0.6;
		cursor: not-allowed;
	}
	.success {
		color: #00ff41;
		font-weight: 500;
		font-size: 1.575rem;
	}
	.error-msg {
		color: #ff3333;
		font-weight: 500;
		font-size: 1.575rem;
	}

	@media (max-width: 768px) {
		.card {
			padding: 24px 20px;
		}
		.layout {
			grid-template-columns: 1fr;
			gap: 24px;
		}
		.info p {
			font-size: 1rem;
		}
		.info ul a {
			font-size: 1rem;
		}
		input,
		textarea {
			font-size: 1rem;
		}
		button {
			font-size: 1rem;
		}
		.success,
		.error-msg {
			font-size: 0.95rem;
		}
	}
</style>
