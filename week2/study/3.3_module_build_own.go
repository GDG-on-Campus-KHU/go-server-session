// < 3.3 직접 만든 패키지 빌드하기 >
package study

// "github.com/hojelentil/go_study/study" //직접 패키지 만들어서 업로드한 뒤에 사용해보기

// 코드 3.3 : 사용자 정의 소수 확인 패키지
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

/* Git 저장소로 올리기
- github 저장소가 public인 경우 : 변경 사항 모두 커밋하고 Git 워크플로우가 요구하는 대로 저장소에 올리면 돤다.
- github 저장소가 private인 경우 : SSH키 이용하여 Git에 인증 -> 글로벌 Git config 명령어 -> sum check 수행하지 않게 설정 -> 새로운 Go 프로젝트 생성
->
*/

// 코드 3.4: 직접 만든 소수 확인 패키지 사용하기
// func Code34() {
// 	e := echo.New()
// 	e.GET("/:number", func(c echo.Context) error {
// 		nstr := c.Param("number")
// 		n, err := strconv.Atoi(nstr)
// 		if err != nil {
// 			return c.String(http.StatusBadRequest, err.Error())
// 		}
// 		return c.String(http.StatusOK, strconv.FormatBool(study.IsPrime(n)))
// 	})

// 	e.Logger.Fatal(e.Start(":8080"))
// }
