package web

import (
	"fmt"
)

type TracesDB string

const (
	TracesDBClickhouse = "clickhouse"
	TracesDBPostgres   = "postgres"
)

type Config struct {
	Postgres   Postgres
	Clickhouse Clickhouse
	Redis      Redis
	TracesDB   TracesDB
	Logger     Logger
	Kafka      Kafka
	Telegram   Telegram
	Port       int
}

type Postgres struct {
	Host           string
	Port           int
	User, Password string
	Database       string
}

type Clickhouse struct {
	Host           string
	Port           int
	User, Password string
	Database       string
	Timeout        int
}

type Redis struct {
	Address        string
	User, Password string
	DB             int
	TTL            int
}

type Logger struct {
	Level string
	Path  string
}

type Kafka struct {
	Brokers []string
	Version string
}

type Telegram struct {
	Enabled bool
	Token   string
	ChatID  int64
}

func (p Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Database)
}

func (c Clickhouse) DSN() string {
	return fmt.Sprintf("tcp://%s:%d/%s?username=%s&password=%s&dial_timeout=%dms", c.Host, c.Port, c.Database, c.User, c.Password, c.Timeout)
}
