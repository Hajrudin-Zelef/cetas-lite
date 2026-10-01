---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/integrations-unifi-5e6e42d0-3
title: "integrations-unifi-5e6e42d0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/integrations-unifi-5e6e42d0.md
source_anchor: ""
source_lines: [115, 152]
sha256: 3ee3f2d9792b1c1ad5edba07e7bc0cbc579b97328363142089234545a146ce78
---

# integrations-unifi-5e6e42d0

Get entities reporting the power utilization for outlets that support metrics (such as the AC outlets on the USP-PDU-Pro).
Device temperature sensor
Get entities reporting the general temperature of a UniFi Network device.
Device state
Get entities reporting the current state of a UniFi Network device.
Device CPU
Get entities reporting the current CPU utilization of a UniFi Network device.
Device memory
Get entities reporting the current memory utilization of a UniFi Network device.
Port Bandwidth sensor
Get entities reporting receiving and transmitting bandwidth per port. These sensors are disabled by default. To enable the bandwidth sensors, on the UniFi integration page, select Configure, go to page 3/3 and enable the bandwidth sensors.
Port link speed sensor
Get entities reporting the link negotiation speed for network device ports. These sensors show the connection speed in megabits per second (Mbit/s) at which each port negotiated its link. Entities are disabled by default.
Light
The Light entities will only be available for UniFi access points that support LED ring customization. Not all access points have this capability.
LED control
Provides control over the LED ring on compatible UniFi access points. Entities appear automatically for devices that support LED customization. The LED state, brightness, and color can be controlled. This feature requires admin privileges.
Changes may take over 5 seconds to apply as the device must adopt a new configuration. The UI updates optimistically.
Firmware updates
This will show if there are firmware updates available for the UniFi network devices connected to the controller. If the configured user has admin privileges, the firmware upgrades can also be installed directly from Home Assistant.
Examples
Community blueprints
The Home Assistant community has created blueprints that use the UniFi Network integration for common use cases like presence-based automations or Wi-Fi scheduling. You can browse them in the blueprints exchange on the community forum.
Data updates
The UniFi Network integrationIntegrations connect and integrate Home Assistant with your devices, services, and more. [Learn more] uses a local push connection (WebSocket) to the UniFi Network application. This means state changes for clients, devices, and network configuration are received in near-real-time as they happen on the controller, without the need for pollingData polling is the process of querying a device or service at regular intervals to check for updates or retrieve data. By defining a custom polling interval, you can control how frequently your system checks for new data, which can help optimize performance and reduce unnecessary network traffic. [Learn more].
If the WebSocket connection is lost, the integration automatically tries to reconnect. While disconnected, entities are marked as unavailable until the connection is restored.
Known limitations
- Ubiquiti SSO cloud users are not supported. You must create a local user in your UniFi OS Console. See the Local user section for instructions.
- Early Access and Release Candidate versions of UniFi Network and UniFi OS are not supported. Only the official Stable Release channel is expected to work with this integration.
- Presence detection is not compatible with MAC Address Randomization, which is enabled by default on most modern smartphones. This feature must be disabled per network on the client device.
- Changes to LED control on access points may take over 5 seconds to apply because the device must adopt a new configuration first.
- Lingering entities: In some edge cases, clients or devices removed from UniFi Network may remain in the Home Assistant device registry and need to be removed manually.
Removing a device in Home Assistant
Integration populates both UniFi devices as well as network clients into Home Assistant. In certain edge cases entities are left lingering even if they are not present in UniFi network anymore. This can lead to an accumulation of entries in the device registry.
To manually remove a device entry, go to the Device Info page and select “Delete” from the Device Info menu.
Only clients/devices which are no longer known by UniFi since the startup or reload of the UniFi integration can be removed.
Debugging integration
If you have problems with the UniFi Network application or integrationIntegrations connect and integrate Home Assistant with your devices, services, and more. [Learn more] you can add debug prints to the log.
