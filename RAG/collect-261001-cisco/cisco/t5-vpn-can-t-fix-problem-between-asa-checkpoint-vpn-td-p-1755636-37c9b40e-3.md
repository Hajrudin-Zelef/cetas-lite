---
id: collect-261001-cisco/cisco/t5-vpn-can-t-fix-problem-between-asa-checkpoint-vpn-td-p-1755636-37c9b40e-3
title: "t5-vpn-can-t-fix-problem-between-asa-checkpoint-vpn-td-p-1755636-37c9b40e"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-vpn-can-t-fix-problem-between-asa-checkpoint-vpn-td-p-1755636-37c9b40e.md
source_anchor: ""
source_lines: [246, 287]
sha256: c4bec60ffc3e870d87a440eccd911fe210a01672ffbe2fc1791f30b6d45ca423
---

# t5-vpn-can-t-fix-problem-between-asa-checkpoint-vpn-td-p-1755636-37c9b40e

Federico, I have another suggestion, why don't you turn on "vpn debug ikeon" on the checkpoint side to see what going on during phase I and II. Checkpoint will tell you exactly how and why it failed, whether it is sending or receiving to and from the ASA.
You then can view the debug file with IKEView.exe. It will tell you exactly where things go wrong.
By the way, what version of checkpoint? Checkpoint on Nokia IPSO or Secureplatform? can you share the output "uname -a" and "fw ver"? Furthermore, are you setting this up in Simplified mode (VPN community) or traditional mode?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-16-2011 08:58 AM
Hi,
Check Point does supernetting by default.
If you have a host in your check point rulebase but the network is defined in the Check Point encryption domain then it will try to establish the SA with the encryption domain network instead of the host.
I suggest you use network to network between check point and ASA then do the filtering in the rulebase.
This way the SA is up between the 2 networks but the rulebase in Check Point will only allow the host.
You can have a VPN Filter ACL in the ASA to acheive the same security. The filter ACL permits only the host and is different from the interesting traffic ACL which is applied to the crypto map and used to establish the SA...
You can also disable supernetting in Check Point but that would affect all Check Point VPNs so be very careful!
You have to do that through the GUI db edit tool not smartdashboard...
Hope this helps
Patrick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-16-2011 09:11 AM
That's why I keep saying you need to turn on "vpn debug ikeon" and look at the file ike.elg file with IKEView.exe and it will tell you where things go wrong.
In Checkpoint VPN community, you have the option to do network or hosts. Keep in mind that if you do hosts instead of network, it will require more processing power because of the SA. That's the beauty of Checkpoint VPN simplify mode.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-16-2011 10:26 AM
Yes I agree, the debugs should tell you what exactly is happening. Same thing on the Cisco side, you should see if you receive a supernetted subnet or a host in the IPSEC phase 2 negotiations...
I think Check Point recommends network to network when establishing an IPSEC tunnel with a non-Check Point gateway.
Patrick
