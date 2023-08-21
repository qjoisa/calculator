package workwithIO

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetExpression() (string, string, string, error) {
	reader := bufio.NewReader(os.Stdin)
	var expression string
	expression, _ = reader.ReadString('\n')
	expression = strings.TrimSpace(expression)
	return BreakingIntoParts(expression)
}

func BreakingIntoParts(expression string) (string, string, string, error) {
	var elements []string = strings.Split(expression, " ")
	if len(elements) > 3 {
		return "", "", "", errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else if len(elements) < 3 {
		return "", "", "", errors.New("Вывод ошибки, так как строка не является математической операцией.")
	}
	return strings.ToUpper(elements[0]), strings.ToUpper(elements[2]), strings.ToUpper(elements[1]), nil
}

func PrintResult(resultArabic int, resultRoman string) {
	if resultRoman != "" {
		fmt.Println(resultRoman)
	} else if resultRoman == "" {
		fmt.Println(resultArabic)
	}
}

func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
