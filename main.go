package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello World!")
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			//writer.WriteHeader(http.StatusBadRequest)
			//writer.Write([]byte("Ooops, something went wrong!"))
			http.Error(writer, "Ooops, something went wrong!", http.StatusBadRequest)
			return
		}
		log.Printf("Data %s", data)
		fmt.Fprintf(writer, "Hello %s!", data)
	})
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GoodBye World!")
	})
	http.ListenAndServe(":9090", nil)
}
