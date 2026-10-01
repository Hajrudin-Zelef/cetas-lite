---
id: collect-261001-general-networking/general-networking/questions-52308-will-all-switches-pass-802-1x-eap-seamlessly-any-caveats-616d7a9e
title: "questions-52308-will-all-switches-pass-802-1x-eap-seamlessly-any-caveats-616d7a9e"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-52308-will-all-switches-pass-802-1x-eap-seamlessly-any-caveats-616d7a9e.md
source_anchor: ""
source_lines: [1, 13]
sha256: 8e38705100fcf681ecb3c5a7b510350278afa18858f7fc97c03f81a7d90bfb77
---

# questions-52308-will-all-switches-pass-802-1x-eap-seamlessly-any-caveats-616d7a9e

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
We have a selection of Unifi, Meraki and Arista switches. We're currently trying to sort out our 802.1x/RADIUS infrastructure
It seems like Meraki has the best 802.1x support (e.g. supporting multi-auth, supporting MAC-address bypass etc.) - is it possible to have the Meraki switches upstream, and then have the other switches downstream pass on the EAP frames seamlessly? Are there any caveats we should be aware of?
I should disable 802.1x on all other downstreams ports and leave them open?
802.1X EAP frames are supposed to only be used between the client (supplicant) and its uplink switch (authenticator). The authenticator then uses a higher layer protocol for communication with the authentication server. The higher layer protocol can be switched and routed as desired.
You cannot use 802.1X authentication across several switches. Essentially, this would only authenticate the downlink port on the authenticating switch - this inherently authenticates all ports further downstream.
Additionally, you shouldn't use 802.1X between switches. Use the edge switch to authenticate edge ports and trust your interlink ports.
Assuming you are trying to use 802.1x to allow or prevent communication by a client through a given switch port (i.e. Port Access Control as defined in the 802.1x standard) the scenario you describe doesn't work. The standard explicitly states that : "The operation of Port-based Access Control assumes that the Ports on which it operates offer a point-to-point connection between a single Supplicant and a single Authenticator", which is not the case if you try to use an upstream switch to deal with the access negotiation.
