package index

import (
	"fmt"
	"go-core-2/homeworks/06-gosearch-v4/pkg/crawler"
	"sort"
	"strings"
)

type crDocs []crawler.Document

// Store preserves the documents data
type Index struct {
	counter int
	docs    crDocs
	ind     map[uint64][]int
}

// New creates a new store instance
func New() *Index {
	return &Index{
		counter: 0,
		docs:    make([]crawler.Document, 0, 50),
		ind:     make(map[uint64][]int, 50),
	}
}

// Append adds document items to the store
func (s *Index) Append(docs []crawler.Document) {
	for _, d := range docs {
		s.counter++
		d.ID = s.counter
		s.docs = append(s.docs, d)
	}
}

// Build creates internal index for the docs
func (s *Index) Build() {
	for _, d := range s.docs {
		s.index(d.ID, d.Title)
	}

	// sort the documents
	s.Sort()
}

// IsEmpty indicates if docs array is empty
func (s *Index) IsEmpty() bool {
	return len(s.docs) == 0
}

// All retrieves the docs items
func (s *Index) All() []crawler.Document {
	return s.docs
}

// Search performs a search by the token passed in
func (s *Index) Search(token *string) []string {
	var d crawler.Document
	res := make([]string, 0, 10)
	h := hash(strings.ToLower(*token))
	ids := s.ind[h]

	for _, id := range ids {
		d = s.binarySearch(id, 0, len(s.docs))
		if d.ID != 0 {
			res = append(res, fmt.Sprintf("%d -> %s -> %s", d.ID, d.URL, d.Title))
		}
	}
	return res
}

// Sort sorts the store's docs array
func (s *Index) Sort() {
	sort.Sort(s.docs)
}

func (s *Index) binarySearch(id, l, r int) crawler.Document {
	if r < l {
		return crawler.Document{}
	}

	mid := l + (r-l)/2

	if s.docs[mid].ID == id {
		return s.docs[mid]
	}

	if id <= s.docs[mid].ID {
		// go left
		return s.binarySearch(id, l, mid-1)
	} else {
		// go right
		return s.binarySearch(id, mid+1, r)
	}
}

func (s *Index) index(id int, title string) {
	var h uint64
	title = strings.TrimRight(title, "\n")
	title = strings.Replace(title, "-", "", -1)
	title = strings.Replace(title, "&", "", -1)

	arr := strings.Split(title, " ")
	for _, t := range arr {
		h = hash(strings.ToLower(t))
		if h > 0 {
			if intArr, ok := s.ind[h]; !ok {
				intArr = make([]int, 0, 5)
				intArr = append(intArr, id)
				s.ind[h] = intArr
			} else {
				intArr = append(intArr, id)
				s.ind[h] = intArr
			}
		}
	}
}

// calculates polynomial hash
func hash(text string) uint64 {
	const (
		a = 123    // base value for hash
		m = 100003 // module on which hash is calculated
	)

	hashval := uint64(0)

	for _, r := range text {
		hashval = (hashval*a + uint64(r)) % m
	}

	return hashval
}

// methods below implement the sort.Interface
func (d crDocs) Len() int           { return len(d) }
func (d crDocs) Less(i, j int) bool { return d[i].ID < d[j].ID }
func (d crDocs) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
