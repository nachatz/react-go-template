package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/time/rate"

	"github.com/go-chi/chi"
	"github.com/nachatz/React-Go-Template/app/internal/api"
	"github.com/nachatz/React-Go-Template/app/internal/config"
	"github.com/nachatz/React-Go-Template/app/internal/handlers"
	"github.com/nachatz/React-Go-Template/app/internal/middleware"
	"github.com/nachatz/React-Go-Template/app/internal/models"
)

func InitRoutes(cfg *config.Config) http.Handler {
	/* InitRoutes - Initializes the routes for the application.
	   @Param *cfg.Config - The configuration object.
	   @Return http.Handler - The router.
	*/
	mux := chi.NewRouter()
	var limiter = rate.NewLimiter(rate.Limit(5), 10)

	// Rate limit all endpoints
	mux.Use(middleware.RateLimitMiddleware(limiter))

	// Non-authenticated endpoints
	mux.Post(api.EndpointAuthToken, func(w http.ResponseWriter, r *http.Request) {
		handlers.GenerateJWTHandler(w, r, cfg.Auth.ClientSecret, cfg.Auth.ClientID)
	})

	// Add authenticated endpoints
	mux.Get(api.EndpointHello, func(w http.ResponseWriter, r *http.Request) {
		middleware.JwtMiddleware(handlers.HelloHandler, cfg.Auth.ClientSecret).ServeHTTP(w, r)
	})

	mux.MethodNotAllowed(methodNotAllowedHandler())

	// log the names of all the endpoints in the router
	logEndpointNames(mux)

	return mux
}

func logEndpointNames(router chi.Router) {
	/* logEndpointNames - Logs the names of all the endpoints in the router.
	   @Param chi.Router - The router to log the endpoint names of.
	*/
	walker := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Println("Endpoint:", route)
		return nil
	}

	if err := chi.Walk(router, walker); err != nil {
		log.Fatal("Error walking the router:", err)
	}
}

func methodNotAllowedHandler() http.HandlerFunc {
	/* methodNotAllowedHandler - Handles the method not allowed error by returning an appropriate response.
	   @Return http.Handler - The handler for method not allowed error.
	*/
	response := models.Response{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    "Method not allowed",
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		json.NewEncoder(w).Encode(response)
	})
}
