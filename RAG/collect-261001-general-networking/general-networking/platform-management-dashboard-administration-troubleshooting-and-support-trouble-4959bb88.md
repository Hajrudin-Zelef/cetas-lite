---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-troubleshooting-and-support-trouble-4959bb88
title: "platform-management-dashboard-administration-troubleshooting-and-support-trouble-4959bb88"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2025-08-01"]
keywords: []
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-troubleshooting-and-support-trouble-4959bb88.md
source_anchor: ""
source_lines: [1, 33]
sha256: 300d7e423e447a2a8c0d363c0655308d7d2392bb203b13f15dcfc78902689252
---

# platform-management-dashboard-administration-troubleshooting-and-support-trouble-4959bb88

Cisco Meraki Local Status Page MS Switches
Click 日本語 for Japanese
All new networks created after August 1, 2025 will require a password to be explicitly set for Local Status Page access. Click here to learn more.
MS Switches LSP Overview
This article covers the features, configuration options, and access methods available on the Cisco Meraki Local Status Page (LSP) for MS switches. It outlines how to use the LSP for monitoring connectivity, adjusting network settings, troubleshooting, and managing various configuration options. By default on MS devices running 17+ firmware, the LSP uses the username of "admin" and the password is the serial number/Cloud ID of the device, older firmware uses the username as the serial number without a password.
Any changes on local status page are only intended to be used when necessary to get a switch connected to the dashboard. Before making any changes on the local status page to bring a switch online, ensure the same changes are applied in the dashboard first. Failing to do so will cause the device’s configuration to be overwritten by the dashboard configuration, which may cause the switch to go offline.
Accessing the Local Status Page
Using a static IP address on a device connected to the back panel port (dedicated management port) is not a supported operation. Customers can submit a feature request/feedback using the 'Give your feedback' option in the Cisco Meraki dashboard.
MS Series (includes CS Catalyst Meraki-Managed mode Pre-IOS XE) Local Status Page Options
Every device's status page includes useful information about the status of the device, limited configuration options (such as setting a static IP), and other tools. This section will cover what is available for each device.
MS switches offer the following information and configuration options on their local status page:
- Connection
Provides information regarding the client's connectivity to the switch, the switch's current network, as well as other cloud connectivity and status information.
- Uplink configuration
    
  - Provides options for setting the IP address of the switch, other addressing settings, or configuring a proxy for HTTP traffic.
  - The Download support data function will allow you to download a special file to submit to Meraki support for additional troubleshooting if you are unable to get the unit online (see more in Support Data Bundle (SDB) article).
  - The packet capture option will assist with troubleshooting Meraki Cloud connectivity. Additionally, there is a packet capture tool found here that will assist with troubleshooting Meraki Cloud connectivity on a switch uplink.
Note: The HTTP proxy allows all default management traffic from the Meraki device to be sent through a proxy. This does not include optional cloud communication, including Auto VPN and 802.1x authentication traffic. HTTP proxy is no longer supported on MS 15+ firmware. Nodes that use HTTP proxy without any other means to connect to dashboard may fail to connect. Starting in MS17+, MS devices will now support HTTP CONNECT proxy.
Note: The local status page packet capture requires a minimum firmware version of MS16 and is only supported on a single physical port. At this time, this is not supported for MS390/Catalyst switches.
Additionally, the packet capture function found on the local status page has a default filter that is specific to Meraki Cloud Connectivity requirements and will not capture or display anything outside of that filter. This filter is not configurable. 
This filter is set to capture the following traffic patterns to/from the switch MAC which were determined to be critical to Meraki Cloud connectivity:
- ARP,
- DHCP (UDP 67/68)
- DNS (TCP/UDP 53)
- ICMP (type 0, 3 and 8)
- UDP 7351
- HTTPS (TCP 443)
- LLDP
- Switch port status
Provides information regarding the configuration and status of ports on this switch.
- Switch ports configuration
Provides options for limited configuration changes on switch ports, including enabled/disabled, native VLAN, and link negotiation.
