/*
 * Ara Platform
 *
 * Ara is a server platform that can deploy Tyk and it's various components into multiple control planes such as K8s and CloudFormation. It is modular and designed to be extended to support multiple specific operational deployments, DNS providers and underlying storage providers
 *
 * API version: v0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"time"
)

// TrafficReport represents the total requests grouped by deployment ID.
type TrafficReport struct {
	Allowed int64 `json:"Allowed,omitempty"`
	// The Base field should not be accessed directly, instead use the helper methods; GetBase & SetBase
	Base          int64     `json:"Base,omitempty"`
	Date          time.Time `json:"Date,omitempty"`
	DepName       string    `json:"DepName,omitempty"`
	DeploymentUID string    `json:"DeploymentUID,omitempty"`
	Diff          int64     `json:"Diff,omitempty"`
	Kind          string    `json:"Kind,omitempty"`
	OrgName       string    `json:"OrgName,omitempty"`
	OrgUID        string    `json:"OrgUID,omitempty"`
	Processed     time.Time `json:"Processed,omitempty"`
	// The Total field should not be accessed directly, instead use the helper methods; GetTotal & SetTotal
	Total     int64 `json:"Total,omitempty"`
	Violation bool  `json:"Violation,omitempty"`
}