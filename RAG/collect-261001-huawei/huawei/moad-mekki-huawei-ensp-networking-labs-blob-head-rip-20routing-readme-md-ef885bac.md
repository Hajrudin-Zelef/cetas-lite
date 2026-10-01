---
id: collect-261001-huawei/huawei/moad-mekki-huawei-ensp-networking-labs-blob-head-rip-20routing-readme-md-ef885bac
title: "moad-mekki-huawei-ensp-networking-labs-blob-head-rip-20routing-readme-md-ef885bac"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/moad-mekki-huawei-ensp-networking-labs-blob-head-rip-20routing-readme-md-ef885bac.md
source_anchor: ""
source_lines: [1, 53]
sha256: d345f98a77f543f9857f26753e550718d02ca066c7d0e40c8611ce5a3fa01189
---

# moad-mekki-huawei-ensp-networking-labs-blob-head-rip-20routing-readme-md-ef885bac

Implement dynamic routing using RIP (Routing Information Protocol) to enable automatic route exchange between multiple routers.
- Dynamic Routing
- RIP Version 2
- Route Advertisement
- Network Convergence
- Multi-router Communication
The topology contains:
- 
3 Routers: 
  - AR1
  - AR2
  - AR3
- 
3 LANs: 
  - LAN 1 → PC1 Network
  - LAN 2 → PC2 Network
  - LAN 3 → PC3 Network
- 
Full router interconnection: 
  - AR1 ↔ AR2
  - AR2 ↔ AR3
  - AR1 ↔ AR3
RIP is configured on all routers to dynamically exchange routing information.
| Device | Interface | IP Address | 
|---|---|---|
| PC1 | - | 192.168.1.2/24 | 
| AR1 | G0/0/0 | 192.168.1.1/24 | 
| AR1 | G0/0/1 | 10.0.12.1/24 | 
| AR1 | G0/0/2 | 10.0.13.1/24 | 
| PC2 | - | 192.168.2.2/24 | 
| AR2 | G0/0/0 | 192.168.2.1/24 | 
| AR2 | G0/0/1 | 10.0.12.2/24 | 
| AR2 | G0/0/2 | 10.0.23.1/24 | 
| PC3 | - | 192.168.3.2/24 | 
| AR3 | G0/0/0 | 192.168.3.1/24 | 
| AR3 | G0/0/1 | 10.0.13.2/24 | 
| AR3 | G0/0/2 | 10.0.23.2/24 | 
- Assign IP addresses
- Enable RIP
- Advertise connected networks
- Use RIP Version 2
- Basic Layer 2 connectivity
✔ PC1 can ping PC2 ✔ PC1 can ping PC3 ✔ PC2 can ping PC3 ✔ RIP routes appear in routing tables
| Test Case | Result | 
|---|---|
| PC1 → PC2 | ✅ Success | 
| PC1 → PC3 | ✅ Success | 
| PC2 → PC3 | ✅ Success | 
- RIP dynamically exchanges routes between routers
- Routers automatically learn remote networks
- Dynamic routing reduces manual configuration effort
Implement:
- OSPF Routing
