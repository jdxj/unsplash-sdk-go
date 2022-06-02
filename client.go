package unsplash_sdk_go

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

const (
	baseURL       = "https://api.unsplash.com"
	AcceptVersion = "v1"
)

type OptFunc func(*option)

func WithDebug(debug bool) OptFunc {
	return func(o *option) {
		o.debug = debug
	}
}

type option struct {
	debug bool
}

func New(ak, sk string, optF ...OptFunc) *Client {
	opt := &option{}
	for _, f := range optF {
		f(opt)
	}

	c := &Client{
		accessKey: ak,
		secretKey: sk,
		opt:       opt,
		rc:        resty.New(),
	}

	if c.opt.debug {
		c.rc.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			log.Printf("--- OnBefore Begin ---")
			defer log.Printf("--- OnBefore End ---")

			return nil
		})

		c.rc.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
			log.Printf("--- OnAfter Begin ---")
			defer log.Printf("--- OnAfter End ---")

			log.Printf("header: %v", response.Header())
			log.Printf("url: %s", response.Request.URL)
			log.Printf("body: %s", response.Body())
			return nil
		})
	}
	return c
}

type Client struct {
	accessKey string
	secretKey string
	opt       *option

	rc *resty.Client
}

func (c *Client) r(ctx context.Context) *resty.Request {
	return c.rc.R().
		SetContext(ctx).
		SetHeader("Authorization", fmt.Sprintf("Client-ID %s", c.accessKey)).
		SetHeader("Accept-Version", AcceptVersion)
}

type Order string

const (
	Latest  Order = "latest"
	Oldest  Order = "oldest"
	Popular Order = "popular"
)

type Pagination struct {
	Page    int   `url:"page,omitempty"`
	PerPage int   `url:"per_page,omitempty"`
	OrderBy Order `url:"order_by,omitempty"`
}

func parseErrorMessage(data []byte) string {
	em := &errorMessage{}
	err := json.Unmarshal(data, em)
	if err != nil {
		log.Printf("parse error message failed: %s", err)
	}
	return em.String()
}

type errorMessage struct {
	Errors []string `json:"errors"`
}

func (em *errorMessage) String() string {
	if em == nil {
		return ""
	}
	return strings.Join(em.Errors, ",")
}

func toValues(i interface{}) url.Values {
	v, err := query.Values(i)
	if err != nil {
		log.Printf("to values err: %s", err)
		return nil
	}
	return v
}
