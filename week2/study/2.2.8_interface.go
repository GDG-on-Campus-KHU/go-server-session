// < 2.2.8 인터페이스 >
package study

// 코드 2.24 : 인터페이스가 없는 구조체
type Person struct {
	Name string
	Age  int
}
type Dog struct {
	Name  string
	Owner *Person
	Age   int
}

func (person *Person) incrementAge() {
	person.Age++
}
func (person *Person) getAge() int {
	return person.Age
}
func (dog *Dog) incrementAge() {
	dog.Age++
}
func (dog *Dog) getAge() int {
	return dog.Age
}

// 코드 2.25 : 두 개의 구조체를 표준화하는 인터페이스
type Living interface {
	incrementAge()
	getAge() int
}

// 코드 2.26: 인터페이스에 맞는 구조체의 인스턴스로 변경하는 함수
func incrementAndPrintAge(being Living) {
	being.incrementAge()
	println(being.getAge())
}

// 코드 2.27: 두 개의 구조체 인스턴스로 incrementAndPrintAge 호출하기
func Code227() {
	harley := Person{"Harley", 21}
	snowy := Dog{"Snowy", &harley, 6}
	incrementAndPrintAge(&harley)
	incrementAndPrintAge(&snowy)
}
