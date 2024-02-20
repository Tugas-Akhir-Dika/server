package entity

type ItemEntity struct {
	Id          int     `db:"id"`
	Title       string  `db:"title"`
	Price       float64 `db:"price"`
	Description string  `db:"description"`
	Category    string  `db:"category"`
	Image       string  `db:"image"`
	Rate        float64 `db:"rate"`
	Count       int     `db:"count"`
}
