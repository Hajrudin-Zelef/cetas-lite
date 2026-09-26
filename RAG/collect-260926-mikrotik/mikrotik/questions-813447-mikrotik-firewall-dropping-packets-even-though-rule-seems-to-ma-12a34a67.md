---
id: collect-260926-mikrotik/mikrotik/questions-813447-mikrotik-firewall-dropping-packets-even-though-rule-seems-to-ma-12a34a67
title: "questions-813447-mikrotik-firewall-dropping-packets-even-though-rule-seems-to-ma-12a34a67"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/questions-813447-mikrotik-firewall-dropping-packets-even-though-rule-seems-to-ma-12a34a67.md
source_anchor: ""
source_lines: [1, 14]
sha256: 84d6893b64ad535f76b6f64d13d84ca20fcbf57258f4a063dd3a3a7a101e8e7c
---

# questions-813447-mikrotik-firewall-dropping-packets-even-though-rule-seems-to-ma-12a34a67

I set up an L2TP VPN server on my Mikrotik. Connecting to the VPN when I'm behind the router works, but once I'm connecting from the WAN side, it doesn't. I logged my firewall to see if I was dropping it with the default drop rule and it was:
WANDROP input: in:ether1-gateway out:(none), src-mac 00:00:5e:00:01:66, proto UDP, 5.6.7.8:38211->1.2.3.4:500, len 412
But my rules are set up like this:
add action=accept chain=input dst-address=1.2.3.4 dst-port=500,1701,4500 in-interface=ether1-gateway log=yes protocol=udp src-address=0.0.0.0
    add action=accept chain=input dst-address=1.2.3.4 in-interface=ether1-gateway log=yes protocol=ipsec-esp src-address=0.0.0.0
    add action=drop chain=input comment="default configuration" in-interface=ether1-gateway log=yes log-prefix=WANDROP
    add action=accept chain=input comment="default configuration" connection-state=established,related
    add action=accept chain=forward comment="default configuration" connection-state=established,related
    add action=drop chain=forward comment="default configuration" connection-state=invalid
    add action=drop chain=forward comment="default configuration" connection-nat-state=!dstnat connection-state=new in-interface=ether1-gateway log=yes log-prefix=DROP
    /ip firewall nat
    add action=dst-nat chain=dstnat dst-address=1.2.3.4 dst-port=80 in-interface=ether1-gateway protocol=tcp to-addresses=192.168.88.232 to-ports=80
    add action=masquerade chain=srcnat comment="default configuration" out-interface=ether1-gateway
not sure what I'm doing wrong or missing.
