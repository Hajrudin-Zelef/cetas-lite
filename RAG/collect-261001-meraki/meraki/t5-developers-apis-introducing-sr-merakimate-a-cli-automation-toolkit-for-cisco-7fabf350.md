---
id: collect-261001-meraki/meraki/t5-developers-apis-introducing-sr-merakimate-a-cli-automation-toolkit-for-cisco-7fabf350
title: "t5-developers-apis-introducing-sr-merakimate-a-cli-automation-toolkit-for-cisco--7fabf350"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-developers-apis-introducing-sr-merakimate-a-cli-automation-toolkit-for-cisco--7fabf350.md
source_anchor: ""
source_lines: [1, 66]
sha256: 9ee019bc373112d008825fbd7e4cea73ec36516adf6d8ccfdd4cafced745f1c9
---

# t5-developers-apis-introducing-sr-merakimate-a-cli-automation-toolkit-for-cisco--7fabf350

🚀 Introducing SR-MerakiMate – A CLI Automation Toolkit for Cisco Meraki
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-28-2025 04:24 AM
Hi Meraki Community ,
I’ve built and open-sourced a project called SR-MerakiMate, a Python-based CLI toolkit that extends what we can do beyond the Dashboard GUI and templates.
Why SR-MerakiMate? While the Meraki Dashboard is excellent for day-to-day tasks, it can be limiting when:
We need bulk automation (hundreds of VLANs, DHCP ranges, firewall rules).
Templates are too rigid and don’t allow per-site flexibility.
We want secure API key handling (Azure Key Vault, masked prompts). We need reporting/export tools that aren’t available natively.
Features YAML & Excel-based bulk configuration (VLANs, DHCP, firewall rules, VPN exclusions). Device status & advanced inventory exports (CSV/XLSX). Policy objects automation (create/delete/group). Site-to-Site VPN viewer (with masked secrets). Troubleshooting assistant with offline knowledge base.
GitHub Repo → https://github.com/srajiwate/SR-Meraki-Mate
Disclaimer: This is a community-driven, proof-of-concept project — not affiliated with Cisco Meraki. Please test responsibly in lab or authorized environments.
Would love to hear from engineers
— what features would you like to see added?
- Labels:
- 
						
							
		
			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-28-2025 12:49 PM
That is a lot of work you have put in! A great effort.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-28-2025 05:40 PM
Thanks came with idea where we use to struggle for day to day operation
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-29-2025 08:34 AM
Thanks Shadab .. this is really usefull when we do mass deployment like same FW rule or traffic shaping across 100 sites .. I feel this is better than API via postman or other tool
