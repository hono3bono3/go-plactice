package main

import (
	"fmt"

	"github.com/hono3bono3/go-plactice/models/animal"
)

func main() {
	dog := animal.Dog{}
	human := animal.Human{
		Name: "John",
	}
	Run(dog)
	Run(human)

}

func Run(a animal.Animal) {
	fmt.Println(a.Say())
}
