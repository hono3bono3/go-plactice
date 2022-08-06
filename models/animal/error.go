package animal

import "fmt"

type NotAnimalProblem struct {
	message string
	code    int
}

func (n NotAnimalProblem) Error() string {
	s := square{}
	func(r rectangle) {

	}(s)
	return fmt.Sprintf("animal error! message: %s, code: %v", n.message, n.code)
}

type rectangle interface {
	height() int
	width() int
}

type square struct {
	length int
}

func (sq *square) width() int {
	return sq.length
}

func (sq *square) height() int {
	return sq.length
}
