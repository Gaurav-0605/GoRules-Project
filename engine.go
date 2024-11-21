package main

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

