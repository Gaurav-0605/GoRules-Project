package model

// RuleInput represents the input structure for the rule
type RuleInput struct {
	Pincode string `json:"pincode"`
}

// RuleOutput represents the output structure for the rule evaluation
type RuleOutput struct {
	DiscountPercent float64 `json:"discountPercent"`
	Message         string  `json:"message"`
}
