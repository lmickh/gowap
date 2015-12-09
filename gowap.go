package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type HealthCheck struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var svc_name string = os.Getenv("SERVICE_NAME")

func check_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	hc := HealthCheck{Status: "OK", Message: svc_name}
	resp, err := json.Marshal(hc)

	if err != nil {
		panic(err)
	} else {
		w.Write(resp)
		return
	}
}

func main() {
	http.HandleFunc("/health-check", check_handler)
	http.ListenAndServe(":8080", nil)
}
