package main

import "fmt"

// Option 是配置选项
type Option func(*config)

// config是需要配置的对象
type config struct {
	host string
	port int
}

// WithHost 设置主机地址
func WithHost(host string) Option {
	return func(c *config) {
		c.host = host
	}
}

// WithPort 设置端口号
func WithPort(port int) Option {
	return func(c *config) {
		c.port = port
	}
}

// NewConfig 创建一个新的配置对象
func NewConfig(opts ...Option) *config {
	c := &config{
		host: "localhost",
		port: 8080,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func CallConfigurePatternHighLevel() {
	c1 := NewConfig()
	fmt.Printf("Default config: %v\n", c1)
	c2 := NewConfig(WithHost("example.com"), WithPort(9000))
	fmt.Printf("Custom config: %v\n", c2)
}
