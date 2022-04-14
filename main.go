package main

import (
	"fmt"
	"net/http"

	"github.com/SStoyanov22/lenslocked/controllers"
	"github.com/SStoyanov22/lenslocked/templates"
	"github.com/SStoyanov22/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "<h1>Page not Found</h1>", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	usersController := controllers.Users{}
	usersController.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	r.Get("/signup", usersController.New)

	r.Post("/signup", usersController.Create)

	r.NotFound(errorHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
