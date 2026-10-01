---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/opnsensemcp-docs-guides-vlan-management-md-at-c2313b49076aa82f49f295079e5d1ee86b381e34-ves
title: "opnsensemcp-docs-guides-vlan-management-md-at-c2313b49076aa82f49f295079e5d1ee86b381e34-vespo92-opnse"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["mcp"]
source: docs/RAG/collect-261001-opnsense-pfsense/opnsensemcp-docs-guides-vlan-management-md-at-c2313b49076aa82f49f295079e5d1ee86b381e34-vespo92-opnse.md
source_anchor: ""
source_lines: [1, 317]
sha256: af3623ef80233358f00d203c08075a6db65bbe1d5f0b731256c9203ace6b72a7
---

# opnsensemcp-docs-guides-vlan-management-md-at-c2313b49076aa82f49f295079e5d1ee86b381e34-vespo92-opnse

Create and manage Virtual LANs (VLANs) to segment your network for security and organization.

VLANs (Virtual Local Area Networks) allow you to:

- Segment network traffic logically
- Improve security through isolation
- Organize devices by function
- Optimize network performance
- Enable multi-tenancy

- Physical interface that supports VLAN tagging (usually igc2 or igc3)
- Switch that supports 802.1Q VLAN tagging
- Understanding of network subnets
- MCP server connected and configured

```
"Create VLAN 50 on interface igc3 for guest network"
```
```
"Set up VLAN 10 on igc2 for IoT devices with description 'Smart Home Devices'"
```
```
"Show me all VLANs"
```
```
"List VLANs on interface igc3"
```
```
"Delete VLAN 50"
```
```
"Remove the IoT VLAN"
```
Isolate guest devices from your main network:

```
"Create a guest network on VLAN 50 with these settings:
- Interface: igc3
- IP range: 192.168.50.0/24
- DHCP: 192.168.50.100-200
- Internet access only
- Block access to LAN"
```
Separate smart home devices:

```
"Set up an IoT network on VLAN 20:
- Interface: igc3
- Subnet: 192.168.20.0/24
- Isolate from main network
- Allow specific services only"
```
Secure network infrastructure:

```
"Create management VLAN 99:
- For switches, APs, and network devices
- Restricted access
- No internet access"
```
Expose services safely:

```
"Create DMZ on VLAN 30:
- For web servers and public services
- Isolated from internal networks
- Internet accessible"
```
Valid VLAN IDs: 1-4094

Common conventions:

- 1: Default/Native VLAN (usually untagged)
- 10-99: Internal networks
- 100-199: Guest networks
- 200-299: IoT/Smart devices
- 900-999: Management VLANs

Choose the right physical interface:

```
"Show network interfaces and their current VLANs"
```
Typical interfaces:

- `igc0` : Usually WAN
- `igc1` : Usually LAN
- `igc2` /`igc3` : Available for VLANs

Plan your subnets:

| VLAN | Purpose | Subnet | DHCP Range | 
|---|---|---|---|
| 10 | IoT | 192.168.10.0/24 | .100-.200 | 
| 20 | Guest | 192.168.20.0/24 | .100-.200 | 
| 30 | DMZ | 192.168.30.0/24 | .100-.150 | 
| 99 | Mgmt | 192.168.99.0/24 | Static only | 

```
"Create VLAN 50 on interface igc3 with description 'Guest Network'"
```
```
"Set IP 192.168.50.1/24 on VLAN 50 interface"
```
```
"Enable DHCP on VLAN 50:
- Range: 192.168.50.100-200
- DNS: 192.168.50.1
- Gateway: 192.168.50.1"
```
```
"Create firewall rules for VLAN 50:
- Allow VLAN 50 to internet
- Block VLAN 50 to LAN
- Block VLAN 50 to other VLANs"
```
```
"Create guest VLAN firewall rules:
1. Allow guest to WAN (internet)
2. Allow guest DNS to firewall
3. Allow guest DHCP to firewall
4. Block guest to LAN net
5. Block guest to firewall (except allowed services)"
```
```
"Set up IoT VLAN firewall:
- Allow IoT to specific cloud services
- Allow LAN to access IoT devices
- Block IoT to LAN initiating connections
- Block IoT to internet except whitelist"
```
Allow specific traffic between VLANs:

