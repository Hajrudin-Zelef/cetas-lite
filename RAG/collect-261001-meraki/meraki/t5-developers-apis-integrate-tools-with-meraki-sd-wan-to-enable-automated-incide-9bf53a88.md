---
id: collect-261001-meraki/meraki/t5-developers-apis-integrate-tools-with-meraki-sd-wan-to-enable-automated-incide-9bf53a88
title: "t5-developers-apis-integrate-tools-with-meraki-sd-wan-to-enable-automated-incide-9bf53a88"
domain: meraki
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-meraki/t5-developers-apis-integrate-tools-with-meraki-sd-wan-to-enable-automated-incide-9bf53a88.md
source_anchor: ""
source_lines: [1, 52]
sha256: 991cf21309650c9900a20fbd6e22d7dbc45c3c913b0d1ee054ab9c901f74f367
---

# t5-developers-apis-integrate-tools-with-meraki-sd-wan-to-enable-automated-incide-9bf53a88

integrate tools with Meraki SD-WAN to enable automated incident creation, alert routing
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-30-2024 10:59 AM
I'm exploring options to integrate tools with Meraki SD-WAN to enable automated incident creation, alert routing, and escalation. The goal is to streamline incident management and integrate alerts with our ticketing system for a unified workflow.
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
10-30-2024 12:30 PM
Check out this thread. In the posted solution you can find some templates and alert settings you can use with webhooks.
Depending on your ticketing system you might need to write a reform script that maps key value pairs from meraki to values your ticketing system understands.
The example in that thread is for Service Now.
https://documentation.meraki.com/General_Administration/Other_Topics/Webhooks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-04-2024 06:05 PM
If you happen to be looking for a packaged solution that does this, I suspect that most if not all of the Performance Monitoring products on the Meraki Marketplace do this.
https://apps.meraki.io/en-US/listing?cat=99861
Disclaimer: I work for one of these companies, but I imagine our rivals have similar ITSM and workflow integrations.
