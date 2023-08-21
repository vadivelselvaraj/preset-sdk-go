package preset

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Returns all workspaces tied to a given Preset team
func (c *PresetClient) GetAllWorkspaces(teamID int, authToken *string) (*[]Workspace, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/teams/%d/workspaces", c.BaseURL, teamID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tr := WorkspaceGetResponse{}
	err = json.Unmarshal(body, &tr)

	if err != nil {
		return nil, err
	}

	workspaces := tr.Payload
	return &workspaces, nil
}

// Returns all the members belonging to a given Preset workspace
func (c *PresetClient) GetWorkspaceMembership(teamID int, workspaceID int, authToken *string) (*[]WorkspaceMembership, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/teams/%d/workspaces/%d/memberships", c.BaseURL, teamID, workspaceID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	wmgr := WorkspaceMembershipGetResponse{}
	err = json.Unmarshal(body, &wmgr)
	if err != nil {
		return nil, err
	}

	memberships := wmgr.Payload
	return &memberships, nil
}

// Updates a given user's workspace role
func (c *PresetClient) UpdateUserWorkspaceRole(teamID int, workspaceID int, userID int, roleIdentifier string, authToken *string) (*WorkspaceMembership, error) {
	// Workspace role identifiers mapping
	workspaceRoleIdentifiers := map[string]string{
		"workspace admin": "Admin",
		"primary contributor": "PresetAlpha",
		"secondary contributor": "PresetBeta",
		"limited contributor": "PresetGamma",
		"viewer": "PresetReportsOnly",
		"dashboard viewer": "PresetDashboardsOnly",
		"no access": "PresetNoAccess",
	}

	// Validate and translate roleIdentifier
	translatedRole, found := workspaceRoleIdentifiers[roleIdentifier]
	if !found {
		return nil, fmt.Errorf("invalid role identifier")
	}

	// Create a map for the request payload
	payload := map[string]interface{}{
		"user_id":        userID,
		"role_identifier": translatedRole,
	}

	// Convert the payload map to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v1/team/%d/workspaces/%d/membership", c.BaseURL, teamID, workspaceID), bytes.NewReader(payloadBytes))
	
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	wmur := WorkspaceMembershipUpdateResponse{}
	err = json.Unmarshal(body, &wmur)
	if err != nil {
		return nil, err
	}

	membership := wmur.Payload
	return &membership, nil
}