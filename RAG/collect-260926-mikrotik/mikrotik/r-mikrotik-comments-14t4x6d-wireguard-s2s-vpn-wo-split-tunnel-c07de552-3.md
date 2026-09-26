---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-14t4x6d-wireguard-s2s-vpn-wo-split-tunnel-c07de552-3
title: "r-mikrotik-comments-14t4x6d-wireguard-s2s-vpn-wo-split-tunnel-c07de552"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/r-mikrotik-comments-14t4x6d-wireguard-s2s-vpn-wo-split-tunnel-c07de552.md
source_anchor: ""
source_lines: [428, 466]
sha256: 98b21e3afcf08806969984389d5b50944f3cf185dc1d4121ca77351ac681041e
---

# r-mikrotik-comments-14t4x6d-wireguard-s2s-vpn-wo-split-tunnel-c07de552

Section des commentaires
You want help with wireguard, but you did not share anything about the specific thing where you need help.
Would you mind sharing:
interface/wireguard/peers/print proplist=allowed-address,interface
Sure thing!
Attached 2 files with configs stripped from non VPN stuff.
Be aware that RB4011 has also:
WireGuard RA VPN interface - that works fine
Few NordVPN configs, that currently are disabled
If I understand correctly, you connect to R1 via the R1 public IP (70.18.138.1) which would try to reach this address due it has no more specific route than the default (0.0.0.0/0) on R2.
So if you want to maintain the wireguard tunnel running and reacht the R1 public interface it would require to add
ip route add dst-address=70.18.138.1 gateway=lte1
u/willyhun, u/adamxp12, u/Whitehawk29 - thanks for your help.
After adding specific route to R1, VPN started to work quite fine, yet not perfect.
Currently most websites work, but not all of them.
For instance speedtest.net takes forever to load.
I thought this might be MTU issue, so I started to sense on both sides what is the biggest I might use.
Pings to the other side of the tunnel allowed me to use regular MTU 1420, but some webstites still took forever to load.
I lowered MTU to 1300 temporarily - this didn't solve the problem.
I think it's not BW problem as I get satisfacory speedtest results
Attached GIFs compare behaviour with and without VPN.
When VPN is on it use MTU 1300 set on both sides and 8.8.8.8 as DNS
Here Speedtest with VPN
Here Speedtest without VPN
Few websites with VPN
You have the solution in your R1 firewall config already for something else (in the mangle hook), why don't you try to apply it on R2 (of course, modified for the actual device).
mangle you say... This is something I don't get
I created a rule like this on both ends:
chain=forward action=change-mss new-mss=1420 passthrough=yes tcp-flags=syn protocol=tcp out-interface=WG_S2S_VPN tcp-mss=1421-65535
No luck
Works! It works! :)
After editing mangle rule on both ends to:
chain=forward action=change-mss new-mss=1380 passthrough=yes tcp-flags=syn protocol=tcp out-interface=WG_S2S_VPN tcp-mss=1381-65535 log=no log-prefix=""
Everything started to work correctly
Thanks for your patience! :)
Do I understand it correctly, that for VPN vonnection that MTU is 1420 I lower payload packet size to 1380 and look between 1381-65535 size with this command?
Magle is something I don't feel too much :)
As someone else mentioned. Adding a default route that goes over the VPN will break the VPN connection as the external IP your connecting with will also be routed.
If your VPN IP is static your easy solution is a static route over the LTE for just that IP. otherwise you will have to do some more complicated route filtering
