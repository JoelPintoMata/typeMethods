package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
	Character *Person
	SonOf     *Person
	Power     int
}

func main() {
	// and to use it:
	goku := &Saiyan{
		Character: &Person{"Gohan"},
		SonOf:     &Person{"Goku"},
		Power:     9001,
	}

	goku.Character.Introduce()
	goku.SonOf.Introduce()
}
