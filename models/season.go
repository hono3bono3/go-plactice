package models

type Season int

const (
	Peak   Season = iota + 1 // 繁忙期
	Normal                   // 通常期
	Off                      // 閑散期
)

func (s Season) Price(price float64) float64 {
	if s == Peak {
		return price + 200
	}
	return price
}
