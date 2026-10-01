---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-operate-and-maintain-monitoring-and-0a50b257-3
title: "platform-management-dashboard-administration-operate-and-maintain-monitoring-and-0a50b257"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-operate-and-maintain-monitoring-and-0a50b257.md
source_anchor: ""
source_lines: [201, 223]
sha256: 500f93d9b185ca399db85a81531efdb5ef9fcb8c441332ff23c03bc15a9442d1
---

# platform-management-dashboard-administration-operate-and-maintain-monitoring-and-0a50b257

| Device Goes Down/Device Comes Online | A camera goes offline | Camera | 
| Device Goes Down/Device Comes Online | A cellular gateway goes offline | Cellular Gateway | 
Troubleshooting
Additional considerations for Syslog
If the syslog server is across a VPN, the source IP will be 6.X.X.X, provided no VLANs are in VPN mode, the MX is in passthrough mode, or the MX is in Routed/NAT Single LAN mode.
Storage allocation
Syslog messages can take up a large amount of disk space, especially when collecting flows. When choosing a host to run the syslog server, ensure enough storage space to hold the logs. Consult the syslog-ng man page for information on keeping logs for only a certain amount of time.
Expected traffic flow
Syslog traffic flows to the server in one of three scenarios, depending on the route type used to reach the server.
Scenario 1 – Reachable via LAN: The MX sources traffic from the VLAN interface where the server resides if the syslog server is on the LAN of the MX. The transit VLAN interface is used if the device is only accessible via static route.
Scenario 2 – Reachable via public interface: The MX sources traffic from the public interface (WAN) if the syslog server is accessible via the WAN link.
Scenario 3 – Reachable via AutoVPN: The MX sources traffic from the interface of the highest VLAN participating in AutoVPN if the syslog server is accessible via AutoVPN. If the traffic passes through the site-to-site AutoVPN connection, the traffic is subject to the Site-to-site outbound firewall rules, so you may need an allow rule.
To configure this allow rule:
- 
    Navigate to Security & SD-WAN > Configure > Site-to-site VPN > Organization-wide settings.
- 
    Select Add a rule.
Additional Webhooks considerations
Refer to Additional Alerting Considerations for event types that are rate limited.
SNMP trap limitations
- SNMP traps cannot be sent for any Systems Manager alerts.
- If a trap and its associated alert are not listed in Section 5.2, Meraki devices do not send traps for it.
- If devices go down and are not communicating with the dashboard, the trap is sent after the amount of time configured for the alert.
