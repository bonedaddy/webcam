// Command webcam serves a WebRTC stream from a camera.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"0f.io/webcam/webrtc"
)

// Embed the UI HTML/JS resources in the Go bindary.
//go:generate embedder -path ui -package main -name uifs

func handleSession(w http.ResponseWriter, r *http.Request) {
	offer, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Couldn't read offer body: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	s := webrtc.NewSession()
	err = s.Remote(offer)
	if err != nil {
		log.Printf("Couldn't set remote description: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := s.Description()
	if err != nil {
		log.Printf("Couldn't get local description: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "%s", answer)
}

func main() {
	addr := flag.String("addr", ":8003", "http address to listen on")
	flag.Parse()
	http.HandleFunc("/session", handleSession)
	http.Handle("/", http.FileServer(uifs))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
