---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/how-to-configure-unifi-firewall-rules-588a0710-2
title: "how-to-configure-unifi-firewall-rules-588a0710"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: ["2025-01"]
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/how-to-configure-unifi-firewall-rules-588a0710.md
source_anchor: ""
source_lines: [29, 62]
sha256: 659c75e4a63c3068b25aa5e2970d0de28f37dc026a926a4ead68e39ab3f3fc64
---

# how-to-configure-unifi-firewall-rules-588a0710

The default matters here more than anywhere else on the page. VPN to Internal is Allow All, and so are VPN to Gateway, Hotspot and DMZ, so a remote client that connects has the run of the network until you write a policy that says otherwise. To limit it, create a policy with the source zone set to VPN, the destination zone set to Internal and the action set to Block, then pick the networks the VPN shouldn’t reach. It lands above the built-in Allow All on its own, so there’s nothing to reorder.
If you’ve read anything written before this (the earlier version of this article included), you were told that VPN traffic needs a LAN Out rule, because the VPN server runs on the router itself rather than behind it. That is finished. As I said in January 2025, you no longer have to remember that it’s LAN Out for VPNs and LAN In for internal.
Teleport VPN on UniFi is the odd one out. It’s Ubiquiti’s zero-configuration VPN, built on WireGuard and turned on under Network Settings > VPN, where an invitation works for a single device and expires after 24 hours. Ubiquiti doesn’t publish which zone a Teleport client lands in, so test any policy you write for it against a connected Teleport client before you rely on it.
Automatic Firewall Rules for UniFi Networks
Policies get written for you in a few cases too. Setting up a VPN server, a port forwarding rule or IPTV streaming all generate additional policies automatically, below whatever you write yourself, and they only show in the Policy Table when View Default Policies is ticked.
Rule Order: How Policies Are Evaluated
Policies are evaluated from top to bottom and the first one that matches wins, which is the one thing that hasn’t changed from the old model. Anything you haven’t matched falls through to the built-in policies at the bottom. The list under the matrix, not the Policy Table, is the one shown in evaluation order, top to bottom. So the practical rule is to put the allow for the traffic you want above the block that denies the rest.
By default a custom policy takes precedence over the built-in ones but follows your other custom policies. You change that from Zones: click the zone pair in the matrix to filter the policies below, then click Reorder under that list. The Policy Table won’t do it for firewall policies. The built-ins themselves (things like “Allow Neighbor Advertisements”) carry a lock icon and can’t be edited or deleted, so your own policy beats one simply by sitting above it, which it already does.
When a rule appears to do nothing, it’s almost always one of three things. The first is return traffic. If you allowed A to B without Auto Allow Return Traffic, B can’t answer and the connection fails, unless another policy already allows the return traffic. The second is direction: a block from zone A to zone B holds even when B to A is allowed. The third is that something above it matched first.
UniFi LAN In vs LAN Out vs LAN Local, and What Replaced Them
The old model asked you to pick a rule type: LAN In caught traffic on its way into the firewall, LAN Out caught traffic heading out to the LAN, and LAN Local only traffic aimed at the firewall itself. Internet In, Internet Out and Internet Local did the same job for WAN traffic (Ubiquiti’s migration table below calls them WAN_IN, WAN_OUT and WAN_LOCAL), with a duplicate set for IPv6 on top of that.
None of those types exist after migration. The direction is the zone pair instead, so you pick a source zone and a destination zone, and one policy can cover IPv4, IPv6 or both. When I first covered the zone-based firewall I said that new users found the difference between LAN In and LAN Out confusing, especially with IPv4 and IPv6 to keep straight, and that those days were gone. Having run it since, that held up.
If you want to know where your old rules went, Ubiquiti published the mapping it uses when it migrates a site. It’s the expansion the migrator runs, one policy for every zone pair in the row, which is why migration multiplies your rules. Building one by hand, keep the single source and destination the traffic really used, and read Internal as whichever zone that network is in now.
| Legacy ruleset | Source zone | Destination zone | 
|---|---|---|
| LAN_IN | Internal | Internal, Hotspot, External, VPN | 
| LAN_OUT | Internal, Hotspot, External, VPN | Internal | 
| LAN_LOCAL | Internal, VPN | Gateway | 
| GUEST_IN | Hotspot | Internal, Hotspot, External, VPN | 
| GUEST_OUT | Internal, Hotspot, External, VPN | Hotspot | 
| GUEST_LOCAL | Hotspot | Gateway | 
| WAN_IN | External | Internal, Hotspot, External, VPN | 
| WAN_OUT | Internal, Hotspot, External, VPN | External | 
| WAN_LOCAL | External | Gateway | 
If your site hasn’t been migrated yet, nothing happens until you open Settings > Security > Traffic & Firewall Rules and click the Upgrade to the New Zone-Based Firewall banner across the top of it; the move is never done for you. The move takes a few seconds and traffic keeps passing while it runs. An automatic backup is taken first, and the notice that appears after the migration carries the link to restore the previous configuration.
Expect more rules afterwards, not fewer. Ubiquiti took a conservative approach so behaviour is identical after the move, which means it generates policies that are redundant or do nothing. It tells you to test first, then remove the ones doing nothing; they’re listed in the Policy Table. On the test gateway I migrated in January 2025, a single LAN In rule blocking IoT to the RFC1918 range came out the other side as four separate policies, so the cleanup is worth an evening once you’re confident the network still behaves. Your old IP and port groups survive as Objects, the cube icon in the Policy Engine group, and the Simple and Advanced rule screens are gone.
UniFi Firewall Rules Best Practices
The UniFi firewall rules best practices I keep coming back to are mostly about zone layout rather than clever policies, because the layout is what decides your defaults. The first is that the Internal zone is sacred: anything you put in it can reach everything else in it, so it’s for devices you genuinely trust. The second is to group VLANs into a few zones rather than creating one zone per VLAN. Networks inside a custom zone can’t talk to each other by default while networks inside Internal can talk to everything, and that one difference does most of your segmentation for you.
In my opinion, two custom zones cover 95 to 98% of what people actually need: one untrusted zone with internet access and one without, the second being the same kind of zone with a block policy to External. IoT gear and cameras go in whichever of the two fits. Then you write the few policies that punch holes for the traffic that has to cross.
- Put the guest network in the Hotspot zone rather than Internal, then uncheck Show Landing Page under the gear icon in Insights > Hotspot and check the captive portal is off on the WiFi network itself. Here’s the full guest network VLAN on UniFi setup.
- Park the default network in a custom zone with a block policy to External if you don’t use it, because any untagged switch port lands a device there.
- Consider both directions. Allowing one direction does not open the other. Auto Allow Return Traffic covers replies, not new connections from the far side.
- Don’t write an allow policy you don’t need. A single VLAN or a single port beats a whole-zone allow.
- Deleting a custom zone deletes every policy attached to it, so check what goes with it before you remove one.
