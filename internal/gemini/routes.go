package gemini

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Handler struct{
    service GeminiService
}

func NewHandler(service GeminiService) *Handler {
    return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *chi.Mux)  {
    r.Get("/gemini", h.geminiHandler)
}

func (s *Handler) geminiHandler (w http.ResponseWriter, r *http.Request) {
    s.Gemini()
    w.Write([]byte("Gemini"))
}

func (h *Handler) Gemini() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("Faça um resumo do livro frankenstein de mary shelley"))
	if err != nil {
        log.Printf("error: %v", err)
	}
    h.service.Gemini()
    fmt.Print("Gemini:")
    printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
