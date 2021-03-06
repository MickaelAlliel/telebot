package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"mickaelalliel.com/telebot/parser/internal/config"
	"mickaelalliel.com/telebot/parser/internal/ent"
)

func NewDatabaseOrFail() *ent.Client {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.DbConfig.Username,
		config.DbConfig.Password,
		config.DbConfig.Host,
		config.DbConfig.Port,
		config.DbConfig.Database)
	if config.DbConfig.UnixSocket != "" {
		log.Printf("using unix socket to connect to db")
		dsn = fmt.Sprintf("postgres:///?host=%s/%s&user=%s&password=%s&database=%s",
			config.DbConfig.UnixSocket,
			config.DbConfig.Host,
			config.DbConfig.Username,
			config.DbConfig.Password,
			config.DbConfig.Database)
	}

	log.Printf("opening connection to postgres db: %s:%d/%s", config.DbConfig.Host, config.DbConfig.Port, config.DbConfig.Database)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
		return nil
	}
	driver := entsql.OpenDB("postgres", db)
	client := ent.NewClient(ent.Driver(driver))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil
	} else {
		log.Println("migrated schema from generated ent files")
	}

	return client
}
