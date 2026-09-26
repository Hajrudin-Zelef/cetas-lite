---
id: collect-260926-mikrotik/mikrotik/firewall-rules-to-seperate-some-vlans
title: "firewall-rules-to-seperate-some-vlans"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/firewall-rules-to-seperate-some-vlans.md
source_anchor: ""
source_lines: [1, 129]
sha256: db63504ce0722afa8a54b7432cb84462ee28501b95df2fbc6c49c48c7f8fba39
---

# firewall-rules-to-seperate-some-vlans

Hi,

I’m using Mikrotik devices for several years now, mostly with routed and unrouted vlans and basic switching.

For some reasons I want to implement a firewall to sperate some vLANs from the others.

In my setup I have a transfer-vlan that is connected to my ha pfsense boxes (in the future I will replace them but the transfer vlan should stay).

The central part of my lan consists of two crs317 which are configured redundant with vrrp and they are responsible for most internal routing.

In preparation of decommisionin g my pfsense I plan to put 2 networks which are internaly handled bei pfsense as a vlan on my mikrotik network. These should be separated from my other vlans but should also be accesible by others.

I’m glad I have some spare hex poe where I could setup a similar routing environment as my production one so I can test almost everything before trying it in production.

I have some experience with pfsense and sophos SG/XG but I donÄt get it using Mikrotik filters.

I read some things about firewalling mostly inspired by http://forum.mikrotik.com/t/inter-vlan-routing-and-default-firewall/146841/1 and http://forum.mikrotik.com/t/best-practice-firewall-inter-vlan-routing/139923/1 and thought this could not be that hard.

So for testing I setup two HEX PoE configured them with the needed vLANs, DHCP and vrrp to mostly match my production. I set up some lists and filter rules but they don’t work as expected.



my testing config looks like this:

