package config

import (
	"context"
	"log"
	"sync"

	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var (
	ctx       = context.Background()
	container = NewTestDatabase(ctx, NewEnv())

	containerInit sync.Once
)

func GetSetup() (context.Context, *postgres.PostgresContainer) {
	containerInit.Do(func() {
		container = NewTestDatabase(ctx, NewEnv())
		ctx = context.Background()

		dsn, _ := container.ConnectionString(ctx, "sslmode=disable")

		MigrateDb(dsn)

		err := container.Snapshot(ctx, postgres.WithSnapshotName("genesis"))
		if err != nil {
			log.Fatal(err)
		}
	})

	return ctx, container
}
