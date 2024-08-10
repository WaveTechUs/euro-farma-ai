package helloworld

import  (
    "github.com/go-chi/chi/v5"
    "encoding/json"
    "net/http"
	"farmaIA/internal/database"
	"log"
)

type Handler struct{
    db database.Service
}

func NewHandler(db database.Service) *Handler {
    return &Handler{db: db} 
}

// @Tags HelloWorld
// @Produce json
// @Success 200 {object} map[string]string
// @Router /helloworld [get]
func (h *Handler) RegisterHandlers(r *chi.Mux) {
	r.Get("/", h.HelloWorldHandler)
}

func (h *Handler) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Teste Farma"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
