/*
 * Netsoc webspaced
 *
 * API for managing next-gen webspaces. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package webspaced
// Webspace Netsoc webspace object
type Webspace struct {
	// Unique database identifier, not modifiable.
	User int32 `json:"user,omitempty"`
	Config Config `json:"config,omitempty"`
	// List of webspace custom domains
	Domains []string `json:"domains,omitempty"`
	// Mapping of external ports to internal container ports (port forwarding)
	Ports map[string]int32 `json:"ports,omitempty"`
}
