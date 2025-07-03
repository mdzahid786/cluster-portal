package middleware

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/mdzahid786/cluster-portal/internal/config"
)

type contextKey string

const UserContextKey  = contextKey("user")

func AuthMiddleware(users []config.User, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("autho: ", users)
		header := r.Header.Get("Authorization")
		if header == "" || !strings.HasPrefix(header, "Basic ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		
		payload, _ := base64.StdEncoding.DecodeString(header[len("Basic "):])
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		username, password := pair[0], pair[1]
		
		for _, user := range users {
			if user.Username == username && user.Password == password {
				// Add user to context
				ctx := context.WithValue(r.Context(), UserContextKey , user)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func GetAuthenticatedUser(r *http.Request) (config.User, bool) {
	user, ok := r.Context().Value(UserContextKey ).(config.User)
	return user, ok
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := GetAuthenticatedUser(r)
		if !ok || user.Role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}



