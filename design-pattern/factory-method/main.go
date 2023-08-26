package main

// Pizza 产品的抽象
type Pizza interface {
	GetName() string
}

type NYPizza struct {
	name string
}

func (p NYPizza) GetName() string {
	return p.name
}

type ChicagoPizza struct {
	name string
}

func (p ChicagoPizza) GetName() string {
	return p.name
}

// PizzaStore 创建者的抽象
type PizzaStore interface {
	OrderPizza(t string) Pizza
}

type NYPizzaStore struct {
}

func (s NYPizzaStore) OrderPizza(t string) Pizza {
	if t == "1" {
		return NYPizza{name: "NY 1"}
	}
	return NYPizza{name: "NY 2"}
}

type ChicagoPizzaStore struct {
}

func (s ChicagoPizzaStore) OrderPizza(t string) Pizza {
	if t == "1" {
		return ChicagoPizza{name: "Chicago 1"}
	}
	return ChicagoPizza{name: "Chicago 2"}
}

func main() {
	NYPizzaStore := NYPizzaStore{}
	println(NYPizzaStore.OrderPizza("1").GetName()) // NY 1

	ChicagoPizzaStore := ChicagoPizzaStore{}
	println(ChicagoPizzaStore.OrderPizza("1").GetName()) // Chicago 1
}
