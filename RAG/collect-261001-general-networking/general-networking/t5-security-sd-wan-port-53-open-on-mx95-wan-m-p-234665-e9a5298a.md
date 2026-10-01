---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-port-53-open-on-mx95-wan-m-p-234665-e9a5298a
title: "t5-security-sd-wan-port-53-open-on-mx95-wan-m-p-234665-e9a5298a"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-port-53-open-on-mx95-wan-m-p-234665-e9a5298a.md
source_anchor: ""
source_lines: [1, 123]
sha256: 9c17679c8e4c84d196752fa4b0dec4711bc47cabeabc96832f09a23f8a27ef8c
---

# t5-security-sd-wan-port-53-open-on-mx95-wan-m-p-234665-e9a5298a

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-06-2024 03:12 AM
Hi All
Our ISP was complaining about port 53 being open with an active dns resolver on it
dnsmasq-2.85
Now this is a new firewall that went live a few days ago, there is no port forwarding rules configured there, so why port 53 is open.. i tested from another public IP and port 53 is indeed open..
- Why this is the case and what other ports are expected to be open?
- how to make sure all ports are closed from internet initiated traffic as these should be?
We don't have any DNS service exposed on the public IP nor any NAT other than WAN standard NAT is on for inbound LANs to access the internet.. the MX is in routed mode obviously
Thanks
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
05-06-2024 04:29 AM
There is a ? On top of the dashboard screen. Then go to firewall info.
But i dont understand you L3 firewall dashboard view. Afaik Inbound rules should be for ipv6 only not for ipv4 traffic. Unless you made/requested some changes to the default settings
Maybe you running the early access: org>early access
"NAT Exceptions with Manual Inbound Firewall"
^^^
Looks like 'open port(s)' its related to above
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-06-2024 03:46 AM
It is a DNS resolver, and has to be able to accept DNS replies.
I think it turns off when you don't have any DHCP scopes including the security appliance as a DNS server.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-06-2024 03:58 AM
I don't have any DHCP scopes set, all VLANs are configured not to respond to any DHCP requests..
Nor are there setups pointing the appliance as a DNS relay or DNS.. I am pointing everywhere towards Google ones
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-06-2024 04:05 AM
Im not running mx95 and i dont see port53 open from outside.
You can take a look at ? > firewall info if that list port 53.
Otherwise create a support ticket and let them find out why its open
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-06-2024 04:19 AM
Hi @jdb1
I don't see a Firewall>Info section under my mx95.. however the inbound L3 rules are set to the default deny..
I don;t want to set an explicit deny for port 53 as this in theory might block all forwarder DNS traffic..
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-06-2024 04:29 AM
There is a ? On top of the dashboard screen. Then go to firewall info.
But i dont understand you L3 firewall dashboard view. Afaik Inbound rules should be for ipv6 only not for ipv4 traffic. Unless you made/requested some changes to the default settings
Maybe you running the early access: org>early access
"NAT Exceptions with Manual Inbound Firewall"
^^^
Looks like 'open port(s)' its related to above
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-06-2024 05:29 AM
Yes indeed, I enabled the early access on this one.. this is why you see the IPV4 rules.. looks more professional for me this way
Thanks for the pointer to the firewall info.. I found a strange snmp inbound rule there that isn't set anywhere in my rules to one public IP from a range routed to the WAN public IP with source the Meraki Networks.. no clue why this is there even though the target IP is not the WAN Ip
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-12-2024 05:38 AM
Hi All
Still with Meraki support on this one, they are trying to figure it out why it is open and responding to such requests..
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-05-2024 06:16 AM
Update on this issue: After spending some time with Meraki support, they did indeed point out to the fact that the early access feature of the L3 inbound tules are the culprit, if no rules are configured, an explicit allow all kicks in..which you don't see in the inbound L3 rules as the default is deny, see previous post screenshot.. confusing..
