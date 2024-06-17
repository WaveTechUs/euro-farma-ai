package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	r.Get("/test", s.TestHandler)

	r.Get("/health", s.healthHandler)

	return r
}

func (s *Server) TestHandler(w http.ResponseWriter, r *http.Request) {
	respQuery, err := s.db.GetTeste()

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	jsonResp, err := json.Marshal(respQuery)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

// HelloWorldHandler retorna uma mensagem simples em JSON.
// @Summary Exemplo de endpoint Hello World
// @Description Retorna uma mensagem simples "Teste Farma"
// @Tags Exemplos
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Teste Farma"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
// healthHandler retorna o status de saúde do banco de dados em JSON.
// @Summary Endpoint de verificação de saúde
// @Description Retorna o status de saúde do banco de dados
// @Tags Saúde
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
