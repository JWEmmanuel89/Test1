// File Name: main.go
package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Create handler function for home
func home(w http.ResponseWriter, r *http.Request) {
	// Write some bio information to page
	w.Write([]byte("Hello, my name is Joshua Emmanuel\n"))
	w.Write([]byte("D.O.B: January 16, 1989.\n"))
	w.Write([]byte("I am currently a member of the Belize Coast Guard.\n"))
	w.Write([]byte("I have 2 sons whom with I enjoy playing video games with and teaching them sports.\n"))
}

// Create handler function for greeting
func greeting(w http.ResponseWriter, r *http.Request) {
	// Write a greeting displaying current time and day on page
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
	// Write a random quote from map to page
	// Create map with random qoutes
	quotes := make(map[string]string)
	quotes["R1"] = "In theory, there is no difference between theory and practice. In practice there is."
	quotes["R2"] = "The problem is not that you feel like shit. The problem is that you are behaving like shit and then blaming it on your feelings."
	quotes["R3"] = "Live as if you were living already for the second time and as if you had acted the first time as wrongly as you are about to act now!"
	quotes["R4"] = "There is no way to happiness -- happiness is the way."
	quotes["R5"] = "How much time are you spending deciding what to spend time on?"
	quotes["R6"] = "I cannot judge what he did, because I did not have his information."
	quotes["R7"] = "If you want to kill any idea in the world today, get a committee working on it."
	quotes["R8"] = "The real problem of humanity is the following: we have paleolithic emotions; medieval institutions; and god-like technology."
	quotes["R9"] = "The most effective form of learning is practice, not planning."
	quotes["R10"] = "Weak people revenge. Strong people forgive. Intelligent people ignore."

	// Print random quote
	k := rand.Intn(len(quotes))
	i := 0
	for _, x := range quotes {
		if i == k {
			w.Write([]byte(x))
		}
		i++
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/greeting", greeting)
	mux.HandleFunc("/random", random)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
