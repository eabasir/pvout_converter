package file_processor

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"pvout_converter/internal/db_manager"
	"pvout_converter/internal/types"
	"strconv"
	"strings"
)

const START_X = -180
const START_Y = -60
const INCREMENT = 1 / 120 // 30 arc-second
const HEADER_LINES = 6

func ProcessFile(file *os.File, month int, db *sql.DB, output_file *os.File) ([]types.PVData, error) {
	reader := bufio.NewReader(file)
	line_counter := 0
	results := make([]types.PVData, 0)
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

		fmt.Printf("processing line: %d\n", line_counter)

		line_results, err := process_line(month, line_counter, string(line))
		if err != nil {
			return nil, err
		}

		if db != nil {
			db_manager.Insert(db, line_results)
		}

		if output_file != nil {
			write_to_csv(output_file, line_results)
		}

	}

	return results, nil
}

func process_line(month int, line_counter int, line string) ([]types.PVData, error) {
	stringSlice := strings.Split(line, " ")
	latitude := START_Y + float64(line_counter)*INCREMENT
	res := make([]types.PVData, 0)
	for idx, str := range stringSlice {
		if str != "nan" && str != "" {
			longitude := START_X + float64(idx)*INCREMENT
			value, err := strconv.ParseFloat(str, 64)
			if err != nil {
				continue
			}
			res = append(res, types.PVData{Month: month, Latitude: latitude, Longitude: longitude, Value: value})
		}
	}
	return res, nil
}

func write_to_csv(file *os.File, data []types.PVData) {
	writer := bufio.NewWriter(file)

	for _, pv := range data {
		fmt.Fprintf(writer, "POINT(%f %f),%f,%d\n", pv.Longitude, pv.Latitude, pv.Value, pv.Month)
	}

	writer.Flush()

}
