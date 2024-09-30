package main

// study, exercise 패키지를 import
import (
	"week1/exercise"
	"week1/study"
)

func main() {
	//함수의 첫 글자가 대문자여야 패키지 외부에서 접근할 수 있음

	study.PrintHelloWorld()
	// study.PrintVariablesTypes()
	//study.PrintVariableConstant()
	//study.PrintArray()
	// study.PrintIfSwitch()
	//study.PrintLoop()
	// study.PrintFunction()
	//study.PrintFmt()

	exercise.Exec1()
	//exercise.Exec2()
	//exercise.Exec3()
}
