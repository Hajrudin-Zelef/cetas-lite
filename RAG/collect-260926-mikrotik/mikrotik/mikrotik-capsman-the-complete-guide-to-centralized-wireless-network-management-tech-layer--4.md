---
id: collect-260926-mikrotik/mikrotik/mikrotik-capsman-the-complete-guide-to-centralized-wireless-network-management-tech-layer--4
title: "Floor 1 CAPs (identified by identity prefix)"
domain: mikrotik
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/mikrotik-capsman-the-complete-guide-to-centralized-wireless-network-management-tech-layer-x-com.md
source_anchor: ""
source_lines: [876, 1106]
sha256: 495f3529914ecf21d3e103c91ea95ca9b38ac3194b9e29518fa44cb64530990b
---

# Floor 1 CAPs (identified by identity prefix)

```
/interface wireless cap
set enabled=yes \
    interfaces=wlan1,wlan2 \
    caps-man-addresses=10.0.1.1,10.0.1.2 \
    bridge=bridge-mgmt
```
#### Controller Failover Behavior

- CAP attempts connection to first address
- If primary fails, CAP connects to secondary
- CAP periodically checks primary availability
- Manual intervention required to sync configurations between controllers

#### Configuration Synchronization

MikroTik does not provide automatic configuration sync between controllers. Options:

1. **Manual Sync:** Export/import configurations periodically
2. **Scripted Sync:** Create script to copy configuration elements
3. **Netinstall:** Restore identical configuration to both

#### Export Configuration for Sync

# On Primary Controller
/caps-man channel export file=capsman-channels
/caps-man datapath export file=capsman-datapaths
/caps-man security export file=capsman-security
/caps-man configuration export file=capsman-configs
/caps-man provisioning export file=capsman-provisioning

### Network Redundancy Considerations

#### VRRP for Gateway Redundancy

# On Primary Router
/interface vrrp
add name=vrrp-mgmt interface=vlan-mgmt vrid=10 priority=200
/ip address
add address=10.0.1.1/24 interface=vrrp-mgmt
# On Secondary Router
/interface vrrp
add name=vrrp-mgmt interface=vlan-mgmt vrid=10 priority=100
/ip address
add address=10.0.1.1/24 interface=vrrp-mgmt

#### Controller Failure Impact Analysis

| Forwarding Mode | Controller Failure Impact | 
|---|---|
| Local Forwarding | Existing clients continue working. No new clients can connect. No roaming. | 
| Manager Forwarding | All traffic stops. Complete network outage for wireless clients. | 

**Recommendation:** Use local forwarding for production networks to minimize controller failure impact.

### Backup and Recovery Procedures

#### Complete Configuration Backup

# Binary backup (includes passwords)
/system backup save name=capsman-controller-full
# Text export (readable, excludes passwords)
/export file=capsman-controller-export

#### Restore Configuration

# Binary restore
/system backup load name=capsman-controller-full
# Text import
/import file=capsman-controller-export.rsc

#### Documentation Requirements

Maintain documentation for:

- Controller IP addresses and credentials
- VLAN assignments and IP schemes
- RADIUS server details
- Certificate information
- Provisioning rule logic
- Channel planning chart
- CAP inventory with locations

## 10. Real-World Deployment Scenarios

### Scenario 1: Small Office (5-10 APs)

#### Network Topology

```
Internet
    |
[ISP Router]
    |
[MikroTik RB4011 - CAPsMAN Controller]
    |
[Managed Switch]
    |
+---+---+---+---+
|   |   |   |   |
CAP CAP CAP CAP CAP
```
#### Key Configuration

- **Controller:** RB4011 (512MB RAM)
- **CAPs:** 5x cAP ac
- **SSIDs:** Corporate (WPA2-PSK), Guest (isolated VLAN)
- **Forwarding:** Local forwarding

#### Configuration Summary

