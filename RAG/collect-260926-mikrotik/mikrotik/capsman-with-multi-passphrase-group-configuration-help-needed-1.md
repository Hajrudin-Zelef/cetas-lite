---
id: collect-260926-mikrotik/mikrotik/capsman-with-multi-passphrase-group-configuration-help-needed-1
title: "NAME      VERSION  BUILD-TIME           SIZE"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2026-01-19", "2026-02-01"]
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/capsman-with-multi-passphrase-group-configuration-help-needed.md
source_anchor: ""
source_lines: [1, 107]
sha256: 26656ca08e2ac0f4504764bdf6a43ed35d8e18ae3d704e569fd61b483a611987
---

# NAME      VERSION  BUILD-TIME           SIZE

Hi all.

I am trying to get the CAPsMAN with multi passphrase group configuration to work and could need some help.

The following VLANs should be available via wifi:

- 1000

- 1010

- 1060

The L2 connectivity is proven working between CAPsMAN device and AP.

The AP has an IP address in the prefix on VLAN 999 (no wifi for this VLAN) over which it should be managed by CAPsMAN (also used for SSH and API calls etc.).

I want to make use of the multi passphrase groups. VLAN assignment seems to work according to the registration table on the CAPsMAN device, but there is no traffic I can see.

I had a working CAPsMAN setup with the old wireless package (if I am not mistaken) but as this is deprecated I wanted to migrate over to the wifi package.

Below is the current configuration (not working) with all the unused interfaces etc. removed. The VLANs are tagged on all network devices along the paht.

**CAPsMAN**

```
[user@capsman] > /system/package/print 
Flags: X - DISABLED
Columns: NAME, VERSION, BUILD-TIME, SIZE
#   NAME      VERSION  BUILD-TIME           SIZE    
0 X wireless  7.21.1   2026-01-19 15:09:07  860.1KiB
1   routeros  7.21.1   2026-01-19 15:09:07  12.8MiB 
[user@capsman] > /interface/export show-sensitive 
# 2026-02-01 16:43:43 by RouterOS 7.21.1
# model = CCR2116-12G-4S+
/interface ethernet
set [ find default-name=sfp-sfpplus1 ] comment=\
"Core;spinesw01.home sfp-sfpplus1;C0000118;" l2mtu=1598
set [ find default-name=sfp-sfpplus2 ] comment=\
"Core;spinesw02.home sfp-sfpplus1;C0000119;" l2mtu=1598
set [ find default-name=sfp-sfpplus3 ] comment=";;;" disabled=yes l2mtu=1598
set [ find default-name=sfp-sfpplus4 ] comment=";;;" disabled=yes l2mtu=1598
/interface bonding
add comment="LACP to spinesw0[12].home" mode=802.3ad name=bond1 slaves=sfp-sfpplus1,sfp-sfpplus2 \
transmit-hash-policy=layer-2-and-3
/interface vlan
add interface=bond1 name=vl-999-mgmt vlan-id=999
add interface=bond1 name=vl-1000-int vlan-id=1000
add interface=bond1 name=vl-1010-media vlan-id=1010
add interface=bond1 name=vl-1060-work vlan-id=1060
/interface vrrp
add interface=vl-1000-int name=vrrp-int priority=254 remote-address=10.10.0.252 \
sync-connection-tracking=yes
add interface=vl-1010-media name=vrrp-media priority=254 remote-address=10.10.1.252 \
sync-connection-tracking=yes
add interface=vl-999-mgmt name=vrrp-mgmt priority=254 remote-address=10.10.99.252 \
sync-connection-tracking=yes
add interface=vl-1060-work name=vrrp-work priority=254 remote-address=10.10.6.252 \
sync-connection-tracking=yes
/interface wifi channel
add disabled=no name=private-2 skip-dfs-channels=10min-cac
add band=5ghz-ax disabled=yes name=private-5 skip-dfs-channels=10min-cac width=20/40/80mhz
/interface wifi datapath
add disabled=no name=int traffic-processing=on-cap vlan-id=1000
add disabled=no name=media traffic-processing=on-cap vlan-id=1010
add disabled=no name=work traffic-processing=on-cap vlan-id=1060
/interface wifi security
add authentication-types=wpa2-psk disabled=no encryption=ccmp ft=yes ft-over-ds=yes \
multi-passphrase-group=private name=wifi-private wps=disable
/interface wifi configuration
add channel=private-2 country=Germany datapath=media disabled=no installation=indoor mode=ap name=media-2 \
security=wifi-private ssid=myssid
add channel=private-2 country=Germany datapath=work disabled=no installation=indoor mode=ap name=work-2 \
security=wifi-private ssid=myssid
add channel=private-2 country=Germany datapath=int disabled=no installation=indoor mode=ap name=int-2 \
security=wifi-private ssid=myssid
add channel=private-5 country=Germany datapath=int disabled=no installation=indoor mode=ap name=int-5 \
security=wifi-private ssid=myssid
add channel=private-5 country=Germany datapath=media disabled=no installation=indoor mode=ap name=\
    media-5 security=wifi-private ssid=myssid
add channel=private-5 country=Germany datapath=work disabled=no installation=indoor mode=ap name=work-5 \
security=wifi-private ssid=myssid
/interface ethernet switch
set 0 l3-hw-offloading=yes
/interface wifi capsman
set ca-certificate=WiFi-CAPsMAN-CA-something enabled=yes interfaces=vl-999-mgmt
/interface wifi provisioning
add action=create-dynamic-enabled disabled=no master-configuration=int-2 name-format=%I-2.4 \
slave-configurations=media-2,work-2 slave-name-format=%I-2.4-virtual supported-bands=2ghz-ax
add action=create-dynamic-enabled disabled=no master-configuration=int-5 name-format=%I-5 \
slave-configurations=media-5,work-5 slave-name-format=%I-5-virtual supported-bands=5ghz-ax
/interface wifi security multi-passphrase
add comment=int disabled=no group=private passphrase=pw1 vlan-id=1000
add comment=media disabled=no group=private passphrase=pw2 vlan-id=1010
add comment=work disabled=no group=private passphrase=pw3 vlan-id=1060
/ip address
add address=10.10.99.253/24 interface=vl-999-mgmt network=10.10.99.0
add address=10.10.99.254/24 interface=vrrp-mgmt network=10.10.99.0
add address=10.10.0.253/24 interface=vl-1000-int network=10.10.0.0
add address=10.10.0.254/24 interface=vrrp-int network=10.10.0.0
add address=10.10.1.253/24 interface=vl-1010-media network=10.10.1.0
add address=10.10.1.254/24 interface=vrrp-media network=10.10.1.0
add address=10.10.6.253/24 interface=vl-1060-work network=10.10.6.0
add address=10.10.6.254/24 interface=vrrp-work network=10.10.6.0
```

**Accesspoint**

