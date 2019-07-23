package main

import (
	"fmt"

	"github.com/lfourky/go-transaction-management/pkg/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lfourky/go-transaction-management/pkg/repository/mysql"
	"github.com/lfourky/go-transaction-management/pkg/service/mail"
	"github.com/lfourky/go-transaction-management/pkg/service/usecase"
)

func main() {

	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=UTC",
		"root",
		"root",
		"localhost:3306",
		"db_name",
	)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}

	mysqlRepo := mysql.NewRepository(db)
	smtpSender := mail.NewSMTPSender(1234, "", "", "", "", "")

	usecase := usecase.New(mysqlRepo, smtpSender)

	userHandler := handler.NewUserHandler(usecase)
	itemHandler := handler.NewItemHandler(usecase)

	_ = userHandler
	_ = itemHandler

}
