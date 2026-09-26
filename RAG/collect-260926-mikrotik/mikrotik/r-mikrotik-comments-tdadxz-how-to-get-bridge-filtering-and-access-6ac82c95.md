---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-tdadxz-how-to-get-bridge-filtering-and-access-6ac82c95
title: "model = RB3011UiAS"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/r-mikrotik-comments-tdadxz-how-to-get-bridge-filtering-and-access-6ac82c95.md
source_anchor: ""
source_lines: [1, 116]
sha256: 64fb2f1ef45154d167f4ece9694c08ea2283f8768243ebe9dcbfc1cae33bec5d
---

# model = RB3011UiAS

How to get bridge filtering and access restriction on vlans working properly 
        
        
        
    
    
    I can't seem to get this right. I am trying to set up 3 segmented networks by vlan: home, public, iot, which then connects to trunk bridge and out to the internet. I have 2 UniFi AP's which are working by tagging vlan traffic, and then in the MikroTik the AP vlans bridge ports have ingress-filters. This is all working and I get DHCP on the ethernet ports as well as the virtual AP's.
No matter how I configure it, I can't get vlan filtering working on the bridges. When I turn bridge filters on, I lose DHCP and access on all devices. I would also appreciate some help setting up the firewall so that my vlan devices cannot access each others' subnets.
OK here's my config, please let me know if anything else doesn't look right.
# model = RB3011UiAS
# serial number = XXXXX
/interface bridge
add name=home_bridge
add name=iot_bridge
add name=public_bridge
add admin-mac=XXXXX auto-mac=no name=trunk_bridge
/interface ethernet
set [ find default-name=ether10 ] poe-out=off
/interface vlan
add interface=ether6 name=vlan_home_ether6 vlan-id=2
add interface=ether7 name=vlan_home_ether7 vlan-id=2
add interface=ether6 name=vlan_iot_ether6 vlan-id=4
add interface=ether7 name=vlan_iot_ether7 vlan-id=4
add interface=ether6 name=vlan_public_ether6 vlan-id=3
add interface=ether7 name=vlan_public_ether7 vlan-id=3
/interface list
add name=WAN
add name=LAN
add name=VLAN
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/ip pool
add name=trunk_dchp_pool ranges=192.168.88.10-192.168.88.254
add name=iot_dhcp_pool ranges=192.168.87.2-192.168.87.254
add name=public_dhcp_pool ranges=192.168.86.2-192.168.86.254
add name=home_dhcp_pool ranges=192.168.1.2-192.168.1.254
/ip dhcp-server
add address-pool=trunk_dchp_pool disabled=no interface=trunk_bridge name=trunk_dhcp
add address-pool=iot_dhcp_pool disabled=no interface=iot_bridge name=iot_dhcp
add address-pool=public_dhcp_pool disabled=no interface=public_bridge name=public_dhcp
add address-pool=home_dhcp_pool disabled=no interface=home_bridge name=home_dhcp
/interface bridge port
add bridge=home_bridge interface=ether2 pvid=2
add bridge=home_bridge interface=ether3 pvid=2
add bridge=home_bridge interface=ether4 pvid=2
add bridge=trunk_bridge interface=ether5
add bridge=trunk_bridge interface=ether6
add bridge=trunk_bridge interface=ether7
add bridge=iot_bridge interface=ether8 pvid=4
add bridge=iot_bridge interface=ether9 pvid=4
add bridge=iot_bridge interface=ether10 pvid=4
add bridge=home_bridge interface=ether1 pvid=2
add bridge=iot_bridge frame-types=admit-only-vlan-tagged ingress-filtering=yes interface=vlan_iot_ether6 pvid=4
add bridge=iot_bridge frame-types=admit-only-vlan-tagged ingress-filtering=yes interface=vlan_iot_ether7 pvid=4
add bridge=public_bridge frame-types=admit-only-vlan-tagged ingress-filtering=yes interface=vlan_public_ether6 pvid=3
add bridge=public_bridge frame-types=admit-only-vlan-tagged ingress-filtering=yes interface=vlan_public_ether7 pvid=3
add bridge=home_bridge frame-types=admit-only-vlan-tagged ingress-filtering=yes interface=vlan_home_ether6 pvid=2
add bridge=home_bridge frame-types=admit-only-vlan-tagged ingress-filtering=yes interface=vlan_home_ether7 pvid=2
/ip neighbor discovery-settings
set discover-interface-list=LAN
/interface bridge vlan
add bridge=home_bridge tagged=vlan_home_ether6,vlan_home_ether7 untagged=ether1,ether2,ether3,ether4 vlan-ids=2
add bridge=iot_bridge tagged=vlan_iot_ether6,vlan_iot_ether7 untagged=ether8,ether9,ether10 vlan-ids=4
add bridge=public_bridge tagged=vlan_public_ether6,vlan_public_ether7 vlan-ids=3
/interface list member
add interface=trunk_bridge list=LAN
add interface=sfp1 list=WAN
add interface=iot_bridge list=VLAN
add interface=public_bridge list=VLAN
add interface=home_bridge list=VLAN
/ip address
add address=192.168.88.1/24 interface=trunk_bridge network=192.168.88.0
add address=192.168.1.1/24 interface=home_bridge network=192.168.1.0
add address=192.168.87.1/24 interface=iot_bridge network=192.168.87.0
add address=192.168.86.1/24 interface=public_bridge network=192.168.86.0
/ip dhcp-client
add interface=ether1
add !dhcp-options disabled=no interface=sfp1
/ip dhcp-server lease
add address=192.168.1.2 client-id=XXXXX mac-address=XXXXX server=home_dhcp
/ip dhcp-server network
add address=192.168.1.0/24 gateway=192.168.1.1
add address=192.168.86.0/24 gateway=192.168.86.1
add address=192.168.87.0/24 gateway=192.168.87.1
add address=192.168.88.0/24 gateway=192.168.88.1
/ip dns
set allow-remote-requests=yes
/ip dns static
add address=192.168.88.1 name=trunk_dns
add address=192.168.87.1 name=iot_dns
add address=192.168.1.1 name=home_dns
add address=192.168.86.1 name=public_dns
/ip firewall filter
add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid
add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
add action=accept chain=input comment="defconf: accept to local loopback (for CAPsMAN)" dst-address=127.0.0.1
add action=drop chain=input comment="defconf: drop all not coming from LAN" in-interface-list=!LAN
add action=accept chain=forward comment="defconf: accept in ipsec policy" ipsec-policy=in,ipsec
add action=accept chain=forward comment="defconf: accept out ipsec policy" ipsec-policy=out,ipsec
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related
add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
add action=drop chain=forward comment="defconf: drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN
add action=accept chain=forward connection-state=new in-interface-list=VLAN out-interface-list=LAN
add action=drop chain=forward connection-state=new in-interface-list=LAN out-interface-list=VLAN
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none out-interface-list=WAN
Section des commentaires
The config parsing on reddit it's a mess, but you need to bridge the ports (I personally bridge all the ports except on custom setups) then add the interface-vlan on the bridge.
Then go to bridge, vlan, define ports and vlans, tagged and untagged, then define the PVID on untagged ports and lastly enable vlan filtering on the bridge in safe mode.
I saw a lot of things pointing to interfaces (dhcp is one) that won't work.
yeah I just spent like 10 minutes trying to fix that, try to refresh
So that is how I did it with the APs, since the virtual APs are setting the vlan ids. However, on the ports, I want the port to set the vlan id and then access the bridge, since the devices are just computers/printers/etc. and don't set their own vlan ids.
It won't get any better, I'm on mobile. try to apply the steps I suggested and give feedback on how it went.
You should have it working at the end of the day.
