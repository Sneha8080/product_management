package models

type Product struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Availability bool    `json:"availability"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
	Quantity     int     `json:"quantity"`
}

const (
	PremiumCategory = "Premium"
	RegularCategory = "Regular"
	BudgetCategory  = "Budget"
)
