package configs

import (
	"github.com/joho/godotenv"
	"strconv"
)

type Configs struct {
	Month   int
	DB_host string
	DB_port int
	DB_user string
	DB_pass string
	DB_name string
}

func GetConfigs() *Configs {
	var envFile, _ = godotenv.Read(".env")

	month, err := strconv.Atoi(envFile["MONTH"])
	db_port, err := strconv.Atoi(envFile["DB_PORT"])

	if err != nil {
		panic(err)
	}

	return &Configs{
		Month:   month,
		DB_host: envFile["DB_HOST"],
		DB_port: db_port,
		DB_user: envFile["DB_USER"],
		DB_pass: envFile["DB_PASS"],
		DB_name: envFile["DB_NAME"],
	}
}
