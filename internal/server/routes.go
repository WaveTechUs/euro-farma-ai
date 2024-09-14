package server

import (
	"farmaIA/cmd/api/swagger"
	"farmaIA/internal/database"
	"farmaIA/internal/gemini"
	"farmaIA/internal/healthcheck"
	"farmaIA/internal/helloworld"
	"farmaIA/internal/survey"
	"farmaIA/internal/user"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
    r.Use(cors.Handler( cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST","PUT","DELETE","OPTIONS"},
    AllowedHeaders: []string{"Content-type"},
    }))
    
    service := database.New()


    healthHandler := healthcheck.NewHandler(service)
    healthHandler.RegisterHandlers(r)
    
    userHandler := user.NewHandler(service)
    userHandler.RegisterHandlers(r)

// @Tags Survey
// @Produce json
// @Success 200 {object} map[string]string
// @Router /survey [get]
    surveyHandler := survey.NewHandler(service)
    surveyHandler.RegisterHandlers(r)

// @Tags HelloWorld
// @Produce json
// @Success 200 {object} map[string]string
// @Router /helloworld [get]
    helloWorldHandler := helloworld.NewHandler(service)
    helloWorldHandler.RegisterHandlers(r)

    geminiHandler := gemini.NewHandler(service)
    geminiHandler.RegisterRoutes(r)
    
    swagger.SwaggerHandler(r)

	return r
}
