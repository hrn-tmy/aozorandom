package main

import (
	"encoding/csv"
	"io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Book struct {
	Author    string
	Title     string
	Publisher string
}

// parseCSV は、CSVデータを解析し条件に合うデータを返します
func parseCSV(r io.Reader) ([]Book, error) {
	reader := transform.NewReader(r, japanese.ShiftJIS.NewDecoder())
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1

	// ヘッダースキップ
	if _, err := csvReader.Read(); err != nil {
		return nil, err
	}

	var books []Book
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(record) < 12 {
			continue
		}
		books = append(books, Book{
			Author:    record[1],
			Title:     record[3],
			Publisher: record[11],
		})
	}
	return books, nil
}