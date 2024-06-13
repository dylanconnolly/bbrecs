package bbrecs

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HomeGetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "WELCOME! This is the Index page")
}
