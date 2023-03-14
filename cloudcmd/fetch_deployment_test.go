package cloudcmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TykTechnologies/cloud-sdk/cloud"
)

func TestCreateDeploymentHeadersAndRows(t *testing.T) {
	tests := []struct {
		name        string
		deployments []cloud.Deployment
		header      []string
		rows        [][]string
	}{
		{
			name: "",
			deployments: []cloud.Deployment{
				{
					Kind:        "Gateway",
					LoadoutName: "",
					Name:        "Deploy 1",
					State:       "success",
					UID:         "cfa0151a-ece8-40cf-a3ff-022851e0537f",
					ZoneCode:    "turkey",
				},
			},
			header: []string{"Name", "UID", "Kind", "Region", "State"},
			rows:   [][]string{{"Deploy 1", "cfa0151a-ece8-40cf-a3ff-022851e0537f", "Gateway", "turkey", "success"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CreateDeploymentHeadersAndRows(tt.deployments)
			assert.Equalf(t, tt.header, got, "CreateDeploymentHeadersAndRows(%v)", tt.deployments)
			assert.Equalf(t, tt.rows, got1, "CreateDeploymentHeadersAndRows(%v)", tt.deployments)
		})
	}
}
