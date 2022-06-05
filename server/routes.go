package server

import (
	"log"
	"net/http"
)

func (mx *MuxServer) initializeRoutes() {
	mx.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit the / route successfully")
		return
	}).Methods("GET")
}
