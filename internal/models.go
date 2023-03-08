package internal

import (
	"time"

	"github.com/TykTechnologies/cloud-sdk/cloud"
)

var (
	userRole   = "role"
	team       = "team"
	org        = "org"
	controller = "controller"
	env        = "env"
)

type OrgInfo struct {
	Organisation cloud.Organisation `json:"Organisation"`
}

type OrgInit struct {
	Controller string `json:"controller"`
	Org        string `json:"org"`
	Team       string `json:"team"`
	Env        string `json:"env"`
}

func (o OrgInit) OrgInitToMap() map[string]string {
	m := make(map[string]string)
	if o.Controller != "" {
		m[controller] = o.Controller
	}
	if o.Org != "" {
		m[org] = o.Org
	}
	if o.Team != "" {
		m[team] = o.Team
	}
	if o.Env != "" {
		m["env"] = o.Env
	}
	return m
}

type UserInfo struct {
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedAt       time.Time `json:"created_at"`
	PasswordUpdated time.Time `json:"password_updated"`
	Email           string    `json:"email"`
	LastName        string    `json:"lastName"`
	AccountID       string    `json:"account_id"`
	FirstName       string    `json:"firstName"`
	ID              string    `json:"id"`
	Roles           []Role    `json:"roles"`
	HubspotID       int       `json:"hubspot_id"`
	IsActive        bool      `json:"is_active"`
	IsEmailVerified bool      `json:"is_email_verified"`
}
type Role struct {
	Role      string `json:"role"`
	OrgID     string `json:"org_id"`
	TeamID    string `json:"team_id"`
	OrgName   string `json:"org_name"`
	TeamName  string `json:"team_name"`
	AccountID string `json:"account_id"`
}

func (r Role) RoleToMap() map[string]string {
	m := make(map[string]string)
	if r.Role != "" {
		m[userRole] = r.Role
	}
	if r.OrgID != "" {
		m[org] = r.OrgID
	}
	if r.TeamID != "" {
		m[team] = r.TeamID
	}
	return m
}

type ZoneResponse struct {
	Payload Payload `json:"Payload"`
	Error   string  `json:"error"`
}
type Payload struct {
	Tags map[string][]string `json:"Tags"`
}

type LoginBody struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	BasicAuthUserName string
	BasicAuthPassword string
}
