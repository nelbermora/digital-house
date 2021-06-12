package samples

import "fmt"

type PersonActions interface {
	Eat()
	Sleep()
	TurnYears()
}
type Person struct {
	Name     string
	LastName string
	Age      int
}

func NewPerson(name, lastName string, age int) PersonActions {
	return &Person{
		Name:     name,
		LastName: lastName,
		Age:      age,
	}
}
func (p *Person) Eat() {
	fmt.Println(p.Name+" está comiendo y su edad es", p.Age)
}
func (p *Person) Sleep() {
	fmt.Println(p.Name+" está durmiendo y su edad es", p.Age)
}
func (p *Person) TurnYears() {
	p.Age = p.Age + 1 // Una forma mas simple de expresarlo es p.Age += 1
	fmt.Println(p.Name+" acaba de cumplir años y ahora su edad es", p.Age)
}
