package facade

import (
	"fmt"
	models "github.com/losev-al/GoLessons/pkg/api/v1"
)

type orderService interface {
	CreateOrder(goods []int) *models.Order
	GetOrder(id int) (*models.Order, error)
}

type warehouse interface {
	GetRest(id int) (int, error)
	Book(id int) error
	UnBook(id int) error
	WriteOf(id int) error
}

type OnlineMarket interface {
	CreateOrder(goods []int) (*models.Order, error)
	ConfirmOrder(id int) error
}

type onlineMarket struct {
	OrderService orderService
	Warehouse warehouse
}

// CreateOrder ...
func (o *onlineMarket) CreateOrder(goods []int) (*models.Order, error) {
	var notEnoughGoods []int
	for _, goodsId := range goods {
		rest, err := o.Warehouse.GetRest(goodsId)
		if rest == 0 || err != nil {
			notEnoughGoods = append(notEnoughGoods, goodsId)
		}
	}
	if len(notEnoughGoods) > 0 {
		return nil, fmt.Errorf("there are no products with id: %v in stock", notEnoughGoods)
	}
	order := o.OrderService.CreateOrder(goods)
	return order, nil
}

// ConfirmOrder book goods in order
func (o *onlineMarket) ConfirmOrder(id int) error {
	order, err := o.OrderService.GetOrder(id)
	if err != nil {
		return err
	}
	for bookIndex, goods := range order.Goods {
		err = o.Warehouse.Book(goods)
		if err != nil {
			for unBookIndex := bookIndex; unBookIndex >= 0; unBookIndex-- {
				o.Warehouse.UnBook(order.Goods[unBookIndex])
			}
			break
		}
	}
	return err
}

// NewOnlineMarket ...
func NewOnlineMarket(o orderService, w warehouse) OnlineMarket {
	return &onlineMarket{
		OrderService: o,
		Warehouse: w,
	}
}
