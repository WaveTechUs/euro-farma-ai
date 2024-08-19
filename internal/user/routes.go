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
    r.Get("/user", h.userHandler)
    r.Get("/user/test", h.userHandlerTest)
    r.Post("/user/test/login", h.userHandlerLoginTest)
    r.Post("/user/register", h.userHandlerRegister)
}

// Get Users
// @Summary      Show all users
// @Description  get string
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  string
// @Router       /user [get]
func (h *Handler) userHandler(w http.ResponseWriter, r *http.Request) {
    users, err := h.service.GetUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return 
    }
	jsonResp, _ := json.Marshal(users)
	_, _ = w.Write(jsonResp)
}
// User Register
// @Tags User 
// @Summary      insert user 
// @Description  post string
// @Tags         User
// @Accept       json
// @Produce      json
// @Success 200 {object} map[string]string
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router /user/register [Post]
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
// User Login Test
// @Tags User
// @Summary      insert user test
// @Description  post string
// @Produce json
// @Failure      400  {object}  string
// @Failure      400  {object}  string
// @Failure      400  {object}  string
// @Success 200 {object} string
// @Router /user/test/login [Post]
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
// User Test
// @Tags User
// @Produce json
// @Success 200 {object} map[string]string
// @Router /user/test [get]
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

