package calculation

import "errors"

func Calculate(a, b int, action string) (int, error) {
	if (a > 10 && a < 0) || (b > 10 && b < 0) {
		return 0, errors.New("Некорректный ввод - переменная должна быть от 1 до 10 включительно")
	}
	switch action {
	case "+":
		return addition(a, b), nil
	case "-":
		return subtraction(a, b), nil
	case "*":
		return multiplication(a, b), nil
	case "/":
		if b != 0 {
			return division(a, b), nil
		} else {
			return 0, errors.New("На 0 делить нельзя")
		}
	default:
		return 0, errors.New("Невозможное действие")
	}
}
func addition(a, b int) int {
	return a + b
}
func subtraction(a, b int) int {
	return a - b
}
func multiplication(a, b int) int {
	return a * b
}
func division(a, b int) int {
	return a / b
}
