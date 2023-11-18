package main

import (
	"context"
	"log"

	"prisma-go-demo/prisma/db"
)

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Println("connect err: ", err)
		return
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			log.Println("disconnnect err: ", err)
		}
	}()

	ctx := context.Background()
	created, err := client.Model.CreateOne(
		db.Model.Value.Set("test"),
	).Exec(ctx)
	if err != nil {
		log.Println("create one err: ", err)
		return
	}
	log.Println("create: ", created)

	find, err := client.Model.FindFirst(db.Model.Value.Equals("test")).Exec(ctx)
	if err != nil {
		log.Println("find first err: ", err)
		return
	}
	log.Println("find: ", find)

	updated, err := client.Model.FindMany(db.Model.ID.Equals(1)).Update(db.Model.Value.Set("update")).Exec(ctx)
	if err != nil {
		log.Println("update err: ", err)
		return
	}
	log.Println("update: ", updated)
}
