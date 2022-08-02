package animal

import "fmt"

type NotAnimalProblem struct {
	message string
	code    int
}

func (n NotAnimalProblem) Error() string {
	return fmt.Sprintf("animal error! message: %s, code: %v", n.message, n.code)
}
