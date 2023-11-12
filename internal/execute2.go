package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

var count2 int
var mutex sync.Mutex

func Count() int {
	mutex.Lock()
	defer mutex.Unlock()
	count2++
	time.Sleep(1000 * time.Nanosecond)
	return count2
}

func Execute2(w http.ResponseWriter, r *http.Request) {
	count := Count()
	json.NewEncoder(w).Encode(count)
	log.Println(count)
}
