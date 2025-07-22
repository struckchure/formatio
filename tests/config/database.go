package config

import (
	"context"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/docker/docker/pkg/ioutils"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
)

func NewTestDatabase(ctx context.Context, env lib.Env) *postgres.PostgresContainer {
	testcontainers.Logger = log.New(&ioutils.NopWriter{}, "", 0)

	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:15-alpine"),
		postgres.WithDatabase(env.PG_DB),
		postgres.WithUsername(env.PG_USER),
		postgres.WithPassword(env.PG_PASSWORD),
		testcontainers.WithWaitStrategy(
			// wait.ForExposedPort().
			// 	WithStartupTimeout(time.Second*10),
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)

	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	return container
}

func NewTestDatabaseConnection(ctx context.Context, env lib.Env, container *postgres.PostgresContainer) *lib.DatabaseConnection {
	dsn, _ := container.ConnectionString(ctx, "sslmode=disable")

	client := db.NewClient(db.WithDatasourceURL(dsn))
	if err := client.Prisma.Connect(); err != nil {
		log.Println("[NewTestDatabaseConnection]: ", err)
	}

	return &lib.DatabaseConnection{Client: client}
}

func MigrateDb(dsn string) {
	os.Setenv("DATABASE_URL", dsn)

	out, err := exec.Command("go", "run", "github.com/steebchen/prisma-client-go", "migrate", "deploy").Output()
	log.Println(string(out))

	if err != nil {
		log.Fatal("MigrateDb: ", err)
	}
}
