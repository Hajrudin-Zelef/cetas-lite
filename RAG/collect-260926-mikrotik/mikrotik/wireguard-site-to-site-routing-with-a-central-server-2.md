---
id: collect-260926-mikrotik/mikrotik/wireguard-site-to-site-routing-with-a-central-server-2
title: "2023-11-08 21:11:31 by RouterOS 7.11.2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-site-to-site-routing-with-a-central-server.md
source_anchor: ""
source_lines: [131, 297]
sha256: 11c1300014dda02be9aef37d7653b8022c4870ca8612bba0c140bc564d3593ea
---

# 2023-11-08 21:11:31 by RouterOS 7.11.2

```
/interface bridge
add name=allinall-bridge vlan-filtering=yes
/interface ethernet
set [ find default-name=sfp-sfpplus1 ] disabled=yes
set [ find default-name=sfp-sfpplus2 ] disabled=yes
set [ find default-name=sfp-sfpplus3 ] disabled=yes
set [ find default-name=sfp-sfpplus4 ] disabled=yes
/interface wireguard
add listen-port=13231 mtu=1420 name=wireguard1
/interface vlan
add interface=allinall-bridge name=unfriendly-vlan vlan-id=99
add interface=allinall-bridge name=backbone-vlan vlan-id=9
add interface=allinall-bridge name=friendly-vlan vlan-id=19
add interface=allinall-bridge name=handling-vlan vlan-id=4
/caps-man configuration
add country="czech republic" datapath.bridge=allinall-bridge \
    .client-to-client-forwarding=no .local-forwarding=yes .vlan-id=99 \
    .vlan-mode=use-tag hide-ssid=no mode=ap name=unfriendly-wlan \
    security.authentication-types=wpa2-psk ssid=Wifi1
add country="czech republic" datapath.bridge=allinall-bridge \
    .client-to-client-forwarding=yes .local-forwarding=yes .vlan-id=19 \
    .vlan-mode=use-tag hide-ssid=yes name=friendly-wlan \
    security.authentication-types=wpa2-psk ssid=Wifi2
/caps-man datapath
add bridge=allinall-bridge client-to-client-forwarding=no local-forwarding=no \
    name=unfriendly-datapath vlan-id=99 vlan-mode=use-tag
/interface list
add comment="all unfriendly interfaces" name=unfriendly-interface-list
/interface lte apn
set [ find default=yes ] ip-type=ipv4 use-network-apn=no
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/ip pool
add name=unfriendly-pool ranges=192.168.98.1-192.168.99.254
add name=backbone-pool ranges=192.168.9.1-192.168.9.254
add name=friendly-pool ranges=192.168.18.1-192.168.19.254
add name=handling-pool ranges=192.168.4.1-192.168.4.13
/ip dhcp-server
add address-pool=unfriendly-pool interface=unfriendly-vlan lease-time=10m name=\
    unfriendly-dhcp
add address-pool=backbone-pool allow-dual-stack-queue=no insert-queue-before=\
    bottom interface=backbone-vlan lease-time=10m name=backbone-dhcp
add address-pool=friendly-pool interface=friendly-vlan lease-time=10m name=\
    friendly-dhcp
add address-pool=handling-pool interface=handling-vlan lease-time=10m name=\
    handling-dhcp
/port
set 0 name=serial0
/routing bgp template
set default disabled=no output.network=bgp-networks
/routing ospf instance
add disabled=no name=default-v2
/routing ospf area
add disabled=yes instance=default-v2 name=backbone-v2
/caps-man manager
set enabled=yes
/caps-man manager interface
add interface=backbone-vlan
/caps-man provisioning
add action=create-dynamic-enabled master-configuration=friendly-wlan \
    name-format=identity slave-configurations=unfriendly-wlan
/interface bridge port
add bridge=allinall-bridge interface=ether3 pvid=9
add bridge=allinall-bridge interface=ether4 pvid=9
add bridge=allinall-bridge interface=ether5 pvid=9
add bridge=allinall-bridge interface=ether6 pvid=9
add bridge=allinall-bridge interface=ether7 pvid=9
add bridge=allinall-bridge interface=ether8 pvid=9
add bridge=allinall-bridge interface=ether9 pvid=19
add bridge=allinall-bridge interface=ether10 pvid=19
add bridge=allinall-bridge interface=ether11 pvid=19
add bridge=allinall-bridge interface=ether12 pvid=19
add bridge=allinall-bridge interface=ether13 pvid=19
add bridge=allinall-bridge interface=ether14 pvid=19
add bridge=allinall-bridge interface=ether15 pvid=19
add bridge=allinall-bridge interface=ether16 pvid=19
add bridge=allinall-bridge interface=ether17 pvid=99
add bridge=allinall-bridge interface=ether18 pvid=99
add bridge=allinall-bridge interface=ether19 pvid=99
add bridge=allinall-bridge interface=ether20 pvid=99
add bridge=allinall-bridge interface=ether21 pvid=99
add bridge=allinall-bridge interface=ether22 pvid=99
add bridge=allinall-bridge interface=ether23 pvid=99
add bridge=allinall-bridge interface=ether24 pvid=99
add bridge=allinall-bridge interface=ether1 pvid=4
/ip settings
set max-neighbor-entries=8192
/ipv6 settings
set disable-ipv6=yes
/interface bridge vlan
add bridge=allinall-bridge tagged=allinall-bridge untagged=\
    ether3,ether4,ether5,ether6,ether7,ether8 vlan-ids=9
add bridge=allinall-bridge tagged=\
    allinall-bridge,ether3,ether4,ether5,ether6,ether7,ether8 untagged=\
    ether9,ether10,ether11,ether12,ether13,ether14,ether15,ether16 vlan-ids=\
    19
add bridge=allinall-bridge tagged=\
    allinall-bridge,ether3,ether4,ether5,ether6,ether7,ether8 untagged=\
    ether17,ether18,ether19,ether20,ether21,ether22,ether23,ether24 vlan-ids=\
    99
add bridge=allinall-bridge tagged=allinall-bridge untagged=ether1 vlan-ids=4
/interface ethernet switch rule
add comment="unfriendly packets must are firewalled" ports=\
    ether3,ether4,ether5,ether6,ether7,ether8 redirect-to-cpu=yes switch=\
    switch1 vlan-header=present vlan-id=99
/interface ovpn-server server
set auth=sha1,md5
/interface wireguard peers
add allowed-address=192.168.224.0/24,192.168.188.0/24,192.168.184.0/24 \
    endpoint-address=[PUBLIC_IP] endpoint-port=13231 interface=wireguard1 \
    persistent-keepalive=5s public-key="..."
/ip address
add address=192.168.96.1/22 interface=unfriendly-vlan network=192.168.96.0
add address=192.168.8.1/23 interface=backbone-vlan network=192.168.8.0
add address=192.168.16.1/22 interface=friendly-vlan network=192.168.16.0
add address=192.168.4.14/28 interface=handling-vlan network=192.168.4.0
add address=192.168.224.2/24 interface=wireguard1 network=192.168.224.0
/ip cloud
set update-time=no
/ip dhcp-client
add interface=ether2
/ip dhcp-server config
set store-leases-disk=15m
/ip dhcp-server network
add address=192.168.4.0/28 dns-server=192.168.4.14 gateway=192.168.4.14 \
    ntp-server=216.239.35.4,195.113.144.201
add address=192.168.8.0/23 dns-server=192.168.8.1 gateway=192.168.8.1 \
    ntp-server=216.239.35.4,195.113.144.201
add address=192.168.16.0/22 dns-server=192.168.16.1 gateway=192.168.16.1 \
    ntp-server=216.239.35.4,195.113.144.201
add address=192.168.96.0/22 dns-server=192.168.96.1 gateway=192.168.96.1 \
    ntp-server=216.239.35.4,195.113.144.201
/ip dns
set allow-remote-requests=yes servers=\
    192.168.4.14,192.168.8.1,192.168.16.1,192.168.96.1
/ip firewall address-list
add address=0.0.0.0/8 comment=RFC6890 list=not_in_internet
add address=172.16.0.0/12 comment=RFC6890 list=not_in_internet
add address=192.168.0.0/16 comment=RFC6890 list=not_in_internet
add address=10.0.0.0/8 comment=RFC6890 list=not_in_internet
add address=169.254.0.0/16 comment=RFC6890 list=not_in_internet
add address=127.0.0.0/8 comment=RFC6890 list=not_in_internet
add address=224.0.0.0/4 comment=Multicast list=not_in_internet
add address=198.18.0.0/15 comment=RFC6890 list=not_in_internet
add address=192.0.0.0/24 comment=RFC6890 list=not_in_internet
add address=192.0.2.0/24 comment=RFC6890 list=not_in_internet
add address=198.51.100.0/24 comment=RFC6890 list=not_in_internet
add address=203.0.113.0/24 comment=RFC6890 list=not_in_internet
add address=100.64.0.0/10 comment=RFC6890 list=not_in_internet
add address=240.0.0.0/4 comment=RFC6890 list=not_in_internet
add address=192.88.99.0/24 comment=RFC3068 list=not_in_internet
/ip firewall filter
add action=fasttrack-connection chain=forward connection-state=\
    established,related hw-offload=yes
add action=accept chain=forward comment="disable forward chain [TEMPORARY]"
add action=accept chain=forward comment="accept established,related" \
    connection-state=established,related
add action=drop chain=forward comment="drop invalid" connection-state=invalid
add action=drop chain=forward comment="incoming not NATed" \
    connection-nat-state=!dstnat connection-state=new in-interface=ether2 \
    log=yes log-prefix=!nat
add action=drop chain=forward comment="incoming without public IP" \
    in-interface=ether2 log=yes log-prefix=!public src-address-list=\
    not_in_internet
add action=jump chain=forward comment="separate chain for ICMP" jump-target=\
    icmp protocol=icmp
