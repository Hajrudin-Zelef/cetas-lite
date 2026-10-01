---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/integrations-unifi-5e6e42d0-2
title: "integrations-unifi-5e6e42d0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-unifi-ubiquiti/integrations-unifi-5e6e42d0.md
source_anchor: ""
source_lines: [35, 114]
sha256: 39aa6ca8bd3345f42cc67f6e7d083beda1ada90b8500e5152999d6286aab32f8
---

# integrations-unifi-5e6e42d0

    If the above My button doesn’t work, you can also perform the following steps manually:
- 
Browse to your Home Assistant instance.
- 
Go to Settings > Devices & services.
- 
In the bottom right corner, select the Add Integration button.
- 
From the list, select UniFi Network.
- 
Follow the instructions on screen to complete the setup.
Whether to verify the SSL certificate of the UniFi Network application. Keep this enabled unless you are using a self-signed certificate in a trusted environment and understand the security risk of disabling certificate verification.
Permissions: The below sections on the features available to your Home Assistant instance assume you have full write access to each device. If the user you are using has limited access to some devices, you will get fewer entities and in many cases, get a read-only sensor instead of an editable switch entityAn entity represents a sensor, actor, or function in Home Assistant. Entities are used to monitor physical properties or to control other entities. An entity is usually part of a device or a service. [Learn more].
Configuration options
All configuration options are offered from the front end. Go to Settings > Devices & services, select the UniFi Network integration, and select Configure.
Create device tracker entities for Ubiquiti network devices such as access points and switches.
Only track wireless clients connected to the selected SSIDs. Leave empty to track clients on all SSIDs.
Number of seconds since last seen before a client is considered away. Defaults to 300 seconds.
Disable the workaround for a UniFi Network bug that sometimes reports wired clients as wireless.
Skip Wi-Fi clients that connect with a locally administered MAC address (like private or randomized Wi-Fi addresses), so no entities are created for them. Wired clients are not affected, and clients you select under Create entities from network clients are still included. Disabled by default.
Select clients whose network access you want to control via switches by adding their MAC addresses.
Enable switches to control DPI (Deep Packet Inspection) restriction groups.
Create bandwidth usage sensors for network clients. Disabled by default.
Button
The Button entities will only be available and usable if the integration has a UniFi Network account with administrator privileges.
Power cycle PoE
Use the Power cycle PoE button entity to power cycle one specific PoE port to cause the connected device to restart.
Restart UniFi device
Use the Restart UniFi device button entity to restart the entire UniFi device. In case the device is a PoE switch, the PoE supply is not affected.
WLAN regenerate password
Use the WLAN regenerate password button entity to generate and apply a new password to the specified WLAN (Wireless Local Area Network). It will be randomly generated with 20 characters, consisting of lowercase letters, uppercase letters, and digits.
Image
Provides QR Code images that can be scanned to easily join a specific WLAN. Entities are disabled by default. This feature requires admin privileges.
Presence detection
This platform allows you to detect presence by looking at devices connected to a Ubiquiti UniFi Network application. By default devices are marked as away 300 seconds after they were last seen.
Troubleshooting and Time Synchronization
If tracked devices continue to show “Home” when not connected/present and show connected in the UniFi Controller, disable 802.11r Fast Roaming. When enabled, various UniFi Controller versions have been observed to fail to declare clients disconnected.
Presence detection is not compatible with Client MAC Address Randomization, enabled by default on most modern SmartPhones. This feature will need to be disabled within the client device settings, usually under the settings for the specific network. If you would rather not track these devices at all, turn on Ignore Wi-Fi clients with private (randomized) MAC addresses in the integration options. Home Assistant then skips these clients instead of creating device trackers that never come back.
Presence detection depends on accurate time configuration between Home Assistant and the UniFi Network application.
If Home Assistant and the UniFi Network application are running on separate machines or VMs ensure that all clocks are synchronized. Failing to have synchronized clocks will lead to Home Assistant failing to mark a device as home.
List of actions
The UniFi Network integrationIntegrations connect and integrate Home Assistant with your devices, services, and more. [Learn more] provides the following actions. Each link below opens a dedicated page with examples, parameters, and a step-by-step UI walkthrough.
- 
Reconnect wireless client ( unifi.reconnect_client )
Tries to get a wireless client to reconnect to the UniFi network.
- 
Remove clients from the UniFi Network ( unifi.remove_clients )
Cleans up short-lived clients from the UniFi Network application.
For an overview of every action across all integrations, see the actions reference.
Switch
Block network access for clients
Allow control of network access to clients configured in the integrationIntegrations connect and integrate Home Assistant with your devices, services, and more. [Learn more] options by adding MAC addresses. Items in this list will have a Home Assistant switch created, using the UniFi Device name, allowing for blocking and unblocking.
PoE port control
Provides per-port PoE control. Entities are disabled by default. This feature requires admin privileges.
Port control
Provides individual control to enable or disable switch ports. Entities are disabled by default. This feature requires admin privileges.
Control DPI Traffic Restrictions
Entities appear automatically for each restriction group. If there are no restrictions in a group, no entityAn entity represents a sensor, actor, or function in Home Assistant. Entities are used to monitor physical properties or to control other entities. An entity is usually part of a device or a service. [Learn more] will be visible. Toggling the switch in Home Assistant will enable or disable all restrictions inside a group.
Control WLAN availability
Entities appear for each WLAN. Changing the state of WLAN will trigger a reconfiguration of affected access points, limiting access to all WLANs exposed by the access point.
Control Port Forwarding Rules
Entities appear for each port Forwarding Rule. The switches can be identified from icon 
Control Traffic Rules
Entities appear for each Traffic Rule. The switches can be identified from icon 
Control Policy Engine rules
Entities appear automatically for Policy Engine rules that block internet access. Turning a switch on enables the corresponding rule in the UniFi Network application. Turning it off disables the rule. Policy Engine configurations that only define routing or Quality of Service do not appear as switches.
Control Policy-Based Routing Rules
Entities appear for each Policy-Based Routing Rule. The switches can be identified from icon 
Control Zone-Based Firewall Policies
Entities appear for each Zone-Based Firewall Policy. The switches can be identified from icon 
Sensor
Bandwidth sensor
Get entities reporting receiving and transmitting bandwidth per network client. These sensors are disabled by default. To enable the bandwidth sensors, on the UniFi integration page, select Configure, go to page 3/3 and enable the bandwidth sensors.
Wired client link speed sensor
Get entities reporting the link speed for wired network clients. This sensor shows the connection speed in megabits per second (Mbit/s) between the wired client and the network switch or gateway. These sensors are disabled by default and are only available for wired clients with an active connection.
Wlan clients sensor
Entities reporting connected clients to a WLAN.
Uptime sensor
Get entities reporting uptime per network client or UniFi Network device.
Power Outlet sensor
