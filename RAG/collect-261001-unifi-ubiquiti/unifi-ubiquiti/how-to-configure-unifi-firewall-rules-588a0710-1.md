---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/how-to-configure-unifi-firewall-rules-588a0710-1
title: "how-to-configure-unifi-firewall-rules-588a0710"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: ["2025-01"]
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/how-to-configure-unifi-firewall-rules-588a0710.md
source_anchor: ""
source_lines: [1, 28]
sha256: 71a6043c3a0f4d089e7ed186625013290b58195954d36043ae4d702e6c33ae58
---

# how-to-configure-unifi-firewall-rules-588a0710

UniFi firewall rules are zone-based policies now. A UniFi firewall setup starts by putting every network, WAN interface and VPN into a zone, then writing policies that allow or block traffic from one zone to another. The old per-rule model with its LAN In, LAN Out and LAN Local types is gone once a site on UniFi Network 9.0 or newer has been migrated to the Policy Engine.
I ran zone-based policies on a test gateway while they were still a release candidate in January 2025, and my own network has run them since 2025, so the screens below are my own. The mapping from the old rule types to the new zone pairs is under UniFi LAN In vs LAN Out vs LAN Local, and What Replaced Them, further down.
Where UniFi Firewall Settings Are Now
The UniFi firewall settings live under Settings, in the Policy Engine group. Two entry points matter: Zones and the Policy Table, every policy on the site in one filterable list. You need a UniFi Cloud Gateway or UniFi Gateway running the zone-based firewall, which means Network 9.0 or newer with gateway firmware 4.1 or newer.
A zone is a group of interfaces, and six zones are built in: Internal, External, Gateway, VPN, Hotspot and DMZ. Every network belongs to exactly one zone and lands in a built-in zone when you create it. You change that either on the network itself or by editing the zone. None of the six can be deleted, and External, Gateway and VPN carry a padlock. You can create your own on top of those, up to 30.
The Zone Matrix is the part worth getting comfortable with. The rows are source zones, the columns are destination zones, and each cell shows what’s allowed for that pair. Clicking a cell filters the policy table below it to the policies governing that one flow. The cell itself reads one of four ways: Allow All, Block All, Allow Return (Ubiquiti’s docs call it Allow Return Traffic: the source can start a conversation and the destination can only reply to it), or Policies, which means a mix of allows and blocks, and that’s how most pairings with External read before you add any policies of your own.
Traffic rules aren’t a separate screen anymore. Ubiquiti pulled firewall policies, routing, QoS, NAT, port forwarding and DNS records into one Policy Engine. The built-in policies in the table (Allow All Traffic, Block Invalid Traffic, Allow Return Traffic, Block All Traffic) carry a lock icon and can’t be edited or deleted.
If your site hasn’t been migrated yet, everything below assumes you’ve clicked the Upgrade to the New Zone-Based Firewall banner first; the LAN In vs LAN Out section below has the path.
How to Configure UniFi Firewall Rules
A UniFi firewall setup is two decisions: which zone each network belongs in, and which policies you write between those zones. There’s one policy dialog now, so the Simple and Advanced split is gone.
1. Open Settings and find the Policy Engine heading in the sidebar. It’s a row of small icons rather than names: the grid icon opens Zones and the list icon beside it is the Policy Table. The grid gives you the zone list with the Zone Matrix underneath it.
2. Decide where each network belongs. Trusted devices stay in Internal, the guest network goes in Hotspot, which also switches on its landing page, and everything you don’t trust (IoT, cameras, a lab VLAN) goes in a zone you create.
3. Select Create Zone, give it a Zone Name, and add the networks that belong to it under Networks / Interfaces, then save it.
4. Look at the matrix again. The new zone reads Block All against Internal, VPN, Hotspot, DMZ and itself, and Allow All to External and Gateway, so the devices in it have internet, DHCP and DNS and nothing else until you write a policy.
5. Create the policy. Click the cell for the pair you’re configuring to filter the table to it, then select Create Policy, or open the Policy Table and create one there. Name it for what it does, set the source and destination zones, then narrow each side with the options that zone offers (Any, Device, Network, IP, MAC or Region). Set Port to Any, Specific (a Service from the list, or a port number) or List, and use the App or Web (domain) fields if you’re matching on those.
6. Pick the action. Allow passes the traffic and gives you the Auto Allow Return Traffic option (which lets the destination answer), Block silently drops it, and Reject blocks it and notifies the sender, so the client fails immediately instead of timing out.
7. Set the optional restrictions if you need them: IP Version (IPv4, IPv6 or both), Protocol, Connection State, a custom schedule, and syslog logging.
8. Save it. A custom policy sits above the built-in policies and below your other custom policies; Rule Order below covers moving it.
The policy almost every network needs is Internal to your untrusted zone, set to Allow with Auto Allow Return Traffic ticked. As soon as it’s saved, the untrusted to Internal cell flips from Block All to Allow Return, the quickest confirmation you did it right.
How to Block Traffic with a Firewall Policy
Blocking is the same dialog with the action set to Block, and there are three ways to stop one network from reaching the others. The blunt one is Network Isolation on the network itself (Settings > Networks), which automatically creates the rules to block inter-VLAN traffic in a click. It works, but it writes policies of its own alongside yours, so I turn it off when I’m managing the zones myself.
The second way is the zone itself, and it’s what I’d suggest for most people. Drop the VLANs you don’t trust into a custom zone and they’re blocked from your other networks and from each other by default, internet still working, including every VLAN you add to that zone later. The third way is a policy between two specific networks. You can do that inside a single zone as well as between zones, so two networks that both sit in Internal can still be separated with an Internal to Internal policy.
For anything narrower, open the smallest hole that does the job. When I give a camera VLAN access to my NAS, the policy names the NAS by IP and the SMB port rather than the whole zone, so those devices reach one service on one box and nothing else. Region, app and domain blocking live on the same source and destination fields. If an app or domain rule doesn’t appear to do anything, it is nearly always a policy above it matching first.
Match Opposite under the address field inverts that selection, and Match Opposite Port does the same for the port, so you can pick the one network or port you care about and have the policy apply to everything except it. I demonstrated it back in January 2025, and it still saves listing every network by hand.
Blocking traffic to the gateway itself needs a lighter touch. The Gateway zone is traffic to and from the gateway: DHCP, DNS, and the HTTPS and SSH management interfaces. I set the destination Port to List, click Create New under it and build one with 80, 443 and 22 in it, then block it from my untrusted zone to the Gateway zone, which closes the management interface and nothing else. Ubiquiti warns that blocking traffic to the Gateway zone may disrupt DHCP and DNS, and a broken DNS looks exactly like a dead internet connection. That’s the first thing I’d check if the internet seems to die right after a gateway policy.
Two things look like firewall rules but aren’t. Device Isolation (ACL) under Settings > Networks is enforced by the switches, and Client Device Isolation under Settings > WiFi stops clients on the same AP and VLAN from seeing each other. Neither appears in the policy table.
Firewall Rules for VPN and Teleport Traffic
VPN traffic has its own zone, which means UniFi VPN firewall rules are written exactly like every other policy. The built-in VPN zone covers remote users coming in over Identity One-Click VPN, WireGuard on UniFi, L2TP and OpenVPN on UniFi. It also covers site-to-site VPN in UniFi over Site Magic, IPsec or OpenVPN.
