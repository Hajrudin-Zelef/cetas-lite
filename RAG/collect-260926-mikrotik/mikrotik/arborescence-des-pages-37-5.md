---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-37-5
title: "arborescence-des-pages-37"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-37.md
source_anchor: ""
source_lines: [506, 562]
sha256: 1f6b99e79f30077220552284069828ff1976e592ff52d652c2c017c81f875c41
---

# arborescence-des-pages-37

Providing HTTP proxy service for authorized users. Authenticated user requests may need to be subject to transparent proxying (the "Universal Proxy" technique and advertisement feature). This HTTP mark is put automatically on the HTTP proxy requests to the servers detected by the HotSpot HTTP proxy (the one that is listening on the 64874 port) as HTTP proxy requests for unknown proxy servers. This is done so that users that have some proxy settings would use the HotSpot gateway instead of the [possibly unavailable outside their network of origin] proxy server users have configured in their computers. This mark is also applied when an advertisement is due to be shown to the user, as well as on any HTTP requests from the users whose profile is configured to transparently proxy their requests.

15 I chain=hs-auth action=jump jump-target=hs-smtp dst-port=25 protocol=tcp

Providing SMTP proxy for authorized users (the same as in rule #13).

### Packet Filtering

From **/ip firewall filter print dynamic** command, you can get something like this (comments follow after each of the rules):

 0 D chain=forward action=jump jump-target=hs-unauth hotspot=from-client,!auth

Any packet that traverse the router from an unauthorized client will be sent to the **hs-unauth** chain. The hs-unauth implements the IP-based Walled Garden filter.

 1 D chain=forward action=jump jump-target=hs-unauth-to hotspot=to-client,!auth

Everything that comes to clients through the router, gets redirected to another chain, called **hs-unauth-to**. This chain should reject unauthorized requests to the clients.

 2 D chain=input action=jump jump-target=hs-input hotspot=from-client

Everything that comes from clients to the router itself, gets to yet another chain, called **hs-input**.

 3 I chain=hs-input action=jump jump-target=pre-hs-input

Before proceeding with [predefined] dynamic rules, the packet gets to the administratively controlled **pre-hs-input** chain, which is empty by default, hence the invalid state of the jump rule.

 4 D chain=hs-input action=accept dst-port=64872 protocol=udp 
 5 D chain=hs-input action=accept dst-port=64872-64875 protocol=tcp 

Allow client access to the local authentication and proxy services (as described earlier).

 6 D chain=hs-input action=jump jump-target=hs-unauth hotspot=!auth

All other traffic from unauthorized clients to the router itself will be treated the same way as the traffic traversing the routers.

```
 7 D chain=hs-unauth action=return protocol=icmp
 8 D ;;; www.mikrotik.com
     chain=hs-unauth action=return dst-address=66.228.113.26 dst-port=80 protocol=tcp
```
Unlike the NAT table where only TCP-protocol related Walled Garden entries were added, in the packet filter **hs-unauth** chain is added to everything you have set in the **/ip hotspot walled-garden ip** menu. That is why although you have seen only one entry in the NAT table, there are two rules here.

 9 D chain=hs-unauth action=reject reject-with=tcp-reset protocol=tcp
10 D chain=hs-unauth action=reject reject-with=icmp-net-prohibited

Everything else that has not been white-listed by the Walled Garden will be rejected. Note the usage of TCP Reset for rejecting TCP connections.

```
11 D chain=hs-unauth-to action=return protocol=icmp
12 D ;;; www.mikrotik.com
     chain=hs-unauth-to action=return src-address=66.228.113.26 src-port=80 protocol=tcp
```
The same action as in rules #7 and #8 is performed for the packets destined to the clients (chain **hs-unauth-to**) as well.

13 D chain=hs-unauth-to action=reject reject-with=icmp-host-prohibited

Reject all packets to the clients with an ICMP reject message.
