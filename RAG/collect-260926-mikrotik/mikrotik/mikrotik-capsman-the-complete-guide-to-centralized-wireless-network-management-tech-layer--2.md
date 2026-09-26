---
id: collect-260926-mikrotik/mikrotik/mikrotik-capsman-the-complete-guide-to-centralized-wireless-network-management-tech-layer--2
title: "Floor 1 CAPs (identified by identity prefix)"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/mikrotik-capsman-the-complete-guide-to-centralized-wireless-network-management-tech-layer-x-com.md
source_anchor: ""
source_lines: [239, 557]
sha256: 2231010b0aac44a897ed976ad8d64ee4dca54d03aea9381980e1c41fa2b95944
---

# Floor 1 CAPs (identified by identity prefix)

```
/caps-man security
add name=security-corporate \
    authentication-types=wpa2-psk \
    encryption=aes-ccm \
    passphrase="YourSecurePassword123!" \
    group-key-update=1h
```
#### WPA2-Enterprise Security Profile

```
/caps-man security
add name=security-enterprise \
    authentication-types=wpa2-eap \
    encryption=aes-ccm \
    eap-methods=passthrough \
    eap-radius-accounting=yes
```
#### WPA3 Security Profile (RouterOS 7+)

```
/caps-man security
add name=security-wpa3 \
    authentication-types=wpa3-psk \
    encryption=aes-ccm \
    passphrase="YourSecurePassword123!" \
    group-key-update=1h
```
### Step 5: Create RADIUS Configuration (Optional)

Required for WPA2-Enterprise deployments:

```
/radius
add service=wireless \
    address=10.0.1.100 \
    secret="RadiusSharedSecret" \
    authentication-port=1812 \
    accounting-port=1813 \
    timeout=3s
/caps-man aaa
set interim-update=5m \
    called-format=mac \
    mac-mode=as-username-and-password
```
### Step 6: Build CAPsMAN Configurations

Configurations combine channels, datapaths, and security profiles into complete wireless settings.

#### Corporate 5GHz Configuration

```
/caps-man configuration
add name=cfg-corporate-5ghz \
    ssid="Corporate-WiFi" \
    mode=ap \
    country="united states" \
    channel=channel-5ghz \
    datapath=datapath-corporate \
    security=security-corporate \
    hw-retries=7 \
    disconnect-timeout=5s \
    max-sta-count=50
```
#### Corporate 2.4GHz Configuration

```
/caps-man configuration
add name=cfg-corporate-2ghz \
    ssid="Corporate-WiFi" \
    mode=ap \
    country="united states" \
    channel=channel-2ghz \
    datapath=datapath-corporate \
    security=security-corporate \
    hw-retries=7 \
    max-sta-count=30
```
#### Guest Network Configuration

```
/caps-man configuration
add name=cfg-guest-5ghz \
    ssid="Guest-WiFi" \
    mode=ap \
    country="united states" \
    channel=channel-5ghz \
    datapath=datapath-guest \
    security=security-guest \
    hide-ssid=no \
    max-sta-count=100
```
### Step 7: Create Provisioning Rules

Provisioning rules assign configurations to CAPs automatically.

#### Basic Provisioning (All CAPs Same Config)

```
/caps-man provisioning
add action=create-dynamic-enabled \
    master-configuration=cfg-corporate-5ghz \
    slave-configurations=cfg-corporate-2ghz \
    name-format=identity \
    comment="Default provisioning for all CAPs"
```
#### Location-Based Provisioning

```
# Floor 1 CAPs (identified by identity prefix)
/caps-man provisioning
add action=create-dynamic-enabled \
    identity-regexp="^CAP-F1-.*" \
    master-configuration=cfg-corporate-5ghz-f1 \
    slave-configurations=cfg-corporate-2ghz-f1 \
    name-format=identity \
    comment="Floor 1 CAPs"
# Floor 2 CAPs
/caps-man provisioning
add action=create-dynamic-enabled \
    identity-regexp="^CAP-F2-.*" \
    master-configuration=cfg-corporate-5ghz-f2 \
    slave-configurations=cfg-corporate-2ghz-f2 \
    name-format=identity \
    comment="Floor 2 CAPs"
```
#### Radio-Specific Provisioning

```
# 5GHz radios only
/caps-man provisioning
add action=create-dynamic-enabled \
    hw-supported-modes=ac \
    master-configuration=cfg-corporate-5ghz \
    name-format=identity \
    comment="5GHz capable radios"
# 2.4GHz radios only
/caps-man provisioning
add action=create-dynamic-enabled \
    hw-supported-modes=gn \
    master-configuration=cfg-corporate-2ghz \
    name-format=identity \
    comment="2.4GHz radios"
```
#### Provisioning Parameters Reference

