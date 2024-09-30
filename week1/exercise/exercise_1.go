package exercise

import "fmt"

func Exec1() {
	var num1, num2 float64
	var operator string

	fmt.Print("첫 번째 숫자를 입력하세요: ")
	fmt.Scan(&num1)

	fmt.Print("두 번째 숫자를 입력하세요: ")
	fmt.Scan(&num2)

	fmt.Print("연산자를 입력하세요 (+, -, *, /): ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f + %f = %f\n", num1, num2, num1+num2)
	case "-":
		fmt.Println(num1, "-", num2, "=", num1-num2)
	case "*":
		fmt.Println(num1, "*", num2, "=", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("0으로 나눌 수 없습니다.")
			return
		} else {
			fmt.Println(num1, "/", num2, "=", num1/num2)
		}
	default:
		fmt.Println("잘못된 연산자입니다.")
	}

}
