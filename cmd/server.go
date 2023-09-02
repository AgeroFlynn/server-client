package main

import (
	"encoding/json"
	"github.com/AgeroFlynn/server-client/pkg/server/api"
	"io"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Read body error", http.StatusBadRequest)
		return
	}

	var dto api.RequestDTO
	if err = json.Unmarshal(body, &dto); err != nil {
		http.Error(w, "Unmarshal error", http.StatusBadRequest)
		return
	}

	// ---
	// do some business logic
	// ---

	result := api.ResponseDTO{Result: "some result"}

	data, err := json.Marshal(&result)
	if err != nil {
		http.Error(w, "Response marshal error", http.StatusBadRequest)
	}

	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Failed to response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/api/v1/method", PostHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
