package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"server/ent"
)

var EC *ent.Client

func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("database"), os.Getenv("password"))
	client, err := ent.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("failed opening db", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	EC = client
}
