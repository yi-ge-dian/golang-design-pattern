package behavior_type_mode

import "fmt"

// Strategy 定义一个策略接口
type Strategy interface {
	Execute(a, b int) int
}

// AdditionStrategy 具体策略1:加法
type AdditionStrategy struct{}

func (s *AdditionStrategy) Execute(a, b int) int {
	return a + b
}

// SubtractionStrategy 具体策略2:减法
type SubtractionStrategy struct{}

func (s *SubtractionStrategy) Execute(a, b int) int {
	return a - b
}

// Context 上下文,用于调用策略
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}
func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func CallPolicyPattern() {
	context := &Context{}
	//使用加法策略
	addition := &AdditionStrategy{}
	context.SetStrategy(addition)
	fmt.Println("10 + 5 =", context.ExecuteStrategy(10, 5))
	//使用减法策略
	subtraction := &SubtractionStrategy{}
	context.SetStrategy(subtraction)
	fmt.Println("10 - 5 =", context.ExecuteStrategy(10, 5))
}
