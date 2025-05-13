package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	//"github.com/avukadin/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/files", func(router chi.Router) {

		// Middleware for authentication in this route
		//router.Use(middleware.Authorization)

		router.Post("/upload", UploadHandler)
		router.Get("/download", DownloadHandler)
	})

	r.Get("/version", versionHandler)
}
