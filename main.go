package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/achintya-7/quiz/api"
	db "github.com/achintya-7/quiz/db/sqlc"
	"github.com/achintya-7/quiz/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	} else {
		fmt.Println("ENV variables loaded")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	} else {
		fmt.Println("Connected to DB")
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store, config)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}


