package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/eabasir/pvout_converter/internal/file_processor"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <month>")
		return
	}

	month_arg := os.Args[1]

	// month_arg := "1"

	month, err := strconv.Atoi(month_arg)

	if err != nil {
		fmt.Println("input month is not a number")
		return
	}

	file_processor.GetFileData(month)

}
