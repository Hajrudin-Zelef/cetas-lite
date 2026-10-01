---
id: collect-261001-cisco/cisco/t5-vpn-asa-8-3-ssl-vpn-nat-issue-td-p-1514316-c1154757-2
title: "sh ru object"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-vpn-asa-8-3-ssl-vpn-nat-issue-td-p-1514316-c1154757.md
source_anchor: ""
source_lines: [22, 231]
sha256: 010de01123e6eb96e94c881a610917aaf7e408e87e49a0e8861bbd712777b6d2
---

# sh ru object

			VPN
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-24-2010 09:04 AM
Stan,
Something like this works for me.
192.168.0.0/24 --- Router -- 172.16.0.0/24 ----ASA ==== cloud ==== Host. (inside the tunnel it get IP address from "over" pool, which is also Connected on inside)
bsns-asa5520-10(config)# clear xlate
INFO: 762 xlates deleted
bsns-asa5520-10(config)# sh run nat
nat (inside,outside) source static any any destination static SHARED SHARED
!
nat (inside,outside) after-auto source dynamic any interface
bsns-asa5520-10(config)# sh run object network
object network LOCAL_NETWORK
 subnet 192.168.0.0 255.255.255.0
object network SHARED
 subnet 172.16.0.0 255.255.255.0
bsns-asa5520-10(config)# sh run ip local pool
ip local pool ANY 10.0.0.100-10.0.0.200
ip local pool OVER 172.16.0.100-172.16.0.155
bsns-asa5520-10(config)# sh run tunne
bsns-asa5520-10(config)# sh run tunnel-group
tunnel-group DefaultWEBVPNGroup general-attributes
 address-pool OVER
If I catch your drift ... bridging inside and outside is not really needed on Cisco equipment as it should work via proxy arp straight out of the box, but I'm not faimilar with neither of vendors' solution for remote access.
Marcin
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-22-2010 03:28 PM
Stan,
Show us exactly what you've tried.
auto-nat from outside to inside with to do dyanmic PAT to interface should be OK.
You might need to indeed add twice NAT to traffic from inside subnets to said PATed (or NATed IPs).
As I said that should work in theory.
Let's see what you tried and on what version and we'll start from there.
Marcin
P.S. Regarding assimetric NAT I believe either Jay or Rama published an article recently.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-22-2010 07:21 PM
I attached config file (a little bit reduced, cut off some unnessessary details).
In short:
VPN SSL is enabled on 
"nat (Collab,Outside) after-auto source dynamic..." creates regular access from Collab to the Internet (outside)
"nat (Collab,Outside) after-auto source static..." creates NAT exemption for vpnpool, so that vpn client can access hosts on Collab
In this variant everything is Ok, as far as I'm not going further than directly copnnected subnets.
The point is that I have several subnets in a cloud behind Collab and those routers are not aware of vpnpool subnet, i.e. no routing back to vpn client.
To make a trick I need vpn client's packets to be PATed to Collab.
If I disable "nat (Collab,Outside) after-auto source static..." and add
nat (Outside,Collab) 1 source dynamic vpn-pools-group interface destination static Collab-netwok-group Collab-netwok-group
or this one
nat (Outside,Collab) 1 source static vpn-pools-group interface destination static Collab-netwok-group Collab-netwok-group
it immediately stops working (notorious "Asymetric NAT rule" error)
Another thing is that vpn traffic originates from 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2010 03:35 AM
Stan,
Here's the way I got it working.
object network Anyconnect
 nat (outside,inside) dynamic interface
object network LOCAL_NET_N
 nat (inside,outside) static LOCAL_NET_N
object network LOCAL_NET_M
 nat (inside,outside) static LOCAL_NET_M
nat (inside,outside) after-auto source dynamic any interface
sh run object
object network Anyconnect
 subnet 10.0.0.0 255.255.255.0
object network LOCAL
object network LOCAL_NET_N
 subnet 192.168.0.0 255.255.255.0
object network LOCAL_NET_M
 subnet 172.16.0.0 255.255.255.0
