---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-16
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [831, 874]
sha256: 84edc8e105b680240c34d858fee183535fdefcb06905145a254551e9daa5e320
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

When a BGP speaker receives updates from multiple autonomous systems that describe different paths to the same destination, it must choose the single best path for reaching that destination. When chosen, the selected path is entered into the BGP routing table and propagated to its neighbors. The decision is based on the value of attributes that the update contains and other BGP-configurable factors.
When a BGP peer learns two EBGP paths for a prefix from a neighboring AS, it chooses the best path and inserts that path in the IP routing table. If BGP multipath support is enabled and the EBGP paths are learned from the same neighboring autonomous systems, instead of a single best path, multiple paths are installed in the IP routing table. Then, during packet switching, per-packet or per-destination load-balancing is performed among the multiple paths. The maximum-paths router configuration command controls the number of paths allowed.
These factors summarize the order in which BGP evaluates the attributes for choosing the best path:
- 
		  If the path specifies a next hop that is inaccessible, drop the update. The BGP next-hop attribute, automatically determined by the software, is the IP address of the next hop that is going to be used to reach a destination. For EBGP, this is usually the IP address of the neighbor specified by the neighbor remote-as router configuration command. You can disable next-hop processing by using route maps or the neighbor next-hop-self router configuration command.
- 
		  Prefer the path with the largest weight (a Cisco proprietary parameter). The weight attribute is local to the router and not propagated in routing updates. By default, the weight attribute is 32768 for paths that the router originates and zero for other paths. Routes with the largest weight are preferred. You can use access lists, route maps, or the neighbor weight router configuration command to set weights.
- 
		  Prefer the route with the highest local preference. Local preference is part of the routing update and exchanged among routers in the same AS. The default value of the local preference attribute is 100. You can set local preference by using the bgp default local-preference router configuration command or by using a route map.
- 
		  Prefer the route that was originated by BGP running on the local router.
- 
		  Prefer the route with the shortest AS path.
- 
		  Prefer the route with the lowest origin type. An interior route or IGP is lower than a route learned by EGP, and an EGP-learned route is lower than one of unknown origin or learned in another way.
- 
		  Prefer the route with the lowest multi -exit discriminator (MED) metric attribute if the neighboring AS is the same for all routes considered. You can configure the MED by using route maps or by using the default-metric router configuration command. When an update is sent to an IBGP peer, the MED is included.
- 
		  Prefer the external (EBGP) path over the internal (IBGP) path.
- 
		  Prefer the route that can be reached through the closest IGP neighbor (the lowest IGP metric). This means that the router will prefer the shortest internal path within the AS to reach the destination (the shortest path to the BGP next-hop).
- 
		  If the following conditions are all true, insert the route for this path into the IP routing table: Both the best route and this route are external. Both the best route and this route are from the same neighboring autonomous system. Maximum-paths is enabled.
- 
		  If multipath is not enabled, prefer the route with the lowest IP address value for the BGP router ID. The router ID is usually the highest IP address on the router or the loopback (virtual) address, but might be implementation-specific.
Route Maps
Within BGP, route maps can be used to control and to modify routing information and to define the conditions by which routes are redistributed between routing domains. See the “Using Route Maps to Redistribute Routing Information” section for more information about route maps. Each route map has a name that identifies the route map (map tag) and an optional sequence number.
BGP Filtering
You can filter BGP advertisements by using AS-path filters, such as the as-path access-list global configuration command and the neighbor filter-list router configuration command. You can also use access lists with the neighbor distribute-list router configuration command. Distribute-list filters are applied to network numbers. See the “Controlling Advertising and Processing in Routing Updates” section for information about the distribute-list command.
You can use route maps on a per-neighbor basis to filter updates and to modify various attributes. A route map can be applied to either inbound or outbound updates. Only the routes that pass the route map are sent or accepted in updates. On both inbound and outbound updates, matching is supported based on AS path, community, and network numbers. Autonomous system path matching requires the match as-path access-list route-map command, community based matching requires the match community-list route-map command, and network-based matching requires the ip access-list global configuration command.
Prefix List for BGP Filtering
You can use prefix lists as an alternative to access lists in many BGP route filtering commands, including the neighbor distribute-list router configuration command. The advantages of using prefix lists include performance improvements in loading and lookup of large lists, incremental update support, easier CLI configuration, and greater flexibility.
Filtering by a prefix list involves matching the prefixes of routes with those listed in the prefix list, as when matching access lists. When there is a match, the route is used. Whether a prefix is permitted or denied is based upon these rules:
- An empty prefix list permits all prefixes.
- An implicit deny is assumed if a given prefix does not match any entries in a prefix list.
- When multiple entries of a prefix list match a given prefix, the sequence number of a prefix list entry identifies the entry with the lowest sequence number.
By default, sequence numbers are generated automatically and incremented in units of five. If you disable the automatic generation of sequence numbers, you must specify the sequence number for each entry. You can specify sequence values in any increment. If you specify increments of one, you cannot insert additional entries into the list; if you choose very large increments, you might run out of values.
BGP Community Filtering
One way that BGP controls the distribution of routing information based on the value of the COMMUNITIES attribute. The attribute is a way to groups destinations into communities and to apply routing decisions based on the communities. This method simplifies configuration of a BGP speaker to control distribution of routing information.
A community is a group of destinations that share some common attribute. Each destination can belong to multiple communities. AS administrators can define to which communities a destination belongs. By default, all destinations belong to the general Internet community. The community is identified by the COMMUNITIES attribute, an optional, transitive, global attribute in the numerical range from 1 to 4294967200. These are some predefined, well-known communities:
- internet—Advertise this route to the Internet community. All routers belong to it.
- no-export—Do not advertise this route to EBGP peers.
- no-advertise—Do not advertise this route to any peer (internal or external).
- local-as—Do not advertise this route to peers outside the local autonomous system.
