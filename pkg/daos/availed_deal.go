package daos

// Dao layer for availed_deals table
type AvailedDealDao struct {
	db any
}

func (a AvailedDealDao) HasUserBoughtDeal(userId, dealId uint) bool {
	// Query availed deal from availed_deals for given userId and dealId
	// If any deal has been availed, return true else return false
	// query = "select * from availed_deals where user_id = ? and deal_id = ? order by id desc limit 1"
	return false
}

func NewAvailedDealsDao() AvailedDealDao {
	return AvailedDealDao{}
}
