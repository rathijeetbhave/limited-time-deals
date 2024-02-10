package pkg

type User struct {
	id    uint
	name  string
	email string
}

type Item struct {
	id    uint
	name  string
	price uint
}

type Seller struct {
	id          uint
	name        string
	shopAddress string
}

type AvailedDeals struct {
	id   uint
	Deal Deal
	User User
}
