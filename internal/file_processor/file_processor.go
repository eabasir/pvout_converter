package file_processor

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PVData struct {
	month     int
	latitude  float64
	longitude float64
	value     float64
}

const FILE_PATH_PREFIX = "input/PVOUT_"
const START_X = -180
const START_Y = -60
const INCREMENT = 1 / 120 // 30 arc-second
const HEADER_LINES = 6

func GetFileData(month int) ([]PVData, error) {

	filename, err := get_file_name(month)

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	fmt.Printf("Reading %s\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	return readFile(file, month)

}

func readFile(file *os.File, month int) ([]PVData, error) {
	reader := bufio.NewReader(file)
	line_counter := 0
	results := make([]PVData, 0)
	total_results := 0
	for {
		line, isPrefix, err := reader.ReadLine()
		line_counter++
		if line_counter <= HEADER_LINES {
			continue
		}
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		for isPrefix {
			var more []byte
			more, isPrefix, err = reader.ReadLine()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				return nil, err
			}
			line = append(line, more...)
		}

		line_results, err := process_line(month, line_counter, string(line))
		if err != nil {
			return nil, err
		}
		total_results += len(line_results)
		fmt.Printf("valid values in line %d: %d\n", line_counter, total_results)

	}

	return results, nil
}

func process_line(month int, line_counter int, line string) ([]PVData, error) {
	stringSlice := strings.Split(line, " ")
	latitude := START_Y + float64(line_counter)*INCREMENT
	res := make([]PVData, 0)
	for idx, str := range stringSlice {
		if str != "nan" && str != "" {
			longitude := START_X + float64(idx)*INCREMENT
			value, err := strconv.ParseFloat(str, 64)
			if err != nil {
				continue
			}
			res = append(res, PVData{month, latitude, longitude, value})
		}
	}
	return res, nil
}

func get_file_name(month int) (string, error) {

	if month < 1 || month > 12 {
		fmt.Println("input month is not in the range 1-12")
		return "", errors.New("month out of range")
	}

	return fmt.Sprintf("%s%02d.asc", FILE_PATH_PREFIX, month), nil

}
