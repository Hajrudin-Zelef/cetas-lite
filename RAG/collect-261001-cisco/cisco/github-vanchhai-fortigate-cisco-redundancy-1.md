---
id: collect-261001-cisco/cisco/github-vanchhai-fortigate-cisco-redundancy-1
title: "Full Network Redundancy (FortiGate Firewall  +  Cisco Core Switch)"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "preemption", "voice"]
source: docs/RAG/collect-261001-cisco/github-vanchhai-fortigate-cisco-redundancy.md
source_anchor: ""
source_lines: [1, 178]
sha256: 98c2d74b7ba37cd3f8fc715673ef1cfa6c244dec148d44336349dbf734a48fdb
---

# Full Network Redundancy (FortiGate Firewall  +  Cisco Core Switch)

## Design Overview (HA Cluster · StackWise Core · LACP · HSRP · Dual ISP)

Reference architecture for a fully redundant enterprise edge.
Dual FortiGate firewalls in Active-Passive HA, dual Cisco Catalyst core switches in StackWise (or VSS) with LACP uplinks, HSRP gateway redundancy, and dual ISP failover with SD-WAN.


## Recommended Redundant Architecture

**Design Principle**

1. Dual FortiGate firewalls — Active-Passive HA with dedicated heartbeat links
2. Dual Cisco Catalyst 9300/9500 core switches — StackWise-Virtual or VSS as a single logical switch
3. LACP (802.3ad) port-channels between FortiGates and the core stack
4. HSRP / VRRP virtual gateway IPs on core SVIs for first-hop redundancy
5. Rapid PVST+ or MST with consistent root bridge priorities
6. Dual ISP with SD-WAN health checks and policy-based failover
7. Out-of-band management network and redundant power (dual PSU + UPS)

## Network Architecture & Addressing

| Segment | Network Address | Default Gateway | Attached Devices / Role | 
| **WAN 1** | `203.0.113.2/30` | `203.0.113.1` | ISP-A, WAN1 | 
| **WAN 2** | `198.51.100.2/30` | `198.51.100.1` | ISP-B, WAN2 | 
| **VLAN 10** | `10.10.10.0/24` | `10.10.10.1` | USER LAN | 
| **VLAN 20** | `10.10.20.0/24` | `10.10.20.1` | SERVER LAN | 
| **VLAN 30** | `10.10.30.0/24` | `10.10.30.1` | VOICE LAN | 
| **VLAN 99** | `10.255.0.0/24` | `10.255.0.10` | MGMT LAN | 
| **VLAN 999** |  |  | NATIVE-UNUSED | 

## FortiGate HA Cluster Configuration

| Parameter | FGT-A | FGT-B | 
| **Hostname** | `FGT-A` | `FGT-B` | 
| **Group ID** | `10` | `10` | 
| **Group Name** | `CORE-HA` | `CORE-HA` | 
| **Mode** | `a-p` (Active-Passive) | `a-p` (Active-Passive) | 
| **Password** | `Fortinet!HA2026` | `Fortinet!HA2026` | 
| **Heartbeat Devices (hbdev)** | `"ha1" 50` ,`"ha2" 100` | `"ha1" 50` ,`"ha2" 100` | 
| **Session Pickup** | `enable` | `enable` | 
| **Session Pickup Connectionless** | `enable` | `enable` | 
| **HA Management Status** | `enable` | `enable` | 
| **HA Management Interface** | `"mgmt1"` | `"mgmt1"` | 
| **HA Management Gateway** | `10.255.0.1` | `10.255.0.1` | 
| **Override** | `enable` | `disable` | 
| **Priority** | `200`*(Primary candidate)* | `100`*(Secondary candidate)* | 
| **Monitor Interfaces** | `"port1"` ,`"port2"` ,`"wan1"` | `"port1"` ,`"port2"` ,`"wan1"` | 

## FortiGate LACP Aggregate Interface Configuration (`agg1`)

| Command / Parameter | Value | Description | 
| **Interface Name** | `agg1` | The name of the logical aggregate interface. | 
| **vdom** | `root` | Assigns the interface to the "root" Virtual Domain. | 
| **type** | `aggregate` | Defines the interface type as an 802.3ad Link Aggregation Group. | 
| **member** | `"port1" "port2"` | Physical interfaces bundled into this aggregate group. | 
| **lacp-mode** | `active` | Actively transmits LACP packets to negotiate the bond with the peer. | 
| **lacp-speed** | `fast` | Requests LACP packets every 1 second (instead of the standard 30 seconds). | 
| **min-links** | `1` | Minimum number of operational physical links required to keep the aggregate interface up. | 

## FortiGate VLAN Sub-interfaces Configuration on `agg1`

| Interface Name | VDOM | Parent Interface | VLAN ID | IP Address / Subnet | Allowed Administrative Access | 
| **vl10-users** | `root` | `agg1` | `10` | `10.10.10.2/24` | `ping` ,`https` ,`ssh` | 
| **vl20-servers** | `root` | `agg1` | `20` | `10.10.20.2/24` | `ping` | 

## Cisco StackWise Virtual (SVL) Configuration

