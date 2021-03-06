/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tsuru

type ServiceBroker struct {
	Name string `json:"Name,omitempty"`

	URL string `json:"URL,omitempty"`

	Config *ServiceBrokerConfig `json:"Config,omitempty"`
}
