package preset

type Team struct {
	ID                    int    `json:"id"`
	Title                 string `json:"title"`
	AdminCount            int    `json:"admin_count"`
	BillingMethod         string `json:"billing_method"`
	BillingStatus         string `json:"billing_status"`
	CreatedOn             string `json:"created_on"`
	DefaultWorkspaceRole  Role   `json:"default_workspace_role"`
	DowngradedAt          string `json:"downgraded_at"`
	FeatureFlags          Flags  `json:"feature_flags"`
	Name                  string `json:"name"`
	PendingPurchaseType   string `json:"pending_purchase_type"`
	PlanCode              string `json:"plan_code"`
	RecurlyAccountID      string `json:"recurly_account_id"`
	SubscriptionStatus    string `json:"subscription_status"`
	Tier                  string `json:"tier"`
	TrialExpiry           string `json:"trial_expiry"`
	UserCount             int    `json:"user_count"`
	WhitelistedEmailDomains []string `json:"whitelisted_email_domains"`
	WorkspaceCount        int    `json:"workspace_count"`
	WorkspaceLimit        int    `json:"workspace_limit"`
	WorkspaceRoles        []Role `json:"workspace_roles"`
}

type Role struct {
	Name           string `json:"name"`
	RoleIdentifier string `json:"role_identifier"`
	RoleName       string `json:"role_name,omitempty"` // omitempty because this field is not present in all roles
}

type Flags struct {
	AiAssistEnabled          bool   `json:"ai_assist_enabled"`
	AlertReports             bool   `json:"alert_reports"`
	AuditLogEnabled          bool   `json:"audit_log_enabled"`
	AuditLogReleased         bool   `json:"audit_log_released"`
	ColorBuilderEnabled      bool   `json:"color_builder_enabled"`
	TeamRolesEnabled         bool   `json:"team_roles_enabled"`
	UsageMetricsDashEnabled  bool   `json:"usage_metrics_dash_enabled"`
	UsageMetricsLink         string `json:"usage_metrics_link"`
	UsageMetricsResourceID   string `json:"usage_metrics_resource_id"`
	UsageMetricsRlsRules     string `json:"usage_metrics_rls_rules"`
	WorkspaceRegionSelectEnabled bool   `json:"workspace_region_select_enabled"`
	WorkspaceRolesEnabled    bool   `json:"workspace_roles_enabled"`
}

type TeamResponse struct {
	Payload []Team `json:"payload"`
}

type TeamMembership struct {
	IsRoleFromGroup bool     `json:"is_role_from_group,omitempty"`
	TeamRole        TeamRole `json:"team_role"`
	User            User     `json:"user"`
}

type TeamRole struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	ID        int    `json:"id"`
	LastName  string `json:"last_name"`
	Onboarded bool   `json:"onboarded"`
	Username  string `json:"username"`
}

type TeamMembershipResponse struct {
	Payload []TeamMembership `json:"payload"`
}

type TeamMembershipUpdateResponse struct {
	Payload TeamMembership `json:"payload"`
}

type Workspace struct {
	ID                   int    `json:"id"`
	Accessible           bool   `json:"accessible"`
	AiAssistActivated    bool   `json:"ai_assist_activated"`
	AllowPublicDashboards bool  `json:"allow_public_dashboards"`
	ChangedOn            string `json:"changed_on"`
	ClusterID            int    `json:"cluster_id"`
	Color                string `json:"color"`
	CreatedOn            string `json:"created_on"`
	DeploymentID         int    `json:"deployment_id"`
	Descr                string `json:"descr"`
	Hostname             string `json:"hostname"`
	Icon                 string `json:"icon"`
	Maintenance          bool   `json:"maintenance"`
	Name                 string `json:"name"`
	Region               string `json:"region"`
	Status               string `json:"status"`
	TeamID               int    `json:"team_id"`
	Title                string `json:"title"`
	WorkspaceStatus      string `json:"workspace_status"`
}

type WorkspaceGetResponse struct {
	Payload []Workspace `json:"payload"`
}

type WorkspaceMembership struct {
	IsRoleFromGroup bool `json:"is_role_from_group,omitempty"`
	User            User `json:"user"`
	WorkspaceRole   WorkspaceRole `json:"workspace_role"`
}

type WorkspaceRole struct {
	Name           string `json:"name"`
	RoleIdentifier string `json:"role_identifier"`
}

type WorkspaceMembershipGetResponse struct {
	Payload []WorkspaceMembership `json:"payload"`
}

type WorkspaceMembershipUpdateResponse struct {
	Payload WorkspaceMembership `json:"payload"`
}
