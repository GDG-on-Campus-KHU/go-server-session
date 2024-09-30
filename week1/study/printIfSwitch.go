package study

import "fmt"

func PrintIfSwitch() {
	// if문
	age := 25
	if age >= 20 {
		fmt.Println("성인입니다.")
	} else {
		fmt.Println("미성년자입니다.")
	}

	// switch문
	name := "Alice"
	switch name {
	case "Bob":
		fmt.Println("안녕하세요, Bob!")
	case "Alice":
		fmt.Println("안녕하세요, Alice!")
	default:
		fmt.Println("안녕하세요! 근데 누구세요?")
	}
}
