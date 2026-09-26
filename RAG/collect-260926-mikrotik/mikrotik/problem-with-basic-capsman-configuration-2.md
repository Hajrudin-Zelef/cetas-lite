---
id: collect-260926-mikrotik/mikrotik/problem-with-basic-capsman-configuration-2
title: "managed by CAPsMAN"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/problem-with-basic-capsman-configuration.md
source_anchor: ""
source_lines: [244, 292]
sha256: 79c6772db9c158a86d5ca9dfb569ea35b3b93435bb32480f045b99c0a6bc1b5a
---

# managed by CAPsMAN

In this topic I get a suggestion that when I use CAPsMAN, it automatically add WLAN interface (in this case CAP interface) to the bridge.

Current config on the CAPsMAN device:

```
/interface wireless
# managed by CAPsMAN
# channel: 2412/20-Ce/gn(20dBm), SSID: abcd, CAPsMAN forwarding
set [ find default-name=wlan1 ] ssid=MikroTik
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/interface wireless cap
set caps-man-addresses=127.0.0.1 enabled=yes interfaces=wlan1
/interface bridge
add name=bridge1 protocol-mode=none
/interface bridge port
add bridge=bridge1 interface=ether1
/caps-man datapath
add bridge=bridge1 name=datapath1
/caps-man configuration
add country="czech republic" datapath=datapath1 mode=ap name=abcd-cfg ssid=abcd
/caps-man manager
set enabled=yes
/caps-man provisioning
add action=create-dynamic-enabled master-configuration=abcd-cfg name-format=identity
```

Current config on the CAP device:

```
/interface wireless
# managed by CAPsMAN
# channel: 2447/20-Ce/gn(20dBm), SSID: abcd, CAPsMAN forwarding
set [ find default-name=wlan1 ] ssid=MikroTik
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/interface wireless cap
set bridge=bridge1 discovery-interfaces=bridge1 enabled=yes interfaces=wlan1
/interface bridge
add name=bridge1 protocol-mode=none
/interface bridge port
add bridge=bridge1 interface=ether1
add bridge=bridge1 interface=ether2
add bridge=bridge1 interface=ether3
add bridge=bridge1 interface=ether4
add bridge=bridge1 interface=ether5
```

It seems that it works fine. Thank everyone for your help and recommendations.
