package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"unicode/utf8"

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
	tree.Handle("POST", "/echo", handleEcho)
	tree.Handle("POST", "/reverse", handleReverse)

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

	var bu bytes.Buffer

	if err := json.NewEncoder(&bu).Encode(responsePack{
		Output: json.RawMessage([]byte(reverse(string(pack.Input)))),
	}); err != nil {
		http.Error(w, "Invalid Data Recieved: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	io.Copy(w, &bu)
}

// handleEcho receives a json data and echos back the given data received.
// Expects: {input:"DATA"} json.
func handleEcho(w http.ResponseWriter, r *http.Request, params map[string]string) {
}

func reverse(s string) string {
	cs := make([]rune, utf8.RuneCountInString(s))
	i := len(cs)
	for _, c := range s {
		i--
		cs[i] = c
	}
	return string(cs)
}
