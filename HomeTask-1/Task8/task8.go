// 8. Дана строка с JSON: {"name":"Ноутбук","price":4500.5}
// Опиши структуру Item с полями Name (string), Price
// (float64) и Count (int) с тегами name, price, count.
// Выполни Unmarshal и выведи структуру. Какое значение
// получило поле Count, которого нет в JSON? Напиши ответ
// комментарием в коде.

package main

import (
	"fmt"
	"encoding/json"
)


type Item struct{
	Name string `json:"name"`
	Price float64 `json:"price"`
	Count int `json:"count"`
}

func main() {
	var item Item
	err := json.Unmarshal([]byte(`{"name":"Iphone", "price":777}`), &item)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", item) // Count = 0: ключа count нет в JSON, поле осталось со значением по умолчанию, ошибки нет
}