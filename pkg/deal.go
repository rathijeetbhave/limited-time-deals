package pkg

import (
	"errors"
	"time"
)

type Deal struct {
	id           uint
	items        []Item
	StartTime    uint64 // epoch seconds
	EndTime      uint64 // epoch seconds
	Limit        uint
	SpecialPrice uint
}

func (d Deal) IsActive() bool {
	currTime := time.Now().Unix()
	return uint64(currTime) >= d.StartTime && uint64(currTime) <= d.EndTime && d.Limit > 0
}

func (d Deal) Avail() error {
	if !d.IsActive() {
		return errors.New("Deal is not Open")
	}

	// delegate this to DB
	d.Limit -= 1

	return nil
}

func NewDeal(items []Item, startTime, endTime uint64, limit uint) Deal {
	return Deal{
		id:           0,
		items:        items,
		StartTime:    startTime,
		EndTime:      endTime,
		Limit:        limit,
		SpecialPrice: 0,
	}
}
