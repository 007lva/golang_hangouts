package main

import "fmt"
import "./models"

type Student struct {
	name      string
	last_name string
	age       int8
}

func (s Student) full_name() string {
	return s.name + " " + s.last_name
}

func (s *Student) setName(new_name string) {
	s.name = new_name
}

func (s *Student) setAge(new_age int8) {
	s.age = new_age
}

func main() {
	student := Student{name: "hugo", last_name: "lopez", age: 15}
	//cambiarNombre(&student)
	student.setName("renzo")
	student.name = "ww"
	full_name := student.full_name()
	fmt.Println(full_name)

	//var text string
	var animal models.Animal
	//animal := models.Animal{}

	animal = models.Dog{}
	//animal.SetName("Fido")
	fmt.Println(animal)

	animal = models.Cat{}

	animal = models.Cuy{}

	animal.Speak()

	animal = models.Horse{}

}

func cambiarNombre(s *Student) {
	s.name = "pancho"
}
