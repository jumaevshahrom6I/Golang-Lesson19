// 4. Опиши структуру Student с полями Name (string), Grade
// (int) и Comment (string). Для поля Comment добавь опцию
// omitempty. Создай двух студентов: одного с комментарием,
// другого без. Сериализуй обоих и сравни, чем отличается
// JSON.

package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string `json:"name"`
	Grade   int    `json:"grade"`
	Comment string `json:"comment,omitempty"`
}

func main() {
	student1 := Student{
		Name: "Shahrom",
		Grade: 5,
		Comment: "You are best",
	}

	student2 := Student{
		Name: "Akbar",
		Grade: 4,
		Comment: "",
	}

	data, err := json.MarshalIndent(student1, "", "  ")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	fmt.Println()


	data, err2 := json.Marshal(student2)
	if err2 != nil{
		fmt.Println(err2)
		return
	}
	fmt.Println(string(data))
}
