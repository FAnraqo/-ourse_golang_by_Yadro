package main

import (
	"fmt"
)

func main() {
	var flag_g = false
	for !flag_g {
		fmt.Print("Введите первое число: ")
		var first int
		flag := false
		for !flag {
			n, err := fmt.Scan(&first)
			if n != 1 || err != nil {
				fmt.Print("Некорректное число. Пожалуйста, введите числовое значение: ")
			} else {
				flag = true
			}
		}

		fmt.Print("Введите операция(+, -, *, /): ")
		var operation string
		flag = false
		for !flag {
			fmt.Scan(&operation)
			if operation == "+" || operation == "-" || operation == "*" || operation == "/" {
				flag = true
			} else {
				fmt.Print("Некорректная операция. Пожалуйста, используйте символы +, -, * или /: ")
			}
		}

		fmt.Print("Введите второе число: ")
		var second int
		flag = false
		for !flag {
			n, err := fmt.Scan(&second)
			if n != 1 || err != nil {
				fmt.Print("Некорректное число. Пожалуйста, введите числовое значение: ")
			} else if operation == "/" && second == 0 {
				fmt.Print("Ошибка: деление на ноль невозможно. Введите другое значение: ")
			} else {
				flag = true
			}
		}

		var itog int
		switch operation {
		case "+":
			itog = first + second
		case "-":
			itog = first - second
		case "*":
			itog = first * second
		case "/":
			itog = first / second
		}

		fmt.Println("Результат:", first, operation, second, "=", itog)

		fmt.Print("Продолжить использование программы? (y - да / n - нет): ")
		var tmp string
		fmt.Scan(&tmp)
		if tmp == "n" || tmp == "т" {
			flag_g = true
			fmt.Println("До скорых встреч))))")
		}

	}
}
