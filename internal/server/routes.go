package server

import (
	"encoding/json"
	"log"
	"farmaIA/internal/services"
	"net/http"
    "farmaIA/internal/healthcheck"
    "farmaIA/internal/helloworld" 
    "farmaIA/internal/database"
    "farmaIA/cmd/api/swagger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
    
    service := database.New()

    healthHandler := healthcheck.NewHandler(service)
    healthHandler.RegisterHandlers(r)
    
// @Tags HelloWorld
// @Produce json
// @Success 200 {object} map[string]string
// @Router /helloworld [get]
    helloWorldHandler := helloworld.NewHandler(service)
    helloWorldHandler.RegisterHandlers(r)
    
    swagger.SwaggerHandler(r)

    r.Get("/survey", s.getSurveyHandler)
    r.Get("/gemini", s.geminiHandler)

	return r
}

// @Tags Survey
// @Produce json
// @Success 200 {object} []types.Survey
// @Router /survey [get]
func (s *Server) getSurveyHandler (w http.ResponseWriter, r *http.Request) {
    result, err := s.db.GetSurveys()
    if err != nil {
        log.Fatalf("Error on getting surveys. Err: %v", err)
    }

	jsonResp, _ := json.Marshal(result)
	_, _ = w.Write(jsonResp)
}

func (s *Server) geminiHandler (w http.ResponseWriter, r *http.Request) {
    services.Gemini()
}

