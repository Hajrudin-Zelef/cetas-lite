---
id: collect-261001-cisco/cisco/t5-vpn-asa-8-3-ssl-vpn-nat-issue-td-p-1514316-c1154757-3
title: "sh ru object"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-vpn-asa-8-3-ssl-vpn-nat-issue-td-p-1514316-c1154757.md
source_anchor: ""
source_lines: [232, 371]
sha256: 7f69db0ccaa6e150c7addce4a467f42d9bbdab0a6b65d002a2566bef7d741ffb
---

# sh ru object

edit: I'm curious how that would work in case of pre 8.3
Anyhoo if I'll have the time I'll check it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2010 03:06 PM
Ok,
Thanks a lot.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-24-2010 02:41 AM
Stan,
First of all I found the solution dirty but it's working for me.
I would say, it's better to assign routble IP addresses to your RA clients and not bother with solution like this.
The NAT RPF checks are going to be addressed soon (redesign to an extent) in new code (I cannot give exact ETA now).
Here's what our local 8.3 NAT wizard come up with to cheat RPF check.
bsns-asa5520-10(config)# sh run nat
nat (outside,inside) source dynamic Anyconnect interface destination static LOCAL_NET_M LOCAL_NET_M
nat (outside,inside) source dynamic Anyconnect interface destination static LOCAL_NET_N LOCAL_NET_N
nat (inside,outside) source static LOCAL_NET_M LOCAL_NET_M destination static interface Anyconnect
nat (inside,outside) source dynamic any interface
bsns-asa5520-10(config)# nat (inside,outside) source static LOCAL_NET_N LOCAL_$
WARNING: All traffic destined to the IP address of the inside interface is being redirected.
WARNING: Users may not be able to access any service enabled on the inside interface.
bsns-asa5520-10(config)#
bsns-asa5520-10(config)#
bsns-asa5520-10(config)# sh run nat
nat (outside,inside) source dynamic Anyconnect interface destination static LOCAL_NET_M LOCAL_NET_M
nat (outside,inside) source dynamic Anyconnect interface destination static LOCAL_NET_N LOCAL_NET_N
nat (inside,outside) source static LOCAL_NET_M LOCAL_NET_M destination static interface Anyconnect
nat (inside,outside) source dynamic any interface
nat (inside,outside) source static LOCAL_NET_N LOCAL_NET_N destination static interface Anyconnect
bsns-asa5520-10(config)# sh run obj
bsns-asa5520-10(config)# sh run object
object network Anyconnect
 subnet 10.0.0.0 255.255.255.0
object network LOCAL_NET_N
 subnet 192.168.0.0 255.255.255.0
object network LOCAL_NET_M
 subnet 172.16.0.0 255.255.255.0
object network ALL
 subnet 0.0.0.0 0.0.0.0
object network INTERFACE
 host 172.16.0.1
bsns-asa5520-10(config)# sh ip
System IP Addresses:
Interface                Name                   IP address      Subnet mask     Method
GigabitEthernet0/0       outside                10.48.66.237    255.255.254.0   CONFIG
GigabitEthernet0/1       inside                 172.16.0.1      255.255.255.0   CONFIG
Current IP Addresses:
Interface                Name                   IP address      Subnet mask     Method
GigabitEthernet0/0       outside                10.48.66.237    255.255.254.0   CONFIG
GigabitEthernet0/1       inside                 172.16.0.1      255.255.255.0   CONFIG
bsns-asa5520-10(config)#                                                                                          
bsns-asa5520-10(config)# sh nat
Manual NAT Policies (Section 1)
1 (outside) to (inside) source dynamic Anyconnect interface destination static LOCAL_NET_M LOCAL_NET_M
    translate_hits = 28, untranslate_hits = 3
2 (outside) to (inside) source dynamic Anyconnect interface destination static LOCAL_NET_N LOCAL_NET_N
    translate_hits = 15, untranslate_hits = 0
3 (inside) to (outside) source static LOCAL_NET_M LOCAL_NET_M destination static interface Anyconnect
    translate_hits = 3, untranslate_hits = 3
4 (inside) to (outside) source dynamic any interface
    translate_hits = 2712, untranslate_hits = 1165
5 (inside) to (outside) source static LOCAL_NET_N LOCAL_NET_N destination static interface Anyconnect
    translate_hits = 1, untranslate_hits = 1
HTH,
Marcin
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-24-2010 06:17 AM
Well, your' right, working but showing "lack of elegance"
In other words, if I have 255 subnets behind vpn gateway I need to create 510 NAT rules.
Maybe it's better to try to terminate VPN on another subinterface, or this is also impossible?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-24-2010 06:44 AM
Stan,
First of all it's poor design choice to mask those users, since you can choose any pool you like
Coming back to "solution" ...
you don't need so many rules, you can pack things into object groups to some extent.
I wouldn't give up on 8.3 NAT just yet, it has it's shortcomings but they will be addressed as market depends it 
Regarding terminating SSL on different interface - usual rule applied. Routing to clients need to point out through the interface you're terminating the clients on.
Maybe you can ellaborate a bit more about why you're trying to PAT users instead of giving them a new routable subnet, or statically NATing client subnet to something else? I might come up with a more elegant (as you put it) solution?
Marcin
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-24-2010 07:18 AM
>> First of all it's poor design choice to mask those users, since you can choose any pool you like
Absolutely agree, but I don't want to amend routing design of existing network (and it's MPLS network and a big one). If I can for ex. choose vpn pool from inside network range (as I can do with others vpn solutions) that will be the answer.
>> Regarding terminating SSL on different interface...
The difference is that this interface won't be included in (inside, outside) dynamic NAT rule and therefore it will work.
Or even better, I think, if I can terminate vpn on inside interface and choose vpn pool range from inside LAN IPs.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-24-2010 08:05 AM
Stan,
Why not usre part of the pool assigned to "inside"? It should indeed work as any other vpn solution in that regard.
Regarding termination on different interface. There is no problem to do this if routing allows, ie. you cannot have users coming in through outside interface being terminated on inside interface, that will just not work to the best of my knowledge and has never worked.
Client "external" IP addresses need to be reachable via inetrface you're terminating them on. ie. for that default route pointing to the outside is good enough if you're terminating on the outside.
Marcin
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-24-2010 08:21 AM
>>> Why not usre part of the pool assigned to "inside"?
You mean just to use IPs from inside network range withing the same subnet? It's not working for me, I think because this address pool is linked to outside interface and thus no routing to this IP range from inside to outside, because they are supposed to be on the (inside). If you mean subnetting inside network then on one hand it's not possible, on the other hand I'll stuck with the same routing issue. The only answer could be choosing bridging to vpn pool which I don't know how to do with ASA and can be easily achieved with MS or old Baynetworks stuff.
