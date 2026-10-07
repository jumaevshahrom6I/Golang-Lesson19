// 2. Опиши структуру User с полями Name (string), Age (int)
// и Email (string). Добавь к каждому полю тег так, чтобы в
// JSON ключи были в нижнем регистре: name, age, email.
// Создай значение, сериализуй и выведи JSON.

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	user1 := User{
		Name: "Shahrom",
		Age: 19,
		Email: "shahromdzumaev61@gmail.com",
	}

	a, err := json.Marshal(user1)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(a))
	fmt.Println()

	user2 := User{
		Name: "Akbar",
		Age: 19,
		Email: "jumaevakbar61@gmail.com",
	}

	a, err2 := json.Marshal(user2)
	if err2 != nil{
		panic(err2)
	}
	fmt.Println(string(a))
	fmt.Println()

	var user3 User
	err3 := json.Unmarshal([]byte(`{"name":"Abdurahmon", "age": 7, "email":"bos@gmail.com"}`), &user3)
	if err3 != nil{
		fmt.Println(err3)
		return
	}
	fmt.Printf("%+v\n", user3)
}