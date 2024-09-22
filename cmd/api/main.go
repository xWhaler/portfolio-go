package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

// BlogPost represents a blog post structure
type BlogPost struct {
	Title   string
	Content string
}

// In-memory storage for blog posts (this will reset each time the server restarts)
var blogPosts []BlogPost

// Template cache
var templates = template.Must(template.ParseFiles("templates/layout.html", "templates/content.html"))

// Home page handler - renders the blog template with layout
func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "layout.html", struct {
		Title string
		Posts []BlogPost
	}{
		Title: "Home - My Blog", // Dynamic title
		Posts: blogPosts,        // Pass the list of blog posts to the template
	})
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

// Blog submission handler (POST request)
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form values
	title := r.FormValue("title")
	content := r.FormValue("content")

	if title == "" || content == "" {
		http.Error(w, "Title or content cannot be empty", http.StatusBadRequest)
		return
	}

	// Create a new blog post and append it to the in-memory slice
	newPost := BlogPost{
		Title:   title,
		Content: content,
	}
	blogPosts = append(blogPosts, newPost)

	// Redirect to the home page to show the updated list of blog posts
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Projects page handler
func projectsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>My Projects</h1><p>Here is a list of my projects...</p>")
}

// CV page handler
func cvHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>My CV</h1><p>Here is my professional CV...</p>")
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", homeHandler).Methods("GET")             // Home page
	r.HandleFunc("/submit", submitHandler).Methods("POST")    // Blog post submission
	r.HandleFunc("/projects", projectsHandler).Methods("GET") // Projects page
	r.HandleFunc("/cv", cvHandler).Methods("GET")             // CV page

	// Start the server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
