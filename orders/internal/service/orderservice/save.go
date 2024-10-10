package orderservice

import (
	"context"

	"github.com/tehrelt/eda/orders/internal/domain/dto"
	"github.com/tehrelt/eda/orders/internal/domain/entity"
)

func (s *OrderService) Save(ctx context.Context, in dto.SaveOrder) (*entity.Order, error) {
	order, err := s.repository.Save(ctx, in)
	if err != nil {
		return nil, err
	}

	if err := s.notifier.Notify(ctx, dto.SaveOrderEvent{
		Order: dto.Order{
			Id:    order.Id,
			Total: order.Total,
		},
		Customer: dto.Customer{
			Id:       order.Customer.Id,
			Username: order.Customer.Username,
		},
		CreatedAt: order.CreatedAt,
	}); err != nil {
		return nil, err
	}

	return order, nil
}
