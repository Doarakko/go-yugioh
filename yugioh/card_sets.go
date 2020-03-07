package yugioh

import "net/http"

// CardSetsService handles communication with the card sets related
// methods of the Yu-Gi-Oh! API by YGOPRODeck.
type CardSetsService struct {
	client *Client
}

// CardSet ...
type CardSet struct {
	Name         string `json:"set_name"`
	Code         string `json:"set_code"`
	CardCounts   int    `json:"num_of_cards"`
	ReleasedDate string `json:"tcg_date"`
}

// List get all card sets name.
// This method can NOT use options.
func (s *CardSetsService) List() ([]CardSet, *http.Response, error) {
	u := defaultBaseURL + "cardsets.php"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	cardSets := new([]CardSet)
	resp, err := s.client.Do(req, cardSets)
	if err != nil {
		return nil, resp, err
	}

	return *cardSets, resp, err
}
