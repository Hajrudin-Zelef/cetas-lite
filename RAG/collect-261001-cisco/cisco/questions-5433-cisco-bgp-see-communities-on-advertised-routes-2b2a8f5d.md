---
id: collect-261001-cisco/cisco/questions-5433-cisco-bgp-see-communities-on-advertised-routes-2b2a8f5d
title: "questions-5433-cisco-bgp-see-communities-on-advertised-routes-2b2a8f5d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-cisco/questions-5433-cisco-bgp-see-communities-on-advertised-routes-2b2a8f5d.md
source_anchor: ""
source_lines: [1, 29]
sha256: 5d9692cca05f9000391c3580482ec68b8c65ef2682fd485ca69602b2c3f505b0
---

# questions-5433-cisco-bgp-see-communities-on-advertised-routes-2b2a8f5d

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
14
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
We have an outbound route-map on an eBGP peer that adds some communities to certain prefixes. Is there any way to prove that this is actually happening from the sender side in classic IOS?
show ip bgp neighbor x.x.x.x advertised-routes does not show them
Even debug ip bgp update out doesn't seem to!
In gns3 I can make the same config and see that it works from my fake upstream side, but I need to be able to verify in the production router from the local side...
Last time I checked/tried to do this, it wasn't possible in IOS. I'm not sure about NX-OS or IOS-XE/XR though. It doesn't help you, but this is possible to do on Juniper gear.
Regarding your configuration - make sure you're using the additive keyword in your route-maps to set the communities, otherwise you will not be adding communities to the list, but rather replacing any existing communities with the one you're setting in your route-maps.
Set up another BGP peer or route-reflector that you control with identical policies . Of course, make sure you're not blackholing any traffic toward this peer.
On router running IOS-XR you can use "show bgp advertised neighbor A.B.C.D or X:X::X",
it will give you additional info about the prefixes sent to the neighbor
(including of course community if configured).
As stated before remeber to use the "additive" keyword inside the route-map statement if you don't want to replace the whole list
and add the "send-community-ebgp" under the right address-family in order
to start sending it to the neighbor
I've been able to successfully confirm BGP attributes are being applied by debugging BGP and then resetting the peer and then watching the debug messages - It is of course disruptive. After you did your debug ip bgp update-out, did you do a hard reset on the peer? I will lab this up and provide a screenshot.
BGP routing table entry for 0.0.0.0/0, version 5
Paths: (1 available, best #1, table default, not advertised to EBGP peer)
Not advertised to any peer
Refresh Epoch 1
65000 65001
1.1.1.1 from 1.1.1.1 (1.1.1.1)
Origin IGP, localpref 100, valid, external, best
Community: no-export
rx pathid: 0, tx pathid: 0x0
