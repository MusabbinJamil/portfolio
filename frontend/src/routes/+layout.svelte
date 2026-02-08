<script lang="ts">
	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href="/logo.jpg" type="image/jpeg" />
</svelte:head>

{@render children()}
