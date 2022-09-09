/*
 * Ara Platform
 *
 * Ara is a server platform that can deploy Tyk and it's various components into multiple control planes such as K8s and CloudFormation. It is modular and designed to be extended to support multiple specific operational deployments, DNS providers and underlying storage providers
 *
 * API version: v0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse20018 struct {
	Error_  string          `json:"Error,omitempty"`
	Payload []TrafficReport `json:"Payload,omitempty"`
	Status  string          `json:"Status,omitempty"`
}
