// 1. Опиши структуру Book с полями Title (string), Author
// (string), Year (int) и Price (float64). Создай значение,
// преврати его в JSON с помощью json.Marshal и выведи
// результат на экран как строку. Не забудь обработать
// ошибку.

package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Price  float64 `json:"price"`
}

func main() {
	book1 := Book{Title: "C#", Author: "Shahrom", Year: 2026, Price: 777}
	a, _ := json.Marshal(book1)
	fmt.Println(string(a))
	fmt.Println()

	book2 := Book{Title: "Golang", Author: "Akbar", Year: 2025, Price: 555}
	b, err1 := json.Marshal(book2)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(string(b))
	fmt.Println()

	book3 := Book{Title: "HTML", Author: "Abdullo", Year: 2024, Price: 888}
	c, err2 := json.Marshal(book3)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(string(c))
	fmt.Println()

	var book4 Book
	err3 := json.Unmarshal([]byte(`{"title":"C++", "author":"Abdurahmon", "year":2027, "price":999.0}`), &book4)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Printf("%+v\n", book4)
	fmt.Printf("%v", book4.Author)
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type User struct {
// 	Name     string `json:"name"`
// 	Age      int    `json:"age"`
// 	Password string `json:"-"`
// 	Email    string `json:"email,omitempty"`
// }

// type Test struct {
// 	Name string
// }

// func main() {
// 	u := User{Name: "Ali", Password: "secret"}

// 	a, _ := json.Marshal(u)
// 	fmt.Println(string(a))

// 	u.Email = "shahromdzumaev61@mail.com"
// 	a, _ = json.Marshal(u)
// 	fmt.Println(string(a))

// 	var u2 User
// 	json.Unmarshal([]byte(`{"name":"Bob","Password":"hack"}`), &u2)
// 	fmt.Printf("%+v\n", u2)
// 	fmt.Println("Name:", u2.Name)
// 	fmt.Println("Age:", u2.Age+1)

// 	u3 := Test{Name: "BOSS"}

// 	b, err := json.Marshal(u3)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(b))

// 	var u4 Test
// 	err2 := json.Unmarshal([]byte(`{"name":"Shahrom"}`), &u4)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 		return
// 	}
// 	fmt.Printf("%+v\n", u4)
// 	fmt.Printf("%#v\n", u4)
// }
