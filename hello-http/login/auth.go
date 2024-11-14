// auth.go
package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var ErrUnauthorized = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	sessionToken, err := r.Cookie("session_token")
	if err != nil || sessionToken.Value == "" {
		fmt.Println("No session token")
		return ErrUnauthorized
	}

	session, ok := sessionDB[sessionToken.Value]

	if !ok {
		fmt.Println("No session in db")
		return ErrUnauthorized
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != session.CSRFToken {
		fmt.Println("CSRF token mismatch", csrf, session.CSRFToken)
		return ErrUnauthorized
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validatePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

	return base64.URLEncoding.EncodeToString(bytes)
}
