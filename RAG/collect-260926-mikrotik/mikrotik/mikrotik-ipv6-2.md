---
id: collect-260926-mikrotik/mikrotik/mikrotik-ipv6-2
title: "in/out-bridge-port matcher not possible when interface (vlan832-orange-internet) is not slave"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/misc/mikrotik-ipv6.md
source_anchor: ""
source_lines: [35, 144]
sha256: f3dd91dab9054f5cbdfe0eefe617f08f52d4c4174f1456d70ecb0f3f961fc117
---

# in/out-bridge-port matcher not possible when interface (vlan832-orange-internet) is not slave

```
/interface bridge
add admin-mac=AA:BB:CC:DD:EE:FF auto-mac=no comment=defconf name=bridge
/interface ethernet
set [ find default-name=sfp-sfpplus1 ] auto-negotiation=no speed=2.5Gbps
/interface vlan
add interface=sfp-sfpplus1 name=vlan832-orange-internet vlan-id=832
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/ip dhcp-client option
add code=60 name=vendorclass value="'sagem'"
add code=77 name=userclass value="'+FSVDSL_livebox.Internet.softathome.Livebox4'"
add code=90 name=authsend value=0xXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
/ip pool
add name=dhcp_pool1 ranges=192.168.88.3-192.168.88.254
/ip dhcp-server
add address-pool=dhcp_pool1 interface=bridge name=dhcpv4
/ipv6 dhcp-client option
add code=15 name=userclass value=0x002b46535644534c5f6c697665626f782e496e7465726e65742e736f66746174686f6d652e4c697665626f7834
add code=16 name=classidentifier value=0x0000040e0005736167656d
add code=11 name=auth value=0xXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
add code=6 name=request value=0x000b001100170018
/interface bridge filter
# in/out-bridge-port matcher not possible when interface (vlan832-orange-internet) is not slave
add action=set-priority chain=output dst-port=67 ip-protocol=udp mac-protocol=ip new-priority=6 out-interface=vlan832-orange-internet passthrough=yes
# in/out-bridge-port matcher not possible when interface (vlan832-orange-internet) is not slave
add action=set-priority chain=output dst-port=547 ip-protocol=udp mac-protocol=ipv6 new-priority=6 out-interface=vlan832-orange-internet passthrough=yes
/interface bridge port
add bridge=bridge comment=defconf interface=ether2
add bridge=bridge comment=defconf interface=ether3
add bridge=bridge comment=defconf interface=ether4
add bridge=bridge comment=defconf interface=ether5
add bridge=bridge comment=defconf interface=ether6
add bridge=bridge comment=defconf interface=ether7
add bridge=bridge comment=defconf interface=ether8
add bridge=bridge interface=ether1
/ip neighbor discovery-settings
set discover-interface-list=LAN
/ipv6 settings
set accept-router-advertisements=yes
/interface list member
add comment=defconf interface=bridge list=LAN
add interface=vlan832-orange-internet list=WAN
add interface=sfp-sfpplus1 list=WAN
/ip address
add address=192.168.88.1/24 comment=defconf interface=bridge network=192.168.88.0
/ip dhcp-client
add disabled=yes interface=ether1
add dhcp-options=vendorclass,userclass,authsend,clientid interface=vlan832-orange-internet
/ip dhcp-server network
add address=192.168.88.0/24 comment=defconf dns-server=192.168.88.1 gateway=192.168.88.1
/ip dns
set allow-remote-requests=yes
/ip dns static
add address=192.168.88.1 comment=defconf name=router.lan
/ip firewall filter
add action=drop chain=input comment="drop invalid input" connection-state=invalid
add action=drop chain=forward comment="drop invalid forward" connection-state=invalid
add action=drop chain=forward comment="drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=!LAN
add action=drop chain=forward comment="force LAN to use the router as DNS server" dst-port=53 in-interface-list=LAN log=yes protocol=udp
add action=accept chain=input comment="accept established,related input" connection-state=established,related
add action=accept chain=forward comment="accept established,related forward" connection-state=established,related
add action=accept chain=input comment="accept outgoing traffic input" in-interface-list=LAN
add action=accept chain=forward comment="accept outgoing traffic forward" in-interface-list=LAN
add action=accept chain=input comment="accept icmp input" protocol=icmp
add action=accept chain=input comment="accept dns input" dst-port=53 in-interface=vlan832-orange-internet protocol=udp
add action=accept chain=input comment="accept dhcp input" dst-port=67 in-interface=vlan832-orange-internet protocol=udp
add action=accept chain=input comment="accept ntp input" dst-port=123 in-interface=vlan832-orange-internet protocol=udp
add action=drop chain=input comment="Drop everything else input"
add action=drop chain=forward comment="Drop everything else forward"
/ip firewall nat
add action=masquerade chain=srcnat out-interface=vlan832-orange-internet src-address=192.168.88.0/24
/ip route
add disabled=yes distance=1 dst-address=0.0.0.0/0 gateway=192.168.1.1 pref-src="" routing-table=main scope=30 suppress-hw-offload=no target-scope=10
/ipv6 dhcp-client
add add-default-route=yes default-route-distance=2 dhcp-options=auth,classidentifier,userclass,request interface=vlan832-orange-internet pool-name=Orange \
    rapid-commit=no request=prefix
/ipv6 firewall address-list
add address=::/128 comment="defconf: unspecified address" list=bad_ipv6
add address=::1/128 comment="defconf: lo" list=bad_ipv6
add address=fec0::/10 comment="defconf: site-local" list=bad_ipv6
add address=::ffff:0.0.0.0/96 comment="defconf: ipv4-mapped" list=bad_ipv6
add address=::/96 comment="defconf: ipv4 compat" list=bad_ipv6
add address=100::/64 comment="defconf: discard only " list=bad_ipv6
add address=2001:db8::/32 comment="defconf: documentation" list=bad_ipv6
add address=2001:10::/28 comment="defconf: ORCHID" list=bad_ipv6
add address=3ffe::/16 comment="defconf: 6bone" list=bad_ipv6
/ipv6 firewall filter
add action=drop chain=input comment="drop invalid input" connection-state=invalid
add action=drop chain=forward comment="drop invalid forward" connection-state=invalid
add action=drop chain=forward comment="defconf: drop packets with bad src ipv6 forward" src-address-list=bad_ipv6
add action=drop chain=forward comment="defconf: drop packets with bad dst ipv6 forward" dst-address-list=bad_ipv6
add action=drop chain=forward comment="defconf: rfc4890 drop hop-limit=1" hop-limit=equal:1 protocol=icmpv6
add action=accept chain=input comment="accept established,related input" connection-state=established,related
add action=accept chain=forward comment="accept established,related forward" connection-state=established,related
add action=accept chain=input comment="accept outgoing traffic input" in-interface-list=LAN
add action=accept chain=forward comment="accept outgoing traffic forward" in-interface-list=LAN
add action=accept chain=input comment="accept ICMPv6 input" protocol=icmpv6
add action=accept chain=input comment="accept UDP traceroute input" port=33434-33534 protocol=udp
add action=accept chain=input comment="accept dhcp input" dst-port=546 in-interface=vlan832-orange-internet protocol=udp src-address=fe80::ba0:bab/128
add action=accept chain=input comment="accept DHCPv6-Client prefix delegation input" dst-port=546 protocol=udp src-address=fe80::/10
add action=drop chain=input comment="Drop everything else input"
add action=drop chain=forward comment="Drop everything else forward"
```

Ai-je oublié qqchse ?


Merci de pour votre aide..
