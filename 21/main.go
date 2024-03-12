package main

import "fmt"

type Food interface {
	Bread()
}

type bread struct{}

func (a *bread) Bread() {
	fmt.Println("EAT BREAD!")
}

type Grandbaba struct{}

func (d *Grandbaba) Eat(t Food) {
	t.Bread()
}

type Other interface {
	Milk()
}

type Drink struct{}

func (h *Drink) Milk() {
	fmt.Println("Drink MILLK")
}

type MilkToBreadAdapter struct {
	drink *Drink
}

func (adapter *MilkToBreadAdapter) Bread() {
	adapter.drink.Milk()
}

func main() {
	// создаем бабулю
	grandbaba := &Grandbaba{}
	bread := &bread{}
	grandbaba.Eat(bread)

	drink := &Drink{}
	TransFer := &MilkToBreadAdapter{
		drink: drink,
	}

	grandbaba.Eat(TransFer)
}
