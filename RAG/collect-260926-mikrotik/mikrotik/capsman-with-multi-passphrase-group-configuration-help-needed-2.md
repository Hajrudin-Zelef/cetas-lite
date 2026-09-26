---
id: collect-260926-mikrotik/mikrotik/capsman-with-multi-passphrase-group-configuration-help-needed-2
title: "NAME      VERSION  BUILD-TIME           SIZE"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2026-01-19", "2026-02-01"]
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/capsman-with-multi-passphrase-group-configuration-help-needed.md
source_anchor: ""
source_lines: [108, 159]
sha256: 8fc3fe06f3161e42297356b0329126bc13b5814820cbba824a12c634e6aa3054
---

# NAME      VERSION  BUILD-TIME           SIZE

```
[user@cap] > /system/package/print 
Columns: NAME, VERSION, BUILD-TIME, SIZE
# NAME       VERSION  BUILD-TIME           SIZE   
0 routeros   7.21.1   2026-01-19 15:09:07  12.8MiB
1 wifi-qcom  7.21.1   2026-01-19 15:09:07  12.1MiB
[user@cap] > /interface/export show-sensitive 
# 2026-02-01 16:44:01 by RouterOS 7.21.1
# model = cAPGi-5HaxD2HaxD
/interface bridge
add mtu=1500 name=general-bridge vlan-filtering=yes
/interface wifi
# managed by CAPsMAN capsmanmacaddress%vl-999-mgmt, traffic processing on CAP
# mode: AP, SSID: myssid, channel: 5680/ax/eCee/D
set [ find default-name=wifi1 ] channel.band=5ghz-ax .skip-dfs-channels=10min-cac .width=20/40/80mhz \
configuration.manager=capsman .mode=ap .ssid=MikroTik-something disabled=no \
security.authentication-types=wpa2-psk,wpa3-psk .ft=yes .ft-over-ds=yes .passphrase=defaultpw
# managed by CAPsMAN capsmanmacaddress%vl-999-mgmt, traffic processing on CAP
# mode: AP, SSID: myssid, channel: 2412/ax/Ce
set [ find default-name=wifi2 ] channel.band=2ghz-ax .skip-dfs-channels=10min-cac .width=20/40mhz \
configuration.manager=capsman .mode=ap .ssid=MikroTik-something disabled=no \
security.authentication-types=wpa2-psk,wpa3-psk .ft=yes .ft-over-ds=yes .passphrase=defaultpw
# managed by CAPsMAN capsmanmacaddress%vl-999-mgmt, traffic processing on CAP
# mode: AP, SSID: myssid
add disabled=no master-interface=wifi1 name=wifi5
# managed by CAPsMAN capsmanmacaddress%vl-999-mgmt, traffic processing on CAP
# mode: AP, SSID: myssid
add disabled=no master-interface=wifi1 name=wifi6
# managed by CAPsMAN capsmanmacaddress%vl-999-mgmt, traffic processing on CAP
# mode: AP, SSID: myssid
add disabled=no master-interface=wifi2 name=wifi7
# managed by CAPsMAN capsmanmacaddress%vl-999-mgmt, traffic processing on CAP
# mode: AP, SSID: myssid
add disabled=no master-interface=wifi2 name=wifi8
/interface ethernet
set [ find default-name=ether1 ] comment="Access;WPS-Corridor-01 1;;"
set [ find default-name=ether2 ] comment=";;;" disabled=yes
/interface vlan
add interface=general-bridge name=vl-999-mgmt vlan-id=999
/interface bridge port
add bridge=general-bridge frame-types=admit-only-vlan-tagged interface=ether1
/interface bridge vlan
add bridge=general-bridge tagged=ether1,general-bridge vlan-ids=999
add bridge=general-bridge tagged=ether1 vlan-ids=1000
add bridge=general-bridge tagged=ether1 vlan-ids=1010
add bridge=general-bridge tagged=ether1 vlan-ids=1060
/interface wifi cap
set caps-man-addresses=10.10.99.253 caps-man-names=fw01.home certificate=request \
discovery-interfaces=vl-999-mgmt enabled=yes slaves-static=yes
/ip address
add address=10.10.99.243/24 interface=vl-999-mgmt network=10.10.99.0
```
