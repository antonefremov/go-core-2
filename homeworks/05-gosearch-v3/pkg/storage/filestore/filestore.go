package filestore

import (
	"bufio"
	"encoding/json"
	"go-core-2/homeworks/05-gosearch-v3/pkg/crawler"
	"io"
	"os"
)

const path = "filestore.txt"

func Save(d []crawler.Document) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}

	err = store(f, b)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

func Retrieve() ([]crawler.Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	b, err := get(f)
	if err != nil {
		return nil, err
	}
	f.Close()

	d := make([]crawler.Document, 0, 5)
	err = json.Unmarshal(b, &d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func store(w io.Writer, b []byte) error {
	_, err := w.Write(b)
	return err
}

func get(r io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(r)
	var b []byte
	for scanner.Scan() {
		b = append(b, []byte(scanner.Text()+"\n")...)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return b, nil
}
