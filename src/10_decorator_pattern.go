package main

import "fmt"

// Component 是一个接口,定义了一个操作
type Component interface {
	Operation() string
}

// ConcreteComponent 是一个具体的组件,实现了 Component接口
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

// Decorator 是一个装饰器函数,接收一个 Component 接口,并返回一个新的带有附加行为的
type Decorator func(Component) Component

// ConcreteDecoratorA 是一个具体的装饰器,给 Component 添加新的行为
func ConcreteDecoratorA(component Component) Component {
	return &decoratorA{component: component}
}

type decoratorA struct {
	component Component
}

func (d *decoratorA) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorA(%s)", d.component.Operation())
}

// ConcreteDecoratorB 是另一个具体的装饰器,给Component添加新的行为
func ConcreteDecoratorB(component Component) Component {
	return &decoratorB{component: component}
}

type decoratorB struct {
	component Component
}

func (d *decoratorB) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorB(%s)", d.component.Operation())
}

func CallDecoratorPattern() {
	component := &ConcreteComponent{}
	decoratedComponentA := ConcreteDecoratorA(component)
	decoratedComponentB := ConcreteDecoratorB(decoratedComponentA)
	fmt.Println(component.Operation())           // 输出: ConcreteComponent
	fmt.Println(decoratedComponentA.Operation()) // 输出: ConcreteDecoratorA(ConcreteComponent)
	fmt.Println(decoratedComponentB.Operation()) // 输出: ConcreteDecoratorB(ConcreteDecoratorA(ConcreteComponent))
}
