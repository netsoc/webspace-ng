/*
 * Netsoc webspaced
 *
 * API for managing next-gen webspaces. 
 *
 * API version: 1.1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package webspaced
// Config Webspace configuration
type Config struct {
	// How many seconds to delay incoming connections to a webspace while starting the container 
	StartupDelay float64 `json:"startupDelay,omitempty"`
	// Incoming SSL-terminated HTTP requests (and SNI passthrough HTTPS connections) will be forwarded to this port 
	HttpPort int32 `json:"httpPort,omitempty"`
	// If true, SSL termination will be disabled and HTTPS connections will forwarded directly 
	SniPassthrough bool `json:"sniPassthrough,omitempty"`
}
