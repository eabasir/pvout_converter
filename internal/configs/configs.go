package configs

import (
	"strconv"

	"github.com/joho/godotenv"
)

type Configs struct {
	Month             int
	DB_host           string
	DB_port           int
	DB_user           string
	DB_pass           string
	DB_name           string
	Skip_db_insertion bool
}

func GetConfigs() *Configs {
	var envFile, _ = godotenv.Read(".env")

	month, err := strconv.Atoi(envFile["MONTH"])
	db_port, err := strconv.Atoi(envFile["DB_PORT"])

	if err != nil {
		panic(err)
	}

	return &Configs{
		Month:             month,
		DB_host:           envFile["DB_HOST"],
		DB_port:           db_port,
		DB_user:           envFile["DB_USER"],
		DB_pass:           envFile["DB_PASS"],
		Skip_db_insertion: envFile["SKIP_DB_INSERTION"] == "true",
	}
}