Based on outside NAT from here:
http://www.cisco.com/en/US/docs/security/asa/asa83/upgrading/migrating.html
If I'll find the time I'll try to optimize it, but cannot promise.
Marcin
edit: Configurtion done on 8.3.2
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2010 12:17 PM
Sorry, not working.
# sh ru object
object network vpn-pool-Outside
 subnet 192.168.110.0 255.255.255.0
object network inside-network
 subnet 10.137.0.0 255.255.0.0
# sh nat
Auto NAT Policies (Section 2)
1 (inside) to (outside) source static inside-network inside-network
    translate_hits = 19, untranslate_hits = 5
2 (outside) to (inside) source dynamic vpn-pool-Outside interface
    translate_hits = 9, untranslate_hits = 0
Manual NAT Policies (Section 3)
1 (inside) to (outside) source dynamic any interface
    translate_hits = 0, untranslate_hits = 0
Once I put auto NAT 1 vpn client can access inside-network but everyone from inside cannot access the Internet.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2010 01:11 PM
Stan,
Please note that this is not EXACTLY what you're looking for (note the static identity)
Possible differences:
- did you clear xlates ? Note that I get both translate and untranstale hits.
- ASA version?
Tomorrow I'll dig into this properly.
Marcin
From my setup:
bsns-asa5520-10# sh run nat
!
object network Anyconnect
 nat (outside,inside) dynamic interface
object network LOCAL_NET_N
 nat (inside,outside) static LOCAL_NET_N
object network LOCAL_NET_M
 nat (inside,outside) static LOCAL_NET_M
!
nat (inside,outside) after-auto source dynamic any interface
bsns-asa5520-10# sh run obj
bsns-asa5520-10# sh run object
object network Anyconnect
 subnet 10.0.0.0 255.255.255.0
object network LOCAL
object network LOCAL_NET_N
 subnet 192.168.0.0 255.255.255.0
object network LOCAL_NET_M
 subnet 172.16.0.0 255.255.255.0
bsns-asa5520-10# sh nat
Auto NAT Policies (Section 2)
1 (inside) to (outside) source static LOCAL_NET_M LOCAL_NET_M
    translate_hits = 3, untranslate_hits = 3
2 (inside) to (outside) source static LOCAL_NET_N LOCAL_NET_N
    translate_hits = 3, untranslate_hits = 3
3 (outside) to (inside) source dynamic Anyconnect interface
    translate_hits = 7, untranslate_hits = 7
Manual NAT Policies (Section 3)
1 (inside) to (outside) source dynamic any interface
    translate_hits = 156438, untranslate_hits = 110871
bsns-asa5520-10#
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2010 01:38 PM
Yes, I cleared xlate, though it can confuse only by keeping old nat translations, i.e. you'll have an access which shouldn't be there, and I'm speaking of loosing connection immediately once I put static NAT from inside to outside (which is practically NAT exemption).
Show version:
-----------------------
Cisco Adaptive Security Appliance Software Version 8.3(2)
Device Manager Version 6.3(3)
Compiled on Fri 30-Jul-10 17:49 by builders
System image file is "disk0:/asa832-k8.bin"
Config file at boot was "startup-config"
CS-FWA up 1 day 6 hours
failover cluster up 1 day 6 hours
Hardware:   ASA5510, 1024 MB RAM, CPU Pentium 4 Celeron 1599 MHz
Internal ATA Compact Flash, 256MB
BIOS Flash M50FW016 @ 0xfff00000, 2048KB
-----------------------
You can even simplify config using only one subnet inside.
Maybe I didn't put it clear, I need just to do this:
SSL VPN client terminated on outside interface should be able to access inside network and be dynamically NATed there through inside interface (all remote access vpn setup examples are speaking about EXEMPTING vpn address pool from NAT, I need the opposite).
At the same time all inside network should be able to access the Internet by dynamic NAT through outside interface.
I'm not sure this is possible though, haven't done that before with Cisco VPN.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2010 01:50 PM
Stan,
I know what you want to do, that's why I said the lines I gave above are NOT EXACTLY what you're looking for and will lab it out tomorrow.
Marcin
