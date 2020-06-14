package orderservice

import (
	"fmt"
	models "github.com/losev-al/GoLessons/pkg/api/v1"
)

type Service interface {
	CreateOrder(goods []int) *models.Order
	GetOrder(id int) (*models.Order, error)
	GetOrderList(state models.State) []*models.Order
}

type service struct {
	orders map[int]*models.Order
}

// CreateOrder ...
func (s *service) CreateOrder(goods []int) (order *models.Order) {
	var max int = 0
	for id, _ := range s.orders {
		if id > max {
			max = id
		}
	}
	order = &models.Order{
		Id: max + 1,
		State: models.New,
		Goods: goods,
	}
	s.orders[order.Id] = order
	return order
}

// GetOrder ...
func (s *service) GetOrder(id int) (order *models.Order, err error) {
	order, exists := s.orders[id]
	if !exists {
		err = fmt.Errorf("order with id = %v don't exists", id)
	}
	return
}

// GetOrderList ...
func (s *service) GetOrderList(state models.State) (result []*models.Order) {
	for _, order := range s.orders {
		if order.State == state {
			result = append(result, order)
		}
	}
	return
}

// NewService ...
func NewService() Service {
	return &service{
		orders: make(map[int]*models.Order),
	}
}

