package unsplash_sdk_go

import (
	"context"
	"fmt"
	"os"
	"testing"
)

var (
	client *Client
)

func TestMain(t *testing.M) {
	client = New(ak, sk, WithDebug(true))
	os.Exit(t.Run())
}

func TestClient_GetPhoto(t *testing.T) {
	rsp, err := client.GetPhoto(context.Background(), &GetPhotoReq{
		ID: "CoqJGsFVJtM",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("raw: %s\n", rsp.Urls.Raw)
	fmt.Printf("full: %s\n", rsp.Urls.Full)
	fmt.Printf("regular: %s\n", rsp.Urls.Regular)
	fmt.Printf("small: %s\n", rsp.Urls.Small)
	fmt.Printf("thumb: %s\n", rsp.Urls.Thumb)
	fmt.Printf("smalls3: %s\n", rsp.Urls.SmallS3)
}

func TestClient_GetRandomPhoto(t *testing.T) {
	rsp, err := client.GetRandomPhoto(context.Background(), nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("raw: %s\n", rsp.Urls.Raw)
	fmt.Printf("full: %s\n", rsp.Urls.Full)
	fmt.Printf("regular: %s\n", rsp.Urls.Regular)
	fmt.Printf("small: %s\n", rsp.Urls.Small)
	fmt.Printf("thumb: %s\n", rsp.Urls.Thumb)
	fmt.Printf("smalls3: %s\n", rsp.Urls.SmallS3)
}
