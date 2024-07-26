package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := (stroka())
	if len(a[1]) > 10 || len(a[3]) > 10 {
		panic("Длина строки не может быть больше 10 символов")
	}

	fmt.Println("Результат:")
	b := calc(a[1], a[2], a[3])
	if len(b) > 40 {
		b = b[:40]
		b = b + "..."
		fmt.Println(strconv.Quote(b))
	} else {
		fmt.Println(strconv.Quote(b))
	}

}

func stroka() []string {
	var a string
	var b []string
	fmt.Println("Введите выражение:")
	a, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	b = strings.Split(a, "\"")

	b[2] = strings.Replace(b[2], " ", "", -1)

	if len(b[2]) > 1 {
		var c []string
		var i string
		c = strings.Split(b[2], "")
		b[2] = c[0]
		c = c[1:]
		c = c[:len(c)-2]
		i = strings.Join(c, "")
		b = append(b, i)

	}

	return b
}

func calc(a, b, c string) string {
	var d string
	switch b {
	case "+":
		d = a + c
	case "-":

		d = strings.Replace(a, c, "", -1)

	case "*":
		e, err := strconv.Atoi(c)
		if e >= 1 && e <= 10 {
			for i := 0; i < e; i++ {
				d += a
			}
			if err != nil {
				fmt.Println("Ошибка")
			}
		} else {
			panic("Число не от 1 до 10, либо множитель строка")
		}

	case "/":
		e, err := strconv.Atoi(c)
		if e >= 1 && e <= 10 {
			f := len(a) / e
			g := strings.Split(a, "")
			g = g[:f]
			d = strings.Join(g, "")
			if err != nil {
				fmt.Println("Ошибка")
			}
		} else {
			panic("Число не от 1 до 10, либо делитель строка")
		}

	}
	return d
}

//test
//"Hi World!" - "World!"
//"Bye-bye!" - "World!"
//"Golang" * 5
//"Example!!!" / 3

//"Golang" * 11
//"Golang" / 11
//"Hi World!" * "World!"
//"Hi World!" / "World!"
//"123456" * 7