```
"Allow printer on VLAN 20 to be accessed from LAN"
```
```
"Allow management VLAN to access all VLANs for administration"
```
```
"Configure DHCP for VLAN 50:
- Range: 192.168.50.100-200
- Lease time: 86400 seconds
- Domain: guest.local"
```
```
"Add static DHCP mapping on VLAN 20:
- MAC: aa:bb:cc:dd:ee:ff
- IP: 192.168.20.50
- Hostname: smart-tv"
```
```
"Set DHCP options for VLAN 30:
- DNS servers: 8.8.8.8, 8.8.4.4
- NTP server: 192.168.30.1
- Domain search: dmz.local"
```
Configure your switch port connected to OPNsense:

```
"The OPNsense port should be configured as:
- Mode: Trunk
- Tagged VLANs: 10,20,30,50,99
- Native VLAN: 1 (or your LAN VLAN)"
```
Configure switch ports for devices:

```
"Device ports should be:
- Mode: Access
- VLAN: (appropriate VLAN ID)
- Example: Guest AP on VLAN 50"
```
**Check physical interface:**

```
"Show status of interface igc3"
```
**Verify VLAN configuration:**

```
"Show configuration for VLAN 50"
```
**Check switch configuration:**

- Ensure port is tagged for VLAN
- Verify VLAN exists on switch
- Check trunk configuration

**Verify DHCP is enabled:**

```
"Show DHCP configuration for VLAN 50"
```
**Check firewall rules:**

```
"Show firewall rules for VLAN 50 interface"
```
Ensure DHCP (port 67/68) is allowed.

**Check NAT rules:**

```
"Show NAT rules for VLAN 50"
```
**Verify firewall allows WAN:**

```
"Check if VLAN 50 can reach WAN"
```
**Check firewall rules:**

```
"Show rules between VLAN 20 and LAN"
```
**Verify routing:**

```
"Show routing table"
```
- Document VLAN purposes
- Plan IP addressing scheme
- Design firewall rules in advance
- Consider growth

Use descriptive names:

- `VLAN10_IoT`
- `VLAN50_Guest`
- `VLAN99_Management`

- Always isolate guest networks
- Restrict IoT device communication
- Use separate management VLAN
- Log inter-VLAN traffic

Keep records of:

- VLAN assignments
- IP ranges
- Firewall rules
- Device assignments

Pass multiple VLANs over single connection:

```
"Configure igc3 as VLAN trunk for VLANs 10,20,30"
```
For service provider deployments:

```
"Enable Q-in-Q on VLAN 100 with S-VLAN 1000"
```
Isolate devices within same VLAN:

```
"Create private VLAN for hotel rooms:
- Primary VLAN: 200
- Isolated VLANs: 201-250"
```
```
"Set up home network VLANs:
- VLAN 1: Main LAN (untagged)
- VLAN 10: IoT devices
- VLAN 20: Guest WiFi
- VLAN 30: Security cameras
- VLAN 99: Network management"
```
```
"Create office VLANs:
- VLAN 10: Workstations
- VLAN 20: VoIP phones
- VLAN 30: Printers
- VLAN 40: Guest WiFi
- VLAN 50: Servers"
```
```
"Set up lab VLANs:
- VLAN 100: Production
- VLAN 200: Development
- VLAN 300: Testing
- VLAN 400: Isolated experiments"
```
```
"Show traffic statistics for VLAN 50"
```
```
"List all devices on VLAN 20"
```
```
"Check status of all VLANs"
```
```
"Apply different DNS servers:
- LAN: Local DNS
- Guest: Public DNS (8.8.8.8)
- IoT: Filtered DNS"
```
```
"Limit guest VLAN to 50Mbps total bandwidth"
```
```
"Enable captive portal on guest VLAN 50"
```
- `createVlan` - Create new VLAN
- `listVlans` - List all VLANs
- `deleteVlan` - Remove VLAN
- `updateVlan` - Modify VLAN settings

- `listInterfaces` - Show available interfaces
- `configureDhcp` - Set up DHCP on VLAN
- `createFirewallRule` - Add firewall rules

- Learn about Firewall Rules for VLAN isolation
- Explore DHCP Management for IP assignment
- Read about DNS Blocking per VLAN
