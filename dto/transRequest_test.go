package dto //ファイル名_test.goでつける　go testコマンドで実行

import ("testing"
		"net/http"
)

//Test〜で始める
func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawal(t *testing.T){
	request := TransRequest{ //Arrange
		TransactionType: "invalid transaction type",
	}
		//Action
	appError := request.Validate2()

		//assert
	if appError.Message != "Transaction type should be withdrawal or deposit" {
		t.Error("Invalid message while testing transaction type")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction type")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T){
	request := TransRequest{
		TransactionType: "deposit",
		Amount: -100,
	}

	appError := request.Validate2()

	if appError.Message != "Amount can not be negative value" {
		t.Error("Invalid message while validating request")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while validating amount")
	}
}