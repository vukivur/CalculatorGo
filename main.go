package main

import (
	"fmt"
	"strconv"
)

func main() {

	var firstNumber string
	var secondNumber string
	var typeOfFirstNumber string
	var typeOfSecondNumber string
	var sign string
	arabicNumbers := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	romeNumbers := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	fmt.Println("Программа Калькулятор. Введите выражение. Принимаются числа от 1 до 10 или от I до X включительно.")
	fmt.Println("Что посчитать?:")
	fmt.Scan(&firstNumber, &sign, &secondNumber)

	// Проверка на корректность знака.
	if sign != "+" && sign != "-" && sign != "*" && sign != "/" {
		fmt.Println("Ошибка ввода данных. Так как введенный знак не является математической операцией.")
	}
	// Определяем тип введенных данных.
	for _, value := range arabicNumbers {
		if value == firstNumber {
			typeOfFirstNumber = "arabic"
		}
	}
	for _, value := range romeNumbers {
		if value == firstNumber {
			typeOfFirstNumber = "rome"
		}
	}
	for _, value := range arabicNumbers {
		if value == secondNumber {
			typeOfSecondNumber = "arabic"
			fmt.Println("secondNumber is: ", typeOfSecondNumber)
		}
	}
	for _, value := range romeNumbers {
		if value == secondNumber {
			typeOfSecondNumber = "rome"
			fmt.Println("secondNumber is: ", typeOfSecondNumber)
		}
	}
	// Проверка на корректный ввод данных.
	if (typeOfFirstNumber != "arabic" && typeOfFirstNumber != "rome") || (typeOfSecondNumber != "arabic" && typeOfSecondNumber != "rome") {
		fmt.Println("Ошибка ввода данных. Введены некорректные данные.")
	}
	// Проверка на разный тип данных.
	if (typeOfFirstNumber == "arabic" && typeOfSecondNumber == "rome") || (typeOfFirstNumber == "rome" && typeOfSecondNumber == "arabic") {
		fmt.Println("Ошибка ввода данных. Так как используются одновременно разные системы счисления.")
	}
	// Считаем результат арабских чисел.
	if typeOfFirstNumber == "arabic" && typeOfSecondNumber == "arabic" {
		// Конвертируем строку в число.
		number1, err := strconv.Atoi(firstNumber)
		if err != nil {
			panic(err)
			//fmt.Println(number1)
		}
		number2, err := strconv.Atoi(secondNumber)
		if err != nil {
			panic(err)
			//fmt.Println(number2)
		}
		switch sign {
		case "+":
			fmt.Println(number1 + number2)
		case "-":
			fmt.Println(number1 - number2)
		case "*":
			fmt.Println(number1 * number2)
		case "/":
			fmt.Println(number1 / number2)
		}
	}
	// Считаем результат римских чисел(сперва результат считаем в форме арабского числа. Потом конвертируем в арабский вид.).
	if typeOfFirstNumber == "rome" && typeOfSecondNumber == "rome" {
		var number1 int
		var number2 int
		romeNumbers100 := [100]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
		for i, value := range romeNumbers {
			if value == firstNumber {
				number1 = i + 1
			}
		}
		for i, value := range romeNumbers {
			if value == secondNumber {
				number2 = i + 1
			}
		}
		switch sign {
		case "+":
			fmt.Println(romeNumbers100[(number1+number2)-1])
		case "-":
			if number1-number2 >= 1 {
				fmt.Println(romeNumbers100[(number1-number2)-1])
			} else {
				fmt.Println("Ошибка ввода данных. Так как в римской системе нет отрицательных чисел.")
			}
		case "*":
			fmt.Println(romeNumbers100[(number1*number2)-1])
		case "/":
			if number1/number2 >= 1 {
				fmt.Println(romeNumbers100[(number1/number2)-1])
			} else {
				fmt.Println("Ошибка ввода данных. Так как в римской системе нет отрицательных чисел.")
			}
		}
	}

}
