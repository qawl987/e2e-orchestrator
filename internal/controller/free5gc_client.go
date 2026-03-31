/*
Copyright 2024 The Nephio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"fmt"
)

// registerUEToFree5GCWebConsole is a placeholder function for registering UE
// to the free5GC WebConsole with the specified QFI.
//
// TODO: Implement the actual HTTP client logic for the free5GC WebConsole REST API.
// The implementation should:
// 1. Construct the proper REST API payload for UE registration
// 2. Include the QFI/5QI parameters in the session management configuration
// 3. Handle authentication with the WebConsole
// 4. Parse and handle the response appropriately
//
// Example payload structure (to be implemented):
//
//	{
//	  "plmnID": "20893",
//	  "ueId": "imsi-208930000000001",
//	  "AuthenticationSubscription": {...},
//	  "SessionManagementSubscriptionData": [{
//	    "singleNssai": {"sst": 1, "sd": "010203"},
//	    "dnnConfigurations": {
//	      "internet": {
//	        "pduSessionTypes": {"defaultSessionType": "IPV4"},
//	        "sscModes": {"defaultSscMode": "SSC_MODE_1"},
//	        "5gQosProfile": {
//	          "5qi": <qfi>,
//	          "arp": {"priorityLevel": 8}
//	        }
//	      }
//	    }
//	  }]
//	}
func registerUEToFree5GCWebConsole(ue string, qfi int) error {
	// TODO: Implement actual HTTP client logic for free5GC WebConsole REST API
	// For now, this is a placeholder that logs the intended action.

	fmt.Printf("[free5GC] Registering UE %s with QFI=%d (placeholder - actual implementation pending)\n", ue, qfi)

	// Placeholder implementation - always succeeds
	// In actual implementation:
	// 1. Build HTTP request to free5GC WebConsole API
	// 2. POST to /api/subscriber/imsi-{ue} endpoint
	// 3. Include authentication headers
	// 4. Handle response and errors

	return nil
}

// updateUEQoSProfile updates the QoS profile for an existing UE in free5GC.
//
// TODO: Implement actual HTTP PUT/PATCH logic for updating UE QoS configuration.
func updateUEQoSProfile(ue string, qfi int, sst uint32, sd uint32) error {
	fmt.Printf("[free5GC] Updating UE %s QoS: QFI=%d, SST=%d, SD=%d (placeholder)\n", ue, qfi, sst, sd)
	return nil
}

// Free5GCClient provides methods to interact with the free5GC WebConsole.
// This is a skeleton for future implementation.
type Free5GCClient struct {
	// BaseURL is the free5GC WebConsole API endpoint.
	BaseURL string

	// Username for WebConsole authentication.
	Username string

	// Password for WebConsole authentication.
	Password string
}

// NewFree5GCClient creates a new Free5GC WebConsole client.
func NewFree5GCClient(baseURL, username, password string) *Free5GCClient {
	return &Free5GCClient{
		BaseURL:  baseURL,
		Username: username,
		Password: password,
	}
}

// RegisterSubscriber registers a new subscriber in free5GC.
//
// TODO: Implement actual HTTP client logic.
func (c *Free5GCClient) RegisterSubscriber(imsi string, qfi int, sst uint32, sd uint32) error {
	// Placeholder - to be implemented with actual HTTP calls
	fmt.Printf("[Free5GCClient] RegisterSubscriber: IMSI=%s, QFI=%d, SST=%d, SD=%d\n", imsi, qfi, sst, sd)
	return nil
}

// UpdateSubscriberQoS updates the QoS profile for an existing subscriber.
//
// TODO: Implement actual HTTP client logic.
func (c *Free5GCClient) UpdateSubscriberQoS(imsi string, qfi int, sst uint32, sd uint32) error {
	// Placeholder - to be implemented with actual HTTP calls
	fmt.Printf("[Free5GCClient] UpdateSubscriberQoS: IMSI=%s, QFI=%d, SST=%d, SD=%d\n", imsi, qfi, sst, sd)
	return nil
}

// DeleteSubscriber removes a subscriber from free5GC.
//
// TODO: Implement actual HTTP client logic.
func (c *Free5GCClient) DeleteSubscriber(imsi string) error {
	// Placeholder - to be implemented with actual HTTP calls
	fmt.Printf("[Free5GCClient] DeleteSubscriber: IMSI=%s\n", imsi)
	return nil
}
