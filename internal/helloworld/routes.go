package helloworld

import  (
    "github.com/go-chi/chi/v5"
    "encoding/json"
    "net/http"
	"log"
)
type Handler struct{
    service HelloWorldService
}

func NewHandler(service HelloWorldService) *Handler {
    return &Handler{service: service} 
}

func (h *Handler) RegisterHandlers(r *chi.Mux) {
	r.Get("/", h.HelloWorldHandler)
}

func (h *Handler) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
    
	jsonResp, err := json.Marshal(h.service.HelloWorld())
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
