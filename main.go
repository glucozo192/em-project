package main

import (
	"context"
	"log"

	"github.com/glu/shopvui/internal/services"
	"github.com/glu/shopvui/util"
	"github.com/jackc/pgx/v4"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := pgx.Connect(ctx, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	//store := db.NewStore(conn)
	//server, err := api.NewServer(config, *store)
	server, err := services.NewServer(config, conn)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
