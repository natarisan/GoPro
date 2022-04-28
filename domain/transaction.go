package domain

import(
	"GOP/dto"
	"github.com/natarisan/gop-libs/errs"
)

type Transaction struct{
	TransId   string  `db:"transaction_id"`
	AccountId string  `db:"account_id"`
	Amount    float64 
	TransType string  `db:"transaction_type"`
	TransDate string  `db:"transaction_date"`
}

func(t Transaction) ToResultTransResponseDto() dto.TransResponse{
	return dto.TransResponse{t.TransId, t.AccountId, t.Amount, t.TransType, t.TransDate}
}

type TransRepository interface{
	Update(Transaction)(*Transaction, *errs.AppError)
}

func (t Transaction) ToDto() dto.TransResponse {
	return dto.TransResponse{
		TransId:   	t.TransId,
		AccountId:  t.AccountId,
		Amount:     t.Amount,
		TransType:  t.TransType,
		TransDate:  t.TransDate,
	}
}