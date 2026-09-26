---
id: collect-260926-mikrotik/mikrotik/using-routeros-winbox-for-default-vlan-and-vlans-3
title: "create interfaces to be used for interaction with individual VLANs"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/using-routeros-winbox-for-default-vlan-and-vlans.md
source_anchor: ""
source_lines: [291, 407]
sha256: 6caf9abe01da16f62bd19d63615aad36a2b014db6c88acb66b1df38e39203b22
---

# create interfaces to be used for interaction with individual VLANs

```
# jul/21/2020 17:23:45 by RouterOS 6.47.1
# software id = 75MP-QGGQ
#
# model = CRS328-24P-4S+
# serial number = D41D0B79CC86
/interface bridge
add admin-mac=C4:AD:34:F6:87:FD auto-mac=no name=DefaultBridge vlan-filtering=\
    yes
/interface vlan
add interface=DefaultBridge name=EnterpriseWiFiVLAN6 vlan-id=6
add interface=DefaultBridge name=GuestWiFiVLAN5 vlan-id=5
add interface=DefaultBridge name=InternalSwitchVLAN99 vlan-id=99
add interface=DefaultBridge name=ManagementVLAN100 vlan-id=100
/interface list
add name=WAN
add name=LAN
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/ip hotspot profile
set [ find default=yes ] html-directory=flash/hotspot
/interface bridge port
add bridge=DefaultBridge interface=ether1
add bridge=DefaultBridge interface=ether2
add bridge=DefaultBridge interface=ether3
add bridge=DefaultBridge interface=ether4
add bridge=DefaultBridge interface=ether5
add bridge=DefaultBridge interface=ether6
add bridge=DefaultBridge interface=ether7
add bridge=DefaultBridge interface=ether8
add bridge=DefaultBridge interface=ether9
add bridge=DefaultBridge interface=ether10
add bridge=DefaultBridge interface=ether11
add bridge=DefaultBridge interface=ether12
add bridge=DefaultBridge interface=ether13
add bridge=DefaultBridge interface=ether14
add bridge=DefaultBridge interface=ether15
add bridge=DefaultBridge interface=ether16
add bridge=DefaultBridge interface=ether17
add bridge=DefaultBridge interface=ether18
add bridge=DefaultBridge interface=ether19
add bridge=DefaultBridge interface=ether20
add bridge=DefaultBridge interface=ether21
add bridge=DefaultBridge interface=ether22
add bridge=DefaultBridge interface=ether23
add bridge=DefaultBridge interface=ether24
add bridge=DefaultBridge interface=sfp-sfpplus1
add bridge=DefaultBridge interface=sfp-sfpplus2
add bridge=DefaultBridge interface=sfp-sfpplus3
add bridge=DefaultBridge interface=sfp-sfpplus4
/interface bridge vlan
add bridge=DefaultBridge tagged=GuestWiFiVLAN5,ether1,ether9,ether17 vlan-ids=5
add bridge=DefaultBridge tagged=EnterpriseWiFiVLAN6 vlan-ids=6
add bridge=DefaultBridge tagged=InternalSwitchVLAN100 vlan-ids=99
add bridge=DefaultBridge tagged=ManagementVLAN100 vlan-ids=100
/interface list member
add interface=ether1 list=LAN
add interface=ether2 list=LAN
add interface=ether3 list=LAN
add interface=ether4 list=LAN
add interface=ether5 list=LAN
add interface=ether6 list=LAN
add interface=ether7 list=LAN
add interface=ether8 list=LAN
add interface=ether9 list=LAN
add interface=ether10 list=LAN
add interface=ether11 list=LAN
add interface=ether12 list=LAN
add interface=ether13 list=LAN
add interface=ether14 list=LAN
add interface=ether15 list=LAN
add interface=ether16 list=LAN
add interface=ether17 list=LAN
add interface=ether18 list=LAN
add interface=ether19 list=LAN
add interface=ether20 list=LAN
add interface=ether21 list=LAN
add interface=ether22 list=LAN
add interface=ether23 list=LAN
add interface=ether24 list=LAN
add interface=sfp-sfpplus1 list=LAN
add interface=sfp-sfpplus2 list=LAN
add interface=sfp-sfpplus3 list=LAN
add interface=sfp-sfpplus4 list=LAN
/ip dhcp-client
add disabled=no interface=DefaultBridge
/system clock
set time-zone-name=America/New_York
/system identity
set name=CRS328-1
/system ntp client
set enabled=yes
/system ntp server
set manycast=no
/system routerboard settings
set boot-os=router-os
/system swos
set identity=CRS328-1 static-ip-address=172.16.1.11
```

Everything is on the default vlan, and ports 1,9 and 17 also have vlan 5 accessible. The default vlan and vlan5 work. However, the default vlan can reach vlan5. vlan5 does not appear to be able to reach the default VLAN.

             
            
           
          
            
            
              My bad, the config is right for vlan 5 and default vlan. The test router I plugged the switch into had the vlans routed together.

             
            
           
          
            
            
              While designing half breed labeled/untagged use, you can for the most part consider untagged one more VLAN, however it accompanies distinctive arranging orders. Also, this is the primary explanation I recomend the all-labeled methodology where one arrangements with untagged just in port setup, the rest SNMP, is then totally done consistently on VLANs.