```
/interface bridge
add name=bridge priority=0x1000 vlan-filtering=yes
/interface vlan
add interface=bridge name=vlan1-Default vlan-id=1
add interface=bridge name=vlan2-Schild vlan-id=2
add interface=bridge name=vlan3-Mama vlan-id=3
add interface=bridge name=vlan4-Gast vlan-id=4
add interface=bridge name=vlan5-SmartHome vlan-id=5
add interface=bridge name=vlan9-Transfer vlan-id=9
/interface vrrp
add interface=vlan1-Default name=vrrp101 priority=150
add interface=vlan2-Schild name=vrrp102 priority=150
add interface=vlan3-Mama name=vrrp103 priority=150
add interface=vlan4-Gast name=vrrp104 priority=150
add interface=vlan5-SmartHome name=vrrp105 priority=150
add interface=vlan9-Transfer name=vrrp109 priority=150
/interface list
add name=LAN
add name=WAN
add name=Schild
add name=Gast
add name=Mama
/ip pool
add name=dhcp_pool1 ranges=192.168.13.1-192.168.13.253
add name=dhcp_pool2 ranges=192.168.20.1-192.168.20.253
add name=dhcp_pool3 ranges=192.168.30.1-192.168.30.253
add name=dhcp_pool4 ranges=192.168.40.1-192.168.40.253
add name=dhcp_pool5 ranges=192.168.50.1-192.168.50.253
/ip dhcp-server
add address-pool=dhcp_pool1 interface=vrrp101 name=dhcp1
add address-pool=dhcp_pool2 interface=vrrp102 name=dhcp2
add address-pool=dhcp_pool3 interface=vrrp103 name=dhcp3
add address-pool=dhcp_pool4 interface=vrrp104 name=dhcp4
add address-pool=dhcp_pool5 interface=vrrp105 name=dhcp5
/interface bridge port
add bridge=bridge interface=ether1
add bridge=bridge interface=ether2
add bridge=bridge interface=ether3
add bridge=bridge interface=ether4
add bridge=bridge interface=ether5
add bridge=bridge interface=sfp1
/interface bridge vlan
add bridge=bridge tagged=bridge vlan-ids=1
add bridge=bridge tagged=bridge,sfp1 vlan-ids=2
add bridge=bridge tagged=bridge,sfp1 vlan-ids=3
add bridge=bridge tagged=bridge,sfp1 vlan-ids=4
add bridge=bridge tagged=bridge,sfp1 vlan-ids=5
add bridge=bridge tagged=bridge,sfp1 vlan-ids=9
/interface list member
add interface=vlan1-Default list=LAN
add interface=vlan2-Schild list=LAN
add interface=vlan3-Mama list=LAN
add interface=vlan4-Gast list=LAN
add interface=vlan5-SmartHome list=LAN
add interface=vlan1-Default list=Schild
add interface=vlan2-Schild list=Schild
add interface=vlan3-Mama list=Mama
add interface=vlan4-Gast list=Gast
add interface=vlan9-Transfer list=WAN
/ip address
add address=192.168.13.61/24 interface=vlan1-Default network=192.168.13.0
add address=192.168.20.61/24 interface=vlan2-Schild network=192.168.20.0
add address=192.168.30.61/24 interface=vlan3-Mama network=192.168.30.0
add address=192.168.40.61/24 interface=vlan4-Gast network=192.168.40.0
add address=192.168.50.61/24 interface=vlan5-SmartHome network=192.168.50.0
add address=192.168.13.254 interface=vrrp101 network=192.168.13.0
add address=192.168.20.254 interface=vrrp102 network=192.168.20.0
add address=192.168.30.254 interface=vrrp103 network=192.168.30.0
add address=192.168.40.254 interface=vrrp104 network=192.168.40.0
add address=192.168.50.254 interface=vrrp105 network=192.168.50.0
add address=192.168.90.254 interface=vrrp109 network=192.168.90.0
add address=192.168.90.61/24 interface=vlan9-Transfer network=192.168.90.0
/ip dhcp-server network
add address=192.168.13.0/24 dns-server=192.168.13.254 gateway=192.168.13.254
add address=192.168.20.0/24 dns-server=192.168.20.254 gateway=192.168.20.254
add address=192.168.30.0/24 dns-server=192.168.30.254 gateway=192.168.30.254
add address=192.168.40.0/24 dns-server=192.168.40.254 gateway=192.168.40.254
add address=192.168.50.0/24 dns-server=192.168.50.254 gateway=192.168.50.254
/ip firewall filter
add action=accept chain=input comment="accept ICMP" protocol=icmp
add action=accept chain=input comment="accept to local loopback (for CAPsMAN)" dst-address=127.0.0.1
add action=accept chain=forward comment="accept in ipsec policy" ipsec-policy=in,ipsec
add action=accept chain=forward comment="accept out ipsec policy" ipsec-policy=out,ipsec
add action=fasttrack-connection chain=forward comment=fasttrack connection-state=established,related hw-offload=yes
add action=accept chain=forward comment="accept established,related, untracked" connection-state=established,related,untracked
add action=drop chain=forward comment="drop invalid" connection-state=invalid
add action=accept chain=input comment="(DEBUG) Accept DHCP request on LAN interfaces" dst-port=67 in-interface-list=LAN protocol=udp src-port=68
add action=accept chain=input comment="(DEBUG) Accept DNS request (UDP) on LAN interfaces" dst-port=53 in-interface-list=LAN protocol=udp
add action=accept chain=input comment="(DEBUG) Accept DNS request (TCP) on LAN interfaces" dst-port=53 in-interface-list=LAN protocol=tcp
add action=accept chain=input comment="(DEBUG) Accept NTP request (UDP) on LAN interfaces" dst-port=123 in-interface-list=LAN protocol=udp
add action=accept chain=input comment="(DEBUG) Accept NTP request (TCP) on LAN interfaces" dst-port=123 in-interface-list=LAN protocol=tcp
add action=accept chain=forward in-interface-list=LAN out-interface-list=WAN
add action=accept chain=forward in-interface-list=Schild log=yes out-interface-list=Mama
add action=drop chain=forward comment="Drop All Else" disabled=yes
add action=drop chain=input comment="drop all not coming from LAN" disabled=yes in-interface-list=!Schild protocol=tcp
```

The second HEX PoE is configured ideticaly, except filter rules, IPs, STP and vrrp priority.

As a Test I put some interfaces in different vLANs and tried to ping the second HEX on its vLAN-interfaces, so that the packets shoould go through the gateway.

I can ping from each vLAN to each vLAN, I don’t see any Filter to be used as the counters don’t increase but when I enable the forward block rule everything gets blocked.

I cannot find my mistake.
