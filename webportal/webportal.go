package webportal

import (
	"fmt"
	"net/http"
)

func Portal(addr string) error {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	return http.ListenAndServe(addr, nil)
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome %s", r.RemoteAddr)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.RequestURI)
}
