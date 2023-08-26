package main

type Pizza interface {
	GetName() string
}

type NYPizza struct {
	Name string
}

func (p NYPizza) GetName() string {
	return p.Name
}

type ChicagoPizza struct {
	Name string
}

func (p ChicagoPizza) GetName() string {
	return p.Name
}

type PizzaStore struct {
}

func (ps PizzaStore) OrderPizza(t string) Pizza {
	// 生产不同口味的披萨
	if t == "Cheese" {
		// 芝士味
		return NYPizza{Name: "Cheese Pizza"}
	} else if t == "Chocolate" {
		// 巧克力味
		return ChicagoPizza{Name: "Chocolate Pizza"}
	}

	return nil
}

func main() {
	pizzaStore := PizzaStore{}

	println(pizzaStore.OrderPizza("Cheese").GetName())    // Cheese Pizza
	println(pizzaStore.OrderPizza("Chocolate").GetName()) // Chocolate Pizza
}
