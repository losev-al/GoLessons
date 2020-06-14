package v1

// Order describes an order in an online store
type Order struct {
	Id int
	State State
	Goods []int
}
