---
id: collect-261001-fortinet/fortinet/dhruvjaiswal-98-fortinet-handbook-blob-head-handbook-labs-02-dhcp-server-configu-b1ab344e
title: "dhruvjaiswal-98-fortinet-handbook-blob-head-handbook-labs-02-dhcp-server-configu-b1ab344e"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/dhruvjaiswal-98-fortinet-handbook-blob-head-handbook-labs-02-dhcp-server-configu-b1ab344e.md
source_anchor: ""
source_lines: [1, 228]
sha256: bd119bb25761a3305f70e4bde5bfa7e8101eca7a3664bbbfe5f0600c83c59d0e
---

# dhruvjaiswal-98-fortinet-handbook-blob-head-handbook-labs-02-dhcp-server-configu-b1ab344e

This lab demonstrates how to configure **FortiGate as a DHCP Server** and automatically assign IP addresses and network parameters to internal clients.

The objective is to configure the FortiGate internal interface with a static IP address, enable the DHCP Server, define a DHCP address pool, and verify that internal VPC clients receive their IP configuration automatically.

This lab was performed in an **EVE-NG virtual environment** using FortiOS version **7.2.8**.

At the end of this lab you will be able to:

- Configure a FortiGate hostname
- Configure the WAN/Management interface
- Verify interface status
- Configure an internal LAN interface
- Configure FortiGate as a DHCP Server
- Configure a DHCP address range
- Configure the default gateway
- Configure DNS settings
- Configure VPC clients as DHCP clients
- Verify DHCP address assignment
- Verify DHCP leases
- Perform basic DHCP troubleshooting

- FortiGate VM
- FortiOS 7.2.8
- EVE-NG
- FortiGate GUI
- FortiGate CLI
- VPC
- IPv4
- DHCP
- Virtual Switching

- FortiGate GUI Navigation
- FortiGate CLI Configuration
- Interface Configuration
- DHCP Server Configuration
- IP Address Management
- DHCP Verification
- DHCP Troubleshooting
- Basic Network Connectivity Testing

| Device | IP Address | 
|---|---|
| EVE-NG | 10.130.193.182 | 
| Laptop | 10.130.193.79 | 
| Gateway | 10.130.193.157 | 

| Interface | IP Address | 
|---|---|
| Port1 / Management | 10.130.193.81 | 
| Port2 / LAN | 192.168.1.1/24 | 

| Parameter | Value | 
|---|---|
| Network | 192.168.1.0/24 | 
| DHCP Server | FortiGate | 
| DHCP Range | 192.168.1.2 – 192.168.1.254 | 
| Subnet Mask | 255.255.255.0 | 
| Default Gateway | 192.168.1.1 | 
| Client | VPC | 

For this lab, the FortiGate VM was deployed in **EVE-NG** using:

```
FortiOS 7.2.8
```
The firewall was allowed to boot completely before beginning the configuration.

Configure a meaningful hostname for easier device identification.

```
config system global
set hostname FW
end
```
Port1 was configured as the WAN/Management interface.

For this lab, the interface receives its IP address using DHCP.

```
config system interface
edit port1
set mode dhcp
set allowaccess https http ssh ping telnet
next
end
```
Port1 received:

```
10.130.193.81/24
```
Use the following command to verify the physical interface:

`get system interface physical`
The interface should show:

```
Mode   : DHCP
IP     : 10.130.193.81
Status : UP
```
Test connectivity between the FortiGate and the laptop.

`execute ping 10.130.193.79`
A successful ping confirms that the FortiGate can reach the laptop through the management network.

Port2 is configured as the internal LAN interface.

Configure:

| Setting | Value | 
|---|---|
| Interface | Port2 | 
| Address | 192.168.1.1/24 | 
| Role | Internal LAN | 

Port2 will act as the gateway for internal clients.

Enable the **DHCP Server** on Port2.

| Setting | Value | 
|---|---|
| DHCP Status | Enabled | 
| Address Range | 192.168.1.2 – 192.168.1.254 | 
| Netmask | 255.255.255.0 | 
| Default Gateway | Same as Interface IP | 
| DNS Server | Same as System DNS | 
| Lease Time | 604800 seconds | 

FortiGate will now automatically provide IP configuration to clients connected to Port2.

The VPC clients connected to the internal network are configured to obtain their IP address automatically.

From the VPC console:

```
dhcp
```
The client sends a DHCP request to FortiGate.

Example:

```
VPCS> dhcp
DDORA IP 192.168.1.2/24 GW 192.168.1.1
```
The exact IP address may vary depending on the available DHCP lease.

After the DHCP process completes, the VPC should receive:

- IP Address
- Subnet Mask
- Default Gateway
- DNS Information

Example:

```
IP Address      : 192.168.1.2
Subnet Mask     : 255.255.255.0
Default Gateway : 192.168.1.1
```
The client is now successfully receiving its network configuration from FortiGate.

The DHCP configuration can be verified from the FortiGate CLI.

`show system dhcp server`
This allows you to verify the configured DHCP server and address pool.

Use the following command to view active DHCP leases:

`execute dhcp lease-list`
This can be used to verify whether the connected VPC has received a DHCP lease.

If the VPC does not receive an IP address, verify the following:

- Port2 interface status
- Port2 IP address
- DHCP Server configuration
- DHCP address range
- VPC network connection
- DHCP lease table
- Client configuration

`get system interface physical``show system dhcp server``execute dhcp lease-list````
ping 192.168.1.1
```
```
VPC
 |
 | DHCP Request
 v
FortiGate Port2
 |
 | DHCP Server
 |
 +---- Check Interface
 |
 +---- Check DHCP Configuration
 |
 +---- Check DHCP Pool
 |
 +---- Check DHCP Lease
 |
 v
IP Address Assigned
```
| Test | Status | 
|---|---|
| FortiGate VM Boot | ✅ | 
| Hostname Configuration | ✅ | 
| Port1 Configuration | ✅ | 
| Port1 Connectivity | ✅ | 
| Port2 Configuration | ✅ | 
| DHCP Server | ✅ | 
| DHCP Pool | ✅ | 
| VPC DHCP Request | ✅ | 
| IP Address Assignment | ✅ | 
| DHCP Lease Verification | ✅ | 

- FortiGate can operate as a **DHCP Server** .
- Port2 was configured as the internal LAN interface.
- `192.168.1.1/24` was configured as the LAN gateway.
- FortiGate was configured with a DHCP pool from `192.168.1.2` to`192.168.1.254` .
- VPC clients successfully obtained their IP configuration automatically.
- DHCP configuration can be verified using the FortiGate CLI.
- DHCP leases can be checked using `execute dhcp lease-list` .
- Proper interface configuration and DHCP scope configuration are essential for successful IP assignment.
