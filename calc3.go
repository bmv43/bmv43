package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b, c, d int
	var op, z string
	var input string
	fmt.Println("Введите выражение")
	fmt.Scanln(&input)

	parts := strings.Split(input, "")
	for _, i := range parts {
		switch i {
		case "+":
			z = i
		case "-":
			z = i
		case "*":
			z = i
		case "/":
			z = i
		}
	}
	parts = strings.Split(input, z)
	var fff [3]string
	fff[0] = parts[0]
	fff[1] = z
	fff[2] = parts[1]

	switch fff[0] {
	case "1":
		a = 1
	case "2":
		a = 2
	case "3":
		a = 3
	case "4":
		a = 4
	case "5":
		a = 5
	case "6":
		a = 6
	case "7":
		a = 7
	case "8":
		a = 8
	case "9":
		a = 9
	case "10":
		a = 10
	case "I":
		b = 1
	case "II":
		b = 2
	case "III":
		b = 3
	case "IV":
		b = 4
	case "V":
		b = 5
	case "VI":
		b = 6
	case "VII":
		b = 7
	case "VIII":
		b = 8
	case "IX":
		b = 9
	case "X":
		b = 10
	default:
		panic("Неверное число")
		return
	}

	switch fff[2] {
	case "1":
		c = 1
	case "2":
		c = 2
	case "3":
		c = 3
	case "4":
		c = 4
	case "5":
		c = 5
	case "6":
		c = 6
	case "7":
		c = 7
	case "8":
		c = 8
	case "9":
		c = 9
	case "10":
		c = 10
	case "I":
		d = 1
	case "II":
		d = 2
	case "III":
		d = 3
	case "IV":
		d = 4
	case "V":
		d = 5
	case "VI":
		d = 6
	case "VII":
		d = 7
	case "VIII":
		d = 8
	case "IX":
		d = 9
	case "X":
		d = 10
	default:
		panic("Неверное число")
		return
	}

	if a != 0 && d != 0 {
		panic("Введите числа из одной системы счисления")
	}

	if b != 0 && c != 0 {
		panic("Введите числа из одной системы счисления")
	}

	op = fff[1]

	res := 0

	if a == 0 && c == 0 {
		switch op {
		case "+":
			res = b + d
		case "-":
			res = b - d
		case "*":
			res = b * d
		case "/":
			res = b / d
		default:
			panic("Неверная операция")
		}
	}

	if res == 1 {
		fmt.Println("I")
	} else if res == 2 {
		fmt.Println("II")
	} else if res == 3 {
		fmt.Println("III")
	} else if res == 4 {
		fmt.Println("IV")
	} else if res == 5 {
		fmt.Println("5")
	} else if res == 6 {
		fmt.Println("VI")
	} else if res == 7 {
		fmt.Println("VII")
	} else if res == 8 {
		fmt.Println("VIII")
	} else if res == 9 {
		fmt.Println("IX")
	} else if res == 10 {
		fmt.Println("X")
	} else if res == 11 {
		fmt.Println("XI")
	} else if res == 12 {
		fmt.Println("XII")
	} else if res == 13 {
		fmt.Println("XIII")
	} else if res == 14 {
		fmt.Println("XIV")
	} else if res == 15 {
		fmt.Println("XV")
	} else if res == 16 {
		fmt.Println("XVI")
	} else if res == 17 {
		fmt.Println("XVII")
	} else if res == 18 {
		fmt.Println("XVIII")
	} else if res == 19 {
		fmt.Println("XIX")
	} else if res == 20 {
		fmt.Println("XX")
	} else if res == 21 {
		fmt.Println("XXI")
	} else if res == 22 {
		fmt.Println("XXII")
	} else if res == 23 {
		fmt.Println("XXIII")
	} else if res == 24 {
		fmt.Println("XXIV")
	} else if res == 25 {
		fmt.Println("XXV")
	} else if res == 26 {
		fmt.Println("XXVI")
	} else if res == 27 {
		fmt.Println("XXVII")
	} else if res == 28 {
		fmt.Println("XXVIII")
	} else if res == 29 {
		fmt.Println("XXIX")
	} else if res == 30 {
		fmt.Println("XXX")
	} else if res == 31 {
		fmt.Println("XXXI")
	} else if res == 32 {
		fmt.Println("XXXII")
	} else if res == 33 {
		fmt.Println("XXXIII")
	} else if res == 34 {
		fmt.Println("XXXIV")
	} else if res == 35 {
		fmt.Println("XXXV")
	} else if res == 36 {
		fmt.Println("XXXVI")
	} else if res == 37 {
		fmt.Println("XXXVII")
	} else if res == 38 {
		fmt.Println("XXXVIII")
	} else if res == 39 {
		fmt.Println("XXXIX")
	} else if res == 40 {
		fmt.Println("XL")
	} else if res == 41 {
		fmt.Println("XLI")
	} else if res == 42 {
		fmt.Println("XLII")
	} else if res == 43 {
		fmt.Println("XLIII")
	} else if res == 44 {
		fmt.Println("XLIV")
	} else if res == 45 {
		fmt.Println("XLV")
	} else if res == 46 {
		fmt.Println("XLVI")
	} else if res == 47 {
		fmt.Println("XLVII")
	} else if res == 48 {
		fmt.Println("XLVIII")
	} else if res == 49 {
		fmt.Println("XLIX")
	} else if res == 50 {
		fmt.Println("L")
	} else if res == 51 {
		fmt.Println("LI")
	} else if res == 52 {
		fmt.Println("LII")
	} else if res == 53 {
		fmt.Println("LIII")
	} else if res == 54 {
		fmt.Println("LIV")
	} else if res == 55 {
		fmt.Println("LV")
	} else if res == 56 {
		fmt.Println("LVI")
	} else if res == 57 {
		fmt.Println("LVII")
	} else if res == 58 {
		fmt.Println("LVIII")
	} else if res == 59 {
		fmt.Println("LIX")
	} else if res == 60 {
		fmt.Println("LX")
	} else if res == 61 {
		fmt.Println("LXI")
	} else if res == 62 {
		fmt.Println("LXII")
	} else if res == 63 {
		fmt.Println("LXIII")
	} else if res == 64 {
		fmt.Println("LXIV")
	} else if res == 65 {
		fmt.Println("LXV")
	} else if res == 66 {
		fmt.Println("LXVI")
	} else if res == 67 {
		fmt.Println("LXVII")
	} else if res == 68 {
		fmt.Println("LXVIII")
	} else if res == 69 {
		fmt.Println("LXIX")
	} else if res == 70 {
		fmt.Println("LXX")
	} else if res == 71 {
		fmt.Println("LXXI")
	} else if res == 72 {
		fmt.Println("LXXII")
	} else if res == 73 {
		fmt.Println("LXXIII")
	} else if res == 74 {
		fmt.Println("LXXIV")
	} else if res == 75 {
		fmt.Println("LXXV")
	} else if res == 76 {
		fmt.Println("LXXVI")
	} else if res == 77 {
		fmt.Println("LXXVII")
	} else if res == 78 {
		fmt.Println("LXXVIII")
	} else if res == 79 {
		fmt.Println("LXXIX")
	} else if res == 80 {
		fmt.Println("LXXX")
	} else if res == 81 {
		fmt.Println("LXXXI")
	} else if res == 82 {
		fmt.Println("LXXXII")
	} else if res == 83 {
		fmt.Println("LXXXIII")
	} else if res == 84 {
		fmt.Println("LXXXIV")
	} else if res == 85 {
		fmt.Println("LXXXV")
	} else if res == 86 {
		fmt.Println("LXXXVI")
	} else if res == 87 {
		fmt.Println("LXXXVII")
	} else if res == 88 {
		fmt.Println("LXXXVIII")
	} else if res == 89 {
		fmt.Println("LXXXIX")
	} else if res == 90 {
		fmt.Println("XC")
	} else if res == 91 {
		fmt.Println("XCI")
	} else if res == 92 {
		fmt.Println("XCII")
	} else if res == 93 {
		fmt.Println("XCIII")
	} else if res == 94 {
		fmt.Println("XCIV")
	} else if res == 95 {
		fmt.Println("XCV")
	} else if res == 96 {
		fmt.Println("XCVI")
	} else if res == 97 {
		fmt.Println("XCVII")
	} else if res == 98 {
		fmt.Println("XCVIII")
	} else if res == 99 {
		fmt.Println("XCIX")
	} else if res == 100 {
		fmt.Println("C")
	}

	if b == 0 && d == 0 {
		switch op {
		case "+":
			res = a + c
			fmt.Println(res)
		case "-":
			res = a - c
			fmt.Println(res)
		case "*":
			res = a * c
			fmt.Println(res)
		case "/":
			res = a / c
			fmt.Println(res)
		default:
			panic("Неверная операция")
		}
	}

	if a == 0 && c == 0 && res < 1 {
		panic("Результат не может быть меньше единицы")
	}

}
