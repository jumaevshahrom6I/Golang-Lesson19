// Задание 1. Каталог товаров с JSON-хранилищем
// Консольная программа. Меню:
// 1. Добавить товар
// 2. Показать все
// 3. Найти по категории
// 4. Удалить по ID
// 5. Сохранить и выйти
//
// Требования:
// 1. Структуры Category, Product (внутри Category) и
// Catalog (список товаров + следующий ID).
// 2. Теги: имена в snake_case; omitempty для описания;
// json:"-" для себестоимости (CostPrice) — она не должна
// попадать в файл.
// 3. Данные хранятся в catalog.json: загружаются при
// старте, сохраняются при выходе.
// 4. Если файла нет — программа стартует с пустым каталогом
// и не падает.
// 5. Все ошибки проверяются и оборачиваются с контекстом
// через %w.
//
// Готово, когда: данные сохраняются между запусками, в
// catalog.json нет себестоимости, программа не падает без
// файла.

package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Category struct {
	Name string `json:"name"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price"`
	CostPrice   float64   `json:"-"`
	Category    Category  `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
}

type Catalog struct {
	Products []Product `json:"products"`
	NextID   int       `json:"next_id"`
}

const fileName = "catalog.json"

var scanner = bufio.NewScanner(os.Stdin)

func ask(text string) string {
	fmt.Print(text)
	if !scanner.Scan() {
		os.Exit(0)
	}
	return strings.TrimSpace(scanner.Text())
}

func load() (Catalog, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Catalog{NextID: 1}, nil
		}
		return Catalog{}, fmt.Errorf("read file: %w", err)
	}

	var c Catalog
	if err := json.Unmarshal(data, &c); err != nil {
		return Catalog{}, fmt.Errorf("parse json: %w", err)
	}
	if c.NextID == 0 {
		c.NextID = 1
	}
	return c, nil
}

func save(c Catalog) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("make json: %w", err)
	}
	if err := os.WriteFile(fileName, data, 0644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	return nil
}

func add(c *Catalog) {
	name := ask("Name: ")
	desc := ask("Description (optional): ")
	cat := ask("Category: ")

	price, err := strconv.ParseFloat(ask("Price: "), 64)
	if err != nil {
		fmt.Println("bad price:", err)
		return
	}
	cost, err2 := strconv.ParseFloat(ask("Cost price: "), 64)
	if err2 != nil {
		fmt.Println("bad cost:", err2)
		return
	}

	p := Product{
		ID:          c.NextID,
		Name:        name,
		Description: desc,
		Price:       price,
		CostPrice:   cost,
		Category:    Category{Name: cat},
		CreatedAt:   time.Now(),
	}
	c.Products = append(c.Products, p)
	c.NextID++
	fmt.Println("added, ID:", p.ID)
}

func printProduct(p Product) {
	fmt.Printf("ID:%d | %s | %.2f | %s | %s\n",
		p.ID, p.Name, p.Price, p.Category.Name, p.Description)
}

func showAll(c Catalog) {
	if len(c.Products) == 0 {
		fmt.Println("empty")
		return
	}
	for _, p := range c.Products {
		printProduct(p)
	}
}

func find(c Catalog) {
	name := ask("Category: ")
	found := false
	for _, p := range c.Products {
		if strings.EqualFold(p.Category.Name, name) {
			printProduct(p)
			found = true
		}
	}
	if !found {
		fmt.Println("not found")
	}
}

func remove(c *Catalog) {
	id, err := strconv.Atoi(ask("ID: "))
	if err != nil {
		fmt.Println("bad ID:", err)
		return
	}
	for i, p := range c.Products {
		if p.ID == id {
			c.Products = append(c.Products[:i], c.Products[i+1:]...)
			fmt.Println("deleted")
			return
		}
	}
	fmt.Println("not found")
}

func main() {
	catalog, err := load()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Println("\n1. Add product")
		fmt.Println("2. Show all")
		fmt.Println("3. Find by category")
		fmt.Println("4. Delete by ID")
		fmt.Println("5. Save and exit")

		switch ask("Choice: ") {
		case "1":
			add(&catalog)
		case "2":
			showAll(catalog)
		case "3":
			find(catalog)
		case "4":
			remove(&catalog)
		case "5":
			if err := save(catalog); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("saved to", fileName)
			return
		default:
			fmt.Println("wrong choice")
		}
	}
}