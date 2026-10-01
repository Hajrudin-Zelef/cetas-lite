---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-how-do-mx-tags-interact-with-configuration-templates-for-vlan-92d956e7
title: "t5-security-sd-wan-how-do-mx-tags-interact-with-configuration-templates-for-vlan-92d956e7"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-how-do-mx-tags-interact-with-configuration-templates-for-vlan-92d956e7.md
source_anchor: ""
source_lines: [1, 59]
sha256: 1865494378224d2bf65ca5c4b30f1555b8571d18766ef57a2cf5eed386b09b3e
---

# t5-security-sd-wan-how-do-mx-tags-interact-with-configuration-templates-for-vlan-92d956e7

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2026 09:21 AM
Hello Meraki Community,
I’m looking for some clarification on how the tagging system works when using Meraki MX security appliances in combination with configuration templates.
Suppose I have a configuration template that assigns VLANs, subnets, and per‑port settings through the Addressing & VLANs section.
Is there any way to use tags to override these VLAN or subnet assignments on a per‑network basis? Or does the configuration template always take precedence and overwrite any local network‑specific VLAN or subnet changes, regardless of the tags applied?
Thanks in advance for any insights!
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2026 09:40 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2026 09:29 AM
When a network is bound to a template, you can override the subnet assigned to a VLAN locally on that network. However, the override must use a subnet from the same subnet pool defined in the template, and you cannot override the VLAN ID itself.
Tags themselves do not provide a mechanism to override VLAN or subnet assignments. Tags are used for device grouping, policy application, and dashboard filtering, but not for configuration overrides.
The configuration template sets the default VLAN and subnet assignments. Local overrides made directly on the network (via Security & SD-WAN > Addressing & VLANs) will take precedence for that network, as long as they comply with the template’s subnet pool rules.
If you unbind a network from the template, any local overrides and template-based configuration are lost, and the network reverts to its own configuration.
Tags do not override VLAN/subnet assignments.
Local configuration changes on a network bound to a template can override the template’s VLAN/subnet settings within the allowed subnet pool, but the template always controls the available options.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2026 09:37 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2026 09:40 AM
