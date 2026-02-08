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

type Experience struct {
	ID          int    `json:"id"`
	Role        string `json:"role"`
	Company     string `json:"company"`
	Period      string `json:"period"`
	Description string `json:"description"`
}

type Education struct {
	ID          int    `json:"id"`
	Degree      string `json:"degree"`
	Institution string `json:"institution"`
	Year        string `json:"year"`
	Details     string `json:"details,omitempty"`
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
	Hero       HeroData
	About      AboutData
	Projects   []Project
	Experience []Experience
	Education  []Education
	Contact    ContactInfo

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
			Name:    "Musab Bin Jamil",
			Title:   "Full Stack Developer & Data Engineer",
			Tagline: "Building robust web applications, data pipelines, and AI-driven solutions.",
		},
		About: AboutData{
			Bio: "Full stack developer with hands-on experience in web development, " +
				"data engineering, and analysis. MS Data Sciences graduate from IBA Karachi, " +
				"currently working as a Junior Full Stack Developer at SmartB Solutions. " +
				"I've built everything from Django backends and ecommerce integrations " +
				"to ETL pipelines and LLM-powered tools. Fast learner, tech savvy, and always eager to improve.",
			Skills: []string{
				"Python", "JavaScript", "C++", "SQL", "HTML/CSS", "Go",
				"Django", "Flask", "React.js", "Node.js", "Svelte",
				"PostgreSQL", "MongoDB", "Docker", "Git", "Linux", "REST APIs",
				"Selenium", "Airflow", "Scikit-learn", "Power BI", "Streamlit",
			},
		},
		Projects: []Project{
			{
				ID:    1,
				Title: "FastMig",
				Description: "A data migration platform where users can clean text-based data " +
					"using evolutionary computation and LLMs. Includes ETL workflow, " +
					"DB connectivity, and other data transformation utilities.",
				TechStack: []string{"Python", "LLMs", "Evolutionary Computation", "ETL"},
				LiveURL:   "http://fastmig.unmashable.online",
			},
			{
				ID:    2,
				Title: "ETL Pipeline using Airflow",
				Description: "A pipeline to extract, transform, and load data from a transactional " +
					"database into a star schema for data analysis and insightful dashboarding.",
				TechStack: []string{"Python", "Apache Airflow", "SQL", "Star Schema"},
			},
			{
				ID:    3,
				Title: "LLM Based OCR Improvement",
				Description: "A prototype that fills in the gaps of Azure Computer Vision OCR " +
					"using semantic imputation of text via the OpenAI API.",
				TechStack: []string{"Python", "OpenAI API", "Azure Computer Vision", "OCR"},
			},
			{
				ID:    4,
				Title: "Persona Based Decision Making",
				Description: "An interface to load data and get feedback based on multiple " +
					"user-defined personas, with verification on LLM answers using probabilistic reasoning.",
				TechStack: []string{"Python", "LLMs", "Probabilistic Reasoning"},
			},
			{
				ID:          5,
				Title:       "Trike AI",
				Description: "An AI agent that can play the game of Trike using reinforcement learning.",
				TechStack:   []string{"Python", "Reinforcement Learning", "AI"},
			},
			{
				ID:          6,
				Title:       "Flask Dashboard",
				Description: "A Flask-based interactive dashboard for data visualization and analytics.",
				TechStack:   []string{"Python", "Flask", "HTML/CSS", "JavaScript"},
				LiveURL:     "https://musabjamil.pythonanywhere.com",
			},
		},
		Experience: []Experience{
			{
				ID:      1,
				Role:    "Junior Full Stack Developer",
				Company: "SmartB Solutions Sdn Bhd",
				Period:  "2024 – Present",
				Description: "Iterative improvements to SmartB's Accounts and Stock tracking workflows. " +
					"Ecommerce integration for Shopee, Lazada, TikTok Shop, Shopify, DHL, SPX Express and NinjaVan. " +
					"Integration of OpenAI for chat-driven sales order generation. " +
					"Bayesian decision support system development. Server-side support and client configuration on custom domains.",
			},
			{
				ID:      2,
				Role:    "Python Developer",
				Company: "Elite Chain L.L.C.",
				Period:  "2024",
				Description: "Sole developer on a Django-based backend project. Built APIs for a Rasa chatbot, " +
					"implemented Celery-based automation, and developed JWT-based user authentication.",
			},
			{
				ID:      3,
				Role:    "Python Developer",
				Company: "Nixaam L.L.C.",
				Period:  "2023 – 2024",
				Description: "Delivered diverse solutions including image processing, HRM systems, " +
					"database design, automation, and web scraping. Contributed to innovation and standardization across projects.",
			},
			{
				ID:      4,
				Role:    "Research Assistant",
				Company: "Karachi Urban Lab",
				Period:  "2023",
				Description: "Cleaned and consolidated datasets using Python, Excel, and Oracle SQL. " +
					"Applied machine learning (Ridge Regression) to handle missing data. " +
					"Managed data storage and reporting via Google Drive and Sheets.",
			},
			{
				ID:      5,
				Role:    "Junior Developer",
				Company: "Wix Digitals",
				Period:  "2023",
				Description: "Developed algorithms on cryptocurrency data. Created interactive dashboards " +
					"using Streamlit and automated scraping bots with Selenium. Utilized Git for version control.",
			},
		},
		Education: []Education{
			{
				ID:          1,
				Degree:      "MS Data Sciences",
				Institution: "Institute of Business Administration (IBA), Karachi",
				Year:        "2026",
				Details:     "Core Courses: Data Analytics & Warehousing, Machine Learning I & II, Deep Learning, Computational Intelligence, Probabilistic Reasoning, DevOps",
			},
			{
				ID:          2,
				Degree:      "M.Sc. Space Science and Technology (With Thesis)",
				Institution: "University of Karachi",
				Year:        "2021",
				Details:     "Focus Areas: Quantum Mechanics, Gravitational Physics, GIS, Astrophysics, Atmospheric Sciences",
			},
			{
				ID:          3,
				Degree:      "B.Sc. (Hons) Space Science and Technology",
				Institution: "University of Karachi",
				Year:        "2020",
			},
		},
		Contact: ContactInfo{
			Email:    "musabjamil3@gmail.com",
			Github:   "https://github.com/MusabbinJamil",
			LinkedIn: "https://www.linkedin.com/in/musab-b-630256b4/",
		},
		Messages: []ContactMessage{},
	}
}
