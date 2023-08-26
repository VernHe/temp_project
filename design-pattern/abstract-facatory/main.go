package main

import "fmt"

// Pizza 生产的披萨
type Pizza struct {
	factory PizzaIngredientFactory
}

func (p Pizza) Prepare() {
	fmt.Printf("正在使用 %s 酱料和 %s 芝士准备披萨\n", p.factory.CreateDough().GetName(), p.factory.CreateCheese().GetName())
	// 此处省略了具体的制作过程
}

// PizzaIngredientFactory 披萨原料工厂的抽象
type PizzaIngredientFactory interface {
	CreateDough() Dough
	CreateCheese() Cheese
	// 其他原料...
}

// Dough 面团接口
type Dough interface {
	GetName() string
}

// NYDough 纽约面团
type NYDough struct {
}

func (d NYDough) GetName() string {
	return "NY Dough"
}

// ChicagoDough 芝加哥面团
type ChicagoDough struct {
}

func (d ChicagoDough) GetName() string {
	return "Chicago Dough"
}

// Cheese 芝士接口
type Cheese interface {
	GetName() string
}

// NYCheese 纽约芝士
type NYCheese struct {
}

func (c NYCheese) GetName() string {
	return "NY Cheese"
}

// ChicagoCheese 芝加哥芝士
type ChicagoCheese struct {
}

func (c ChicagoCheese) GetName() string {
	return "Chicago Cheese"
}

// NYPizzaIngredientFactory 纽约披萨原料工厂
type NYPizzaIngredientFactory struct {
}

func (f NYPizzaIngredientFactory) CreateDough() Dough {
	return NYDough{}
}

func (f NYPizzaIngredientFactory) CreateCheese() Cheese {
	return NYCheese{}
}

// ChicagoPizzaIngredientFactory 芝加哥披萨原料工厂
type ChicagoPizzaIngredientFactory struct {
}

func (f ChicagoPizzaIngredientFactory) CreateDough() Dough {
	return ChicagoDough{}
}

func (f ChicagoPizzaIngredientFactory) CreateCheese() Cheese {
	return ChicagoCheese{}
}

func main() {
	NYPizza := Pizza{factory: NYPizzaIngredientFactory{}}
	NYPizza.Prepare() // 正在使用 NY Dough 酱料和 NY Cheese 芝士准备披萨

	ChicagoPizza := Pizza{factory: ChicagoPizzaIngredientFactory{}}
	ChicagoPizza.Prepare() // 正在使用 Chicago Dough 酱料和 Chicago Cheese 芝士准备披萨
}
