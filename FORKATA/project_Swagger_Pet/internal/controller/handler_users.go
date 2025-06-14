package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"golang/project_Swagger_Pet/internal/service"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt/v5"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Responder interface {
	OutputJSON(w http.ResponseWriter, responseData interface{})
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

type Respond struct{}

func NewResponder() Responder {
	return &Respond{}
}

func (r *Respond) OutputJSON(w http.ResponseWriter, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		log.Println("responder json encode error:", err)
	}
}

func (r *Respond) ErrorBadRequest(w http.ResponseWriter, err error) {
	log.Println("http response bad request status code:", err)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		log.Println("response writer error on write:", err)
	}
}

func (r *Respond) ErrorInternal(w http.ResponseWriter, err error) {
	log.Println("http response internal server error:", err)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{
		Success: false,
		Message: err.Error(),
	}); err != nil {
		log.Println("response writer error on write:", err)
	}
}

type HandleUser struct {
	handler   *service.UserService
	responder Responder
}

func NewHandleUser(service *service.UserService, responder Responder) *HandleUser {
	return &HandleUser{
		handler:   service,
		responder: responder,
	}
}

// @Summary      Get User by ID
// @Description  get user
// @Accept       json
// @Produce      json
// @Param        id  path string true  "User id"
// @Success      200  {object}  Response
// @Failure      500  {string}  Response
// @Router       /users/{username} [get]
func (h *HandleUser) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user, err := h.handler.GetUser(r.Context(), username)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    user,
	})
}

// @Summary      Create User
// @Description  create
// @Accept       json
// @Produce      json
// @Param        user  body service.UserStruct true  "User data"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users [post]
func (h *HandleUser) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var user service.UserStruct

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		log.Println("Ошибка в обработке тела запроса")
		h.responder.ErrorBadRequest(w, fmt.Errorf("query parameter is required"))
		return
	}

	err = h.handler.CreateUser(r.Context(), user)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return	
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    user,
	})
}

// @Summary      Update User data Base
// @Description  update
// @Accept       json
// @Produce      json
// @Param        user  body service.UserStruct true  "User data"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users [put]
func (h *HandleUser) HandlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	var user service.UserStruct

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		log.Println("Ошибка в обработке тела запроса")
		h.responder.ErrorBadRequest(w, fmt.Errorf("query parameter is required"))
		return
	}

	err = h.handler.UpdateUser(r.Context(), user)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return	
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    user,
	})
}

// @Summary      Delete User data Base
// @Description  update
// @Accept       json
// @Produce      json
// @Param        id  path string true  "User id"
// @Success      200  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users/{username} [delete]
func (h *HandleUser) HandlerDelete(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	err := h.handler.DeleteUser(r.Context(), username)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
	})
}


// @Summary      List of users
// @Description  get users
// @Accept       json
// @Produce      json
// @Param        limit  query int false  "Лимит" default(10)
// @Param        offset query int false  "Смещение" default(0)
// @Success      200  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users [get]
func (h *HandleUser) HandlerListUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	usernameStr := query.Get("username")
	passwordStr := query.Get("password")

	if usernameStr == "" || passwordStr == "" {
		h.responder.ErrorBadRequest(w, fmt.Errorf("username and password are required query parameters"))
		return
	}

	user, err := h.handler.GetByCredentials(r.Context(), usernameStr, passwordStr)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    user,
	})
}

func (h *HandleUser) JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			h.responder.ErrorBadRequest(w, fmt.Errorf("authorization header is required"))
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			h.responder.ErrorBadRequest(w, fmt.Errorf("invalid authorization header format"))
			return
		}

		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("Jkdf04iaosfj9049a409dfpobntp"), nil
		})

		if err != nil {
			h.responder.ErrorBadRequest(w, fmt.Errorf("invalid token: %v", err))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if exp, ok := claims["exp"].(float64); ok {
				if time.Now().Unix() > int64(exp) {
					h.responder.ErrorBadRequest(w, fmt.Errorf("token expired"))
					return
				}
			}
			ctx := context.WithValue(r.Context(), "user", claims["sub"])
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			h.responder.ErrorBadRequest(w, fmt.Errorf("invalid token"))
		}
	})
}

// @Summary      User Login
// @Description  Logs user into the system
// @Accept       json
// @Produce      json
// @Param        username  query string true  "The user name for login"
// @Param        password  query string true  "The password for login"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users/login [get]
func (h *HandleUser) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	username := query.Get("username")
	password := query.Get("password")

	if username == "" || password == "" {
		h.responder.ErrorBadRequest(w, fmt.Errorf("username and password"))
		return
	}

	user, err := h.handler.GetByCredentials(r.Context(), username, password)
	if err != nil {
		h.responder.ErrorBadRequest(w, fmt.Errorf("ОШИБКА в credentials"))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Name,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("Jkdf04iaosfj9049a409dfpobntp"))
	if err != nil{
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data: map[string]string{
			"token": tokenString,
		},
	})
}