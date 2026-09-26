---
id: collect-260926-mikrotik/mikrotik/vakur87-dual-wan-mikrotik-two-internet-lines-part-1-c7a4f22764e6-3ce619e1-1
title: "vakur87-dual-wan-mikrotik-two-internet-lines-part-1-c7a4f22764e6-3ce619e1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "ethernet"]
source: docs/RAG/lot-mikrotik/forum/misc/vakur87-dual-wan-mikrotik-two-internet-lines-part-1-c7a4f22764e6-3ce619e1.md
source_anchor: ""
source_lines: [1, 103]
sha256: c444adf1981fbd34615d868f4a682fc251e65a4c9e8681e8003b542481a25319
---

# vakur87-dual-wan-mikrotik-two-internet-lines-part-1-c7a4f22764e6-3ce619e1

Dual-WAN MikroTik, Two Internet Lines [Part 1]
One line has a real address. The other has speed. Here is how to run both.
Plenty of homes and small offices end up with two internet connections.
Often they are very different. One is a business line. It is not fast, but it comes with a real, fixed IP address. People can reach you on it from outside.
The other is a consumer line. It is fast and cheap. But it sits behind the provider’s own NAT, so you get an address that only exists inside their network. Nobody outside can reach you through it.
Neither line is good enough alone. Together they are exactly what you want.
This article covers the routing side. Part 2 adds WireGuard on top, which is where it gets interesting.
where it gets interesting.
Who wrote this, and what that means for you
I am not a network engineer. I am a technical person who wanted a working home network and had to figure out a MikroTik router to get it.
Everything here comes from my own setup. It works, and I checked what I could against MikroTik’s documentation. But there may be cleaner ways to do this, and there are certainly parts of RouterOS I do not fully understand.
Read it as one person’s notes, not as a reference. Test on your own equipment before you rely on any of it.
A note on the addresses in this article
Every IP address, interface name, and port number below is an example. They are placeholders chosen to make the explanation readable.
Do not paste these commands as they are. Your provider gives you a different address on the static line. Your local network is probably a different subnet. Your interfaces may be numbered differently.
Before you run anything, work out your own values:
- the address and gateway your provider assigned to the static line
- the subnet your local network uses
- which physical ports each line is plugged into
Then substitute them throughout. The logic of each rule stays the same. Only the numbers change.
The addresses in play
Four different kinds of address appear in this article. Mixing them up is the most common source of confusion, so here they are up front.
- What: Static WAN address
 Example: 203.0.113.42/24, gateway 203.0.113.1
 Where it comes from: Assigned by the provider, written into the router by hand. Never changes.
Reachable from outside?: Yes
- What: Dynamic WAN address
 Example: 100.64.10.5, gateway 100.64.10.1
 Where it comes from: Handed out by the provider over DHCP. Can change at any time.
Reachable from outside?: No
- What: LAN addresses
 Example: 192.168.88.0/24, router at .1, devices from .100 upward
 Where it comes from: Your own choice. The router hands them out over DHCP.
Reachable from outside?: No
- What: Gateway addresses
 Example: 203.0.113.1 and 100.64.10.1
 Where it comes from: The provider’s router at the far end of each line.
Reachable from outside?: Not yours to use
Two of these deserve a closer look.
The dynamic address is not a real internet address. 100.64.10.5 belongs to a range reserved for provider-side NAT. Hundreds of other customers are sharing one real address somewhere upstream. You can reach out. Nothing can reach in. There is no port forwarding, no VPN endpoint, no incoming anything. Some providers call this CGNAT.
The static address is a real internet address. 203.0.113.42 is yours alone. Anything on the internet can send packets to it. That is useful and it is also a responsibility, which is why the firewall section matters.
Your LAN addresses never leave your building. The router swaps them for a WAN address on the way out and swaps them back on the way in. That swap is the NAT section below.
One more distinction, because it causes trouble later. The address the router uses when it sends a packet out is called the source address. Which line a packet leaves by is a separate question. These two are decided at different moments, and Part 2 is entirely about what happens when they disagree.
What you are building
ether1  →  wan-fast    DHCP, 100.64.10.5    provider NAT, 400 Mbps, no way in
ether2  →  wan-static  Static 203.0.113.42/24, gateway 203.0.113.1
bridge  →  LAN         192.168.88.0/24
The goal is simple to state.
Everyday traffic leaves through the fast line. If the fast line dies, everything moves to the slow line automatically. When the fast line comes back, everything moves back.
Naming your interfaces
Do this first. It costs nothing and it saves you every time you read a config later.
/interface ethernet
set [find default-name=ether1] name=wan-fast
set [find default-name=ether2] name=wan-static
Every rule from now on refers to wan-fast and wan-static. Six months later you will still know what they mean.
Addresses
The fast line uses DHCP. The provider hands you an address.
/ip dhcp-client
add interface=wan-fast add-default-route=yes default-route-distance=1 use-peer-dns=no
Two things there are worth explaining.
default-route-distance=1 sets the priority of the route this line creates. Lower numbers win. We will come back to this.
use-peer-dns=no stops the provider's DNS servers from being pushed into your router. Set your own instead. Provider DNS servers change without warning and are often slow.
The static line is configured by hand.
/ip address
add address=203.0.113.42/24 interface=wan-static
Note that no route appears yet. Adding an address only tells the router about the local network. It does not tell it how to reach the rest of the internet.
The default route, and what distance means
Your router needs to know where to send traffic that is not for the local network. That is the default route.
You have two possible paths, so you add two routes.
/ip route
add dst-address=0.0.0.0/0 gateway=203.0.113.1 distance=2 check-gateway=ping
The fast line already added its own route through DHCP, at distance 1.
Now look at both.
/ip route print detail where dst-address=0.0.0.0/0DAd dst-address=0.0.0.0/0 gateway=100.64.10.1  distance=1
 s  dst-address=0.0.0.0/0 gateway=203.0.113.1  distance=2 check-gateway=ping
The letter A means active. Only the first route has it.
This is how failover works in RouterOS. Both routes exist. Only the one with the lowest distance is used. The other sits there waiting.
A route without A is not broken. It is a backup doing its job.
That last sentence is worth repeating, because it trips people up. A backup route showing no A flag while the primary is healthy is the correct and expected state.
Making failover actually work
Adding a second route is not enough on its own.
Imagine the fast line’s cable is still plugged in, the interface is still up, but the provider has a fault upstream. The route stays active because nothing looks broken locally. Traffic keeps going into a black hole.
The check-gateway=ping setting fixes this. The router pings the gateway regularly. If the pings stop coming back, the route is pulled and the next one takes over.
Add it to the DHCP line too.
/ip dhcp-client
set [find interface=wan-fast] add-default-route=yes default-route-distance=1
/ip route
set [find gateway=100.64.10.1] check-gateway=ping
Note that a DHCP route is created fresh each time the lease renews. If you want check-gateway to survive that, a small script on lease renewal is the usual approach.
Testing it
Do not assume failover works. Test it.
Get A Guy on the Internet’s stories in your inbox
Join Medium for free to get updates from this writer.
Unplug the fast line. Then check the routes again.
/ip route print detail where dst-address=0.0.0.0/0
The A flag should have moved to the static route within a few seconds. Browse something to confirm.
Plug it back in. The flag should move back.
If it does not, you have a failover setup that exists only on paper. Better to find that out now.
NAT
Your LAN uses private addresses. They cannot travel on the internet. The router rewrites them on the way out.
On a router with MikroTik’s default configuration, most of this already exists. Check before you add anything.
/interface list print
/interface list member print
