package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/djamaile/mango/pkg/releases"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(releases.CollectVizReleases())
	fmt.Println("test")
}
