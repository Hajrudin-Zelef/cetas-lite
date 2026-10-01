---
id: collect-261001-fortinet/fortinet/2016-10-ha-with-fortigate-vm-and-third-party-products-13495236
title: "HA with FortiGate-VM and third-party products"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/2016-10-ha-with-fortigate-vm-and-third-party-products-13495236.md
source_anchor: ""
source_lines: [1, 20]
sha256: f29e6823c86ca94597339dbe8d3b5137ef5f6badaa6069cc355d8f5a26f09765
---

# HA with FortiGate-VM and third-party products

*Source : https://www.fortinetguru.com/2016/10/ha-with-fortigate-vm-and-third-party-products/*

HA with FortiGate-VM and third-party products
This chapter provides information about operating FortiOS VM cluster and operating FortiGate clusters with third party products such as layer-2 and layer-3 switches.
FortiGate–VM for VMware HA configuration
If you want to combine two or more FortiGate-VM instances into a FortiGate Clustering Protocol (FGSP) High Availability (HA) cluster the VMware server’s virtual switches used to connect the heartbeat interfaces must operate in promiscuous mode. This permits HA heartbeat communication between the heartbeat interfaces. HA heartbeat packets are non-TCP packets that use Ethertype values 0x8890, 0x8891, and 0x8890. The FGCP uses link-local IP4 addresses in the 169.254.0.x range for HA heartbeat interface IP addresses.
To enable promiscuous mode in VMware:
1. In the vSphere client, select your VMware server in the left pane and then select the Configuration tab in the right pane.
2. In Hardware, select Networking.
3. Select Properties of a virtual switch used to connect heartbeat interfaces.
4. In the Properties window left pane, select vSwitch and then select Edit.
5. Select the Security tab, set Promiscuous Mode to Accept, then select OK.
6. Select Close.
You must also set the virtual switches connected to other FortiGate interfaces to allow MAC address changes and to accept forged transmits. This is required because the FGCP sets virtual MAC addresses for all FortiGate interfaces and the same interfaces on the different VM instances in the cluster will have the same virtual MAC addresses.
To make the required changes in VMware:
3. Select Properties of a virtual switch used to connect FortiGate VM interfaces.
4. Set MAC Address ChangestoAccept.
5. Set Forged Transmits to Accept.
