package lab

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func change(a *User, b *User) {
	temp := a.Name
	a.Name = b.Name
	b.Name = temp
}

func sorting(list []User) {
	for i := range list {
		for j := i + 1; j < len(list); j++ {
			if list[i].Name > list[j].Name {
				change(&list[i], &list[j])
			}
		}
	}
}

func Lab1() {
	list := []User{
		{"Paul", 19},
		{"John", 21},
		{"Jane", 35},
		{"Abraham", 25},
	}

	sorting(list)
	for _, user := range list {
		fmt.Println(user.Name)
	}
}
