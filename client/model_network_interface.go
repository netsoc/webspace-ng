/*
 * Netsoc webspaced
 *
 * API for managing next-gen webspaces. 
 *
 * API version: 1.1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package webspaced
// NetworkInterface Webspace network interface status
type NetworkInterface struct {
	// MAC address
	Mac string `json:"mac"`
	Mtu int32 `json:"mtu"`
	State string `json:"state"`
	Counters InterfaceCounters `json:"counters"`
	Addresses []InterfaceAddress `json:"addresses"`
}
