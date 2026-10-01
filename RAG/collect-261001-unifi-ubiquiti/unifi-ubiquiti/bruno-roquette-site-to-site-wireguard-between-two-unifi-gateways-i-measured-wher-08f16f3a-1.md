---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/bruno-roquette-site-to-site-wireguard-between-two-unifi-gateways-i-measured-wher-08f16f3a-1
title: "bruno-roquette-site-to-site-wireguard-between-two-unifi-gateways-i-measured-wher-08f16f3a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["SpaceX"]
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-unifi-ubiquiti/bruno-roquette-site-to-site-wireguard-between-two-unifi-gateways-i-measured-wher-08f16f3a.md
source_anchor: ""
source_lines: [1, 89]
sha256: 34ffef5f68f1fa1b0a05d20df0bcb7bc7cf8def1f9f39e6598b6d0a60411c5cf
---

# bruno-roquette-site-to-site-wireguard-between-two-unifi-gateways-i-measured-wher-08f16f3a

Site-to-Site WireGuard Between Two UniFi Gateways: I Measured Where It Breaks
Two UCG Ultras on the same bench. Every claim below is a measurement, not a reading. Where I got something wrong, I say so and show what corrected it.
If you have two UniFi gateways in two locations and you want them to talk to each other, the internet gives you three answers:
- “Just use Site Magic.” It works. Nobody tells you the price.
- “You can do it with plain WireGuard.” You can get close. Where it stops is undocumented.
- “It’s impossible.” Nearly right, for the wrong reason.
I put two UCG Ultras on a bench, ran both paths end to end, and measured all three. Here is what actually happens.
The setup
- 2 × UniFi Cloud Gateway Ultra, UniFi Network 10.5.67
- Site A: 8 VLANs, zone-based firewall, dual WAN (PPPoE + DHCP-with-MAC-clone), both public, both on dynamic DNS
- Site B: same design, different second octet — deliberately non-overlapping subnets — single WAN, behind carrier-grade NAT
- A test host that exists only at Site B, so nothing could answer from the wrong side
That last point matters more than it sounds. Half the bad advice online comes from tests that accidentally succeeded locally.
INTERNET
                 |
   +-------------+--------------+
   |                            |
ISP 1 (PPPoE)            ISP 2 (DHCP + MAC clone)      ISP 3 (CGNAT)
public IP, dynamic       public IP, dynamic            NO public IP
ddns: vpn.<domain>       ddns: vpn2.<domain>           no inbound, ever
   |                            |                            |
   +------------+---------------+                            |
                |                                            |
        +-------+--------+                            +------+-------+
        |   SITE A       |                            |   SITE B     |
        |   UCG Ultra    | <-- the tunnel we want --> |   UCG Ultra  |
        |   dual WAN     |                            |  single WAN  |
        +-------+--------+                            +------+-------+
                |                                            |
     8 VLANs, 10.10.<vlan>.0/24                  8 VLANs, 10.20.<vlan>.0/24
     zone-based firewall                         same zones, same design
