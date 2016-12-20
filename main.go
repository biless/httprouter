package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	_ "httprouter"
	"net/http"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, gjson.Parse(json).String())
	})

	http.ListenAndServe(":8080", nil)
}
