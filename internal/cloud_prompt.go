package internal

import "github.com/TykTechnologies/cloud-sdk/cloud"

//go:generate mockgen -source=cloud_prompt.go -destination=./mocks/cloud_prompt.go -package=mock CloudPrompt
type CloudPrompt interface {
	RegionPrompt(regions []string) (string, error)
	OrgPrompt(orgs []cloud.Organisation) (*cloud.Organisation, error)
	TeamPrompt(teams []cloud.Team) (*cloud.Team, error)
	EnvPrompt(envs []cloud.Loadout) (*cloud.Loadout, error)
	PerformActionPrompt() bool
}
