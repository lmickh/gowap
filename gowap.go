package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type HealthCheck struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Metric  string `json:"metric"`
}

var svc_name string = os.Getenv("SERVICE_NAME")

func check_handler(w http.ResponseWriter, r *http.Request) {
	// Query consul KV for metric
	metric_url := fmt.Sprintf("http://172.17.0.1:8500/v1/kv/app/%s/metric", svc_name)
	metric_resp, err := http.Get(metric_url + "?raw")
	if err != nil {
		panic(err)
	}
	defer metric_resp.Body.Close()
	value, err := ioutil.ReadAll(metric_resp.Body)

	// Create health-check response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	hc := HealthCheck{Status: "OK", Message: svc_name, Metric: string(value)}
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
