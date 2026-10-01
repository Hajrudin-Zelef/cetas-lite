---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/r-opnsense-comments-10im9qd-nat-port-forwarding-not-working-at-all-82db6a34
title: "NAT Port Forwarding not working at all"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/r-opnsense-comments-10im9qd-nat-port-forwarding-not-working-at-all-82db6a34.md
source_anchor: ""
source_lines: [1, 19]
sha256: 7b56bcaabaee8341d29f16332c2fda0664d49e74ea79a7de5bff03b8b36d5c89
---

# NAT Port Forwarding not working at all

I am on the edge of desperation

I have now read quite a few posts and how:

but still do not get to the goal.

The internal traffic works as desired over all the VLANS, etc.. But when I want to share my gaming server, I just can not get it myself and according to opn-ports tools, the port is still shown as closed.

Now I ask myself, what am I missing?

I have a router from my ISP and there DMZ mode enabled, so it should forward everything to the opnsense. In the Opnsense I have entered the NAT port forwarding as in the forum above, from this was directly set up a rule in the WAN.

Under Firewall->Settings-> Advanced I have set the marks for Reflection for port forwards and Automatic outbound NAT for Reflection.

Despite this, I can not access it via my ext. IP.

Opnsense is still quite new territory for me, maybe that's why :P
