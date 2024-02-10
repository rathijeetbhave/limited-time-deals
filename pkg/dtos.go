package pkg

type UpdateDealRequest struct {
	dealId  uint
	limit   uint
	endTime uint64
}

type ClaimDealRequest struct {
	userId uint
	dealId uint
}
