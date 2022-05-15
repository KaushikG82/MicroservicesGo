package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/KaushikG82/MicroservicesGo.git/handlers"
)

func main() {
	Log := log.New(os.Stdout, "product-api", log.LstdFlags)
	HelloHandler := handlers.NewHello(Log)
	GoodByeHandler := handlers.NewGoodBye(Log)

	sm := http.NewServeMux()
	sm.Handle("/", HelloHandler)
	sm.Handle("/goodbye", GoodByeHandler)

	//http.ListenAndServe(":9090", sm)
	Server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := Server.ListenAndServe()
		if err != nil {
			Log.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	Log.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	Server.Shutdown(tc)
}
