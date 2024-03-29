package cloudcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"

	"github.com/TykTechnologies/tykctl/internal"
)

type Permission int

const (
	TeamMember Permission = iota
	TeamAdmin
	OrgAdmin
)

type CloudRbac struct {
	MinAllowedUser Permission
	Config         internal.UserConfig
}

func NewCloudRbac(minAllowedUser Permission, config internal.UserConfig) CloudRbac {
	return CloudRbac{
		MinAllowedUser: minAllowedUser,
		Config:         config,
	}
}

// CloudRbac is a will run during preRun to check if a role is allowed to perform an action
// this will prevent you from sending a request that will fail due to lack of permissions.
func (c CloudRbac) CloudRbac(_ *cobra.Command, _ []string) error {
	role := c.Config.GetCurrentUserRole()

	if role == "" {
		return ErrorNoRoleFound
	}

	if !slices.Contains(cloudRoles, role) {
		return fmt.Errorf("%s is invalid", role)
	}

	if c.MinAllowedUser == TeamMember || role == "org_admin" {
		return nil
	}

	if c.MinAllowedUser == TeamAdmin && role == "team_admin" {
		return nil
	}

	return fmt.Errorf("user with role %s is not allowed to perform this action", role)
}
