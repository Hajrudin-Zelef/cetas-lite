---
id: collect-260926-mikrotik/mikrotik/questions-831174-how-to-configure-a-mikrotik-hap-ac-lite-router-as-a-layer-2-swi-084d44a1
title: "mar/05/2017 00:49:46 by RouterOS 6.34.2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-831174-how-to-configure-a-mikrotik-hap-ac-lite-router-as-a-layer-2-swi-084d44a1.md
source_anchor: ""
source_lines: [1, 91]
sha256: fc1155ca5327d6a872d6fa60a186e020f7b6c73101cc889fdd62d789b189d859
---

# mar/05/2017 00:49:46 by RouterOS 6.34.2

Here is an overview of the network topology:
Both the Fritz!Box 7340 and MikroTik hAP ac lite router currently act as DHCP servers, effectively splitting the network in two LANs.
The MikroTik router was also configured as a bridge for the 192.168.88.xxx LAN.
In order to avoid the trouble of double NAT, I would like to reconfigure the MikroTik hAP ac lite as a Layer 2 switch. There are advantages of doing so over double NATing.
All DHCP and NAT should be handled exclusively by the Fritz!Box. Furthermore, the MikroTik hAP ac lite should have a static IP address in the same local IP address range so it can still be managed over its web interface.
The Fritz!Box configuration is not an issue. However, I am an utter newbie managing MikroTik's RouterOS and I found no instructions on the web concerning this specific problem. Hence, my question:
What are the (detailed) steps to configure a MikroTik hAP ac lite as a Layer 2 switch, removing its DHCP server whilst keeping a static IP address for management?
By the way, I am on GNU/Linux, so using Winbox is not an option for me.
MikroTik support pages
On the MikroTik support pages, I came across an article about how to configure a DHCP relay. However, it is not clear to me whether a DHCP relay is the same thing as a Layer 2 switch or just a part of it?
Crosslink to the MikroTik forum
http://forum.mikrotik.com/viewtopic.php?f=2&t=117434
Export
# mar/05/2017 00:49:46 by RouterOS 6.34.2
# software id = XXXX-XXXX
#
/interface bridge
add admin-mac=6C:3B:::: auto-mac=no comment=defconf name=bridge
/interface ethernet
set [ find default-name=ether2 ] name=ether2-master
set [ find default-name=ether3 ] master-port=ether2-master
set [ find default-name=ether4 ] master-port=ether2-master
set [ find default-name=ether5 ] master-port=ether2-master
/ip neighbor discovery
set ether1 discover=no
set bridge comment=defconf
/interface wireless security-profiles
add authentication-types=wpa2-psk eap-methods="" management-protection=allowed \
    mode=dynamic-keys name=WLAN1 supplicant-identity="" wpa2-pre-shared-key=\
    *****************
add authentication-types=wpa2-psk eap-methods="" management-protection=allowed \
    mode=dynamic-keys name=WLAN2 supplicant-identity="" wpa2-pre-shared-key=\
    *****************
/interface wireless
set [ find default-name=wlan1 ] band=2ghz-b/g/n channel-width=20/40mhz-Ce \
    country=belgium disabled=no distance=indoors frequency=auto mode=ap-bridge \
    security-profile=WLAN1 ssid=WLAN1 wireless-protocol=802.11
set [ find default-name=wlan2 ] band=5ghz-a/n channel-width=20/40mhz-Ce \
    country=belgium disabled=no distance=indoors frequency=auto mode=\
    station-pseudobridge security-profile=WLAN2 ssid=WLAN2 \
    wireless-protocol=802.11
/ip pool
add name=default-dhcp ranges=192.168.88.10-192.168.88.254
/ip dhcp-server
add address-pool=default-dhcp disabled=no interface=bridge name=defconf
/interface bridge port
add bridge=bridge comment=defconf interface=ether2-master
add bridge=bridge comment=defconf interface=wlan1
add bridge=bridge comment=defconf interface=wlan2
/ip address
add address=192.168.88.1/24 comment=defconf interface=bridge network=\
    192.168.88.0
/ip dhcp-client
add comment=defconf dhcp-options=hostname,clientid interface=ether1
add default-route-distance=0 dhcp-options=hostname,clientid disabled=no \
    interface=wlan2
/ip dhcp-server network
add address=192.168.88.0/24 comment=defconf gateway=192.168.88.1
/ip dns
set allow-remote-requests=yes
/ip dns static
add address=192.168.88.1 name=router
/ip firewall filter
add chain=input comment="defconf: accept ICMP" protocol=icmp
add chain=input comment="defconf: accept establieshed,related" \
    connection-state=established,related
add action=drop chain=input comment="defconf: drop all from WAN" in-interface=\
    ether1
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" \
    connection-state=established,related
add chain=forward comment="defconf: accept established,related" \
    connection-state=established,related
add action=drop chain=forward comment="defconf: drop invalid" connection-state=\
    invalid
add action=drop chain=forward comment=\
    "defconf:  drop all from WAN not DSTNATed" connection-nat-state=!dstnat \
    connection-state=new in-interface=ether1
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" disabled=yes \
    out-interface=ether1
add action=masquerade chain=srcnat out-interface=bridge
/system clock
set time-zone-name=Europe/Brussels
/system routerboard settings
set cpu-frequency=650MHz protected-routerboot=disabled
/tool mac-server
set [ find default=yes ] disabled=yes
add interface=bridge
/tool mac-server mac-winbox
set [ find default=yes ] disabled=yes
add interface=bridge
