package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

const URL = "https://www.aozora.gr.jp/index_pages/list_person_all.zip"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("コマンドを入力してください。")
		fmt.Println("例： aozora <作品名または著者名>")
		os.Exit(1)
	}
	key := os.Args[1]

	data, err := fetchCSV()
	if err != nil {
		log.Fatal(err)
	}

	books, err := parseCSV(bytes.NewReader(data))

	var matched []Book
	for _, b := range books {
		if strings.Contains(b.Author, key) || strings.Contains(b.Title, key) {
			matched = append(matched, b)
		}
	}

	if len(matched) == 0 {
		fmt.Printf("「%s」の作品が見つかりませんでした\n", key)
		os.Exit(1)
	}

	pick := matched[rand.Intn(len(matched))]

	fmt.Printf("作家: %s\n", pick.Author)
	fmt.Printf("作品: %s\n", pick.Title)
	fmt.Printf("出版社: %s\n", pick.Publisher)
}
