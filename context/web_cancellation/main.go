package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func webServer(w http.ResponseWriter, r *http.Request) {
	time := time.NewTicker(10 * time.Second)

	select {
	case <-r.Context().Done():
		log.Println("Error when processing request:", r.Context().Err())
		return
	case <-time.C:
		log.Println("Writing response....")
		_, err := io.WriteString(w, "Hello context")
		if err != nil {
			log.Println("Error when writing responce", err)
		}
		return
	}

}

func main() {
	http.HandleFunc("/", webServer)
	log.Println("Starting web server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
