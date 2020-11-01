package yugioh

import (
	"fmt"
	"net/http"
)

// CheckDBVersionService ...
type CheckDBVersionService struct {
	client *Client
}

// DBInformation ...
type DBInformation struct {
	Version     string `json:"database_version"`
	LastUpdated string `json:"last_update" validate:"datetime"`
}

// One ...
func (s *CheckDBVersionService) One() (DBInformation, *http.Response, error) {
	u := defaultBaseURL + "checkDBVer.php"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return DBInformation{}, nil, err
	}

	dbInformation := make([]DBInformation, 1)
	resp, err := s.client.Do(req, &dbInformation)
	if err != nil {
		return DBInformation{}, resp, err
	}
	if len(dbInformation) == 0 {
		err = fmt.Errorf("data does not exist")
		return DBInformation{}, resp, err
	}

	err = Validate(dbInformation[0])
	if err != nil {
		return DBInformation{}, resp, err
	}

	return dbInformation[0], resp, err
}
