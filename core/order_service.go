package core

import "errors"

type OrderService interface {
	CreateOrder(order Order) error
}

type OrderServiceIml struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &OrderServiceIml{repo: repo}
}

func (s *OrderServiceIml) CreateOrder(order Order) error {
	if order.Total < 0 {
		return errors.New("total must be positive")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}

	return nil
}
