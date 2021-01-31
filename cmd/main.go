package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=UTC",
	// 	"root",
	// 	"root",
	// 	"localhost:3306",
	// 	"db_name",
	// )
	// db, err := gorm.Open("mysql", connString)
	// if err != nil {
	// 	panic(err)
	// }

	// mysqlRepo := mysql.NewRepository(db)
	// smtpSender := mail.NewSMTPSender(1234, "", "", "", "", "")

	// usecase := usecase.New(mysqlRepo, smtpSender)

	// userHandler := handler.NewUserHandler(usecase)
	// itemHandler := handler.NewItemHandler(usecase)

	// _ = userHandler
	// _ = itemHandler

}
