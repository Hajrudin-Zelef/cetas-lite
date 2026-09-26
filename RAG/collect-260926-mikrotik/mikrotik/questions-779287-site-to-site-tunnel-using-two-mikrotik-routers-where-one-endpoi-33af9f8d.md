---
id: collect-260926-mikrotik/mikrotik/questions-779287-site-to-site-tunnel-using-two-mikrotik-routers-where-one-endpoi-33af9f8d
title: "questions-779287-site-to-site-tunnel-using-two-mikrotik-routers-where-one-endpoi-33af9f8d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vpn/questions-779287-site-to-site-tunnel-using-two-mikrotik-routers-where-one-endpoi-33af9f8d.md
source_anchor: ""
source_lines: [1, 33]
sha256: 25a840b304890e1d18ced7494e74f17e6e1c46bee98f9aab7a289ae1de5a4827
---

# questions-779287-site-to-site-tunnel-using-two-mikrotik-routers-where-one-endpoi-33af9f8d

I would like to interconnect two offices where one has a public static IP address (main office) and the second one is behind NAT (no public IP) because there is just an LTE modem.
I am able to create a one-way VPN connection from the LTE modem into the main office but is it possible to make the TCP communication between the two offices bi-directional? So that people from the main office can for example RDP to the branch office?
(I'm using two MikroTik Routerboards and a PPTP connection. I should be able to change to L2TP if needed.).
UPDATE:
I'm providing details on request:
Main office: LAN: 192.168.16.0/24
Public IP: MAIN_OFFICE_IP
Branch office LAN: 192.168.1.0/24
Public IP: [DHCP from ISP]
BRANCH OFFICE configuration:
two network interfaces
one PPTP client
absolutely basic Firewall and NAT  
/interface ethernet
set [ find default-name=ether1 ] name=ether1-gateway
set [ find default-name=ether2 ] arp=proxy-arp name=ether2-master-local
/interface pptp-client
add add-default-route=yes allow=pap,chap,mschap1,mschap2 connect-to=MAIN_OFFICE_IP default-route-distance=1 dial-on-demand=no disabled=no keepalive-timeout=60 max-mru=1450 max-mtu=1450 mrru=1600 name=\
MAIN_OFFICE_VPN password=******** profile=default-encryption user=MAIN_OFFICE_USER
/ip firewall filter
add chain=input comment="default configuration" protocol=icmp
add chain=input comment="default configuration" connection-state=established
add chain=input comment="default configuration" connection-state=related
add action=drop chain=input comment="default configuration" in-interface=ether1-gateway
add chain=forward comment="default configuration" connection-state=established
add chain=forward comment="default configuration" connection-state=related
add action=drop chain=forward comment="default configuration" connection-state=invalid
/ip firewall nat
add action=masquerade chain=srcnat comment="default configuration" out-interface=ether1-gateway to-addresses=0.0.0.0
add action=masquerade chain=srcnat out-interface=MAIN_OFFICE_VPN
add action=dst-nat chain=dstnat dst-address=BRANCH_IP dst-port=100 protocol=tcp
/ip route
add distance=1 dst-address=192.168.16.0/24 gateway=OFFICE_VPN routing-mark=“MAIN_OFFICE_VPN”
