package yugioh

import (
	"fmt"
	"net/http"
)

// CardsService handles communication with the cards related
// methods of the Yu-Gi-Oh! API by YGOPRODeck.
type CardsService struct {
	client *Client
}

// Cards ...
type Cards struct {
	Data []Card `json:"data"`
	Meta Meta   `json:"meta"`
}

// Card if a piece of response info is empty or null then it will NOT show up.
type Card struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"desc"`
	Race        string  `json:"race"`
	Archetype   string  `json:"archetype"`
	Sets        []Set   `json:"card_sets"`
	Images      []Image `json:"card_images"`

	// array, but this API currently only returns one
	Prices []Prices `json:"card_prices"`
	Misc   []Misc   `json:"misc_info"`

	// Monster Card only
	Atk       int    `json:"atk"`
	Def       int    `json:"def"`
	Level     int    `json:"level"`
	Attribute string `json:"attribute"`

	// Pendulum Monster Card only
	Scale int `json:"scale"`

	// Link Monster Card only
	Link int `json:"linkval"`
	// Top, Bottom, Left, Right, Bottom-Left, Bottom-Right, Top-Left, Top-Right
	LinkMarkers []string `json:"linkmarkers"`

	// Ban Card only
	BanListInfo BanListInfo `json:"banlist_info"`
}

// Set ...
type Set struct {
	Name       string `json:"set_name"`
	Code       string `json:"set_code"`
	Rarity     string `json:"set_rarity"`
	RarityCode string `json:"set_rarity_code"`

	// Dollar
	Price string `json:"set_price"`
}

// Prices card prices
// Only Cardmarket price is euro.
type Prices struct {
	// Euro
	Cardmarket float32 `json:"cardmarket_price,string"`

	// Dollar
	TCGPlayer    float32 `json:"tcgplayer_price,string"`
	Ebay         float32 `json:"ebay_price,string"`
	Amazon       float32 `json:"amazon_price,string"`
	CoolStuffInc float32 `json:"coolstuffinc_price,string"`
}

// BanListInfo if card not in ban list, it will NOT show up.
type BanListInfo struct {
	TCG  string `json:"ban_tcg"`
	OCG  string `json:"ban_ocg"`
	Goat string `json:"ban_goat"`
}

// Image card image.
type Image struct {
	ID       int32  `json:"id"`
	URL      string `json:"image_url"`
	SmallURL string `json:"image_url_small"`
}

// Misc if you want to get, set "yes" to Misc param
type Misc struct {
	// The Old/Temporary/Translated name
	BetaName string `json:"beta_name"`

	// The number of times a card has been viewed in ygoprodeck (does not include API/external views)
	Views int `json:"views"`

	// The available formats the card is in (tcg, ocg, goat, ocg goat, duel links, rush duel or speed duel).
	Formats []string `json:"formats"`

	// The original date the card was released in the TCG.
	TCGDate string `json:"tcg_date"`

	// The original date the card was released in the OCG.
	OCGDate string `json:"ocg_date"`

	// If the card is treated as another card.
	// For example, Harpie Lady 1,2,3 are treated as Harpie Lady.
	TreatedAs string `json:"treated_as"`
}

// Meta about pagination
type Meta struct {
	CurrentRows    int `json:"current_rows"`
	TotalRows      int `json:"total_rows"`
	RowsRemaining  int `json:"rows_remaining"`
	TotalPages     int `json:"total_pages"`
	PagesRemaining int `json:"pages_remaining"`

	// omitted on last page
	NextPage       string `json:"next_page,omitempty"`
	NextPageOffset int    `json:"next_page_offset,omitempty"`

	// omitted if not on the last page
	PreviousPage       string `json:"previous_page,omitempty"`
	PreviousPageOffset int    `json:"previous_page_offset,omitempty"`
}

// CardsListOptions specifies the optional parameters to various CardsService.List methods.
type CardsListOptions struct {
	ID        int32  `url:"name,omitempty"`
	Name      string `url:"name,omitempty"`
	Q         string `url:"fname,omitempty"`
	Type      string `url:"type,omitempty"`
	Race      string `url:"race,omitempty"`
	Archetype string `url:"archetype,omitempty"`
	Set       string `url:"set,omitempty"`
	BanList   string `url:"banlist,omitempty"`

	Sort string `url:"sort,omitempty" validate:"omitempty,oneof=atk def name type level id new"`

	// goat, ocg goat, speed duel, rush duel, duel links
	// Duel Links is not 100% accurate but is close.
	Format string `url:"format,omitempty"`

	// it can use "lt" (less than), "lte" (less than equals to), "gt" (greater than), "gte" (greater than equals to)
	Atk   string `url:"atk,omitempty"`
	Def   string `url:"def,omitempty"`
	Level string `url:"level,omitempty"`

	Attribute string `url:"attribute,omitempty" validate:"omitempty,oneof=dark earth fire light water wind divine"`

	// Pendulum Monster Card only
	Scale int `url:"scale,omitempty"`

	// Link Monster Card only
	Link       int    `url:"link,omitempty"`
	LinkMarker string `url:"linkmarker,omitempty"`

	// If set "yes", return Misc
	Misc string `url:"misc,omitempty" validate:"omitempty,eq=yes"`

	Staple string `url:"staple,omitempty"`

	Language string `url:"language,omitempty"`

	Num    int `url:"num,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// List the cards
func (s *CardsService) List(opt *CardsListOptions) (*Cards, *http.Response, error) {
	err := Validate(opt)
	if err != nil {
		return nil, nil, err
	}

	u, err := addOptions("cardinfo.php", opt)
	if err != nil {
		return nil, nil, err
	}
	// prevent to offset from being omitted
	if opt.Num > 0 && opt.Offset == 0 {
		u = fmt.Sprintf("%s&offset=0", u)
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	cards := new(Cards)
	resp, err := s.client.Do(req, cards)
	if err != nil {
		return nil, resp, err
	}

	return cards, resp, err
}
