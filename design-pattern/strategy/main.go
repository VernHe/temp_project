package main

// Duck 鸭子接口
// ------------------ 通过继承的方式 ------------------
//type Duck interface {
//	fly()
//	quack()
//}
//
//// MallardDuck 绿头鸭
//type MallardDuck struct {
//	Name string
//}
//
//func (d MallardDuck) fly() {
//	println("用翅膀飞起来了")
//}
//
//func (d MallardDuck) quack() {
//	println("呱呱呱")
//}
//
//// RedheadDuck 红头鸭
//type RedheadDuck struct {
//	Name string
//}
//
//func (d RedheadDuck) fly() {
//	println("坐火箭飞起来了")
//}
//
//func (d RedheadDuck) quack() {
//	println("呱呱呱")
//}

// Duck 鸭子接口
// ------------------ 通过组合的方式 ------------------
type Duck struct {
	Name          string
	flyBehavior   FlyBehavior   // 飞行行为
	quackBehavior QuackBehavior // 叫声行为
}

// FlyBehavior 飞行行为接口
type FlyBehavior interface {
	fly()
}

type FlyWithWings struct {
}

func (f FlyWithWings) fly() {
	println("用翅膀飞起来了")
}

type FlyWithRocket struct {
}

func (f FlyWithRocket) fly() {
	println("坐火箭飞起来了")
}

// QuackBehavior 叫声行为接口
type QuackBehavior interface {
	quack()
}

type Quack struct {
}

func (q Quack) quack() {
	println("呱呱呱")
}

type Squeak struct {
}

func (q Squeak) quack() {
	println("吱吱吱")
}

func main() {
	// ------------------ 通过组合的方式 ------------------
	ducks := []Duck{
		Duck{
			Name:          "绿头鸭",
			flyBehavior:   FlyWithWings{},
			quackBehavior: Quack{},
		},
		Duck{
			Name:          "红头鸭",
			flyBehavior:   FlyWithRocket{},
			quackBehavior: Squeak{},
		},
	}

	for _, duck := range ducks {
		duck.flyBehavior.fly()
		duck.quackBehavior.quack()
	}
	// 输出:
	// 用翅膀飞起来了
	// 呱呱呱
	// 坐火箭飞起来了
	// 吱吱吱
}
