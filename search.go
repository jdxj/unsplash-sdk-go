package unsplash_sdk_go

import (
	"context"
	"fmt"
)

// todo: 枚举
type SearchReq struct {
	Query         string   `url:"query,omitempty"`
	Collections   []string `url:"collections,omitempty"`
	ContentFilter string   `url:"content_filter,omitempty"`
	Color         string   `url:"color,omitempty"`
	Orientation   string   `url:"orientation,omitempty"`

	Lang string `url:"lang,omitempty"`

	Pagination
}

type SearchRsp struct {
}

func (c *Client) Search(ctx context.Context, req *SearchReq) (*SearchRsp, error) {
	rsp, err := c.r(ctx).
		SetQueryParamsFromValues(toValues(req)).
		Get(baseURL + "/search/photos")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("search err, status: %s", rsp.Status())
	}
	return &SearchRsp{}, nil
}