| Configuration Step | Switch-1 (Primary Candidate) | Switch-2 (Secondary Candidate) | 
| **Domain ID** | `10` | `10` | 
| **SVL Link (SVL1)** | `TenGigabitEthernet1/1/1-2` | `TenGigabitEthernet2/1/1-2` | 
| **Dual-Active Detection (DAD)** | `TenGigabitEthernet1/0/48` | `TenGigabitEthernet2/0/48` | 
| **Command Mode** | `stackwise-virtual` | `stackwise-virtual` | 
| **Final Action** | `write memory` &`reload` | `write memory` &`reload` | 

## Cisco Port-Channel Configuration to FortiGate Cluster

| Logical Interface | Physical Members | Description | Mode | Allowed VLANs | Native VLAN | Security/STP | 
| **Port-channel10** | `Te1/0/1` ,`Te2/0/1` | `>>> FGT-A agg1 <<<` | Trunk | `10, 20, 30, 99` | `999` | `spanning-tree guard root` | 
| **Port-channel20** | `Te1/0/2` ,`Te2/0/2` | `>>> FGT-B agg1 <<<` | Trunk | `10, 20, 30, 99` | `999` | `spanning-tree guard root` | 

### Interface Membership Details

| Physical Interface | Description | Channel Group | LACP Mode | 
| **TenGigabitEthernet1/0/1(CORE-SW1)** | FGT-A port1 | `10` | `active` | 
| **TenGigabitEthernet2/0/1(CORE-SW2)** | FGT-A port2 | `10` | `active` | 
| **TenGigabitEthernet1/0/2(CORE-SW1)** | FGT-B port1 | `20` | `active` | 
| **TenGigabitEthernet2/0/2(CORE-SW2)** | FGT-B port2 | `20` | `active` | 

## Cisco SVI and HSRPv2 Configuration

| Parameter | VLAN 10 (Users) | VLAN 20 (Servers) | 
| **Interface Name** | `Vlan10` | `Vlan20` | 
| **Description** | `USERS GATEWAY` | `SERVERS GATEWAY` | 
| **Physical IP Address** | `10.10.10.252 255.255.255.0` | `10.10.20.252 255.255.255.0` | 
| **HSRP Version** | `Version 2` | `Version 2` | 
| **HSRP Group ID** | `10` | `20` | 
| **Virtual IP (Gateway)** | `10.10.10.1` | `10.10.20.1` | 
| **HSRP Base Priority** | `110` | `110` | 
| **Preemption** | `Enabled` (with 60s minimum delay) | `Enabled` (immediate) | 
| **HSRP Timers** | Hello: `250 msec` / Hold:`750 msec` | *Default (Hello: 3s / Hold: 10s)* | 
| **Authentication** | MD5 Key-string: `CORE-HSRP` | *None* | 
| **Object Tracking** | Track Object `1` (Decrements priority by`20` ) | *None* | 

### Active-Passive cluster with session synchronization

- Mode: Active-Passive (A-P) — single forwarding unit, hot standby
- Group: same HA group ID, group name and password on both units
- Priority: higher value (e.g. 200) becomes primary; preempt enabled
- Heartbeat: two dedicated interfaces (ha1, ha2) — direct fiber/copper
- Session pickup: stateful failover for TCP/UDP and IPsec tunnels
- Monitor interfaces: WAN + LAN tracked; link loss triggers failover
- Override: set 'override enable' on primary to keep it stable
- Sync: configuration, kernel routes, session table, FIB, IPsec SAs
- Failover target: < 1 second for L2/L3, ~3 s for IPsec re-keying

- Configuration (objects, policies)
- Session table (stateful)
- Routing table (FIB / kernel)
- IPsec SAs and FortiGuard cache
- Certificates and local users

## FortiGate CLI — HA Configuration (Apply on BOTH units, sync handles the rest)

```
# Hostname & HA cluster on FGT-A
config system global
    set hostname FGT-A
end
 
config system ha
    set group-id 10
    set group-name CORE-HA
    set mode a-p
    set password Fortinet     #HA2026
    set hbdev "ha1" 50 "ha2" 100
    set session-pickup enable
    set session-pickup-connectionless enable
    set ha-mgmt-status enable
    config ha-mgmt-interfaces
        edit 1
            set interface "mgmt1"
            set gateway 10.255.0.1
        next
    end
    set override enable
    set priority 200
    set monitor "port1" "port2" "wan1"
end
```
```
# Hostname & HA cluster on FGT-B
config system global
    set hostname FGT-B
end
 
config system ha
    set group-id 10
    set group-name CORE-HA
    set mode a-p
    set password Fortinet     #HA2026
    set hbdev "ha1" 50 "ha2" 100
    set session-pickup enable
    set session-pickup-connectionless enable
    set ha-mgmt-status enable
    config ha-mgmt-interfaces
        edit 1
            set interface "mgmt1"
            set gateway 10.255.0.1
        next
    end
    set override disable
    set priority 100
    set monitor "port1" "port2" "wan1"
end
```
### FortiGate — Interfaces, LACP & VLANs (Aggregated uplink to the Cisco core stack)

