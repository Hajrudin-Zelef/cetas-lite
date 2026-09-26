---
id: collect-260926-mikrotik/mikrotik/mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca-6
title: "mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2020-12-27", "2021-03-15", "2021-03-16", "2021-03-17", "2021-03-19"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca.md
source_anchor: ""
source_lines: [641, 693]
sha256: 0aca62d112ef8574b4e1cc83091ebc4137a355508575c36f68483b7c5a85c612
---

# mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca

Have been using your solution for quit some time running very stable. As of the newest stable version 6.48 I get an error “IPsec-SA expired before finishing rekey” every 30 minutes. This is the first version I encounter this error.

Do you have any clue why this is happening?

## Emiel · December 27, 2020 at 09:10

Found out it has to do with an update on IKEv2:

https://forum.mikrotik.com/viewtopic.php?f=2&t=147769#p740153

## gliappis · March 15, 2021 at 18:17

Hello, i want t set up the same site to site vpn.

At the isp routers, do you have make any port forwording? if i put the mikrotik as dmz host?

Thank you.

## Pessoft · March 15, 2021 at 23:27

Hello, yes, you need to forward the ports necessary for VPN to the Mikrotik. For IPsec IKEv2 with NAT traversal it will be port 4500/udp. Easy way is to do it using DMZ host as you mentioned, but that will forward all ports, so make sure you do not expose some service on Internet that you do not want to. Some routers also like to terminate IPsec connections if not specifically disabled in configuration, so check also your ISP router configuration or query your ISP for it, if your VPN packets do not come through.

## gliappis · March 16, 2021 at 13:49

Thanks! i put the mikrotiks at dmz! Now i have the ipsec established!

But i can’t ping mikrotik’s subnet hosts! i can ping local interface both of mikrotiks, but when i ping host i receive timeout! any ideas?

## gliappis · March 16, 2021 at 13:50

mikrotik 1 subnet 192.168.1.0/24, local ip 192.168.1.1

mikrotik 2 subnet 192.168.2.0/24, local ip 192.168.2.1

i can ping each other, but not the host from dhcp each side

## Pessoft · March 17, 2021 at 00:00

Most likely it is a firewall topic. Double check if there is a correct filter rule in forward chain which accepts forward between networks, as mentioned in guide, on both Mikrotiks. Also check if there is not some rule in the same chain which would affect the packets before accept rule. Next thing is to check nat rule in srcnat and dstnat chains: either you create accept rule in the same chain before any other nat rule is present ( i.e. before masquerade rule in srcnat chain ) or make sure that rules in nat table ignore IPsec. If all is ok, then check using Torch tool, where the packets get and if there are returning packets – this way you can identify if it is outgoing or incoming traffic that you need to fix. Also try ping from some DHCP host to IP of Mikrotik on remote site, which might help to identify where your packets get lost. Hope that helps.

## gliappis · March 17, 2021 at 12:55

i put at firewall raw, 2 routes for bypassing. Also i have to disable windows firewall, because win 10, does not allow ping from another subnet, only local.

And it works. The other think i have notice is that i have slow speed. The isp at both ends are 120 down/120 up, at i share a big file(3gb) and i saw speeds 3mb/sec. Any idea?

## Pessoft · March 19, 2021 at 01:43

It should be possible to configure Windows 10 firewall, so it accepts pings from more/all subnets and firewall could be enabled again then.

Regarding speed: Check real Internet connection speed at first, for example by speedtest.net to the server close to the location of remote Mikrotik. This should show if it is a general issue or issue with tunnel. Check also connection speed without using wireless network as that my impact the speed as well. Mikrotik also has a Bandwidth test, maybe you can give it a go as well to test connection between those with and without tunnel.

## Comments are closed.
