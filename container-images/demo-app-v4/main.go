package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const filePath = "/var/local/demo.txt"

func demoH(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if _, statErr := os.Stat(filePath); statErr == nil {
			data, readErr := os.ReadFile(filePath)
			if readErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(readErr)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("Data stored on pod %s: %s\n", os.Getenv("HOSTNAME"), data)))
		} else if errors.Is(statErr, os.ErrNotExist) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("No data found on pod %s\n", os.Getenv("HOSTNAME"))))
			return
		}
	case "POST":
		file, openErr := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
		if openErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(openErr)
			return
		}
		_, copyErr := io.Copy(file, req.Body)
		if copyErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(copyErr)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Data stored on pod %s\n", os.Getenv("HOSTNAME"))))
	}
}

func main() {
	http.HandleFunc("/", demoH)

	http.ListenAndServe(":8080", nil)
}
