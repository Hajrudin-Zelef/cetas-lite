---
id: collect-260926-mikrotik/mikrotik/vakur87-dual-wan-mikrotik-two-internet-lines-part-1-c7a4f22764e6-3ce619e1-2
title: "vakur87-dual-wan-mikrotik-two-internet-lines-part-1-c7a4f22764e6-3ce619e1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/vakur87-dual-wan-mikrotik-two-internet-lines-part-1-c7a4f22764e6-3ce619e1.md
source_anchor: ""
source_lines: [104, 177]
sha256: 06c53bc23706b8cd2dec120184c6fcaba6e0c93510fd2a7d49fff9d967d81079
---

# vakur87-dual-wan-mikrotik-two-internet-lines-part-1-c7a4f22764e6-3ce619e1

/ip firewall nat print
You will usually find two lists called WAN and LAN, both marked defconf. You will also find bridge already in LAN, ether1 already in WAN, and a masquerade rule already in place.
If that is what you see, the only thing missing is your second line.
/interface list member
add interface=wan-static list=WAN
If you started from a blank configuration, create the whole set instead.
/interface list
add name=WAN
add name=LAN
/interface list member
add interface=wan-fast list=WAN
add interface=wan-static list=WAN
add interface=bridge list=LAN/ip firewall nat
add chain=srcnat action=masquerade out-interface-list=WAN
Using an interface list rather than naming each interface means one NAT rule covers both lines. When failover happens, NAT follows automatically.
A word of caution about that list. Putting an interface in the WAN list has effects beyond NAT. Firewall rules reference it too. Adding an interface there is not a small change, even though it looks like one.
Firewall
The default RouterOS firewall already does the right thing here. The rule that matters is this one, and it is present out of the box.
/ip firewall filter
add chain=input action=drop in-interface-list=!LAN comment="drop all not coming from LAN"
You do not need to add it. Check that it is there and that it sits below your accept rules.
Read it as: block anything arriving from somewhere that is not the local network.
Your static line has a real address, so it will be scanned constantly. Within minutes of going live you will see attempts on SSH, Telnet, and a spread of random ports. That rule is what stops them.
Check the counter after a day.
/ip firewall filter print stats
The number will be larger than you expect.
When you need policy routing
Everything so far uses one route at a time. That covers most needs.
Sometimes you want two paths active at once. Some traffic here, other traffic there.
That needs a second routing table.
/routing table add name=via-static fib
/ip route add dst-address=0.0.0.0/0 gateway=203.0.113.1 routing-table=via-static
A routing table is just a separate list of routes. The fib flag means the router will actually use it to forward packets, not merely store it.
By itself this table does nothing. Nothing is sent to it yet.
To send traffic there, you mark it.
/ip firewall mangle
add chain=prerouting action=mark-routing new-routing-mark=via-static \
    src-address=192.168.88.50 passthrough=no
Read that as: traffic from 192.168.88.50 should have its route looked up in via-static instead of the usual table.
The mangle section of the firewall does not block or allow anything. It attaches labels. Other parts of the router read those labels later.
The trap in policy routing
Your new table holds exactly one route: the default.
Now think about 192.168.88.50 trying to reach 192.168.88.60, a machine on the same network.
That packet gets marked. The router looks in via-static. The only route there is the default. So the router sends a packet meant for the machine next door out to the internet.
MikroTik’s documentation warns about this. A custom table that cannot find a destination will send traffic somewhere wrong.
The fix is to exclude local destinations from marking.
/ip firewall address-list
add list=no-policy-route address=192.168.88.0/24
add list=no-policy-route address=203.0.113.0/24/ip firewall mangle
set [find new-routing-mark=via-static] dst-address-list=!no-policy-route
Now the rule reads: traffic from this machine, going anywhere that is not local, uses the other table.
Add every local network you have to that list. Your LAN, your WAN subnets, any VPN subnets. Anything the router should deliver itself rather than send away.
The side effect nobody warns you about
RouterOS has a shortcut called fasttrack. Once a connection is established and known, fasttrack moves its packets through the router with far less work. It is the reason a small router can push several hundred megabits.
Fasttrack skips things. It skips mangle rules.
The moment you mark a connection for policy routing, that connection is no longer eligible for the shortcut. Every packet takes the long path.
That means more work per packet. On smaller hardware, the router’s CPU becomes the ceiling before the line does. Watch it during a speed test.
/system resource print
If CPU sits at 100%, you are measuring the router, not your connection.
One more thing to rule out. If your measured speed is lower than expected, check whether the router has traffic shaping configured.
/queue simple print stats where dropped!="0/0"
Queues are not part of the default configuration, so this is only relevant if you or someone before you added them. A queue sitting right at its limit with packets being dropped is the answer, and it is not the provider.
Confirming which line is carrying traffic
Routes tell you which line the router intends to use. To see which one it is actually using, watch the interface.
/interface monitor-traffic wan-fast
Start a large download on a computer on your LAN. If the numbers move here, that line is carrying it. If they do not, check wan-static instead.
This is the one check worth running after every change in this article.
Where this leaves you
You now have two lines. Traffic uses the fast one. If it fails, everything moves to the slow one and comes back afterwards. NAT and the firewall handle both.
There is one thing this setup cannot do yet.
Nothing from outside can reach you.
Your static line has a real address, and it is sitting there unused for inbound connections. The obvious next step is a VPN, so you can reach your own network from anywhere.
That is where the design above quietly breaks. Not the firewall. Not the VPN configuration. The routing.
Part 2 covers what goes wrong, why the standard fix does not apply, and how to rearrange this so both halves work.
