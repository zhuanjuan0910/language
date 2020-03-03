package main

import "fmt"

type Animal struct {
	Name string
}
type Dog struct {
	Animal //类似继承，可以继承Animal的属性和方法
}
type Sayer interface {
	Say(string)
	SayHi()
}

func (a *Animal) Say(message string) {
	fmt.Printf("Animal[%v] say: %v\n", a.Name, message)
}

func (a *Animal) SayHi() {
	a.Say("Hi")
}

func (d *Dog) Say(message string) {
	fmt.Printf("Dog[%v] say: %v\n", d.Name, message)
}

func main() {
	var sayer Sayer

	sayer = &Dog{Animal{Name: "Yoda"}}
	sayer.Say("hello world")
	sayer.SayHi()
}
