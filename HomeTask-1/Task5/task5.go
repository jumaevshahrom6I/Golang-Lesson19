// 5. Опиши структуру Account с полями Login (string),
// Password (string) и Balance (float64). Для поля Password
// поставь тег json:"-", чтобы пароль не попадал в JSON.
// Создай значение, сериализуй и убедись, что пароля в
// выводе нет.

package main

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Login    string  `json:"login"`
	Password string  `json:"-"`
	Balance  float64 `json:"balance"`
}

func main() {
	account1 := Account{
		Login: "user2",
		Password: "1234",
		Balance: 100,
	}

	data, err2 := json.MarshalIndent(account1, "", "   ")
	if err2 != nil{
		fmt.Println(err2)
		return
	}
	fmt.Println(string(data))
}