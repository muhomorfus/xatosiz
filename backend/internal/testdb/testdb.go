package testdb

import (
	"context"
	"fmt"
	"os"

	"github.com/pressly/goose"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func startContainer() (testcontainers.Container, string, int, error) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "test",
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "password",
		},
	}

	dbContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		},
	)

	if err != nil {
		return nil, "", 0, fmt.Errorf("start db: %w", err)
	}

	host, _ := dbContainer.Host(context.Background())
	port, _ := dbContainer.MappedPort(context.Background(), "5432")

	return dbContainer, host, port.Int(), nil
}

func setup(host string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=user password=password dbname=test sslmode=disable", host, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open gorm: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get db: %w", err)
	}

	fmt.Println(os.Getwd())
	if err = goose.Up(sqlDB, "./deployments/migrations"); err != nil {
		return nil, fmt.Errorf("up migrations: %w", err)
	}

	text, err := os.ReadFile("./deployments/testdb/testdb.sql")
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	if err := db.Exec(string(text)).Error; err != nil {
		return nil, fmt.Errorf("exec: %w", err)
	}

	return db, nil
}

func New() (testcontainers.Container, *gorm.DB, error) {
	c, host, port, err := startContainer()
	if err != nil {
		return nil, nil, fmt.Errorf("start container: %w", err)
	}

	db, err := setup(host, port)
	if err != nil {
		return nil, nil, fmt.Errorf("setup db: %w", err)
	}

	return c, db, nil
}
