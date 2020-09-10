package main

import (
	"DZ6/server"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	lg := NewLogger()

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:Anthony@localhost:27017"))
	if err != nil {
		lg.WithError(err).Fatal("Не удалось соединиться с БД")
		return
	}

	defer client.Disconnect(ctx)
	db := client.Database("messeges")
	srv, err := server.New(lg, ctx, db)
	if err != nil {
		lg.WithError(err).Fatal("Server start err.")
		return
	}
	lg.Debug("Server is started ...")

	srv.Serve(":8080")
}
