package server

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "farmaIA/cmd/api/swagger"
	"farmaIA/internal/services"

=======
	"net/http"
    "farmaIA/internal/healthcheck"
    "farmaIA/internal/helloworld" 
    "farmaIA/internal/database"
    "farmaIA/cmd/api/swagger"
>>>>>>> b85a42f (feat: Added individual routes file)
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
    
    service := database.New()

<<<<<<< HEAD
    swaggerUrl := fmt.Sprintf("http://localhost:%v/swagger/doc.json", os.Getenv("PORT"))
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerUrl),
	))

	r.Get("/test", s.TestHandler)
	r.Get("/health", s.healthHandler)
=======
    healthHandler := healthcheck.NewHandler(service)
    healthHandler.RegisterHandlers(r)
    
    helloWorldHandler := helloworld.NewHandler(service)
    helloWorldHandler.RegisterHandlers(r)
    
    swagger.SwaggerHandler(r)
>>>>>>> b85a42f (feat: Added individual routes file)

    r.Get("/survey", s.getSurveyHandler)
    r.Get("/gemini", s.geminiHandler)

	return r
}

// @Tags Test
// @Produce json
// @Success 200 {object} map[string]string
// @Router /test [get]
func (s *Server) TestHandler(w http.ResponseWriter, r *http.Request) {
	respQuery, err := s.db.GetTeste()

<<<<<<< HEAD
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	jsonResp, err := json.Marshal(respQuery)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
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
=======

>>>>>>> b85a42f (feat: Added individual routes file)
