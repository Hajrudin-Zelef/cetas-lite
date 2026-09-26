---
id: collect-260926-mikrotik/mikrotik/questions-713797-setting-up-mikrotik-router-with-native-ipv6-with-only-a-single-5caa9df6
title: "questions-713797-setting-up-mikrotik-router-with-native-ipv6-with-only-a-single--5caa9df6"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-713797-setting-up-mikrotik-router-with-native-ipv6-with-only-a-single--5caa9df6.md
source_anchor: ""
source_lines: [1, 10]
sha256: 8e2cc0b4ff18928bc85da084532ffca5520f722185d746bc5b02328965bb695e
---

# questions-713797-setting-up-mikrotik-router-with-native-ipv6-with-only-a-single--5caa9df6

I am at a colo provider that supplies a single IPv6 /64 block.
The goal was to route the provided /64 of IPv6 addresses to the hosts behind the Mikrotik running RouterOS 6.24.
Some Mikrotik examples and that I found always had the user getting a /48 or at least a /64 and another small block to connect with the gateway, or blogger major.io describing the possibility, however not recommended to use the link address to connect with the uplink router.
I didn't have access to this so I tried to do it another way.
What I had tried was a router IPv6 address on the gateway port as a /126 block aaaa.bbbb.cccc.dddd::2/126 to talk to the uplink router at aaaa.bbbb.cccc.dddd::1/126.
Then I created another router IPv6 address on the master port behind the firewall with the mask aaaa.bbbb.cccc.dddd:8000:1/65. I also configured neighbour discovery so that the clients could autoconfigure.
From the router terminal, I was able to ping the internet, and ping the hosts behind the firewall. From the hosts, I was able to ping the router addresses on both sides of the firewall but not when it needed to go to to the uplink.
From another network, I could ping the external router addresses in front of the firewall, but could not access the aaaa.bbbb.cccc.dddd:8000:1/65 that had a static route entry in place the master port behind the firewall. I had no rules in my firewall during testing.
Is my theory wrong, or is there problem with this model being used on the Microtik?
aaaa.bbbb.cccc.dddd::2/126by your provider?
