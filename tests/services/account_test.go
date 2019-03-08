package test

import (
	"testing"

	"github.com/brunoleonel/payment/app/http/resources"

	"github.com/brunoleonel/payment/app/services"

	mocks "github.com/brunoleonel/payment/tests/mocks/repositories"
	"github.com/jinzhu/gorm"

	"github.com/stretchr/testify/assert"
)

func TestShouldFindAccount(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	repositoryMock := mocks.NewAccountRepositoryMock(db)
	service := services.NewAccountService(repositoryMock)
	result, _ := service.Find(1)

	assert.NotNil(result)
	assert.Equal(int64(1), result.AccountID)
	assert.Equal(float32(100), result.AvailableCreditLimit)
	assert.Equal(float32(100), result.AvailableWithdrawalLimit)
}

func TestShouldntFindAccount(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	repositoryMock := mocks.NewAccountRepositoryMock(db)
	service := services.NewAccountService(repositoryMock)
	_, err := service.Find(2)

	assert.Error(err)
	assert.Equal(err.Code, 404)
}

func TestShouldCreateAccount(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	repositoryMock := mocks.NewAccountRepositoryMock(db)
	service := services.NewAccountService(repositoryMock)

	resource := &resources.Account{
		AvailableCreditLimit: resources.AvailableCreditLimit{
			Amount: 100,
		},
		AvailableWithdrawalLimit: resources.AvailableWithdrawalLimit{
			Amount: 100,
		},
	}

	result := service.Create(resource)

	assert.Equal(int64(1), result.ID)
	assert.Equal(float32(100), result.AvailableCreditLimit.Amount)
	assert.Equal(float32(100), result.AvailableWithdrawalLimit.Amount)
}

func TestShouldUpdateAccount(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	repositoryMock := mocks.NewAccountRepositoryMock(db)
	service := services.NewAccountService(repositoryMock)

	resource := &resources.Account{
		ID: 1,
		AvailableCreditLimit: resources.AvailableCreditLimit{
			Amount: 100,
		},
		AvailableWithdrawalLimit: resources.AvailableWithdrawalLimit{
			Amount: 100,
		},
	}

	result, _ := service.Update(resource.ID, resource)

	assert.Equal(int64(1), result.ID)
	assert.Equal(float32(200), result.AvailableCreditLimit.Amount)
	assert.Equal(float32(200), result.AvailableWithdrawalLimit.Amount)
}

func TestShouldGetAccountLimits(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	repositoryMock := mocks.NewAccountRepositoryMock(db)
	service := services.NewAccountService(repositoryMock)

	result := service.GetLimits()
	assert.NotEmpty(result)
}

func TestShouldAdjustAccountCreditLimit(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	repositoryMock := mocks.NewAccountRepositoryMock(db)
	service := services.NewAccountService(repositoryMock)

	err := service.AdjustAccountLimit(1, 100, 1)
	assert.Nil(err)
}

func TestShouldAdjustAccountWithdrawalLimit(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	repositoryMock := mocks.NewAccountRepositoryMock(db)
	service := services.NewAccountService(repositoryMock)

	err := service.AdjustAccountLimit(1, 100, 3)
	assert.Nil(err)
}

func TestShouldntAdjustAccountCreditLimit(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	repositoryMock := mocks.NewAccountRepositoryMock(db)
	service := services.NewAccountService(repositoryMock)

	err := service.AdjustAccountLimit(2, 100, 1)
	assert.Error(err)
	assert.Equal(404, err.Code)
}

func TestShouldntAdjustAccountWithdrawalLimit(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	repositoryMock := mocks.NewAccountRepositoryMock(db)
	service := services.NewAccountService(repositoryMock)

	err := service.AdjustAccountLimit(2, 100, 3)
	assert.Error(err)
	assert.Equal(404, err.Code)
}
