package domain

import "fmt"

type Product interface {
	GetID() int
	GetModel() string
	GetInStock() int
	GetPrice() int
	ReduceStock(quantity int) error
	Clone() Product
}

type BaseProduct struct {
	ID      int    `json:"id"`
	Model   string `json:"model"`
	InStock int    `json:"in_stock"`
	Price   int    `json:"price"`
}

func (bp *BaseProduct) ReduceStock(quantity int) error {
	if bp.InStock < quantity {
		return fmt.Errorf("not enough stock")
	}
	bp.InStock -= quantity
	return nil
}
func (bp *BaseProduct) Clone() Product {
	return &BaseProduct{
		ID:      bp.ID,
		Model:   bp.Model,
		Price:   bp.Price,
		InStock: bp.InStock,
	}

}

func (bp BaseProduct) GetID() int {
	return bp.ID
}
func (bp BaseProduct) GetModel() string {
	return bp.Model
	//Sonyx, Kurumi - chair
	//Unibi, Facito - wardrobe
}
func (bp BaseProduct) GetInStock() int {
	return bp.InStock
}
func (bp BaseProduct) GetPrice() int {
	return bp.Price
}

type Chair struct {
	BaseProduct
	Material string `json:"material"`
	Type     string `json:"type"`
}

type Wardrobe struct { //шкаф
	BaseProduct
	Material string `json:"material"`
	Type     string `json:"type"`
}

type Conditioner struct {
	BaseProduct
	Version string `json:"version"`
	Type    string `json:"type"`
}
