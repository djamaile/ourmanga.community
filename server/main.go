package main

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/djamaile/mango/pkg/releases"
	"log"
	"net/http"
)

func vizHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(releases.CollectVizReleases())
}

func yenHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(releases.CollectYenPressReleases())
}

func sevenSeasHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(releases.CollectSevenSeasReleases())
}

func darkHorseHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(releases.CollectDarkHorseReleases())
}

func kodanshaHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(releases.CollectKodanshaReleases())
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/releases/viz", vizHandler)
	r.HandleFunc("/releases/yenpress", yenHandler)
	r.HandleFunc("/releases/sevenseas", sevenSeasHandler)
	r.HandleFunc("/releases/darkhorse", darkHorseHandler)
	r.HandleFunc("/releases/kodansha", kodanshaHandler)

	log.Println("Listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}


