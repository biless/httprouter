package main

import (
	"fmt"
	_ "github.com/julienschmidt/httprouter"
	"github.com/tidwall/gjson"
	"net/http"
	"httprouter"
	"log"
)


const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {

	r := httprouter.New()

	r.Prefix("/api").Use().Path("").Use().Get(func() {

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, gjson.Parse(json).String())
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
