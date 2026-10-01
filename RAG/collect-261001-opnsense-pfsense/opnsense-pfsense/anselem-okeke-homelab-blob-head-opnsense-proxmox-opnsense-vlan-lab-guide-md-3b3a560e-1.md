---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/anselem-okeke-homelab-blob-head-opnsense-proxmox-opnsense-vlan-lab-guide-md-3b3a560e-1
title: "anselem-okeke-homelab-blob-head-opnsense-proxmox-opnsense-vlan-lab-guide-md-3b3a560e"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/collect-261001-opnsense-pfsense/anselem-okeke-homelab-blob-head-opnsense-proxmox-opnsense-vlan-lab-guide-md-3b3a560e.md
source_anchor: ""
source_lines: [1, 493]
sha256: d64cab902a55f779eddf9b2f889c554ebe9faac2967a8839a30af5ac001babe8
---

# anselem-okeke-homelab-blob-head-opnsense-proxmox-opnsense-vlan-lab-guide-md-3b3a560e

This lab builds an enterprise-style segmented network using:

- **Proxmox VE** as the virtualization platform
- **OPNsense** as a virtual firewall/router
- **Windows Server** in a server VLAN
- **Windows 11** in a client VLAN
- **FritzBox/home router** as the upstream internet router

The goal is to understand:

- WAN vs LAN
- VLAN tagging
- DHCP per VLAN
- DNS forwarding/resolution
- Inter-VLAN routing
- Firewall rules
- Why enterprise networks separate servers, clients, guests, and management devices

```
Internet
   |
FritzBox / Home Router
192.168.0.1
   |
Proxmox vmbr0
   |
OPNsense WAN
192.168.0.57
   |
OPNsense Firewall / Router
   |
OPNsense LAN trunk on vmbr1
   |
------------------------------------------------
| VLAN 20 - SERVERS | Gateway 192.168.20.1     |
| VLAN 30 - CLIENTS | Gateway 192.168.30.1     |
| VLAN 10 - MGMT    | Gateway 192.168.10.1     |
| VLAN 40 - GUEST   | Gateway 192.168.40.1     |
------------------------------------------------
   |
Windows Server → VLAN 20 → 192.168.20.x
Windows 11     → VLAN 30 → 192.168.30.x
```
WAN means the outside/upstream side of the firewall.

In this lab:

```
OPNsense WAN = connected to vmbr0 = FritzBox/home network
```
Example:

```
OPNsense WAN IP: 192.168.0.57
FritzBox gateway: 192.168.0.1
```
LAN is the inside side of OPNsense.

In this lab:

```
OPNsense LAN = connected to vmbr1
```
The LAN interface carries multiple VLANs:

```
VLAN 20 = SERVERS
VLAN 30 = CLIENTS
VLAN 10 = MGMT
VLAN 40 = GUEST
```
A VLAN separates traffic logically on the same virtual/physical network.

```
No VLAN tag  = untagged LAN
VLAN tag 20 = SERVERS
VLAN tag 30 = CLIENTS
```
Firewall rules answer:

```
Who can talk to who?
```
In OPNsense, rules are usually placed on the interface where traffic **enters** the firewall.

Example:

```
Windows 11 → Internet
```
The packet enters OPNsense through:

```
CLIENTS interface
```
So the firewall rule goes under:

```
Firewall → Rules → CLIENTS
```
Direction is:

```
in
```
Even though the client is going “out” to the internet, from OPNsense’s view the packet first enters through the CLIENTS interface.

`vmbr0` connects to the physical network/FritzBox.

Used by:

```
OPNsense WAN NIC
Proxmox management
Optional direct home-network VMs
```
`vmbr1` is an internal Proxmox bridge.

Used by:

```
OPNsense LAN NIC
Windows Server VLAN 20
Windows 11 VLAN 30
Other lab VMs
```
`vmbr1` does not need a physical port.

Correct `vmbr1` config:

```
auto vmbr1
iface vmbr1 inet manual
        bridge-ports none
        bridge-stp off
        bridge-fd 0
        bridge-vlan-aware yes
        bridge-vids 2-4094
```
Important:

```
vmbr1 must be VLAN-aware
```
Without VLAN-aware enabled, VLAN-tagged VMs will not receive DHCP from OPNsense.

Create a new VM:

| Setting | Value | 
|---|---|
| VM name | `opnsense-fw` | 
| OS | OPNsense ISO | 
| CPU | 2 cores | 
| RAM | 4 GB recommended | 
| Disk | 20 GB | 
| NIC 1 | WAN | 
| NIC 2 | LAN / VLAN trunk | 

| Proxmox NIC | Bridge | VLAN tag | OPNsense name | Purpose | 
|---|---|---|---|---|
| net0 | vmbr0 | empty | vtnet0 | WAN | 
| net1 | vmbr1 | empty | vtnet1 | LAN trunk | 

Important:

```
Do not put VLAN tags on the OPNsense LAN NIC in Proxmox.
```
OPNsense must receive VLANs internally on `vtnet1`.

Correct:

```
OPNsense net0 → vmbr0 → WAN
OPNsense net1 → vmbr1 → LAN trunk
```
Boot the OPNsense ISO.

When asked to import config:

```
Leave blank and press Enter
```
Login to installer:

```
Username: installer
Password: opnsense
```
Choose:

```
Install (UFS)
```
UFS is simpler and lighter for a small lab VM.

