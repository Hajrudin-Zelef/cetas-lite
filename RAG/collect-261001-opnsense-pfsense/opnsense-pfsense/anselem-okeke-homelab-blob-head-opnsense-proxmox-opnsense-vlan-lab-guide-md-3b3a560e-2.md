---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/anselem-okeke-homelab-blob-head-opnsense-proxmox-opnsense-vlan-lab-guide-md-3b3a560e-2
title: "anselem-okeke-homelab-blob-head-opnsense-proxmox-opnsense-vlan-lab-guide-md-3b3a560e"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-opnsense-pfsense/anselem-okeke-homelab-blob-head-opnsense-proxmox-opnsense-vlan-lab-guide-md-3b3a560e.md
source_anchor: ""
source_lines: [494, 934]
sha256: 9dca0aa8eb0ce859e5ea8402e1eb9131ccfa98cf2459de79cee63f0344fd146a
---

# anselem-okeke-homelab-blob-head-opnsense-proxmox-opnsense-vlan-lab-guide-md-3b3a560e

- Wrong VLAN tag
- Proxmox bridge not VLAN-aware
- OPNsense DHCP not listening on the VLAN interface
- Missing DHCP range
- Proxmox firewall checkbox enabled
- Wrong NIC is active in Windows
- Multiple adapters are confusing the test

If Windows has multiple adapters, it may have:

```
192.168.0.x  → FritzBox/direct home network
192.168.1.x  → untagged OPNsense LAN
192.168.20.x → SERVERS VLAN
192.168.30.x → CLIENTS VLAN
100.x.x.x    → Tailscale
```
For clean testing, disable extra adapters and keep only the lab adapter active.

PowerShell:

```
Get-NetAdapter
Disable-NetAdapter -Name "Ethernet" -Confirm:$false
```
Use the exact adapter name shown by `Get-NetAdapter`.

New OPNsense interfaces are blocked by default.

That means VLAN interfaces like:

```
SERVERS
CLIENTS
GUEST
MGMT
```
need rules before normal traffic works.

DHCP may work, but ping, browsing, DNS, and Web UI access may fail until rules are added.

This is one of the most important concepts.

If Windows 11 goes to the internet:

```
Windows 11 → OPNsense CLIENTS interface → WAN → Internet
```
The traffic first **enters** OPNsense on the CLIENTS interface.

So the rule is:

```
Interface: CLIENTS
Direction: in
```
Even though the final destination is outside/internet.

Most OPNsense interface rules are created as:

```
Direction: in
```
Start broad first to verify everything works.

Later you can make rules stricter.

Go to:

```
Firewall → Rules → CLIENTS → Add
```
Set:

| Field | Value | 
|---|---|
| Action | `Pass` | 
| Interface | `CLIENTS` | 
| Direction | `in` | 
| TCP/IP Version | `IPv4` | 
| Protocol | `any` | 
| Source | `CLIENTS network` | 
| Destination | `any` | 
| Gateway | `default` | 
| Description | `Allow CLIENTS to any` | 

Save and apply.

Important source choice:

```
CLIENTS network = 192.168.30.0/24
CLIENTS address = only 192.168.30.1
```
Use:

```
CLIENTS network
```
Go to:

```
Firewall → Rules → SERVERS → Add
```
Set:

| Field | Value | 
|---|---|
| Action | `Pass` | 
| Interface | `SERVERS` | 
| Direction | `in` | 
| TCP/IP Version | `IPv4` | 
| Protocol | `any` | 
| Source | `SERVERS network` | 
| Destination | `any` | 
| Gateway | `default` | 
| Description | `Allow SERVERS to any` | 

Save and apply.

```
ipconfig /all
ping 192.168.20.1
ping 8.8.8.8
nslookup google.com
```
Expected:

```
192.168.20.1 ping works
8.8.8.8 ping works
DNS lookup works
```
```
ipconfig /all
ping 192.168.30.1
ping 8.8.8.8
nslookup google.com
```
Expected:

```
192.168.30.1 ping works
8.8.8.8 ping works
DNS lookup works
```
From CLIENTS VLAN:

```
https://192.168.30.1
```
From SERVERS VLAN:

```
https://192.168.20.1
```
From untagged LAN:

```
https://192.168.1.1
```
The firewall rule allows the traffic.

NAT allows private VLAN traffic to go out through WAN.

Default OPNsense usually uses automatic outbound NAT.

Traffic path:

```
Windows 11
192.168.30.x
   |
OPNsense CLIENTS
192.168.30.1
   |
OPNsense WAN
192.168.0.57
   |
FritzBox
192.168.0.1
   |
Internet
```
If:

```
ping 8.8.8.8 works
nslookup google.com fails
```
then routing/NAT works, but DNS is the issue.

If:

```
ping 192.168.30.1 works
ping 8.8.8.8 fails
```
then check:

- CLIENTS firewall rule
- Outbound NAT
- OPNsense WAN gateway
- FritzBox connectivity

Meaning:

```
DHCP failed.
```
Check:

- VM NIC VLAN tag
- VM NIC bridge
- Proxmox firewall checkbox
- `vmbr1` VLAN-aware
- OPNsense VLAN parent
- DHCP range
- Dnsmasq listening interface

Likely cause:

```
vmbr1 is not VLAN-aware
```
Fix:

