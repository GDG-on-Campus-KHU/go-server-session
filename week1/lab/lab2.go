package lab

import "fmt"

func Lab2() {
	var number int

	fmt.Print("숫자를 입력하세요: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println(number, "는 짝수입니다.")
	} else {
		fmt.Println(number, "는 홀수입니다.")
	}
}
