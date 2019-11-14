package yugioh

import "net/http"

// RandomCardsService handles communication with the random cards related
// methods of the Yu-Gi-Oh! API by YGOPRODeck.
type RandomCardsService struct {
	client *Client
}

// One get random one card.
// This method can NOT use options.
func (s *RandomCardsService) One() (*Card, *http.Response, error) {
	u := defaultBaseURL + "randomcard.php"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	cards := new([]Card)
	resp, err := s.client.Do(req, cards)
	if err != nil {
		return nil, resp, err
	}
	return &(*cards)[0], resp, err
}
