---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/how-to-port-forward-on-unifi-devices-488a0645-2
title: "how-to-port-forward-on-unifi-devices-488a0645"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["SpaceX"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/how-to-port-forward-on-unifi-devices-488a0645.md
source_anchor: ""
source_lines: [33, 46]
sha256: 149696677ddec450c780021876a7e5453a67a47475935c2ca087b8783fe3d514
---

# how-to-port-forward-on-unifi-devices-488a0645

- Direct Remote Connection. This lives in the console settings under Control Plane, and it exposes port 443 on the gateway itself, so your router’s login page answers to anyone who hits your public IP. This should be off. If you want to manage UniFi remotely, use Ubiquiti’s built-in remote access through the UniFi site manager, which doesn’t need an open port.
When Port Forwarding Doesn’t Work
If the rule is in and the service still isn’t reachable from outside, it’s almost always one of these, in this order:
- You don’t have a public WAN IP. If your ISP puts you behind carrier-grade NAT (common on cellular, fiber in some regions, and Starlink), or your UniFi gateway sits behind another router that’s still doing NAT, port forwarding on the UniFi side can’t work. Check the WAN address on the Internet page: anything starting with 100.64 through 100.127, or a private address like 192.168 or 10.x, means you’re behind another NAT. Double NAT is fixable by putting the upstream device in bridge mode. CGNAT is not, and that’s when you want Teleport or a Cloudflare tunnel instead, since neither needs an open port.
- The service isn’t listening, or its own firewall is. Confirm you can reach it from another device on the LAN using the forward IP and port first. If that fails, the problem is on the device, not the gateway. Windows Firewall and Docker port mappings are the usual suspects.
- Wrong protocol or wrong WAN. A UDP service forwarded as TCP does nothing, and a rule tied to WAN2 on a single-WAN setup does nothing either.
- You’re testing from inside. Test from mobile data. If it works from outside and not from inside on your public IP, that’s a NAT reflection quirk and not a broken rule, and the fix is to use the local address when you’re home or set up local DNS for it.
- Your IP changed. Residential connections don’t keep the same public IP forever. Set up dynamic DNS on your UniFi gateway so people connect to a hostname instead of a number that moves.
The Alternatives I’d Try First
If you’re port forwarding to reach your own network from outside, stop and set up a VPN instead. UniFi makes it about as easy as it gets. WireGuard is the one I run and recommend, OpenVPN is there if a client needs it, and Teleport is Ubiquiti’s own option that works even behind double NAT because it doesn’t need a port at all. A VPN does open one port for itself, but everything behind it requires a key, which is a completely different security posture from a service answering to the world.
If you’re exposing a web service to other people, look at a Cloudflare tunnel before you forward anything. It is free for this, it needs no open port, and Cloudflare sits in front of your service taking the abuse. The catch is that a tunnel set up on your main LAN gives Cloudflare a path to everything on that LAN, so the right way is to run it on an isolated VLAN in the DMZ zone, which is how I set it up in my Cloudflare Tunnels and UniFi firewall video. The tunnel itself is a single container, and my Cloudflare tunnel guide covers that part. Check that what you’re exposing is allowed under Cloudflare’s terms before you rely on it, because media streaming is the case people get wrong most often.
And if you’re doing any of this on an older USG or a non-UniFi router in front of your UniFi gear, the zone-based firewall, IDS/IPS, and the one-click VPNs above all need a current UniFi gateway. The UniFi Express 7 covers a normal home for around $200, and the Cloud Gateway Fiber is the step up for multi-gig connections. I compared the whole lineup in my best UniFi router guide.
Final Thoughts
UniFi port forwarding is a short form in the Policy Table, and the form is the easy part. The work is everything around it: a fixed IP on the target, the From field or a country policy to shrink who can reach it, the device on its own VLAN in the DMZ, IDS/IPS set to block, UPnP and Direct Remote Connection off, and an occasional look at the Port Forwarding filter to make sure the list still matches what you meant to expose. Do that and a forwarded port is a manageable risk instead of an open door. And keep in mind that security changes constantly and none of this guarantees anything, so audit that list every so often, especially after you’ve added a new device.
