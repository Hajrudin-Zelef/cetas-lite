---
id: collect-261001-cisco/cisco/questions-2238-open-a-range-of-tcp-ports-in-cisco-ios-nat-9a5f176a
title: "questions-2238-open-a-range-of-tcp-ports-in-cisco-ios-nat-9a5f176a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "rotary"]
source: docs/RAG/collect-261001-cisco/questions-2238-open-a-range-of-tcp-ports-in-cisco-ios-nat-9a5f176a.md
source_anchor: ""
source_lines: [1, 46]
sha256: 172c5a0dca09560b2afb27bda5f0569dffb3590f744e28e514346eebdebbcf22
---

# questions-2238-open-a-range-of-tcp-ports-in-cisco-ios-nat-9a5f176a

(EDIT)
It seems that inside->outside works as expected, as seen in the answer below, but outside->inside actually does not, it allows everything, as OP suggested.
Adding 'reversible' in the NAT line it starts to honor the route-map for outside->inside, unfortunately it does not seem to work with ports:
- permit ip any host 194.100.7.226 works
- permit tcp any any works
- permit tcp any any eq 80 no match, does not work
- permit tcp any eq 80 any match, does not work
- permit tcp any eq 80 host 194.100.7.226 match, does not work
- permit tcp any eq 0 host 194.100.7.226 works
At '194.100.7.226' I'm doing 'telnet 91.198.120.222 80', that is my source is 194.100.7.226:ephemeral destination is 91.198.120.222:80. As the example #1 works, we can conclude that reversible actually 'reverses' the ACL, so that it works in same manner both directions, which makes sense.  
When the the connection match but does not work, in 'deny any any log-input line I get this:
  .Jul  7 07:58:59.118 UTC: %SEC-6-IPACCESSLOGP: list MOO denied tcp
  91.198.120.2(0) (Tunnel101 ) -> 194.100.7.226(0), 1 packet
So it really seems like L4 protocol type is carried, but ports are not carried during the NAT reversal. So outside->inside ranges do not work.
As suggested in question Cisco 867 forward UDP port range this works for outside->inside
ip nat pool MOO 91.198.120.2 91.198.120.2 prefix-length 30 type rotary
ip nat inside destination list MOO pool MOO
ip access-list extended MOO
 permit tcp any any range 22 100
 deny   ip any any log-input
It's bit ghetto I feel, as you don't have good control on the outside IP. Pool is the inside IP, outside IP is router outside IP.
Original answer of inside->outside working with ports:
ip nat inside source static 91.198.120.2 91.198.120.222 route-map MOO
!
ip access-list extended MOO
 permit icmp any any
 permit tcp any any range 22 telnet
!
route-map MOO permit 100
 match ip address MOO
!
route-map MOO deny 200
!
@91.198.120.2 I'm doing:
- telnet testhost 22
- telnet testhost 23
- telnet testhost 24
At testhost I can observe:
1   0.000000 91.198.120.222 -> 194.100.7.226 TCP 74 50925 > ssh [SYN] Seq=0 Win=14600 Len=0 MSS=1350 SACK_PERM=1 TSval=7995067 TSecr=0 WS=128
2   9.838471 91.198.120.222 -> 194.100.7.226 TCP 74 41586 > telnet [SYN] Seq=0 Win=14600 Len=0 MSS=1350 SACK_PERM=1 TSval=7997586 TSecr=0 WS=128
5  16.773181 91.198.120.2 -> 194.100.7.226 TCP 74 53307 > 24 [SYN] Seq=0 Win=14600 Len=0 MSS=1350 SACK_PERM=1 TSval=7999327 TSecr=0 WS=128
Tested on:
bu.ip.fi#sh ver | i ^Cisco
Cisco IOS Software, C880 Software (C880DATA-UNIVERSALK9-M), Version 15.1(2)T5, RELEASE SOFTWARE (fc1)
Cisco 881G (MPC8300) processor (revision 1.0) with 236544K/25600K bytes of memory.
bu.ip.fi#
