import type { PageLoad } from './$types';
import type { HeroData, AboutData, Project, ContactInfo } from '$lib/types';

export const prerender = false;
export const ssr = false;

export const load: PageLoad = async ({ fetch }) => {
	const [heroRes, aboutRes, projectsRes, contactRes] = await Promise.all([
		fetch('/api/hero'),
		fetch('/api/about'),
		fetch('/api/projects'),
		fetch('/api/contact')
	]);

	const hero: HeroData = await heroRes.json();
	const about: AboutData = await aboutRes.json();
	const projects: Project[] = await projectsRes.json();
	const contact: ContactInfo = await contactRes.json();

	return { hero, about, projects, contact };
};
