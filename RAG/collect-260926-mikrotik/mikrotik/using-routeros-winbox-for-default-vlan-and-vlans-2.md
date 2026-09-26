---
id: collect-260926-mikrotik/mikrotik/using-routeros-winbox-for-default-vlan-and-vlans-2
title: "create interfaces to be used for interaction with individual VLANs"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/using-routeros-winbox-for-default-vlan-and-vlans.md
source_anchor: ""
source_lines: [138, 290]
sha256: ebfd49012d2f61b290bfb097ad5029d33bb7dcaf55984f5e4ca7f5f40baf54ab
---

# create interfaces to be used for interaction with individual VLANs

The WAP plugged into the Mikrotik on ether1 isn’t happy. Most devices are connecting, dropping and connecting constantly on the SSID assigned to the Guest WiFi (vlan5). The other SSID assigned to the default VLAN, well nothing can connect to it. However, the management utility and the management IP address that is assigned to the WAP over the default VLAN work. Moving it back to the existing switch makes everything happy.

I’m going to dig out a POE power injector just to make sure its not a power issue from Mikrotik;s POE support. I’m also going port based vlan to assign 10 and 5 to ports 5 and 10 for testing with my laptop and the WAPs. I just can’t do that today.

             
            
           
          
            
            
              I didn’t write you can’t use “default VLANs” on Mikrotik, I just wrote that for me things are easier if everything is tagged internally. And yes, to do it like that you need a VLAN dedicated for that. On the upside, using this technique you can “partition” your switch to two (or more) parts, all of them untagged externally but separated from each other … and you don’t need multiple bridges for that (remember, only single bridge can be HW offloaded).

CRS is a switch with routing capabilities … so it might route traffic between VLANs if configuration allows.

To be able to give you further assistance, post actual configuration of CRS (run */export hide-sensitive* and post results inside [__code] [/code] environment). When doing inter-vlan ping test, have your test setup (CRS and the two devices) separated from the rest of network to be sure the problem isn’t elsewhere.

             
            
           
          
            
            
              ```
# jul/21/2020 10:19:13 by RouterOS 6.47.1
# software id = 75MP-QGGQ
#
# model = CRS328-24P-4S+
# serial number = D41D0B79CC86
/interface bridge
add admin-mac=C4:AD:34:F6:87:FD auto-mac=no name=DefaultBridge vlan-filtering=\
    yes
/interface vlan
add interface=DefaultBridge name=GuestWiFiVLAN5 vlan-id=5
add interface=DefaultBridge name=UntaggedToVLAN10 vlan-id=10
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
add bridge=DefaultBridge tagged=DefaultBridge untagged=ether1,ether9,ether23 \
    vlan-ids=10
add bridge=DefaultBridge tagged=DefaultBridge,ether1,ether9,ether23 vlan-ids=5
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
/ip address
add address=172.16.10.200 interface=GuestWiFiVLAN5 network=255.255.255.0
/ip dhcp-client
add disabled=no interface=DefaultBridge
/system clock
set time-zone-name=America/New_York
/system identity
set name=CRS328-1
/system ntp client
set enabled=yes primary-ntp=172.16.1.90 secondary-ntp=172.16.1.91
/system package update
set channel=testing
/system routerboard settings
set boot-os=router-os
/system swos
set identity=CRS328-1 static-ip-address=172.16.1.11
```

I really just want VLAN 5 and default on selected ports. Internal to switch VLANs are going to get messy, especially when we can switch to all VLANs

             
            
           
          
            
            
              Your setup is mostly fine, but probably not what you want. As it’s got IP addresses in both VLANs, it will route between the subnets if it gets opportunity. You wrote that you want to use device as switch, therefore you should set IP address only on management interface (which currently it should be either DefaultBridge or UntaggedToVLAN10 depending on which concept you’re going to go with). And DefaultBridge *interface* doesn’t have to be member of VLAN with which it doesn’t interact (where it doesn’t have IP address set).

You can take the opportunity to prepare for the all-tagged setup by pursuing the “internally tagged” setup you already started. Just configure DHCP client on UntaggedToVLAN10 interface and it should be fine.

BTW, when manually setting IP address, using CIDR netmask is mandatory, so your setup should be */ip address add address=172.16.10.200/24 interface=GuestWiFiVLAN5* (setting network is optional, in your case it’s actually wrong, it should be network address not netmask)

             
            
           
          
            
            
              So how would I do this without VLAN 10 and using just the default VLAN?

             
            
           
          
            
            
              So I change things up to this:

Ignore VLANS 6,99 and 100

