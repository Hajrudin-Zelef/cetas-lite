---
id: collect-261001-meraki/meraki/t5-network-platform-api-terraform-meraki-mx-l3-firewall-rules-order-td-p-5412783-dbeb1083-1
title: "t5-network-platform-api-terraform-meraki-mx-l3-firewall-rules-order-td-p-5412783-dbeb1083"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-meraki/t5-network-platform-api-terraform-meraki-mx-l3-firewall-rules-order-td-p-5412783-dbeb1083.md
source_anchor: ""
source_lines: [1, 22]
sha256: 14d2b282bcddb6b3eb312f15977213d190de398b05bef9ed93098b77836bf608
---

# t5-network-platform-api-terraform-meraki-mx-l3-firewall-rules-order-td-p-5412783-dbeb1083

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-16-2025 11:59 PM
Hi gang,
I'm working on a full IaC deployment of a Meraki organization using terraform.
I notice some issues with the firewall L3 rule ordering when applying the terraform code.
Since terraform applies all code at the same time unless a dependency is decleared the 10 or so starting rules i have end up in a random order. For the most part i dont care about rule order, but i got some deny rules that must be placed in a spesific sequence.
I notice that the API endpoint for L3 FW rules also does not contain any parameters for sequence.
Has anyone worked around this in a way that is scaleable?
Also if some meraki employees read this, is it possible to add a feature request for firewall sequence numbering?
Thanks in advance!
Solved! Go to Solution.
- Labels:
- 
						
							
		
