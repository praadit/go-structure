package main

import (
	"fmt"
	"go-best-practice/src/config"
	"go-best-practice/src/models/dto"
	"go-best-practice/src/models/entity"
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

	Test[*entity.Book, *dto.BookData](&newBook)

	server := config.InitServer()

	server.Logger.Fatal(server.Start(":" + appConf.Service.Port))
}

func Test[T any, K any](en entity.IBaseEntity[T, K]) {
	fmt.Printf("%v", en.ToResponse())
}
