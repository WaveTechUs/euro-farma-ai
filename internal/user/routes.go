package user

import (
    "github.com/go-chi/chi/v5"
    "encoding/json"
    "net/http"
	"log"
    "golang.org/x/crypto/bcrypt"
)

type Handler struct{
    service UserService
}

func NewHandler(service UserService) *Handler {
    return &Handler{service: service}
}

// @Tags User
// @Produce json
// @Success 200 {object} map[string]string
// @Router /user [get]
func (h *Handler) RegisterHandlers(r *chi.Mux)  {
    r.Get("/user", h.userHandler)
    r.Get("/user/test", h.userHandlerTest)
    r.Post("/user/register", h.userHandlerRegister)
}

func (h *Handler) userHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(h.service.User())
	_, _ = w.Write(jsonResp)
}

func (h *Handler) userHandlerRegister(w http.ResponseWriter, r *http.Request){
    var data map[string]string

    if err := json.NewDecoder(r.Body).Decode(&data); err != nil{
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    

    cPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
    newUser := User{
        Name: data["name"],
        Email: data["email"],
        Password: cPassword,
        Role: data["role"],
    }

    w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func (h *Handler) userHandlerTest(w http.ResponseWriter, r *http.Request)  {
	resp := make(map[string]string)
	resp["id"] = "1"
	resp["name"] = "jose"
	resp["email"] = "j@j.com"
	resp["password"] = "123"
	resp["role"] = "admin"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

