---
id: collect-261001-fortinet/fortinet/2017-07-fortigate-vm-high-availability-vmware-configuration-182e16a0
title: "FortiGate VM High Availability VMware configuration"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/2017-07-fortigate-vm-high-availability-vmware-configuration-182e16a0.md
source_anchor: ""
source_lines: [1, 10]
sha256: 232d5cdcbecc6130226f6549913d422ee83bd17d93d5eba6c78bd6458a8aca1f
---

# FortiGate VM High Availability VMware configuration

*Source : https://www.fortinetguru.com/2017/07/fortigate-vm-high-availability-vmware-configuration/*

If you want to combine two or more FortiGate-VM instances into a FortiGate Clustering Protocol (FGCP) High Availability (HA) cluster the VMware server’s virtual switches used to connect the heartbeat interfaces must operate in promiscuous mode. This permits HA heartbeat communication between the heartbeat interfaces. HA heartbeat packets are non-TCP packets that use Ethertype values 0x8890, 0x8891, and 0x8890. The FGCP uses link-local IPv4 addresses in the 169.254.0.x range for HA heartbeat interface IP addresses.
To enable promiscuous mode in VMware:
You must also set the virtual switches connected to other FortiGate interfaces to allow MAC address changes and to accept forged transmits. This is required because the FGCP sets virtual MAC addresses for all FortiGate interfaces and the same interfaces on the different VM instances in the cluster will have the same virtual MAC addresses.
To make the required changes in VMware:
You can now proceed to power on your FortiGate VM. There are several ways to do this:
Select the Console tab to view the console. To enter text, you must click in the console pane. The mouse is then captured and cannot leave the console screen. As the FortiGate console is text-only, no mouse pointer is visible. To release the mouse, press Ctrl-Alt.
