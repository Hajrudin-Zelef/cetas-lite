---
id: collect-261001-huawei/huawei/yaweffah0-huawei-mpls-lab-48f0fc0f
title: "yaweffah0-huawei-mpls-lab-48f0fc0f"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/yaweffah0-huawei-mpls-lab-48f0fc0f.md
source_anchor: ""
source_lines: [1, 9]
sha256: 91f646dcda8db7b0ef2b18ea72bd1827bc29d3ae021cfb26d3178f38321bd308
---

# yaweffah0-huawei-mpls-lab-48f0fc0f

In this lab, we will setup a Service Provider MPLS network and Provide MPLS VPN services. This lab will be done with the ENSP software. Base topology file will be provided configured with device hostnames, loopbacks and Point to Point links.
You can track progress of project on notion. https://yaweffah.notion.site/MPLS-LAB-29afe946580e4d03ab463dc05feacf01
ISIS will be is as our IGP to establish basic network connectivity between all RRs, Ps and PEs.
Loopback address will be used as the LSR IDs for each node. MPLS RSVP TE will be enabled to make use of Fast Reroute and Constrained Shortest Path First(CSPF)
There will be three RRs in three different locations in the topology. All Ps and PEs will connect to the 3 RRs separately. the RRs will not be in a cluster.
Various configurations tasks to be completed are summarized below.
IGP CONFIGURATION Configure IGP ISIS Enable ISIS Configure Network Entity (NSAP Addresses) for routers Connfigure ISIS Level-2 only Configure Wide metrics only Configure automatic calculations of link cost Enable traffic engineering Enable ISIS for Backbone interfaces including Loopback Configure Point to Point link type for all ISIS interfaces Configure bfd for ISIS interfaces
MPLS CONFIGURATION Configure MPLS LSR ID Enable MPLS Configure MPLS-TE Configure MPLS TE Fast Reroute Enable RSVP-TE Enable MPLS CSPF Enable MPLS TE and MPLS RSVP-TE on backbone interfaces Configure Tunnel Interfaces for all PEs to mesh each other IP address of Tunnel shuld be same as Loopback Address Configure MPLS-TE rsvp as tunnel protocol Enable Fast Re-route for Tunnels
BGP CONFIGURATION Configure 2 RRs in cluster mode Configure BGP peer between RR and PEs Enable Peer tracking Ensure routes are not reflected between clients Enable BGP VPN4 on all BGP peers Provison l2VPN and L3VPN Services
