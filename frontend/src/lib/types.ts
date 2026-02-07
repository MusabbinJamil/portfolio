export interface HeroData {
	name: string;
	title: string;
	tagline: string;
}

export interface AboutData {
	bio: string;
	skills: string[];
}

export interface Project {
	id: number;
	title: string;
	description: string;
	techStack: string[];
	liveUrl?: string;
	githubUrl?: string;
}

export interface ContactInfo {
	email: string;
	github: string;
	linkedin: string;
}
