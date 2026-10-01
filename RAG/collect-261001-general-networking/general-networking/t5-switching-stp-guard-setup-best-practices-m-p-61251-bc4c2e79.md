---
id: collect-261001-general-networking/general-networking/t5-switching-stp-guard-setup-best-practices-m-p-61251-bc4c2e79
title: "t5-switching-stp-guard-setup-best-practices-m-p-61251-bc4c2e79"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-stp-guard-setup-best-practices-m-p-61251-bc4c2e79.md
source_anchor: ""
source_lines: [1, 161]
sha256: f4a031f906edca4bd52a7f72504d1f6e91f4072486cda55a150075ad4052f281
---

# t5-switching-stp-guard-setup-best-practices-m-p-61251-bc4c2e79

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-01-2018 10:34 PM
Curious what the consensus is on STP guard settings for ports on Meraki switches. We've turned on BPDU guard for all access ports. However, I was wondering under what circumstances Root or Loop guard would be used. We have a few 3rd party switches uplinked to some of our Meraki switches (trunk ports). Would Root or Loop guard be worthwhile to activate?
The same question goes for fiber uplinks - from Meraki switches to a core. Is there a best practice on what STP guard settings should be? Or is "disabled" the norm?
Thanks for your input. Happy to provide more topology details if need be.
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2019 11:38 AM
Thanks for bringing this question forward! We have published new documentation on STP guard configuration that incorporates STP guard recommendations. Check it out and let us know what you think!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2019 01:11 PM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2019 01:39 PM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2019 02:32 PM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2019 01:36 PM
CJones - Thank you for posting this. It is very helpful.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-03-2019 07:04 AM
Redsector had a great answer.
However that said, I don't use any of these settings because the Meraki already has RSTP on by default. I definitely don't configure them on Meraki-Meraki links because the expectation is to use RSTP. In my mind you should only use these spanning tree options if the port is connected to a switch that doesn't support RSTP.
As for root guard - I set the priority on my core switches, a stack of MS425s, to 0, and that stopped the inter-vendor squabbling over who thinks it's root.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-07-2019 02:16 PM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-06-2019 11:11 AM
I have my root priority set to the core switch in the network but a few locations are still somehow wanting to use a core switch in another network.
Would that be something to be concerned with?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-09-2019 11:08 AM
I always do the following:
- BPDU guard on all client ports and access point ports if they are Meraki (Meraki AP's don't send BPDU's).
- Root guard on all downlinks from CORE to access layer
- I would have wanted to put loopguard on uplinks of access layer switches but Meraki won't let me because we use the management inline with the network.
Also and this is important. If you have a MX warmspare with the four uplinks from the switch network towards those MX's that you DON'T enable bpdu guard on those ports leading to the MX and never ever use drop untagged traffic on the MX because that causes a loop.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-04-2020 03:59 PM
@joey.debraI see you mentioned MX upstream in HA mode. I have a MS225 switch stack of 4 switches with 4 uplinks going up from the the stack to the MX84s. All 4 trunks have loopguard enabled. Is this what you were suggesting with your comment about 4 uplinks to a warm spare MX? Or anybody else on the thread. Is this what you would recommend for a HA setup? I guess I could have burned more ports to the other two switches but seemed a little like overkill. **Forgot to mention that both trunks on SW1 are forwarding and both trunks on SW2 are blocking.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-05-2020 10:05 AM
Yes your picture is completely correct in terms of HA setup and STP behavior.
In a stack only have four uplinks where two will be blocked because the BPDU's leave 1/0/1 and 1/0/2 with a lower Port-ID than 2/0/1 and 2/0/2 and the MX just forwards the STP messages since it does not participate in it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-21-2020 10:19 PM
May you please elaborate on the exact STP port related settings for the uplinks. I have the exact setup but with 2x MS120 (without a stack). I have tried almost any STP configuration on my 5 uplinks (including the one between the switches) but Im still getting a loop when I add the 2nd link from the firewalls.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-22-2020 07:44 AM
Now granted this is my lab at the moment but all of my switch trunks from my diagram are set like this. My MX is set to a trunk on native vlan 99, but I will eventually prune back the vlans back to only my vlans for the site.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-23-2020 05:14 AM
@deguc004, even without stacked switches it should work.
Since packets are bounced through the MX'es you should NOT enable any features like rootguard, loopguard or bpduguard on the links towards the MX'es.
The direct link between your switches can have root guard enabled in case of non-stack.
Make sure you don't have a link between your MX'es and you allow the same VLANs across the 4 links including the correct native.  DO NOT USE drop untagged!.  You can use a different VLAN native config between the switches.
Ensure one of your MS120's is root bridge for the L2 network.
You should see two links blocked on the switch that is not root bridge.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-24-2020 09:05 AM
