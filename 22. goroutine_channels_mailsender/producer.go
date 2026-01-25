package main

import (
	"encoding/csv"
	"log"
	"os"
)

func sendReceipients(filename string, ch chan User) error {
	defer close(ch)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer file.Close()

	r := csv.NewReader(file)
	records, error := r.ReadAll()
	if error != nil {
		log.Fatal(error)
		return error
	}

	for _, record := range records[1:] {
		// Publish to channel or message queue
		ch <- User{
			Name:  record[0],
			Email: record[1],
		}
	}

	return nil
}
