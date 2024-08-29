package env

import (
	"os"

	"github.com/joho/godotenv"
)

func New() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func Port() string {
	return os.Getenv("PORT")
}

func JWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func DBUser() string {
	return os.Getenv("PG_USER")
}

func DBPassword() string {
	return os.Getenv("PG_PASSWORD")
}

func DBName() string {
	return os.Getenv("PG_DB")
}

func DBHost() string {
	return os.Getenv("PG_HOST")
}

func DBConn() string {
	return os.Getenv("PG_CONN")
}

