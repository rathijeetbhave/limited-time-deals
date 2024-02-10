package pkg

import "context"

type IGenericDao interface {
	SaveToDB(in interface{}) error
	GetDealById(dealId uint) (Deal, error)
	EndDeal(dealId uint) error
	UpdateDeal(deal Deal) error
}

type IGenericService interface {
	CreateDeal(in interface{}) error
	EndDeal(dealId uint) error
	UpdateDeal(dealId uint, limit uint, endTime uint64) error
	ClaimDeal(ctx context.Context, req ClaimDealRequest) error
}
