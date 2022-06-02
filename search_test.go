package unsplash_sdk_go

import (
	"context"
	"testing"
)

func TestClient_Search(t *testing.T) {
	_, err := client.Search(context.Background(), &SearchReq{
		Query:         "apple",
		Collections:   nil,
		ContentFilter: "",
		Color:         "",
		Orientation:   "",
		Pagination:    Pagination{},
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
