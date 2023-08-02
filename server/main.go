package main

import (
	"fmt"
	"log"
	"os"
	"server/db"
	"server/ent"
	"server/router"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	var envName string
	if len(os.Args) > 0 {
		switch os.Args[1] {
		case "--dev":
			envName = ".env.dev"
		case "--test":
			envName = ".env.test"
		case "--prod":
			envName = ".env.prod"
		}
	}
	err := godotenv.Load(envName)
	if err != nil {
		log.Fatal(err)
	}
	db.InitDB()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	G := router.Setup()
	err := G.Run(":9000")
	defer func(EC *ent.Client) {
		err := EC.Close()
		if err != nil {
			return
		}
	}(db.EC)
	if err != nil {
		return
	}
}
