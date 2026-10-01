---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1167790-freepbx-17-behind-unifi-firewall-cbba3885
title: "questions-1167790-freepbx-17-behind-unifi-firewall-cbba3885"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1167790-freepbx-17-behind-unifi-firewall-cbba3885.md
source_anchor: ""
source_lines: [1, 6]
sha256: ef306fe4aaeaee8e9e115b5def82d3ea697f59de731fb69a1b0c2d2c0283fef5
---

# questions-1167790-freepbx-17-behind-unifi-firewall-cbba3885

I have been running FreePBX for years now inside my Unifi network (USG Pro, Unifi switch, 3 AC Pro access points). I’m currently running version 16 (upgraded from 15) and am now looking to upgrade to version 17. I have already set everything up and restored a backup and all seems to be working fine except for my Trunks that are rejected when connecting. In v16 I’m using chan_sip trunks that I need to change to chan_pjsip for v17. So first I thought my problems were with version 17 and I went back to my v16 install and try to add the pjsip trunks there, but I got the same result. I tried on a server of one of my clients which is a Vultr virtual server which connected without problem, so I began to think it must be my install being behind a firewall.
asterisk -x "pjsip show registrations"
 CheapConnect-pjsip/sip:sip.cheapconnect.net            CheapConnect-pjsip         Rejected         (exp. 24s)
 Frynga_prive-pjsip/sip:sip.frynga.com                  Frynga_prive-pjsip         Rejected         (exp. 25s)
Any tips? Is there need to open up my firewall?
SIP/ALG is turned off. I have also tried on another network with a simple firewall (the standard one a provider gives you) and there it also works fine.
