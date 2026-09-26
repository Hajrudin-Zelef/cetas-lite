---
id: collect-260926-mikrotik/mikrotik/mikrotik-capsman-the-complete-guide-to-centralized-wireless-network-management-tech-layer--3
title: "Floor 1 CAPs (identified by identity prefix)"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "latency", "memory"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/mikrotik-capsman-the-complete-guide-to-centralized-wireless-network-management-tech-layer-x-com.md
source_anchor: ""
source_lines: [558, 875]
sha256: e36c725c028d4ffb9914cc2d09a8ac84e34e720186593553aa146fe13f86a8c3
---

# Floor 1 CAPs (identified by identity prefix)

```
/caps-man channel
add name=channel-5ghz-auto \
    band=5ghz-n/ac \
    control-channel-width=20mhz \
    extension-channel=eeCe \
    reselect-interval=1h \
    skip-dfs-channels=yes
```
Note: Leave frequency empty for automatic selection.

#### Per-CAP Channel Assignment

```
/caps-man provisioning
add identity-regexp="^CAP-F1-01" \
    master-configuration=cfg-corp-ch36 \
    comment="Channel 36 for CAP-F1-01"
add identity-regexp="^CAP-F1-02" \
    master-configuration=cfg-corp-ch44 \
    comment="Channel 44 for CAP-F1-02"
```
### Load Balancing Between Access Points

#### Limit Clients Per Radio

/caps-man configuration
set cfg-corporate-5ghz max-sta-count=40
/caps-man configuration
set cfg-corporate-2ghz max-sta-count=25

#### Band Steering (2.4GHz to 5GHz)

# Lower tx-power on 2.4GHz to encourage 5GHz
/caps-man channel
set channel-2ghz tx-power=14
# Higher tx-power on 5GHz
/caps-man channel
set channel-5ghz tx-power=20

### Multiple SSID Configuration

#### Corporate + Guest + IoT Setup

```
# Datapaths with different VLANs
/caps-man datapath
add name=dp-corporate vlan-id=10 vlan-mode=use-tag local-forwarding=yes
add name=dp-guest vlan-id=20 vlan-mode=use-tag local-forwarding=yes client-to-client-forwarding=no
add name=dp-iot vlan-id=30 vlan-mode=use-tag local-forwarding=yes
# Configurations for each SSID
/caps-man configuration
add name=cfg-corporate ssid="CORP-SECURE" datapath=dp-corporate security=sec-wpa2-ent
add name=cfg-guest ssid="GUEST" datapath=dp-guest security=sec-wpa2-psk
add name=cfg-iot ssid="IoT-Devices" datapath=dp-iot security=sec-iot hide-ssid=yes
# Provisioning with all SSIDs
/caps-man provisioning
add master-configuration=cfg-corporate \
    slave-configurations=cfg-guest,cfg-iot \
    action=create-dynamic-enabled
```
### Hotspot Integration

```
# Create datapath for hotspot
/caps-man datapath
add name=dp-hotspot \
    local-forwarding=no \
    bridge=bridge-hotspot
# Configure hotspot on controller
/ip hotspot
add name=hotspot-guest \
    interface=bridge-hotspot \
    address-pool=hotspot-pool \
    profile=hsprof-guest
# Guest configuration with hotspot datapath
/caps-man configuration
add name=cfg-hotspot \
    ssid="Free-WiFi" \
    datapath=dp-hotspot \
    security.authentication-types=""
```
## 7. Monitoring and Troubleshooting

### Monitoring Tools and Commands

#### View Registered CAPs

/caps-man remote-cap print

Output shows:

- CAP identity and MAC address
- State (running, disabled)
- Version
- Radio details

#### View Active Interfaces

/caps-man interface print

#### View Connected Clients

/caps-man registration-table print

Output shows:

- Client MAC address
- Interface name
- Signal strength (rx/tx)
- Data rates
- Uptime
- Packets/bytes transferred

#### Monitor Real-Time Client Activity

/caps-man registration-table print interval=1

#### View Client Statistics

/caps-man registration-table print stats

### SNMP Monitoring

#### Enable SNMP on Controller

/snmp
set enabled=yes contact="admin@company.com" location="DataCenter"
/snmp community
set public addresses=10.0.1.0/24 read-access=yes

#### Key SNMP OIDs for CAPsMAN

| OID | Description | 
|---|---|
| .1.3.6.1.4.1.14988.1.1.1.2 | Wireless registration table | 
| .1.3.6.1.4.1.14988.1.1.1.3 | Wireless interface statistics | 
| .1.3.6.1.2.1.2.2 | Interface table (standard MIB-II) | 

### Common Issues and Solutions

#### Issue: CAP Not Connecting to Controller

Troubleshooting steps:

