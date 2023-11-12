package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

var count1 int

func Execute1(w http.ResponseWriter, r *http.Request) {
	count1++
	json.NewEncoder(w).Encode(count1)
	log.Println(count1)
}
