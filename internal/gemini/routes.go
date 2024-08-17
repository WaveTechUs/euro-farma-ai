package gemini

import (
	"context"
	"encoding/json"
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
	r.Post("/gemini", h.geminiHandler)
}

func (s *Handler) geminiHandler(w http.ResponseWriter, r *http.Request) {
	input := s.getBody(r)
	jsonResp := s.Gemini(input)
	_, _ = w.Write(jsonResp)
}

func (h *Handler) Gemini(body GeminiRequest) []byte {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := configureGemini(client)
	cs := model.StartChat()

	makeHistory(cs, body.History)

	res, err := cs.SendMessage(ctx, genai.Text(body.Text))
	if err != nil {
		log.Fatal(err)
	}

	jsonResp := getJsonReponse(res)
	return jsonResp
}

func makeHistory(cs *genai.ChatSession, gr []History) {
	cs.History = make([]*genai.Content, len(gr))

	for i, item := range gr {
		cs.History[i] = &genai.Content{
			Parts: []genai.Part{
				genai.Text(item.OldText),
			},
			Role: item.Role,
		}
	}
}

func configureGemini(client *genai.Client) *genai.GenerativeModel {
	model := client.GenerativeModel("gemini-1.5-flash")
	return model
}

func getJsonReponse(resp *genai.GenerateContentResponse) []byte {
	var jsonResult []byte

	for _, c := range resp.Candidates {
		if c.Content != nil {
			jsonBytes, err := json.Marshal(*c.Content)
			if err != nil {
				log.Fatal(err)
			}

			jsonResult = jsonBytes
		}
	}
	return jsonResult
}

func (s *Handler) getBody(r *http.Request) GeminiRequest {
    gr := GeminiRequest{}
    err := json.NewDecoder(r.Body).Decode(&gr)
    if err != nil {
        log.Fatalf("Error on decoding request. Err: %v", err)
    }
    return gr
}
