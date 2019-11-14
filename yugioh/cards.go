package yugioh

import (
	"net/http"
)

// CardsService handles communication with the cards related
// methods of the Yu-Gi-Oh! API by YGOPRODeck.
type CardsService struct {
	client *Client
}

// Card ...
// If a piece of response info is empty or null then it will NOT show up.
type Card struct {
	ID          int32   `json:"id,string"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"desc"`
	Race        string  `json:"race"`
	Archetype   string  `json:"archetype"`
	Sets        []Set   `json:"card_sets"`
	Images      []Image `json:"card_images"`
	Prices      Prices  `json:"card_prices"`

	// Monster Card only
	Atk       int    `json:"atk,string"`
	Def       int    `json:"def,string"`
	Level     int    `json:"level,string"`
	Attribute string `json:"attribute"`

	// Pendulum Monster Card only
	Scale int `json:"scale,string"`

	// Link Monster Card only
	Link int `json:"linkval,string"`
	// Top, Bottom, Left, Right, Bottom-Left, Bottom-Right, Top-Left, Top-Right
	LinkMarkers []string `json:"linkmarkers"`

	// Ban Card only
	BanListInfo BanListInfo `json:"banlist_info"`
}

// Set ...
type Set struct {
	Name   string `json:"set_name"`
	Code   string `json:"set_code"`
	Rarity string `json:"set_rarity"`

	// Dollar
	Price float32 `json:"set_price,string"`
}

// Prices card prices
// Only Cardmarket price is euro.
type Prices struct {
	// Euro
	Cardmarket float32 `json:"cardmarket_price,string"`

	// Dollar
	Tcgplayer float32 `json:"tcgplayer_price,string"`
	Ebay      float32 `json:"ebay_price,string"`
	Amazon    float32 `json:"amazon_price,string"`
}

// BanListInfo if card not in ban list, it will NOT show up.
type BanListInfo struct {
	Tcg  string `json:"ban_tcg"`
	Ocg  string `json:"ban_ocg"`
	Goat string `json:"ban_goat"`
}

// Image card image.
type Image struct {
	ID       int32  `json:"id,string"`
	URL      string `json:"image_url"`
	SmallURL string `json:"image_url_small"`
}

// CardsListOptions specifies the optional parameters to various CardsService.List methods.
type CardsListOptions struct {
	ID        int32  `url:"name,omitempty"`
	Name      string `url:"name,omitempty"`
	KeyWord   string `url:"fname,omitempty"`
	Type      string `url:"type,omitempty"`
	Race      string `url:"race,omitempty"`
	Archetype string `url:"archetype,omitempty"`
	Set       string `url:"set,omitempty"`
	BanList   string `url:"banlist,omitempty"`

	// atk, def, name, type, level, id, new
	Sort string `url:"sort,omitempty"`

	// goat, ocg goat, speed duel, duel links
	// Duel Links is not 100% accurate but is close.
	Format string `url:"format,omitempty"`

	// English (Official), French
	Language string `url:"la,omitempty"`

	// Monster Card only
	Atk   int `url:"atk,omitempty"`
	Def   int `url:"def,omitempty"`
	Level int `url:"level,omitempty"`
	// dark, earth, fire, light, water, wind, divine
	Attribute string `url:"attribute,omitempty"`

	// Pendulum Monster Card only
	Scale int `json:"scale,omitempty"`

	// Link Monster Card only
	Link       int    `url:"link,omitempty"`
	LinkMarker string `url:"linkmarker,omitempty"`
}

// List the cards
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
