package user

import (
	"encoding/json"
	utils "farmaIA/pkg/utils"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	service UserService
}

func NewHandler(service UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterHandlers(r *chi.Mux) {
	r.Get("/user", h.userHandler)
	r.Post("/user/login", h.userHandlerLogin)
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
	err := utils.JwtAuth(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
        return
	}
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
func (h *Handler) userHandlerRegister(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cPassword, _ := utils.Encryption(data["password"])

	newUser := User{
		Name:     data["name"],
		Email:    data["email"],
		Password: cPassword,
		Role:     data["role"],
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.service.PostUser(newUser)
}

// User Login
// @Tags User
// @Summary      insert user test
// @Description  post string
// @Produce json
// @Failure      400  {object}  string
// @Failure      400  {object}  string
// @Failure      400  {object}  string
// @Success 200 {object} string
// @Router /user/login [Post]
func (h *Handler) userHandlerLogin(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.service.GetUser(data["email"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if data["email"] != user.Email {
		http.Error(w, "Email invalido", http.StatusBadRequest)
		return
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		http.Error(w, "Senha invalido", http.StatusBadRequest)
		return
	}

	claims := utils.JwtClaims(int(user.Id), user.Role)
	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		http.Error(w, "Não foi possivel logar", http.StatusInternalServerError)
	}
	err = utils.JwtAuth(r, false)
	if err != nil {
		utils.SetCookie(token, w)     
	}

    w.WriteHeader(http.StatusOK)
    jsonResp, _ := json.Marshal(user)
    _, _ = w.Write(jsonResp)
}
