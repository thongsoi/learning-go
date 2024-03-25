http.SetCookie(w, &http.Cookie{
	Name:  "session_token",
	Value: "some_token",
	Path:  "/",
})