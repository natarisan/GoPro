package domain

import ("github.com/natarisan/gop-libs/errs"
		"GOP/dto"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func(c Customer) statusAsText() string{
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

//Customer構造体は、このメソッドを持つ。
func (c Customer) ToDto() dto.CustomerResponse{
	
	return dto.CustomerResponse{
		Id: c.Id,
		Name: c.Name,
		City: c.City,
		Zipcode: c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status: c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}