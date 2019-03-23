package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 123})
	// 忽略的字段为 0 或 空
	fmt.Println(Books{title:"Go 语言"})

	var book1 Books
	var book2 Books
	book1.title = "Redis 设计与实现"
	book2.title = "高性能 MySQL"
	println(book1.title)
	println(book2.title)
}
