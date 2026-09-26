---
id: collect-260926-mikrotik/mikrotik/mikrotik-capsman-the-complete-guide-to-centralized-wireless-network-management-tech-layer--1
title: "Floor 1 CAPs (identified by identity prefix)"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "license", "parameters"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/mikrotik-capsman-the-complete-guide-to-centralized-wireless-network-management-tech-layer-x-com.md
source_anchor: ""
source_lines: [1, 238]
sha256: d40d87bce376e4729753b9373b19b0533da3d3dcf54c65b306c46043c479da28
---

# Floor 1 CAPs (identified by identity prefix)

Managing multiple wireless access points individually wastes time and creates configuration inconsistencies. MikroTik CAPsMAN (Controlled Access Point System Manager) solves this problem by providing centralized wireless management without additional licensing costs.

## 1. Understanding MikroTik CAPsMAN Architecture

### Core Components of CAPsMAN

CAPsMAN consists of two primary components:

- **CAPsMAN Controller (Manager):** The central router that stores wireless configurations and manages all connected access points. Any RouterOS device with sufficient resources can function as a controller.
- **CAP (Controlled Access Point):** Wireless devices that receive their configuration from the controller. CAPs operate as lightweight access points without local wireless configuration.

### How CAP Discovery Works

CAPs locate their controller using two discovery methods:

| Discovery Method | Protocol | Use Case | Requirements | 
|---|---|---|---|
| Layer 2 (MAC) | Broadcast/Multicast | Same broadcast domain | CAP and Controller on same LAN | 
| Layer 3 (IP) | UDP port 5246/5247 | Across routed networks | IP connectivity to controller | 

### Control and Data Plane Separation

CAPsMAN separates wireless management into two planes:

- **Control Plane:** Handles CAP provisioning, configuration distribution, and management communication. Uses DTLS encryption.
- **Data Plane:** Handles client traffic. Can operate in local forwarding or manager forwarding mode.

### CAPsMAN vs. Standalone AP Management

| Aspect | CAPsMAN | Standalone APs | 
|---|---|---|
| Configuration Location | Centralized on controller | Individual on each AP | 
| New AP Deployment | Automatic provisioning | Manual configuration | 
| Policy Changes | Single point update | Update each AP individually | 
| Client Roaming | Controller-assisted | Client-driven only | 
| Monitoring | Centralized dashboard | Access each AP separately | 
| Failure Impact | Controller failure affects management | Single AP failure isolated | 

## 2. Key Benefits of CAPsMAN for Enterprise Networks

### Simplified Wireless Configuration Management

- Configure all access points from one location
- Push changes to dozens of APs simultaneously
- Maintain consistent settings across the entire network
- Reduce human error from repetitive configuration

### Automatic Provisioning of New Access Points

- New CAPs receive configuration automatically upon connection
- Zero-touch deployment reduces installation time
- Provisioning rules match CAPs to appropriate configurations
- Replacement APs inherit settings without manual intervention

### Centralized Security and Policy Enforcement

- Apply uniform security settings across all APs
- Manage RADIUS integration from controller
- Implement access lists centrally
- Audit security configurations from single interface

### Seamless Client Roaming

- Controller maintains client session information
- Layer 2 roaming between CAPs on same datapath
- Reduced client disconnection during movement
- PMK caching for faster 802.1X reauthentication

### Cost-Effective Scaling

- No licensing fees for controller functionality
- Any RouterOS device can serve as controller
- Add APs without additional software costs
- Scale from 5 to 100+ APs on same platform

## 3. Prerequisites and Planning

### Hardware Requirements

#### CAPsMAN Controller Recommendations

| Deployment Size | Number of CAPs | Recommended Controller | Minimum RAM | 
|---|---|---|---|
| Small | 1-10 | hAP ac², RB750Gr3 | 256 MB | 
| Medium | 11-30 | RB4011, CCR1009 | 512 MB | 
| Large | 31-100+ | CCR1036, CCR2004 | 1 GB+ | 

#### Compatible CAP Devices

- **cAP Series:** cAP ac, cAP XL ac (ceiling mount)
- **wAP Series:** wAP ac, wAP 60G (outdoor)
- **hAP Series:** hAP ac², hAP ac³ (multi-purpose)
- **Audience:** Audience, Audience LTE (tri-band)
- **Any RouterOS device** with wireless interface

### Software Requirements

