package main

import (
	"strconv"
)

type Product struct {
	Name, Category string
	Price          float64
}

var ProductList = []*Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flag", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
	{"Bling-Bling King", "Chess", 1200},
	{"Aqua Shoes", "Watersports", 29.50},
	{"Tennis Ball", "Tennis", 10.00},
	{"Tennis Racket", "Tennis", 200.00},
	{"Tennis Shoes", "Tennis", 100.00},
	{"Tennis Net", "Tennis", 50.00},
	{"Tennis Court", "Tennis", 5000},
	{"Tennis Bag", "Tennis", 25.00},
	{"Tennis Ball Machine", "Tennis", 1000.00},
}

type ProductGroup []*Product

type ProductData = map[string]ProductGroup

var Products = make(ProductData)

func ToCurrency(value float64) string {
	return "$" + strconv.FormatFloat(value, 'f', 2, 64)
}

func init() {
	for _, p := range ProductList {
		if _, ok := Products[p.Category]; ok {
			Products[p.Category] = append(Products[p.Category], p)
		} else {
			Products[p.Category] = ProductGroup{p}
		}
	}
}
