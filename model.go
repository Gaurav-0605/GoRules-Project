package main

// Input data structure
type RuleInput struct {
	Pincode string `json:"pincode"`
}

// Output data structure
type RuleOutput struct {
	DiscountPercent float64 `json:"discountPercent"`
	Message         string  `json:"message"`
}

