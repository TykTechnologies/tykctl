package cloudcmd

import (
	"fmt"
	"github.com/TykTechnologies/tykctl/internal"
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCloudRbacCloudRbac(t *testing.T) {
	tests := []struct {
		name           string
		MinAllowedUser Permission
		Role           string
		ExpectedError  error
	}{
		{
			name:           "Test all role are allowed",
			MinAllowedUser: TeamMember,
			Role:           "team_member",
			ExpectedError:  nil,
		},
		{
			name:           "Test Only OrgAdmins Are Allowed",
			MinAllowedUser: OrgAdmin,
			Role:           "team_admin",
			ExpectedError:  fmt.Errorf("user with role %s is not allowed to perform this action", "team_admin"),
		},
		{
			name:           "Test Only Org Admin are allowed ",
			MinAllowedUser: OrgAdmin,
			Role:           "team_member",
			ExpectedError:  fmt.Errorf("user with role %s is not allowed to perform this action", "team_member"),
		},
		{
			name:           "Test user with permission team admin and above allowed",
			MinAllowedUser: TeamAdmin,
			Role:           "team_member",
			ExpectedError:  fmt.Errorf("user with role %s is not allowed to perform this action", "team_member"),
		},
		{
			name:           "Test user with permission team admin is allowed",
			MinAllowedUser: TeamAdmin,
			Role:           "team_admin",
			ExpectedError:  nil,
		},
		{
			name:           "Test Org Admin is allowed when Min Is Team Admin",
			MinAllowedUser: TeamAdmin,
			Role:           "org_admin",
			ExpectedError:  nil,
		},
		{
			name:           "Test Invalid role Not allowed",
			MinAllowedUser: TeamMember,
			Role:           "invalid_role",
			ExpectedError:  fmt.Errorf("%s is invalid", "invalid_role"),
		},
		{
			name:           "Test Empty Role",
			MinAllowedUser: TeamMember,
			Role:           "",
			ExpectedError:  ErrorNoRoleFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockUserConfig(ctrl)
			c := NewCloudRbac(tt.MinAllowedUser, m)
			m.EXPECT().GetCurrentUserRole().Return(tt.Role)
			cmd := internal.NewCmd("test").NoArgs(nil)
			err := c.CloudRbac(cmd, []string{})
			assert.Equal(t, tt.ExpectedError, err)
		})
	}
}
