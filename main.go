package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"sigaa-mock/mocks"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	log.Println("SIGAA mock running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != "Bearer apitoken123" {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = io.WriteString(w, "unauthorized")
		return
	}

	q := r.URL.Query()
	username := q.Get("username")
	if at := strings.Index(username, "@"); at != -1 {
		username = username[:at]
	}
	passwordB64 := q.Get("password")
	passwordBytes, _ := base64.StdEncoding.DecodeString(passwordB64)
	password := string(passwordBytes)

	user, ok := mocks.Scenarios[username]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = io.WriteString(w, fmt.Sprintf("invalid credentials for %s/%s", username, password))
		return
	}

	if password != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = io.WriteString(w, fmt.Sprintf("invalid credentials for %s/%s", username, password))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user.Payload)
}
