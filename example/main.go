package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	cookiemanager "github.com/Zhima-Mochi/go-authentication-service/service/cookieManager"
	sessionmanager "github.com/Zhima-Mochi/go-authentication-service/service/sessionManager"
)

func main() {
	ctx := context.Background()
	cookieManager := cookiemanager.NewCookieManager("session-id", 3600)
	sessionManager := sessionmanager.NewSessionManager()
	// example user data
	userData := map[string]interface{}{
		"userID": "david_recardo",
		"token":  "1234567890",
	}

	// Create new session
	session, _ := sessionManager.CreateSession(ctx, userData)

	http.HandleFunc("/setCookie", func(w http.ResponseWriter, r *http.Request) {
		cookieManager.SetCookie(w, session.GetID())
		fmt.Fprintf(w, "cookie set")
	})

	http.HandleFunc("/getCookie", func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("session-id")
		if cookie == nil {
			fmt.Fprintf(w, "cookie not found")
			return
		}
		fmt.Fprintf(w, "Cookie: %s=%s", "session-id", cookie.Value)
	})

	http.HandleFunc("/deleteCookie", func(w http.ResponseWriter, r *http.Request) {
		cookieManager.DeleteCookie(w)
		fmt.Fprintf(w, "cookie deleted")
	})

	http.HandleFunc("/getSession", func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("session-id")
		if cookie == nil {
			fmt.Fprintf(w, "cookie not found")
			return
		}
		session, err := sessionManager.GetSession(ctx, cookie.Value)
		if err != nil {
			fmt.Fprintf(w, "session not found")
			return
		}
		fmt.Fprintf(w, "Session: %s", session)
	})

	server := &http.Server{
		Addr:        ":8080",
		ReadTimeout: 5 * time.Second,
	}
	fmt.Println("Server started on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
