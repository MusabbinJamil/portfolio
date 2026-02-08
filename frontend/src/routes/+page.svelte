<script lang="ts">
	import type { PageData } from './$types';
	import Navbar from '$lib/components/Navbar.svelte';
	import Hero from '$lib/components/Hero.svelte';
	import About from '$lib/components/About.svelte';
	import Experience from '$lib/components/Experience.svelte';
	import Projects from '$lib/components/Projects.svelte';
	import Education from '$lib/components/Education.svelte';
	import Contact from '$lib/components/Contact.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import MatrixRain from '$lib/components/MatrixRain.svelte';
	import '../app.css';

	let { data }: { data: PageData } = $props();
</script>

<MatrixRain />
<Navbar />
<main>
	<Hero hero={data.hero} />
	<About about={data.about} />
	<Experience experience={data.experience} />
	<Projects projects={data.projects} />
	<Education education={data.education} />
	<Contact contact={data.contact} />
</main>
<Footer />
