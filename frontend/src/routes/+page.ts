import type { PageLoad } from './$types';
import type { HeroData, AboutData, Project, Experience, Education, ContactInfo } from '$lib/types';

export const prerender = false;
export const ssr = false;

export const load: PageLoad = async ({ fetch }) => {
	const [heroRes, aboutRes, projectsRes, experienceRes, educationRes, contactRes] = await Promise.all([
		fetch('/api/hero'),
		fetch('/api/about'),
		fetch('/api/projects'),
		fetch('/api/experience'),
		fetch('/api/education'),
		fetch('/api/contact')
	]);

	const hero: HeroData = await heroRes.json();
	const about: AboutData = await aboutRes.json();
	const projects: Project[] = await projectsRes.json();
	const experience: Experience[] = await experienceRes.json();
	const education: Education[] = await educationRes.json();
	const contact: ContactInfo = await contactRes.json();

	return { hero, about, projects, experience, education, contact };
};
