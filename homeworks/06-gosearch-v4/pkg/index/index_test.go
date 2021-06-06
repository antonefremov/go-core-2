package index

import (
	"go-core-2/homeworks/06-gosearch-v4/pkg/crawler"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestIndex_Append(t *testing.T) {
	i := New()
	docs := []crawler.Document{{ID: 10}, {ID: 20}}
	i.Append(docs)
	got := len(i.docs)
	want := 2
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}

func TestIndex_Build(t *testing.T) {
	i := New()
	docs := []crawler.Document{{ID: 20, Title: "B"}, {ID: 5, Title: "A"}, {ID: 15, Title: "B"}, {ID: 25, Title: "C"}}
	i.Append(docs)
	i.Build()
	if len(i.ind[uint64('a')]) != 1 {
		t.Fatalf("получили %d, ожидалось %d", len(i.ind[uint64('a')]), 1)
	}
	if len(i.ind[uint64('b')]) != 2 {
		t.Fatalf("получили %d, ожидалось %d", len(i.ind[uint64('b')]), 2)
	}
}

func TestIndex_All(t *testing.T) {
	i := New()
	docs := []crawler.Document{{ID: 20, Title: "B"}, {ID: 5, Title: "A"}, {ID: 15, Title: "B"}}
	i.Append(docs)
	i.Build()
	isEmpty := i.IsEmpty()
	if isEmpty {
		t.Fatalf("получили %v, ожидалось %v", isEmpty, false)
	}
	res := i.All()
	if len(res) != 3 {
		t.Fatalf("получили %d, ожидалось %d", len(res), 3)
	}
}

func TestIndex_Search(t *testing.T) {
	i := New()
	docs := []crawler.Document{{ID: 20, Title: "B"}, {ID: 5, Title: "A"}, {ID: 15, Title: "B"}}
	i.Append(docs)
	i.Build()
	searchStr := "B"
	res := i.Search(&searchStr)
	if len(res) != 2 {
		t.Fatalf("получили %d, ожидалось %d", len(res), 2)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	data := sampleData()
	s := New()
	s.Append(data)
	s.Build()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1_000_000)
		res := s.Search(&data[n].Title)
		_ = res
	}
}

func sampleData() crDocs {
	rand.Seed(time.Now().UnixNano())
	var res crDocs
	var d crawler.Document
	for i := 0; i < 1_000_000; i++ {
		d.ID = rand.Intn(1_000_000)
		d.Title = RandStringBytesMaskImpr(10)
		res = append(res, d)
	}

	sort.Slice(res, func(i, j int) bool { return res[i].ID < res[j].ID })
	return res
}

// позаимствовано с https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandStringBytesMaskImpr(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
