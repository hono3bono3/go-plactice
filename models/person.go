package models

type Person struct {
	FirstName string
	LastName  string
}

func NewPerson(fitst, last string) *Person {
	return &Person{
		FirstName: fitst,
		LastName:  last,
	}
}

func (p *Person) ChangeFirstName(first string) {
	p.FirstName = first
}

func (p *Person) ChangeLastName(last string) {
	p.LastName = last
}
