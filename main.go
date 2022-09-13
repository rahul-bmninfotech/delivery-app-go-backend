package main

import (
	r2 "DeliveryApp/router"
	"log"
	"net/http"
)

func main() {
	router := r2.NewDeliveryRouter()
	log.Fatal(http.ListenAndServe(":8087", router))
}
