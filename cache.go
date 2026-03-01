package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const cacheExpiry = 7 * 24 * time.Hour

// cachePath は、キャッシュパスを取得します
func cachePath() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	cacheDir := filepath.Join(dir, "aozora")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(cacheDir, "list.csv"), nil
}

// isCacheValid は、CSVファイルがキャッシュパスに存在するかチェックします
func isCacheValid(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return time.Since(info.ModTime()) < cacheExpiry
}

// loadCache は、キャッシュからデータを読み込みます
func loadCache(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// saveCache は、キャッシュパスにデータを書き込みます
func saveCache(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

// download は、外部リンクからデータを取得します
func download() ([]byte, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, err
	}

	for _, f := range zipReader.File {
		if !strings.HasSuffix(f.Name, ".csv") {
      continue
    }

		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer rc.Close()

		return io.ReadAll(rc)
	}

	return nil, fmt.Errorf("ZIPの中にファイルが見つかりません")
}

// fetchCSV は、CSVデータを取得します
func fetchCSV() ([]byte, error) {
	path, err := cachePath()
	if err != nil {
		return nil, err
	}

	if isCacheValid(path) {
		fmt.Println("キャッシュから読み込み中")
		return loadCache(path)
	}

	fmt.Println("作品リストをダウンロード中")
	data, err := download()
	if err != nil {
		return nil, err
	}

	if err := saveCache(path, data); err != nil {
		return nil, err
	}

	return data, nil
}