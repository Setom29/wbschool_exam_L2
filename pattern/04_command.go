package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// https://golangbyexample.com/command-design-pattern-in-golang/

import "fmt"

// buttom

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// command
type command interface {
	execute()
}

// offCommand
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// onCommand
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// device
type device interface {
	on()
	off()
}

// tv
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

// main
// func main() {
// 	tv := &tv{}
// 	onCommand := &onCommand{
// 		device: tv,
// 	}
// 	offCommand := &offCommand{
// 		device: tv,
// 	}
// 	onButton := &button{
// 		command: onCommand,
// 	}
// 	onButton.press()
// 	offButton := &button{
// 		command: offCommand,
// 	}
// 	offButton.press()
// }
