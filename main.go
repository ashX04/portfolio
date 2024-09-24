package main

import (
	"html/template"
	"log"
	"net/http"
)

type Resume struct {
	Name            string
	Contact         string
	Summary         string
	ProfessionalExp []string
	Education       []string
	Skills          []string
}

func main() {
	// Define the resume data
	resume := Resume{
		Name:    "Your Name",
		Contact: "your.email@example.com | Your Phone Number | LinkedIn | GitHub",
		Summary: "I’m a passionate web developer with 3 years of experience creating user-friendly web experiences.",
		ProfessionalExp: []string{
			"1.Dreamers Inc.	founder|Apr 21|Present|",
			"I started the company as a game developer but transitioned to web development",
			"Unity,Unreal,3d art,music etc. for games I was developing",
			"Golang for beautiful cli tools and websites",
			"Elixir Phoenix framework,Gleam for massively concurrent webapps hosted on home server using coolify",
			"Linux administration to make my server secure",
			"Automated repetitive tasks by writing backend scripts, improving software usability and efficiency.",
			"Python for machine learning models and api",
			"2.StareOut Studio	Developer|Feb 19|Apr 19|",
			"Actively contributed to the development of responsive games using modern frameworks.",
			"Conducted testing and debugging to ensure smooth user experiences with benchmarks for user retention rates.",
			"Took part in many Brainstorming Sessions for new games being developed by the studio",
			"Used C#,Python,Blender Script",
			"3.UltraVirtualStudios	Developer|Jan 18|Dec 18|",
			"I Collaborated within an international team to develop a space horror-themed project.",
			"Collaborated with design and backend teams to ensure a cohesive and seamless user experience.",
			"Created 3d Art of monsters for World building,Marketing etc.",
		},
		Education: []string{
			"Degree in Field of Study - University Name",
		},
		Skills: []string{
			"HTML, CSS, JavaScript",
			"Golang, Elixir",
			"React, Phoenix",
			"Tailwind CSS",
		},
	}

	// Load templates
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from: %s", r.RemoteAddr)
		tmpl.Execute(w, resume)
	})

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}

}
