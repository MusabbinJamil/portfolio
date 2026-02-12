<script lang="ts">
	import '../../app.css';

	interface AnalyticsEvent {
		timestamp: string;
		ip: string;
		userAgent: string;
		eventType: string;
		target: string;
		label: string;
		referrer: string;
		country: string;
		city: string;
	}

	let key = $state('');
	let events: AnalyticsEvent[] = $state([]);
	let total = $state(0);
	let error = $state('');
	let authenticated = $state(false);
	let loading = $state(false);

	let pageviews = $derived(events.filter(e => e.eventType === 'pageview').length);
	let clicks = $derived(events.filter(e => e.eventType === 'click').length);
	let countries = $derived([...new Set(events.map(e => e.country).filter(Boolean))]);

	async function login(e: SubmitEvent) {
		e.preventDefault();
		error = '';
		loading = true;

		try {
			const res = await fetch(`/api/analytics?key=${encodeURIComponent(key)}`);
			if (!res.ok) {
				error = 'Invalid key';
				loading = false;
				return;
			}
			const data = await res.json();
			events = (data.events || []).reverse();
			total = data.total || 0;
			authenticated = true;
		} catch {
			error = 'Failed to connect';
		}
		loading = false;
	}

	function formatTime(ts: string): string {
		return new Date(ts).toLocaleString();
	}

	function parseUA(ua: string): string {
		if (!ua) return '-';
		if (ua.includes('curl')) return 'curl';
		const browser = ua.match(/(Chrome|Firefox|Safari|Edg|Opera)\/[\d.]+/);
		const os = ua.match(/\((.*?)\)/);
		const parts = [];
		if (browser) parts.push(browser[1]);
		if (os) {
			const osStr = os[1].split(';')[0].trim();
			parts.push(osStr);
		}
		return parts.join(' / ') || ua.slice(0, 40);
	}
</script>

<svelte:head>
	<title>Analytics</title>
</svelte:head>

{#if !authenticated}
	<div class="login-container">
		<div class="login-box">
			<h1>&gt; analytics_</h1>
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
			<h1>&gt; analytics_</h1>
			<span class="total">{total} events</span>
		</div>

		<div class="stats">
			<div class="stat-card">
				<div class="stat-value">{pageviews}</div>
				<div class="stat-label">pageviews</div>
			</div>
			<div class="stat-card">
				<div class="stat-value">{clicks}</div>
				<div class="stat-label">clicks</div>
			</div>
			<div class="stat-card">
				<div class="stat-value">{countries.length}</div>
				<div class="stat-label">countries</div>
			</div>
		</div>

		<div class="table-wrapper">
			<table>
				<thead>
					<tr>
						<th>time</th>
						<th>country</th>
						<th>city</th>
						<th>type</th>
						<th>target</th>
						<th>browser</th>
					</tr>
				</thead>
				<tbody>
					{#each events as event}
						<tr>
							<td>{formatTime(event.timestamp)}</td>
							<td>{event.country || '-'}</td>
							<td>{event.city || '-'}</td>
							<td><span class="badge {event.eventType}">{event.eventType}</span></td>
							<td>{event.target || event.label || '-'}</td>
							<td class="ua">{parseUA(event.userAgent)}</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
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
		border: 1px solid #003300;
		border-radius: 12px;
		padding: 48px;
		text-align: center;
		max-width: 400px;
		width: 100%;
	}

	.login-box h1 {
		color: #00ff41;
		font-size: 1.8rem;
		margin-bottom: 32px;
	}

	input {
		width: 100%;
		padding: 12px 16px;
		background: #0a0a0a;
		border: 1px solid #003300;
		border-radius: 8px;
		color: #00ff41;
		font-family: 'Inter', monospace;
		font-size: 1rem;
		outline: none;
		margin-bottom: 16px;
	}

	input:focus {
		border-color: #00ff41;
	}

	input::placeholder {
		color: #008f11;
	}

	button {
		width: 100%;
		padding: 12px;
		background: transparent;
		border: 1px solid #00ff41;
		border-radius: 8px;
		color: #00ff41;
		font-family: 'Inter', monospace;
		font-size: 1rem;
		cursor: pointer;
	}

	button:hover {
		background: #00ff4115;
	}

	button:disabled {
		opacity: 0.5;
		cursor: not-allowed;
	}

	.error {
		color: #ff4141;
		margin-top: 16px;
		font-size: 0.9rem;
	}

	.dashboard {
		max-width: 1100px;
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
		color: #00ff41;
		font-size: 1.8rem;
	}

	.total {
		color: #008f11;
		font-size: 1rem;
	}

	.stats {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		gap: 16px;
		margin-bottom: 32px;
	}

	.stat-card {
		border: 1px solid #003300;
		border-radius: 8px;
		padding: 20px;
		text-align: center;
	}

	.stat-value {
		color: #00ff41;
		font-size: 2rem;
		font-weight: 700;
	}

	.stat-label {
		color: #008f11;
		font-size: 0.85rem;
		margin-top: 4px;
	}

	.table-wrapper {
		overflow-x: auto;
	}

	table {
		width: 100%;
		border-collapse: collapse;
		font-size: 0.85rem;
	}

	th {
		text-align: left;
		padding: 12px 16px;
		color: #00ff41;
		border-bottom: 1px solid #003300;
		white-space: nowrap;
		font-weight: 600;
	}

	td {
		padding: 10px 16px;
		border-bottom: 1px solid #0a1a0a;
		color: #00cc33;
		white-space: nowrap;
	}

	tr:hover td {
		background: #001a00;
	}

	.badge {
		display: inline-block;
		padding: 2px 8px;
		border-radius: 4px;
		font-size: 0.75rem;
	}

	.badge.pageview {
		background: #00ff4120;
		color: #00ff41;
	}

	.badge.click {
		background: #41ff0020;
		color: #41ff00;
	}

	.ua {
		max-width: 200px;
		overflow: hidden;
		text-overflow: ellipsis;
	}

	@media (max-width: 768px) {
		.login-box {
			padding: 32px 24px;
			margin: 0 16px;
		}

		.stats {
			grid-template-columns: 1fr;
		}

		.dashboard {
			padding: 24px 12px;
		}

		table {
			font-size: 0.75rem;
		}

		th, td {
			padding: 8px 10px;
		}
	}
</style>
