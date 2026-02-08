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

export interface Experience {
	id: number;
	role: string;
	company: string;
	period: string;
	description: string;
}

export interface Education {
	id: number;
	degree: string;
	institution: string;
	year: string;
	details?: string;
}

export interface ContactInfo {
	email: string;
	github: string;
	linkedin: string;
}
