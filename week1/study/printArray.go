package study

import "fmt"

func PrintArray() {
	fruits := [3]string{"사과", "바나나", "포도"} //배열 선언

	fmt.Println(fruits)
	fmt.Println(fruits[0])

	var names []string //슬라이스 선언
	names = []string{"신건우", "허윤지", "강해린"}
	names = append(names, "김민지") //append 함수를 사용하여 배열에 요소 추가

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	sports := make([]string, 3) //슬라이스 길이 지정
	fmt.Println("초기 길이:", len(sports), "초기 용량:", cap(sports))
	sports[0] = "축구"
	sports[1] = "야구"
	sports[2] = "농구"
	fmt.Println(sports)

	idol := make([]string, 0, 2) //슬라이스 길이와 용량 지정
	fmt.Println("초기 길이:", len(idol), "초기 용량:", cap(idol))
	// 요소를 추가
	idol = append(idol, "데이식스")
	idol = append(idol, "아이브")
	fmt.Println("추가 후 길이:", len(idol), "추가 후 용량:", cap(idol))
	// 용량이 넘는 요소를 추가 (현재 용량의 약 2배로 증가)
	idol = append(idol, "뉴진스")
	fmt.Println("추가 후 길이:", len(idol), "추가 후 용량:", cap(idol))
	fmt.Println(idol)

}
