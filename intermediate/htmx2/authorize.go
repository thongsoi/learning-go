func requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the user is authenticated
		// In this example, we check if the session_token cookie is set
		// In a real application, you should use a secure way to check if the user is authenticated
		cookie, err := r.Cookie("session_token")
		if err != nil || cookie.Value == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the user is authenticated, call the original handler
		next.ServeHTTP(w, r)
	})
}