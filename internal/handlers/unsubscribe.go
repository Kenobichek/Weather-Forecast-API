package handlers

import (
	"log"
	"net/http"
)

func Unsubscribe(w http.ResponseWriter, r *http.Request) {
	log.Print("Unsubscribe")
}
