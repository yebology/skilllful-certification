package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/yebology/skillful-certification/constant"
)

func LoadEnv() {

	envPaths := []string{
		filepath.Join(".", ".env"),
		filepath.Join("..", ".env"),
		filepath.Join("..", "..", ".env"),
	}

	for _, path := range envPaths {
		if _, err := os.Stat(path); err == nil {
			err = godotenv.Load(path)
			if err == nil {
				log.Println(constant.SuccessLoadEnvFile)
				return
			}
		}
	}
}

func LoadEnvConfig(name string) string {

	return os.Getenv(name)

}
