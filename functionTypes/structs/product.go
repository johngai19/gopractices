package main

type Product struct {
	name, category string
	price          float64
}

func (p Product) getName() string {
	return p.name
}
func (p Product) getCost(annual bool) float64 {
	if annual {
		return p.price * 12
	}
	return p.price
}
