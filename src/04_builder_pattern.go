package main

import "fmt"

// Product_ 类
type Product_ struct {
	PartA string
	PartB string
	PartC string
}

// Builder 接口
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetProduct() *Product_
}

// 具体建造者类

type ConcreteBuilder struct {
	product *Product_
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

func (c *ConcreteBuilder) GetProduct() *Product_ {
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

func CallBuilderPattern() {
	concreteBuilder := &ConcreteBuilder{product: &Product_{}}
	director := &Director{builder: concreteBuilder}
	director.Construct()
	product := concreteBuilder.GetProduct()
	fmt.Println(product.PartA, product.PartB, product.PartC)
}
