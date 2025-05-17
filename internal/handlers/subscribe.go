package handlers

import (
	"log"
	"net/http"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	log.Print("Subscribe")
}
