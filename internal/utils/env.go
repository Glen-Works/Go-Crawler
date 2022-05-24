package utils

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetEnvData(parameter string) string {
	return os.Getenv(parameter)
}

// func ChechEnv() {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = godotenv.Load(dir + "/.env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }
