---
id: collect-261001-meraki/meraki/meraki-dashboard-api-go-c8a7f037
title: "meraki-dashboard-api-go-c8a7f037"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "license", "mit license", "parameters"]
source: docs/RAG/collect-261001-meraki/meraki-dashboard-api-go-c8a7f037.md
source_anchor: ""
source_lines: [1, 136]
sha256: df1a6e7b134c9334b1b54fbadc65c0215e75df68864512dddd685571fc5d100a
---

# meraki-dashboard-api-go-c8a7f037

dashboard-api-go is a Go client library for the Meraki Dashboard API.
import meraki "github.com/meraki/dashboard-api-go/sdk"
The dashboard-api-go makes it easier to work with the Meraki Dashboard RESTFul APIs from Go.
It supports version 1.33.0
The first thing you need to do is to generate an API client. There are two options to do it:
- Parameters
- Environment variables
The client could be generated with the following parameters:
- baseURL : The base URL, FQDN or IP, of the MERAKI instance.
- dashboardApiKey : The meraki_key for access to API.
- debug : Boolean to enable debugging
- userAgent : String, set the User-Agent Format (AplicationName VendorName).
client, err = meraki.NewClientWithOptions("https://api.meraki.com/",
		"MerakiKey",
		"true", "AplicationName VendorName Client")
	if err != nil {
		fmt.Println(err)
		return
	}
nResponse, _, err := client.Administered.GetAdministeredIDentitiesMe()
	if err != nil {
		fmt.Println(err)
		return
	}
- Support for fetching all items with perpage=-1
 A new feature has been introduced to the API endpoints, enabling clients to fetch all available items in a single request by setting theperpage parameter to-1 . This enhancement allows you to retrieve the full dataset without needing to make multiple paginated requests.
- When perpage is set to-1 , the server will return all available items for that endpoint, bypassing the pagination logic.
- If a positive integer is passed for perpage , the endpoint will continue using traditional pagination and return only the number of items specified byperpage .
func main() {
	var err error
	fmt.Println("Authenticating")
	client, err = meraki.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	nResponse, _, err := client.Organizations.GetOrganizationDevices("828099381482762270", &meraki.GetOrganizationDevicesQueryParams{
		PerPage:        -1,
		TagsFilterType: "withAnyTags",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if nResponse != nil {
		fmt.Println("\n <Count>: ", len(*nResponse))
		fmt.Printf("%v", *nResponse)
		return
	}
	fmt.Println("There's no data on response")
To retrieve all items from the endpoint, set the perpage parameter to -1:
The client can be configured with the following environment variables:
- MERAKI_BASE_URL : The base URL, FQDN or IP, of the MERAKI instance.
- MERAKI_DASHBOARD_API_KEY : The meraki_key for access to API.
- MERAKI_DEBUG : Boolean to enable debugging
- MERAKI_USER_AGENT : String, set the User-Agent Format (AplicationName VendorName Client).
Client, err = meraki.NewClient()
devicesCount, _, err := Client.Devices.GetDeviceCount()
Here is an example of how we can generate a client, get a device count and then a list of devices filtering them using query params.
client, err = meraki.NewClientWithOptions("https://api.meraki.com/",
		"MerakiKey",
		"true", "AplicationName VendorName Client")
	if err != nil {
		fmt.Println(err)
		return
	}
	nResponse, _, err := client.Organizations.GetOrganizations()
	if err != nil {
		fmt.Println(err)
		return
	}
	if nResponse != nil {
		fmt.Println(nResponse)
		return
	}
	fmt.Println("There's no data on response")
The client allows you to configure automatic retry (backoff) behavior for rate-limited (HTTP 429) responses using the SetBackoff method. This is useful to control how and how many times the SDK retries requests when the API responds with rate limits.
// Example of custom backoff configuration
maxRetries := 5 // Maximum number of retries (total attempts = maxRetries + 1)
maxRetryDelay := 2 * time.Second // Base wait time for exponential backoff between retries
maxRetryJitter := 3 * time.Second // Maximum random jitter to add to the backoff
useRetryHeader := true // Whether to respect the Retry-After header if present
err := client.SetBackoff(&maxRetries, &maxRetryDelay, &maxRetryJitter, &useRetryHeader)
if err != nil {
    fmt.Println("Error configuring backoff:", err)
}
- maxRetries (*int ): Maximum number of retries before failing. Ifnil , the default is 2 (3 total attempts).
- maxRetryDelay (*time.Duration ): Base wait time for exponential backoff. Ifnil , the default is 1 second.
- maxRetryJitter (*time.Duration ): Maximum random jitter to add to the backoff. Ifnil , the default is 3 seconds.
- useRetryHeader (*bool ): Iftrue , the client will respect theRetry-After header in the response to determine the wait time before retrying. Ifnil , the default isfalse .
If you do not call SetBackoff, the following defaults are used:
- maxRetries : 2 (3 total attempts)
- maxRetryDelay : 1 second
- maxRetryJitter : 3 seconds
- useRetryHeader : false
- Backoff is exponential: the wait time doubles on each attempt, plus a random jitter.
- If the response includes a Retry-After header anduseRetryHeader istrue , that value will be used as the wait time.
- The backoff mechanism is automatically applied to all calls that receive an HTTP 429 (rate limit).
package main
import (
    "fmt"
    "time"
    meraki "github.com/meraki/dashboard-api-go/sdk"
)
func main() {
    client, err := meraki.NewClient()
    if err != nil {
        fmt.Println(err)
        return
    }
    maxRetries := 5
    maxRetryDelay := 2 * time.Second
    maxRetryJitter := 3 * time.Second
    useRetryHeader := true
    err = client.SetBackoff(&maxRetries, &maxRetryDelay, &maxRetryJitter, &useRetryHeader)
    if err != nil {
        fmt.Println("Error configuring backoff:", err)
        return
    }
    // All calls will now use this backoff configuration
    nResponse, _, err := client.Organizations.GetOrganizations()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(nResponse)
}
| SDK versions | MERAKI Dashboard version supported | 
|---|---|
| 2.y.z | 1.33.0 | 
| 3.y.z | 1.44.1 | 
| 4.y.z | 1.53.0 | 
| 5.y.z | 1.57.0 | 
All notable changes to this project will be documented in the CHANGELOG file.
The development team may make additional name changes as the library evolves with the MERAKI Dashboard APIs.
This library is distributed under the MIT license found in the LICENSE file.
