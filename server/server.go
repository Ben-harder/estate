package server

import (
	"log"
	"net/http"
	"strconv"
)

func NewServer(port int) *server {
	svr := &server{}
	svr.port = strconv.Itoa(port)
	return svr
}

type ServerInterface interface {
	ListenAndServe(port int)
}

type server struct {
	port string
}

func (svr *server) ListenAndServe() {
	log.Println("Starting server on port", svr.port)

	err := http.ListenAndServe(":"+svr.port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
