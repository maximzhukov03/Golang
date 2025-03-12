package main

import (
	"fmt"
)

// Пример кода, если необходимо
func generateMathString(operands []int, operator string) string {
    res := operands[0]
    stringMath := fmt.Sprintf("%d", res) // Начинаем с первого операнда

    for _, operand := range operands[1:] {
        switch operator {
        case "+":
            res += operand
        case "-":
            res -= operand
        case "*":
            res *= operand
        case "/":
            if operand != 0 {
                res /= operand
            } else {
                return "Ошибка"
            }
        }
        stringMath += fmt.Sprintf(" %s %d", operator, operand) // Добавляем оператор и операнд
    }

    return fmt.Sprintf("%s = %d", stringMath, res) // Возвращаем строку с результатом
}

// Пример результата выполнения программы:
func main() {
    fmt.Println(generateMathString([]int{2, 4, 6}, "+")) // "2 + 4 + 6 = 12"
}