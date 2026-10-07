// 7. Дана строка с JSON:
// {"name":"Алия","age":20,"city":"Душанбе"} Опиши структуру
// Person, у которой есть только поля Name и Age (поля city
// в структуре нет). Выполни Unmarshal и выведи результат.
// Что произошло с лишним ключом city? Напиши ответ
// комментарием в коде.

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var person Person
	err := json.Unmarshal([]byte(`{"name":"Shahrom", "age":19, "city":"Dushanbe"}`), &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("All:: %+v\n", person) // city проигнорирован: в структуре нет такого поля, ошибки нет
}
