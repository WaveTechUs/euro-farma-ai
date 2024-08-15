package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Handler struct {
	service GeminiService
}

func NewHandler(service GeminiService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *chi.Mux) {
	r.Get("/gemini", h.geminiHandler)
	r.Post("/gemini", h.geminiHandlerChat)
}

func (s *Handler) geminiHandlerChat(w http.ResponseWriter, r *http.Request) {
    geminiRequest := GeminiRequest{}
    err := json.NewDecoder(r.Body).Decode(&geminiRequest)
    if err != nil {
        log.Fatalf("Error on decoding request. Err: %v", err)
    }

    fmt.Println(geminiRequest.Text)
}

func (s *Handler) geminiHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp := s.Gemini()
	_, _ = w.Write(jsonResp)
}

func (h *Handler) Gemini() []byte {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

    model := configureGemini(client)
	resp, err := model.GenerateContent(ctx, genai.Text("Faça um resumo do livro frankenstein de mary shelley"))
	if err != nil {
		log.Printf("error: %v", err)
	}

	jsonResp := stringResponse(resp)
	return jsonResp
}

func configureGemini(client *genai.Client) *genai.GenerativeModel {
	model := client.GenerativeModel("gemini-1.5-flash")
	model.GenerationConfig = genai.GenerationConfig{
		ResponseMIMEType: "application/json",
	}
	return model
}

func stringResponse(resp *genai.GenerateContentResponse) []byte {
	var jsonResult []byte

	for _, c := range resp.Candidates {
		if c.Content != nil {
			jsonBytes, err := json.MarshalIndent(*c.Content, "", "")
			if err != nil {
				log.Fatal(err)
			}

			jsonResult = jsonBytes
		}
	}

	return jsonResult
}
