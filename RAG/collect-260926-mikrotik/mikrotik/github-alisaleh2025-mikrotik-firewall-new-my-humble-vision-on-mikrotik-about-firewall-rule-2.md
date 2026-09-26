---
id: collect-260926-mikrotik/mikrotik/github-alisaleh2025-mikrotik-firewall-new-my-humble-vision-on-mikrotik-about-firewall-rule-2
title: "github-alisaleh2025-mikrotik-firewall-new-my-humble-vision-on-mikrotik-about-firewall-rules"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/github-alisaleh2025-mikrotik-firewall-new-my-humble-vision-on-mikrotik-about-firewall-rules.md
source_anchor: ""
source_lines: [3, 119]
sha256: 3f0084e648be913ce105541f43a4bb85f81b34435d91a1a0ab9efd6eba0ac588
---

# github-alisaleh2025-mikrotik-firewall-new-my-humble-vision-on-mikrotik-about-firewall-rules

```
  /ip firewall filter
  add chain=input action=passthrough comment="Input chain to router"
  add chain=input action=jump jump-target=Icmp-Chain protocol=icmp
  add chain=input action=jump jump-target=Protected-chain
  add chain=input action=jump jump-target=Input-Outside in-interface-list=Outside
  add chain=input action=jump jump-target=Input-Inside in-interface-list=Inside
  add chain=input action=jump jump-target=Drop-chain in-interface-list=!Outside disabled=yes
  add chain=input action=jump jump-target=Drop-chain in-interface-list=!Inside disabled=yes
```
```
    /ip firewall filter
    add chain=Input-Outside action=passthrough comment="Input chain to router from outside"
    add chain=Input-Outside action=accept connection-state=established
    add chain=Input-Outside action=accept connection-state=related
    add chain=Input-Outside action=accept connection-state=untracked
    add chain=Input-Outside action=jump jump-target=Drop-chain connection-state=invalid
    add chain=Input-Outside action=accept ipsec-policy=in,ipsec action=accept
    add chain=Input-Outside action=jump jump-target=Input-Outside-Allow
    add chain=Input-Outside action=jump jump-target=Drop-chain
    /ip firewall filter
    add chain=Input-Outside-Allow  action=passthrough comment="Allow input chain to router from outside"
    add chain=Input-Outside-Allow  action=jump jump-target=BruteForce protocol=tcp dst-port=22 connection-nat-state=dstnat
    add chain=Input-Outside-Allow  action=jump jump-target=BruteForce protocol=tcp dst-port=8291
```
```
    /ip firewall filter
    add chain=Input-Inside action=passthrough comment="Input chain to router from inside" 
    add chain=Input-Inside action=accept connection-state=established 
    add chain=Input-Inside action=accept connection-state=related 
    add chain=Input-Inside action=accept connection-state=untracked 
    add chain=Input-Inside action=jump jump-target=Drop-chain connection-state=invalid
    add chain=Input-Inside action=jump jump-target=Input-Inside-All
    add chain=Input-Inside action=jump jump-target=Drop-chain
    /ip firewall filter
    add chain=Input-Inside-All action=passthrough comment="Input chain from inside to router"
    add chain=Input-Inside-All action=accept protocol=udp dst-port=53 
    add chain=Input-Inside-All action=accept protocol=tcp dst-port=53
    add chain=Input-Inside-All action=accept src-address-list=MNGMNT protocol=tcp dst-port=22
    add chain=Input-Inside-All action=accept src-address-list=MNGMNT protocol=tcp dst-port=8291
    add chain=Input-Inside-All action=jump jump-target=Drop-chain
```
```
    /ip firewall filter
    add chain=forward action=passthrough comment="Forward chain"
    add chain=forward action=fasttrack-connection connection-state=established connection-mark=!ipsec 
    add chain=forward action=fasttrack-connection connection-state=related connection-mark=!ipsec 
    add chain=forward action=accept connection-state=established
    add chain=forward action=accept connection-state=related
    add chain=forward action=accept connection-state=untracked
    add chain=forward action=jump jump-target=Drop-chain connection-state=invalid
    add action=jump chain=forward src-address-list=no_forward_ipv4 jump-target=Drop-Chain
    add action=ju8mp chain=forward dst-address-list=no_forward_ipv4 jump-target=Drop-Chain
    add chain=forward action=jump jump-target=Protected-chain
    add chain=forward in-interface-list=Outside action=jump jump-target=Forward-Outside
    add chain=forward in-interface-list=Inside action=jump jump-target=Forward-Inside
```
```
    /ip firewall filter
    add chain=Forward-Outside action=passthrough comment="Forward chain from outside to inside"
    add chain=Forward-Outside action=accept ipsec-policy=in,ipsec action=accept
    add chain=Forward-Outside action=jump jump-target=Drop-chain connection-nat-state=!dstnat connection-state=new
    add chain=Forward-Outside action=jump jump-target=Drop-chain in-interface-list=Outside out-interface-list=!Inside src-address-list=not_in_internet
    add chain=Forward-Outside connection-nat-state=dstnat action=accept
    add chain=Forward-Outside action=jump jump-target=Drop-chain
```
```
    /ip firewall filter
    add chain=Forward-Inside action=passthrough comment="Forward chain from inside to outside"
    add chain=Forward-Inside action=jump jump-target=Drop-chain in-interface-list=Inside out-interface-list=!Inside dst-address-list=not_in_internet
    add chain=Forward-Inside action=accept in-interface-list=Inside out-interface-list=Outside
    add chain=Forward-Inside action=accept in-interface-list=Inside out-interface-list=Inside
```
```
    /ip firewall filter
    add chain=Icmp-Chain action=passthrough comment="ICMP chain" 
    add chain=Icmp-Chain action=accept protocol=icmp icmp-options=0:0 comment="echo reply"
    add chain=Icmp-Chain action=accept protocol=icmp icmp-options=3:0 comment="net unreachable"
    add chain=Icmp-Chain action=accept protocol=icmp icmp-options=3:1 comment="host unreachable"
    add chain=Icmp-Chain action=accept protocol=icmp icmp-options=3:4 comment="host unreachable fragmentation required"
    add chain=Icmp-Chain action=accept protocol=icmp icmp-options=8:0 comment="allow echo request"
    add chain=Icmp-Chain action=accept protocol=icmp icmp-options=11:0 comment="allow time exceed"
    add chain=Icmp-Chain action=accept protocol=icmp icmp-options=12:0 comment="allow parameter bad"
    add chain=Icmp-Chain action=jump jump-target=Drop-chain
```
```
    /ip firewall filter
    add chain=BruteForce action=passthrough comment="BruteForce chain to router"
    add chain=BruteForce action=add-src-to-address-list address-list=Banned src-address-list=Stage_3 address-list-timeout=1d  comment=Blacklist 
    add chain=BruteForce action=add-src-to-address-list address-list=Stage_3 src-address-list=Stage_2,!secured address-list-timeout=1h  comment="Third attempt" 
    add chain=BruteForce action=add-src-to-address-list address-list=Stage_2 src-address-list=Stage_1 address-list-timeout=15m  comment="Second attempt" 
    add chain=BruteForce action=add-src-to-address-list address-list=Stage_1 address-list-timeout=5m  comment="First attempt" 
    add chain=Bruteforce action=accept 
```
```
    /ip firewall filter
    add chain=Protected-chain action=passthrough comment="Protect chain"
    add chain=Protected-chain action=jump jump-target=Detect-DDoS connection-state=new
    add chain=Protected-chain action=jump jump-target=Detect-SynFlood connection-state=new protocol=tcp tcp-flags=syn
    add chain=Protected-chain action=jump jump-target=Detect-SynAttack connection-state=new
    add chain=Protected-chain action=jump jump-target=Detect-PortScanners connection-state=new
    /ip firewall filter
    add chain=Detect-DDoS action=passthrough comment="DDoS detect chain" 
    add chain=Detect-DDoS action=return dst-limit=32,32,src-and-dst-addresses/10s 
    add chain=Detect-DDoS action=add-dst-to-address-list address-list=ddos-targets address-list-timeout=10m 
    add chain=Detect-DDoS action=add-src-to-address-list address-list=ddos-attackers address-list-timeout=10m 
    /ip firewall filter
    add chain=Detect-SynFlood action=passthrough comment="SynFlood detect chain"
    add chain=Detect-SynFlood action=return connection-state=new protocol=tcp tcp-flags=syn limit=200,5:packet
    add chain=Detect-SynFlood action=add-dst-to-address-list address-list=Banned address-list-timeout=1w3d
    /ip settings 
    set tcp-syncookies=yes
    /ip firewall filter
    add chain=Detect-SynAttack action=passthrough comment="SynAttack detect chain"
    add chain=Detect-SynAttack action=return protocol=tcp tcp-flags=syn,ack dst-limit=32,32,src-and-dst-addresses/10s
    add chain=Detect-SynAttack action=add-dst-to-address-list address-list=Banned address-list-timeout=1w3d
    /ip firewall filter
