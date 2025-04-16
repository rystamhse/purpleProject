package auth

import (
	"fmt"
	"net/http"
	"purpleProject/configs"
	"purpleProject/pkg/req"
	"purpleProject/pkg/res"
)

type AuthConfigDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthConfigDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())

}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		res.SendJSON(w, body, 201)
	}
}
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			fmt.Println(err)
			return
		}
		res.SendJSON(w, body, 201)
	}
}
