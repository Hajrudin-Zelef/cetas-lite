---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1b2ewwo-using-sslvpn-reduce-your-security-footprint-block-8d4c8ad7-3
title: "r-fortinet-comments-1b2ewwo-using-sslvpn-reduce-your-security-footprint-block-8d4c8ad7"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit", "research"]
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1b2ewwo-using-sslvpn-reduce-your-security-footprint-block-8d4c8ad7.md
source_anchor: ""
source_lines: [131, 150]
sha256: 3dd884e1d4ce27de3b2d9d55ccbc8ec2b0f95a6f488325bcb80161fe25b09f14
---

# r-fortinet-comments-1b2ewwo-using-sslvpn-reduce-your-security-footprint-block-8d4c8ad7

I use ASN info to look them up. For example: https://asn.ipinfo.app/AS8100
You can then go to Resources -> Blacklists -> List -> Text and you get this https://asn.ipinfo.app/api/text/list/AS8100
You can then create an external feed, set it to update maybe once an hour or more, then add that external feed into your firewall policy. It's great and easy to do.
I saw a comment in a different thread where someone was creating lists of IPs from https://bgp.he.net/ and then creating pastebins with the ips.
Did you happen to crosscheck these with existing ISDB for malicious IPs or with any other public threatfeeds?
Block TOR networks if you haven’t already.
Ive noticed a few IPs listed as TOR exit nodes. Havent seen many of them yet though. Trick is that some of the nodes are listed as ISPs not Hosting providers. I try to be careful with blocking ASNs of ISPs as that might cause problems for traveling employees.
Fortinet ISDB and External IP block list like abuseIP or another other doesnt help much in this case?
Its easier to manage my own lists. When botnets and attackers started spamming attempts on the ivanti exploit, it opened my eyes to how many silent hosts were just waiting to become active.
Since we have no reason for hosting ASNs to login to our VPN, there is no reason to keep their IPs unblocked. If tomorrow there is a 10/10 CVE from fortinet, im already sitting better than many others as networks known to be used for these attacks are already blacked out for me.
I used to volunteer as a moderator for a phpbb forum, and for a couple of months we had a persistent spammer who kept posting successfully no matter how many times we blocked their address.
I spent an evening examining the attacks, and determined initially that they used a different IP address each time, but they were all withing a common address block. I then verified that nobody else was posting from that address block, so we were safe to block off the entire thing.
It didn't help. The spammer kept getting through anyway, for several other distinct address blocks.
I examined WHOIS records for those address blocks, and they all had a common owner. I then searched WHOIS for that owner, and discovered more than 50 additional large blocks of addresses owned by the same spammer.
I wrote up the results of my research and sent it to the owner of the phpbb forum, and recommended he block all of them. He did so, and the spammer never posted again afterward.
MS-ISAC and cisa already have the list complied and you can hook your gate into it by using the external threat feeds in security fabric
u/CallEither683 : Can you refer to exact link of MS-ISAC and cisa?
Thanks for posting this. I've now implemented this in my environment and am seeing way less spam in my logs.
What a great idea! I will definitely be using this.
Excellent idea this and some brilliant variations on theme. I’m going to start doing this for all my ssl von endpoints as well.
