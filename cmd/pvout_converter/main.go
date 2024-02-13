package main

import (
	"errors"
	"fmt"
	"os"
	"pvout_converter/internal/configs"
	"pvout_converter/internal/db_manager"
	"pvout_converter/internal/file_processor"
)

const FILE_PATH_PREFIX = "input/PVOUT_"

func main() {

	configs := configs.GetConfigs()

	filename, err := get_file_name(configs.Month)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Reading %s\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if !configs.Skip_db_insertion {
		db, err := db_manager.ConnectDB(configs)
		if err != nil {
			panic(err)
		}
		defer db_manager.CloseDB(db)

		file_processor.ProcessFile(file, configs.Month, db, nil)
	} else {

		output_file_name := fmt.Sprintf("output_%02d.csv", configs.Month)
		output_file, err := os.Create(output_file_name)
		if err != nil {
			panic(err)
		}
		defer output_file.Close()

		file_processor.ProcessFile(file, configs.Month, nil, output_file)
	}

}

func get_file_name(month int) (string, error) {

	if month < 1 || month > 12 {
		fmt.Println("input month is not in the range 1-12")
		return "", errors.New("month out of range")
	}

	return fmt.Sprintf("%s%02d.asc", FILE_PATH_PREFIX, month), nil

}
