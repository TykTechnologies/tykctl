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

func TestCombineGetsValues(t *testing.T) {
	tests := []struct {
		name          string
		getkeys       []string
		getEnvValKeys []string
		want          []string
	}{
		{
			name:          "Test has both envVar and get keys",
			getkeys:       []string{"ExtraContext.Data.EnvData.TYK_DB_EMAILBACKEND_CODE", "ExtraContext.Data.EnvData.TYK_DB_EMAILBACKEND_DEFAULTFROMEMAIL"},
			getEnvValKeys: []string{"TYK_DB_EMAILBACKEND_ENABLEEMAILNOTIFICATIONS"},
			want:          []string{"ExtraContext.Data.EnvData.TYK_DB_EMAILBACKEND_CODE", "ExtraContext.Data.EnvData.TYK_DB_EMAILBACKEND_DEFAULTFROMEMAIL", "ExtraContext.Data.EnvData.TYK_DB_EMAILBACKEND_ENABLEEMAILNOTIFICATIONS"},
		},
		{
			name:          "Test has envVar only",
			getkeys:       []string{},
			getEnvValKeys: []string{"TYK_DB_EMAILBACKEND_ENABLEEMAILNOTIFICATIONS", "test"},
			want:          []string{"ExtraContext.Data.EnvData.TYK_DB_EMAILBACKEND_ENABLEEMAILNOTIFICATIONS", "ExtraContext.Data.EnvData.test"},
		},
		{
			name:          "Test has getVars only",
			getkeys:       []string{"ExtraContext.Data.EnvData.TYK_DB_EMAILBACKEND_CODE", "ExtraContext.Data.EnvData.TYK_DB_EMAILBACKEND_DEFAULTFROMEMAIL"},
			getEnvValKeys: nil,
			want:          []string{"ExtraContext.Data.EnvData.TYK_DB_EMAILBACKEND_CODE", "ExtraContext.Data.EnvData.TYK_DB_EMAILBACKEND_DEFAULTFROMEMAIL"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := combineGetsValues(tt.getkeys, tt.getEnvValKeys)
			assert.Equal(t, tt.want, results)
			assert.Equal(t, len(tt.want), len(results))
		})
	}
}
