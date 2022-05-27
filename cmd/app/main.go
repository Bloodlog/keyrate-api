package main

import (
	"fmt"
	"key-rate-api/config"
	"key-rate-api/docs"
	"key-rate-api/internal/app/router"
	"log"
	"net/http"
)

func main() {
	appConf := config.AppConfig()
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	log.Printf("Starting server %s\n", address)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := router.SetUpRouter()

	s := &http.Server{
		Addr:         address,
		Handler:      r,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	s.ListenAndServe()
}
