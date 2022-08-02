package animal

type Animal interface {
	Say() string
}

type Dog struct {
	Animal
}

type Human struct {
	Name string
}

func (a Dog) Say() string {
	return "wan! wan!"
}

func (a Human) Say() (n string) {
	n = "My name is " + a.Name
	return
}
