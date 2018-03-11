package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"
)

type ServerStatus struct {
	Name         string    `json:"app_name"`
	CustomerName string    `json:"customer_name"`
	DateStart    time.Time `json:"date_start"`
	PodName      string    `json:"pod_name"`
	Status       int       `json:"status"`
}

func main() {

	customer := flag.String("customer", "unknown", "the name of the customer")
	podName := flag.String("pod", "unknown", "the name of the pod")
	flag.Parse()

	app := ServerStatus{"hello-world", *customer, time.Now().UTC(), *podName, http.StatusOK}

	data, _ := json.Marshal(app)

	http.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request on /alive")
		w.Write(data)
	})

	log.Println("Server will start on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
