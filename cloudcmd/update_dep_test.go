package cloudcmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TykTechnologies/cloud-sdk/cloud"
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
