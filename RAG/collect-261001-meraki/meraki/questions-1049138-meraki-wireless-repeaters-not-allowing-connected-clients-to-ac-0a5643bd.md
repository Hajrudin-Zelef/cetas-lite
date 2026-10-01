---
id: collect-261001-meraki/meraki/questions-1049138-meraki-wireless-repeaters-not-allowing-connected-clients-to-ac-0a5643bd
title: "questions-1049138-meraki-wireless-repeaters-not-allowing-connected-clients-to-ac-0a5643bd"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-1049138-meraki-wireless-repeaters-not-allowing-connected-clients-to-ac-0a5643bd.md
source_anchor: ""
source_lines: [1, 2]
sha256: 6e9797a01962b4658347237ff127e53f81e76f60fbd6ef543a43c30b01520d94
---

# questions-1049138-meraki-wireless-repeaters-not-allowing-connected-clients-to-ac-0a5643bd

I am trying to set up a Meraki Mesh network(Mesh1) at a small remote office. I have 1 MX65W and 3 MR36's, 1 acting as the gateway and the other 2 as repeaters. The SSID I'm broadcasting is set to operate in bridge mode, and it's connected to a VLAN that is connected to a VPN tunnel back to the main office network(Corp1). From my laptop on Corp1, I can ping devices that are connected directly to the MX65 via a wired connection, and I can ping devices that are connected wirelessly to the Gateway AP. However, I can't ping devices that are connected to either of the 2 repeaters.
The 3 AP's appear to have Meshed successfully and devices connected to the repeaters can still access the internet and are getting IP's in the correct subnet, they just can't seem to communicate with anything on Corp1. Is there some kind of security setting I'm missing? How do I get the repeaters to communicate with the corporate network like the Gateway is already doing?
