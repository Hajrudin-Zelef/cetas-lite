---
id: collect-260926-mikrotik/mikrotik/wireguard-roadwarior-plus-vlan-configuration
title: "wireguard-roadwarior-plus-vlan-configuration"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-roadwarior-plus-vlan-configuration.md
source_anchor: ""
source_lines: [1, 70]
sha256: afc7e216f90db111e7c701d3f506d12623bc669a4875b24a549de63782916e68
---

# wireguard-roadwarior-plus-vlan-configuration

(1) You define 4 VLANS to BR1  but only have 3 pools?   3 DHCP Servers?  3 DHCP-Server Networks?

This leads me to believe you dont really have a BASE VLAN…

Thus I will assume you actually  have a trusted VLAN on the  BLUE vlan and I get rid of 99

(2) All bridge ports are trunk ports, I would add ingress-filtering=yes to all the /interface bridge ports.

(3) Since there is no difference between how any VLAN-ID is assigned to bridge ports/vlans, your  /interface bridge vlan settings can be simplified.

*/interface bridge vlan
add bridge=BR1 tagged=BR1,ether2,ether3,ether4,ether5 **vlan-ids=20,30,40***

(4) Why do you have keep alive settings on the ROUTER,  its the  client devices which need persistent-keep-alive, YOU can remove them on the peers settings on the router.

(5) Mac server  alone, is not a secure access method, set to **NONE.**

*/tool mac-server
set allowed-interface-list=**BASE***

(6)  What is the purpose of these two masquerade rules??

add action=masquerade chain=srcnat src-address=10.1.1.0/24 ??

add action=masquerade chain=srcnat src-address=10.1.8.0/24 ??

(7)  Firewall rules are out of order  - fixed AND lets think about this from a consistent standpoint.  Why bother have a management segment if you let every tom dick and harry full access to the router??   I suppose the wireguard user is you as an admin, THus do this…

- add WG interface to BASE interface list
- add WG interface to VLAN interface list
- add blue-vlan trusted to the BASE interface list
- create firewall address list:  list=Authorized  add AdminDesktop_IP, AdminLaptop_IP, AdminIPad/iphone, RemoteAdminWG-IP

```
/ip firewall filter
add action=accept chain=input comment="Allow Estab & Related & untracked" \
    connection-state=established,related,untracked
add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
add action=accept chain=input comment="wg handshake" dst-port=13231 protocol=udp
add action=accept chain=input comment="Allow Admin" in-interface-list=BASE src-address=list=Authorized
add action=accept chain=input comment="User services" dst-port=53,123 protocol=udp in-interface-list=VLAN
add action=accept chain=input comment="User services" dst-port=53 protocol=tcp in-interface-list=VLAN
add action=drop chain=input comment=Drop
add action=fasttrack-connection chain=forward comment=Fasttrack \
    connection-state=established,related hw-offload=yes
add action=accept chain=forward comment="Allow Estab & Related & Untracked" \
    connection-state=established,related,untracked
add action=accept chain=forward comment="Internet traffic" in-interface-list=VLAN out-interface-list=WAN
add action=acccept chain=forward comment="admin to vlans"  in-interface-list=BASE out-interface-list=VLAN src-address-list=Authorized
add action=drop chain=forward comment=Drop
```

…

(8) This statement is completely bogus.

***What I have noticed as a bonus part of VLAN router configuration, now we having no internet going trought the WG RW Client which is configured :
AllowedIPs = 0.0.0.0/0, ::/0***

First, if there is no intention to provide internet to wireguard clients when they hit the router,  then they should NOT use 0.0.0.0/0 as their allowed IP settings.

It should be  10.10.1.8.0/24,SUBNETA,SUBNETB     ( depending if there are any subnets  they need to reach ).

Secondly if there is no intention to provide internet to wireguard clients, then dont assign them to the VLAN interface as I did above for  interface list members, and then they wont be included in the VLAN to WAN rule and thus not get internet.

Thirdly lets say you did have one wireguard client that needs internet and two that dont, then just create another firewall rule

*add chain=forward action=accept  in-interface=WG out-interface-list=WAN  src=address=wgIP.*

All to say is that you need to know what the requirements are and articulate them prior to making the config and/or asking advice.
