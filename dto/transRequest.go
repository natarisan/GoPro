package dto

import(
	"strings"
	"GOP/errs"
)

type TransRequest struct{
	AccountId       string `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string `json:"transaction_type"`
	TransactionDate string `json:"transaction_date"`
	CustomerId      string `json:"-"`
}

func (r TransRequest) Validate2() *errs.AppError {
	if r.Amount < 0 {
		return errs.NewValidationError("Amount can not be negative value")
	}
	if strings.ToLower(r.TransactionType) != "withdrawal" && strings.ToLower(r.TransactionType) != "deposit"{
		return errs.NewValidationError("Transaction type should be withdrawal or deposit")
	}
	return nil
}