# Complete Small Office Setup
# Channels
/caps-man channel
add name=ch-5ghz band=5ghz-n/ac control-channel-width=20mhz extension-channel=eeCe
add name=ch-2ghz band=2ghz-g/n control-channel-width=20mhz
# Datapaths
/caps-man datapath
add name=dp-corp bridge=bridge local-forwarding=yes vlan-id=10 vlan-mode=use-tag
add name=dp-guest bridge=bridge local-forwarding=yes vlan-id=20 vlan-mode=use-tag client-to-client-forwarding=no
# Security
/caps-man security
add name=sec-corp authentication-types=wpa2-psk encryption=aes-ccm passphrase="CorpSecure2024!"
add name=sec-guest authentication-types=wpa2-psk encryption=aes-ccm passphrase="GuestAccess2024"
# Configurations
/caps-man configuration
add name=cfg-corp-5g ssid="CorpNet" channel=ch-5ghz datapath=dp-corp security=sec-corp country="united states"
add name=cfg-corp-2g ssid="CorpNet" channel=ch-2ghz datapath=dp-corp security=sec-corp country="united states"
add name=cfg-guest-5g ssid="GuestNet" channel=ch-5ghz datapath=dp-guest security=sec-guest country="united states"
# Provisioning
/caps-man provisioning
add master-configuration=cfg-corp-5g slave-configurations=cfg-corp-2g,cfg-guest-5g action=create-dynamic-enabled
# Enable Manager
/caps-man manager set enabled=yes

### Scenario 2: Multi-Floor Enterprise (20-50 APs)

#### Network Topology

```
         [Core Switch]
              |
    +---------+---------+
    |         |         |
[Floor 1] [Floor 2] [Floor 3]
 Switch    Switch    Switch
    |         |         |
  CAPs      CAPs      CAPs
(10 APs)  (15 APs)  (10 APs)
              |
       [CCR1036 - Controller]
```
#### Key Configuration

- **Controller:** CCR1036 (4GB RAM)
- **CAPs:** 35x cAP ac/cAP XL ac
- **SSIDs:** Corporate (802.1X), Guest (Hotspot), IoT
- **VLANs:** Mgmt (99), Corp (10), Guest (20), IoT (30)
- **RADIUS:** Microsoft NPS for 802.1X

#### Channel Planning Strategy

Floor 1: Channels 36, 44, 149, 157 (5GHz)
Floor 2: Channels 40, 48, 153, 161 (5GHz)
Floor 3: Channels 36, 44, 149, 157 (5GHz)
2.4GHz: Channels 1, 6, 11 (rotate per AP)

#### Provisioning by Floor

/caps-man provisioning
add identity-regexp="^F1-.*" master-configuration=cfg-corp-f1-5g slave-configurations=cfg-corp-f1-2g,cfg-guest-5g,cfg-iot action=create-dynamic-enabled
add identity-regexp="^F2-.*" master-configuration=cfg-corp-f2-5g slave-configurations=cfg-corp-f2-2g,cfg-guest-5g,cfg-iot action=create-dynamic-enabled
add identity-regexp="^F3-.*" master-configuration=cfg-corp-f3-5g slave-configurations=cfg-corp-f3-2g,cfg-guest-5g,cfg-iot action=create-dynamic-enabled

### Scenario 3: Multi-Site with Centralized Controller

#### Network Topology

```
[Headquarters]              [Branch 1]              [Branch 2]
     |                          |                       |
[Controller]----[VPN/MPLS]----[Router]----[VPN/MPLS]----[Router]
     |                          |                       |
   CAPs                       CAPs                    CAPs
(20 APs)                    (5 APs)                 (5 APs)
```
#### Key Configuration

- **WAN Consideration:** Use Layer 3 discovery over VPN
- **Forwarding:** Local forwarding mandatory (avoids WAN traffic for data)
- **Bandwidth:** Management traffic minimal (~1 Kbps per CAP)

#### Branch CAP Configuration

```
# Branch CAP settings
/interface wireless cap
set enabled=yes \
    interfaces=wlan1,wlan2 \
    caps-man-addresses=10.100.1.1 \
    bridge=bridge-local
# Ensure local bridge exists for forwarding
/interface bridge
add name=bridge-local
/interface bridge port
add bridge=bridge-local interface=ether1
```
#### Site-Specific Configurations

# Create site-specific datapaths
/caps-man datapath
add name=dp-hq-corp bridge=bridge local-forwarding=yes vlan-id=10
add name=dp-branch1-corp bridge=bridge local-forwarding=yes vlan-id=10
add name=dp-branch2-corp bridge=bridge local-forwarding=yes vlan-id=10
# Provisioning by site identity
/caps-man provisioning
add identity-regexp="^HQ-.*" master-configuration=cfg-hq-5g action=create-dynamic-enabled
add identity-regexp="^BR1-.*" master-configuration=cfg-branch1-5g action=create-dynamic-enabled
add identity-regexp="^BR2-.*" master-configuration=cfg-branch2-5g action=create-dynamic-enabled

### Scenario 4: Hotel/Hospitality Environment

#### Requirements

- High client density per AP
- Guest captive portal with terms acceptance
- Bandwidth limits per client
- Client isolation for security

#### Key Configuration

