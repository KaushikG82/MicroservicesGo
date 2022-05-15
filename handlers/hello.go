package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	Log *log.Logger
}

func NewHello(Log *log.Logger) *Hello {
	return &Hello{Log}
}

func (h *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.Log.Println("Hello World!")
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Something went wrong!!", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(writer, "Hello %s", data)
}
