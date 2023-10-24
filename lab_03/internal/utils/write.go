package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteFile(filePath string, data [][]string, header []string) {
	file, err := os.Create(filePath)
	defer file.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(header); err != nil {
		log.Fatalln("failed while writing to file", err)
	}
	for _, record := range data {
		if err := writer.Write(record); err != nil {
			log.Fatalln("failed writing to file", err)
		}
	}
}
