package auth

import (
	"net/http"
	"v1/familyManager/configs"
	"v1/familyManager/internal/family"
	"v1/familyManager/pkg/jwt"
	"v1/familyManager/pkg/middleware"
	"v1/familyManager/pkg/req"
	"v1/familyManager/pkg/res"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
	router.Handle("GET /iam", middleware.IsAuthed(handler.WhoIAm(), deps.Config))
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			res.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}
		email, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := LoginResponse{
			Token: token,
		}
		res.Json(w, data, http.StatusCreated)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			res.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}
		email, err := handler.AuthService.Register(body.Email, body.Password, body.FirstName, body.LastName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := RegisterResponse{
			Token: token,
		}
		res.Json(w, data, 201)
	}
}

func (handler *AuthHandler) WhoIAm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.Context().Value(middleware.ContextEmailKey).(string)
		user, err := handler.AuthService.UserRepository.GetByEmail(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if user.FamilyID != nil {
			familyId, _ := user.FamilyID.(int64)
			family, err := family.NewFamilyRepository(handler.AuthService.UserRepository.Storage).GetByID(familyId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			user.FamilyID = family.Name
		}
		res.Json(w, user, 200)
	}
}
