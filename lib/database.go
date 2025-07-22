package lib

import (
	"fmt"
	"log"
	"strconv"

	"pkg.formatio/prisma/db"
)

type BaseListFilterArgs struct {
	Skip   int    `query:"skip" swag-validate:"optional"`
	Take   int    `query:"take" swag-validate:"optional"`
	Search string `query:"search" swag-validate:"optional"`
	SortBy string `query:"sortBy" swag-validate:"optional"`
}

type DatabaseConnection struct {
	Client *db.PrismaClient
}

func NewDatabaseConnection(env Env) *DatabaseConnection {
	PG_PORT, _ := strconv.Atoi(env.PG_PORT)

	client := db.NewClient(db.WithDatasourceURL(GetPostgresConnectionString(PostgresConnectionParam{
		User:      env.PG_USER,
		Password:  env.PG_PASSWORD,
		Host:      env.PG_HOST,
		Port:      PG_PORT,
		Database:  env.PG_DB,
		ExtraArgs: fmt.Sprintf("sslmode=%s", env.DB_SSL_MODE),
	})))
	if err := client.Prisma.Connect(); err != nil {
		log.Println(err)
	}

	return &DatabaseConnection{Client: client}
}
