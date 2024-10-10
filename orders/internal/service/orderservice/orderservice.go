package orderservice

import (
	"context"

	"github.com/tehrelt/eda/orders/internal/domain/dto"
	"github.com/tehrelt/eda/orders/internal/domain/entity"
)

type OrderRepository interface {
	Save(ctx context.Context, order dto.SaveOrder) (*entity.Order, error)
	SummaryByCustomer(ctx context.Context, customerId string) (*dto.OrderSummary, error)
}

type EventNotifier interface {
	Notify(ctx context.Context, in dto.SaveOrderEvent) error
}

type OrderService struct {
	repository OrderRepository
	notifier   EventNotifier
}

func New(repository OrderRepository, notifier EventNotifier) *OrderService {
	return &OrderService{
		repository: repository,
		notifier:   notifier,
	}
}
