/*
 * Netsoc webspaced
 *
 * API for managing next-gen webspaces. 
 *
 * API version: 1.1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package webspaced
// InterfaceCounters Counters for a network interface
type InterfaceCounters struct {
	BytesReceived int64 `json:"bytesReceived"`
	BytesSent int64 `json:"bytesSent"`
	PacketsReceived int64 `json:"packetsReceived"`
	PacketsSent int64 `json:"packetsSent"`
}
