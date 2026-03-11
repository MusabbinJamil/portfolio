<script lang="ts">
	import '../../app.css';

	interface Message {
		timestamp: string;
		name: string;
		email: string;
		message: string;
	}

	let key = $state('');
	let messages: Message[] = $state([]);
	let total = $state(0);
	let error = $state('');
	let authenticated = $state(false);
	let loading = $state(false);

	async function login(e: SubmitEvent) {
		e.preventDefault();
		error = '';
		loading = true;

		try {
			const res = await fetch(`/api/messages?key=${encodeURIComponent(key)}`);
			if (!res.ok) {
				error = 'Invalid key';
				loading = false;
				return;
			}
			const data = await res.json();
			messages = (data.messages || []).reverse();
			total = data.total || 0;
			authenticated = true;
		} catch {
			error = 'Failed to connect';
		}
		loading = false;
	}

	function formatTime(ts: string): string {
		if (!ts) return '-';
		return new Date(ts).toLocaleString();
	}
</script>

<svelte:head>
	<title>Messages</title>
</svelte:head>

{#if !authenticated}
	<div class="login-container">
		<div class="login-box">
			<h1>messages</h1>
			<form onsubmit={login}>
				<input
					type="password"
					bind:value={key}
					placeholder="enter access key"
					autocomplete="off"
				/>
				<button type="submit" disabled={loading}>
					{loading ? 'connecting...' : 'access'}
				</button>
			</form>
			{#if error}
				<p class="error">{error}</p>
			{/if}
		</div>
	</div>
{:else}
	<div class="dashboard">
		<div class="header">
			<h1>messages</h1>
			<span class="total">{total} messages</span>
		</div>

		{#if messages.length === 0}
			<p class="empty">No messages yet.</p>
		{:else}
			<div class="message-list">
				{#each messages as msg}
					<div class="message-card">
						<div class="message-meta">
							<span class="name">{msg.name}</span>
							<a href="mailto:{msg.email}" class="email">{msg.email}</a>
							<span class="time">{formatTime(msg.timestamp)}</span>
						</div>
						<p class="message-body">{msg.message}</p>
					</div>
				{/each}
			</div>
		{/if}
	</div>
{/if}

<style>
	.login-container {
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.login-box {
		border: 1px solid #e8dff5;
		border-radius: 12px;
		padding: 48px;
		text-align: center;
		max-width: 400px;
		width: 100%;
		background: #ffffff;
		box-shadow: 0 4px 24px rgba(196, 167, 231, 0.12);
	}

	.login-box h1 {
		color: #907aa9;
		font-size: 1.8rem;
		margin-bottom: 32px;
	}

	input {
		width: 100%;
		padding: 12px 16px;
		background: #faf5ff;
		border: 1px solid #e8dff5;
		border-radius: 8px;
		color: #575279;
		font-family: 'Inter', monospace;
		font-size: 1rem;
		outline: none;
		margin-bottom: 16px;
	}

	input:focus {
		border-color: #c4a7e7;
	}

	input::placeholder {
		color: #9893a5;
	}

	button {
		width: 100%;
		padding: 12px;
		background: transparent;
		border: 1px solid #c4a7e7;
		border-radius: 8px;
		color: #907aa9;
		font-family: 'Inter', monospace;
		font-size: 1rem;
		cursor: pointer;
	}

	button:hover {
		background: rgba(196, 167, 231, 0.1);
	}

	button:disabled {
		opacity: 0.5;
		cursor: not-allowed;
	}

	.error {
		color: #eb6f92;
		margin-top: 16px;
		font-size: 0.9rem;
	}

	.dashboard {
		max-width: 800px;
		margin: 0 auto;
		padding: 40px 20px;
	}

	.header {
		display: flex;
		align-items: center;
		justify-content: space-between;
		margin-bottom: 32px;
	}

	.header h1 {
		color: #907aa9;
		font-size: 1.8rem;
	}

	.total {
		color: #9893a5;
		font-size: 1rem;
	}

	.empty {
		color: #9893a5;
		text-align: center;
		padding: 40px;
	}

	.message-list {
		display: flex;
		flex-direction: column;
		gap: 16px;
	}

	.message-card {
		border: 1px solid #e8dff5;
		border-radius: 8px;
		padding: 20px;
		background: #ffffff;
		box-shadow: 0 4px 24px rgba(196, 167, 231, 0.12);
	}

	.message-card:hover {
		border-color: #c4a7e7;
	}

	.message-meta {
		display: flex;
		align-items: center;
		gap: 16px;
		margin-bottom: 12px;
		flex-wrap: wrap;
	}

	.name {
		color: #907aa9;
		font-weight: 600;
		font-size: 1rem;
	}

	.email {
		color: #9893a5;
		font-size: 0.85rem;
	}

	.time {
		color: #c4a7e7;
		font-size: 0.8rem;
		margin-left: auto;
	}

	.message-body {
		color: #575279;
		font-size: 0.9rem;
		line-height: 1.6;
		white-space: pre-wrap;
	}

	@media (max-width: 768px) {
		.login-box {
			padding: 32px 24px;
			margin: 0 16px;
		}

		.dashboard {
			padding: 24px 12px;
		}

		.message-meta {
			gap: 8px;
		}

		.time {
			margin-left: 0;
			width: 100%;
		}
	}
</style>
