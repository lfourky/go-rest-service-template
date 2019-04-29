package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lfourky/go-transaction-management/pkg/repository/mysql"
	"github.com/lfourky/go-transaction-management/pkg/service"
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

	// Pass these to handlers or whatevs
	userService := service.NewUser(mysqlRepo)
	itemService := service.NewItem(mysqlRepo)
	cartService := service.NewCart(mysqlRepo)

	_, _ = userService.DemonstrateTransaction()
	_, _ = itemService.DemonstrateTransaction()
	_, _ = cartService.DemonstrateTransaction()
}