Select the VM disk:

```
ada0 / QEMU HARDDISK / 20GB
```
Do not select the CD/DVD device.

When installation is complete:

1. Set the root password.
2. Select `Complete Install` .
3. Remove/detach the ISO from Proxmox.
4. Reboot/start the OPNsense VM from disk.

When OPNsense asks for interface assignment:

```
Do you want to configure LAGGs now? n
Do you want to configure VLANs now? n
```
Assign:

```
WAN = vtnet0
LAN = vtnet1
Optional = blank
```
Final console should show something like:

```
LAN (vtnet1) -> 192.168.1.1/24
WAN (vtnet0) -> DHCP: 192.168.0.57/24
```
Meaning:

```
WAN = home/FritzBox side
LAN = internal lab side
```
From a VM on the untagged LAN:

```
https://192.168.1.1
```
Login:

```
Username: root
Password: your OPNsense password
```
Browser certificate warning is normal in a lab.

Important:

If your machine is on the FritzBox side, for example:

```
192.168.0.x
```
then:

```
https://192.168.1.1
```
will not work.

That is because `192.168.1.1` is on the OPNsense LAN side.

Go to:

```
Interfaces → Devices → VLAN
```
Create VLANs using parent:

```
vtnet1 [LAN]
```
| Field | Value | 
|---|---|
| Parent | `vtnet1 [LAN]` | 
| VLAN tag | `20` | 
| Description | `SERVERS` | 

| Field | Value | 
|---|---|
| Parent | `vtnet1 [LAN]` | 
| VLAN tag | `30` | 
| Description | `CLIENTS` | 

Optional future VLANs:

| VLAN | Name | Gateway | 
|---|---|---|
| 10 | MGMT | `192.168.10.1/24` | 
| 40 | GUEST | `192.168.40.1/24` | 

Click:

```
Save
Apply
```
Go to:

```
Interfaces → Assignments
```
Add the VLAN interfaces.

They may appear as:

```
vlan01
vlan02
vlan03
```
After adding, OPNsense may name them:

```
OPT1
OPT2
OPT3
```
Rename them.

Go to:

```
Interfaces → OPT1
```
Set:

| Field | Value | 
|---|---|
| Enable Interface | checked | 
| Description | `SERVERS` | 
| IPv4 Configuration Type | `Static IPv4` | 
| IPv4 address | `192.168.20.1/24` | 
| IPv6 Configuration Type | `None` | 
| Block private networks | unchecked | 
| Block bogon networks | unchecked | 

Save and apply.

Go to:

```
Interfaces → OPT2
```
Set:

| Field | Value | 
|---|---|
| Enable Interface | checked | 
| Description | `CLIENTS` | 
| IPv4 Configuration Type | `Static IPv4` | 
| IPv4 address | `192.168.30.1/24` | 
| IPv6 Configuration Type | `None` | 
| Block private networks | unchecked | 
| Block bogon networks | unchecked | 

Save and apply.

OPNsense 26.x may use:

```
Services → Dnsmasq DNS & DHCP
```
Go to:

```
Services → Dnsmasq DNS & DHCP → General
```
Check:

```
Enable
```
For interfaces, select:

```
LAN
SERVERS
CLIENTS
```
Do not select WAN.

Save and apply.

Go to:

```
Services → Dnsmasq DNS & DHCP → DHCP ranges
```
Add:

| Field | Value | 
|---|---|
| Interface | `SERVERS` | 
| Start address | `192.168.20.100` | 
| End address | `192.168.20.200` | 
| Subnet mask | automatic | 
| Lease time | default | 

Save and apply.

Add:

| Field | Value | 
|---|---|
| Interface | `CLIENTS` | 
| Start address | `192.168.30.100` | 
| End address | `192.168.30.200` | 
| Subnet mask | automatic | 
| Lease time | default | 

Save and apply.

In Proxmox:

```
Windows Server VM → Hardware → Network Device → Edit
```
Set:

| Field | Value | 
|---|---|
| Bridge | `vmbr1` | 
| VLAN Tag | `20` | 
| Model | `VirtIO` or`Intel E1000` | 
| Firewall | unchecked | 

Important:

```
The Proxmox firewall checkbox caused DHCP problems in this lab.
Keep it unchecked while learning.
```
Inside Windows Server:

```
ipconfig /release
ipconfig /renew
ipconfig /all
```
Expected:

```
IPv4 Address:    192.168.20.x
Subnet Mask:     255.255.255.0
Default Gateway: 192.168.20.1
DHCP Server:     192.168.20.1
DNS Server:      192.168.20.1
```
In Proxmox:

```
Windows 11 VM → Hardware → Network Device → Edit
```
Set:

| Field | Value | 
|---|---|
| Bridge | `vmbr1` | 
| VLAN Tag | `30` | 
| Model | `Intel E1000` or`VirtIO` if driver works | 
| Firewall | unchecked | 

Inside Windows 11:

```
ipconfig /release
ipconfig /renew
ipconfig /all
```
Expected:

```
IPv4 Address:    192.168.30.x
Subnet Mask:     255.255.255.0
Default Gateway: 192.168.30.1
DHCP Server:     192.168.30.1
DNS Server:      192.168.30.1
```
If Windows shows:

```
169.254.x.x
```
that means:

```
Windows asked for DHCP but no DHCP server replied.
```
Possible causes:

