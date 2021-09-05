package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/djamaile/mango/pkg/releases"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

func tokyopopHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(releases.CollectTokyoPopReleases())
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/releases/viz", vizHandler)
	r.HandleFunc("/api/releases/yenpress", yenHandler)
	r.HandleFunc("/api/releases/sevenseas", sevenSeasHandler)
	r.HandleFunc("/api/releases/darkhorse", darkHorseHandler)
	r.HandleFunc("/api/releases/kodansha", kodanshaHandler)
	r.HandleFunc("/api/releases/tokyopop", tokyopopHandler)

	handler := cors.Default().Handler(r)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
