---
id: collect-261001-general-networking/general-networking/questions-1634635-pc-in-vlan-cant-ping-its-own-gateway-5cd13fa1
title: "questions-1634635-pc-in-vlan-cant-ping-its-own-gateway-5cd13fa1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["decode", "ethernet"]
source: docs/RAG/collect-261001-general-networking/questions-1634635-pc-in-vlan-cant-ping-its-own-gateway-5cd13fa1.md
source_anchor: ""
source_lines: [1, 29]
sha256: 1c163e3c77d0c9d24b69b0830ef64ed0f6fda95f9238d0b466bb4eb762eacef1
---

# questions-1634635-pc-in-vlan-cant-ping-its-own-gateway-5cd13fa1

My VLAN's can't ping their own gateway and beyond.

1 - For VLANs to be able to ping their own gateway, do I have to enable the Parent VLAN and assign an IP to it in the OPNsense firewall ?

2 - Do I have to enable DHCP for devices in the VLAN, I know sounds foolish but just checking as all videos I found regarding this on OPNSense show as DHCP enabled which I'm not using.

All rules are allowed on all interfaces, and I'm not using any physical switch, this is a VMware Workstation setup, I have a Windows VM and OPNsense VM.

**Parent VLAN Interface**

**Server VLAN**

**Server Ping Failure**

**tcpdump on em3 shows the following**

```
root@firewallwm:~ # tcpdump -e -n -i em3
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on em3, link-type EN10MB (Ethernet), capture size 262144 bytes
18:26:52.651787 00:0c:29:ae:a2:10 > ff:ff:ff:ff:ff:ff, ethertype ARP (0x0806), length 60: Request who-has 192.168.28.35 tell 192.168.28.47, length 46
18:26:53.316297 00:0c:29:ae:a2:10 > ff:ff:ff:ff:ff:ff, ethertype ARP (0x0806), length 60: Request who-has 192.168.28.35 tell 192.168.28.47, length 46
18:26:54.316412 00:0c:29:ae:a2:10 > ff:ff:ff:ff:ff:ff, ethertype ARP (0x0806), length 60: Request who-has 192.168.28.35 tell 192.168.28.47, length 46
```
I checked non device has the MAC address **00:0c:29:ae:a2:10**.

**Firewall Settings**

**Windows VM**
