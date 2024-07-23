package service_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/lucaszatta/frete-rapido-v2/internal/quote/models"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/repository/mocks"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/service"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	//Act
	serviceUT := service.New(nil)

	//ASSERT
	assert.NotNil(t, serviceUT)
	assert.Equal(t, "*service.quoteService", fmt.Sprintf("%T", serviceUT))
}

func TestService_Create(t *testing.T) {
	//Arrange

	qs := &[]models.Quote{
		{
			Name:     "foo",
			Service:  "service",
			Deadline: 3,
			Price:    10.0,
		},
	}

	errToReturn := errors.New("any_error")

	t.Run("success", func(t *testing.T) {
		repository := mocks.NewQuoteRepository(t)
		repository.On("Create", qs).Return(nil)

		serviceUT := service.New(repository)
		//Act
		q, err := serviceUT.Create(qs)

		//Assert
		assert.Nil(t, err)
		assert.Equal(t, q, qs)
		repository.AssertNumberOfCalls(t, "Create", 1)
		repository.AssertCalled(t, "Create", qs)
	})

	t.Run("error", func(t *testing.T) {
		repository := mocks.NewQuoteRepository(t)
		repository.On("Create", qs).Return(errToReturn)

		serviceUT := service.New(repository)
		//Act
		q, err := serviceUT.Create(qs)

		//Assert
		assert.Nil(t, q)
		assert.Equal(t, err, errToReturn)
		repository.AssertNumberOfCalls(t, "Create", 1)
		repository.AssertCalled(t, "Create", qs)
	})
}

func TestService_GetLastQuotes(t *testing.T) {
	//Arrange
	lastqs := &[]models.LastQuotes{
		{
			Name:     "carrier",
			Count:    2,
			PriceSum: 15.0,
			PriveAvg: 7.0,
		},
		{
			Name:     "carrier 2",
			Count:    2,
			PriceSum: 40.0,
			PriveAvg: 20.0,
		},
		{
			Name:     "carrier 3",
			Count:    1,
			PriceSum: 10.0,
			PriveAvg: 10.0,
		},
	}

	errToReturn := errors.New("any_error")
	lastqparams := "5"

	t.Run("success all", func(t *testing.T) {
		repository := mocks.NewQuoteRepository(t)
		repository.On("GetLastQuotes", lastqparams).Return(lastqs, nil)

		serviceUT := service.New(repository)
		//Act
		q, err := serviceUT.GetLastQuotes(lastqparams)

		//Assert
		assert.Nil(t, err)
		assert.Equal(t, q, lastqs)
		repository.AssertNumberOfCalls(t, "GetLastQuotes", 1)
		repository.AssertCalled(t, "GetLastQuotes", lastqparams)
	})

	t.Run("error", func(t *testing.T) {
		repository := mocks.NewQuoteRepository(t)
		repository.On("GetLastQuotes", lastqparams).Return(nil, errToReturn)

		serviceUT := service.New(repository)
		//Act
		q, err := serviceUT.GetLastQuotes(lastqparams)

		//Assert
		assert.Nil(t, q)
		assert.Equal(t, err, errToReturn)
		repository.AssertNumberOfCalls(t, "GetLastQuotes", 1)
		repository.AssertCalled(t, "GetLastQuotes", lastqparams)
	})
}

func TestService_GetMaxMinQuotes(t *testing.T) {
	//Arrange
	maxq := &models.Quote{
		Name:     "expensive",
		Service:  "service",
		Deadline: 2,
		Price:    50.5,
	}

	minq := &models.Quote{
		Name:     "cheapest",
		Service:  "service",
		Deadline: 10,
		Price:    15,
	}

	errToReturn := errors.New("any_error")

	t.Run("success all", func(t *testing.T) {
		repository := mocks.NewQuoteRepository(t)
		repository.On("GetMaxMinQuotes").Return(maxq, minq, nil)

		serviceUT := service.New(repository)
		//Act
		max, min, err := serviceUT.GetMaxMinQuotes()

		//Assert
		assert.Nil(t, err)
		assert.Equal(t, max, maxq)
		assert.Equal(t, min, minq)
		repository.AssertNumberOfCalls(t, "GetMaxMinQuotes", 1)
		repository.AssertCalled(t, "GetMaxMinQuotes")
	})

	t.Run("error", func(t *testing.T) {
		repository := mocks.NewQuoteRepository(t)
		repository.On("GetMaxMinQuotes").Return(nil, nil, errToReturn)

		serviceUT := service.New(repository)
		//Act
		max, min, err := serviceUT.GetMaxMinQuotes()

		//Assert
		assert.Nil(t, max)
		assert.Nil(t, min)
		assert.Equal(t, err, errToReturn)
		repository.AssertNumberOfCalls(t, "GetMaxMinQuotes", 1)
		repository.AssertCalled(t, "GetMaxMinQuotes")
	})
}
