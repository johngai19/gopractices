package store

var standardTax = newTaxRate(0.25, 20)

// Product represents a product in the store.
type Product struct {
	Name, Category string //Go examines the first letter of the names of the fields
	price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{
		Name:     name,
		Category: category,
		price:    price,
	}
}

func (p *Product) Price() float64 {
	return standardTax.calcTax(p)
}
func (p *Product) SetPrice(price float64) {
	p.price = price
}
