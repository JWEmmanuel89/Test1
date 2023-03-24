// File Name: main.go
package main

import (
	"log"
	"net/http"
	"time"
)

// Create handler function for home
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, my name is Joshua Emmanuel\n"))
	w.Write([]byte("D.O.B: January 16, 1989.\n"))
	w.Write([]byte("I am currently a member of the Belize Coast Guard.\n"))
	w.Write([]byte("I have 2 sons whom I enjoy playing video games with and teaching them sports.\n"))
}

// Create handler function for greeting
func greeting(w http.ResponseWriter, r *http.Request) {
	weekday := time.Now().Weekday()
	now := time.Now()
	w.Write([]byte("The time is curently "))
	w.Write([]byte(now.Format(time.Kitchen)))
	w.Write([]byte("\n"))
	w.Write([]byte("Enjoy the rest of your "))
	w.Write([]byte(weekday.String()))
}

// Create handler function for random
func random(w http.ResponseWriter, r *http.Request) {
}

func main() {
	webapp := http.NewServeMux()
	webapp.HandleFunc("/", home)
	webapp.HandleFunc("/greeting", greeting)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", webapp)
	log.Fatal(err)
}
