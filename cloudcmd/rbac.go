package cloudcmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type Permission int

const (
	TeamMember Permission = iota
	TeamAdmin
	OrgAdmin
)

type CloudRbac struct {
	MinAllowedUser Permission
}

func NewCloudRbac(MinAllowedUser Permission) CloudRbac {
	return CloudRbac{
		MinAllowedUser: MinAllowedUser,
	}
}

// CloudRbac is a will run during preRun to check if a role is allowed to perform an action.
func (c CloudRbac) CloudRbac(cmd *cobra.Command, args []string) error {
	role := getCurrentUserRole()
	if c.MinAllowedUser == TeamMember || role == "org_admin" {
		return nil
	}
	if c.MinAllowedUser == TeamAdmin && role == "team_admin" {
		return nil
	}
	return fmt.Errorf("user with role %s is not allowed to perform this action", role)
}
