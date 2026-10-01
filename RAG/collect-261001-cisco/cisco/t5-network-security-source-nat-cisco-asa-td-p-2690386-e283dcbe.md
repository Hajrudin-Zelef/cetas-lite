---
id: collect-261001-cisco/cisco/t5-network-security-source-nat-cisco-asa-td-p-2690386-e283dcbe
title: "t5-network-security-source-nat-cisco-asa-td-p-2690386-e283dcbe"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-network-security-source-nat-cisco-asa-td-p-2690386-e283dcbe.md
source_anchor: ""
source_lines: [1, 143]
sha256: 9f2443f80741025d8ccd9393614313edaaf31fe9cb883330384dcc2717baf8e5
---

# t5-network-security-source-nat-cisco-asa-td-p-2690386-e283dcbe

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-27-2015 12:44 PM - edited 03-11-2019 11:11 PM
Hi
I am trying to understand the concept on source NAT (not sure if this is the same as Twice NAT).
So I have attached a sample topology. I am in the process of migrating from watchguard firewalls to cisco ASAs, and during the migration I have come across this issue I am trying to get my head around. Im pretty new to configuring firewalls.
So If external Host A is trying to access my internal Server A via ASA 2, the traffic comes in, but on return it will hit the default gateway on the core switch which points to ASA 1. I was told that I can configure Source NAT, to force that traffic to return via the same firewall which involves natting on the Inside and Outside interfaces.
Just wondering if anyone is able to shed some light on this or knows of a good link.
Thankyou
Solved! Go to Solution.
- Labels:
- 
						
							
		
			NGFW Firewalls
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-28-2015 06:37 AM
Hi,
I believe that from outside anyone can connect to your server at 88.88.88.254.
So in that case, the NAT statement will be :
nat (outside,inside) source dynamic any interface destination static 88.88.88.254 10.10.10.50
Ensure that we have already allowed the required traffic through access rule in inbound direction.
This will certainly make it work.
Regards,
Pulkit Saxena
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-27-2015 01:12 PM
Hi,
Looking into the topology, it seems that we need to get the return traffic go through ASA 2 directly.
For that we need to do source and destination NAT both.
First of all, to answer your question regarding source NAT and twice NAT.
Source NAT simply means to NAT the source IP. For instance, all inside users when go to internet gets translated to outside interface IP.
Twice NAT also called as manual NAT is a feature on code 8.3 and above where in a single NAT statement you can NAT the source and destination both.
In this scenario of your's, the statement's syntax should be like :
nat (outside,inside) source dynamic any interface destination static mapped-ip real-ip .
This will ensure that your return traffic goes to firewall 2.
Some ASA NAT translation links that will be helpful :
https://supportforums.cisco.com/document/33921/asa-pre-83-83-nat-configuration-examples
http://www.cisco.com/c/en/us/td/docs/security/asa/asa83/configuration/guide/config/nat_overview.html
Thanks,
Pulkit Saxena
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-28-2015 03:59 AM
Hi Pulkit, thanks for the response.
I am using asa version 9.3 and currently on ASDM. So am I right in thinking I should configure NAT on both the source interface and destination interface?
What would be my translated source addresses for each. The translated destination I assume will stay the same
Thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-28-2015 04:06 AM
Hi,
There will be a single NAT statement only.
Try going into CLI through ASDM, and apply the NAT statement that I have given above.
In this particular scenario, source will get translated to inside interface IP.
Destination will get translated from the mapped IP to the actual IP.
If you can provide me all actual IP addresses, I can help you with the NAT statement.
Regards,
Pulkit Saxena
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-28-2015 04:26 AM
Hi Pulkit
So addresses are as follows, example IPs of course
Inside Server 10.10.10.50
ASA2 Inside Int 10.10.10.1
ASA2 Outside Int 88.88.88.254
Outside Host 98.98.98.50
Thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-28-2015 06:37 AM
Hi,
I believe that from outside anyone can connect to your server at 88.88.88.254.
So in that case, the NAT statement will be :
nat (outside,inside) source dynamic any interface destination static 88.88.88.254 10.10.10.50
Ensure that we have already allowed the required traffic through access rule in inbound direction.
This will certainly make it work.
Regards,
Pulkit Saxena
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-29-2015 03:23 AM
Thanks Pulkit
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-26-2015 09:44 PM
Hi Pulkhit,
I have got similar kinda issue. But the difference is Real Source (x.x.x.x) and destinations (x.x.x.x) IP's are belong to Same Subnet and NAT already exist to translate destination's subnet IP's to other IP's (y.y.y.y) so they can talk to other networks.
Requirement :
Host 1.1.1.1 in Environment A needs to talk to node (1.1.1.20) in Environment B
Environment B inside interface of ASA (image using 9.13) already translating 1.1.1.20 to 3.3.3.20 using static nat entering from outside interface.
If you have any solution to it let me know.
Warm Regards,
