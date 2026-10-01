---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/art-of-wifi-unifi-network-application-api-client-e466c2e4-1
title: "art-of-wifi-unifi-network-application-api-client-e466c2e4"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-unifi-ubiquiti/art-of-wifi-unifi-network-application-api-client-e466c2e4.md
source_anchor: ""
source_lines: [1, 166]
sha256: f158d4279210d4086452fd2d5004548040737843b562102c89c09ab7aba355e3
---

# art-of-wifi-unifi-network-application-api-client-e466c2e4

A modern PHP API client for the official UniFi Network Application API, built on Saloon with a fluent interface for easy integration and powerful features.
This client provides a clean, intuitive way to interact with your UniFi Network Application, supporting all major operations including site management, device control, client monitoring, network configuration, WiFi management, and more.
It is not a direct successor to the UniFi API client which has been developed for the legacy, "unofficial" UniFi API. At this point in time, the "unofficial" API supports more endpoints than the official API. If your integration requirements can be met with the official API, we recommend using this client. For a richer set of features, we currently recommend using the legacy UniFi API client.
- Built on Saloon v3 for robust HTTP communication
- Fluent interface with method chaining for elegant code
- Full support for the official UniFi Network Application API (v10.1.84+)
- Comprehensive coverage of all API endpoints
- Strongly typed using PHP 8.1+ features
- Easy to use for beginners, flexible for advanced users
- Well-documented with inline PHPDoc for IDE auto-completion
- PSR-4 autoloading
- Sends a User-Agent header with every request for easy troubleshooting
- PHP 8.1 or higher
- Composer
- A UniFi OS Server or UniFi OS console with API key access to the Network Application
- Network access to your UniFi Controller
Install via Composer:
composer require art-of-wifi/unifi-network-application-api-client
You must generate an API key from your UniFi Network Application to use this client:
- Log into your UniFi Network Application
- Navigate to Settings → Integrations or straight to Integrations from the sidebar with the latest versions of the UI
- Click Create New API Key
- Give it a descriptive name and save the key securely
- Use this key when initializing the client
- API keys are site-specific and tied to your user account
- The account generating the API key must have appropriate permissions
- API keys can be revoked at any time from the Integrations page
- For local controllers with self-signed certificates, you may need to disable SSL verification (not recommended for production)
Here's the simplest example to get you started:
<?php
require_once 'vendor/autoload.php';
use ArtOfWiFi\UnifiNetworkApplicationApi\UnifiClient;
// Initialize the client
$apiClient = new UnifiClient(
    baseUrl: 'https://192.168.1.1',  // Your controller URL
    apiKey: 'your-api-key-here',      // Your generated API key
    verifySsl: false                   // Set to true for production with valid SSL
);
// Get all sites
$response = $apiClient->sites()->list();
$sites = $response->json();
// Display site names
foreach ($sites['data'] ?? [] as $site) {
    echo "Site: {$site['name']}\n";
}
That's it! You're now connected to your UniFi Network Application.
Most UniFi API operations require a site ID. You can set this once, and it will be used for all subsequent calls:
<?php
use ArtOfWiFi\UnifiNetworkApplicationApi\UnifiClient;
$apiClient = new UnifiClient('https://192.168.1.1', 'your-api-key');
// Set the site ID (get this from the sites list)
$apiClient->setSiteId('550e8400-e29b-41d4-a716-446655440000');
// Now all operations use this site automatically
$devices = $apiClient->devices()->listAdopted();
$clients = $apiClient->clients()->list();
Note: The site ID is a UUID (not the short site name). You can retrieve site IDs using $apiClient->sites()->list().
<?php
// List all adopted devices
$response = $apiClient->devices()->listAdopted();
$devices = $response->json();
// Get a specific device by ID
$device = $apiClient->devices()->get('device-uuid-here');
// Get device statistics
$stats = $apiClient->devices()->getStatistics('device-uuid-here');
// Execute an action on a device (only RESTART is documented)
$apiClient->devices()->executeAction('device-uuid-here', [
    'action' => 'RESTART'
]);
// Adopt a pending device by MAC address
$apiClient->devices()->adopt('00:11:22:33:44:55');
// Adopt a device, ignoring the device limit
$apiClient->devices()->adopt('00:11:22:33:44:55', ignoreDeviceLimit: true);
// Remove (unadopt) a device
$apiClient->devices()->remove('device-uuid-here');<?php
// List all connected clients
$response = $apiClient->clients()->list();
$clients = $response->json();
// Get details for a specific client
$apiClient->clients()->get('client-uuid-here');
// Authorize a guest client
// Requires a lookup for the client's MAC address using $apiClient->clients()->list() with an appropriate filter first.
// Until the Official API supports client device creation, this approach does imply you cannot pre-authorize guests using
// the API because they need to be connected to the network first.
$apiClient->clients()->executeAction('client-uuid-here', [
    'action' => 'AUTHORIZE_GUEST_ACCESS',
    'timeLimitMinutes' => 60  // Grant access for 60 minutes
]);<?php
// List all networks
$networks = $apiClient->networks()->list();
// Create a new UNMANAGED network (simple VLAN)
$apiClient->networks()->create([
    'management' => 'UNMANAGED',
    'name' => 'Guest Network',
    'enabled' => true,
    'vlanId' => 10
]);
// Update a network
$apiClient->networks()->update('network-uuid-here', [
    'name' => 'Updated Guest Network'
]);
// Delete a network
$apiClient->networks()->delete('network-uuid-here');<?php
// List all WiFi broadcasts (SSIDs)
$wifiNetworks = $apiClient->wifiBroadcasts()->list();
// Create a new WiFi network (requires complex nested structure - see examples)
// Refer to examples/05-wifi-management.php for complete structure
// Update WiFi settings
$apiClient->wifiBroadcasts()->update('wifi-uuid', [
    'name' => 'Updated WiFi Name'
]);
// Delete a WiFi network
$apiClient->wifiBroadcasts()->delete('wifi-uuid-here');<?php
// Create vouchers for guest access
$apiClient->hotspot()->createVouchers([
    'count' => 10,
    'timeLimitMinutes' => 480,
    'authorizedGuestLimit' => 1  // How many guests can use same voucher
]);
// List all vouchers
$vouchers = $apiClient->hotspot()->listVouchers();
// Delete a voucher
$apiClient->hotspot()->deleteVoucher('voucher-uuid-here');
Access a variety of resources for the UniFi Network Application configuration:
<?php
// List available WAN interfaces
$wans = $apiClient->supportingResources()->listWanInterfaces();
// List DPI (Deep Packet Inspection) categories
$dpiCategories = $apiClient->supportingResources()->listDpiCategories();
// List DPI applications
$dpiApps = $apiClient->supportingResources()->listDpiApplications();
// List countries (for regulatory compliance)
$countries = $apiClient->supportingResources()->listCountries();
// List RADIUS profiles
$radiusProfiles = $apiClient->supportingResources()->listRadiusProfiles();
// List device tags
$deviceTags = $apiClient->supportingResources()->listDeviceTags();
// List site-to-site VPN tunnels
$vpnTunnels = $apiClient->supportingResources()->listSiteToSiteVpnTunnels();
// List VPN servers
$vpnServers = $apiClient->supportingResources()->listVpnServers();<?php
// List firewall zones
$zones = $apiClient->firewall()->listZones();
// Create a firewall zone
$apiClient->firewall()->createZone([
    'name' => 'DMZ',
    'networkIds' => []  // Array of network UUIDs
]);
// List firewall policies
$policies = $apiClient->firewall()->listPolicies();
// Create a firewall policy
$apiClient->firewall()->createPolicy([
    'name' => 'Block IoT to LAN',
    'enabled' => true,
    'action' => 'BLOCK',
    'source' => ['zoneId' => 'source-zone-uuid'],
    'destination' => ['zoneId' => 'destination-zone-uuid'],
]);
// Partially update a firewall policy (PATCH - only send changed fields)
$apiClient->firewall()->patchPolicy('policy-uuid', [
    'enabled' => false
]);
// Get/update firewall policy ordering between two zones
$ordering = $apiClient->firewall()->getPolicyOrdering(
    sourceFirewallZoneId: 'source-zone-uuid',
    destinationFirewallZoneId: 'destination-zone-uuid'
);
