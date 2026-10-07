// 9. Опиши структуру Address с полями City и Street
// (string) и структуру Person с полями Name (string), Age
// (int) и Address типа Address. Задай теги: name, age,
// address, city, street. Создай значение Person, сериализуй
// его через MarshalIndent, а затем полученный JSON преврати
// обратно в новую структуру через Unmarshal и выведи город.

package main

import (
	"fmt"
	"encoding/json"
)

type Address struct{
	City string `json:"city"`
	Street string `json:"street"`
}

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Address Address `json:"address"`
}

func main() {
	person := Person{
		Name: "Shahrom",
		Age: 19 ,
		Address: Address{City: "Dushanbe", Street: "50-solagi TJK"},
	}

	data, err := json.MarshalIndent(person, "", "   ")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))


	var person2 Person
	if err2 := json.Unmarshal(data, &person2); err2 != nil{
		fmt.Println(err2)
		return
	}
	fmt.Printf("%+v\n", person2)
}