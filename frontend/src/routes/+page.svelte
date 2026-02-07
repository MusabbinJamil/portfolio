<script lang="ts">
	import type { PageData } from './$types';
	import Navbar from '$lib/components/Navbar.svelte';
	import Hero from '$lib/components/Hero.svelte';
	import About from '$lib/components/About.svelte';
	import Projects from '$lib/components/Projects.svelte';
	import Contact from '$lib/components/Contact.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import '../app.css';

	let { data }: { data: PageData } = $props();
</script>

<Navbar />
<main>
	<Hero hero={data.hero} />
	<About about={data.about} />
	<Projects projects={data.projects} />
	<Contact contact={data.contact} />
</main>
<Footer />
