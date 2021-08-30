package api

import (
	"encoding/json"
	"net/http"

	"github.com/djamaile/mango/pkg/releases"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(releases.CollectVizReleases())
}
