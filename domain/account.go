package domain

import(
	"GOP/dto"
	"GOP/errs"
)

type Account struct{
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string	`db:"status"`
}

//dtoを参照してdomainでレスポンス生成→serviceへ飛ばす
func(a Account) ToNewAccountResponseDto() dto.NewAccountResponse{
	return dto.NewAccountResponse{a.AccountId}
}

//このドメインのリポジトリはこのinterfaceを持つという宣言
type AccountRepository interface{
	Save(Account)(*Account, *errs.AppError)
	SaveTransaction(transaction Transaction)(*Transaction, *errs.AppError)
	FindBy(accountId string)(*Account, *errs.AppError)
}

func(a Account) CanWithdraw(amount float64) bool{
	if a.Amount < amount {
		return false
	}
	return true
}