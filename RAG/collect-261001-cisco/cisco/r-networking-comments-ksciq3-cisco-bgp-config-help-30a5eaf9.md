---
id: collect-261001-cisco/cisco/r-networking-comments-ksciq3-cisco-bgp-config-help-30a5eaf9
title: "r-networking-comments-ksciq3-cisco-bgp-config-help-30a5eaf9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/r-networking-comments-ksciq3-cisco-bgp-config-help-30a5eaf9.md
source_anchor: ""
source_lines: [1, 45]
sha256: 5132b45e5cc642b0f9c33060f53a64abd2df351d57a2acdf252f6348fd09b293
---

# r-networking-comments-ksciq3-cisco-bgp-config-help-30a5eaf9

Cisco BGP Config help! 
        
    Hi all,
I'm trying to setup some BGP routing from my cumulus leafspine network to a Cisco Nexus.
I can see the vlans hosted on the cumulus ones, when I use 'show bgp all' but nothing on 'show routes' or 'show routes bgp. On the cumulus side, I can see my networks from the nexus but no traffic flows over.
This nexus hosts some vlans, it had no BGP setup before so we enabled the feature and setup a the config as follows:
router bgp 65001
address-family ipv4 unicast
network 172.16.0.0/16
network 172.24.96.0/24
neighbor fe80::9a03:9bff:fefb:4870 remote-as 65101
address-family ipv4 unicast
route-map RFC5549 out
soft-reconfiguration inbound
address-family ipv6 unicast
neighbor fe80::ba59:9fff:fe59:2b0 remote-as 65102
address-family ipv4 unicast
route-map RFC5549 out
soft-reconfiguration inbound
address-family ipv6 unicast
My understanding is this is route map related? Without routemaps is traffic blocked? Or should it all be allowed by default?
I've put the RFC5549 in due to sending ipv4 routes over ipv6 links. I'm running BGP unnumbered on cumulus. and followed this connection/configuration guide for the links: https://support.cumulusnetworks.com/hc/en-us/articles/212561648-Configuring-BGP-Unnumbered-with-Cisco-IOS
Cisco is completely out of my area of expertise so if anyone knows what I'm missing that would be awesome.
Thanks in advance
Section des commentaires
What is the content of your route-map called RFC5549?
I have discussed this with another and since removed this as its not required anymore. Thanks for the tips tho.
Soft configuration for bgp is not recommended and requires doing route refresh for routing changes. I would not recommend using it for your configuration
Too handy to not use though. show BGP neighbor x.x.x.x received routes is a godsend when troubleshooting towards third parties. Is there another way to display pre-policy routes?
No, but using soft inbound keeps two copies of the routing table in memory to see the output of the pre and post filtering. Working in BGP for an ISP, bgp soft config is considered by our team to be deprecated and i can think of the times where ive had to explain soft config and refresh to other bgp peers that my policy of advising against has only hardened with time. Advice is free to ignore though.
Note taken, will remove that too. Thanks.
Like someone mentioned, can you include the route-map config?
Two useful command to see what routes are actually being advertised to and from Nexus:
"Show bgp all neighbor x.x.x.x received-routes" "Show bgp all neighbor x.x.x.x advertised-routes"
Did you use activate command?
NX-OS doesn't have the activate command - trust me I looked :D
Activate isn't an nx-os command
OK, didn't know that.
Will tell us if the two are peered and how many prefixes are being received. If you can send the output we can go from there.
I'm only sharing IPV4 routes. but I'm using IPV6 local links for BGP
switch-r6a# show ip bgp summary BGP summary information for VRF default, address family IPv4 Unicast BGP router identifier 172.16.50.2, local AS number 65001 BGP table version is 45, IPv4 Unicast config peers 2, capable peers 2 10 network entries and 18 paths using 2080 bytes of memory BGP attribute entries [7/1008], BGP AS path entries [6/60] BGP community entries [0/0], BGP clusterlist entries [0/0] Neighbor V AS MsgRcvd MsgSent TblVer InQ OutQ Up/Down State/PfxRcd fe80::9a03:9bff:fefb:4870 4 65101 27630 27486 45 0 0 21:34:41 8 fe80::ba59:9fff:fe59:2b0 4 65102 27572 27416 45 0 0 21:34:41 8 switch-r6a# show ipv6 bgp summary BGP summary information for VRF default, address family IPv6 Unicast
So you're receiving 8 prefixes, are those not the ones you're looking for?
to see what those prefixes are (post filtering). If you keep soft-reconfig in your bgp config, you can see the routes before the filtering happens. That command is:
Without soft-reconfig you can only see them after filtering with the first command. The downside to soft-reconfig is that it uses more memory, but with only 8 routes the difference is negligible.
In your OP you asked if it could be route-map related, yes - but it's impossible to tell without seeing the route-map. If you remove it from the BGP config, it will allow the routes through unfiltered, which is fine for troubleshooting.
