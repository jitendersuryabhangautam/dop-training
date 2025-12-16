package api

// User is a simple struct that matches the JSON we expect in the request body.
// Example JSON:
// { "name": "Alice", "email": "alice@example.com", "age": 25 }
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
