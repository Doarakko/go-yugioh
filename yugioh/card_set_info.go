package yugioh

import "net/http"

// CardSetInfoService ...
type CardSetInfoService struct {
	client *Client
}

// CardSetInfo ...
type CardSetInfo struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float32 `json:"set_price,string"`
	SetName string  `json:"set_name"`
	Code    string  `json:"set_code"`
	Rarity  string  `json:"set_rarity"`
}

// CardSetInfoOneOptions ...
type CardSetInfoOneOptions struct {
	Code string `url:"setcode" validate:"required"`
}

// One ...
func (s *CardSetInfoService) One(opt *CardSetInfoOneOptions) (*CardSetInfo, *http.Response, error) {
	err := Validate(opt)
	if err != nil {
		return nil, nil, err
	}

	u, err := addOptions("cardsetsinfo.php", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
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
