package main

import (
	"errors"
	"fmt"
)

// есть интерфейс, который ожидает определённую имплементацию скалярного произведения
type dotProductService interface {
	dotProduct([]int, []int) (float32, error)
}

// у нас уже есть тип calculator, реализующий нужный функционал, но его метод возвращает значение не того типа,
// что требует интерфейс dotProductService
type calculator struct{}

func (c *calculator) dotProduct(lhs []int, rhs []int) (int, error) {
	if len(lhs) != len(rhs) {
		return 0, errors.New("vectors must have the same dimensionality")
	}

	var res int
	for i := range lhs {
		res += lhs[i] * rhs[i]
	}
	return res, nil
}

// поэтому мы реализуем адаптер, который оборачивает нашу структуру и делает её совместимой с интерфейсом
type calculatorAdapter struct {
	calc calculator
}

func (ca *calculatorAdapter) dotProduct(lhs []int, rhs []int) (float32, error) {
	res, err := ca.calc.dotProduct(lhs, rhs)
	return float32(res), err
}

func newCalculatorAdapter(calc calculator) *calculatorAdapter {
	return &calculatorAdapter{calc: calc}
}

func main() {
	calc := new(calculator)
	service := dotProductService(newCalculatorAdapter(*calc))

	vec := []int{1, 2, 3, 4, 5}

	if dotProduct, err := service.dotProduct(vec, vec); err != nil {
		panic(err)
	} else {
		fmt.Println(dotProduct)
	}
}
