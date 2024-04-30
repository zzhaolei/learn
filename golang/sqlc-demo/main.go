package main

import (
	"context"
	"database/sql"
	"log"

	"sqlc-demo/model/dto"
	"sqlc-demo/model/mysql/author"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUrl := "root:12qwaszx@tcp(mysql:3306)/demo?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	queries := author.New(db)

	{
		authors, err := queries.ListAuthors(ctx)
		if err != nil {
			log.Fatal("获取 list 失败：", err)
		}
		log.Println("获取 list：", authors)
	}

	var id int64
	{
		bio := "测试用户"
		result, err := queries.CreateAuthor(ctx, author.CreateAuthorParams{
			Name: "测试用户",
			Bio:  &bio,
			Config: &dto.AuthorConfig{
				Name: "lei",
				Age:  18,
			},
		})
		if err != nil {
			log.Fatal("创建用户失败：", err)
		}
		id, err = result.LastInsertId()
		if err != nil {
			log.Fatal("获取插入 id 失败：", err)
		}
		log.Println("", id)
	}

	{
		resp, err := queries.GetAuthor(ctx, id)
		if err != nil {
			log.Fatal("获取指定用户失败：", err)
		}
		log.Println("用户：", resp)
		log.Println("用户 config：", resp.Config)

		err = queries.DeleteAuthor(ctx, id)
		if err != nil {
			log.Fatal("删除指定用户失败：", err)
		}
	}
}
