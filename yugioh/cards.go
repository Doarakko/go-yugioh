package yugioh

import (
	"net/http"
)

// CardsService ...
type CardsService struct {
	client *Client
}

// Card ...
type Card struct {
	ID          int32  `json:"id,string"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"desc"`
	Race        string `json:"race"`
	Archetype   string `json:"archetype"`
	Sets        []Set  `json:"card_sets"`
	// CardImages  CardImages  `json:"card_images"`
	// CardPrices []CardPrice `json:"card_prices"`
	// Monster
	Atk       int    `json:"atk,string"`
	Def       int    `json:"def,string"`
	Level     int    `json:"level,string"`
	Attribute string `json:"attribute"`
	// Pendulum Monster
	Scale int `json:"scale,string"`
	// Link Monster
	Link        int      `json:"linkval,string"`
	LinkMarkers []string `json:"linkmarkers"`
	// Ban card
	BanListInfo BanListInfo `json:"banlist_info"`
}

// Set ...
type Set struct {
	Name   string  `json:"set_name"`
	Code   string  `json:"set_code"`
	Rarity string  `json:"set_rarity"`
	Price  float32 `json:"set_price,string"`
}

// CardPrice ...
type CardPrice struct {
	CardmarketPrice float32 `json:"cardmarket_price,string"`
	TcgplayerPrice  float32 `json:"tcgplayer_price,string"`
	EbayPrice       float32 `json:"ebay_price,string"`
	AmazonPrice     float32 `json:"amazon_price,string"`
}

// BanListInfo ...
type BanListInfo struct {
	Tcg  string `json:"ban_tcg"`
	Ocg  string `json:"ban_ocg"`
	Goat string `json:"ban_goat"`
}

// CardImages ...
type CardImages struct {
	ID            int32  `json:"id"`
	ImageURL      string `json:"image_url"`
	ImageURLSmall string `json:"image_url_small"`
}

// CardsListOptions ...
type CardsListOptions struct {
	ID          int32  `url:"name,omitempty"`
	Name        string `url:"name,omitempty"`
	NameKeyWord string `url:"fname,omitempty"`
	Type        string `url:"type,omitempty"`
	Race        string `url:"race,omitempty"`
	Archetype   string `url:"archetype,omitempty"`
	CardSet     string `url:"set,omitempty"`
	BanList     string `url:"banlist,omitempty"`
	Sort        string `url:"sort,omitempty"`
	Format      string `url:"format,omitempty"`
	Language    string `url:"la,omitempty"`
	// Monster
	Atk       int    `url:"atk,omitempty"`
	Def       int    `url:"def,omitempty"`
	Level     int    `url:"level,omitempty"`
	Attribute string `url:"attribute,omitempty"`
	// Pendulum Monster
	Scale int `json:"scale,omitempty"`
	// Link Monster
	Link       int    `url:"link,omitempty"`
	LinkMarker string `url:"linkmarker,omitempty"`
}

// List ...
func (s *CardsService) List(opt *CardsListOptions) ([]Card, *http.Response, error) {
	u, err := addOptions("cardinfo.php", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	cards := new([]Card)
	resp, err := s.client.Do(req, cards)
	if err != nil {
		return nil, resp, err
	}
	return *cards, resp, err
}
