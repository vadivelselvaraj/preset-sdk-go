package preset

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllWorkspaces_SuccessfulResponse(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with mock data
		response := []byte(`{
			"payload": [
				{
					"id": 1,
					"accessible": true,
					"ai_assist_activated": true,
					"allow_public_dashboards": false,
					"changed_on": "2023-08-01T12:00:00Z",
					"cluster_id": 123,
					"color": "blue",
					"created_on": "2023-08-01T10:00:00Z",
					"deployment_id": 456,
					"descr": "Test Workspace 1",
					"hostname": "workspace1.example.com",
					"icon": "icon1",
					"maintenance": false,
					"name": "Workspace1",
					"region": "us-west",
					"status": "active",
					"team_id": 1,
					"title": "Workspace 1",
					"workspace_status": "ready"
				},
				{
					"id": 2,
					"accessible": true,
					"ai_assist_activated": false,
					"allow_public_dashboards": true,
					"changed_on": "2023-08-02T12:00:00Z",
					"cluster_id": 456,
					"color": "green",
					"created_on": "2023-08-02T10:00:00Z",
					"deployment_id": 789,
					"descr": "Test Workspace 2",
					"hostname": "workspace2.example.com",
					"icon": "icon2",
					"maintenance": true,
					"name": "Workspace2",
					"region": "us-east",
					"status": "inactive",
					"team_id": 1,
					"title": "Workspace 2",
					"workspace_status": "disabled"
				}
			]
		}`)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}

	workspaces, err := client.GetAllWorkspaces(1, nil)
	assert.NoError(t, err)
	assert.NotNil(t, workspaces)
	assert.Len(t, *workspaces, 2) // Check the number of workspaces in the response
	assert.Equal(t, 1, (*workspaces)[0].ID)
	assert.Equal(t, "Workspace1", (*workspaces)[0].Name)
}


func TestGetAllWorkspaces_InternalServerError(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a 500 internal server error
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}

	_, err := client.GetAllWorkspaces(1, nil)
	assert.Error(t, err)
}

func TestGetWorkspaceMembership_SuccessfulResponse(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with mock data
		response := []byte(`{
			"payload": [
				{
					"is_role_from_group": false,
					"user": {
						"email": "user1@example.com",
						"first_name": "User",
						"id": 123,
						"last_name": "One",
						"onboarded": true,
						"username": "user1"
					},
					"workspace_role": {
						"name": "Primary Contributor",
						"role_identifier": "PresetAlpha"
					}
				},
				{
					"is_role_from_group": true,
					"user": {
						"email": "user2@example.com",
						"first_name": "User",
						"id": 456,
						"last_name": "Two",
						"onboarded": false,
						"username": "user2"
					},
					"workspace_role": {
						"name": "Secondary Contributor",
						"role_identifier": "PresetBeta"
					}
				}
			]
		}`)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}

	memberships, err := client.GetWorkspaceMembership(1, 2, nil)
	assert.NoError(t, err)
	assert.NotNil(t, memberships)
	assert.Len(t, *memberships, 2) // Check the number of workspace memberships in the response
	// Add more assertions based on the expected response structure
	assert.Equal(t, false, (*memberships)[0].IsRoleFromGroup)
	assert.Equal(t, "user1@example.com", (*memberships)[0].User.Email)
	assert.Equal(t, "PresetBeta", (*memberships)[1].WorkspaceRole.RoleIdentifier)
}

func TestGetWorkspaceMembership_InternalServerError(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a 500 internal server error
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}

	_, err := client.GetWorkspaceMembership(1, 2, nil)
	assert.Error(t, err)
}

func TestUpdateUserWorkspaceRole_SuccessfulResponse(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with mock data
		response := []byte(`{
			"payload": {
				"is_role_from_group": false,
				"user": {
					"email": "user@example.com",
					"first_name": "User",
					"id": 123,
					"last_name": "One",
					"onboarded": true,
					"username": "user1"
				},
				"workspace_role": {
					"name": "Primary Contributor",
					"role_identifier": "PresetAlpha"
				}
			}
		}`)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}

	workspaceMembership, err := client.UpdateUserWorkspaceRole(1, 2, 123, "primary contributor", nil)
	assert.NoError(t, err)
	assert.NotNil(t, workspaceMembership)
	// Add more assertions based on the expected response structure
	assert.Equal(t, false, workspaceMembership.IsRoleFromGroup)
	assert.Equal(t, "user@example.com", workspaceMembership.User.Email)
	assert.Equal(t, "Primary Contributor", workspaceMembership.WorkspaceRole.Name)

}

func TestUpdateUserWorkspaceRole_InvalidRoleIdentifier(t *testing.T) {
	client := &PresetClient{
		BaseURL:    "mockBaseURL",
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}

	_, err := client.UpdateUserWorkspaceRole(1, 2, 123, "invalid role", nil)
	assert.Error(t, err)
}

func TestUpdateUserWorkspaceRole_InternalServerError(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a 500 internal server error
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}

	_, err := client.UpdateUserWorkspaceRole(1, 2, 123, "primary contributor", nil)
	assert.Error(t, err)
}
