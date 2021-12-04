package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_entity "github.com/vitorath/clean-architecture-golang/entity/mock"
)

func TestProcessTransactionWhenItsValid(t *testing.T) {
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

	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "approved", "").Return(nil)

	output, err := processTransaction.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransactionWhenItsInvalid(t *testing.T) {
	input := TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    1001,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       "rejected",
		ErrorMessage: "you don't have limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)
	processTransaction := NewProcessTransaction(repositoryMock)

	errorMessage := errors.New("you don't have limit for this transaction")
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "rejected", errorMessage.Error()).Return(nil)

	output, err := processTransaction.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
