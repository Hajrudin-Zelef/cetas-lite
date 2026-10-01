---
id: collect-261001-general-networking/general-networking/dhcp-servers-and-relays-2
title: "dhcp-servers-and-relays"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/dhcp-servers-and-relays.md
source_anchor: ""
source_lines: [160, 178]
sha256: dc4a332674925640aad3735cb9f24ada03354d2370659ab49ef6d25271ebda05
---

# dhcp-servers-and-relays

To prevent users in the from changing their IP addresses and causing IP address conflicts or unauthorized use of IP addresses, you can bind an IP address to a specific MAC address using DHCP.

Use the CLI to reserve an IP address for a particular client identified by its device MAC address and type of connection. The DHCP server then always assigns the reserved IP address to the client. The number of reserved addresses that you can define ranges from 10 to 200 depending on the FortiGate model.

After setting up a DHCP server on an interface by going to **S****ys****t****e****m > Network > Interface**, select the blue arrow next to **A****d****va****n****ce****d** to expand the options. If you know the MAC address of the system select **C****r****ea****t****e New** to add it, or if the system has already connected, locate it in the list, select its check box and select A**d****d from DHCP Client List**.

You can also match an address to a MAC address in the CLI. In the example below, the IP address 10.10.10.55 for User1 is assigned to MAC address 00:09:0F:30:CA:4F.


config system dhcp reserved-address edit User1

set ip 10.10.10.55

set mac 00:09:0F:30:CA:4F

set type regular end

jorge
before, muy english is not good enough. muy cuestión is how can i block all routers if someone connect them. muy la is. FORTIGATE DHCP-AP-PCs.
