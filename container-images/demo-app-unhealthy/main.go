package main

import (
	"fmt"
	"net/http"
	"os"
)

var i = 1

func demoH(w http.ResponseWriter, req *http.Request) {
	if i >= 5 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Some Error Occurred!\n"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello from %s\n", os.Getenv("HOSTNAME"))))
	}
	i = i + 1
}

func main() {
	http.HandleFunc("/", demoH)

	http.ListenAndServe(":8080", nil)
}
