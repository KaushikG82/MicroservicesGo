package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	Log *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	g.Log.Printf("Going to say Byee")
	writer.Write([]byte("Byee !"))
}
