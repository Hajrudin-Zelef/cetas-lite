---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-vxlan-0007-html-65e6afe0
title: "hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-vxlan-0007-html-65e6afe0"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-vxlan-0007-html-65e6afe0.md
source_anchor: ""
source_lines: [1, 4]
sha256: b72ea3f3a5139e207eff59db1035031fa8a7d3b96cf828cda5ecbed3f687883a
---

# hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-vxlan-0007-html-65e6afe0

In Figure 1, an enterprise deploys its departments in different areas. To facilitate management and maintenance, departments with the same service requirements are planned in the same network segment, while those with different service requirements are planned in different network segments. End users in the same or different departments need to communicate with each other. For example, R&D departments 1 and 2 need to communicate in the same network segment; the R&D department 2 and marketing department need to communicate across different network segments.
VXLAN provides Layer 2 interconnection for dispersed physical sites. For example, in Figure 1, Router1 and Router2 are VXLAN Layer 2 gateway, and they establish a VXLAN tunnel to enable end users in R&D departments 1 and 2 to communicate with each other in the same network segment.
VXLAN provides Layer 3 interconnection for tenants in different sites. For example, when the R&D department 2 wants to communicate with the marketing department, Router3 functions as the VXLAN Layer 3 gateway to establish VXLAN tunnels with Router2 and Router4 respectively.
After static VXLAN tunnels are established between the routers, they dynamically learn flow table information, such as MAC address entries and ARP entries. After flow table information is learned, end users in the same or different network segments can communicate with each other over the VXLAN tunnels.
