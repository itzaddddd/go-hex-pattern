package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockOrderRepo struct {
	saveFunc func(order Order) error
}

func (m *mockOrderRepo) Save(order Order) error {
	return m.saveFunc(order)
}

func TestCreateOrder(t *testing.T) {
	repo := &mockOrderRepo{
		saveFunc: func(order Order) error {
			return nil
		},
	}
	service := NewOrderService(repo)

	testcases := []struct {
		name   string
		total  float64
		expect error
	}{
		{
			name:   "create order successful",
			total:  10,
			expect: nil,
		},
		{
			name:   "create order failed, total < 0",
			total:  -1,
			expect: errors.New("total must be positive"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			order := Order{Total: tc.total}
			err := service.CreateOrder(order)

			assert.Equal(t, err, tc.expect)

		})
	}

	t.Run("repository error", func(t *testing.T) {
		repo := &mockOrderRepo{
			saveFunc: func(order Order) error {
				return errors.New("database error")
			},
		}
		service := NewOrderService(repo)
		order := Order{Total: 10}
		err := service.CreateOrder(order)

		assert.Error(t, err)

	})
}
