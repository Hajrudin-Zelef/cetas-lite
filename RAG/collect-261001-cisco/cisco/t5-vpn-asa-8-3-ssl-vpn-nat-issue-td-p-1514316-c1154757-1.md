---
id: collect-261001-cisco/cisco/t5-vpn-asa-8-3-ssl-vpn-nat-issue-td-p-1514316-c1154757-1
title: "sh ru object"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-vpn-asa-8-3-ssl-vpn-nat-issue-td-p-1514316-c1154757.md
source_anchor: ""
source_lines: [1, 21]
sha256: a355db5832d48efa2df4dcf33ab45a386fd72c4d5f972656264953d7fea44902
---

# sh ru object

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-22-2010 01:38 PM
Need help in figuring out how to setup anyconnect VPN with VPN client NATed into internal network.
There're a lot articles about opposite - how to disable NAT for vpn pool.
I need to create VPN gateway to complex interna lnetwork, vpnpool is out of regular subnet range of that network, so it'll be routing issues witout NAT.
So I need vpn clients connected to <outside> to be PATed to <inside>. The problem is that there's also dynamic PAT rule from <inside> to <outside> for regular Iternet acccess which results in "Asymetric NAT rules..." error.
Creating different Twice NAT rules and moving them on top/bottom doesn't make any difference. There're also some hidden rules from vpn setup which couldn't be seen.
v8.3 seems is trying to destroy confidence in Cisco firewalls...
Thank you.
Solved! Go to Solution.
- Labels:
- 
						
							
		
