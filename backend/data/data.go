package data

import "sync"

type HeroData struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Tagline string `json:"tagline"`
}

type AboutData struct {
	Bio    string   `json:"bio"`
	Skills []string `json:"skills"`
}

type Project struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	TechStack   []string `json:"techStack"`
	LiveURL     string   `json:"liveUrl,omitempty"`
	GithubURL   string   `json:"githubUrl,omitempty"`
}

type ContactInfo struct {
	Email    string `json:"email"`
	Github   string `json:"github"`
	LinkedIn string `json:"linkedin"`
}

type ContactMessage struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Store struct {
	Hero     HeroData
	About    AboutData
	Projects []Project
	Contact  ContactInfo

	mu       sync.Mutex
	Messages []ContactMessage
}

func (s *Store) AddMessage(msg ContactMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Messages = append(s.Messages, msg)
}

func NewStore() *Store {
	return &Store{
		Hero: HeroData{
			Name:    "Musab",
			Title:   "Full-Stack Developer",
			Tagline: "I build things for the web with clean code and sharp design.",
		},
		About: AboutData{
			Bio: "I'm a developer passionate about creating clean, efficient, " +
				"and user-friendly web applications. I enjoy working across the " +
				"full stack, from designing intuitive interfaces to building " +
				"robust backend services.",
			Skills: []string{
				"Go", "JavaScript", "TypeScript", "Svelte",
				"React", "Node.js", "PostgreSQL", "Linux",
				"Docker", "Git",
			},
		},
		Projects: []Project{
			{
				ID:          1,
				Title:       "Portfolio Site",
				Description: "This very site — built with SvelteKit and Go.",
				TechStack:   []string{"SvelteKit", "Go", "CSS"},
				GithubURL:   "https://github.com/musab/portfolio",
			},
			{
				ID:          2,
				Title:       "Task Tracker",
				Description: "A minimal task management app with drag-and-drop.",
				TechStack:   []string{"React", "Node.js", "PostgreSQL"},
				LiveURL:     "https://tasks.example.com",
				GithubURL:   "https://github.com/musab/task-tracker",
			},
			{
				ID:          3,
				Title:       "CLI Weather Tool",
				Description: "A command-line tool that fetches weather data.",
				TechStack:   []string{"Go", "OpenWeather API"},
				GithubURL:   "https://github.com/musab/weather-cli",
			},
		},
		Contact: ContactInfo{
			Email:    "musab@example.com",
			Github:   "https://github.com/musab",
			LinkedIn: "https://linkedin.com/in/musab",
		},
		Messages: []ContactMessage{},
	}
}
