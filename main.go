package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"golang.org/x/text/transform"

	"golang.org/x/text/encoding/japanese"
)

const URL = "https://www.aozora.gr.jp/index_pages/list_person_all.zip"

type Book struct {
	Author    string
	Title     string
	Publisher string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("コマンドを入力してください。")
		fmt.Println("例： aozora 夏目")
		os.Exit(1)
	}
	author := os.Args[1]

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		log.Fatal(err)
	}

	books := []Book{}
	for _, f := range zipReader.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()

		reader := transform.NewReader(rc, japanese.ShiftJIS.NewDecoder())
		csvReader := csv.NewReader(reader)
		csvReader.FieldsPerRecord = -1

		// ヘッダー行をスキップ
		if _, err := csvReader.Read(); err != nil {
			log.Fatal(err)
		}

		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			author := record[1]
			title := record[3]
			publisher := record[11]

			books = append(books, Book{
				Author:    author,
				Title:     title,
				Publisher: publisher,
			})
		}
	}

	var matched []Book
	for _, b := range books {
		if strings.Contains(b.Author, author) {
			matched = append(matched, b)
		}
	}

	if len(matched) == 0 {
		fmt.Printf("「%s」の作品が見つかりませんでした\n", author)
		os.Exit(1)
	}

	pick := matched[rand.Intn(len(matched))]

	fmt.Printf("作品: %s\n", pick.Title)
	fmt.Printf("出版社: %s\n", pick.Publisher)
}
