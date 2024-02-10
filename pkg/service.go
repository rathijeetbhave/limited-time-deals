package pkg

import (
	"errors"
)

type GenericService struct {
	dao IGenericDao
}

func (s GenericService) CreateDeal(in interface{}) error {
	return s.dao.SaveToDB(in)
}

func (s GenericService) EndDeal(dealId uint) error {
	_, err := s.dao.GetDealById(dealId)
	if err != nil {
		return err
	}

	return s.dao.EndDeal(dealId)
}

func (s GenericService) UpdateDeal(req UpdateDealRequest) error {
	deal, err := s.dao.GetDealById(req.dealId)
	if err != nil {
		return err
	}

	if req.endTime > 0 {
		deal.EndTime = req.endTime
	}

	if req.limit > 0 {
		deal.Limit = req.limit
	}

	return s.dao.UpdateDeal(deal)
}

func (s GenericService) ClaimDeal(req UpdateDealRequest) error {
	deal, err := s.dao.GetDealById(req.dealId)
	if err != nil {
		return err
	}

	if !deal.IsActive() {
		return errors.New("Deal is not active")
	}

	deal.Limit -= 1

	err = s.dao.UpdateDeal(deal)
	if err != nil {
		return err
	}

	// create new object in availed_deals table
	return s.dao.SaveToDB(req)

}

func NewGenericService(dao IGenericDao) GenericService {
	return GenericService{dao: dao}
}
