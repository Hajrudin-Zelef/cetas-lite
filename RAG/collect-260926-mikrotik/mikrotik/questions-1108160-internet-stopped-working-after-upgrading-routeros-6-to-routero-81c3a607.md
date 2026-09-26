---
id: collect-260926-mikrotik/mikrotik/questions-1108160-internet-stopped-working-after-upgrading-routeros-6-to-routero-81c3a607
title: "questions-1108160-internet-stopped-working-after-upgrading-routeros-6-to-routero-81c3a607"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-1108160-internet-stopped-working-after-upgrading-routeros-6-to-routero-81c3a607.md
source_anchor: ""
source_lines: [1, 115]
sha256: ba41e80b78ef6e5121b1f70cf614a55b8cc089aea35007ad57f786b06f18334c
---

# questions-1108160-internet-stopped-working-after-upgrading-routeros-6-to-routero-81c3a607

I would like upgrade RouterOS 6 to RouterOS 7, everything went fine, all interfaces went up, but the Internet did not appear for users. As I understand it, somewhere need to change routes / something else for the new OS7? The config that works now is below. Where and what to change for RouterOS 7?
/interface bridge
add name=VLAN_99
add name=VLAN_100
add name=VLAN_200
/interface ethernet
set [ find default-name=ether1 ] comment=ISP1 name=ether1-wan
set [ find default-name=ether2 ] comment=Upl_SNR_sw1
set [ find default-name=ether6 ] comment=ISP2 name=ether6-wan
set [ find default-name=ether7 ] comment=Upl_SNR_sw2
/interface pppoe-client
add allow=pap,mschap1 disabled=no interface=ether1-wan max-mtu=1480 name=\
    pppoe_isp1 password=*** use-peer-dns=yes user=*
add allow=pap,mschap1 disabled=no interface=ether6-wan name=pppoe_isp2 \
    password=** use-peer-dns=yes user=****
/interface vlan
add interface=ether2 name=eth2-vlan99 vlan-id=99
add interface=ether2 name=eth2-vlan100 vlan-id=100
add interface=ether7 name=eth7-vlan99 vlan-id=99
add interface=ether7 name=eth7-vlan100 vlan-id=100
add interface=ether7 name=eth7-vlan200 vlan-id=200
/interface list
add name=WAN
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/ip pool
add name=POOL_99 ranges=192.168.3.30-192.168.3.254
add name=POOL_100 ranges=192.168.1.30-192.168.1.254
add name=POOL_200 ranges=192.168.2.30-192.168.2.254
/ip dhcp-server
add address-pool=POOL_99 disabled=no interface=VLAN_99 lease-time=2d name=\
    DHCP_99
add address-pool=POOL_100 disabled=no interface=VLAN_100 lease-time=2d name=\
    DHCP_100
add address-pool=POOL_200 disabled=no interface=VLAN_200 lease-time=2d name=\
    DHCP_200
/user group
set full policy="local,telnet,ssh,ftp,reboot,read,write,policy,test,winbox,pas\
    sword,web,sniff,sensitive,api,romon,dude,tikapp"
/interface bridge port
add bridge=VLAN_99 interface=eth2-vlan99
add bridge=VLAN_100 interface=eth2-vlan100
add bridge=VLAN_100 interface=eth7-vlan100
add bridge=VLAN_200 interface=eth7-vlan200
add bridge=VLAN_99 disabled=yes interface=ether3
add bridge=VLAN_99 interface=eth7-vlan99
/ip neighbor discovery-settings
set discover-interface-list=!dynamic
/interface list member
add interface=ether1-wan list=WAN
add interface=ether6-wan list=WAN
/ip address
add address=192.168.1.1/24 interface=VLAN_100 network=192.168.1.0
add address=192.168.2.1/24 interface=VLAN_200 network=192.168.2.0
add address=192.168.3.1/24 interface=VLAN_99 network=192.168.3.0
/ip dns
set allow-remote-requests=yes servers=8.8.8.8
/ip firewall address-list
add address=172.16.0.0/12 list=PRIVATE_NETWORKS
add address=192.168.0.0/16 list=PRIVATE_NETWORKS
/ip firewall filter
add action=accept chain=input comment=:::::::::Established/Related \
    connection-state=established,related
add action=accept chain=input comment=:::::::::GRE in-interface-list=WAN \
    protocol=gre
add action=accept chain=input comment=:::::::::L2TP dst-port=1701 \
    in-interface-list=WAN protocol=udp
add action=accept chain=input comment=:::::::::IPsec dst-port=500,4500 \
    in-interface-list=WAN protocol=udp
add action=accept chain=input comment=:::::::::IPsec in-interface-list=WAN \
    protocol=ipsec-esp
add action=accept chain=input comment=:::::::::Winbox/SSH dst-port=8291,22 \
    in-interface-list=WAN protocol=tcp src-address-list=CONSOLE
add action=accept chain=input comment=":::::::::Echo Request" icmp-options=\
    8:0-255 protocol=icmp
add action=accept chain=input comment=":::::::::Echo Reply" icmp-options=\
    0:0-255 protocol=icmp
add action=accept chain=input comment=":::::::::Destination Unreachable" \
    icmp-options=3:0-255 protocol=icmp
add action=accept chain=input comment=":::::::::Time Exceeded" icmp-options=\
    11:0-255 protocol=icmp
add action=accept chain=forward dst-address=192.168.1.0/24 src-address=\
    192.168.3.0/24
add action=accept chain=forward dst-address=192.168.3.0/24 src-address=\
    192.168.1.0/24
add action=accept chain=forward dst-address=192.168.2.0/24 src-address=\
    192.168.3.0/24
add action=drop chain=input comment=":::::::::Input Drop" in-interface-list=\
    WAN
add action=reject chain=forward comment=\
    ":::::::::Reject Direct Internet Access" dst-address-list=!EXCLUSION \
    out-interface-list=WAN reject-with=icmp-admin-prohibited \
    src-address-list=PRIVATE_NETWORKS
add action=accept chain=forward comment=:::::::::Established/Related \
    connection-state=established,related
add action=drop chain=forward comment=":::::::::Forward Drop" \
    connection-nat-state=!dstnat in-interface-list=WAN
/ip firewall mangle
add action=mark-routing chain=prerouting dst-address=!192.168.0.0/16 \
    new-routing-mark=ISP1 passthrough=yes src-address=192.168.3.0/24
add action=mark-routing chain=prerouting dst-address=!192.168.0.0/16 \
    new-routing-mark=ISP1 passthrough=yes src-address=192.168.1.0/24
add action=mark-routing chain=prerouting dst-address=!192.168.0.0/16 \
    new-routing-mark=ISP2 passthrough=yes src-address=192.168.2.0/24
/ip firewall nat
add action=masquerade chain=srcnat out-interface-list=WAN
add action=masquerade chain=srcnat dst-address=!192.168.0.0/16 out-interface=\
    pppoe_isp1 src-address=192.168.3.2-192.168.3.254
add action=masquerade chain=srcnat dst-address=!192.168.0.0/16 out-interface=\
    pppoe_isp1 src-address=192.168.1.2-192.168.3.254
add action=masquerade chain=srcnat dst-address=!192.168.0.0/16 out-interface=\
    pppoe_isp2 src-address=192.168.2.2-192.168.2.254
/ip route
add check-gateway=ping distance=1 gateway=pppoe_isp1 routing-mark=ISP1
add check-gateway=ping distance=1 gateway=pppoe_isp2 routing-mark=ISP2
