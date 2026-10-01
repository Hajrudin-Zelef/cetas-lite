---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1836339-why-is-traffic-not-passing-through-opnsense-firewall-in-hyper-e717a645
title: "Why is traffic not passing through OPNsense firewall in Hyper V?"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1836339-why-is-traffic-not-passing-through-opnsense-firewall-in-hyper--e717a645.md
source_anchor: ""
source_lines: [1, 8]
sha256: 05ac48672d6d2ffd607cf29772cf25dfb07736c1e5235f73a1189d1b3328ea69
---

# Why is traffic not passing through OPNsense firewall in Hyper V?

*Score : 1 | Source : https://superuser.com/questions/1836339/why-is-traffic-not-passing-through-opnsense-firewall-in-hyper-v*

I'm configuring OPNsense in Hyper-V. It is connected to the DMZ on one interface, and the internal network on the other, with internal switches for the VMs. There are VLANs setup and working fine.
The OPNsense Firewall is for virtual machines on this Hyper-V host. The problem is I can reach the Internet from the VMs. I can also connect multiple virtual machines on different virtual switches through the OPNsense instance, but I cannot apply that configuration so that they can reach a device on the internal network such as a NAS. I can ping the NAS as expected from the OPNsense firewall.
I know that Hyper-V has some functionality to block fake or false routers, but that should be an opt in feature.
Network diagram:
