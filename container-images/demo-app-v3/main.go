package main

import (
	"fmt"
	"net/http"
	"os"
)

func demoH(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello from %s, this is v3 speaking!\n", os.Getenv("HOSTNAME"))))
}

func main() {
	http.HandleFunc("/", demoH)

	http.ListenAndServe(":8080", nil)
}
