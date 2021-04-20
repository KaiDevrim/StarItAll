package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	// Root route
	// Simply returns a link to the login route

	//rootHandler
	http.HandleFunc("/", rootHandler)

	// Login route
	http.HandleFunc("/login/github/", githubLoginHandler)

	// Github callback
	http.HandleFunc("/login/github/callback", githubCallbackHandler)

	// Route where the authenticated user is redirected to
	http.HandleFunc("/loggedin", func(w http.ResponseWriter, r *http.Request) {
		loggedinHandler(w, r, "")
	})

	// Listen and serve on port from env
	fmt.Println("[ UP ON PORT " + os.Getenv("PORT") + " ]")
	log.Panic(
		http.ListenAndServe(":"+os.Getenv("PORT"), nil),
	)
}

func getImportantInfo() {
	user := os.Args[1]
	req, reqerr := http.NewRequest("GET", "https://api.github.com/"+user, nil)
}

func star(repos string) {

}
