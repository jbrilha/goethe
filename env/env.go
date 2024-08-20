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

func DBConn() string {
	return os.Getenv("PG_CONN")
}

func JWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
