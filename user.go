package unsplash_sdk_go

import (
	"encoding/json"
	"time"
)

type UserLinks struct {
	Self      string `json:"self"`
	Html      string `json:"html"`
	Photos    string `json:"photos"`
	Likes     string `json:"likes"`
	Portfolio string `json:"portfolio"`
	Following string `json:"following"`
	Followers string `json:"followers"`
}

type ProfileImage struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type Social struct {
	InstagramUsername string          `json:"instagram_username"`
	PortfolioUrl      string          `json:"portfolio_url"`
	TwitterUsername   string          `json:"twitter_username"`
	PaypalEmail       json.RawMessage `json:"paypal_email"`
}
type User struct {
	Id                string    `json:"id"`
	UpdatedAt         time.Time `json:"updated_at"`
	Username          string    `json:"username"`
	Name              string    `json:"name"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	TwitterUsername   string    `json:"twitter_username"`
	PortfolioUrl      string    `json:"portfolio_url"`
	Bio               string    `json:"bio"`
	Location          string    `json:"location"`
	InstagramUsername string    `json:"instagram_username"`
	TotalCollections  int       `json:"total_collections"`
	TotalLikes        int       `json:"total_likes"`
	TotalPhotos       int       `json:"total_photos"`
	AcceptedTos       bool      `json:"accepted_tos"`
	ForHire           bool      `json:"for_hire"`

	Links        UserLinks    `json:"links"`
	ProfileImage ProfileImage `json:"profile_image"`
	Social       Social       `json:"social"`
}

type Position struct {
	Latitude  interface{} `json:"latitude"`
	Longitude interface{} `json:"longitude"`
}

type Location struct {
	Title    json.RawMessage `json:"title"`
	Name     json.RawMessage `json:"name"`
	City     json.RawMessage `json:"city"`
	Country  json.RawMessage `json:"country"`
	Position Position        `json:"position"`
}
