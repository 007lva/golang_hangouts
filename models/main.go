package models

import "fmt"

type Animal interface {
	Speak()
}

//inicio perro
type Dog struct {
	name string
}

func (d Dog) Speak() {
	fmt.Println("guauguau")
}

func (d *Dog) SetName(new_name string) {
	d.name = new_name
}

//fin perro

//inicio gato
type Cat struct {
	name string
}

func (c Cat) Speak() {
	fmt.Println("miaumiau")
}

//fin gato

type Cuy struct {
}

func (c Cuy) Speak() {
	fmt.Println("????")
}

type Horse struct {
}

func (h Horse) Speak() {
	fmt.Println("????")
}
