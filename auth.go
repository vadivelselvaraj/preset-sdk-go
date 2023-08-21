package preset

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Gets a new JWT access token to interact with Preset/Superset APIs
func (c *PresetClient) GetAccessToken() (*AuthResponse, error) {
	if c.Auth.TokenName == "" || c.Auth.Secret == "" {
		return nil, fmt.Errorf("missing API token name and/or secret")
	}
	rb, err := json.Marshal(c.Auth)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/auth", c.BaseURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
