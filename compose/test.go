package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	const port = "8180"
	log.Printf("Webapp running on port %s", port)
	http.HandleFunc("/", index)
	http.ListenAndServe(fmt.Sprintf(`:%s`, port), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	color := os.Getenv("BG_COLOR")
	log.Printf("req from %s\n", r.Host)
	fmt.Fprintf(w, `<html><head><style>body{text-align:center;background-color:%s;}</style><link rel="icon" type="image/png" href="data:image/png;base64,iVBORw0KGgo="><title>Hello</title></head><body><h1>Hello Docker!</h1></body></html>`, color)
}
