package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stephanvebrian/e-pocket/pocket-engine/handler"
	accountLogicImpl "github.com/stephanvebrian/e-pocket/pocket-engine/logic/account"
	transferHistoryLogicImpl "github.com/stephanvebrian/e-pocket/pocket-engine/logic/transactionhistory"
	transferLogicImpl "github.com/stephanvebrian/e-pocket/pocket-engine/logic/transfer"
	userLogicImpl "github.com/stephanvebrian/e-pocket/pocket-engine/logic/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbConn, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN: "postgres://epocket:secret@postgres:5432/epocket?sslmode=disable",
		}),
		&gorm.Config{},
	)
	if err != nil {
		fmt.Printf("error when trying to connect to database: %+v\n", err)
	}

	transferLogic := transferLogicImpl.New(transferLogicImpl.TransferLogicOptions{
		DB: dbConn,
	})
	accountLogic := accountLogicImpl.New(accountLogicImpl.AccountLogicOptions{
		DB: dbConn,
	})
	userLogic := userLogicImpl.New(userLogicImpl.UserLogicOptions{
		DB: dbConn,
	})
	transactionHistoryLogic := transferHistoryLogicImpl.New(transferHistoryLogicImpl.TransactionHistoryLogicOptions{
		DB: dbConn,
	})

	handler := handler.New(handler.HandlerOptions{
		TransferLogic:           transferLogic,
		AccountLogic:            accountLogic,
		UserLogic:               userLogic,
		TransactionHistoryLogic: transactionHistoryLogic,
	})

	handler.RegisterRoutes()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3010", handler.GetRouter()))

	fmt.Println("Server started on port 3010")
}
