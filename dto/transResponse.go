package dto

type TransResponse struct {
	TransId       string `json:"transaction_id"`
	AccountId     string `json:"account_id"`
	Amount 	      float64 `json:"amount"`           
	TransType     string  `json:"transaction_type`
	TransDate     string `json:"transaction_date`
}