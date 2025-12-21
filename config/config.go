package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {} // empty struct cuz we don't need to store anything

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(fileNames ...string) Config {
	err := godotenv.Load(fileNames...)
	if err != nil {
		panic("Error loading .env file")
	}

	return &configImpl{}
}