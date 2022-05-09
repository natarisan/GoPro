package domain

import(
	"database/sql"
	"os"
	"GOP/dto"
	"encoding/base64"
	_ "github.com/go-sql-driver/mysql"
	"github.com/natarisan/gop-libs/errs"
	"github.com/natarisan/gop-libs/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func(d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError){
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil{
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customers not found")
		} else {
			logger.Error("Error while querying customer table" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError){
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func(d CustomerRepositoryDb) PostImage(req dto.PostImageRequest) *errs.AppError {
	customer_id := req.CustomerId + "a"
	//フォルダ存在確認　フォルダ作成
	_, er := os.Stat(customer_id)
	if er != nil {
		if os.IsNotExist(er) {
			if err := os.Mkdir(customer_id, 0777); err != nil {
				return errs.NewUnexpectedError("Unexpected mkdir error" + err.Error())
			}
		}
	}
	//画像デコード
	base64Image := req.Image
	data, _ := base64.StdEncoding.DecodeString(base64Image)
	file, err2 := os.Create("./" + customer_id + "/" + customer_id + ".jpg")
	if err2 != nil {
		return errs.NewUnexpectedError("Unexpected create image error" + err2.Error())
	}
	defer file.Close()
	//画像保存
	file.Write(data)
	return nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb{
	return CustomerRepositoryDb{dbClient}
}