package main

import "fmt"

type Animal interface {
	Speak()
}

type animal struct {
	Name  string
	Breed string
	Color string
	Sound string
}

func (a *animal) Speak() {
	fmt.Printf("%s say %s", a.Name, a.Sound)
	fmt.Println()
}

type Cat struct {
	animal
}

func NewCat() Animal {
	return &Cat{
		animal: animal{
			Name:  "Whiskers",
			Breed: "Munchkin",
			Color: "White",
			Sound: "Meowww",
		},
	}
}

type Dog struct {
	animal
}

func NewDog() Animal {
	return &Dog{
		animal: animal{
			Name:  "Brian",
			Breed: "Labrador",
			Color: "Gold",
			Sound: "Bark",
		},
	}
}

func CreateAnimal(animal string) Animal {
	switch animal {
	case "cat":
		return NewCat()
	case "dog":
		return NewDog()
	default:
		return nil
	}
}


func main() {
	cat := CreateAnimal("cat")
	dog := CreateAnimal("dog")


	cat.Speak()
	dog.Speak()
}