package technical

import "fmt"

type Config struct {
	Postgres Postgres
	Logger   Logger
	Kafka    Kafka
}

type Postgres struct {
	Host           string
	Port           int
	User, Password string
	Database       string
}

type Logger struct {
	Level string
	Path  string
}

type Kafka struct {
	Brokers []string
	Version string
}

func (p Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Database)
}
