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

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type User struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// func main() {
// 	users := []User{
// 		{Name: "Ali", Age: 20},
// 		{Name: "Bob", Age: 30},
// 	}

// 	f, err := os.Create("users.json")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()

// 	enc := json.NewEncoder(f)
// 	enc.SetIndent("", "  ")

// 	if err := enc.Encode(users); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("записано в users.json")

// 	// 1) без SetIndent: одна строка
// 	enc2 := json.NewEncoder(os.Stdout)
// 	enc2.Encode(users) // [{"name":"Ali"}]

// 	// 3) prefix ">>": виден в начале строк (кроме первой)
// 	enc3 := json.NewEncoder(os.Stdout)
// 	enc3.SetIndent(">>", "  ")
// 	enc3.Encode(users)

// 	// 4) Marshal с отступами
// 	fmt.Println()
// 	fmt.Println()
// 	b, _ := json.MarshalIndent(users, ">>", "  ")
// 	fmt.Println(string(b))
// }

// json.NewEncoder(f) создаёт кодировщик, который пишет JSON сразу в f (в файл), а не возвращает []byte, как Marshal.
// enc.SetIndent("", " ") включает форматирование: первый аргумент — префикс строки, второй — отступ (2 пробела).
// enc.Encode(users) превращает users в JSON и записывает в файл. Возвращает error.

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type Book struct{
// 	Title string `json:"title"`
// 	Author string `json:"author"`
// 	Year int `json:"year"`
// 	Price float64 `json:"price"`
// 	Note string `json:"note,omitempty"`
// }

// func save(filename string, v any) error{
// 	data, err := json.MarshalIndent(v, "", "  ")
// 	if err != nil{
// 		return fmt.Errorf("Upakovka v JSON: %w", err)
// 	}

// 	if err := os.WriteFile(filename, data, 0644); err != nil{
// 		return fmt.Errorf("Zapis faila: %w", err)
// 	}
// 	return nil
// }

// func load(filname string, v any) error{
// 	data, err := os.ReadFile(filname)
// 	if err != nil{
// 		return fmt.Errorf("Chteniya fayla: %w", err)
// 	}

// 	if err := json.Unmarshal(data, v); err != nil{
// 		return fmt.Errorf("Razbor JSON: %w", err)
// 	}
// 	return nil
// }

// func main() {
// 	books := []Book{
// 		{Title: "C#", Author: "Shahrom", Year: 2026, Price: 777, Note: "best"},
// 		{Title: "Golang", Author: "Akbar", Year: 2025, Price: 555, Note: "good"},
// 		{Title: "HTML", Author: "Abdullo", Year: 2024, Price: 888}, // Note пустой
// 		{Title: "C++", Author: "Abdurahmon", Year: 2027, Price: 999, Note: "hard"},
// 	}

// 	if err := save("shahrom.json", books); err != nil{
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("Soxroneno v shahrom.json")

// 	var book []Book
// 	if err := load("shahrom.json", &book); err != nil{
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println()
// 	fmt.Printf("Zagruzheno: %d | %+v\n", len(book), book)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type User struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// func main() {
// 	users := []User{{Name: "Ali", Age: 20}, {Name: "Bob", Age: 30}}

// 	// 1) NewEncoder: пишем в консоль (os.Stdout)
// 	enc := json.NewEncoder(os.Stdout)

// 	// 2) Encode: одна строка
// 	if err := enc.Encode(users); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// 3) SetIndent: красиво
// 	enc.SetIndent("", "  ")
// 	if err := enc.Encode(users); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// в файл: тот же Encoder, другой Writer
// 	f, err := os.Create("users.json")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()

// 	enc2 := json.NewEncoder(f)
// 	enc2.SetIndent("", "  ")
// 	if err := enc2.Encode(users); err != nil {
// 		fmt.Println(err)
// 	}
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type Movie struct {
// 	Title  string  `json:"title"`
// 	Year   int     `json:"year"`
// 	Rating float64 `json:"rating"`
// 	Note   string  `json:"note,omitempty"`
// }

// func main() {
// 	movies := []Movie{
// 		{Title: "Inception", Year: 2010, Rating: 8.8, Note: "mind-bending"},
// 		{Title: "Interstellar", Year: 2014, Rating: 8.7, Note: "space"},
// 		{Title: "The Matrix", Year: 1999, Rating: 8.7}, // Note пустой
// 		{Title: "Gladiator", Year: 2000, Rating: 8.5, Note: "epic"},
// 		{Title: "Joker", Year: 2019, Rating: 8.4}, // Note пустой
// 	}

// 	enc := json.NewEncoder(os.Stdout)

// 	if err := enc.Encode(movies); err != nil{
// 		fmt.Println(err)
// 		return
// 	}

// 	enc.SetIndent("", "  ")
// 	if err := enc.Encode(movies); err != nil{
// 		fmt.Println(err)
// 		return
// 	}

// 	f, err := os.Create("boss.json")
// 	if err != nil{
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()

// 	enc2 := json.NewEncoder(f)
// 	enc2.SetIndent("", "  ")
// 	if err := enc2.Encode(movies); err != nil{
// 		fmt.Println(err)
// 		return
// 	}
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type Car struct {
// 	Brand string `json:"brand"`
// }

// func main() {
// 	c := Car{Brand: "BMW"}
// 	enc := json.NewEncoder(os.Stdout)

// 	enc.Encode(c) // до SetIndent: {"brand":"BMW"}

// 	enc.SetIndent("", "    ")
// 	enc.Encode(c) // после: с отступом 4 пробела

// 	// Marshal: результат в памяти, печатаем сами
// 	b, _ := json.Marshal(c)
// 	fmt.Println(string(b))
// }

// package main
// import (
// 	"fmt"
// 	"encoding/json"
// 	"os"
// )

// type Phone struct {
// 	Brand string  `json:"brand"`
// 	Year  int     `json:"year"`
// 	Price float64 `json:"price"`
// 	Color string  `json:"color,omitempty"`
// }

// func main() {
// 	phones := []Phone{
// 	{Brand: "iPhone", Year: 2023, Price: 999, Color: "black"},
// 	{Brand: "Samsung", Year: 2022, Price: 799, Color: "blue"},
// 	{Brand: "Xiaomi", Year: 2021, Price: 399},  // Color пустой
// 	{Brand: "Google Pixel", Year: 2024, Price: 899, Color: "white"},
// 	{Brand: "Nokia", Year: 2020, Price: 199},   // Color пустой
// 	}

// 	f, err := os.Create("phones.json")
// 	if err != nil{
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()

// 	enc := json.NewEncoder(f)
// 	enc.SetIndent("", "   ")
// 	if err := enc.Encode(phones); err != nil{
// 		fmt.Println(err)
// 		return
// 	}

// 	enc2 := json.NewEncoder(os.Stdout)
// 	enc2.SetIndent("", "   ")
// 	if err := enc2.Encode(phones); err != nil{
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Finished")
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Car struct {
// 	Brand string `json:"brand"`
// 	Year  int    `json:"year"`
// 	Note  string `json:"note,omitempty"`
// }

// func main() {
// 	// 1) из файла
// 	f, err := os.Open("cars.json")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()

// 	var cars []Car
// 	dec := json.NewDecoder(f)
// 	if err := dec.Decode(&cars); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("count:", len(cars))
// 	for _, c := range cars {
// 		fmt.Printf("%s | %d | %s\n", c.Brand, c.Year, c.Note)
// 	}

// 	// 2) из строки + DisallowUnknownFields
// 	r := strings.NewReader(`{"brand":"Kia","year":2022,"color":"red"}`)
// 	var car Car
// 	dec2 := json.NewDecoder(r)
// 	dec2.DisallowUnknownFields() // до Decode
// 	if err := dec2.Decode(&car); err != nil {
// 		fmt.Println("error:", err) // json: unknown field "color"
// 		return
// 	}
// 	fmt.Println(car)
// }


// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Phone struct {
// 	Brand string  `json:"brand"`
// 	Year  int     `json:"year"`
// 	Price float64 `json:"price"`
// 	Color string  `json:"color,omitempty"`
// }

// func main() {
// 	f, err := os.Open("phones.json")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()

// 	var phones []Phone
// 	dec := json.NewDecoder(f)
// 	if err := dec.Decode(&phones); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Printf("Count: %d\n", len(phones))
// 	for _, p := range phones {
// 		fmt.Printf("Brand: %s | Year: %d | Price: %0.2f | Color: %s\n", p.Brand, p.Year, p.Price, p.Color)
// 	}

// 	test1 := `{"brand":"Iphone18", "year":2026, "price":27000, "ram": 8}`
// 	r := strings.NewReader(test1)
// 	var a Phone
// 	dec2 := json.NewDecoder(r)
// 	if err2 := dec2.Decode(&a); err2 != nil{
// 		fmt.Println(err2)
// 		return
// 	}
// 	fmt.Printf("%+v\n", a)
	
// 	test3 := `{"brand":"Iphone19", "year":2026, "price":27000, "color":"green"}`
// 	var c Phone
// 	dec4 := json.NewDecoder(strings.NewReader(test3))
// 	dec4.DisallowUnknownFields()
// 	if err3 := dec4.Decode(&c); err3 != nil{
// 		fmt.Println(err3)
// 		return
// 	}
// 	fmt.Printf("%+v\n", c)
	
// 	test2:= `{"brand":"Iphone18", "year":2026, "price":27000, "ram": 8}`
// 	var b Phone
// 	dec3 := json.NewDecoder(strings.NewReader(test2))
// 	dec3.DisallowUnknownFields()
// 	if err3 := dec3.Decode(&b); err3 != nil{
// 		fmt.Println(err3)
// 		return
// 	}
// 	fmt.Printf("%+v\n", b)
// }