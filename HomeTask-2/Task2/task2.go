// Задание 2. Погода или пользователи из API
// Возьмите любой бесплатный API без ключа, например:
// - https://jsonplaceholder.typicode.com/posts?_limit=10
// - https://jsonplaceholder.typicode.com/comments?postId=1
//
// 1. Опишите структуру под ответ.
// 2. Сделайте запрос, проверьте err и StatusCode.
// 3. Разберите JSON в срез структур и выведите красиво.
// 4. Сохраните результат в файл result.json через
// MarshalIndent.
//
// (по желанию)
// - Добавьте в Product поле CreatedAt time.Time и
// посмотрите, как оно выглядит в JSON.
// - Добавьте в меню пункт «Загрузить товары из API» и
// сохраните их в каталог.


package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts?_limit=10")
	if err != nil {
		fmt.Println(fmt.Errorf("request: %w", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("bad status:", resp.StatusCode)
		return
	}

	var posts []Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		fmt.Println(fmt.Errorf("parse json: %w", err))
		return
	}

	for _, p := range posts {
		fmt.Printf("#%d (user %d)\n  %s\n  %s\n\n", p.ID, p.UserID, p.Title, p.Body)
	}

	data, err2 := json.MarshalIndent(posts, "", "  ")
	if err2 != nil {
		fmt.Println(fmt.Errorf("make json: %w", err2))
		return
	}
	if err3 := os.WriteFile("result.json", data, 0644); err3 != nil {
		fmt.Println(fmt.Errorf("write file: %w", err3))
		return
	}
	fmt.Println("saved to result.json")
}