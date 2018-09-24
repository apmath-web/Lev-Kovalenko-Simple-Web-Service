package handlers

import (
	"fmt"
	"net/http"
)

func Hello_world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow world!")
}
