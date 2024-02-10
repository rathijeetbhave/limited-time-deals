package pkg

import (
	"testing"
	"time"
)

type isActiveTest struct {
	name             string
	dealStart        uint64
	dealEnd          uint64
	limit            uint
	expectedResponse bool
}

func TestDeal_IsActive(t *testing.T) {
	isActiveTests := []isActiveTest{
		{
			name:      "CurrentTimeLessThanStartTime__DealYetToStart",
			dealStart: uint64(time.Now().AddDate(100, 0, 0).Unix()),
			dealEnd:   uint64(time.Now().AddDate(101, 0, 0).Unix()),
		},
		{
			name:      "CurrentTimeGreaterThanEndTime__DealAlreadyOver",
			dealStart: uint64(time.Now().AddDate(-2, 0, 0).Unix()),
			dealEnd:   uint64(time.Now().AddDate(-1, 0, 0).Unix()),
		},
		{
			name:      "DealLimitExceeded",
			dealStart: uint64(time.Now().AddDate(-2, 0, 0).Unix()),
			dealEnd:   uint64(time.Now().AddDate(1, 0, 0).Unix()),
		},
		{
			name:             "SuccessFlow",
			dealStart:        uint64(time.Now().AddDate(-2, 0, 0).Unix()),
			dealEnd:          uint64(time.Now().AddDate(1, 0, 0).Unix()),
			limit:            10,
			expectedResponse: true,
		},
	}

	for _, test := range isActiveTests {
		t.Run(test.name, func(t *testing.T) {
			d := NewDeal(nil, test.dealStart, test.dealEnd, test.limit)
			if d.IsActive() != test.expectedResponse {
				t.Errorf("Expected and actual response is not same for test : %s. Expected : %t, Actual : %t", test.name, test.expectedResponse, d.IsActive())
				return
			}
		})
	}
}
