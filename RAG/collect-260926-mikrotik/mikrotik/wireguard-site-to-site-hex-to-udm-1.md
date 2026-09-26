---
id: collect-260926-mikrotik/mikrotik/wireguard-site-to-site-hex-to-udm-1
title: "UDM Pro Site C"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-site-to-site-hex-to-udm.md
source_anchor: ""
source_lines: [1, 109]
sha256: 5fc3a7e4eb9d386e3168627e0635bac8678554094b99000a4f8463af02018c18
---

# UDM Pro Site C

I really am trying to understand and learn – and (as always) appreciate the help!

With the code below, there appears to be active VPN connections between sites A, B and C, but it’s still not completely working.

Using the Ping tool at the Hex at SITE-B I can ping the UDM at SITE-C (192.168.0.1) and devices on the LAN behind the UDM.

Using the same Ping tool at the Hex at SITE-B I cannot ping the Hex at Site-A (192.168.2.1) or any devices behind the Hex.

Using the Ping tool at the Hex at SITE-A I can ping the UDM at SITE-C (192.168.0.1) and devices on the LAN behind the UDM.

Using the same Ping tool at the Hex at SITE-A I cannot ping the Hex at Site-A (192.168.88.1) or any devices behind the Hex.

From the UDM at SITE-C I cannot ping anything at SITE-A or SITE-B

SITE-C UDM

```
# UDM Pro Site C
Address = 10.10.20.1/32
SaveConfig = true
ListenPort = 51820
PrivateKey = WBj6xxxxx
[Peer]
# SITE A
PublicKey = xx27xxxxx
AllowedIPs = 10.10.10.0/24, 192.168.2.0/24
Endpoint = 22.22.22.22:51820
[Peer]
# SITE B
PublicKey = zoZtixxxxxx
AllowedIPs = 10.10.30.0/24, 192.168.88.0/24
Endpoint = 33.33.33.33:51820
```

SITE-B Hex

```
# Hex site B
/interface wireguard
add listen-port=51820 mtu=1420 name=wireguard1
/interface list member
add comment=defconf interface=bridge list=LAN
add comment=defconf interface=ether1 list=WAN
add interface=wireguard1 list=WAN
/interface wireguard peers
add allowed-address=192.168.2.0/24 comment=212 endpoint-address=\
    SITE-A.dyndns.org endpoint-port=51820 interface=wireguard1 public-key=\
    "xx27ccccc"
add allowed-address=192.168.0.0/24,192.168.5.0/24 comment=355 \
    endpoint-address=SITE-C.dyndns.org endpoint-port=51820 interface=\
    wireguard1 public-key="4HEOxxxxxxxx"
/ip address
add address=192.168.88.1/24 comment=defconf interface=bridge network=\
    192.168.88.0
add address=10.10.30.1/24 interface=wireguard1 network=10.10.30.0
/ip firewall address-list
add address=SITE-C.dyndns.org list=mtdale
add address=SITE-A.dyndns.org list=212
/ip firewall filter
add action=accept chain=input comment=\
    "defconf: accept established,related,untracked" connection-state=\
    established,related,untracked
add action=accept chain=input comment="allow incoming wireguard connections" \
    dst-port=51820 log=yes protocol=udp
add action=accept chain=input comment="Alow wireguard to router" \
    in-interface=wireguard1 log=yes
add action=accept chain=forward comment="Allow wireguard to subnet" \
    dst-address=192.168.88.0/24 in-interface=wireguard1 log=yes
add action=accept chain=forward in-interface=wireguard1 protocol=udp
add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
add action=drop chain=input comment="defconf: drop invalid" connection-state=\
    invalid
add action=accept chain=input comment=\
    "defconf: accept to local loopback (for CAPsMAN)" dst-address=127.0.0.1
add action=accept chain=input disabled=yes src-address-list=mtdale
add action=accept chain=input src-address-list=212
add action=drop chain=input comment="defconf: drop all not coming from LAN" \
    in-interface-list=!LAN
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" \
    connection-state=established,related hw-offload=yes
add action=accept chain=forward comment=\
    "defconf: accept established,related, untracked" connection-state=\
    established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" \
    connection-state=invalid log=yes
add action=accept chain=forward in-interface-list=LAN out-interface-list=WAN
add action=accept chain=forward comment="allow port forwarding" \
    connection-nat-state=dstnat log=yes
add action=drop chain=forward comment="Drop all else"
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" \
    ipsec-policy=out,none out-interface-list=WAN
add action=dst-nat chain=dstnat disabled=yes dst-port=9000,8080,554,1935,8035 \
    in-interface=wireguard1 log=yes protocol=tcp to-addresses=192.168.88.35
add action=dst-nat chain=dstnat dst-port=9000,8080,554,1935,8035 protocol=tcp \
    src-address-list=212 to-addresses=192.168.88.35
add action=dst-nat chain=dstnat comment=cam dst-port=8080,9000,554,1935,8035 \
    protocol=tcp src-address-list=mtdale to-addresses=192.168.88.35
/ip route
add disabled=no dst-address=192.168.2.0/24 gateway=wireguard1 routing-table=\
    main suppress-hw-offload=no
add disabled=no dst-address=192.168.0.0/24 gateway=wireguard1 routing-table=\
    main suppress-hw-offload=no
add disabled=no dst-address=192.168.1.0/24 gateway=wireguard1 routing-table=\
    main suppress-hw-offload=no
```

SITE-A Hex

