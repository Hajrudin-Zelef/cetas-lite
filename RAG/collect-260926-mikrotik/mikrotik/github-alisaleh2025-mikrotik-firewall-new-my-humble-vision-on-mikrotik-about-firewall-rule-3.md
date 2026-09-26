---
id: collect-260926-mikrotik/mikrotik/github-alisaleh2025-mikrotik-firewall-new-my-humble-vision-on-mikrotik-about-firewall-rule-3
title: "github-alisaleh2025-mikrotik-firewall-new-my-humble-vision-on-mikrotik-about-firewall-rules"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/github-alisaleh2025-mikrotik-firewall-new-my-humble-vision-on-mikrotik-about-firewall-rules.md
source_anchor: ""
source_lines: [120, 173]
sha256: 4ec312e9b309b002a8da695ecbfcbc2fa19f229b82fda564ca06675a2f7ab65f
---

# github-alisaleh2025-mikrotik-firewall-new-my-humble-vision-on-mikrotik-about-firewall-rules

    add chain=Detect-Portscanners action=passthrough comment="Port scanner detect chain" action=passthrough
    add chain=Detect-Portscanners action=add-src-to-address-list address-list=Banned address-list-timeout=none-dynamic protocol=tcp psd=21,3s,3,1
```
```
    /ip firewall filter
    add chain=Drop-chain action=passthrough comment="Main drop chain"
    add chain=Drop-chain action=drop
```
```
    /ip firewall raw
    add chain=prerouting action=passthrough comment="Drop from prerouting"
    add chain=prerouting action=drop src-address-list=Banned
    add chain=prerouting action=drop src-address-list=ddos-attackers dst-address-list=ddos-targets
```
```
    /ip firewall mangle
    add action=mark-connection chain=input comment="mark ipsec connections to exclude them from fasttrack" ipsec-policy=in,ipsec new-connection-mark=ipsec
    add action=mark-connection chain=output comment="mark ipsec connections to exclude them from fasttrack" ipsec-policy=out,ipsec new-connection-mark=ipsec 
    add action=mark-connection chain=forward comment="mark ipsec connections to exclude them from fasttrack" ipsec-policy=out,ipsec new-connection-mark=ipsec 
    add action=mark-connection chain=forward comment="mark ipsec connections to exclude them from fasttrack" ipsec-policy=in,ipsec new-connection-mark=ipsec
```
```
    /ip firewall nat
    add chain=srcnat action=accept ipsec-policy=out,ipsec
    add chain=dstnat action=accept ipsec-policy=in,ipsec
    add chain=srcnat action=masquerade out-interface-list=Outside
```
```
    /ip firewall address-list
    add list=ddos-attackers
    add list=ddos-targets
    add list=Banned
    /ip firewall address-list
    add address=0.0.0.0/8 comment="RFC6890" list=not_in_internet
    add address=172.16.0.0/12 comment="RFC6890" list=not_in_internet
    add address=192.168.0.0/16 comment="RFC6890" list=not_in_internet
    add address=10.0.0.0/8 comment="RFC6890" list=not_in_internet
    add address=169.254.0.0/16 comment="RFC6890" list=not_in_internet
    add address=127.0.0.0/8 comment="RFC6890" list=not_in_internet
    add address=224.0.0.0/4 comment="Multicast" list=not_in_internet
    add address=198.18.0.0/15 comment="RFC6890" list=not_in_internet
    add address=192.0.0.0/24 comment="RFC6890" list=not_in_internet
    add address=192.0.2.0/24 comment="RFC6890" list=not_in_internet
    add address=198.51.100.0/24 comment="RFC6890" list=not_in_internet
    add address=203.0.113.0/24 comment="RFC6890" list=not_in_internet
    add address=100.64.0.0/10 comment="RFC6890" list=not_in_internet
    add address=240.0.0.0/4 comment="RFC6890" list=not_in_internet
    add address=192.88.99.0/24 comment="RFC 3068" list=not_in_internet
    /ip firewall address-list
    add address=0.0.0.0/8 comment="RFC6890" list=no_forward_ipv4
    add address=169.254.0.0/16 comment="RFC6890" list=no_forward_ipv4
    add address=224.0.0.0/4 comment="multicast" list=no_forward_ipv4
    add address=255.255.255.255/32 comment="RFC6890" list=no_forward_ipv4
```
