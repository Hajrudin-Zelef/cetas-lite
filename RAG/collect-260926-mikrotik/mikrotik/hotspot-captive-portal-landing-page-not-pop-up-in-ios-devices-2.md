---
id: collect-260926-mikrotik/mikrotik/hotspot-captive-portal-landing-page-not-pop-up-in-ios-devices-2
title: "hotspot-captive-portal-landing-page-not-pop-up-in-ios-devices"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/hotspot-captive-portal-landing-page-not-pop-up-in-ios-devices.md
source_anchor: ""
source_lines: [10, 121]
sha256: 353b0a8c63ed29f0f1ea2eb2274b9b2522bff5ef396ce95802afb13fa65b3cf0
---

# hotspot-captive-portal-landing-page-not-pop-up-in-ios-devices

```
/interface bridge
add ingress-filtering=no name=bridge-trunk protocol-mode=none vlan-filtering=yes
/interface ethernet
set [ find default-name=ether1 ] comment=WAN
/interface vlan
add interface=bridge-trunk name=vlan10_CP vlan-id=10
add interface=bridge-trunk name=vlan30_Private vlan-id=30
add interface=ether1 name=vlan126_B2B_Management vlan-id=126
add interface=ether1 name=vlan127_B2B_Internet vlan-id=127
/caps-man datapath
add bridge=bridge-trunk client-to-client-forwarding=no comment="Hotspot Datapath" local-forwarding=no name=vlan10 vlan-id=10 vlan-mode=use-tag
add bridge=bridge-trunk client-to-client-forwarding=no comment="Private Customer Datapath" local-forwarding=no name=vlan30 vlan-id=30 vlan-mode=use-tag
/caps-man configuration
add channel.band=2ghz-b/g/n .control-channel-width=20mhz .tx-power=30 country=**ELIDED** datapath=vlan10 distance=dynamic installation=any mode=ap name=hotspot-vlan10 rates.basic=1Mbps,2Mbps,5.5Mbps,11Mbps,6Mbps,9Mbps,12Mbps,18Mbps,24Mbps,36Mbps,48Mbps,54Mbps \
    .ht-basic-mcs=mcs-0,mcs-1,mcs-2,mcs-3,mcs-4,mcs-5,mcs-6,mcs-7,mcs-8,mcs-9,mcs-10,mcs-11,mcs-12,mcs-13,mcs-14,mcs-15,mcs-16,mcs-17,mcs-18,mcs-19,mcs-20,mcs-21,mcs-22,mcs-23 .ht-supported-mcs=\
    mcs-0,mcs-1,mcs-2,mcs-3,mcs-4,mcs-5,mcs-6,mcs-7,mcs-8,mcs-9,mcs-10,mcs-11,mcs-12,mcs-13,mcs-14,mcs-15,mcs-16,mcs-17,mcs-18,mcs-19,mcs-20,mcs-21,mcs-22,mcs-23 .supported=1Mbps,2Mbps,5.5Mbps,11Mbps,6Mbps,9Mbps,12Mbps,18Mbps,24Mbps,36Mbps,48Mbps,54Mbps \
    .vht-basic-mcs="" ssid=**ELIDED**
add channel.band=2ghz-b/g/n .control-channel-width=20mhz .tx-power=30 country=**ELIDED** datapath=vlan30 distance=dynamic installation=any mode=ap name=private-vlan30 rates.basic=1Mbps,2Mbps,5.5Mbps,11Mbps,6Mbps,9Mbps,12Mbps,18Mbps,24Mbps,36Mbps,48Mbps,54Mbps \
    .ht-basic-mcs=mcs-0,mcs-1,mcs-2,mcs-3,mcs-4,mcs-5,mcs-6,mcs-7,mcs-8,mcs-9,mcs-10,mcs-11,mcs-12,mcs-13,mcs-14,mcs-15,mcs-16,mcs-17,mcs-18,mcs-19,mcs-20,mcs-21,mcs-22,mcs-23 .ht-supported-mcs=\
    mcs-0,mcs-1,mcs-2,mcs-3,mcs-4,mcs-5,mcs-6,mcs-7,mcs-8,mcs-9,mcs-10,mcs-11,mcs-12,mcs-13,mcs-14,mcs-15,mcs-16,mcs-17,mcs-18,mcs-19,mcs-20,mcs-21,mcs-22,mcs-23 .supported=1Mbps,2Mbps,5.5Mbps,11Mbps,6Mbps,9Mbps,12Mbps,18Mbps,24Mbps,36Mbps,48Mbps,54Mbps \
    .vht-basic-mcs="" ssid=**ELIDED**
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/ip hotspot profile
add dns-name=**ELIDED** hotspot-address=192.168.150.1 http-cookie-lifetime=1d login-by=cookie,https name=hsprof1 ssl-certificate=**ELIDED**.crt use-radius=yes
/ip pool
add name=pool_portal ranges=192.168.150.5-192.168.151.254
add name=dhcp_pool1 ranges=192.168.20.2-192.168.20.254
add name=dhcp_pool2 ranges=192.168.160.2-192.168.160.254
/ip dhcp-server
add address-pool=pool_portal interface=vlan10_CP lease-time=2h name=dhcp-captive-portal
add address-pool=dhcp_pool1 interface=bridge-trunk lease-time=2h name=dhcp-management
add address-pool=dhcp_pool2 interface=vlan30_Private lease-time=2h name=dhcp-private-lan
/ip hotspot
add address-pool=pool_portal addresses-per-mac=1 disabled=no interface=vlan10_CP name=hotspot1 profile=hsprof1
/port
set 0 name=serial0
/queue type
set 0 pfifo-limit=5000
set 1 pfifo-limit=5000
set 9 pfifo-limit=1000
/caps-man manager
set ca-certificate=auto certificate=auto enabled=no
/caps-man provisioning
add action=create-dynamic-enabled master-configuration=hotspot-vlan10 slave-configurations=private-vlan30
/interface bridge port
add bridge=bridge-trunk interface=ether2
add bridge=bridge-trunk interface=ether3
add bridge=bridge-trunk interface=ether4
add bridge=bridge-trunk interface=ether5
add bridge=bridge-trunk interface=ether6
add bridge=bridge-trunk interface=ether7
add bridge=bridge-trunk interface=ether8
/ip neighbor discovery-settings
set discover-interface-list=LAN
/interface bridge vlan
add bridge=bridge-trunk tagged=ether2,ether3,ether4,ether5,ether6,ether7,ether8,bridge-trunk vlan-ids=10
add bridge=bridge-trunk tagged=bridge-trunk,ether2,ether3,ether4,ether5,ether6,ether7,ether8 vlan-ids=30
/interface list member
add interface=bridge-trunk list=LAN
add interface=ether1 list=WAN
/ip address
add address=192.168.88.2/24 interface=ether2 network=192.168.88.0
add address=102.216.175.200/23 interface=vlan127_B2B_Internet network=102.216.174.0
add address=192.168.150.1/23 interface=vlan10_CP network=192.168.150.0
add address=192.168.20.1/24 interface=bridge-trunk network=192.168.20.0
add address=192.168.160.1/24 interface=vlan30_Private network=192.168.160.0
/ip dhcp-client
add add-default-route=no interface=vlan126_B2B_Management
/ip dhcp-server network
add address=192.168.20.0/24 dns-server=10.100.4.2,10.100.4.10 gateway=192.168.20.1 netmask=24
add address=192.168.150.0/23 dns-server=10.100.4.2,10.100.4.10 gateway=192.168.150.1 netmask=23
add address=192.168.160.0/24 dns-server=10.100.4.2,10.100.4.10 gateway=192.168.160.1 netmask=24
/ip dns
set allow-remote-requests=yes cache-max-ttl=1d cache-size=4096KiB servers=10.100.4.2,10.100.4.10
/ip dns static
/ip firewall filter
add action=passthrough chain=unused-hs-chain comment="place hotspot rules here" disabled=yes
add action=fasttrack-connection chain=forward connection-state=established,related hw-offload=yes
add action=accept chain=forward comment="Allow already established & related connections" connection-state=established,related
add action=accept chain=forward comment="Allow access to router from known network"
add action=accept chain=forward comment="Allow acces to internal machines" dst-address-list=Internal_LAN
add action=drop chain=forward comment="Drop anything else"
add action=accept chain=input comment="Allow Established & related connections" connection-state=established,related
add action=accept chain=input comment="Allow access to router from known network" src-address=192.168.0.0/16
add action=accept chain=input comment="Allow remote ssh & WinBOX" dst-port=18291,32222,443 protocol=tcp src-address=10.100.0.0/16
add action=accept chain=input comment="Allow ping" protocol=icmp
add action=accept chain=input comment="Allow SNMP" dst-port=161 protocol=udp src-address=10.100.0.0/16
add action=accept chain=input comment="Allow pptp connections" disabled=yes dst-port=1723 protocol=tcp
add action=accept chain=input disabled=yes protocol=gre
add action=accept chain=input comment="Allow Web Acces" disabled=yes dst-port=80 protocol=tcp src-address=10.0.0.0/8
add action=drop chain=input comment="Drop anything else"
/ip firewall mangle
add action=mark-connection chain=prerouting comment="VOIP Traffic" connection-type=sip disabled=yes new-connection-mark=VOIP passthrough=yes
add action=mark-connection chain=prerouting comment="VOIP Traffic" disabled=yes dst-address-list=WM-VOIP dst-port=5060-5070,10000-20000 new-connection-mark=VOIP passthrough=yes protocol=udp
add action=set-priority chain=prerouting comment="VOIP Priority" connection-mark=VOIP disabled=yes new-priority=5 passthrough=no
add action=mark-connection chain=prerouting comment="Remaining Traffic" connection-mark=no-mark disabled=yes new-connection-mark=NO-VOIP passthrough=yes
add action=set-priority chain=prerouting comment="Remaining Traffic Priority" connection-mark=NO-VOIP disabled=yes new-priority=3 passthrough=no
/ip firewall nat
add action=passthrough chain=unused-hs-chain comment="place hotspot rules here" disabled=yes
add action=passthrough chain=unused-hs-chain comment="place hotspot rules here" disabled=yes
add action=masquerade chain=srcnat comment=Private-LAN out-interface=vlan126_B2B_Management src-address=192.168.160.0/24
add action=masquerade chain=srcnat out-interface=vlan127_B2B_Internet src-address=192.168.160.0/24
add action=masquerade chain=srcnat comment=Hotspot-LAN out-interface=vlan126_B2B_Management src-address=192.168.150.0/23
add action=masquerade chain=srcnat out-interface=vlan127_B2B_Internet src-address=192.168.150.0/23
/ip firewall service-port
set ftp disabled=yes
set tftp disabled=yes
set h323 disabled=yes
set pptp disabled=yes
/ip hotspot user