```
pve01 → System → Network → vmbr1 → VLAN aware checked
```
Or edit:

`nano /etc/network/interfaces`
Ensure:

```
bridge-vlan-aware yes
bridge-vids 2-4094
```
Apply:

`ifreload -a`
If GUI keeps reverting, restart the OPNsense VM or Proxmox networking after confirming config.

Possible causes:

- Change was not applied
- Pending change was reverted
- Wrong bridge was edited
- Proxmox GUI did not persist the setting

Fix manually:

`nano /etc/network/interfaces`
For `vmbr1`:

```
auto vmbr1
iface vmbr1 inet manual
        bridge-ports none
        bridge-stp off
        bridge-fd 0
        bridge-vlan-aware yes
        bridge-vids 2-4094
```
Apply:

`ifreload -a`
Likely cause:

```
OPNsense firewall rule missing
```
New VLAN interfaces are blocked by default.

Add:

```
Firewall → Rules → VLAN_INTERFACE
Action: Pass
Source: VLAN network
Destination: any
```
You may be on the wrong network.

Use the OPNsense gateway IP for your current VLAN:

| Current client IP | OPNsense UI | 
|---|---|
| `192.168.1.x` | `https://192.168.1.1` | 
| `192.168.20.x` | `https://192.168.20.1` | 
| `192.168.30.x` | `https://192.168.30.1` | 
| `192.168.0.x` | WAN side, normally blocked | 

If Windows has:

```
192.168.0.x
Gateway 192.168.0.1
```
then it is using FritzBox directly.

That does not test the OPNsense VLAN lab.

For OPNsense CLIENTS VLAN, Windows should have:

```
192.168.30.x
Gateway 192.168.30.1
```
For OPNsense SERVERS VLAN:

```
192.168.20.x
Gateway 192.168.20.1
```
Release current DHCP lease:

`ipconfig /release`
Request new DHCP lease:

`ipconfig /renew`
Show all network config:

`ipconfig /all`
Flush DNS cache:

`ipconfig /flushdns`
Test gateway:

`ping 192.168.30.1`
Test internet routing:

`ping 8.8.8.8`
Test DNS:

`nslookup google.com`
Test specific port:

`Test-NetConnection 192.168.20.162 -Port 3389`
Show adapters:

`Get-NetAdapter`
Disable adapter:

`Disable-NetAdapter -Name "Ethernet" -Confirm:$false`
Enable adapter:

`Enable-NetAdapter -Name "Ethernet" -Confirm:$false`
Show network config:

`cat /etc/network/interfaces`
Edit network config:

`nano /etc/network/interfaces`
Apply network config without reboot:

`ifreload -a`
Restart a VM only:

```
VM → Shutdown
VM → Start
```
Do not reboot the whole Proxmox host unless necessary.

| NIC | Bridge | VLAN Tag | Purpose | 
|---|---|---|---|
| net0 | vmbr0 | empty | WAN | 
| net1 | vmbr1 | empty | LAN trunk | 

| NIC | Bridge | VLAN Tag | Purpose | 
|---|---|---|---|
| netX | vmbr1 | 20 | SERVERS VLAN | 

| NIC | Bridge | VLAN Tag | Purpose | 
|---|---|---|---|
| netX | vmbr1 | 30 | CLIENTS VLAN | 

For this lab:

```
Unchecked
```
Use OPNsense firewall rules instead.

You can explain the lab like this:

I built a Proxmox-based network lab using OPNsense as a virtual firewall/router. I connected OPNsense with two NICs: one WAN interface to the FritzBox/home network and one LAN trunk interface to an internal VLAN-aware Proxmox bridge. I then created separate VLANs for servers and clients, configured OPNsense interfaces as VLAN gateways, added DHCP scopes per VLAN, and assigned Windows Server to VLAN 20 and Windows 11 to VLAN 30 using Proxmox VLAN tags. After troubleshooting VLAN-aware bridge settings, Proxmox firewall blocking, DHCP timeouts, and Windows adapter routing issues, I implemented OPNsense firewall rules to allow controlled traffic from each VLAN. This gave me a practical understanding of VLAN segmentation, DHCP, DNS, NAT, firewall rule direction, and enterprise-style network isolation.


```
VLAN = logical separation
OPNsense = routing/firewall/DHCP/DNS/NAT
Proxmox = virtual switching and VLAN tagging
vmbr1 must be VLAN-aware
OPNsense LAN NIC is trunk/untagged in Proxmox
VMs receive VLAN tags in Proxmox
New OPNsense VLAN interfaces are blocked by default
Firewall rules are placed where traffic enters the firewall
Direction is usually "in"
DHCP working does not mean firewall traffic is allowed
169.254.x.x means DHCP failed
192.168.0.x means direct FritzBox/home network
192.168.20.x means SERVERS VLAN
192.168.30.x means CLIENTS VLAN
```
After basic rules work, improve security:

Allow:

```
CLIENTS → Internet
CLIENTS → Windows Server only on required ports
```
Block:

```
CLIENTS → MGMT
CLIENTS → other internal networks unless needed
```
Allow:

```
GUEST → Internet
```
Block:

```
GUEST → RFC1918 private networks
10.0.0.0/8
172.16.0.0/12
192.168.0.0/16
```
Allow:

```
MGMT → all infrastructure
```
Restrict:

