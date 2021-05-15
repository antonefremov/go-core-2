package storage

import "go-core-2/homeworks/05-gosearch-v3/pkg/crawler"

type Interface interface {
	Retrieve() []crawler.Document
	Save([]crawler.Document) error
}
