package yugioh

import "net/http"

// CardSetInfoService ...
type CardSetInfoService struct {
	client *Client
}

// CardSetInfo ...
type CardSetInfo struct {
	ID     int     `json:"id"`
	Name   string  `json:"set_name"`
	Code   string  `json:"set_code"`
	Rarity string  `json:"set_rarity"`
	Price  float32 `json:"set_price,string"`
}

// CardSetInfoOneOptions ...
type CardSetInfoOneOptions struct {
	Code string `url:"setcode"`
}

// One ...
func (s *CardSetInfoService) One(opt *CardSetInfoOneOptions) (*CardSetInfo, *http.Response, error) {
	u, err := addOptions("cardsetsinfo.php", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	cardSetInfo := new(CardSetInfo)
	resp, err := s.client.Do(req, cardSetInfo)
	if err != nil {
		return nil, resp, err
	}

	return cardSetInfo, resp, err
}
