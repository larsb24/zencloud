package main

import (
	"fmt"
	"net/http"
	"os"

	"zencloud-backend/internal/handlers"
	"zencloud-backend/pkg/utils/environment"

	"github.com/go-chi/chi"
)

func main() {
	var port string = environment.GetPort()
	var host string = "localhost"
	var socket string = host + ":" + port

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting server on", socket)

	err := http.ListenAndServe(socket, r)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}

}
