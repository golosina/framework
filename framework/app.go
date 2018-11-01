package framework

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

type App struct {
	Router *Router
}

func New() *App {

	loadEnv()

	r := NewRouter()
	return &App{r}
}

func (a *App) Start() {
	// Bind to a port and pass our router in
	log.Println("Server listening at localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
