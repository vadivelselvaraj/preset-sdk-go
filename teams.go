package preset

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// TeamRoleEnum defines the valid role IDs
type TeamRoleEnum int

const (
	TEAM_ADMIN TeamRoleEnum = 1
	TEAM_USER TeamRoleEnum = 2
)

// Returns all Preset teams that the user admin token has access to
func (c *PresetClient) GetAllTeams(authToken *string) (*[]Team, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/teams", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tr := TeamResponse{}
	err = json.Unmarshal(body, &tr)

	if err != nil {
		return nil, err
	}

	teams := tr.Payload
	return &teams, nil
}

// Returns all the members belonging to a given team
func (c *PresetClient) GetTeamMembership(teamID int, workspaceID int, authToken *string) (*[]TeamMembership, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/team/%d/memberships", c.BaseURL, teamID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tmr := TeamMembershipResponse{}
	err = json.Unmarshal(body, &tmr)
	if err != nil {
		return nil, err
	}

	memberships := tmr.Payload
	return &memberships, nil	
}

// Updates a given user's team role
func (c *PresetClient) UpdateUserTeamRole(teamID int, userID int, roleID TeamRoleEnum, authToken *string) (*TeamMembership, error) {
	// Validate roleID
	if roleID != TEAM_ADMIN && roleID != TEAM_USER {
		return nil, fmt.Errorf("invalid role ID")
	}

	// Create a map for the request payload
	payload := map[string]interface{}{
		"team_role_id": roleID,
	}

	// Convert the payload map to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v1/team/%d/memberships/%d", c.BaseURL, teamID, userID), bytes.NewReader(payloadBytes))
	
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tmur := TeamMembershipUpdateResponse{}
	err = json.Unmarshal(body, &tmur)
	if err != nil {
		return nil, err
	}

	membership := tmur.Payload
	return &membership, nil
}

// Deletes a member from the team
func (c *PresetClient) DeleteTeamMembership(teamID int, userID int, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1/team/%d/memberships/%d", c.BaseURL, teamID, userID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	// Define a struct for parsing the response JSON
	type ApiResponse struct {
		Payload map[string]interface{} `json:"payload"`
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return err
	}

	if len(apiResponse.Payload) == 0 {
		return nil
	}

	return errors.New(string(body))
}
