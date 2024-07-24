package patterns

import (
	"errors"
	"fmt"
)

type IProduct interface {
	SetName(name string)
	SetPrice(price float32)
	GetInfo() string
}

func ProductFactory(category string) (IProduct, error) {
	switch category {
	case "PCComponents":
		return NewPCComponents(), nil
	case "ReadyMadePC":
		return NewReadyMadePC(), nil
	case "Appliances":
		return NewAppliances(), nil
	default:
		return nil, errors.New("category not recognized")
	}
}

type PCComponents struct {
	Product
}

func NewPCComponents() IProduct {
	return &PCComponents{
		Product: Product{category: "PCComponents"},
	}
}

func (p *PCComponents) SetName(name string) {
	p.Product.SetName(name)
}

func (p *PCComponents) SetPrice(price float32) {
	p.Product.SetPrice(price)
}

func (p *PCComponents) GetInfo() string {
	return p.Product.GetInfo()
}

type ReadyMadePC struct {
	Product
}

func NewReadyMadePC() IProduct {
	return &ReadyMadePC{
		Product: Product{category: "ReadyMadePC"},
	}
}

func (p *ReadyMadePC) SetName(name string) {
	p.Product.SetName(name)
}

func (p *ReadyMadePC) SetPrice(price float32) {
	p.Product.SetPrice(price)
}

func (p *ReadyMadePC) GetInfo() string {
	return p.Product.GetInfo()
}

type Appliances struct {
	Product
}

func NewAppliances() IProduct {
	return &Appliances{
		Product: Product{category: "Appliances"},
	}
}

func (p *Appliances) SetName(name string) {
	p.Product.SetName(name)
}

func (p *Appliances) SetPrice(price float32) {
	p.Product.SetPrice(price)
}

func (p *Appliances) GetInfo() string {
	return p.Product.GetInfo()
}

type Product struct {
	category string
	name     string
	price    float32
}

func (p *Product) SetName(name string) {
	p.SetName(name)
}

func (p *Product) SetPrice(price float32) {
	p.SetPrice(price)
}

func (p *Product) GetInfo() string {
	return fmt.Sprintf("Product info: name - %s | price - %f", p.name, p.price)
}
