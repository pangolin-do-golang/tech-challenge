package order

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/product"
)

type Service struct {
	OrderRepository        IOrderRepository
	OrderProductRepository IOrderProductRepository
	CartService            cart.IService
	ProductService         *product.Service
}

func NewOrderService(repo IOrderRepository, orderProductRepository IOrderProductRepository, cartService cart.IService, productService *product.Service) IOrderService {
	return &Service{
		OrderRepository:        repo,
		OrderProductRepository: orderProductRepository,
		CartService:            cartService,
		ProductService:         productService,
	}
}

func (s *Service) Get(id uuid.UUID) (*Order, error) {
	return s.OrderRepository.Get(id)
}

func (s *Service) GetAll() ([]Order, error) {
	return s.OrderRepository.GetAll()
}

func (s *Service) Create(clientID uuid.UUID) (*Order, error) {
	c, err := s.CartService.GetFullCart(clientID)
	if err != nil {
		return nil, err
	}

	if len(c.Products) == 0 {
		return nil, fmt.Errorf("empty cart")
	}

	order := &Order{
		ClientID: clientID,
		Status:   Status.Created,
	}

	o, err := s.OrderRepository.Create(order)
	if err != nil {
		return nil, err
	}

	var total float64
	for _, p := range c.Products {
		stockProduct, err := s.ProductService.GetByID(p.ProductID)
		if err != nil {
			return nil, err
		}

		productTotal := stockProduct.Price * float64(p.Quantity)

		orderProduct := &Product{
			ClientID:  clientID,
			ProductID: p.ProductID,
			Quantity:  p.Quantity,
			Comments:  p.Comments,
			Total:     productTotal,
		}

		err = s.OrderProductRepository.Create(context.Background(), o.ID, orderProduct)
		if err != nil {
			return nil, err
		}

		total += productTotal
	}

	o.TotalAmount = total
	o.Status = Status.Preparing
	err = s.OrderRepository.Update(o)
	if err != nil {
		return nil, err
	}

	if err = s.CartService.Cleanup(clientID); err != nil {
		return nil, err
	}

	return o, nil
}
