package domain

type Plant struct {
	Id          int64
	Name        string
	Category    string
	Description string
	PricePerDay float32
	IsAvailable bool
}
