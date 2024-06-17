package main

import (
	"fmt"

	"github.com/dylanconnolly/bbrecs/http"
)

func main() {
	fmt.Println("BB RECS")
	s := http.NewServer()
	s.Serve()
}
