package main

import (
	"errors"
	"flag"
	"fmt"
	"go-core-2/homeworks/05-gosearch-v3/pkg/crawler/spider"
	"go-core-2/homeworks/05-gosearch-v3/pkg/index"
	"go-core-2/homeworks/05-gosearch-v3/pkg/storage"
	"go-core-2/homeworks/05-gosearch-v3/pkg/storage/filestore"
	"os"
)

const path = "filestore.txt"

type gosearch struct {
	scanner *spider.Service
	sites   []string
	depth   int
	index   *index.Index
	store   storage.Interface
}

func main() {
	var token = flag.String("s", "", "search for a particular word/token")
	flag.Parse()
	if *token == "" {
		fmt.Println("exiting as no token to search for was provided by input")
		return
	}

	fileExists := true
	s := new()
	f, err := os.Open(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		fileExists = false
		fmt.Println("Error opening file:", err.Error()) // ничего страшного не произошло, просто информируем пользователя
	}
	if errors.Is(err, os.ErrNotExist) {
		fileExists = false
		fmt.Println("File doesn't exist") // опять же FYI
	}
	defer f.Close()

	if fileExists {
		docs, err := s.store.Retrieve(f)
		if err != nil {
			fileExists = false
			fmt.Println("Couldn't retrieve information from file") // тоже ФУЙ
		}
		s.index.Append(docs)
	}

	if !fileExists {
		fmt.Println("We couldn't retrieve pre-saved results (see the output above), so hold on tight!") // и тут ФУЙ
	}

	fmt.Println("Processing...")

	if s.index.IsEmpty() {
		for _, url := range s.sites {
			od, err := s.scanner.Scan(url, s.depth)
			if err != nil {
				panic(err) // а вот тут уже всё серьезно и стоит жестко выйти (наверное)
			}
			s.index.Append(od)
		}
	}

	// build index for the documents
	s.index.Build()

	fmt.Println("Search results:")
	res := s.index.Search(token)
	for _, d := range res {
		fmt.Println(d)
	}

	// save the indexed docs into the file storage
	f, err = os.Create(path) // а можно ли переиспользовать f? defer f.Close() ведь сработает только после выхода из текущей функции?
	if err != nil {
		fmt.Println("Couldn't create a file to store results", err) // FYI
	}
	err = s.store.Save(f, s.index.All())
	if err != nil {
		fmt.Println("Couldn't save results", err) // FYI
	}
	f.Close() // а если был defer f.Close() выше, то нужна ли инструкция f.Close() здесь?
}

func new() *gosearch {
	s := gosearch{}
	s.sites = []string{"https://go.dev", "https://golang.org/"}
	s.depth = 2
	s.scanner = spider.New()
	s.index = index.New()
	s.store = filestore.New()
	return &s
}
