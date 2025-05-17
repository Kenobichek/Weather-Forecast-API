package handlers

import (
	"log"
	"net/http"
)

func Confirm(w http.ResponseWriter, r *http.Request) {
	log.Print("Confirm")
}
