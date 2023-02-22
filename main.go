package main

import (
	"context"
	"devtrekker/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	h := handlers.NewHandler()
	router := mux.NewRouter()

	router.HandleFunc("/telephone", h.GetTelephones).Methods("GET")
	router.HandleFunc("/telephone/{id:[0-9]+}", h.GetTelephoneById).Methods("GET")
	router.HandleFunc("/telephone", h.UploadTelephone).Methods("POST")
	router.HandleFunc("/telephone/{id:[0-9]+}", h.DeleteTelephone).Methods("DELETE")

	s := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      router,            // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		log.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
