// < 2.2.7 구조체 >
package study

import "fmt"

// 코드 2.19
type User struct { // User의 바이트 크기: string의 바이트 크기 + int의 바이트 크기 = 16바이트
	Name string
	Age  int
}

func nameAndAge(uid int) User {
	switch uid {
	case 0:
		return User{"Baheer Kamal", 24}
	case 1:
		return User{"Tanmay Bakshi", 16}
	default:
		return User{"", -1}
	}
}
func Code219() {
	user := nameAndAge(1)
	fmt.Println(user)
}

// 코드 2.20
func incrementAge(user User) {
	user.Age++
	fmt.Println(user.Age)
}
func Code220() {
	kathy := User{"Kathy", 19}
	incrementAge(kathy)
	fmt.Println(kathy.Age)
}

// 코드 2.21
func incrementAgePointer(user *User) {
	user.Age++
	fmt.Println(user.Age)
}
func Code221() {
	kathy := User{"Kathy", 19}
	incrementAgePointer(&kathy)
	fmt.Println(kathy.Age)
}

// 코드 2.22
func (user User) incrementAge() {
	user.Age++
	fmt.Println(user.Age)
}

func Code222() {
	kathy := User{"Kathy", 19}
	kathy.incrementAge()
	fmt.Println(kathy.Age)
}

// 코드 2.23
func (user *User) incrementAgePointerReceiver() {
	user.Age++
	fmt.Println(user.Age)
}
func Code223() {
	kathy := User{"Kathy", 19}
	kathy.incrementAgePointerReceiver()
	fmt.Println(kathy.Age)
}
