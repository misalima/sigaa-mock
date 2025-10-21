package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	log.Println("SIGAA mock running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// simples check do Authorization header (Bearer <token>)
	auth := r.Header.Get("Authorization")
	if auth != "Bearer apitoken123" {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = io.WriteString(w, "unauthorized")
		return
	}

	q := r.URL.Query()
	username := q.Get("username")
	passwordB64 := q.Get("password")
	passwordBytes, _ := base64.StdEncoding.DecodeString(passwordB64)
	password := string(passwordBytes)

	// logic example: se username/password batem, retorna 201 com payload
	if username == "jose" && password == "Test1234!" {
		w.WriteHeader(http.StatusCreated)
		payload := map[string]interface{}{
			"id_usuario": 1,
			"id_pessoa":  10,
			"nome":       "José",
			"email":      "jose@example.com",
			"cpf":        "12345678900",
			"perfil":     []string{"docente", "discente"},
		}
		_ = json.NewEncoder(w).Encode(payload)
		return
	}

	// caso padrão: 401 com corpo
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = io.WriteString(w, fmt.Sprintf("invalid credentials for %s/%s", username, password))
}
