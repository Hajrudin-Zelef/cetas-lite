---
id: collect-261001-cisco/cisco/r-networking-comments-143qk55-help-with-cisco-bgp-configuration-afc32203
title: "r-networking-comments-143qk55-help-with-cisco-bgp-configuration-afc32203"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/r-networking-comments-143qk55-help-with-cisco-bgp-configuration-afc32203.md
source_anchor: ""
source_lines: [1, 75]
sha256: cccc1f5b011bb1294f47c79708651e4aaabe91f8fc51d9842a25d7187ad81b27
---

# r-networking-comments-143qk55-help-with-cisco-bgp-configuration-afc32203

Help with Cisco BGP configuration 
        
        
        
    
    
    On an existing Cisco BGP configuration I want to add a new BGP peer as backup and remove an old one:
router bgp 54321
 bgp router-id 9.9.9.9
 neighbor 50.50.50.50 remote-as 7777
 neighbor 50.50.50.50 description peer1
 neighbor 68.68.68.68 remote-as 7777
 neighbor 68.68.68.68 description peer2
 neighbor 68.68.68.68 ebgp-multihop 255
 !
 address-family ipv4
  network 9.9.9.0
  neighbor 50.50.50.50 activate
  neighbor 50.50.50.50 weight 300
  neighbor 50.50.50.50 prefix-list default-route in
  neighbor 50.50.50.50 prefix-list ARIN-BLOCK out
  neighbor 68.68.68.68 activate
  neighbor 68.68.68.68 weight 100
  neighbor 68.68.68.68 prefix-list default-route in
  neighbor 68.68.68.68 prefix-list deny-all out
 exit-address-family
!
ip prefix-list ARIN-BLOCK seq 5 permit 9.9.9.0/24
!
ip prefix-list default-route seq 5 permit 0.0.0.0/0
!
ip prefix-list deny-all seq 5 deny 0.0.0.0/0
Removing 68.68.68.68 and adding 12.12.12.12 it would look something like this:
router bgp 54321
 bgp router-id 9.9.9.9
 neighbor 50.50.50.50 remote-as 7777
 neighbor 50.50.50.50 description peer1
 neighbor 12.12.12.12 remote-as 7777
 neighbor 12.12.12.12 description peer2
 !
 address-family ipv4
  network 9.9.9.0
  neighbor 50.50.50.50 activate
  neighbor 50.50.50.50 weight 300
  neighbor 50.50.50.50 prefix-list default-route in
  neighbor 50.50.50.50 prefix-list ARIN-BLOCK out
  neighbor 12.12.12.12 activate
  neighbor 12.12.12.12 weight 200
  neighbor 12.12.12.12 prefix-list default-route in
  neighbor 12.12.12.12 prefix-list deny-all out
 exit-address-family
Questions:
Why does only the primary peer has "prefix-list ARIN-BLOCK out" is this a misconfiguration on the original setup that missed the advertised block on the secondary peer?
Same question for "prefix-list deny-all out" on the backup peer
Shouldn't both neighbors have the same config except for weight (300 for primary)?
EDIT Forgot to add, the new peer is directly connected and shouldn't need multihop so I removed that line
Section des commentaires
It can be difficult to influence inbound traffic. Techniques like AS path prepending and using provider's communities to directly set their local pref are usually effective but sometimes not 100% effective. A provider may have an incentive to send as much traffic as possible out a link that is billed based on bandwidth usage, so they can just set local pref as high as they want and you can't really do anything about it.
It appears the administrator decided that they want to be 100% sure inbound traffic is arriving via 50.50.50.50, so they are not even advertising their block out to the other peer. Presumably they would manually flip the outbound prefix filter if they needed to failover. A more graceful way to automatically accomplish this would be conditional route advertisement, in which the router would only advertise the block to 12.12.12.12 if a particular prefix (probably the 0/0 route via 50.50.50.50) is lost from the BGP table.
But with all that said, if you are dual-homing to a single provider (AS "7777") then you should be able to just use path prepending or a community to influence inbound traffic instead of relying on manual failover. Multi-homing to two different providers is where you usually have more difficulty.
The new 12.12.12.12 peer is actually a different AS, it's a different provider. I forgot to add that on the new config, so this is actually multihoming to two different providers.
Would I have asynchronous routing issues if I advertise ARIN-BLOCK to both peers (different AS)? how can I accomplish failover in this scenario?
Unless you have a specific reason not to want to receive inbound traffic on that circuit, I would just advertise the ARIN-BLOCK so that you have automatic failover. You can AS prepend the route to try to influence return traffic to come in the other circuit if that's desired. Routing on the internet is asynchronous anyways, so you shouldn't have any issues. If it's really important to only receive traffic on this backup circuit upon failover, then I would use conditional route advertisement.
You’d actually gain high availability by doing so. A lot of internet pathing is asynchronous fyi.
Do you require active-standby internet services? The best approach is to let BGP and upstreams do their thing. Treat your advertisements the same unless the business or technology requires it.
BGP communities are the most friendly way to do inbound traffic engineering. AS Path prepends are a way to do it. The worst case, in terms of stewardship, and most assured way is disaggregation but you need a shorter prefix.
If you want to apply the same prefix-list/route-map to multiple peers, I
recommend using peer-groups to ensure that the announcement logic is
applied to every peer within that group.
set default route to new preferred peer
advertise the whole block (/24) and also create a more specific advertisement (/25) for the link you want to be the main one,
wait 20 minutes for propagation
watch the consumption of the circuit migrate to the new circuit
turn off the circuit you no longer want to advertise
Unfortunately there will be an unavailability for a few minutes for some destinations, the only way to solve this is using BFD but even so there may be unavailability
