---
id: collect-260926-mikrotik/mikrotik/capsman
title: "CAPsMAN"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/capsman.md
source_anchor: ""
source_lines: [1, 73]
sha256: 3dbfffb6e17bcad00304923df8e0c4ec68f7679ed828672486bab72dc8f3eede
---

# CAPsMAN

## CAPsMAN â Centralized AP Management

CAPsMAN (Controlled Access Point system Manager) allows a MikroTik router to centrally manage multiple access points.

## Enable CAPsMAN Manager

On the controller router:

```
/caps-man manager set enabled=yes
```
## Create a Channel

```
/caps-man channel
add name=channel-2ghz frequency=2412 band=2ghz-b/g/n tx-power=17 width=20mhz
add name=channel-5ghz frequency=5180 band=5ghz-a/n/ac tx-power=23 width=40mhz
```
## Create Security Profile

```
/caps-man security
add name=security-wpa2 \
    authentication-types=wpa2-psk \
    encryption=aes-ccm \
    passphrase="YourWiFiPassword"
```
## Create Datapath

```
/caps-man datapath
add name=datapath-local \
    bridge=bridge-local \
    client-to-client-forwarding=yes
```
## Create Configuration

```
/caps-man configuration
add name=config-2ghz \
    ssid="YourNetwork" \
    channel=channel-2ghz \
    security=security-wpa2 \
    datapath=datapath-local \
    multicast-helper=full
add name=config-5ghz \
    ssid="YourNetwork-5G" \
    channel=channel-5ghz \
    security=security-wpa2 \
    datapath=datapath-local
```
## Provisioning Rules

```
/caps-man provisioning
add action=create-dynamic-enabled \
    master-configuration=config-2ghz \
    slave-configurations=config-5ghz
```
## On the Access Point (CAP)

```
/interface wireless cap
set enabled=yes caps-man-addresses=192.168.1.1 interfaces=wlan1,wlan2
```
## Verify

```
/caps-man remote-cap print
/caps-man interface print
```
