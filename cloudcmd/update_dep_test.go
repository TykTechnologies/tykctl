package cloudcmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
)

func TestHandleEnvVariables(t *testing.T) {
	tests := []struct {
		name                 string
		sets                 []string
		deployment           *cloud.Deployment
		expectedError        error
		expectedExtraContext *cloud.MetaDataStore
	}{
		{
			name: "Test Values are set successfully",
			sets: []string{"TYK_DB_EMAILBACKEND_ENABLEEMAILNOTIFICATIONS=true", "TYK_DB_EMAILBACKEND_CODE=sendgrid"},
			deployment: &cloud.Deployment{
				Name: "success",
			},
			expectedError:        nil,
			expectedExtraContext: &cloud.MetaDataStore{Data: map[string]map[string]interface{}{"EnvData": {"TYK_DB_EMAILBACKEND_CODE": "sendgrid", "TYK_DB_EMAILBACKEND_ENABLEEMAILNOTIFICATIONS": true}}},
		},
		{
			name: "Test with existing values",
			sets: []string{"DB=tyk", "Email=sendgrid"},
			deployment: &cloud.Deployment{
				ExtraContext: &cloud.MetaDataStore{Data: map[string]map[string]interface{}{"EnvData": {"company": "tyk", "DB": "ty"}}},
			},
			expectedError:        nil,
			expectedExtraContext: &cloud.MetaDataStore{Data: map[string]map[string]interface{}{"EnvData": {"Email": "sendgrid", "company": "tyk", "DB": "tyk"}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := handleEnvVariables(tt.deployment, tt.sets)
			assert.Equal(t, tt.expectedExtraContext, tt.deployment.ExtraContext)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestHandleDeploymentDynamicVars(t *testing.T) {
	cmd := internal.NewCmd("test").
		WithFlagAdder(false, setValues).
		WithFlagAdder(false, envValues).
		WithCommands()
	cmd.SetArgs([]string{
		fmt.Sprintf("--set=%s", "ExtraContext.Data.EnvData.TYK_DB_EMAILBACKEND_CODE=sendgrid"),
		fmt.Sprintf("--envVar=%s", "TYK_DB_EMAILBACKEND_ENABLEEMAILNOTIFICATIONS=true"),
		fmt.Sprintf("--set=%s", "ExtraContext.Data.EnvData.TYK_DB_EMAILBACKEND_DEFAULTFROMEMAIL=no-reply@tyk.io"),
		fmt.Sprintf("--envVar=%s", "TYK_DB_EMAILBACKEND_CODE=sendgri"),
		fmt.Sprintf("--envVar=%s", "DB=tm"),
	})
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	deployment := &cloud.Deployment{
		ExtraContext: &cloud.MetaDataStore{Data: map[string]map[string]interface{}{"EnvData": {"company": "tyk", "DB": "ty"}}},
	}
	err = handleDeploymentDynamicVars(deployment, cmd.Flags())
	assert.Equal(t, nil, err)
	assert.Equal(t, &cloud.MetaDataStore{Data: map[string]map[string]interface{}{"EnvData": {"DB": "tm", "company": "tyk", "TYK_DB_EMAILBACKEND_DEFAULTFROMEMAIL": "no-reply@tyk.io", "TYK_DB_EMAILBACKEND_CODE": "sendgri", "TYK_DB_EMAILBACKEND_ENABLEEMAILNOTIFICATIONS": true}}}, deployment.ExtraContext)
}
