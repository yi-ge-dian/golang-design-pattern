package main

import "fmt"

// Product 接口定义
type Product interface {
	GetName() string
}

// ProductA 类
type ProductA struct{}

func (p *ProductA) GetName() string {
	return "Product A"
}

// ProductB 类
type ProductB struct{}

func (p *ProductB) GetName() string {
	return "Product B"
}

// Factory 工厂类
type Factory struct{}

func (f *Factory) CreateProduct(name string) Product {
	switch name {
	case "A":
		return &ProductA{}
	case "B":
		return &ProductB{}
	default:
		return nil
	}
}

func main02() {
	factory := &Factory{}
	productA := factory.CreateProduct("A")
	productB := factory.CreateProduct("B")
	fmt.Println(productA.GetName()) // Output: Product A
	fmt.Println(productB.GetName()) // Output: Product B
}
