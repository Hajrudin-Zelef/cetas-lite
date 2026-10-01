---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-operate-and-maintain-monitoring-and-0a50b257-2
title: "platform-management-dashboard-administration-operate-and-maintain-monitoring-and-0a50b257"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-operate-and-maintain-monitoring-and-0a50b257.md
source_anchor: ""
source_lines: [95, 200]
sha256: d562a53bea10590d6ba2d0a159ad2e883c3495cbf778d986d83e8d7d21ffe3d8
---

# platform-management-dashboard-administration-operate-and-maintain-monitoring-and-0a50b257

Webhooks support all configurable alert types available under Network-wide > Configure > Alerts, including alerts for all product types you own or operate. The webhooks architecture consists of the Meraki cloud and a cloud-accessible HTTP or HTTPS receiver (server). Several standalone cloud webhook services (for example, hook.io and Zapier) can receive or forward webhooks.
Set up Webhooks on the dashboard
- 
    Navigate to Network-Wide > Configure > Alerts.
- 
    Locate the Webhooks section.
- 
    Select Add an HTTP server and configure the server URLs.
- 
    Specify the HTTP servers as recipients for dashboard alerts.
Zapier and Built.io can integrate and automate API applications. Either can be used to set up Meraki dashboard Webhooks.
For more information about setting up and using Webhooks, refer to Cisco Meraki Webhooks.
Configure SNMP
Simple Network Management Protocol (SNMP) lets network administrators query devices for various types of information. Meraki allows SNMP polling to gather information from the dashboard or directly from Meraki devices, including MR access points, MS switches, and MX security appliances. Third-party network monitoring tools can use SNMP to monitor certain parameters on Meraki devices.
SNMP device information
You can poll the following information from the Meraki dashboard:
- Device MAC address
- Device serial number
- Device name
- Device status (online or offline)
- Device last contacted—date and time
- Mesh status (gateway or repeater)
- Public IP address
- Product code (for example, MR18-HW)
- Product description (for example, Meraki Cloud-controller 802.11n AP)
- Name of the network the device resides in (dashboard network)
- Packets/bytes in/out on each physical interface
Standard MIB support
SNMP servers use a Management Information Base (MIB), a database of SNMP Object Identifiers (OIDs). OIDs identify the managed objects that inform the SNMP server which values to poll from the device.
The Meraki dashboard and SNMP-capable Meraki devices support the following standard MIBs:
- SNMPv2-MIB .1.3.6.1.2.1.1
- IF-MIB .1.3.6.1.2.1
Meraki proprietary MIB
The Meraki dashboard has its own proprietary MIB that lets SNMP servers pull data and information from the dashboard.
- 
    Navigate to Organization > Configure > Settings > SNMP to locate and download the MIB document.
The Meraki Cloud MIB cannot poll information from Meraki devices directly, because it is designed strictly for dashboard information.
Configure dashboard SNMP polling
- 
    Navigate to Organization > Configure > Settings > SNMP.
- 
    Choose an SNMP configuration option: SNMP Version 2C or Version 3.
- 
    After you enable SNMP, send SNMP requests to the host defined directly under the enable setting. The community string and a sample command to extract information via SNMP requests display as well.
There are three versions available for configuration for SNMP. These versions are: Version 1, Version 2c, and Version 3
- Versions 1 and 2c allow for a simple community string to be defined. This string will be used between the SNMP server and reporting devices to validate the connection
- Version 3 includes authentication and encryption for added security. Version 3 requires that a username and password be defined
Cisco strongly recommends disabling SNMP V2C and only using SNMP V3 with Authentication mode SHA and Privacy mode AES128 for the highest level of security.
Configure IP restrictions to restrict SNMP access to particular IP addresses in your environment when using v2c. SNMP versions 1 and 2 send the community string in clear text, so IP restrictions prevent unauthorized SNMP access if another party intercepts or learns the community string. IP restrictions are recommended for v3 but not required.
Configure local device SNMP polling
You can poll individual Meraki devices locally. In this scenario, SNMP traffic stays within the local network, and each device is polled from the network management system.
- 
    Navigate to Network-wide > Configure > General > SNMP.
