---
id: collect-261001-cisco/cisco/questions-2683-problems-connecting-catalyst-2960g-to-netgear-prosafe-via-fibre-0082c955
title: "questions-2683-problems-connecting-catalyst-2960g-to-netgear-prosafe-via-fibre-0082c955"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/questions-2683-problems-connecting-catalyst-2960g-to-netgear-prosafe-via-fibre-0082c955.md
source_anchor: ""
source_lines: [1, 3]
sha256: b0448d1bcb745b8d7f7360a78bbc893e6d82d28bd6a4bbc816bb546e239299e8
---

# questions-2683-problems-connecting-catalyst-2960g-to-netgear-prosafe-via-fibre-0082c955

I've got a catalyst 2960G, connected to a Netgear Prosafe GS718T switch (cheap manageable switch). Idea is I have a couple of users that need access to a test environment - hence the cheap switch. It's connected via fibre through a 1Gbit SFP over multimode, layer 1 looks ok, interface on the catalyst and the netgear are both up/up, but I can't get any traffic over it.
The int on the 2960G is configured as default except it's in access mode (tried it with DTP on - doesn't work) on a specific VLAN, and the netgear is in it's default config, so it should just be acting as a dumb switch with no tagging on the link to the cat. However I do a #sh mac address table on the cat and it doesn't return anything against the interface facing the netgear, and I try to ping to a good IP at the cat end from the netgear end and get nothing. Again interfaces are all up/up and the cat isn't blocking multiple mac addresses of one port (ie no port-security). 
Anyone else tried to use on of these Netgear GS switches with a cisco switch or similar over it's fibre int and got it working? The web interface is horrible, been trying to see if I can do anything in it, really want to avoid trunking it if I can as it's not needed and it looks like a real hassle on the netgear. Ideally I'd be using another catalyst but we've got none spare, and this netgear is meant to support all the ethernet standards.
