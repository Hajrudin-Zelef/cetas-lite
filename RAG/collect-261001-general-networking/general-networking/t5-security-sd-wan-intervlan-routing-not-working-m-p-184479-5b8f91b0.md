---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-intervlan-routing-not-working-m-p-184479-5b8f91b0
title: "t5-security-sd-wan-intervlan-routing-not-working-m-p-184479-5b8f91b0"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-intervlan-routing-not-working-m-p-184479-5b8f91b0.md
source_anchor: ""
source_lines: [1, 172]
sha256: a97943d0423c6dc62f806886ce8ec3542e798da48c027b841543be57dbabc2e0
---

# t5-security-sd-wan-intervlan-routing-not-working-m-p-184479-5b8f91b0

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-11-2023 04:11 AM
Could you help me understand why I can't have internal communication between my VLANS, I have an mx64.
The situation is as follows, I have a desktop directly connected to the meraki that is on vlan 1 and has the ip 192.168.1.2 if on that desktop I use a VM that is on the same network (192.168.1.0/24) I can have normal communication.
However, if I put this VM on the VLAN 30 LABS network (192.168.30.0/24) I cannot have any type of communication, not even with the gateway (MX IP 192.168.30.1). However, from the desktop (ip 192.168.1.2) I can ping the gateway, but no other ip from the VLAN 30 range.
Below is how the VLANs are configured:
Meraki is currently with Client tracking in IP Address mode, even in Mac Address mode, the communication does not work:
There are no firewall rules or group policies blocking communication:
Group policies
ARP VLAN 30:
Tests vlan 30 not working
Test vlan 1 working
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
02-12-2023 04:45 PM
I found the problem.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-11-2023 04:19 AM
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-11-2023 04:37 AM
Hello,
I'm using VMware. I've already tried to do the tests by tagging the vlan or leaving it without tagging, but I wasn't successful anyway.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-11-2023 04:40 AM
Forget it, VMware running directly on Windows will not work, only with VMware ESXI.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-11-2023 05:07 AM
But why not?
This to me is simple routing on L3.
If I connect the windows desktop to the ISP's router (a very simple router, by the way) the communication works normally...
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-11-2023 05:09 AM
It' VMware limitation.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-11-2023 08:37 AM
Not sure about the VMware on Windows (Im assuming you are using VMWare workstation?). Did you try changing the MX Port settings for the port connected to this VM? Change it from Trunk port to Access Port (VLAN 30). Im assuming you are changing the virtual machine (the same one) from VLAN 1 to VLAN 30. If that is not the case and you have a switch port uplinked to the Meraki make sure you have presented VLAN 30 over that.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-11-2023 08:44 AM
I am using VMware workstation. With the use of vmware workstation, there is no point in changing the trunk port to access, as that would be one more reason for the connection not to work between the vlans.
I don't have any switches between the MX and the desktop, it's directly connected to the MX.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-11-2023 09:41 AM
@DscRafael it's a VMware limitation, on Microsoft hyperv it works well
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-11-2023 03:48 PM
VMware limitation? Are you sure what you're talking about? If I replicate the same vlans connected to an ASA 5505 it works correctly. It's a limitation of MX
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-12-2023 05:21 AM
Yep, I'm sure.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-12-2023 05:26 AM
So use a Cisco ASA, or open a support case.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-12-2023 06:34 AM
And just so you know, I'm talking about VMware Workstation, okay? I've already performed several network migrations with virtualized environments (VMware EXI, Microsoft HyperV, Xen, etc) and never had any problems.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-12-2023 01:22 PM
This sounds like a VM issue and not a MX problem. Everything appears fine on your MX from the screenshots and inter VLAN routing should work fine. If you put standard (not VM) clients on each VLAN and test I assume everything is ok?
I have a slightly different setup than you. But I run VM Fusion on a Mac here with two CSR1000Vs on different VLANs. Routing works fine through my MX. My Mac is on a native VLAN 172 and the two VMs are tagged to use VLAN 85 & 86. I did have to mess around with a few configs to get it working, but it's fine now.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-12-2023 04:45 PM
I found the problem.
