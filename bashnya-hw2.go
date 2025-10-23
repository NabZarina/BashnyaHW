package main

import "fmt"

func main() {
	var N int
	fmt.Print("Введите целое число")
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("Ошибка при чтении числа:", err)
		return
	}
	if N >= 12307 {
		fmt.Println("Введенное число больше 12307")
		return
	}

	for N < 12307 {
		if N < 0 {
			N = N * (-1)
		} else if N%7 == 0 {
			N = N * 39
		} else if N%9 == 0 {
			N = N*13 + 1
			continue
		} else {
			N = (N + 2) * 3
		}
		if N%13 == 0 && N%9 == 0 {
			fmt.Println("service error")
			break
		} else {
			N = N + 1
		}
	}
	fmt.Println("Результат работы цикла:", N)
}
