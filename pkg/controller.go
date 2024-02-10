package pkg

import (
	"context"
	"errors"
	"time"
)

type DealController struct {
	service IGenericService
}

func (c DealController) CreateDeal(ctx context.Context, deal Deal) error {
	if deal.Limit <= 0 {
		return errors.New("Deal should have limit > 0")
	}

	currTime := time.Now().Unix()

	if deal.StartTime < uint64(currTime) || deal.EndTime < uint64(currTime) {
		return errors.New("Deal start and end time should be in future")
	}

	if len(deal.items) == 0 {
		return errors.New("Deal should be applicable on an item")
	}

	if deal.StartTime > deal.EndTime {
		return errors.New("Invalid data")
	}

	return c.service.CreateDeal(deal)
}

func (c DealController) EndDeal(ctx context.Context, dealId uint) error {
	if dealId <= 0 {
		return errors.New("DealId cannot be negative")
	}

	return c.service.EndDeal(dealId)
}

func (c DealController) UpdateDeal(ctx context.Context, req UpdateDealRequest) error {
	if req.dealId <= 0 {
		return errors.New("DealId cannot be negative")
	}

	if req.limit <= 0 && req.endTime <= 0 {
		return errors.New("Invalid data")
	}

	return c.service.UpdateDeal(req.dealId, req.limit, req.endTime)
}

func (c DealController) ClaimDeal(ctx context.Context, req ClaimDealRequest) error {
	if req.dealId == 0 || req.userId == 0 {
		return errors.New("Invalid data")
	}

	return c.service.ClaimDeal(ctx, req)
}
