package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrgInitOrgInitToMap(t *testing.T) {

	tests := []struct {
		name    string
		orgInit OrgInit
		want    map[string]string
	}{
		{
			name: "Test has no org",
			orgInit: OrgInit{
				Controller: "https://tyk.io/eu-east",
				Org:        "",
				Team:       "My Team Here",
			},
			want: map[string]string{
				"controller": "https://tyk.io/eu-east",
				"team":       "My Team Here",
			},
		},
		{
			name: "Test Success",
			orgInit: OrgInit{
				Controller: "https://tyk.io/us-west",
				Org:        "DX org",
				Team:       "My Team",
			},
			want: map[string]string{"controller": "https://tyk.io/us-west",
				"team": "My Team",
				"org":  "DX org",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.orgInit.OrgInitToMap(), "OrgInitToMap()")
		})
	}
}

func TestRoleToMap(t *testing.T) {
	tests := []struct {
		name string
		role Role
		want map[string]string
	}{

		{
			name: "Test has empty team",
			role: Role{
				Role:      "org_admin",
				OrgID:     "89084736352",
				TeamID:    "",
				OrgName:   "Second Org",
				TeamName:  "Chose Team",
				AccountID: "94847472747",
			},
			want: map[string]string{
				"role": "org_admin",
				"org":  "89084736352",
			},
		},
		{
			name: "Test Success",
			role: Role{
				Role:      "team_member",
				OrgID:     "56ui84736",
				TeamID:    "jkituy6rt",
				OrgName:   "My Org",
				TeamName:  "New Team",
				AccountID: "890o47373",
			},
			want: map[string]string{
				"role": "team_member",
				"org":  "56ui84736",
				"team": "jkituy6rt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.role.RoleToMap(), "RoleToMap()")
		})
	}
}
