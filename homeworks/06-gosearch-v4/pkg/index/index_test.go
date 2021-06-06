package index

import (
	"go-core-2/homeworks/06-gosearch-v4/pkg/crawler"
	"testing"
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