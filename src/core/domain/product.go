package domain

import "time"

// Domínio da aplicação
type Product struct {
	id              int
	name            string
	price           float32
	code            string
	expiration_date time.Time
}

func (p Product) Id() int {
	return p.id
}

func (p Product) Nome() string {
	return p.name
}

func (p Product) Code() string {
	return p.code
}

func (p Product) Price() float32 {
	return p.price
}

func (p Product) ExpirationDate() time.Time {
	return p.expiration_date
}

func newProduct(id int, name string, price float32, code string, expiration_date time.Time) *Product {
	return &Product{
		id:              id,
		name:            name,
		price:           price,
		code:            code,
		expiration_date: expiration_date,
	}
}
