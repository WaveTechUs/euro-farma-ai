package gemini

import (
	"context"
	"encoding/json"
	S "farmaIA/internal/survey"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
	"strings"
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

func (h *Handler) geminiHandler(w http.ResponseWriter, r *http.Request) {
	input := h.getBody(r)
	jsonResp := h.Gemini(input)
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

	h.makeHistory(cs, body.History)

	res, err := cs.SendMessage(ctx, genai.Text(body.Text))
	if err != nil {
		log.Fatal(err)
	}

	jsonResp := getJsonReponse(res)
	return jsonResp
}

func (h *Handler) makeHistory(cs *genai.ChatSession, gr []History) {
	gr = h.setPreInformation(gr)
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

func (h *Handler) setPreInformation(gr []History) []History {
	surveys, err := h.service.GetSurveys()
	if err != nil {
		log.Fatalf("Error on getting surveys. Err: %v", err)
	}

	surveysStr := SurveysToString(surveys)
	information := "Se você não souber a resposta diga que não sabe, e responde apenas o que está no seu escopo, o seu escopo é: Respostas sobre pesquisas farmacêuticas e com base nos seguintes dados."
	information += surveysStr

	command := History{OldText: information, Role: "user"}
	respCommand := History{OldText: "Ok", Role: "model"}

	gr = append([]History{command, respCommand}, gr...)
	return gr
}

func ToString(s S.Survey) string {
	return fmt.Sprintf("Id: %d\nCreatedAt: %s\nDescription: %s\nTopic: %s\nStatus: %s\nSummary: %s\nConclusions: %s\nMethod: %s\nKeywords: %s\n",
		s.Id, s.CreatedAt, s.Description, s.Topic, s.Status, s.Summary, s.Conclusions, s.Method, s.Keywords)
}

func SurveysToString(surveys []S.Survey) string {
	var builder strings.Builder
	for _, survey := range surveys {
		builder.WriteString(ToString(survey))
		builder.WriteString("\n")
	}
	return builder.String()
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
