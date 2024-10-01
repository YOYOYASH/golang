package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title : "Go Intro",
		Author : "YOYOYASH",
		Copies: 1,
	}

	want := 0
	res,error := bookstore.Buy(b)
	got := res.Copies
	if error != nil {
		t.Fatal(error.Error())
	}
	if want != got {
		t.Errorf("Expected %d copies, got %d copies",want,got)
	}
}