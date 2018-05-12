package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/heartbeat", heartbeatHandler)
}

func heartbeatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
