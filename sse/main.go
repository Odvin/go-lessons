package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")

	tokens := []string{"this", "is", "a", "live", "event", "from", "the", "server"}

	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", token)
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		time.Sleep(time.Second)
	}
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "SSE API server port")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	mux := http.NewServeMux()
	mux.HandleFunc("/events", events)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting server on %s", srv.Addr)

	err := srv.ListenAndServe()
	logger.Fatal(err)
}
