package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Wait(w http.ResponseWriter, r *http.Request) {
	parameters, ok := r.URL.Query()["delay"]
	if !ok {
		fmt.Fprintf(w, "Parametr 'delay' is missing.")
		return
	}
	delay, err := strconv.ParseUint(parameters[0], 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Parametr 'delay' is wrong. It must be an Uint64 number.")
		return
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Fprintf(w, "Sleep for "+strconv.FormatUint(delay, 10)+"ms completed")
}
