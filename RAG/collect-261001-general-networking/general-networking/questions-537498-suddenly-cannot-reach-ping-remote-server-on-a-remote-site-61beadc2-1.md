---
id: collect-261001-general-networking/general-networking/questions-537498-suddenly-cannot-reach-ping-remote-server-on-a-remote-site-61beadc2-1
title: "questions-537498-suddenly-cannot-reach-ping-remote-server-on-a-remote-site-61beadc2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-537498-suddenly-cannot-reach-ping-remote-server-on-a-remote-site-61beadc2.md
source_anchor: ""
source_lines: [1, 9]
sha256: c1efd9b57a9d33dc213a9352399a66a7af27b521ac29b20ab0d00c0218ba0109
---

# questions-537498-suddenly-cannot-reach-ping-remote-server-on-a-remote-site-61beadc2

We have 2 sites linked together with VPN tunnel (Fortigate 60C devices). On each site I have ESXi server with a couple of VMs. Normally, everything works fine.
Site 1 (S1) subnet is 192.168.254.0/24, with Machine A1, A2 on ESXi1
 
Site 2 (S2) subnet is 192.168.253.0/24, with Machine B1, B2 on ESXi2
All ping between those machines works normally through VPN tunnel.
Suddently, S1-A1 cannot ping S2-B1 anymore, but S2-B1 still ping S1-A1.
All pings (using IP addresses) accross all machines (VMs and ESXi) works except from S1-A1 -> S2-B1.
Traceroute results were:
 
