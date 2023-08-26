package main

import "fmt"

// Command 命令对像
type Command interface {
	Execute()
}

// TurnOnLightCommand 开灯命令
type TurnOnLightCommand struct {
	Light *Light
}

// Execute 执行命令
func (c TurnOnLightCommand) Execute() {
	c.Light.On()
}

// TurnOffLightCommand 关灯命令
type TurnOffLightCommand struct {
	Light *Light
}

func (c TurnOffLightCommand) Execute() {
	c.Light.Off()
}

// TurnOnFanCommand 开风扇命令
type TurnOnFanCommand struct {
	Fan *Fan
}

func (c TurnOnFanCommand) Execute() {
	c.Fan.On()
}

// TurnOffFanCommand 关风扇命令
type TurnOffFanCommand struct {
	Fan *Fan
}

func (c TurnOffFanCommand) Execute() {
	c.Fan.Off()
}

// ConcreteCommand 宏命令，可以执行多个命令
type ConcreteCommand struct {
	Commands []Command
}

func (c ConcreteCommand) Execute() {
	for _, command := range c.Commands {
		command.Execute()
	}
}

// Light 灯
type Light struct {
	Name string
}

func (l Light) On() {
	fmt.Println(l.Name + " on")
}

func (l Light) Off() {
	fmt.Println(l.Name + " off")
}

// Fan 风扇
type Fan struct {
	Name string
}

func (f Fan) On() {
	fmt.Println(f.Name + " on")
}

func (f Fan) Off() {
	fmt.Println(f.Name + " off")
}

// SimpleRemoteControl 简单的遥控器，只有一个按钮
type SimpleRemoteControl struct {
	Commands Command
}

// ButtonWasPressed 按钮被按下
func (rc SimpleRemoteControl) ButtonWasPressed() {
	rc.Commands.Execute()
}

// NormalRemoteControl 普通的遥控器，有多个按钮
type NormalRemoteControl struct {
	Commands []Command
}

// ButtonWasPressed 按钮被按下
func (rc NormalRemoteControl) ButtonWasPressed(index int) {
	rc.Commands[index].Execute()
}

func main() {
	// 简单的遥控器
	fmt.Println("简单的遥控器")
	light := Light{Name: "Living Room"}
	remoteControl := SimpleRemoteControl{Commands: TurnOnLightCommand{Light: &light}}
	remoteControl.ButtonWasPressed()

	// 普通的遥控器
	fmt.Println("普通的遥控器")
	fan := Fan{Name: "Living Room"}
	normalRemoteControl := NormalRemoteControl{Commands: []Command{TurnOnLightCommand{Light: &light}, TurnOffLightCommand{Light: &light}, TurnOnFanCommand{Fan: &fan}, TurnOffFanCommand{Fan: &fan}}}
	normalRemoteControl.ButtonWasPressed(0)
	normalRemoteControl.ButtonWasPressed(1)
	normalRemoteControl.ButtonWasPressed(2)
	normalRemoteControl.ButtonWasPressed(3)

	// 通过宏，让一个按钮可以执行多个命令
	fmt.Println("通过宏，让一个按钮可以执行多个命令")
	normalRemoteControl.Commands = []Command{ConcreteCommand{Commands: []Command{TurnOnLightCommand{Light: &light}, TurnOnFanCommand{Fan: &fan}}}, ConcreteCommand{Commands: []Command{TurnOffLightCommand{Light: &light}, TurnOffFanCommand{Fan: &fan}}}}
	normalRemoteControl.ButtonWasPressed(0)
}
