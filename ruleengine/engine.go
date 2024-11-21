package ruleengine

import "GoRules-Project/model"

// EvaluateRules applies the rules based on the pincode
func EvaluateRules(input model.RuleInput) model.RuleOutput {
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

	return model.RuleOutput{
		DiscountPercent: discount,
		Message:         message,
	}
}
