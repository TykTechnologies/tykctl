package cloudcmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/TykTechnologies/cloud-sdk/cloud"
)

func RegionPrompt(regions []string) (string, error) {
	prompt := &survey.Select{
		Message: "Select your home region",
		Options: regions,
	}
	var region string
	err := survey.AskOne(prompt, &region, survey.WithValidator(survey.Required))
	if err != nil {
		return "", err
	}
	return region, nil
}
func OrgPrompt(orgs []cloud.Organisation) (*cloud.Organisation, error) {
	var organizations []string
	for _, org := range orgs {
		organizations = append(organizations, fmt.Sprintf("%s:%s", org.Name, org.UID))
	}
	prompt := &survey.Select{
		Message: "Select a default organization",
		Options: organizations,
	}
	var selectedIndex int
	err := survey.AskOne(prompt, &selectedIndex, survey.WithValidator(survey.Required))
	if err != nil {
		return nil, err
	}
	return &orgs[selectedIndex], nil
}

func teamPrompt(teams []cloud.Team) (*cloud.Team, error) {
	var teamString []string
	for _, team := range teams {
		teamString = append(teamString, fmt.Sprintf("%s:%s", team.Name, team.UID))
	}
	prompt := &survey.Select{
		Message: "Select a default team",
		Options: teamString,
	}
	var selectedIndex int
	err := survey.AskOne(prompt, &selectedIndex, survey.WithValidator(survey.Required))
	if err != nil {
		return nil, err
	}
	return &teams[selectedIndex], nil
}

func envPrompt(envs []cloud.Loadout) (*cloud.Loadout, error) {
	var loadoutString []string
	for _, loadout := range envs {
		loadoutString = append(loadoutString, fmt.Sprintf("%s:%s", loadout.Name, loadout.UID))
	}
	prompt := &survey.Select{
		Message: "Select a default environment",
		Options: loadoutString,
	}
	var selectedIndex int
	err := survey.AskOne(prompt, &selectedIndex, survey.WithValidator(survey.Required))
	if err != nil {
		return nil, err
	}
	return &envs[selectedIndex], nil
}

func emailPrompt() (string, error) {
	email := ""
	prompt := &survey.Input{
		Message: "Enter dashboard user email",
	}
	err := survey.AskOne(prompt, &email, survey.WithValidator(survey.Required))
	if err != nil {
		return "", err
	}
	return email, nil
}

func passwordPrompt() (string, error) {
	password := ""
	prompt := &survey.Password{
		Message: "Enter dashboard user password",
	}
	err := survey.AskOne(prompt, &password, survey.WithValidator(survey.Required))
	if err != nil {
		return "", err
	}
	return password, nil
}
