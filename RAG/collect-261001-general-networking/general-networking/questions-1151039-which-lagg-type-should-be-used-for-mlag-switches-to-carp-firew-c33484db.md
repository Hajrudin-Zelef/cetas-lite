---
id: collect-261001-general-networking/general-networking/questions-1151039-which-lagg-type-should-be-used-for-mlag-switches-to-carp-firew-c33484db
title: "Which LAGG type should be used for MLAG switches to CARP firewalls?"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1151039-which-lagg-type-should-be-used-for-mlag-switches-to-carp-firew-c33484db.md
source_anchor: ""
source_lines: [1, 36]
sha256: 4062dfc43940ac5e5eb498aa9e0c9a3dc2f1859f1873be9133cb47f9b1f3c478
---

# Which LAGG type should be used for MLAG switches to CARP firewalls?

*Score : 4 | Source : https://serverfault.com/questions/1151039/which-lagg-type-should-be-used-for-mlag-switches-to-carp-firewalls*

My network setup involves two firewalls in a Common Address Redundancy Protocol (CARP) group, each connected to an MLAG (Multi-Chassis Link Aggregation) configuration of Mikrotik switches. Onward ports on the switches are bonded using LACP.
                   VIP (WAN)
                      |
          -------------------------
         |                         |
         |                         |
  ------------     pfSync    ------------
 | Firewall 1 | <---------> | Firewall 2 |
  ------------               ------------
         |                         |
         |                         |
         |                         |
  ------------      MLAG     ------------
 |  Switch 1  | <---------> |  Switch 2  |
  ------------               ------------
         |                         |
          -------------------------
                       |
                   LACP (LAN)
Question: How should the switch ports to the firewalls be configured?
MLAG allows the switches to appear to external hardware as though they are a single switch. LAGG are created across the physical switches (i.e. an LACP bond will have 1 port of the bond on each switch so that an entire switch can go down and the connection survive). How should the ports connected to the firewalls be bonded? Clearly LACP is not correct because the firewalls will not negotiate this correctly on their end as the firewalls act independently of each other. My options seem to be:
The challenge here is ensuring that the switch's active link aligns with the CARP master firewall. Without an automatic mechanism to align the active switch port with the CARP master, this setup could lead to misalignment where the active switch port is connected to the standby firewall.
This guarantees that the active CARP firewall always receives the traffic, regardless of which one is master, but if a firewall leaves the CARP group (say because CARP is disabled for maintainance) it will start handling duplicate traffic.

---

### Reponse (acceptee) — score 4

(M)LAG is used for aggregating links on the data link layer (L2), where there are multiple physical paths between different L2 nodes.
CARP (or HSRP, VRRP) is a network-layer redundancy protocol, using an active-passive model. It creates a virtual IP address (VIP) that moves from a previously active, failed L3 node to a previously passive, standby one. Usually the failover is announced to the network by gratuitous ARP (GARP) with the MAC address of the standby node.
That said, CARP doesn't interact with L2 any more than that. Specifically, it doesn't require nor use any LAG setup. On the contrary, LAG would only interfere with CARP and may obstruct it.
So, don't use any LAG towards the firewalls, regardless of whether you connect them to a single switch or two stacked ones. Only when you use multiple links to each firewall, then you should use LACP on each of them.
