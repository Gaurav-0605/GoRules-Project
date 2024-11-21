package main

import (
	"encoding/json"
	"net/http"
)

// Sample pincode-based discount rules
var pincodeDiscounts = map[string]float64{
	"110001": 10, // Example pincode with 10% discount
	"110002": 15, // Example pincode with 15% discount
	"110003": 5,  // Example pincode with 5% discount
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
