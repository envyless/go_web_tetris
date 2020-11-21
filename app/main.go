package main // import "github.com/johngrib/go-http-helloworld"

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	//run_game()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("/ request")
		fmt.Fprintf(w, "Hello World Changed !!!\n")
	})

	// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("/ping request")
	// 	fmt.Fprintf(w, "pong\n")
	// })

	http.HandleFunc("/process", process)

	server := &http.Server{
		Addr: ":3000",
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, "Hello World")
}

//tetris in web
var blocks [10][20]int
var strUI string

