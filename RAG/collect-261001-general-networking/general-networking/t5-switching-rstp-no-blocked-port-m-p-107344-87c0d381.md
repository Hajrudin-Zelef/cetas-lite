---
id: collect-261001-general-networking/general-networking/t5-switching-rstp-no-blocked-port-m-p-107344-87c0d381
title: "t5-switching-rstp-no-blocked-port-m-p-107344-87c0d381"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-rstp-no-blocked-port-m-p-107344-87c0d381.md
source_anchor: ""
source_lines: [1, 110]
sha256: 98eca60081d7e084eda17b863bc1f993e7e49b08e02219c2b46ad46883ea8373
---

# t5-switching-rstp-no-blocked-port-m-p-107344-87c0d381

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-20-2021 11:04 AM
Hi all,
I have a constellation in which 4 switches are connected in a ring - but there is no loop although on a Meraki switch port, RSTP is deactivated on the interface! No idea but maybe the different native vlan on the trunk plays a role?
thanks in advance for any help!
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
01-26-2021 12:13 PM
In that case you don’t have a loop
On the left-hand side of the diagram you only have VLANs 2 and 3, in the top right segment you have all VLANs on the wire, and in the bottom right you only have VLAN1 (since it’s an Access Port in VLAN1 it only carries VLAN1). 
RSTP (if enabled) runs on a port regardless of whether it’s an access port or a trunk port. The BPDUs it sends are always un-tagged, but it’s not part of a VLAN, although it does impact all VLANs. In your scenario, if RSTP was enabled on the port you have indicated it is disabled on, it would likely go into blocking state on one of the uplinks of the bottom switch; this is even though there isn’t an actual Layer 2 loop, there is a loop of ‘un-tagged’ traffic.
I wouldn’t call it good practice, but it is working as would be expected in this scenario.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-20-2021 11:48 AM
As per below, if the port has RSTP disabled then it won’t participate in STP. It’s this config that’s stopping the loop
| Option | Description | 
| Enabled | RSTP must be enabled globally (see "Enable RSTP Globally") for any ports to be able to participate in Spanning Tree processes. When RSTP is enabled globally, RSTP will be enabled at the port level by default. A disabled port can be re-enabled by selecting Enabled. While RSTP is enabled on a switch port, that port is able to participate in Spanning Tree processes. It is recommended that RSTP be enabled on all ports. | 
| Disabled | RSTP may be disabled at the port level. Disabling RSTP on a port removes the port from any STP processing including any STP guard configuration. Disabling RSTP on a port is not recommended unless the client device connected to the port is incompatible with STP. If RSTP is disabled globally, all ports will have RSTP disabled and cannot have it enabled. | 
https://www.linkedin.com/in/darrenoconnor
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-20-2021 10:28 PM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-25-2021 04:57 AM
someone any ideas on that?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-25-2021 01:52 PM
Since only VLANs 2 and 3 are allowed on the trunks on the left-hand side of the diagram, a loop has definitely been formed for these two VLANs based on the diagram. Are these VLANs in use? Are these any other ports on the four switches that RSTP has put into blocking state? What model switches are they, has someone enabled ‘broadcast storm controls’ to make the network workable (although obviously not correct)? Are there any messages in the Event Logs?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-26-2021 07:02 AM
I`ve just looked at the configuration on the bottom switch again and I`ve to say that I made a mistake in the sketch! the port (towards where RSTP is disabled) is an access port untagged in vlan 1
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-26-2021 12:13 PM
In that case you don’t have a loop
On the left-hand side of the diagram you only have VLANs 2 and 3, in the top right segment you have all VLANs on the wire, and in the bottom right you only have VLAN1 (since it’s an Access Port in VLAN1 it only carries VLAN1). 
RSTP (if enabled) runs on a port regardless of whether it’s an access port or a trunk port. The BPDUs it sends are always un-tagged, but it’s not part of a VLAN, although it does impact all VLANs. In your scenario, if RSTP was enabled on the port you have indicated it is disabled on, it would likely go into blocking state on one of the uplinks of the bottom switch; this is even though there isn’t an actual Layer 2 loop, there is a loop of ‘un-tagged’ traffic.
I wouldn’t call it good practice, but it is working as would be expected in this scenario.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-02-2021 02:51 AM
@BRUCE NEWTON thank you man for that detailed explanation! I appreciate that help on this topic very much!
Just for my better understanding - when RSTP is disabled on a switchport, does this automatically mean that there is some sort of BPDU-Filter active?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-02-2021 02:59 AM
When RSTP is disabled on a port, the port takes no part in spanning-tree, so it always passes traffic. It won’t send BPDUs, and it won’t process incoming BPDUs (so essentially they get dropped). The recommendation is generally not to disable RSTP as if you do get a loop it will take down your network.
