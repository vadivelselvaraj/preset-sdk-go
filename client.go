package preset

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// BaseURL - Default Preset URL
const APIURL string = "https://api.app.preset.io"
const ManagerURL string = "https://manage.app.preset.io"

// PresetClient
type PresetClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct
type AuthStruct struct {
	TokenName string `json:"name"`
	Secret string `json:"secret"`
}

type AuthPayload struct {
	AccessToken string `json:"access_token"`
}

type AuthResponse struct {
	Payload AuthPayload `json:"payload"`
}

// NewClient
func NewClient(host, tokenName, secret *string) (*PresetClient, error) {
	c := PresetClient{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		BaseURL: APIURL,
	}

	if host != nil {
		c.BaseURL = *host
	}

	// If tokenName or secret not provided, return empty client
	if tokenName == nil || secret == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		TokenName: *tokenName,
		Secret: *secret,
	}

	ar, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Payload.AccessToken

	return &c, nil
}

func (c *PresetClient) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
		req.Header.Set("Authorization", token)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
