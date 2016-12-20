package main

import (
	"github.com/tidwall/gjson"
	"fmt"
	"httprouter"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Parse(json).Get("name.last")
	fmt.Println(value)
	httprouter.Router{}
}