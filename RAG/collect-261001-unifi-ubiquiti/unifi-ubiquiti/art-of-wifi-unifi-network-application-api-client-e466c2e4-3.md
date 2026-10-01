---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/art-of-wifi-unifi-network-application-api-client-e466c2e4-3
title: "art-of-wifi-unifi-network-application-api-client-e466c2e4"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "license", "mit license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/art-of-wifi-unifi-network-application-api-client-e466c2e4.md
source_anchor: ""
source_lines: [360, 515]
sha256: cbc17f4c1bd24ba6bd879cceb11f0ba0bebefd1afef5f50c82ce446e627ac424
---

# art-of-wifi-unifi-network-application-api-client-e466c2e4

- deviceId (UUID) -eq ,ne ,in ,notIn ,isNull ,isNotNull
- metadata.origin (STRING) -eq ,ne ,in ,notIn
According to the official API specification, the following properties are filterable for firewall policies:
- id (UUID) -eq ,ne ,in ,notIn
- name (STRING) -eq ,ne ,in ,notIn ,like
- source.zoneId (UUID) -eq ,ne ,in ,notIn
- destination.zoneId (UUID) -eq ,ne ,in ,notIn
- metadata.origin (STRING) -eq ,ne ,in ,notIn
According to the official API specification, the following properties are filterable for DNS policies:
- type (STRING) -eq ,ne ,in ,notIn (Valid values:A ,AAAA ,CNAME ,MX ,TXT ,SRV ,FORWARD_DOMAIN )
- id (UUID) -eq ,ne ,in ,notIn
- enabled (BOOLEAN) -eq ,ne
- domain (STRING) -eq ,ne ,in ,notIn ,like
- ipv4Address (STRING) -eq ,ne ,in ,notIn
- ipv6Address (STRING) -eq ,ne ,in ,notIn
- targetDomain (STRING) -eq ,ne ,in ,notIn ,like
- mailServerDomain (STRING) -eq ,ne ,in ,notIn ,like
- text (STRING) -eq ,ne ,in ,notIn ,like
- serverDomain (STRING) -eq ,ne ,in ,notIn ,like
- ipAddress (STRING) -eq ,ne ,in ,notIn
- ttlSeconds (INTEGER) -eq ,ne ,gt ,ge ,lt ,le
- priority (INTEGER) -eq ,ne ,gt ,ge ,lt ,le
- service (STRING) -eq ,ne ,in ,notIn
- protocol (STRING) -eq ,ne ,in ,notIn
- port (INTEGER) -eq ,ne ,gt ,ge ,lt ,le
- weight (INTEGER) -eq ,ne ,gt ,ge ,lt ,le
For full filtering syntax documentation, see the Network Application API documentation in your controller.
All methods return a Saloon Response object with helpful methods:
<?php
$response = $apiClient->sites()->list();
// Get response as array
$data = $response->json();
// Get response as object
$data = $response->object();
// Check HTTP status
$status = $response->status();
$isSuccessful = $response->successful();
// Access headers
$contentType = $response->header('Content-Type');
// Get raw body
$body = $response->body();
IMPORTANT: The UniFi API returns different response structures depending on the endpoint type:
List Endpoints (e.g., list(), listAdopted(), listVouchers()) return paginated data:
<?php
// List endpoints return data in a 'data' array with pagination metadata
$response = $apiClient->devices()->listAdopted();
$result = $response->json();
// Access the array of items
$devices = $result['data'];  // Array of devices
// Pagination metadata is also available
$total = $result['total'] ?? null;
$offset = $result['offset'] ?? null;
$limit = $result['limit'] ?? null;
// Iterate through items
foreach ($result['data'] as $device) {
    echo $device['name'];
}
Single Item Endpoints (e.g., get($uuid), getVoucher($uuid)) return the object directly:
<?php
// Single item endpoints return the object WITHOUT a 'data' wrapper
$response = $apiClient->devices()->get($deviceId);
$device = $response->json();
// Access properties directly (NO 'data' key!)
echo $device['name'];           // ✓ Correct
echo $device['ipAddress'];      // ✓ Correct
// NOT like this:
// echo $device['data']['name']; // ✗ Wrong! Will cause errors
Key Differences:
- List endpoints: Use $result['data'] to access items + pagination metadata
- Single item endpoints: Access properties directly, no data wrapper, no pagination metadata
Quick Reference:
| Method Type | Response Structure | Example Access | 
|---|---|---|
| list() ,listAdopted() , etc. | { "data": [...], "total": X, ... } | $response->json()['data'][0] | 
| get($id) ,getVoucher($id) , etc. | { "id": "...", "name": "...", ... } | $response->json()['name'] | 
Handle API errors gracefully using try-catch blocks:
<?php
use Saloon\Exceptions\Request\RequestException;
use Saloon\Exceptions\Request\ClientException;
use Saloon\Exceptions\Request\ServerException;
try {
    $response = $apiClient->devices()->get('invalid-uuid-here');
    $device = $response->json();
} catch (ClientException $e) {
    // 4xx errors (client errors like 404 Not Found)
    echo "Client error: " . $e->getMessage();
    echo "Status code: " . $e->getResponse()->status();
} catch (ServerException $e) {
    // 5xx errors (server errors)
    echo "Server error: " . $e->getMessage();
} catch (RequestException $e) {
    // General request errors
    echo "Request failed: " . $e->getMessage();
}
For production environments with valid SSL certificates, enable SSL verification:
<?php
$apiClient = new UnifiClient(
    baseUrl: 'https://unifi.example.com',
    apiKey: 'your-api-key',
    verifySsl: true  // Verify SSL certificates
);
For local development with self-signed certificates, you can disable verification:
<?php
$apiClient = new UnifiClient(
    baseUrl: 'https://192.168.1.1',
    apiKey: 'your-api-key',
    verifySsl: false  // Skip SSL verification (not recommended for production)
);
You can retrieve the client library version at runtime, which is useful for troubleshooting and logging:
<?php
echo $apiClient->getVersion(); // e.g., "1.0.0"
The version is also sent with every API request as a User-Agent header (unifi-api-client-php/1.0.0), which can help
when debugging API issues in controller logs. The version constant is also available directly
via UnifiConnector::VERSION.
The client provides access to the following resources:
| Resource | Description | 
|---|---|
| applicationInfo() | General application information and metadata | 
| sites() | Site management and listing | 
| devices() | Device management, monitoring, actions, adoption, and removal | 
| clients() | Connected client management and guest authorization | 
| networks() | Network configuration (VLANs, DHCP, etc.) | 
| wifiBroadcasts() | WiFi network (SSID) management | 
| hotspot() | Guest voucher management | 
| firewall() | Firewall zone and policy management, including policy ordering | 
| aclRules() | Access Control List (ACL) rule management, including rule ordering | 
| dnsPolicies() | DNS policy management (A, AAAA, CNAME, MX, TXT, SRV, forward domains) | 
| trafficMatchingLists() | Port and IP address lists for firewall policies | 
| supportingResources() | Reference data (WAN interfaces, DPI categories, countries, RADIUS profiles, device tags) | 
See the examples/ directory for complete working examples:
- Basic Usage - Getting started with the client
- Device Management - Working with UniFi devices (including adopt/remove)
- Client Operations - Managing connected clients
- Network Configuration - Creating and managing networks
- WiFi Management - WiFi broadcast configuration
- Error Handling - Proper exception handling
- Firewall Policies & DNS Policies - Firewall policies, ACL ordering, and DNS policies
If you're migrating from the legacy UniFi API client, this new Saloon-based client offers:
- Modern PHP 8.1+ syntax with typed properties
- Fluent interface for more readable code
- Better error handling with exceptions
- Comprehensive IDE support through PHPDoc
- Based on the official API (not the legacy private API)
The main differences:
- Authentication: Uses API keys instead of username/password login
- Return values: Returns Saloon Response objects instead of arrays directly
- Method names: More descriptive and consistent naming
- Site handling: Explicit site ID setting with setSiteId()
Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.
If you encounter any issues or have questions:
- Check the examples directory for working code samples
- Review the official UniFi API documentation within your controller
- Open an issue on GitHub
This library is developed and maintained by Art of WiFi and is developed for the official UniFi Network Application API.
Built with Saloon by Sammyjo20.
This project is licensed under the MIT License - see the LICENSE file for details.
