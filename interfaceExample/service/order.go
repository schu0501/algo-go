package service

import "fmt"

type DataService interface {
	Save(data string) error
}

type OrderService struct {
	Saver DataService
}

func (s *OrderService) CreateOrder(orderId string) {
	fmt.Println("CreateOrder....")
	err := s.Saver.Save(orderId)
	if err != nil {
		return
	}
}
