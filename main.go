package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "WELCOME! This is the Index page")
}

func main() {
	fmt.Println("BB RECS")
	router := httprouter.New()
	router.GET("/", HomeHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
