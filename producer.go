package main

import (
	"encoding/csv"
	"os"
)

func loadRecipient(filePath string, ch chan Recipient, subject, body, fromName string) error {
	defer close(ch)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] { // skip header row
		ch <- Recipient{
			Name:     record[0],
			Email:    record[1],
			Subject:  subject,
			Body:     body,
			FromName: fromName,
		}
	}

	return nil
}
