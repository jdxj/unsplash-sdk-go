package unsplash_sdk_go

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
	SmallS3 string `json:"small_s3"`
}

type PhotoLinks struct {
	Self             string `json:"self"`
	Html             string `json:"html"`
	Download         string `json:"download"`
	DownloadLocation string `json:"download_location"`
}

type Exif struct {
	Make         string `json:"make"`
	Model        string `json:"model"`
	Name         string `json:"name"`
	ExposureTime string `json:"exposure_time"`
	Aperture     string `json:"aperture"`
	FocalLength  string `json:"focal_length"`
	Iso          int    `json:"iso"`
}

type Meta struct {
	Index bool `json:"index"`
}

type Photo struct {
	Id             string    `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	PromotedAt     time.Time `json:"promoted_at"`
	Width          int       `json:"width"`
	Height         int       `json:"height"`
	Color          string    `json:"color"`
	BlurHash       string    `json:"blur_hash"`
	Description    string    `json:"description"`
	AltDescription string    `json:"alt_description"`
	Likes          int       `json:"likes"`
	LikedByUser    bool      `json:"liked_by_user"`
	PublicDomain   bool      `json:"public_domain"`
	Views          int       `json:"views"`
	Downloads      int       `json:"downloads"`

	Urls                   Urls            `json:"urls"`
	Links                  PhotoLinks      `json:"links"`
	Categories             json.RawMessage `json:"categories"`
	CurrentUserCollections json.RawMessage `json:"current_user_collections"`
	Sponsorship            json.RawMessage `json:"sponsorship"`
	TopicSubmissions       json.RawMessage `json:"topic_submissions"`
	User                   User            `json:"user"`
	Exif                   Exif            `json:"exif"`
	Location               Location        `json:"location"`
	Mete                   Meta            `json:"meta"`
	Tags                   json.RawMessage `json:"tags"`
	TagsPreview            json.RawMessage `json:"tags_preview"`
	// todo: RelatedCollections
	Topics json.RawMessage `json:"topics"`
}

type PhotosReq struct {
}

type PhotosRsp struct {
}

func (c *Client) Photos() {

}

type GetPhotoReq struct {
	ID string
}

func (c *Client) GetPhoto(ctx context.Context, req *GetPhotoReq) (*Photo, error) {
	rsp, err := c.r(ctx).
		SetPathParam("id", req.ID).
		SetResult(&Photo{}).
		Get(baseURL + "/photos/{id}")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("status: %s, message: %v",
			rsp.Status(), parseErrorMessage(rsp.Body()))
	}
	return rsp.Result().(*Photo), nil
}

type GetRandomPhotoReq struct {
	Collections   []string `url:"collections,omitempty"`
	Topics        []string `url:"topics,omitempty"`
	Username      string   `url:"username,omitempty"`
	Query         string   `url:"query,omitempty"`
	Orientation   string   `url:"orientation,omitempty"`
	ContentFilter string   `url:"content_filter,omitempty"`
	Count         int      `url:"count,omitempty"`
}

func (c *Client) GetRandomPhoto(ctx context.Context, req *GetRandomPhotoReq) (*Photo, error) {
	rsp, err := c.r(ctx).
		SetQueryParamsFromValues(toValues(req)).
		SetResult(&Photo{}).
		Get(baseURL + "/photos/random")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("status: %s, message: %v",
			rsp.Status(), parseErrorMessage(rsp.Body()))
	}
	return rsp.Result().(*Photo), nil
}
