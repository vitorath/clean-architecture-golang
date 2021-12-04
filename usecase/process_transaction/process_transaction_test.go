package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_entity "github.com/vitorath/clean-architecture-golang/entity/mock"
)

func newSut(t *testing.T) (TransactionDtoInput, TransactionDtoOutput, *ProcessTransaction, *mock_entity.MockTransactionRepository) {
	input := TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    200,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       "approved",
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)
	processTransaction := NewProcessTransaction(repositoryMock)
	return input, expectedOutput, processTransaction, repositoryMock
}

func TestProcessTransactionWhenItsValid(t *testing.T) {
	input, expectedOutput, processTransaction, repositoryMock := newSut(t)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "approved", "").Return(nil)

	output, err := processTransaction.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransactionWhenItsInvalid(t *testing.T) {
	input, expectedOutput, processTransaction, repositoryMock := newSut(t)
	input.Amount = 1001
	expectedOutput.Status = "rejected"
	expectedOutput.ErrorMessage = "you don't have limit for this transaction"
	errorMessage := errors.New("you don't have limit for this transaction")
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "rejected", errorMessage.Error()).Return(nil)

	output, err := processTransaction.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
