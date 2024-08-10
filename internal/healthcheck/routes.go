package healthcheck

import  (
    "github.com/go-chi/chi/v5"
    "encoding/json"
    "net/http"
)

type Handler struct{
    service HealthCheckService
}

func  NewHandler(service HealthCheckService) *Handler {
    return &Handler{service: service} 
}

// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *Handler) RegisterHandlers(r *chi.Mux) {
	r.Get("/health", h.healthHandler)
}

func (h *Handler) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(h.service.Health())
	_, _ = w.Write(jsonResp)
}