| Parameter | Description | 
|---|---|
| action | create-dynamic-enabled (auto-enable), create-enabled (static), none | 
| master-configuration | Primary radio configuration (typically 5GHz) | 
| slave-configurations | Secondary radio configuration (typically 2.4GHz) | 
| identity-regexp | Match CAP by system identity | 
| radio-mac | Match CAP by specific MAC address | 
| hw-supported-modes | Match by radio capabilities (ac, gn, an) | 
| name-format | How to name created interfaces (identity, cap, prefix) | 

## 5. CAP Device Setup and Deployment

### Preparing CAP Devices

#### Step 1: Reset CAP to Default

/system reset-configuration no-defaults=yes skip-backup=yes

#### Step 2: Set System Identity

/system identity set name="CAP-F1-01"

#### Step 3: Configure Basic IP Connectivity

# Create bridge for management
/interface bridge
add name=bridge-mgmt
/interface bridge port
add bridge=bridge-mgmt interface=ether1
# Get IP via DHCP
/ip dhcp-client
add interface=bridge-mgmt disabled=no

#### Step 4: Upgrade Firmware

/system package update
set channel=stable
check-for-updates
download
# Reboot after download
/system reboot

### Configure CAP Mode

#### Layer 2 Discovery (Same Subnet)

```
/interface wireless cap
set enabled=yes \
    interfaces=wlan1,wlan2 \
    discovery-interfaces=bridge-mgmt \
    bridge=bridge-mgmt
```
#### Layer 3 Discovery (Routed Network)

```
/interface wireless cap
set enabled=yes \
    interfaces=wlan1,wlan2 \
    caps-man-addresses=10.0.1.1 \
    bridge=bridge-mgmt
```
#### CAP Configuration Parameters

| Parameter | Description | 
|---|---|
| enabled | Enable CAP functionality | 
| interfaces | Wireless interfaces to be controlled by CAPsMAN | 
| discovery-interfaces | Interfaces for Layer 2 discovery broadcasts | 
| caps-man-addresses | Controller IP addresses for Layer 3 discovery | 
| bridge | Bridge for local forwarding traffic | 
| lock-to-caps-man | Bind CAP to specific controller | 

### Certificate-Based CAP Authentication

#### Generate CA Certificate on Controller

/certificate
add name=capsman-ca common-name=capsman-ca key-usage=key-cert-sign,crl-sign
sign capsman-ca

#### Generate Controller Certificate

/certificate
add name=capsman-controller common-name=capsman-controller
sign capsman-controller ca=capsman-ca

#### Generate CAP Certificate

/certificate
add name=cap-01 common-name=cap-01
sign cap-01 ca=capsman-ca

#### Enable Certificate Requirement on Controller

```
/caps-man manager
set require-peer-certificate=yes \
    certificate=capsman-controller \
    ca-certificate=capsman-ca
```
#### Configure CAP with Certificate

# Export CA certificate from controller
/certificate export-certificate capsman-ca
# Import on CAP (after copying file)
/certificate import file-name=capsman-ca.crt
# Configure CAP with certificate
/interface wireless cap
set certificate=cap-01

### Zero-Touch Provisioning with DHCP

Configure DHCP server to provide CAPsMAN controller address:

# On DHCP Server
/ip dhcp-server option
add name=capsman-option code=138 value="'10.0.1.1'"
/ip dhcp-server network
set [find] dhcp-option=capsman-option

CAPs will automatically discover the controller using DHCP option 138.

## 6. Advanced CAPsMAN Features and Optimization

### Access List Configuration

Control which clients can connect based on MAC address or signal strength.

#### MAC-Based Access Control

```
# Allow specific devices
/caps-man access-list
add mac-address=AA:BB:CC:DD:EE:FF \
    action=accept \
    comment="Executive laptop"
# Block specific devices
add mac-address=11:22:33:44:55:66 \
    action=reject \
    comment="Blocked device"
# Default policy (allow all others)
add action=accept \
    comment="Default allow"
```
#### Signal Strength-Based Access Control

```
# Reject weak signal clients (force roaming)
/caps-man access-list
add signal-range=-75..0 \
    action=accept \
    comment="Accept clients with good signal"
add signal-range=-120..-76 \
    action=reject \
    comment="Reject weak signal clients"
```
#### Time-Based Access Control

```
/caps-man access-list
add interface=cfg-guest \
    time=8h-18h,mon,tue,wed,thu,fri \
    action=accept \
    comment="Guest access during business hours only"
add interface=cfg-guest \
    action=reject \
    comment="Deny guest access outside hours"
```
### Dynamic Channel and Power Management

#### Automatic Channel Selection

