package storage

import (
	"go-core-2/homeworks/06-gosearch-v4/pkg/crawler"
	"io"
)

type Interface interface {
	Retrieve(io.Reader) ([]crawler.Document, error)
	Save(io.Writer, []crawler.Document) error
}
