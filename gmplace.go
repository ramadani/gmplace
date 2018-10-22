package gmplace

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

func Autocomplete() {

}