- **RouterOS Version:** 6.22 or higher (recommend 6.49+ or 7.x)
- **Required Package:** wireless (included in default installation)
- **Version Matching:** Controller and CAP should run same RouterOS version
- **License:** No additional license required

### Network Design Considerations

#### VLAN Planning

- **Management VLAN:** Isolate CAPsMAN control traffic
- **Corporate VLAN:** Internal employee wireless access
- **Guest VLAN:** Isolated internet-only access
- **IoT VLAN:** Segmented device networks

#### IP Addressing Scheme Example

Management Network:  10.0.1.0/24   (Controller: 10.0.1.1, CAPs: DHCP)
Corporate WLAN:      10.0.10.0/24  (VLAN 10)
Guest WLAN:          10.0.20.0/24  (VLAN 20)
IoT WLAN:            10.0.30.0/24  (VLAN 30)

#### Firewall Considerations

Allow these ports for CAPsMAN communication:

- **UDP 5246:** CAPsMAN control (CAPWAP)
- **UDP 5247:** CAPsMAN data (when using manager forwarding)
- **Layer 2:** MAC-based discovery uses broadcast

### Pre-Deployment Checklist

1. Verify RouterOS version compatibility
2. Document VLAN assignments and IP scheme
3. Perform wireless site survey
4. Calculate AP density requirements
5. Plan channel allocation (avoid overlap)
6. Document SSID and security requirements
7. Identify RADIUS server details (if using 802.1X)
8. Establish naming convention for CAPs

## 4. Step-by-Step CAPsMAN Controller Configuration

### Step 1: Enable CAPsMAN Manager

Enable the CAPsMAN controller functionality on your router:

#### CLI Configuration

/caps-man manager
set enabled=yes

#### Winbox Configuration

1. Navigate to CAPsMAN → Manager
2. Check “Enabled”
3. Click OK

### Step 2: Create CAPsMAN Channels

Define radio frequency settings for your wireless networks.

#### 2.4GHz Channel Configuration

```
/caps-man channel
add name=channel-2ghz \
    frequency=2412 \
    band=2ghz-g/n \
    control-channel-width=20mhz \
    extension-channel=disabled \
    tx-power=20
```
#### 5GHz Channel Configuration

```
/caps-man channel
add name=channel-5ghz \
    frequency=5180 \
    band=5ghz-n/ac \
    control-channel-width=20mhz \
    extension-channel=eeCe \
    tx-power=23
```
#### Channel Parameters Explained

| Parameter | Description | Recommended Value | 
|---|---|---|
| frequency | Center frequency in MHz | 2412/2437/2462 (2.4G), 5180/5220/5745 (5G) | 
| band | Wireless standard | 2ghz-g/n or 5ghz-n/ac | 
| control-channel-width | Primary channel width | 20mhz (2.4G), 20mhz (5G for compatibility) | 
| extension-channel | Channel bonding direction | disabled (2.4G), eeCe for 80MHz (5G) | 
| tx-power | Transmit power in dBm | 17-20 (indoor), 23-27 (outdoor) | 

### Step 3: Configure CAPsMAN Datapaths

Datapaths define how client traffic flows through the network.

#### Local Forwarding Datapath (Recommended)

Client traffic bridges locally at the CAP. Controller handles only management.

```
/caps-man datapath
add name=datapath-corporate \
    bridge=bridge-corporate \
    local-forwarding=yes \
    client-to-client-forwarding=yes \
    vlan-mode=use-tag \
    vlan-id=10
add name=datapath-guest \
    bridge=bridge-guest \
    local-forwarding=yes \
    client-to-client-forwarding=no \
    vlan-mode=use-tag \
    vlan-id=20
```
#### Manager Forwarding Datapath

All client traffic tunnels to the controller. Use when central inspection is required.

```
/caps-man datapath
add name=datapath-centralized \
    bridge=bridge-main \
    local-forwarding=no \
    client-to-client-forwarding=yes
```
#### Datapath Parameters Explained

| Parameter | Description | 
|---|---|
| bridge | Bridge interface for client traffic on controller (manager forwarding) or reference for CAP (local forwarding) | 
| local-forwarding | yes = traffic bridges at CAP, no = traffic tunnels to controller | 
| client-to-client-forwarding | Allow wireless clients to communicate directly | 
| vlan-mode | use-tag = apply VLAN tag to traffic | 
| vlan-id | VLAN tag for this datapath | 

### Step 4: Create Security Configurations

#### WPA2-PSK Security Profile

