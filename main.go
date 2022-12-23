package main

import (
	"fmt"
	"go-best-practice/internal/config"
	"go-best-practice/internal/entity"
	"go-best-practice/internal/request"
	"time"
)

func main() {
	appConf := config.InitConfig()

	config.InitDb(appConf)
	newBook := entity.Book{}
	newBook.Id = 1
	newBook.BookName = "Test"
	newBook.CreateDt = time.Now()
	newBook.UpdateDt = time.Now()

	Test[*entity.Book, *request.BookData](&newBook)

	server := config.InitServer()

	server.Logger.Fatal(server.Start(":" + appConf.Service.Port))
}

func Test[T any, K any](en entity.IBaseEntity[T, K]) {
	fmt.Printf("%v", en.ToResponse())
}
