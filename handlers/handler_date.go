package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func Date(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Local()
	fmt.Fprintf(w, currentTime.String())
}
