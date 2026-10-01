---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/bruno-roquette-site-to-site-wireguard-between-two-unifi-gateways-i-measured-wher-08f16f3a-2
title: "bruno-roquette-site-to-site-wireguard-between-two-unifi-gateways-i-measured-wher-08f16f3a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-unifi-ubiquiti/bruno-roquette-site-to-site-wireguard-between-two-unifi-gateways-i-measured-wher-08f16f3a.md
source_anchor: ""
source_lines: [90, 141]
sha256: 89d1aa016c59dc349c28bc2b0cfdfaf720c57db0a6050b2968d32778b2b5ec3d
---

# bruno-roquette-site-to-site-wireguard-between-two-unifi-gateways-i-measured-wher-08f16f3a

- ListenPort in the client’s.conf is accepted and persisted, but the port is never opened. Two clients dialing each other will never handshake. This one is especially expensive because nothing errors.
- A VPN client lands in the External zone and the API refuses to move it (NetworksCantBeReassignedToDifferentZone ). That’s not a bug — the tunnel is modeled as a virtual WAN.
- Don’t “fix” that by opening External -> LAN . That zone contains your actual WAN interfaces. Qualify by source with an address group instead.
- create_allow_respond is rejected onExternal -> Gateway (FirewallPolicyCreateRespondTrafficPolicyNotAllowed ).
- Creating a WireGuard peer isn’t automatable. I swept 21 REST collections and no peer’s public key appears in any of them. v2/.../vpn/users shows them, but it’s a status endpoint — read-only in practice.
And one reading trap that will silently waste your afternoon: a zone PUT signals success by returning the object (with _id), not by returning meta.rc. Code that checks meta sees “fine” on a call that changed nothing.
Part 3 — Site Magic / SD-WAN: it works, and here’s the bill
Four fields on one cloud page, and the exact test that failed above passes:
-> remote gateway   HTTPS   200 in 34 ms
-> remote host :8899        200 · 1422 bytes · 31 ms
That’s the honest headline: the official path works and it is fast. The cost is in four places nobody warns you about.
1. The UCG Ultra cannot be a Hub. The screen says “No sites are available to be set as Hubs” and offers no explanation. The documentation lists who can: EFG, UDM Pro Max, UDM SE, UDM Pro, UCG Fiber, UCG Industrial, UDW. So you use Mesh, which with two sites behaves identically — except it does not expose the VPN Tunnels field, which is what you’d use to pin a tunnel to a specific WAN on a dual-WAN site.
2. “Enable Subnet Overlap with SNAT” ships ON. It rewrites the remote site into 172.16.x. If your subnets don’t overlap — mine deliberately don’t — this throws away your entire addressing scheme for nothing. Uncheck it.
3. The cloud lags behind a WAN change. Until it catches up, the site shows greyed out and Mesh refuses to include it, without saying why.
4. It depends on Ubiquiti’s cloud. That’s the actual trade: the work disappears, the dependency appears. Whether that’s acceptable is a judgment call, not a technical one — but you should make it knowingly.
Where the firewall applies: both ends
The zone-based firewall documentation says the destination gateway evaluates tunnel traffic. That’s true, and it’s half the picture. The source gateway filters too.
⚠️ I concluded the opposite and I was wrong. I fired 25 pings, watched the destination counters stay at zero, and wrote down “the filter only happens at the source.” Then a real VNC session ran across the tunnel and the same counters moved immediately:
<zone-policy>-to-trusted     91,340 hits
<zone-policy>-to-gateway      4,159 hits
A flat counter under light traffic proves nothing about whether a rule is live. I’d been treating absence of evidence as evidence, on a sample far too small to say anything.
The clean way to isolate the source side — same machine, same cable, only the VLAN changes:
source VLAN        reaches other site?  internet?
-----------------  -------------------  -------------------------------
the permitted one  [ok] everything      [ok]
any other          [NO] nothing         [ok] (the experiment's control)
That control column matters. Without it, “nothing works” is indistinguishable from “I unplugged something.”
Part 4 — WAN failover: 78 seconds, unattended
I swapped the two ISPs between the sites — meaning both public IPs changed at the same moment, which is harsher than any real-world outage:
tunnel down     ~78 s, recovered on its own
back up with    200 in 33 ms, 0% loss
DDNS            followed
Dynamic WAN doesn’t break SD-WAN. The two ends re-find each other through the cloud, not by address — which is precisely the dependency you’re paying for. Note what this means both ways: you get resilience you didn’t configure, and you get it because a third party is brokering the rendezvous.
⚠️ Two things will stall the swap if you don’t plan for them: PPPoE must be reconfigured on the receiving gateway, and an ISP that requires MAC cloning will not hand out an address without it. Both cost me a bench outage before I learned to stage them first.
One last trap that isn’t a network problem
Once SD-WAN links your sites, a phone’s WireGuard profile actively hurts when you’re physically at one of them. Its AllowedIPs covers both networks, so local traffic hairpins out to the far site and back. The symptom is that the phone reaches nothing while everything else on the same Wi-Fi is fine. Turning the VPN off fixes it instantly, and it looks like a catastrophic failure for about five minutes first.
Verdict
                             manual WireGuard                             Site Magic / SD-WAN
---------------------------  -------------------------------------------  ---------------------
gateway<->gateway both ways  [NO] no - client NAT, unfixable from the UI  [ok] yes
independent of vendor cloud  [ok]                                         [NO] no
survives both WANs changing  -                                            [ok] 78 s, unattended
honest effort                days                                         minutes
If you need two UniFi gateways to route to each other in both directions, Site Magic is currently the only path, and the price is a cloud dependency. The manual route isn’t hard — it’s closed, at the NAT on the VPN client interface, and no amount of firewall or route work opens it.
That the workaround exists as an undocumented checkbox below the fold of a page you can only reach by search, and still isn’t sufficient, is the part I’d like Ubiquiti to fix. Three years and 196 replies on the feature request suggest I’m not the first to ask.
References
- Feature Request: UniFi WireGuard Site-to-Site — 196 replies over three years, no official response. Contains the SSH workaround (wg showconf /wg setconf ) and the warning that any UI change wipes it
- Add Feature: WireGuard AllowedIPs in the GUI for Client — the missing field, requested separately
- Unifi Site to Site WireGuard VPN (without Site Magic) — Lawrence Systems forum; ends without anyone demonstrating working bidirectional traffic
- Zone-Based Firewalls in UniFi — official
- Setting Up SD-WAN with UniFi Site Manager and Fabrics — official; this is where the Hub-capable hardware list lives
Discussion thread on r/Ubiquiti: https://www.reddit.com/r/Ubiquiti/s/jmYfUAqf04
