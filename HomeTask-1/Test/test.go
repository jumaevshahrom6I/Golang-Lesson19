// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type User struct {
// 	Name  string   `json:"name"`
// 	Roles []string `json:"roles"`
// }

// func main() {
// 	u := User{Name: "Ali", Roles: []string{"admin"}}

// 	// 1) prefix "", indent 2 пробела
// 	b, _ := json.MarshalIndent(u, "", "  ")
// 	fmt.Println(string(b))

// 	// 2) prefix "", indent 4 пробела
// 	b, _ = json.MarshalIndent(u, "", "    ")
// 	fmt.Println(string(b))

// 	// 3) prefix "", indent таб
// 	b, _ = json.MarshalIndent(u, "", "\t")
// 	fmt.Println(string(b))
// 	fmt.Println()
// 	fmt.Printf("%q\n", b)
// 	fmt.Println()

// 	// 4) prefix "# ", indent 2 пробела
// 	b, _ = json.MarshalIndent(u, "# ", "  ")
// 	fmt.Println(string(b))

// 	// 5) indent "--" (чтобы увидеть уровни)
// 	b, _ = json.MarshalIndent(u, "", "--")
// 	fmt.Println(string(b))
// }




package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	users := []User{
		{Name: "Ali", Age: 20},
		{Name: "Bob", Age: 30},
	}

	f, err := os.Create("users.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")

	if err := enc.Encode(users); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("записано в users.json")

	// 1) без SetIndent: одна строка
	enc2 := json.NewEncoder(os.Stdout)
	enc2.Encode(users) // [{"name":"Ali"}]

	// 3) prefix ">>": виден в начале строк (кроме первой)
	enc3 := json.NewEncoder(os.Stdout)
	enc3.SetIndent(">>", "  ")
	enc3.Encode(users)

	// 4) Marshal с отступами
	fmt.Println()
	fmt.Println()
	b, _ := json.MarshalIndent(users, ">>", "  ")
	fmt.Println(string(b))
}

// json.NewEncoder(f) создаёт кодировщик, который пишет JSON сразу в f (в файл), а не возвращает []byte, как Marshal.
// enc.SetIndent("", " ") включает форматирование: первый аргумент — префикс строки, второй — отступ (2 пробела).
// enc.Encode(users) превращает users в JSON и записывает в файл. Возвращает error.
