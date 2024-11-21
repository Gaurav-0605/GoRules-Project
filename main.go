package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"GoRules-Project/ruleengine"
)

func main() {
	// Setting up the router
	router := mux.NewRouter()
	router.HandleFunc("/evaluate", ruleengine.RuleHandler).Methods("POST")

	// Enable CORS
	http.ListenAndServe(":8080", enableCors(router))
}

// Enable CORS for the server (if necessary)
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
