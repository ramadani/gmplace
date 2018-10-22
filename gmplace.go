package gmplace

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"
)

// GoogleMapsPlaceBaseAPI base api of google maps place
const GoogleMapsPlaceBaseAPI = "https://maps.googleapis.com/maps/api/place"

// Prediction entity
type Prediction struct {
	ID                   string               `json:"id"`
	Description          string               `json:"description"`
	MatchedSubstrings    []MatchedSubstring   `json:"matched_substrings"`
	PlaceID              string               `json:"place_id"`
	Reference            string               `json:"reference"`
	StructuredFormatting StructuredFormatting `json:"structured_formatting"`
	Terms                []Term               `json:"terms"`
	Types                []string             `json:"types"`
}

// MatchedSubstring entity
type MatchedSubstring struct {
	Length int `json:"length"`
	Offset int `json:"offset"`
}

// StructuredFormatting entity
type StructuredFormatting struct {
	MainText                  string             `json:"main_text"`
	MainTextMatchedSubstrings []MatchedSubstring `json:"main_text_matched_substrings"`
	SecondaryText             string             `json:"secondary_text"`
}

// Term entity
type Term struct {
	Offset int    `json:"offset"`
	Value  string `json:"value"`
}

// AutocompleteResult Google Maps Place Autocomplete results
type AutocompleteResult struct {
	Predictions []Prediction `json:"predictions"`
	Status      string       `json:"status"`
}

// GmPlace go package struct
type GmPlace struct {
	key        string
	httpClient *http.Client
}

// Autocomplete Google Maps Place Autocomplete
// https://developers.google.com/places/web-service/autocomplete
func (p *GmPlace) Autocomplete(input string) (*AutocompleteResult, error) {
	var result *AutocompleteResult

	u, err := url.Parse(GoogleMapsPlaceBaseAPI)
	if err != nil {
		log.Fatal(err)
	}

	u.Path = path.Join(u.Path, "/autocomplete/json")
	q := u.Query()
	q.Set("key", p.key)
	q.Set("input", input)
	u.RawQuery = q.Encode()

	resp, err := p.httpClient.Get(u.String())
	if err != nil {
		log.Println(err)
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buf, &result)

	return result, err
}

// New Google Maps Place
func New(key string) *GmPlace {
	return &GmPlace{key, &http.Client{
		Timeout: time.Second * 10,
	}}
}
