// go를 사용한 로그인, 로그아웃 구현
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID             int
	Email          string
	HashedPassword string
}

type Session struct {
	ID           int
	UserID       int
	SessionToken string
	CSRFToken    string
}

var users = []User{
	{ID: 1, Email: "abcd@gmail.com", HashedPassword: "1234"},
}

var sessionDB = map[string]Session{}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", register)
	mux.HandleFunc("POST /login", login)
	mux.HandleFunc("POST /logout", logout)
	mux.HandleFunc("GET /profile", profile)

	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	log.Println("Starting server on :3000")
	server.ListenAndServe()
}

func register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// check valid email and password
	if email == "" || len(password) < 8 {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// check existing user
	for _, user := range users {
		if user.Email == email {
			http.Error(w, "Email already exists", http.StatusBadRequest)
			return
		}
	}

	// hash password

	hashedPassword, _ := hashPassword(password)

	// add user

	id := len(users) + 1
	user := User{ID: id, Email: email, HashedPassword: hashedPassword}
	users = append(users, user)

	fmt.Fprintln(w, "User created successfully")

}

func login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// find user
	var foundUser *User
	for _, user := range users {
		if user.Email == email {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	// compare password
	if !validatePassword(password, foundUser.HashedPassword) {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}
	// generate session token and set cookie

	sessionToken := generateToken(32)
	csrfToken := generateToken(32)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour), // 1 day
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	// store session in db

	session := Session{
		ID:           len(sessionDB) + 1,
		UserID:       foundUser.ID,
		SessionToken: sessionToken,
		CSRFToken:    csrfToken,
	}

	sessionDB[sessionToken] = session

	fmt.Fprintln(w, "Login successful")

}

func logout(w http.ResponseWriter, r *http.Request) {
	if err := Authorize(r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	sessionToken, _ := r.Cookie("session_token")
	delete(sessionDB, sessionToken.Value)

	// clear session
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	fmt.Fprintln(w, "Logout successful")
}

func profile(w http.ResponseWriter, r *http.Request) {

	if err := Authorize(r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	sessionToken, _ := r.Cookie("session_token")

	session, _ := sessionDB[sessionToken.Value]

	var foundUser *User
	for _, user := range users {
		if user.ID == session.UserID {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Email: %s\n", foundUser.Email)
}
