package main

import (
	"fmt"
	_ "github.com/julienschmidt/httprouter"
	_ "github.com/pquerna/ffjson/fflib/v1"
	"github.com/tidwall/gjson"
	_ "httprouter"
	_ "log"
	"net/http"
	"log"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func print1(i int, next Next) {
	println("func1", i)
	next()
}

func print2(i int, next Next) {
	println("func2", i)
	next()
}

func print3(i int, next Next) {
	println("func3", i)
	next()
}

type Next func()

type Handle func(int, Next)

var now Handle = nil
var next Handle = nil


func main() {
	//_ := []int{1,2,3,4}
	//r := httprouter.New()

	m := []Handle{}

	m = append([]Handle{print1}, m...)
	m = append([]Handle{print2}, m...)
	m = append([]Handle{print3}, m...)
	//r.Prefix("/api").Use().Path("").Use().Get(func() {
	//
	//})

	temp := func() {
		println("last")
	}

	// 还是需要倒叙构建
	for i, v := range m {
		println(i, v)
		s := temp
		t := v
		in := i
		temp = func() {
			t(in, s)
		}
	}

	for i, v := range m {
		println(i, v)
		s := temp
		t := v
		in := i + 3
		temp = func() {
			t(in, s)
		}
	}

	temp()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, gjson.Parse(json).String())
	})

	// TODO: 感觉还是要参考一下底层实现
	//log.Fatal(http.ListenAndServe(":8080", r))
}
