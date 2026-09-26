---
id: collect-260926-mikrotik/mikrotik/ipsec-tunnel-configuration
title: "on side1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/ipsec-tunnel-configuration.md
source_anchor: ""
source_lines: [1, 19]
sha256: 53c79744b23e371fed8da45f22fff8fa88f2f0b4735f57fd9597e59fbb3a147f
---

# on side1

Hello,

you may try this, if you have static ip-addresses on your pppoe-client interfaces

# on side1

```
/ip ipsec peer add addr=93.138.77.119 secret="your_very_strong_secret" nat-traversal=yes
/ip ipsec policy add src-addr=172.16.1.0/24 dst-addr=192.168.2.0/24 sa-src-addr=78.0.208.170 sa-dst-addr=93.138.77.119 tunnel=yes
/ip firewall nat add place-before=0 chain=srcnat action=accept src-addr=172.16.1.0/24 dst-addr=192.168.2.0/24 out-interface=INTERNET comment="NAT bypass for IPsec"
```

# on side2

```
/ip ipsec peer add addr=78.0.208.170 secret="your_very_strong_secret" nat-traversal=yes
/ip ipsec policy add src-addr=192.168.2.0/24 dst-addr=172.16.1.0/24 sa-src-addr=93.138.77.119 sa-dst-addr=78.0.208.170 tunnel=yes
/ip firewall nat add place-before=0 chain=srcnat action=accept src-addr=192.168.2.0/24 dst-addr=172.16.1.0/24 out-interface=INTERN comment="NAT bypass for IPsec"
```
