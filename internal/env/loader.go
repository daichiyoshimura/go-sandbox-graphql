package env

import (
	"errors"
	"fmt"
	"os"
	_ "github.com/joho/godotenv/autoload"
)

func Load() (*DB, *Server, error) {
	db, err := loadDB()
	if err != nil {
		return nil, nil, err
	}

	srv, err := loadServer()
	if err != nil {
		return nil, nil, err
	}

	return db, srv, nil
}

func load(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return "", errors.New(fmt.Sprintf("%s must be set", key))
	}
	return v, nil
}
