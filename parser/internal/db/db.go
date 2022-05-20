package db

import (
	"context"
	"fmt"
	"log"

	"mickaelalliel.com/telebot/parser/ent"
	"mickaelalliel.com/telebot/parser/internal/config"
)

func NewDatabaseOrFail() *ent.Client {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		config.DbConfig.Driver,
		config.DbConfig.Username,
		config.DbConfig.Password,
		config.DbConfig.Host,
		config.DbConfig.Port,
		config.DbConfig.Database)

	client, err := ent.Open(config.DbConfig.Driver, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
		return nil
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil
	}

	return client
}
