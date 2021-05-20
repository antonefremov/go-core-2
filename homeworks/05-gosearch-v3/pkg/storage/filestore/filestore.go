package filestore

import (
	"bufio"
	"encoding/json"
	"go-core-2/homeworks/05-gosearch-v3/pkg/crawler"
	"io"
)

type Filestore struct{}

func New() *Filestore {
	return &Filestore{}
}

func (f *Filestore) Save(w io.Writer, d []crawler.Document) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}

	err = store(w, b)
	if err != nil {
		return err
	}
	return nil
}

func (f *Filestore) Retrieve(r io.Reader) ([]crawler.Document, error) {
	b, err := get(r)
	if err != nil {
		return nil, err
	}

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
