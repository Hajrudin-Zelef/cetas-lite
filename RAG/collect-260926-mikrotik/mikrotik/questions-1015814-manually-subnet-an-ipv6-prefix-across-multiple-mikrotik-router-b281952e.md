---
id: collect-260926-mikrotik/mikrotik/questions-1015814-manually-subnet-an-ipv6-prefix-across-multiple-mikrotik-router-b281952e
title: "questions-1015814-manually-subnet-an-ipv6-prefix-across-multiple-mikrotik-router-b281952e"
domain: mikrotik
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["distribution"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1015814-manually-subnet-an-ipv6-prefix-across-multiple-mikrotik-router-b281952e.md
source_anchor: ""
source_lines: [1, 97]
sha256: dafc0c69cfb95ecc26a36ce23a4dfde9133b248eaafae7fb8664201de2dec69a
---

# questions-1015814-manually-subnet-an-ipv6-prefix-across-multiple-mikrotik-router-b281952e

MikroTik's DHCPv6 Server- which also delivers PD functionality- has been a work-in-progress for years and so I wanted to avoid any dependencies on it. This HowTo documents how I achieved this to save others solving the same problem anew.
So if you're configuring your network to enable hosts to auto-configure IPv6 GUA addresses across multiple routers, but also want an alternative to MikroTik's DHCPv6 Server implementation, continue reading...
- HowTo SCOPE:
This HowTo will instruct on manually subnetting a ::/48 Prefix & configuring (2) routers with Neighbor Discovery, RIPng enabled and a few static routes. Hosts will configure a GUA address on the subnet of the interface that they connect to.
Prefix Distribution, DHCPv6 Server and DHCPv6 Client are NOT covered as they are not used in the configuration.
I also don't detail how the ::/48 routes out of Router #1 to the Internet as that can be network-specific. ie, the ::/48 Hurricane Electric assigned to me (FREE!) routes out of a 6to4 tunnel whereas you might have a ::/48 assigned directly by your ISP and have none of that.
IPv6 security is a subject better dealt with separately and outside the scope of this HowTo. If you don't have an IPv6 FW configured, after getting everything working you can just disable the IPv6 interfaces until you configure a set of sensible rules.
- COMPATIBILITY:
Since this configuration only uses Neighbor Discovery, manually addressed interfaces, RIPng & some Static Routing, it should be immune to breakage in future RouterOS upgrades.
Clients tested to successfully auto-addressed using this solution:
- Raspberry Pi4: Running Buster using DHCP
- MacBook: Running OSX Catalina 10.15.3, again, using default DHCP
- IOS & iPadOS 13.41
- Windows: A Virtualbox VM Windows 10 VM running on my MacBook using bridging
- IPv6 TESTING TOOLS:
A few IPv6-compatible tools to help you test & troubleshoot your config:
- Browser: https://test-ipv6.com.  Note: Firefox mobile incompatible w/ IPv6 addresses.
- IOS: Hurricane Electric IOS app ping &traceroute IPv6 addresses from iPhones & iPads
- OSX:
  - ifconfig -a
  - netstat -r -f inet6
  - ping6
  - traceroute6
- Linux:
  - ip -6 addr show
  - ip -6 route show
  - route -6 -n
  - ping6
  - traceroute6 -r
2001:4860:4860:8888 is a good address to use to test external connectivity; Google's DNS.
- GLOBAL UNICAST ADDRESSING ("GUA") MANUAL IPv6 ADDRESSING:
Using a ::/48 Prefix of:
2001:db8:1d4f::/48
I'll illustrate how I subnetted a ::/48 Hurricane Electric assigned me and manually addressed the interfaces on (2) routers. Process however would be the same for a non-Hurricane assigned ::/48
3.1: POINT-TO-POINT LINKS:
In this example, our (2) routers are connected between each other on ether2 using Local-Link fe80::/10 addresses which auto configure themselves. Router Announcements used for auto-address configuration are made using these Local-Link fe80::/10 addresses, NOT a Global Unicast Addresses ("GUA").
Note on P-2-P links in RoS v6: Although a ::/126 or a ::/127 would seem the obvious choice, support for ::/127's will only come in RoS v7.
3.2 Router #1: (RB4011)
Directly connected to Internet, exposing multiple SSIDs for wireless clients.  Although there are many interfaces configured for IPv6 on my RB4011, to simplify the example we'll only use:
- ether2: Uplink to Router #2. No GUA address required.
- wlan1: 2001:db8:1d4f:10::1/64
- wlan2: 2001:db8:1d4f:11::1/64
- wlan3: 2001:db8:1d4f:12::1/64
- etc.... /ipv6 address
add address=2001:db8:1d4f:10::1 interface=wlan1
add address=2001:db8:1d4f:11::1 interface=wlan2
add address=2001:db8:1d4f:12::1 interface=wlan3
3.3 Router #2: (RB951-2n)
Connected to Router #1, exposing single SSID for wireless clients.
- ether2: 2001:db8:1d4f:20::1/64 Uplink to Router #1. NOTE: GUA address added for WebFig access
- ether3: 2001:db8:1d4f:21::1/64
- ether4: 2001:db8:1d4f:22::1/64
- ether5: 2001:db8:1d4f:23::1/64
- wlan1:  2001:db8:1d4f:24::1/64 /ipv6 address
add address=2001:db8:1d4f:20::1 interface=ether2-master
add address=2001:db8:1d4f:21::1 interface=ether3
add address=2001:db8:1d4f:22::1 interface=ether4
add address=2001:db8:1d4f:23::1 interface=ether5
add address=2001:db8:1d4f:24::1 interface=wlan1
3.4 Additional Routers:
Were there a third router, we'd use 2001:db8:1d4f:30::X/64, incrementing subnet by 10 and use a host address of "1".  If exposing lots of SSIDs, it's suggested that you increment the subnets in even numbered multiples greater than 10 to keep things tidy.
3.5 Host Addressing:
Once routers GUA addresses are configured on the routers and other steps which follow are completed, network hosts will auto-configure a GUA address via SLAAC from same subnet as the router interface they're connecting to. ie, using the above addressing plan for Router #2, a host connecting to Router #2's AP wlan1 will auto-configure a GUA address of:
- 2001:db8:1d4f:24::PrivacyExtensionConfiguredByHost
- IPv6 ROUTING:
4.1 RIPng: ALL Routers
- Routing > RIPng: Enable RIPng on all router interfaces on all routers /routing ripng interface
add
- Routing > RIPng > RIPng Settings: Enable "Redistribute Static Routes" /routing ripng
set redistribute-static=yes
4.2 Static Routes: Router #1 ONLY
- IPv6 > Routes: On Router #1, Add static route(s) to uplink interfaces for Router #2 (and any additional routers) connected to Router #1.  Using the addressing plan, this would be to:
2001:db8:1d4f:**20**::1/64
    /ipv6 route
    add distance=1 dst-address=2001:db8:1d4f:20::1/128 gateway=\
        ether2-AP2-RB951-2n
- NEIGHBOR DISCOVERY ("ND")
Neighbor Discovery (ND) is used for link-layer address resolution (similar to ARP) & Address Auto-Config.  ND is an integral part of ALL IPv6 address auto-configuration- SLAAC, DHCPv6 Server & Prefix Delegation (PD).
IPv6 > ND > Add New: On each router, add only the interfaces clients will connect to and use for SLAAC auto-address configuration.  Please note the MTU's are set for 1280 as that is the size used by Hurricane Electric who routes the ::/48 they assigned to me.
For Router #1 this would be:
For Router #2 this would be:
SECURITY NOTE: There's a security issue using ND RA's (Router Announcements) for auto-address configuration.  An attacker with access to connected network could inject a RA into the network, triggering a device to add an IPv6 address or default route.
- Set IPv6 DNS & NTP Sources:
DNS:
IP > DNS: Add an IPv6 DNS source, such as 2001:4860:4860::8888 (Google) to the IPv4 one:
/ip dns
set allow-remote-requests=yes servers=8.8.8.8,2001:4860:4860::8888
NTP:
System > NTP Client: Add an IPv6 NTP source, such as 2610:20:6f15:15::27 to the IPv4 one:
/system ntp client
set enabled=yes primary-ntp=216.239.35.0 secondary-ntp=2610:20:6f15:15::27
- ALLOW ADMIN ACCESS:
If you want to use an IPv6 address for Webfig or SSH access, don't forget to update:
IP > Services and add a subnet(s) to allow for administrative access to the MikroTik.
Conclusion:
By this point you should have clients auto-address configuring IPv6 GUA addresses.  If you spot any errors/ommissions, please let me know so I can update the documentation.
Don't forget to spend some time configuring the IPv6 > Firewall and tightening-up security, or at least disable the IPv6 interfaces until you have the time to.  Hope you found this useful and it helped you get up to speed quickly.
