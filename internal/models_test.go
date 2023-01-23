package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrgInit_OrgInitToMap(t *testing.T) {
	type fields struct {
		Controller string
		Org        string
		Team       string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OrgInit{
				Controller: tt.fields.Controller,
				Org:        tt.fields.Org,
				Team:       tt.fields.Team,
			}
			assert.Equalf(t, tt.want, o.OrgInitToMap(), "OrgInitToMap()")
		})
	}
}
