// 3. Опиши структуру Product с полями ID (int), Name
// (string), Price (float64) и InStock (bool). С помощью
// тегов задай ключи id, name, price и in_stock. Создай
// значение и выведи его в красивом виде через
// json.MarshalIndent (отступ — два пробела).

package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	InStock bool    `json:"in_stock"`
}

func main() {
	product1 := Product{
		ID: 1,
		Name: "Shahrom",
		Price: 5,
		InStock: true,
	}

	a, err := json.MarshalIndent(product1, "", "  ")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(a))
}