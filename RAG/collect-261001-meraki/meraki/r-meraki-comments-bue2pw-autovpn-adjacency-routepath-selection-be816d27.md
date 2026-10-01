---
id: collect-261001-meraki/meraki/r-meraki-comments-bue2pw-autovpn-adjacency-routepath-selection-be816d27
title: "r-meraki-comments-bue2pw-autovpn-adjacency-routepath-selection-be816d27"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-bue2pw-autovpn-adjacency-routepath-selection-be816d27.md
source_anchor: ""
source_lines: [1, 9]
sha256: e3e45a2200a37aac5ac1ca5111a921d6719441f0bfbb42d0a7cc9837a5b3ed46
---

# r-meraki-comments-bue2pw-autovpn-adjacency-routepath-selection-be816d27

Auto-VPN Adjacency Route/Path Selection
Hey all, I'm piloting the use some MX's here and I'm trying to get some clarification on something I was told from support. I was told that AutoVPN will always try and use the publicly addressable interface of Auto-VPN Peers to form an adjacency. However, I'm not finding this to be true. Looking at the diagram below, assuming all MX's are acting as One-Arm Concentrators...
The resulting flow is that the MX's traverse tunnel traffic between their private interfaces using the Cloud Connect Circuit as the path. This is actually desirable, I just thought I would need an ACL to block the public IP's to force this to happen. I want to make sure I understand why that's not needed. Thanks so much for any clarification you can offer!
EDIT: Just to be clear, in the example everything in Azure has routes for everything on Prem and vice versa. Also, is there a way to tell in the portal what path is being taken? The only thing I can find to do is perform a packet capt to see if the pub or priv IP is being used.
Section des commentaires
Auto VPN will use both the public IP address it has and also the private IP it has (for MPLS lines). If you take a capture on the uplink of the MX you should see UDP peering data going out for both the private IPs and public IPs of the other MX devices. You can filter using the UDP port that each MX is dynamically. You can find the port on the VPN status page and look at what port its using
Thanks for the response. I did see that a capture showed connections to both the public and private IP's over the ESP encapsulated UDP port. My questions is how the MX handles preference/priority. I'm seeing it prefer the inside route rather than the outside route while I was told by the trial engineer the reverse would be true.
Also, is the behavior the same when the MX is acting in Routed Mode?
It will always prefer the private IP first. This is due to how it was designed to form its adjacency over an MPLS link. Then if it fails to peer over its private IP, it uses the Public IP. This works the same for both routed and one arm mode
