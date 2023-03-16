package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/achintya-7/quiz/api"
	db "github.com/achintya-7/quiz/db/sqlc"
	"github.com/achintya-7/quiz/gapi"
	"github.com/achintya-7/quiz/pb"
	"github.com/achintya-7/quiz/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	runGRPC(config, store)
}

func runGRPC(config util.Config, store *db.Store) {
	grpcServer := grpc.NewServer()
	server, err := gapi.NewServer(store, config)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	pb.RegisterQuizServiceServer(grpcServer, server)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot listen to port", err)
	}

	fmt.Println("GRPC server listening on port", config.GRPCServerAddress)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}

func runHTTP(config util.Config, store *db.Store) {
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
