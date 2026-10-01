---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/r-opnsense-comments-12uylzn-lan-to-internet-traffic-intermittently-blocked-60f56386
title: "LAN to internet traffic intermittently blocked"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/r-opnsense-comments-12uylzn-lan-to-internet-traffic-intermittently-blocked-60f56386.md
source_anchor: ""
source_lines: [1, 19]
sha256: 88934050060ba9d3efe5b388737c54b8397d181a4ba6fb939aee1da7a894744f
---

# LAN to internet traffic intermittently blocked

I am on the latest 23.1.6 OPNsense and have a very simply firewall config. It is basically stock on the LAN side. I am seeing some traffic from LAN to the internet intermittently being blocked by the firewall and cannot explain why this is even happening as there is a rule to allow all LAN net traffic to any destination.

Basically stock LAN rules

Here is an example of a filterlog entry showing a block

4,,,02f4bab031b57d1e30553ce08e0ec131,em2,match,block,in,4,0x0,,64,34116,0,none,6,tcp,52,192.168.47.29,142.250.189.163,23724,443,0,FA,3413329432,512495170,4096,,nop;nop;TS

Note, em2 is the LAN interface.

I cannot figure out why it would block any internet bound traffic from LAN when there is a rule clearly allowing it.

Also, I think that first 4 is the rule number. How the heck do I find out what is rule 4?

no comments yet

Be the first to share what you think!
