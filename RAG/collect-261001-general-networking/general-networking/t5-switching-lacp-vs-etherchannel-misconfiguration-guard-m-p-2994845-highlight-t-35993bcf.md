---
id: collect-261001-general-networking/general-networking/t5-switching-lacp-vs-etherchannel-misconfiguration-guard-m-p-2994845-highlight-t-35993bcf
title: "of ports needs to be the same"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/t5-switching-lacp-vs-etherchannel-misconfiguration-guard-m-p-2994845-highlight-t-35993bcf.md
source_anchor: ""
source_lines: [1, 130]
sha256: 473096de6aa33cb212590c8ec63f3f9d6d0da9a3033ab2fdeb237fe39132eb4d
---

# of ports needs to be the same

LACP vs Etherchannel Misconfiguration guard?
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-30-2016 09:12 PM - edited 03-08-2019 06:49 AM
My understanding is that LACP will detect many misconfigurations on the connecting switch. Why is "spanning-tree etherchannel guard misconfig" needed then? What does one cover that the other doesn't?
- Labels:
- 
						
							
		
			Other Switching
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-01-2016 01:16 AM
LACP will pick up mis-configured speed and duplex settings between ports, and that is about it. LACP has no interest in spanning tree.
"spanning-tree misconfiguration guard" is not needed for LACP. You would use it if you wanted to protect your spanning tree root, regardless of weather you are using LACP or not.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-01-2016 05:29 AM
I'm sorry, I meant "spanning-tree etherchannel guard misconfig". I corrected my original post to clarify.
Also, I'm not aware of a "spanning-tree misconfiguration guard". Are you sure you're not thinking of root guard ("spanning-tree guard root")?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-05-2016 04:05 PM
Etherchannel misconfiguration guard or "Etherchannel Guard" is a feature that protects you when you incorrectly configure an etherchannel. So, how can you incorrectly configure an etherchannel? Well remember for an etherchannel to be configured correctly, the parameters need to be setup correctly on both ends.
Speed must be the same
Duplex must be the same
# of ports needs to be the same
1 side is set to a trunk while other side is not
Type of port (L2 or L3) need to be the same
By incorrectly configuring or not configuring one of these settings, you can run into issues. When the switch detects one of these types of errors (list above is not all inclusive) then it will place the port-channel interface and all associated interfaces into an "err-disabled" state. This is to notify the administrator that the switch protected you from configuring your etherchannel incorrectly. Hope this helps
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-05-2016 07:57 PM
Hello
My understanding is its used to protect against unwarranted mis-configuration of an aggregate (etherchannel)
 
 Example: One side is configured for ethchannel and the other isn’t /or has some configuration missing thus isn’t compatible to establish a peering and with this command the switch will error disable the port-channel.
 
 res
 Paul
Please rate and mark as an accepted solution if you have found any of the information provided useful.
This then could assist others on these forums to find a valuable answer and broadens the community’s global network.
Kind Regards
Paul
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-31-2017 11:53 PM
You are correct that LACP will detect most misconfigurations and remove links from bundling in those cases. In that case, STP will go back to regular forwarding/blocking modes.
In environments that use a mix of Spanning-Tree flavors, this feature is something to watch out for due to potential of false positives. Personally, I recommend disabling it altogether, or at least making sure auto-recovery is enabled should it kick in.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-02-2019 11:25 AM
With etherchannel misconfig guard, a switch will detect when it recieves a bpdu on a port (independently of the port channel interface) that is configured to be in an etherchannel and will place the port in a error disable state. This will happen when you explicitly set one side of an etherchannel to manual.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-06-2019 09:36 AM
You can use EtherChannel guard to detect an EtherChannel misconfiguration between the switch and a connected device. A misconfiguration can occur if the switch interfaces are configured in an EtherChannel, but the interfaces on the other device are not. A misconfiguration can also occur if the channel parameters are not the same at both ends of the EtherChannel. If the switch detects a misconfiguration on the other device, EtherChannel guard places the switch interfaces in the error-disabled state, and displays an error message.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-24-2021 01:49 PM
hi Esa_fresa,
-------------
LACP will detect only consistency problems between two switch regarding the PO creation. It will not perform configuration verification.
Example:
if I try to connect two interface (binded into a PO in mode LACP) with two single interfaces or two interfaces binded with a different LAG protocol, the LACP port channel not come UP (this is a consistency check)
but
if I try to connect two interface (binded into a PO in mode LACP) configured in mode trunk on SW-1, with two interface (binded into a PO in mode LACP) configured in mode switchport on SW-2, all PO come UP correctly. (LACP not perform config. verification)
-------------
ETHERCHANNEL MISCONFIG GUARD is an STP feature that protect the network environment from possible loop situation, this feature helps when a PO is in MODE ON (static binding)
Example:
if I try to connect two interface (binded into a PO in mode ON) on SW-1, with two single interfaces on SW-2, ETHERCHANNEL MISCONFIG GUARD put SW-1 interfaces in errdisable, in order to avoid loops.
