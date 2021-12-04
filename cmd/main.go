package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/vitorath/clean-architecture-golang/adapter/repository"
	"github.com/vitorath/clean-architecture-golang/adapter/repository/fixtures"
	usecase "github.com/vitorath/clean-architecture-golang/usecase/process_transaction"
)

func main() {
	migrationsDir := os.DirFS("adapter/repository/fixtures/sql")
	db := fixtures.Up(migrationsDir)
	defer db.Close()

	repo := repository.NewTransactionRepositoryDb(db)
	processTransactionUseCase := usecase.NewProcessTransaction(repo)

	input := usecase.TransactionDtoInput{
		ID:        "1",
		AccountID: "",
		Amount:    1100,
	}

	output, err := processTransactionUseCase.Execute(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	outputJson, _ := json.Marshal(output)
	fmt.Println(string(outputJson))
}
