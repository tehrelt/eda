package orderservice

import (
	"context"

	"github.com/tehrelt/eda/orders/internal/domain/dto"
)

func (s *OrderService) SummaryByCustomer(ctx context.Context, customerId string) (*dto.OrderSummary, error) {
	res, err := s.repository.SummaryByCustomer(ctx, customerId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
