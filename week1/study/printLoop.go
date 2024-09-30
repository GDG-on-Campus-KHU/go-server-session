package study

import "fmt"

func PrintLoop() {
	//For-in 루프
	names := []string{"Alice", "Bob", "Charlie"}
	for i := range names {
		fmt.Println(i, names[i])
	}

	// 오프셋 카운터에 직접 접근
	for i, v := range names {
		fmt.Println(i, v)
	}
	// 인덱스를 무시하고 값만 가져오기
	for _, v := range names {
		fmt.Println(v)
	}

	// c언어 스타일 루프
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//While 문과 비슷한 for 루프 (Go엔 while문이 없음)
	i := 0
	for i < 1004 {
		i++
	}
	fmt.Println(i)

	// 무한 루프 (while(true)처럼 사용)
	j := 0
	for {
		j++
		if j >= 2024 {
			break // i가 5 이상이면 루프 종료
		}
	}
	fmt.Println(j)

}
