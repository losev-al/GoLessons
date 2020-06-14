package warehouseservice

import (
	"fmt"
	models "github.com/losev-al/GoLessons/pkg/api/v1"
)

// Service describe warehouse
type Service interface {
	GetRest(id int) (int, error)
	Book(id int) error
	UnBook(id int) error
	WriteOf(id int) error
	Add(id int, count int)
}

type service struct {
	warehouse map[int]*models.SavedProduct
}

// GetRest ...
func (s *service) GetRest(id int) (int, error) {
	goods, err := s.getGoods(id)
	if err != nil {
		return 0, err
	}
	return goods.StockBalance - goods.Booked, nil
}

// Book book a good if can
func (s *service) Book(id int) (err error) {
	goods, err := s.getGoods(id)
	if err != nil {
		return
	}
	if goods.StockBalance == goods.Booked {
		err = fmt.Errorf("goods with id = %v is out of stock", id)
		return
	}
	goods.Booked++
	return
}

// UnBook removes the good from the reservation
func (s *service) UnBook(id int) (err error) {
	goods, err := s.getGoods(id)
	if err != nil {
		return
	}
	if goods.Booked == 0 {
		err = fmt.Errorf("don't exists booking goods with id = %v", id)
		return
	}
	goods.Booked--
	return
}

// WriteOf performs the shipment of the goods
func (s *service) WriteOf(id int) (err error) {
	goods, err := s.getGoods(id)
	if err != nil {
		return
	}
	if goods.Booked > 0 {
		err = fmt.Errorf("don't exists booking goods with id = %v", id)
		return
	}
	goods.StockBalance--
	goods.Booked--
	return
}

// Add ...
func (s *service) Add(id int, count int) {
	goods, exists := s.warehouse[id]
	if !exists {
		goods = &models.SavedProduct{
			Id:           id,
			StockBalance: 0,
			Booked:       0,
		}
		s.warehouse[id] = goods
	}
	goods.StockBalance += count
	return
}

func (s *service) getGoods(id int) (goods *models.SavedProduct, err error) {
	goods, exists := s.warehouse[id]
	if !exists {
		err = fmt.Errorf("goods with id = %v don't exists", id)
	}
	return
}

// NewService ...
func NewService() Service {
	return &service{
		warehouse: make(map[int]*models.SavedProduct),
	}
}

