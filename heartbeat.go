package heartbeat

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/heartbeat", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
