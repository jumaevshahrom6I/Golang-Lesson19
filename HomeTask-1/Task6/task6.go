// 6. Дана строка с JSON: {"title":"Go для
// начинающих","author":"Иван Петров","year":2023} Опиши
// подходящую структуру с тегами и преврати строку в
// структуру с помощью json.Unmarshal. Выведи каждое поле
// отдельной строкой.

package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func main() {
	var book Book
	err := json.Unmarshal([]byte(`{"title":"C#", "author":"Shahrom", "year":2007}`), &book)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("Title: %v\n", book.Title)
	fmt.Printf("Author: %v\n", book.Author)
	fmt.Printf("Year: %v\n", book.Year)
}