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
		views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "home-page.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "contact-page.gohtml"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "faq-page.gohtml"))))

	r.NotFound(errorHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

	fmt.Println("HA GAAAY!!!")
}
