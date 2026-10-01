---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/uptimeisbesttime-ha-unifi-network-deb42788-1
title: "Install Ruff (user env or venv)"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/uptimeisbesttime-ha-unifi-network-deb42788.md
source_anchor: ""
source_lines: [1, 107]
sha256: 73eab157f6a02336ada54b7d51bcbe311d122794585c3ad55586eee08f4605a2
---

# Install Ruff (user env or venv)

Home Assistant custom integration for UniFi Network that uses UniFi's official Integration API. It discovers UniFi devices and connected clients, exposes presence detection (device trackers), and provides comprehensive monitoring sensors and control buttons.
- Local polling against the UniFi Network Integration API (no cloud)
- Config flow in the UI (no YAML required)
- Automatic device and client discovery with pagination support
- Selective feature enabling (devices and/or clients)
- Entity categories appropriately marked as Diagnostic or Config
- Service for cleaning up stale client devices from device registry
- Four platforms:
  - device_tracker - Presence detection for clients
  - sensor - Monitoring and diagnostics for devices
  - button - Device and port controls
  - update - Firmware update information
- UniFi Client Tracker: Reports home when client is currently connected,not_home otherwise. Attributes include IP address, MAC address, last seen timestamp, connected at timestamp (when available), andsource_type=router . Automatically created for all discovered clients when client tracking is enabled.
- 
Device State: Current operational state (Online, Offline, etc.)
- 
System Statistics (per device): 
  - Uptime (timestamp showing when device started)
  - Load Average (1, 5, and 15 minute averages)
  - CPU Utilization (%)
  - Memory Utilization (%)
- 
Uplink Statistics (per device): 
  - Uplink RX Rate (bps, suggested display as Mbps)
  - Uplink TX Rate (bps, suggested display as Mbps)
- 
Radio Statistics (per device, per available radio frequency): 
  - TX Retries (%) — created for each available radio frequency (e.g., 2.4GHz, 5GHz, 6GHz)
- 
Port Statistics (per device, per physical port): 
  - Port State (Up, Down, etc.) with additional attributes for port details
- 
PoE Port Statistics (per device, per PoE-capable port): 
  - PoE State (Providing Power, Off, etc.) with PoE standard and type information
- PoE Port Power Cycle (per device, per PoE-capable port): Triggers power cycle action on PoE ports. Button is automatically available only for ports with PoE capability.
- Restart Device (per device): Triggers a restart action on the device. Button is only available when device is online.
- Firmware Update (per device): Shows current firmware version and indicates if firmware updates are available. Displays "Unknown" for latest version when an update is available (UniFi API doesn't provide specific version information). Note: Firmware installation through Home Assistant is not supported - use the UniFi Network Application for updates.
- Remove Stale Clients (unifi_network.remove_stale_clients ): Removes devices from the Home Assistant device registry that are no longer in the known clients or devices list. Useful for cleaning up devices that were previously tracked but are no longer present in the UniFi network. Can target specific config entries or process all UniFi Network integrations.
Update interval: 30 seconds by default.
The integration uses a comprehensive API client that supports additional UniFi Network features that could be implemented in future versions:
- Hotspot voucher management (create, delete, monitor vouchers)
- Guest access controls and authorization
- Advanced device and client actions beyond PoE control
- VPN and Teleport client monitoring
Current implementation focuses on core monitoring and basic device control functionality.
- UniFi Network Application: UniFi OS / Network Application version that supports the Integration API
  - Must have "Integrations" feature available in Network settings
  - API Key generation capability (see Configuration section below for version-specific instructions)
- Network Access: Home Assistant must have network access to your UniFi Network Application
  - Typically runs on port 443 (HTTPS) or 8443
  - Integration API endpoint: /proxy/network/integration
- Supported UniFi Devices: Any UniFi network devices managed by your Network Application
  - Switches, Access Points, Gateways, etc.
  - PoE functionality requires PoE-capable switch ports
Note: HACS installation via the default store is not available yet. You can install manually (see below). If you want to use HACS today, add this repository as a Custom Repository in HACS.
- Ensure you have HACS installed in your Home Assistant instance
- Add this repository as a custom repository in HACS:
  - Open HACS in Home Assistant
  - Click on "Integrations"
  - Click the three dots in the top right corner
  - Select "Custom repositories"
  - Add the repository URL: https://github.com/wittypluck/ha-unifi-network
  - Select category: "Integration"
  - Click "Add"
- Search for "Unifi Network (Local API)" in HACS
- Click "Download"
- Restart Home Assistant
- Go to Settings → Devices & Services → Add Integration and search for "Unifi Network (Local API)"
- Copy the custom_components/unifi_network folder from this repository to your Home Assistant config directory undercustom_components (final path:<config>/custom_components/unifi_network ).
- Restart Home Assistant.
- Go to Settings → Devices & Services → Add Integration and search for "Unifi Network (Local API)".
- 
In Home Assistant, go to Settings → Devices & Services → Add Integration → search for "Unifi Network (Local API)".
- 
Connection Setup: Enter your UniFi Network API credentials: 
  - Base URL: Your UniFi Network Integration API endpoint
    - Format: https://<unifi-host-or-ip>/proxy/network/integration
    - Example: https://192.168.1.1/proxy/network/integration
  - Format: 
  - API Key: Create an API key in the UniFi Network Application:
    - UniFi OS 4.4.x and earlier: UniFi Network → Settings → Control Plane → Integrations → Create API key
    - UniFi OS 5.0.x and later: UniFi Network → Integrations (next to Settings at bottom left) → Create New Api Key
  - Verify SSL Certificate: Enable for production, disable for self-signed certificates
- Base URL: Your UniFi Network Integration API endpoint
- 
Site Selection: Choose which UniFi site to monitor from the automatically discovered list.
- 
Feature Selection: Choose which features to enable: 
  - Unifi Devices sensors and actions: Monitor UniFi network infrastructure devices (switches, access points, gateways, etc.) with comprehensive sensors, buttons, and firmware update information
  - Track Clients: Monitor connected client devices (computers, phones, IoT devices, etc.) with device tracker entities
  - You can enable one or both features based on your monitoring needs
The integration will automatically discover all devices and clients using API pagination to ensure complete coverage. Entities are created dynamically based on device capabilities (e.g., PoE sensors and buttons only appear for ports with PoE support, radio sensors only for devices with wireless radios).
- 
To configure options after initial setup, go to Settings → Devices & Services → Unifi Network → click the gear icon.
- 
Optional Filters: Configure filters to be used when enumerating devices and clients. See API documentation for filter syntax. 
  - Devices Filter: See API documentation for filterable properties.
    - Example: and(not(ipAddress.eq('192.168.1.5')), not(ipAddress.eq('192.168.1.10'))) ignores 192.168.1.5 and 192.168.1.10
  - Example: 
  - Clients Filter: See API documentation for filterable properties.
    - Example: macAddress.eq('00:1a:2b:3c:4d:5e') only tracks 00:1A:2B:3C:4D:5E
    - Note: filter does not appear to match uppercase MAC addresses
  - Example: 
- Devices Filter: See API documentation for filterable properties.
- SSL Certificates: If using self-signed certificates, disable SSL verification in the integration settings or ensure your Home Assistant host trusts the UniFi certificate.
- API Permissions: The API Key should have sufficient privileges for read access to devices, clients, statistics, and port control actions.
- Presence Detection Logic:
