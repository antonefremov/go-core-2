package storage

import (
	"go-core-2/homeworks/05-gosearch-v3/pkg/crawler"
	"io"
)

type Interface interface {
	Retrieve(io.Reader) ([]crawler.Document, error)
	Save(io.Writer, []crawler.Document) error
}
