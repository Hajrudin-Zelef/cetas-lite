---
id: collect-261001-cisco/cisco/t5-network-security-asa-nat-best-practice-td-p-4031530-4f2d2ba1
title: "t5-network-security-asa-nat-best-practice-td-p-4031530-4f2d2ba1"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-network-security-asa-nat-best-practice-td-p-4031530-4f2d2ba1.md
source_anchor: ""
source_lines: [1, 73]
sha256: 5bdc91fbfd7e548e6e4f337d816f89d4cb730fcbe13624fd753d40d578727dce
---

# t5-network-security-asa-nat-best-practice-td-p-4031530-4f2d2ba1

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-18-2020 04:18 AM
Hi
I am in the process of reconfiguring all the outside access rules and NATs as we are migrating to a new public IP range. My question is about the best practice when configuring the NAT and access rules. I want to only use manual NATs.
1) Should the outside in access rule have the destination as the mapped public IP (so any to public ip) or the real IP address (any to real ip) of the internal server
2) Should the nat rule (although bidirectional) be inside to outside (real inside real outside translated mapped inside real outside) or the other way around. I know the rule will be bidirectional and I can make it unidirectional but what works as best practice.
Thanks
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Firewalls
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-18-2020 04:25 AM
You would always define the real IP address in the ACL.
Best Practice is to be consistent with your NAT rules. Source should be highest security level to lowest - e.g "nat (inside,outside) ...."
HTH
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-18-2020 02:26 PM
The biggest rule, as mentioned by RJI is to be consistant with your NAT and ACL configurations. However, there are some rules I try to follow as best as possible (though it is not easily done in some situations)
1. Configure NAT rules based on an inside to outside traffic flow (i.e. higher security level to lower security level)
2. Always define NAT source and destination interfaces (do not use "any" for an interface)
3. Try to be as specifc as possible with the IPs / subnets and ports in ACLs (this is particularly difficult as server administrators do not always know the traffic flow of their applications.)
4. Restrict access between internal devices (a PC needs to reach the AD, DHCP, DNS and printers, etc., but doesnt need to reach other PCs..usually)
ACLs require the use of the real IP address of an internal host.
Please remember to select a correct answer and rate helpful posts
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-18-2020 04:25 AM
You would always define the real IP address in the ACL.
Best Practice is to be consistent with your NAT rules. Source should be highest security level to lowest - e.g "nat (inside,outside) ...."
HTH
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-18-2020 02:26 PM
The biggest rule, as mentioned by RJI is to be consistant with your NAT and ACL configurations. However, there are some rules I try to follow as best as possible (though it is not easily done in some situations)
1. Configure NAT rules based on an inside to outside traffic flow (i.e. higher security level to lower security level)
2. Always define NAT source and destination interfaces (do not use "any" for an interface)
3. Try to be as specifc as possible with the IPs / subnets and ports in ACLs (this is particularly difficult as server administrators do not always know the traffic flow of their applications.)
4. Restrict access between internal devices (a PC needs to reach the AD, DHCP, DNS and printers, etc., but doesnt need to reach other PCs..usually)
ACLs require the use of the real IP address of an internal host.
Please remember to select a correct answer and rate helpful posts
