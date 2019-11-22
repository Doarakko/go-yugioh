package yugioh

import "net/http"

// ArchetypesService handles communication with the archetype related
// methods of the Yu-Gi-Oh! API by YGOPRODeck.
type ArchetypesService struct {
	client *Client
}

// Archetype ...
type Archetype struct {
	Name string `json:"archetype_name"`
}

// List get all archetypes name.
// This method can NOT use options.
func (s *ArchetypesService) List() ([]Archetype, *http.Response, error) {
	u := defaultBaseURL + "archetypes.php"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	archetypes := new([]Archetype)
	resp, err := s.client.Do(req, archetypes)
	if err != nil {
		return nil, resp, err
	}

	return *archetypes, resp, err
}
