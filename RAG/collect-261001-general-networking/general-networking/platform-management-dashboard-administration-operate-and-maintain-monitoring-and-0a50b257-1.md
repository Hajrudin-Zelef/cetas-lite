---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-operate-and-maintain-monitoring-and-0a50b257-1
title: "platform-management-dashboard-administration-operate-and-maintain-monitoring-and-0a50b257"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "license", "training"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-operate-and-maintain-monitoring-and-0a50b257.md
source_anchor: ""
source_lines: [1, 94]
sha256: 19aa72bae3cb38045ac1930be58c1586747e3341ee5dda2122c327325051b864
---

# platform-management-dashboard-administration-operate-and-maintain-monitoring-and-0a50b257

How to Configure Meraki Device Reporting via Syslog, SNMP, and API
Click 日本語 for Japanese
Learn more with these free online training courses on the Meraki Learning Hub:
Overview
This article explains how to configure the three device reporting methods available on the Cisco Meraki dashboard: Syslog, API (including Webhooks), and SNMP. Aside from the Meraki Event Log available on the dashboard, these methods let you gather device information and events for reporting and monitoring purposes.
Each method offers distinct benefits for network and device reporting. Use the comparison below to choose the reporting method that best fits your use case.
Choosing a reporting method
The following list compares the capabilities supported by each reporting method.
| Network Event Information | Syslog | API / Webhooks | SNMP | 
|---|---|---|---|
| Device Flows | X |  |  | 
| Client Connectivity |  | X | X | 
| Configuration Changes |  | X | X | 
| Real-time information gathering | X | X | X | 
| Real-time network statistics | X | X | X | 
| Device information gathering | X | X | X | 
| Detailed device statistics |  | X | X | 
| Network-wide information gathering | X | X | X | 
| Organization-wide information gathering |  | X | X | 
| Proactive alerts for critical events |  | X | X | 
| Automation capabilities |  | X | X | 
| Application integration |  | X |  | 
| Cloud-centric |  | X |  | 
| Scalability |  | X |  | 
Step-by-step instructions
Configure a Syslog server
A syslog server stores messages for reporting from MX WAN appliances, MR access points, and MS switches.
Supported message roles by device:
- MX WAN Appliance supports four roles: Event Log, IDS Alerts, URLs, and Flows.
- MR access points support the same roles except IDS Alerts.
- MS switches currently support Event Log messages only.
Add the syslog server
- 
    Navigate to Network-Wide > Configure > General.
- 
    Locate the Reporting section, which contains the Syslog server configuration.
- 
    Select Add a syslog server to define a new server.
- 
    Configure the IP address of your syslog server, the UDP port the server listens on, and the roles you want reported to the server.
In Appliance-only, Switch-only, and Wireless-only dashboard networks, syslog servers appear under the Logging section instead of Reporting.
Encrypted (TLS) syslog will be available for configuration at a future date.
Enable firewall rule logging (optional)
If you enable the Appliance Flows role for Meraki MX reporting, you can enable or disable logging for individual firewall rules.
- 
    Navigate to Security & SD-WAN > Configure > Firewall.
- 
    Set logging under the Logging column.
Configure API reporting
Meraki devices support API calls to gather statistics and other information from your networks. The dashboard API is a powerful, flexible, open-ended tool for many reporting use cases.
Generate a Dashboard API key
The Meraki dashboard API is enabled by default on all organizations. The API key associates with the dashboard administrator account that generates it and inherits that account's permissions. You can generate, revoke, and regenerate your API key on your profile.
- 
    Select the avatar icon in the top right-hand corner of the dashboard, then open the My Profile page.
- 
    Select Generate new API key. A window displays your unique API key.
- 
    Record the API key immediately. Once this window closes, you cannot view the full key on the Meraki dashboard again.
- 
    After copying the key, select the I have stored my new API key checkbox.
- 
    Select Done.
The page refreshes. Under the API access section, only the last 4 digits of the key display for security.
Only you can view your unique API key. No one else with access to the Meraki dashboard—including Meraki Support—can view your API key.
API endpoints for device and network reporting
The following endpoints support device and network reporting:
| API Endpoint | API Call | Description | 
| Client Security Events | Get Network Appliance Client Security Events | Gathers data on all security events for a specified dashboard network | 
| Clients | Get Device Clients | Gathers data on all client devices connected to specific Meraki device, up to a maximum of one month | 
| Devices | Get Network Device Loss and Latency History | Displays data on uplink loss percentage and latency (in milliseconds) for Meraki security appliance | 
| Devices | Get Network Device Performance | Returns the performance score for a specified device. (Only MX security appliances are supported for this API call at this time) | 
| Devices | Get Network Device Uplink | Returns the current uplink information of a specified device. | 
| Devices | Get Network Devices | Lists all devices in the specified network | 
| MV Sense | Get Device Camera Analytics Live | Returns live state form camera of analytics zones configured | 
| MV Sense | Get Device Camera Analytics Overview | Returns overview of aggregate analytics data for timespan specified | 
| MV Sense | Get Device Camera Analytics Recent | Returns the most recent record for analytics zones configured | 
| MV Sense | Get Device Camera Analytics Zone History | Returns historical records for analytic zones configured | 
| MV Sense | Get Device Camera Analytics Zones | Returns all configured analytic zones for specified camera | 
| Networks | Get Network Air Marshal | Lists the Air Marshal scan results for the specified network | 
| Networks | Get Network Traffic | Gathers traffic analysis data for the specified network | 
| Organizations | Get Organization Device Statuses | Lists the current status is every Meraki device in the specified Organization | 
| Organizations | Get Organization License State | Returns the license state for the specified Organization | 
| Organizations | Get Organization Uplinks Loss and Latency | Returns the uplink loss and latency measurements for every MX in the Organization. (From 2 - 7 minutes ago) | 
| Splash Login Attempts | Get Network Splash Login Attempts | Lists the number of login attempts for a splashpage on the specified network | 
| Wireless Health | Get Network Clients Connection Stats | Gathers aggregated connectivity information for wireless clients in the specified network | 
| Wireless Health | Get Network Clients Latency Stats | Gathers aggregated wireless latency for clients on the specified network | 
| Wireless Health | Get Network Connection Stats | Gathers aggregated connectivity data for the specified network | 
| Wireless Health | Get Network Devices Connection Status | Lists client connectivity data on a per-AP basis for the specified network | 
| Wireless Health | Get Network Devices Latency Stats | Lists the amount of wireless latency data on a per-AP basis for the specified network | 
| Wireless Health | Get Network Failed Connections | Lists all failed client connection events on the specified network based on time frame given | 
| Wireless Health | Get Network Latency Stats | Lists aggregated wireless latency information for the entirety of the specified network | 
To learn more about using API with the Meraki dashboard, please visit Cisco Meraki DevNet at https://developer.cisco.com/meraki/
Configure Meraki Dashboard Webhooks
Meraki Webhooks are a lightweight way to subscribe to alerts sent from the Meraki Cloud when an event triggers a configured dashboard alert. They include a JSON-formatted message sent to a unique URL, where you can process, store, or use them to trigger automations. This solution provides a rapid way to set up a hosted service that receives and stores webhook alert data.
