package server

import (
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type MuxServer struct {
	port     string
	router   *mux.Router
	handlers http.Handler
}

func NewMuxServer(ctx context.Context, db *sql.DB, app_port string) MuxServer {
	mxRouter := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	// Cors wrapper around the routes handler
	cors_wrapped_handlers := c.Handler(mxRouter)

	return MuxServer{
		port:     app_port,
		router:   mxRouter,
		handlers: cors_wrapped_handlers,
	}

}

func (mx *MuxServer) Run() {
	mx.initializeRoutes()
	log.Printf("Server starting on %s\n", mx.port)
	log.Fatal(http.ListenAndServe(":"+mx.port, mx.handlers))
}
