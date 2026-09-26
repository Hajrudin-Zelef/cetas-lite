---
id: collect-260926-mikrotik/mikrotik/help-with-ipsec-gre-tunnel-site-to-site-2
title: "These are the IPs of both ends of the IPsec tunnel."
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/ipsec/help-with-ipsec-gre-tunnel-site-to-site.md
source_anchor: ""
source_lines: [188, 321]
sha256: 9b025fa8c95068f4178caee0b939e36c838576441e2a0bb9b8f4b4744b3665c1
---

# These are the IPs of both ends of the IPsec tunnel.

```
/interface bridge
add name=ike2-gre-loopback-A
add name=ike2-gre-loopback-C
/interface ethernet
set [ find default-name=ether1 ] disable-running-check=no
set [ find default-name=ether2 ] disable-running-check=no
/interface gre
add local-address=10.200.2.1 name=gre-A remote-address=10.200.2.2
add local-address=10.200.1.1 name=gre-C remote-address=10.200.1.2
/interface wireguard
add listen-port=13231 mtu=1420 name=wireguard1
/interface list
add name=LAN
add name=WAN
/ip ipsec mode-config
add address=10.200.1.2 name=ike-C split-include=10.200.1.1/32 \
    system-dns=no
add address=10.200.2.2 name=ike-A split-include=10.200.2.1/32 \
    system-dns=no
/ip ipsec policy group
add name=ike2-gre
/ip ipsec profile
add dh-group=modp2048,modp1536,modp1024 enc-algorithm=aes-256,aes-192,aes-128 \
    hash-algorithm=sha256 name=ike2-gre
/ip ipsec peer
add exchange-mode=ike2 name=ike2-server passive=yes profile=ike2-gre \
    send-initial-contact=no
/ip ipsec proposal
add auth-algorithms=sha512,sha256 enc-algorithms=\
    aes-256-cbc,aes-256-ctr,aes-256-gcm,aes-192-ctr,aes-192-gcm,aes-128-gcm \
    lifetime=2h name=ike2-gre pfs-group=none
/ip neighbor discovery-settings
set discover-interface-list=none
/ip settings
set tcp-syncookies=yes
/interface list member
add interface=ether1 list=WAN
add interface=ether2 list=LAN
add interface=wireguard1 list=LAN
/ip address
add address=XXX interface=ether1 network=XXX
add address=192.168.5.1/24 interface=ether2 network=192.168.5.0
add address=192.168.105.1/24 interface=wireguard1 network=192.168.105.0
add address=10.200.1.1 comment=site-C interface=ike2-gre-loopback-C \
    network=10.200.1.1
add address=10.100.2.1/30 interface=gre-site-A network=10.100.2.0
add address=10.200.2.1 comment=site-A interface=ike2-gre-loopback-A \
    network=10.200.2.1
add address=10.100.1.1/30 interface=gre-site-C network=10.100.1.0
/ip cloud
set ddns-enabled=yes ddns-update-interval=30m
/ip dns
set allow-remote-requests=yes servers=8.8.8.8,1.1.1.1
/ip firewall address-list
add list=ddos-attacker
add list=ddos-target
/ip firewall filter
add action=accept chain=input comment=\
    "defconf: accept established,related,untracked" connection-state=\
    established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection-state=\
    invalid
add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
add action=accept chain=input comment=\
    "defconf: accept to local loopback (for CAPsMAN)" dst-address=127.0.0.1
add action=accept chain=input comment=\
    "Accept UDP of IPsec ports" dst-port=1701,500,4500 \
    in-interface-list=WAN protocol=udp
add action=accept chain=input comment="Accept protocol IPsec" protocol=\
    ipsec-esp
add action=drop chain=input comment="defconf: drop all not coming from LAN" \
    disabled=yes in-interface-list=!LAN
add action=jump chain=forward connection-state=new in-interface-list=WAN \
    jump-target=detect-ddos
add action=accept chain=forward comment="defconf: accept incoming IPsec" \
    ipsec-policy=in,ipsec
add action=accept chain=forward comment="defconf: accept outgoing IPsec" \
    ipsec-policy=out,ipsec
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" \
    connection-state=established,related hw-offload=yes
add action=accept chain=forward comment=\
    "defconf: accept established,related,untracked" connection-state=\
    established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" \
    connection-state=invalid
add action=drop chain=forward comment=\
    "defconf: drop all from WAN not DSTNATed" connection-nat-state=!dstnat \
    connection-state=new in-interface-list=WAN
add action=return chain=detect-ddos dst-limit=64,64,src-and-dst-addresses/10s
add action=add-dst-to-address-list address-list=ddos-target \
    address-list-timeout=10m chain=detect-ddos
add action=add-src-to-address-list address-list=ddos-attacker \
    address-list-timeout=10m chain=detect-ddos
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" \
    ipsec-policy=out,none out-interface-list=WAN
/ip firewall raw
add action=drop chain=prerouting dst-address-list=ddos-target \
    src-address-list=ddos-attacker
/ip ipsec identity
add auth-method=digital-signature certificate=CA.ike2.cert comment=site-C \
    generate-policy=port-strict match-by=certificate mode-config=\
    ike-gre-site-C peer=ike2-server policy-template-group=ike2-gre \
    remote-certificate=site_C.cert
add auth-method=digital-signature certificate=CA.ike2.tacf comment=site-A \
    generate-policy=port-strict match-by=certificate mode-config=\
    ike-gre-site-A peer=ike2-server policy-template-group=ike2-gre \
    remote-certificate=site_A.cert
/ip ipsec policy
add dst-address=10.200.0.0/16 group=ike2-gre proposal=ike2-gre src-address=\
    10.200.0.0/16 template=yes
/ip route
add gateway=XXX
# Route to site C
add disabled=no distance=1 dst-address=192.168.6.0/24 gateway=10.100.1.2 \
    routing-table=main scope=30 suppress-hw-offload=no target-scope=10
# Route to site A
add disabled=no distance=1 dst-address=192.168.7.0/24 gateway=10.100.2.2 \
    routing-table=main scope=30 suppress-hw-offload=no target-scope=10
/system identity
set name=site-B
/system ntp client
set enabled=yes
/system ntp client servers
add address=0.....pool.ntp.org
add address=1.....pool.ntp.org
add address=3.....pool.ntp.org
add address=4.....pool.ntp.org
add address=2.....pool.ntp.org
```

### Site C (client/initiator in IPsec/GRE connection)

