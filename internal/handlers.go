package internal

import (
	"log"
	"net/http"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	log.Print("GetWeather")
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	log.Print("Subscribe")
}

func Confirm(w http.ResponseWriter, r *http.Request) {
	log.Print("Confirm")
}

func Unsubscribe(w http.ResponseWriter, r *http.Request) {
	log.Print("Unsubscribe")
}