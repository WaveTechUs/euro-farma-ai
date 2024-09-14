package survey

import (
	"encoding/json"
	"net/http"

    "github.com/go-chi/chi/v5"
)

type Handler struct {
	service SurveyService
}

func NewHandler(service SurveyService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterHandlers(r *chi.Mux) {
	r.Get("/survey", h.surveyHandler)
}

// @Tags Survey
// @Produce json
// @Success 200 {object} map[string]string
// @Router /survey [get]
func (h *Handler) surveyHandler(w http.ResponseWriter, r *http.Request) {
	surveys, err := h.service.GetSurveys()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResp, _ := json.Marshal(surveys)
	_, _ = w.Write(jsonResp)
}
