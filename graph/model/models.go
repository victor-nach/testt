package model

//URL contains the mapping between a url and the shortcode
type URL struct {
	URL  string `json:"url"`
	Code string `json:"code"`
}
