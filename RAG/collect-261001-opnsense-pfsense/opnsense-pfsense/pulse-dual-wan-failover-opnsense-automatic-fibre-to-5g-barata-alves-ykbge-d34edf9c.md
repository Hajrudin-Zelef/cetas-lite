---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/pulse-dual-wan-failover-opnsense-automatic-fibre-to-5g-barata-alves-ykbge-d34edf9c
title: "pulse-dual-wan-failover-opnsense-automatic-fibre-to-5g-barata-alves-ykbge-d34edf9c"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "latency", "throughput"]
source: docs/RAG/collect-261001-opnsense-pfsense/pulse-dual-wan-failover-opnsense-automatic-fibre-to-5g-barata-alves-ykbge-d34edf9c.md
source_anchor: ""
source_lines: [1, 29]
sha256: f4f51c5d4165a71c584916a1e8b0ae20d6099c1809772584e43e0f19af9f48e0
---

# pulse-dual-wan-failover-opnsense-automatic-fibre-to-5g-barata-alves-ykbge-d34edf9c

Dual-WAN Failover on OPNsense: Automatic Fibre-to-5G Backup (Tested)
A single internet line is a single point of failure. One outage in the middle of a video call, a large upload, or an online exam, and you are stuck reconnecting by hand at the worst possible moment. The fix is a second connection that takes over on its own — and with OPNsense, you can build it for the cost of a backup SIM and an evening of work.
This guide walks through a working dual-WAN failover setup on OPNsense: fibre as the primary link, 5G as the backup. When fibre drops, traffic moves to 5G in seconds, with no manual switching. When fibre comes back, it switches back. I built it, then validated it the only way that counts — by physically pulling the fibre cable and timing the cutover.
Why one connection is a risk you can remove
For a homelab, a small office, or anyone who works from home, "the internet is down" is rarely just an inconvenience. It interrupts meetings, breaks long uploads, and — if you study online like I do — can put an exam at risk. A second, physically independent path removes that single point of failure. Fibre and 5G fail for different reasons and travel different last miles, so it is unlikely that both go down at the same instant.
Two things to set expectations up front. This is failover, not load balancing — one line carries everything while the other waits. And it is failover, not SD-WAN — more on that distinction below, because it matters and most write-ups gloss over it.
The setup at a glance
Primary (Tier 1): fibre, on the WAN interface. Backup (Tier 2): 5G fixed wireless, on a second interface (OPT1). OPNsense monitors both links and routes LAN traffic through a gateway group that prefers fibre and falls back to 5G only when fibre is declared down. The steps below assume both WANs already exist as interfaces.
Step 1 — Give each gateway its own monitor IP
Under System ▸ Gateways ▸ Configuration, assign each gateway a distinct, external monitor IP. This is what lets OPNsense decide that a link is actually dead. I use 1.1.1.1 for fibre and 8.8.8.8 for 5G. Keep gateway monitoring enabled — without separate monitor IPs, failover never triggers, because the firewall has no independent way to tell the lines apart.
Step 2 — Create the gateway group
Under System ▸ Gateways ▸ Group ▸ Add, build the group. Name it clearly, e.g. GW_FAILOVER. Set fibre to Tier 1 and 5G to Tier 2. Set the trigger level to Member down. "Member down" means traffic only moves to 5G once the fibre gateway is declared dead — a clean primary/backup split, rather than flapping on minor latency.
Step 3 — The LAN rule most guides forget
This is the step people skip, and then wonder why nothing happens. Under Firewall ▸ Rules ▸ LAN, edit your "allow LAN to any" rule (or add one at the top), open Advanced, and set the Gateway to GW_FAILOVER. This is what activates policy-based routing. Without it, your LAN traffic ignores the gateway group entirely and keeps using the default route — so the failover you carefully configured does nothing at all. If you take one thing from this article, take this.
Step 4 — Outbound NAT and DNS
Outbound NAT in automatic mode already covers both interfaces. If you run Hybrid or Manual mode, make sure there is an outbound NAT rule for both fibre and 5G, or traffic over the backup will not be translated. For DNS, local Unbound resolves without issue. If you forward to your ISP's resolver instead, point it at 1.1.1.1 / 8.8.8.8 under System ▸ Settings ▸ General, so you are not tied to the fibre's DNS at the exact moment the fibre is the thing that is down.
Test it by breaking it
Do not trust failover until you have broken it on purpose. Physically unplug the fibre and time how long until 5G takes over. With "Member down", expect roughly 5–10 seconds. Plug it back in and confirm traffic returns to fibre. In my own test, the cutover landed inside that window — and the measured 5G throughput comfortably exceeded what the operator advertised.
What survives the switch
The cutover takes a few seconds, and active TCP sessions break — the connection drops and re-establishes on the new path. Anything where the session lives in a token rather than the IP address (most HTTPS web apps) survives a simple page reload without losing state. Long-lived streams and downloads take the hit. Knowing this in advance means no surprises.
Failover vs SD-WAN: be honest about scope
It is worth being precise, because the terms get mixed up. What this setup does is link-down failover with static policy-based routing. It keeps you online when the primary line dies. What it is not is SD-WAN: there is no real-time path selection by latency, jitter, or packet loss, and no central orchestration across multiple sites. True SD-WAN lives in the control plane; this lives in the data plane. For a homelab or a small office, link-down failover is almost always what you actually need — and it runs on hardware you already have, with no licence.
A note on hardware
I ran this on a second-hand, fanless appliance with no video output, which meant installing OPNsense entirely over a serial console. That is a useful skill in itself for repurposing end-of-life enterprise gear, but it is not a requirement: the failover configuration above is identical on any OPNsense-compatible box with two WAN interfaces.
Get the complete, step-by-step version
The above is the whole concept and the key steps. If you want the full walkthrough — including the serial-console install from scratch, every exact menu path, a troubleshooting table for when the switch does not trigger, and the same content in both English and Portuguese — I packaged it into a practical PDF guide:
It is the exact setup I built and tested, written so you can follow it on a headless appliance without guessing where a setting lives.
Have you set up a backup line for your network? If you run the cable-cut test, I would like to hear your cutover time.
https://njba.gumroad.com/l/opnsense-dual-wan-failover
