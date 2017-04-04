package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dimfeld/httptreemux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	tree := httptreemux.New()
	tree.Handle("GET", "/echo", handleEcho)
	tree.Handle("GET", "/reverse", handleReverse)

	srv := &http.Server{Addr: ":" + port, Handler: tree}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")

	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Server gracefully stopped")
}

//================================================================================

type dataPack struct {
	Input json.RawMessage `json:"input"`
}

type responsePack struct {
	Output json.RawMessage `json:"output"`
}

// handleReverse receives json data from the giving request and attempts to reverse
// the giving string.
// Expects: {input:"DATA"} json.
func handleReverse(w http.ResponseWriter, r *http.Request, params map[string]string) {
	var pack dataPack

	if err := json.NewDecoder(r.Body).Decode(&pack); err != nil {
		http.Error(w, "Invalid Data Recieved: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Recieved: %+q\n", pack)
}

// handleEcho receives a json data and echos back the given data received.
// Expects: {input:"DATA"} json.
func handleEcho(w http.ResponseWriter, r *http.Request, params map[string]string) {

}
