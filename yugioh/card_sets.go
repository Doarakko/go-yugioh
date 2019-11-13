package yugioh

import "net/http"

// CardSetsService ...
type CardSetsService struct {
	client *Client
}

// CardSet ...
type CardSet struct {
	Name string `json:"Set Name"`
}

// List ...
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
