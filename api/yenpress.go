package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func YenHandler(w http.ResponseWriter, r *http.Request) {
	indexHTML, err := ioutil.ReadFile("/var/task/api/pages/viz-2021-9-3.html")
	check(err)
	fmt.Println(indexHTML)
	w.Write(indexHTML)
}
