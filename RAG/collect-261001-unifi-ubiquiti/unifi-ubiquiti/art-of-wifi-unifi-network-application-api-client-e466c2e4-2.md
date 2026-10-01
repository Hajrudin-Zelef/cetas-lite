---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/art-of-wifi-unifi-network-application-api-client-e466c2e4-2
title: "art-of-wifi-unifi-network-application-api-client-e466c2e4"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/art-of-wifi-unifi-network-application-api-client-e466c2e4.md
source_anchor: ""
source_lines: [167, 359]
sha256: 6923e6e86792ce9c8c76d3ae5aa5d575a62e8ad8718887abb3047610c25d49ee
---

# art-of-wifi-unifi-network-application-api-client-e466c2e4

$apiClient->firewall()->updatePolicyOrdering(
    sourceFirewallZoneId: 'source-zone-uuid',
    destinationFirewallZoneId: 'destination-zone-uuid',
    data: ['orderedFirewallPolicyIds' => ['policy-1-uuid', 'policy-2-uuid']]
);
// List ACL rules
$rules = $apiClient->aclRules()->list();
// Create an ACL rule (requires complex structure - see API docs)
$apiClient->aclRules()->create([
    'type' => 'IPV4',  // or 'MAC'
    'name' => 'Block Social Media',
    'enabled' => true,
    'action' => 'BLOCK',  // or 'ALLOW'
    'index' => 1000,
    // ... additional filters required
]);
// Get/update ACL rule ordering
$ordering = $apiClient->aclRules()->getOrdering();
$apiClient->aclRules()->updateOrdering([
    'orderedAclRuleIds' => ['rule-1-uuid', 'rule-2-uuid', 'rule-3-uuid']
]);<?php
// List all DNS policies
$policies = $apiClient->dnsPolicies()->list();
// Create a DNS A record
$apiClient->dnsPolicies()->create([
    'type' => 'A',
    'enabled' => true,
    'domain' => 'myapp.local',
    'ipv4Address' => '192.168.1.100',
    'ttlSeconds' => 3600,
]);
// Create a DNS CNAME record
$apiClient->dnsPolicies()->create([
    'type' => 'CNAME',
    'enabled' => true,
    'domain' => 'alias.local',
    'targetDomain' => 'myapp.local',
    'ttlSeconds' => 3600,
]);
// Update a DNS policy
$apiClient->dnsPolicies()->update('dns-policy-uuid', [
    'enabled' => false,
    'ipv4Address' => '192.168.1.200',
]);
// Delete a DNS policy
$apiClient->dnsPolicies()->delete('dns-policy-uuid');
All list endpoints use offset-based pagination:
<?php
// Example: List adopted devices with pagination
$response = $apiClient->devices()->listAdopted(
    offset: 100,  // Skip first 100 results
    limit: 50     // Get 50 results
);
$data = $response->json();
// Response structure for list endpoints:
// {
//   "offset": 100,
//   "limit": 50,
//   "count": 50,        // Number of items in current response
//   "totalCount": 1000, // Total items available
//   "data": [...]       // Array of results
// }
The offset parameter specifies how many results to skip, while limit specifies the maximum number of results to return. All paginated endpoints follow this pattern.
The UniFi API supports advanced filtering on many endpoints. You can use either raw filter strings or the type-safe filter builders.
For type-safe, IDE-friendly filtering with autocomplete, use the fluent filter builders:
<?php
use ArtOfWiFi\UnifiNetworkApplicationApi\Filters\Devices\DeviceFilter;
use ArtOfWiFi\UnifiNetworkApplicationApi\Filters\Clients\ClientFilter;
use ArtOfWiFi\UnifiNetworkApplicationApi\Enums\ClientType;
use ArtOfWiFi\UnifiNetworkApplicationApi\Enums\ClientAccessType;
// Simple filtering - find access points
$devices = $apiClient->devices()->listAdopted(
    filter: DeviceFilter::name()->like('AP-*')
);
// Find devices by model
$devices = $apiClient->devices()->listAdopted(
    filter: DeviceFilter::model()->in(['U6-LR', 'U6-PRO', 'U6-ENTERPRISE'])
);
// Complex filtering with AND - wireless guest clients
$clients = $apiClient->clients()->list(
    filter: ClientFilter::and(
        ClientFilter::type()->equals(ClientType::WIRELESS),
        ClientFilter::accessType()->equals(ClientAccessType::GUEST)
    )
);
// Complex filtering with OR - APs or Switches
$devices = $apiClient->devices()->listAdopted(
    filter: DeviceFilter::or(
        DeviceFilter::model()->like('U6*'),
        DeviceFilter::model()->like('USW*')
    )
);
// Multiple conditions - devices needing updates
$devices = $apiClient->devices()->listAdopted(
    filter: DeviceFilter::and(
        DeviceFilter::firmwareUpdatable()->equals(true),
        DeviceFilter::supported()->equals(true)
    )
);
// Null checks - clients with IP addresses
$clients = $apiClient->clients()->list(
    filter: ClientFilter::ipAddress()->isNotNull()
);
// Set operations - devices with WiFi 6
$devices = $apiClient->devices()->listAdopted(
    filter: DeviceFilter::features()->contains('wifi6')
);
// Preset filters for common use cases
$aps = $apiClient->devices()->listAdopted(
    filter: DeviceFilter::accessPoints()
);
$wirelessGuests = $apiClient->clients()->list(
    filter: ClientFilter::wirelessGuests()
);
Countries Filter (Easy to Test!):
The Countries endpoint is perfect for testing filters as it works without needing a site ID and has lots of data:
use ArtOfWiFi\UnifiNetworkApplicationApi\Filters\SupportingResources\CountriesFilter;
// Find United States
$countries = $apiClient->supportingResources()->listCountries(
    filter: CountriesFilter::unitedStates()
);
// Find countries with "Kingdom" in the name
$countries = $apiClient->supportingResources()->listCountries(
    filter: CountriesFilter::name()->like('*Kingdom*')
);
// Find multiple specific countries
$countries = $apiClient->supportingResources()->listCountries(
    filter: CountriesFilter::code()->in(['US', 'GB', 'CA', 'AU'])
);
// Find North American countries
$countries = $apiClient->supportingResources()->listCountries(
    filter: CountriesFilter::northAmerica()
);
Available Filter Classes:
- DeviceFilter - For device filtering
- ClientFilter - For client filtering
- NetworkFilter - For network filtering
- SitesFilter - For site filtering
- FirewallPolicyFilter - For firewall policy filtering
- DnsPolicyFilter - For DNS policy filtering (with presets for record types)
- CountriesFilter - For country filtering (easy to test!)
- DpiCategoriesFilter - For DPI category filtering
- DpiApplicationsFilter - For DPI application filtering
- SiteToSiteVpnTunnelsFilter - For VPN tunnel filtering
- VpnServersFilter - For VPN server filtering
- RadiusProfilesFilter - For RADIUS profile filtering
- DeviceTagsFilter - For device tag filtering
Available Enums:
- ClientType - WIRED, WIRELESS, VPN, TELEPORT
- ClientAccessType - DEFAULT, GUEST
Benefits of Filter Builders:
- Full IDE autocomplete support
- Type safety - catches errors at development time
- Better readability for complex filters
- Property-specific methods for easy discovery
- Preset filters for common use cases
- Automatic value escaping and formatting
If preferred, you can also use raw filter strings from the Network Application API documentation in your controller:
<?php
// Filter devices by name
$response = $apiClient->devices()->listAdopted(
    filter: "name.like('AP-*')"
);
// Filter clients by type and access (wireless guests only)
$response = $apiClient->clients()->list(
    filter: 'and(type.eq("WIRELESS"), access.type.eq("GUEST"))'
);
According to the official API specification, the following properties are filterable for clients:
- id (UUID) -eq ,ne ,in ,notIn
- type (STRING) -eq ,ne ,in ,notIn (Valid values:WIRED ,WIRELESS ,VPN ,TELEPORT )
- macAddress (STRING) -isNull ,isNotNull ,eq ,ne ,in ,notIn
- ipAddress (STRING) -isNull ,isNotNull ,eq ,ne ,in ,notIn
- connectedAt (TIMESTAMP) -isNull ,isNotNull ,eq ,ne ,gt ,ge ,lt ,le
- access.type (STRING) -eq ,ne ,in ,notIn (Valid values:DEFAULT ,GUEST )
- access.authorized (BOOLEAN) -isNull ,isNotNull ,eq ,ne
According to the official API specification, the following properties are filterable for adopted devices:
- id (UUID) -eq ,ne ,in ,notIn
- macAddress (STRING) -eq ,ne ,in ,notIn
- ipAddress (STRING) -eq ,ne ,in ,notIn
- name (STRING) -eq ,ne ,in ,notIn ,like
- model (STRING) -eq ,ne ,in ,notIn
- state (STRING) -eq ,ne ,in ,notIn
- supported (BOOLEAN) -eq ,ne
- firmwareVersion (STRING) -isNull ,isNotNull ,eq ,ne ,gt ,ge ,lt ,le ,like ,in ,notIn
- firmwareUpdatable (BOOLEAN) -eq ,ne
- features (SET(STRING)) -isEmpty ,contains ,containsAny ,containsAll ,containsExactly
- interfaces (SET(STRING)) -isEmpty ,contains ,containsAny ,containsAll ,containsExactly
According to the official API specification, the following properties are filterable for networks:
- management (STRING) -eq ,ne ,in ,notIn
- id (UUID) -eq ,ne ,in ,notIn
- name (STRING) -eq ,ne ,in ,notIn ,like
- enabled (BOOLEAN) -eq ,ne
- vlanId (INTEGER) -eq ,ne ,gt ,ge ,lt ,le ,in ,notIn
