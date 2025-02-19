package main

import (
	"log"
	"net/http"

	"github.com/furkankarayel/URL_Shortener/config"
	"github.com/furkankarayel/URL_Shortener/internal/api"
	"github.com/furkankarayel/URL_Shortener/internal/cache"
	"github.com/furkankarayel/URL_Shortener/internal/db"
	"github.com/furkankarayel/URL_Shortener/internal/ui"
	"github.com/furkankarayel/URL_Shortener/internal/urlshortener"
)

func main() {

	// Set up routes
	topLevelRoutes := make(map[string]*api.Route)

	config, cfgErr := config.NewConfig(".env")
	if cfgErr != nil {
		log.Fatal("Failed to load config", cfgErr)
	}

	db, dbErr := db.NewDB(config)
	if dbErr != nil {
		log.Fatal("Failed to connect to database", dbErr)
	}
	urlCache := cache.NewURLCache()

	topLevelRoutes[""] = ui.New()
	topLevelRoutes["url"] = urlshortener.New(db, urlCache)

	svr := api.New(topLevelRoutes)
	err := http.ListenAndServe(":8080", svr)
	log.Println(err)
	log.Println("Server is running")

}
