package main

import "fmt"

// 抽象工厂

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// 具体工厂

type WinFactory struct{}

func (w *WinFactory) CreateButton() Button {
	return &WinButton{}
}
func (w *WinFactory) CreateCheckbox() Checkbox {
	return &WinCheckbox{}
}

type MacFactory struct{}

func (m *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (m *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

// 抽象产品

type Button interface {
	Click()
}
type Checkbox interface {
	Check()
}

// 具体产品

type WinButton struct{}

func (w *WinButton) Click() {
	fmt.Println("Windows button clicked")
}

type WinCheckbox struct{}

func (w *WinCheckbox) Check() {
	fmt.Println("Windows checkbox checked")
}

type MacButton struct{}

func (m *MacButton) Click() {
	fmt.Println("Mac button clicked")
}

type MacCheckbox struct{}

func (m *MacCheckbox) Check() {
	fmt.Println("Mac checkbox checked")
}

func CallAbstractFactoryPattern() {
	var factory GUIFactory
	// 根据需要选择具体工厂
	factory = &WinFactory{}
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()
	button.Click()
	checkbox.Check()
}
