import type { PageServerLoad } from './$types';
import type { HeroData, AboutData, Project, ContactInfo } from '$lib/types';

const API_BASE = process.env.API_URL || 'http://localhost:3000';

export const load: PageServerLoad = async ({ fetch }) => {
	const [heroRes, aboutRes, projectsRes, contactRes] = await Promise.all([
		fetch(`${API_BASE}/api/hero`),
		fetch(`${API_BASE}/api/about`),
		fetch(`${API_BASE}/api/projects`),
		fetch(`${API_BASE}/api/contact`)
	]);

	const hero: HeroData = await heroRes.json();
	const about: AboutData = await aboutRes.json();
	const projects: Project[] = await projectsRes.json();
	const contact: ContactInfo = await contactRes.json();

	return { hero, about, projects, contact };
};
