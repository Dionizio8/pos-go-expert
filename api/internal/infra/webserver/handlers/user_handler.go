package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Dionizio/pos-go-expert/api/internal/dto"
	"github.com/Dionizio/pos-go-expert/api/internal/entity"
	"github.com/Dionizio/pos-go-expert/api/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Menssage string `json:"message"`
}

type UserHandler struct {
	UserDB        database.UserInterface
	Jwt           *jwtauth.JWTAuth
	JwtExperiesIn int
}

func NewUserHanlder(userDB database.UserInterface, jwt *jwtauth.JWTAuth, jwtExperiesIn int) *UserHandler {
	return &UserHandler{UserDB: userDB, Jwt: jwt, JwtExperiesIn: jwtExperiesIn}
}

// GetJWT godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body     dto.GetJWTInput  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msgErr := Error{Menssage: err.Error()}
		json.NewEncoder(w).Encode(msgErr)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		msgErr := Error{Menssage: err.Error()}
		json.NewEncoder(w).Encode(msgErr)
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		msgErr := Error{Menssage: err.Error()}
		json.NewEncoder(w).Encode(msgErr)
		return
	}

	_, tokenstring, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExperiesIn)).Unix(),
	})

	w.Header().Set("Content-Type", "appliction/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.GetJWTOutput{AccessToken: tokenstring})
}

// Create user godoc
// @Summary     Create user
// @Description Create user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       request body dto.CreateUserInput true "user request"
// @Success     201
// @Failure     500 {object} Error
// @Router      /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msgErr := Error{Menssage: err.Error()}
		json.NewEncoder(w).Encode(msgErr)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msgErr := Error{Menssage: err.Error()}
		json.NewEncoder(w).Encode(msgErr)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
