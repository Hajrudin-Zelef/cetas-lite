---
id: collect-260926-mikrotik/mikrotik/allclass-building-a-zero-downtime-dual-isp-failover-on-mikrotik-47408e09364e-b87e7810
title: "allclass-building-a-zero-downtime-dual-isp-failover-on-mikrotik-47408e09364e-b87e7810"
domain: mikrotik
role: reference
task: reference
actors: ["SpaceX"]
dates: []
keywords: ["cybersecurity"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/allclass-building-a-zero-downtime-dual-isp-failover-on-mikrotik-47408e09364e-b87e7810.md
source_anchor: ""
source_lines: [1, 49]
sha256: 23ae285191f2935f212865eded082bf5fc1c65d805ac3fcd42bc556361664bc9
---

# allclass-building-a-zero-downtime-dual-isp-failover-on-mikrotik-47408e09364e-b87e7810

Building a Zero-Downtime Dual-ISP Failover on MikroTik
How I engineered automatic failover between a fiber and satellite connection, and hardened the router while I was at it
Every office eventually learns the same lesson the hard way: a single internet connection is a single point of failure. Ours was no different. One fiber outage was all it took to make “we should really set up a backup connection” jump from the someday-list to this-week.
This is a walkthrough of how I built an automatic dual-ISP failover setup on a MikroTik router, fiber as primary, Starlink as secondary, along with the firewall hardening I did alongside it, since touching the edge router is always a good excuse to tighten everything else too.
The goal
Simple on paper: if the primary fiber connection drops, traffic should shift to the secondary link automatically, with no one needing to touch the router. When fiber comes back, traffic should shift back. No manual intervention, no support ticket, no “hey is the internet down for you too?” Slack message.
In practice, this touches routing, health-checking, and NAT all at once, and getting the failback behavior right (not just failover) is where most naive setups fall short.
Architecture
┌─────────────┐
   Fiber ─────┤             │
   (Primary)  │  MikroTik   ├──── LAN
              │   Router    │
   Starlink ──┤             │
   (Secondary)└─────────────┘
Two WAN interfaces, one router, one LAN. The router needed to:
- Continuously verify that the primary link was actually reachable (not just “up,” but genuinely passing traffic)
- Automatically re-route through the secondary link the moment the primary failed a health check
- Automatically revert once the primary was healthy again
- Keep NAT and DNS working correctly regardless of which link was active
Health-checking with Netwatch
MikroTik’s Netwatch tool is the backbone of this setup, it periodically pings a target and runs a script when the target goes up or down. The trick is picking the right target: pinging the ISP's own gateway isn't reliable, since a gateway can stay reachable even when upstream connectivity is dead. I pinged a small set of highly available public IPs beyond the ISP's edge instead, so a "down" result actually meant "the internet is down," not just "the first hop is fine."
Get Chukwuebuka Okeke’s stories in your inbox
Join Medium for free to get updates from this writer.
Roughly, the logic looked like this:
/tool netwatch
add host=<reliable-external-ip> interval=10s \
  up-script="/ip route enable [find comment=\"primary-route\"]; \
             /ip route disable [find comment=\"backup-route\"]" \
  down-script="/ip route disable [find comment=\"primary-route\"]; \
               /ip route enable [find comment=\"backup-route\"]"
Two static default routes exist at all times, one via fiber, one via Starlink, with different distances (priorities). Netwatch doesn’t create routes on the fly; it just enables and disables the ones already sitting in the routing table. That distinction mattered a lot during testing, a script that tries to add/remove routes dynamically is far more fragile than one that just flips existing routes on and off.
The failback problem nobody talks about
Failover is the easy half. Failback is where things get interesting.
A naive setup fails back to the primary link the instant it starts responding to a single ping, which is a problem if the fiber connection is flapping (up for ten seconds, down for five, repeat). Without care, you get a router that flip-flops between links every time the primary is unstable rather than actually recovered, which is worse for user experience than just staying on the backup link.
The fix was making the “up” condition stricter than the “down” condition, requiring a longer run of successful checks before switching back, rather than a single successful ping. This kind of asymmetry (fast to fail over, slow to fail back) is a pattern worth carrying into any health-check system, not just this one.
Firewall hardening, since we were already in there
Touching the edge router is always a good moment to review everything else facing the internet. A few things went in alongside the failover work:
- Bogon filtering — dropping traffic from IP ranges that should never legitimately appear on the public internet (private ranges, reserved space, etc.) arriving on the WAN side. This is free security with essentially zero downside.
- SSH brute-force escalation — a staged address-list system: an IP gets added to a “stage 1” list after a failed attempt, escalates to “stage 2” after a second, and lands on a week-long blacklist after a third. This turns an unlimited-attempt attack surface into a self-defeating one.
- Port scanner detection and blacklisting — using MikroTik’s PSD (port scan detection) matcher to catch and blacklist hosts probing multiple ports in a short window.
- DNS-based content filtering — routing DNS queries through a filtering resolver at the firewall level rather than relying on every individual device to be configured correctly.
None of these are exotic. What matters is that they’re layered, no single rule is doing all the work, and a gap in one doesn’t expose the whole router.
What I’d tell someone doing this for the first time
- Test failover by physically pulling the cable, not just disabling the interface in software. Software-level tests don’t always replicate what a genuine outage looks like to the router.
- Log every failover/failback event with a timestamp. When someone asks “was the internet actually down yesterday at 3pm,” you want an answer, not a guess.
- Treat failback as a separate design problem from failover, because it is one.
- Document the routing distances and why they’re set the way they are. Six months later, “why is the distance 2 in this network entry” is not obvious without a note next to it.
The result: an internet connection that survives a fiber outage without anyone in the building noticing, and a router that’s meaningfully harder to compromise than before I opened the terminal.
I’m a SOC analyst and network engineer working across cybersecurity, observability, and infrastructure. If you’re working through a similar dual-WAN setup or firewall hardening project, I’d love to hear how you approached it.
