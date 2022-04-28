package service

import(
	"GOP/dto"
	"github.com/natarisan/gop-libs/errs"
	"GOP/domain"
	"time"
)

type AccountService interface{
	NewAccount(dto.NewAccountRequest)(*dto.NewAccountResponse,*errs.AppError)
	MakeTransaction(request dto.TransRequest)(*dto.TransResponse, *errs.AppError)
}

type DefaultAccountService struct{
	repo domain.AccountRepository
}

func(s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError){
	err := req.Validate()
	if err != nil{
		return nil, err
	}
	a := domain.Account{
		AccountId:    "",
		CustomerId:   req.CustomerId,
		OpeningDate:  time.Now().Format("2006-01-02 15:04:05"),
		AccountType:  req.AccountType,
		Amount:       req.Amount,
		Status:       "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService{
	return DefaultAccountService{repo}
}

func(s DefaultAccountService) MakeTransaction(req dto.TransRequest)(*dto.TransResponse, *errs.AppError){
	err := req.Validate2()
	if err != nil{
		return nil, err
	}
	if req.TransactionType == "withdrawal" {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil{
			return nil, err
		}
		if !account.CanWithdraw(req.Amount){
			return nil, errs.NewValidationError("Insufficient balance in your account")
		}
	}	
	t := domain.Transaction{
		AccountId: req.AccountId,
		Amount:    req.Amount,
		TransType: req.TransactionType,
		TransDate: time.Now().Format("2006-01-02 15:04:05"),
	}	
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}