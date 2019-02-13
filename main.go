package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nbari/violetear"
)

var loginForm []byte = []byte(`
<html>
	<body>
		<form action="/login" method="POST">
			<p>
				<label for="username">Username</label>
				<input type="text" id="username" name="username" />
			</p>
			<p>
				<label for="password">Password</label>
				<input type="password" id="password" name="password"/>
			</p>
			<p>
				<input type="submit" />
			</p>
		</form>
	</body>
</html>
`)

func getLogin(w http.ResponseWriter, r *http.Request) {
	w.Write(loginForm)
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	fmt.Fprintf(w, "<b>username:</b> %s<br/><b>password:</b> %s", username, password)
}

func main() {
	router := violetear.New()
	router.LogRequests = true
	router.RequestID = "Request-ID"

	router.HandleFunc("/login", getLogin, "GET")
	router.HandleFunc("/login", postLogin, "POST")

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   7 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(srv.ListenAndServe())

}
