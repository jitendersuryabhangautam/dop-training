package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// usersPostHandler handles POST /users requests.
// Steps:
// 1) Check that the method is POST
// 2) Read and decode the JSON body into a User
// 3) Do a few basic validations
// 4) Return a simple JSON response
func usersPostHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method for this endpoint
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// We will respond with JSON
	w.Header().Set("Content-Type", "application/json")

	// Decode the JSON body into our User struct
	var u User
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // if extra fields are sent, reject them (beginner tip)
	if err := dec.Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid JSON: " + err.Error(),
		})
		return
	}

	// Basic validations (beginner-friendly)
	if u.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "name is required"})
		return
	}
	if u.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "email is required"})
		return
	}
	if u.Age <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "age must be a positive number"})
		return
	}

	// Simulate creating the user (in a real app we would save to a database)
	// We'll just create a simple ID based on the length of fields to show something happening.
	id := fmt.Sprintf("usr_%d", len(u.Name)+len(u.Email)+u.Age)

	// Send back a friendly JSON response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "user created",
		"id":      id,
		"user":    u,
	})
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "pong",
	})
}

// StartServer boots a minimal HTTP server and registers routes under /users.
// Port can be provided via addr (e.g., ":8080"). If empty, defaults to ":8080".
func StartServer(addr string) error {
	if addr == "" {
		addr = ":8080"
	}

	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/users", usersPostHandler)

	// Print a message so beginners can see where to send requests
	fmt.Printf("Server running at http://localhost%s!\n", addr)
	// Pass nil to ListenAndServe to use the default mux
	return http.ListenAndServe(addr, nil)
}
