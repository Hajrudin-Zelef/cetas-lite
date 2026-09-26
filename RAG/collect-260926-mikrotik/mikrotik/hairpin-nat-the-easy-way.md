---
id: collect-260926-mikrotik/mikrotik/hairpin-nat-the-easy-way
title: "hairpin-nat-the-easy-way"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/hairpin-nat-the-easy-way.md
source_anchor: ""
source_lines: [1, 54]
sha256: d0f96733b777244b3c4b78248b8ebb985f1eb161c8cc08eb9b332e726f7b3fc0
---

# hairpin-nat-the-easy-way

Decided to write a simple guide on Hairpin NAT, because quite a lot of users struggle to understand how to set it up.

**I am not a networking professional and I am open to any criticism on how to implement it in a better way.**

Official wiki page by Mikrotik regarding Hairpin NAT: https://wiki.mikrotik.com/wiki/Hairpin_NAT

**Step 1 - add LANs to “address-list”**

List all your LANs like this:

```
/ip firewall address-list add address=192.168.10.0/24 comment=Management list=LANs
/ip firewall address-list add address=192.168.11.0/24 comment=Work list=LANs
/ip firewall address-list add address=192.168.12.0/24 comment=Security list=LANs
/ip firewall address-list add address=192.168.13.0/24 comment=Home list=LANs
/ip firewall address-list add address=192.168.14.0/24 comment=Guest list=LANs
```

**Step 2 - add WANs to “address-list”**

If you have a single dynamic IP - add your “/ip cloud” domain to address-list named “WANs” and Mikrotik will automatically resolve it to IP. Using custom script in “/ip dhcp-client” is another option in order to keep WAN IP address in address-list updated.

If you have multiple WANs - it gets a little more complicated. I’ve written a simple solution for multiple dynamic WANs here: http://forum.mikrotik.com/t/hairpin-with-2-wan/145618/1

List all your WANs like this:

```
/ip firewall address-list add address=123.123.123.123 list=WANs
```

**Step 3 - mark connections from LANs to WANs**

Use this rule:

```
/ip firewall mangle add action=mark-connection chain=prerouting comment="Mark connections for hairpin NAT" dst-address-list=WANs new-connection-mark="Hairpin NAT" passthrough=yes src-address-list=LANs
```

**Step 4 - perform Hairpin NAT**

Use this rule, placed before any other NAT rule:

```
/ip firewall nat add action=masquerade chain=srcnat comment="Hairpin NAT" connection-mark="Hairpin NAT" place-before=0
```

**Step 5 - port forwarding**

Setup port-forwarding like this:

```
/ip firewall nat add action=dst-nat chain=dstnat comment="Port forward: something1" dst-address-list=WANs dst-port=5001 protocol=tcp to-addresses=192.168.0.8 to-ports=5001
/ip firewall nat add action=dst-nat chain=dstnat comment="Port forward: something2" dst-address-list=WANs dst-port=5002 protocol=tcp to-addresses=192.168.0.9 to-ports=5002
```
