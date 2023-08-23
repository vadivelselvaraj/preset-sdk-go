package preset

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUserTeamRole_InvalidRoleID(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}	

	_, err := client.UpdateUserTeamRole(1, 2, 3, nil) // Invalid role ID
	assert.EqualError(t, err, "invalid role ID")
}

func TestUpdateUserTeamRole_SuccessfulResponse(t *testing.T) {
	// Mock a successful response payload
	mockResponse := []byte(`{"payload": {"user": {"id": 456, "email": "foo@example.com", "first_name": "foo", "last_name": "foo", "onboarded": false, "username": "samlp|rivian|foo@example.com"}, "team_role": {"id": 1, "name": "Admin"}}}`)

	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response
		w.WriteHeader(http.StatusOK)
		w.Write(mockResponse)
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}	
	
	membership, err := client.UpdateUserTeamRole(1, 456, TEAM_ADMIN, nil)
	assert.NoError(t, err)
	assert.NotNil(t, membership)
	assert.Equal(t, 1, membership.TeamRole.ID)
}

func TestUpdateUserTeamRole_InternalServerError(t *testing.T) {
	// Mock a 500 internal server error
	mockResponse := []byte(`Internal Server Error`)
		
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate an internal error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(mockResponse)
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}

	_, err := client.UpdateUserTeamRole(1, 456, TEAM_ADMIN, nil)
	assert.Error(t, err)
}

func TestGetAllTeams_SuccessfulResponse(t *testing.T) {
	// Mock a successful response payload
	mockResponse := []byte(`
	{
		"payload": [
			{
				"admin_count": 13,
				"billing_method": "INVOICE",
				"billing_status": "CURRENT",
				"created_on": "2021-05-17T21:36:16.216107",
				"default_workspace_role": {
					"name": "REPORTS_ONLY",
					"role_identifier": "PresetReportsOnly",
					"role_name": "Viewer"
				},
				"downgraded_at": null,
				"feature_flags": {
					"ai_assist_enabled": false,
					"alert_reports": true,
					"audit_log_enabled": true,
					"audit_log_released": false,
					"color_builder_enabled": true,
					"team_roles_enabled": true,
					"usage_metrics_dash_enabled": true,
					"usage_metrics_link": "https://2fa9923f.us2a.app.preset.io",
					"usage_metrics_resource_id": "93d85214-3761-43be-90ef-bb0851b3fb2f",
					"usage_metrics_rls_rules": "",
					"workspace_region_select_enabled": true,
					"workspace_roles_enabled": true
				},
				"id": 831,
				"name": "65006831",
				"pending_purchase_type": null,
				"plan_code": null,
				"recurly_account_id": null,
				"subscription_status": "PAID",
				"tier": "ENTERPRISE",
				"title": "Test Team",
				"trial_expiry": "2021-06-14T15:16:54.299781",
				"user_count": 1020,
				"whitelisted_email_domains": [
					"foo.com"
				],
				"workspace_count": 3,
				"workspace_limit": 3,
				"workspace_roles": [
					{
						"name": "Workspace Admin",
						"role_identifier": "Admin"
					},
					{
						"name": "Primary Contributor",
						"role_identifier": "PresetAlpha"
					},
					{
						"name": "Secondary Contributor",
						"role_identifier": "PresetBeta"
					},
					{
						"name": "Limited Contributor",
						"role_identifier": "PresetGamma"
					},
					{
						"name": "Viewer",
						"role_identifier": "PresetReportsOnly"
					},
					{
						"name": "Dashboard Viewer",
						"role_identifier": "PresetDashboardsOnly"
					},
					{
						"name": "No Access",
						"role_identifier": "PresetNoAccess"
					}
				]
			}
		]
	}`)	

	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with an empty array of teams
		w.WriteHeader(http.StatusOK)
		w.Write(mockResponse)
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}

	teams, err := client.GetAllTeams(nil)
	assert.NoError(t, err)
	assert.NotNil(t, teams)
	assert.Len(t, *teams, 1)
	assert.Equal(t, 831, (*teams)[0].ID)
	assert.Len(t, (*teams)[0].WorkspaceRoles, 7)
}

func TestGetAllTeams_InternalServerError(t *testing.T) {
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

	_, err := client.GetAllTeams(nil)
	assert.Error(t, err)
}

func TestGetTeamMembership_SuccessfulResponse(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with mock data
		response := []byte(`{
			"payload": [
				{
					"is_role_from_group": false,
					"team_role": {
						"id": 1,
						"name": "Admin"
					},
					"user": {
						"email": "test1@example.com",
						"first_name": "John",
						"id": 1670,
						"last_name": "Doe",
						"onboarded": true,
						"username": "auth0|60a2eb45f3e56c00686cb008"
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

	memberships, err := client.GetTeamMembership(1, 2, nil)
	assert.NoError(t, err)
	assert.NotNil(t, memberships)
	assert.Len(t, *memberships, 1)
	assert.Equal(t, 1, (*memberships)[0].TeamRole.ID)
	assert.Equal(t, 1670, (*memberships)[0].User.ID)
}

func TestGetTeamMembership_InternalServerError(t *testing.T) {
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

	_, err := client.GetTeamMembership(1, 2, nil)
	assert.Error(t, err)
}

func TestDeleteTeamMembership_SuccessfulResponse(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with mock data
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"payload": {}}`))
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}

	err := client.DeleteTeamMembership(1, 456, nil)
	assert.NoError(t, err)
}

func TestDeleteTeamMembership_InternalServerError(t *testing.T) {
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

	err := client.DeleteTeamMembership(1, 456, nil)
	assert.Error(t, err)
}
