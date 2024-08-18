package user

import (
    "github.com/go-chi/chi/v5"
    "encoding/json"
    "net/http"
	"log"
    "golang.org/x/crypto/bcrypt"
    utils "farmaIA/pkg/utils"
)

type Handler struct{
    service UserService
}

func NewHandler(service UserService) *Handler {
    return &Handler{service: service}
}

func (h *Handler) RegisterHandlers(r *chi.Mux)  {
// @Tags User
// @Produce json
// @Success 200 {object} map[string]string
// @Router /user [get]
    r.Get("/user", h.userHandler)
// @Tags User Test
// @Produce json
// @Success 200 {object} map[string]string
// @Router /user/test [get]
    r.Get("/user/test", h.userHandlerTest)
// @Tags User Login Test
// @Produce json
// @Success 200 {object} map[string]string
// @Router /user/test/login [Post]
    r.Post("/user/test/login", h.userHandlerLoginTest)
    r.Post("/user/register", h.userHandlerRegister)
}

func (h *Handler) userHandler(w http.ResponseWriter, r *http.Request) {
    users, err := h.service.GetUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return 
    }
	jsonResp, _ := json.Marshal(users)
	_, _ = w.Write(jsonResp)
}

func (h *Handler) userHandlerRegister(w http.ResponseWriter, r *http.Request){
    var data map[string]string

    if err := json.NewDecoder(r.Body).Decode(&data); err != nil{
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    cPassword, _ := utils.Encryption(data["password"]) 

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
    // Adicionar no Banco
}

func (h *Handler) userHandlerLoginTest(w http.ResponseWriter, r *http.Request){
    var data map[string]string  

    if err := json.NewDecoder(r.Body).Decode(&data); err != nil{
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    var user map[string]string
    user = h.service.GetUserTest() 
    
    if data["email"] != user["email"]{
        http.Error(w, "Email invalido", http.StatusBadRequest)
        return
    }
    cPassword, _ := utils.Encryption(user["password"]) 
    if err := bcrypt.CompareHashAndPassword(cPassword, []byte(data["password"])); err != nil{ 
        http.Error(w, "Senha invalido", http.StatusBadRequest)
        return
    }

	jsonResp, err := json.Marshal("Seja bem vindo")
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
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

