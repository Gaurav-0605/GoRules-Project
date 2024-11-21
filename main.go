package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// Input data structure
type RuleInput struct {
	Pincode string `json:"pincode"`
}

// Output data structure
type RuleOutput struct {
	DiscountPercent float64 `json:"discountPercent"`
	Message         string  `json:"message"`
}

// Sample pincode-based discount rules
var pincodeDiscounts = map[string]float64{
	"110001": 10, // Example pincode with 10% discount
	"110002": 15, // Example pincode with 15% discount
	"110003": 5,  // Example pincode with 5% discount
}

// Rule evaluation logic based only on pincode
func evaluateRules(input RuleInput) RuleOutput {
	var discount float64
	var message string

	// Check if the pincode is eligible for a discount
	discount = pincodeDiscounts[input.Pincode]

	if discount > 0 {
		message = "You get a discount for your pincode!"
	} else {
		// No pincode discount
		discount = 0
		message = "No discount available for this pincode."
	}

	return RuleOutput{
		DiscountPercent: discount,
		Message:         message,
	}
}

// Handler for the API
func ruleHandler(w http.ResponseWriter, r *http.Request) {
	var input RuleInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Evaluate rules based on pincode
	result := evaluateRules(input)

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
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

func main() {
	// Setting up the router
	router := mux.NewRouter()
	router.HandleFunc("/evaluate", ruleHandler).Methods("POST")

	// Enable CORS
	http.ListenAndServe(":8080", enableCors(router))
}