1. Verify network connectivity:
# On CAP /ping 10.0.1.1
2. Check CAP configuration:
/interface wireless cap print
3. Verify CAPsMAN is enabled:
# On Controller /caps-man manager print
4. Check firewall rules:
# Allow CAPsMAN traffic /ip firewall filter add chain=input protocol=udp dst-port=5246-5247 action=accept
5. Check RouterOS version match

#### Issue: Clients Cannot Connect

Troubleshooting steps:

1. Verify interface is running:
/caps-man interface print where disabled=no
2. Check security configuration:
/caps-man security print
3. Verify RADIUS connectivity (if using):
/radius print /radius monitor 0
4. Check access lists:
/caps-man access-list print

#### Issue: Poor Roaming Performance

Solutions:

- Enable access list with signal threshold:
/caps-man access-list add signal-range=-70..0 action=accept add signal-range=-120..-71 action=reject
- Use same datapath for all APs
- Enable local forwarding to reduce latency
- Verify overlapping coverage between APs

#### Issue: Intermittent Disconnections

Solutions:

- Check for channel interference
- Increase disconnect timeout:
/caps-man configuration set cfg-corporate disconnect-timeout=10s
- Verify CAP-to-controller link stability
- Check for power issues on CAPs

### Logging and Debugging

#### Enable CAPsMAN Debug Logging

/system logging
add topics=caps,debug action=memory
add topics=wireless,debug action=memory

#### View Logs

/log print where topics~"caps"

#### Enable Remote Logging

/system logging action
add name=remote-syslog target=remote remote=10.0.1.200
/system logging
add topics=caps action=remote-syslog

## 8. Security Best Practices

### Securing the Management Plane

#### Restrict CAPsMAN Access

/ip firewall filter
# Allow CAPsMAN from management network only
add chain=input src-address=10.0.1.0/24 protocol=udp dst-port=5246-5247 action=accept comment="CAPsMAN from mgmt"
add chain=input protocol=udp dst-port=5246-5247 action=drop comment="Block other CAPsMAN"

#### Management VLAN Isolation

# Create management VLAN
/interface vlan
add name=vlan-mgmt vlan-id=99 interface=bridge
# Place CAPs on management VLAN
/interface bridge vlan
add bridge=bridge tagged=bridge,ether1,ether2 vlan-ids=99

#### Disable Unnecessary Services on CAPs

/ip service
set telnet disabled=yes
set ftp disabled=yes
set www disabled=yes
set api disabled=yes
set api-ssl disabled=yes
set winbox address=10.0.1.0/24
set ssh address=10.0.1.0/24

### Wireless Security Recommendations

#### Use WPA3 Where Possible (RouterOS 7+)

```
/caps-man security
add name=sec-wpa3 \
    authentication-types=wpa3-psk \
    encryption=aes-ccm \
    passphrase="ComplexPassword123!"
```
#### Enable Protected Management Frames

/caps-man configuration
set cfg-corporate security.pmf=required

#### Implement 802.1X with RADIUS

```
# RADIUS server configuration
/radius
add address=10.0.1.100 secret="RadiusSecret123" service=wireless
# Security profile for EAP
/caps-man security
add name=sec-enterprise \
    authentication-types=wpa2-eap \
    encryption=aes-ccm \
    eap-methods=passthrough
# AAA settings
/caps-man aaa
set interim-update=5m \
    called-format=mac:ssid \
    mac-mode=as-username
```
#### Disable Legacy Protocols

# Use only modern standards
/caps-man channel
set channel-5ghz band=5ghz-onlyac
/caps-man configuration
set cfg-corporate supported-rates-b="" basic-rates-b=""

### Security Audit Checklist

1. Verify WPA2/WPA3 encryption enabled
2. Confirm PMF is enabled (required or preferred)
3. Validate RADIUS server connectivity
4. Review access list entries
5. Check firewall rules for CAPsMAN traffic
6. Verify certificate-based CAP authentication
7. Audit guest network isolation
8. Review password complexity for PSK networks
9. Confirm management access restrictions
10. Validate VLAN assignments

### Regular Maintenance Tasks

- **Weekly:** Review logs for anomalies
- **Monthly:** Backup configurations
- **Quarterly:** Review and rotate passwords
- **Semi-annually:** Update firmware
- **Annually:** Full security audit

#### Backup Configuration Script

```
/system scheduler
add name=weekly-backup interval=7d on-event={
    /system backup save name=([/system identity get name] . "-" . [:pick [/system clock get date] 0 10])
    /export file=([/system identity get name] . "-config-" . [:pick [/system clock get date] 0 10])
}
```
## 9. High Availability and Redundancy

### Dual CAPsMAN Controller Setup

CAPs can be configured with primary and backup controller addresses.

#### Configure CAP with Multiple Controllers

