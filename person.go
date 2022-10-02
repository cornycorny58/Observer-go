package awesomeProject1

import "fmt"

type Person struct {
	name string
}

func newPerson(name string) *Person {
	p := new(Person)
	p.name = name
	return p
}

func (p *Person) handleEvent(vacancies []string) {
	fmt.Printf("Hi dear %s", p.name)
	fmt.Println("Vacancies updated: ")
	for _, vacancies := range vacancies {
		fmt.Println(vacancies)
	}
}
