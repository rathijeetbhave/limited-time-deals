package daos

import (
	"awesomeProject/pkg"
	"errors"
)

type DealDao struct {
	dbConn any
}

func (dao DealDao) GetDealById(dealId uint) (pkg.Deal, error) {
	// If deal does not exist in DB return error
	if dealId == 0 {
		return pkg.Deal{}, errors.New("Deal does not exist")
	}

	// Fire GET query to DB to get deal by ID
	// query = "Select * from deal where id = ?"
	return pkg.Deal{}, nil
}

func (dao DealDao) EndDeal(dealId uint) error {
	// Fire DB query to update endTime to start of epoch time
	// This way we will know that the deal is ended manually
	// query = "update deal set end_time = '1970-00-00 00:00:00' where id = ?"
	return nil
}

func (dao DealDao) UpdateDeal(deal pkg.Deal) error {
	// Fire DB query to update a deal
	// query = "Update deal set limit = ?, endTime = ? where id = ?"
	return nil
}