Why this exact shape is the whole problem
Read that diagram again with the WireGuard limitation in mind, because the two facts interlock and that’s the part nobody spells out:
- Site B is behind CGNAT, so it can never accept an inbound connection. It is forced to be the WireGuard client — the side that dials.
- UniFi NATs the WireGuard client interface. So the side that dials is the side that can’t be reached back.
The topology you’re forced into is exactly the topology that doesn’t work. There is no re-arrangement that helps, because the roles aren’t a choice: whichever end lacks a public IP must dial, and whichever end dials gets NATed. Site A having two public IPs doesn’t rescue it either — the constraint is at the far end, and the far end has none.
That’s why “just make the other one the server” — the most common suggestion in every thread — is not advice. It’s the configuration you already had.
Part 1 — Manual WireGuard, and exactly where it dies
The obvious build: WireGuard Server on one gateway, VPN Client on the other. The tunnel comes up. Handshake completes. The peer shows online. Then:
direction                     result
----------------------------  -----------------------------------
ICMP, client -> server        [ok] passes
ICMP, server -> client's LAN  [ok] passes (after the fixes below)
TCP, server -> client's LAN   [NO] connects, delivers nothing
The symptom is genuinely cruel. nc reports the port open. The TCP handshake completes. And curl times out with 0 bytes.
It looks exactly like an MTU problem. It is not an MTU problem.
The real cause, and other people who hit it
From the Ubiquiti feature request thread — 196 replies, three years, no official response — user rbrunka nails it:
UCG behind Starlink CGNAT as WireGuard client to a self-hosted VPS hub. The tunnel establishes fine and outbound policy-based routing works perfectly, but inbound traffic initiated by the hub to the UCG’s LAN is silently dropped — root cause is NAT applied by UniFi on the WG client interface.
And quasides on the Lawrence Systems forum, more bluntly: the WireGuard client is implemented behind NAT, so site-to-site is impossible; the server side can receive, but can’t initiate.
The client can reach the server. The server cannot reach back. That is the whole story, and it is not written anywhere in the product.
The workaround that gets further than any published attempt
UniFi does have a NAT rule type with an Exclude option, which turns off masquerading on a chosen interface. Finding it is its own small adventure:
Settings -> search NAT -> Policy Table -> NAT (/settings/policy-table?preset=nat-rules). /settings/nat does not exist. The Exclude checkbox sits below the fold of the form, so it’s easy to miss entirely.
For anyone automating, the API shape is:
{
  "type": "MASQUERADE",
  "exclude": true,
  "out_interface": "<network _id>",
  "source_filter":      {"filter_type": "NONE", "firewall_group_ids": [],
                         "invert_address": false, "invert_port": false},
  "destination_filter": {"filter_type": "NONE", "firewall_group_ids": [],
                         "invert_address": false, "invert_port": false},
  "protocol": "all", "ip_version": "IPV4", "rule_index": 0,
  "setting_preference": "manual", "pppoe_use_base_interface": false,
  "enabled": true, "logging": false
}
I found that shape by creating one rule in the UI and reading it back, because it isn’t documented either.
And it genuinely helps. Before the rule, the gateway answered !H (destination host unreachable) and nothing left the box. After it, ICMP crosses and TCP completes its handshake. That is further than any write-up I could find.
Get bruno roquette’s stories in your inbox
Join Medium for free to get updates from this writer.
It still doesn’t carry data. Source NAT was an obstacle, not the obstacle.
The test that settles it
Clean target — a machine that exists only at Site B, running python3 -m http.server:
from the local LAN (no tunnel)     200 · 1422 bytes ·  6 ms
across the tunnel, from Site A     000 ·    0 bytes ·  timeout, 3 of 3
⚠️ A methodology warning that cost me an hour. My first attempt aimed at macOS AirPlay Receiver on port 7000. It refuses connections whose source is outside the local subnet — so it produced a perfect false negative that looked like a network failure. Your test target must accept any source address.
Hypotheses I killed, and how
hypothesis                      how it died
------------------------------  ----------------------------------------------------------------------------
MTU                             ping succeeds up to 1408 B payload; the HTTP response was 1.4 kB
MSS clamping                    tried MTU 1380 and mss_clamp custom 1300 - no change
firewall at destination         the allow rule had 10,471 hits, and UniFi auto-created the (Return) policies
wrong key / AllowedIPs          compared byte for byte
service-side source filter      re-ran against http.server, which filters nothing
asymmetric routing on the host  test host had Wi-Fi off, wired only
Part 2 — Seven API traps, if you’re going to automate this anyway
- Static routes must be interface-route , pointing at the VPN network’s_id . As anexthop-route , traffic never leaves.static-route_interface will happily accept"wan" (wrong) and reject the interface name withIdInvalid .
- The VPN Client requires an Endpoint — without one you getapi.err.WireguardInvalidConfiguration . It dials; it never listens.
