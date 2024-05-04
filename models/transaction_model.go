package models

// Transaction represents a banking transaction
type Transaction struct {
	ID          uint    `json:"id"`
	AccountID   uint    `json:"account_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	// Add more fields as needed
}
