package main

import (
	"bbrecs/api"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "WELCOME! This is the Index page")
}

func main() {
	fmt.Println("BB RECS")
	s := api.NewServer()
	s.Serve()
}
