---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-972593-opnsense-pfsense-on-proxmox-with-lacp-vlans-and-multi-wan-82f4509f
title: "OPNsense/pfsense on Proxmox with LACP, VLANs and Multi-WAN"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-972593-opnsense-pfsense-on-proxmox-with-lacp-vlans-and-multi-wan-82f4509f.md
source_anchor: ""
source_lines: [1, 13]
sha256: 3486f2ac7344a79c3307398ff150ad605069e6ce673ca94c59265a5de7982f4a
---

# OPNsense/pfsense on Proxmox with LACP, VLANs and Multi-WAN

*Score : 1 | Source : https://serverfault.com/questions/972593/opnsense-pfsense-on-proxmox-with-lacp-vlans-and-multi-wan*

I want to set up an OPNsense router (of pfsense if OPNsense sould not work as desired) on Proxmox. I have multiple WANs as well as multiple local VLANs. Since I have more VLANs than physical ports I obviously will have to use VLAN-trunking. And since routing between VLANs on a single gigabit line will be slow I want to use LACP to bond multiple ports to serve as my trunk.
My multiple WAN connections (4 at the moment) have a sum of less than 1 gigabit. I could use VLANs here as well in order to save on ports (= more ports for the LACP).
So, a simple setup (ignoring proxmox of now) would look like this
Now for the difficult part: I want to set this up as a VM on proxmox. One reason being, that, should the hardware fail, I could use a different server to run the VM on. But not all my proxmox hosts have the same number of ethernet ports. Therefore I am planning the following setup:
So now, in case I have to switch to another host, I could replace the LACP bond with a simple VLAN aware bridge over only one gigabit port. (lower speed, but things would still work)
Is there any fault in my setup? Something I missed?
Am I right, that OPNsense would not see any change when switching the host = no config changes needed
Do I need to use Open vSwitch in order for other VMs to be able to use the router?
What is the expected overhead of this setup compared to a bare-metal setup?
