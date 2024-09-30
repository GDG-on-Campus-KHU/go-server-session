package study

import "fmt"

// 전역변수
var globalVar = "전역변수"

// 전역상수
const globalConst = "전역상수"

func PrintVariableConstant() {
	// 지역변수
	localVar := "지역변수"

	// 지역상수
	const localConst = "지역상수"

	fmt.Println(globalVar)
	fmt.Println(globalConst)
	fmt.Println(localVar)
	fmt.Println(localConst)
}
