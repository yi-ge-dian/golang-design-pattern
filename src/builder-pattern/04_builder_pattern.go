package main

import "fmt"

// Product 类
type Product struct {
	PartA string
	PartB string
	PartC string
}

// Builder 接口
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetProduct() *Product
}

// 具体建造者类

type ConcreteBuilder struct {
	product *Product
}

func (c *ConcreteBuilder) BuildPartA() {
	c.product.PartA = "Part A"
}
func (c *ConcreteBuilder) BuildPartB() {
	c.product.PartB = "Part B"
}
func (c *ConcreteBuilder) BuildPartC() {
	c.product.PartC = "Part C"
}

func (c *ConcreteBuilder) GetProduct() *Product {
	return c.product
}

// Director类

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func main() {
	concreteBuilder := &ConcreteBuilder{product: &Product{}}
	director := &Director{builder: concreteBuilder}
	director.Construct()
	product := concreteBuilder.GetProduct()
	fmt.Println(product.PartA, product.PartB, product.PartC)
}
