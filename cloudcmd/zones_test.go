package cloudcmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TykTechnologies/tykctl/internal"
)

func TestZonesTable(t *testing.T) {
	tests := []struct {
		name     string
		response internal.Payload
		headers  []string
		rows     [][]string
	}{
		{
			name: "Check Success",
			response: internal.Payload{
				Tags: map[string][]string{
					"aws-ap-southeast-1": {"Home", "Gateway"},
					"aws-eu-central-1":   {"Gateway"},
				},
			},
			headers: []string{"Name", "Support Home", "Support Gateway"},
			rows:    [][]string{{"aws-ap-southeast-1", "true", "true"}, {"aws-eu-central-1", "false", "true"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedHeader, expectedRows := ZonesTable(tt.response)
			assert.Equalf(t, tt.headers, expectedHeader, "ZonesTable(%v)", tt.response)
			assert.ElementsMatchf(t, tt.rows, expectedRows, "ZonesTable(%v)", tt.response)
		})
	}
}
