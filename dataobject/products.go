package dataobject

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID			int `json:"id"`
	Name		string `json:"name"`
	Description	string `json:"description"`
	Price		float32 `json:"price"`
	SKU			string	`json:"sku"`
	CreatedOn	string `json:"-"`
	UpdatedOn	string `json:"-"`
	DeletedOn	string `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode()
}

func GetProducts() Products {
	return productList
}

var productList = []*Product {
	&Product{
		ID:			1,
		Name:		"Latte",
		Description:	"Frothy milky coffee"
		Price:		3.90,
		SKU:			"caf2345",
		CreatedOn:	time.Now().UTC().string(),
		UpdatedOn:	time.Now().UTC().string(),
	},
	&Product{
		ID:			2,
		Name:		"Espresso",
		Description:"strong coffee without milk"
		Price:		3.25,
		SKU:		"caf1499",
		CreatedOn:	time.Now().UTC().string(),
		UpdatedOn:	time.Now().UTC().string(),
	},
	&Product{
		ID:			3,
		Name:		"Moka",
		Description:"Coffee with milk"
		Price:		4.90,
		SKU:		"caf4567",
		CreatedOn:	time.Now().UTC().string(),
		UpdatedOn:	time.Now().UTC().string(),
	},
}