- 
    Configure the local SNMP settings.
Cisco strongly recommends using SNMP V3 with Privacy mode AES128 for the highest level of security.
Configure SNMP traps
SNMP traps are proactive SNMP messages sent when specific networking events take place. They are useful for real-time alerting in your network environment and can be sent from the Meraki cloud. SNMP traps are closely related to the alerts you can configure for your network. SNMP traps use SHA1 for authentication and AES for privacy.
Consider using webhooks instead of SNMP traps if your network environment allows, because webhooks have greater coverage.
Enable SNMP traps
- 
    Navigate to Network-Wide > Configure > Alerts.
- 
    Locate the SNMP Traps section, which is set to Disabled.
- 
    Enable SNMP Version 2c or Version 3.
- 
    To use V3 (username/password), select Add an SNMP user.
- 
    Input the desired username and password, which must also be configured on the SNMP server.
- 
    Enter the Receiving server IP address, which must be a public IP address.
- 
    Configure the receiving server ports with UDP port 161 or 162 (the default ports SNMP servers listen on).
Cisco strongly recommends using SNMP v3 for the highest level of security.
If the SNMP server is behind a NAT device, configure a port forwarding rule to allow the SNMP traffic through. This is due to SNMP traps being sent from the Meraki cloud controller. Specify the correct LAN IP address of the SNMP server and the UDP ports it listens on.
Define SNMP traps to be sent
SNMP traps are closely tied to the alerts configured under Network-Wide > Configure > Alerts. To generate SNMP traps, configure SNMP as a recipient for the alert. You can:
- Input SNMP as a default recipient so all enabled alerts generate a trap, or
- Configure SNMP on a per-alert basis.
Once you enable and configure SNMP traps to send, enable and set up the alerts that trigger them. The following list maps each SNMP trap to its dashboard alert and alert category:
| SNMP Trap Sent | Meraki Dashboard Alert | Alert Category | 
| Settings Changed | Configuration settings are changed | Network-Wide | 
| VPN Connectivity Change | A VPN connection comes up or goes down | Network-Wide | 
| Foreign AP Detected | A rogue access point is detected | Network-Wide | 
| Device Goes Down/Device Comes Online | A gateway goes offline | Wireless | 
| Device Goes Down/Device Comes Online | A repeater goes offline | Wireless | 
| Gateway to Repeater | A gateway becomes a repeater | Wireless | 
| Device Goes Down/Device Comes Online | A WAN appliance goes offline | WAN Appliance | 
| Uplink Status Changed | Primary uplink status changes | WAN Appliance | 
| No DHCP leases | The DHCP lease pool is exhausted | WAN Appliance | 
| IP Conflict | An IP conflict is detected | WAN Appliance | 
| Cellular Network Up/Cellular Network Down | Cellular connection state changes | WAN Appliance | 
| Rogue DHCP Server | A rogue DHCP is detected | WAN Appliance | 
| Warm Spare Failover Detected | A warm spare failover occurs | WAN Appliance | 
| Malware Blocked | Malware is blocked | WAN Appliance | 
| Malware Detected | Malware is downloaded | WAN Appliance | 
| Device Goes Down/Device Comes Online | A switch goes offline | Switch | 
| New DHCP Server Alert | A new DHCP server is detected on the network | Switch | 
| Port Disconnected/Port Connected | Any port goes down | Switch | 
| Port Cable Error | Any port detects a cable error | Switch | 
| Port Speed Change | Any port changes link speed | Switch | 
| Power Supply Down/Power Supply Up | A power supply goes down | Switch | 
| Redundant Power Supply Backup/Redundant Power Supply Back to Primary | A redundant power supply is powering a switch | Switch | 
| UDLD Error | Unidirectional link detection (UDLD) errors exist on a port | Switch | 
| Critical Temperature | A switch is operating at critical temperature | Switch | 
