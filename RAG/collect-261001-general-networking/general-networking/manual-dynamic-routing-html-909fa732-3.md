---
id: collect-261001-general-networking/general-networking/manual-dynamic-routing-html-909fa732-3
title: "manual-dynamic-routing-html-909fa732"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution"]
source: docs/RAG/collect-261001-general-networking/manual-dynamic-routing-html-909fa732.md
source_anchor: ""
source_lines: [144, 217]
sha256: 4003c8a4cb6a0654ddd4c49dca9f933b0abd0f7496af01e407d554090d7bb175
---

# manual-dynamic-routing-html-909fa732

Interfaces in OSPF are for defining how each interface participates in OSPF routing. Key settings include Area, Hello Interval and Dead Interval for neighbor relationships and Cost for path preference. Networks and Interfaces cannot have the same Area, only one of them can be defined in the Backbone Area. For simpler networks, using Interfaces is recommended.
| Options | Description | 
|---|---|
| Enabled | Enable / Disable | 
| Name | The name of your Prefix-List. If there should be multiple entries for the same prefix list, give them all the same name. | 
| Number | The ACL sequence number (10-99). | 
| Action | Set permit for match or deny to negate the rule. | 
| Network | The network pattern you want to match. | 
Note
Prefix Lists in OSPF serve as a filter to control the distribution of specific IP ranges within the network. Though not as common as in BGP, prefix filtering can prevent or allow propagation of defined network routes.
| Options | Description | 
|---|---|
| Enabled | Enable / Disable | 
| Name | Route-map name for matching and setting patterns, enabled via redistribution. | 
| Action | Set permit for match or deny to negate the rule. | 
| ID | Route-map ID between 10 and 99. Entries added in order of insertion. | 
| Prefix List | Allows for matching based on prefix lists, multiple selections enabled. | 
| Set | Free text for setting options, e.g., “local-preference 300” or “community 1:1”. | 
Note
Route Maps act like conditional filters, allowing you to set and modify OSPF route attributes based on match criteria. They can combine prefix lists for detailed route manipulation.
Open Shortest Path First (OSPF) is a widely used link-state routing protocol designed for IP networks within a single autonomous system (AS). Operating as an interior gateway protocol (IGP), OSPF builds a network topology map by gathering link-state information from routers, allowing it to create an optimal routing table for IP packet delivery. OSPFv2 (RFC 2328) supports IPv4, while v3 (RFC 5340) extends support to IPv6.
BGP (Border Gateway Protocol)
| Options | Description | 
|---|---|
| Enable | This will activate the BGP service. | 
| BGP AS Number | Your AS Number here. | 
| BGP AD Distance | Adjust BGP administrative distance, typically set to 20. Useful if you want to prefer OSPF-learned routes. | 
| Router ID | Optional fixed router ID for BGP. | 
| Graceful Restart | Enable BGP graceful restart as per RFC 4724, allowing packet forwarding during protocol restoration. | 
| Network | Defines connected networks to be advertised over BGP. Disable Network Import-Check to announce all networks. | 
| Network Import-Check | By default, only networks present in the routing table are advertised. Disable to announce all configured networks. | 
| Bestpath | Route selection modifiers that influence the best path. | 
| Enforce First AS | Deny an update received from an external BGP (eBGP) peer that does not list its autonomous system number at the beginning of the AS_PATH in the incoming update. | 
| Log Neighbor Changes | Enable extended logging of BGP neighbor changes. | 
| Maximum Paths | Maximum number of equal-cost paths for EBGP multipath (ECMP). Leave empty to use FRR default (1). | 
| Maximum Paths (IBGP) | Maximum number of equal-cost paths for IBGP multipath (ECMP). Leave empty to use FRR default (1). | 
| Route Redistribution | Select other routing sources to redistribute to other nodes. Can be combined with a Route Map per redistribution. | 
| Options | Description | 
|---|---|
| Enabled | Enable/Disable | 
| Description | Optional description for the neighbor. | 
| Peer-IP | Specify the IP address of the BGP neighbor. | 
| Remote AS mode | “Use Remote AS Number” will use the number specified in the “Remote AS” field, while “External” or “Internal” will ignore it in favor of the alternative “remote-as internal” and “remote-as external” settings. | 
| Remote AS | AS (Autonomous System) number of the neighbor, required for establishing a BGP session. | 
| Local AS | Optional local AS number (1-4294967295) presented only to this neighbor (per-neighbor local-as). Useful for migration/interoperability scenarios requiring a different local AS than the global BGP AS. | 
| BGP MD5 Password | Password used for MD5 authentication of BGP connections to enhance security. | 
| Weight | Default weight for routes from this neighbor; higher weight increases path preference within the same AS. | 
| Local Initiater IP | Specify the local IP address used to establish connections with the neighbor. Only relevant for MD5 authentication. | 
| Update-Source Interface | Interface where BGP sessions are sourced from, typically required when using loopback addresses. | 
| IPv6 link-local interface | Interface for IPv6 link-local neighbors, used primarily for link-local IPv6 addressing. | 
| Next-Hop-Self | Sets the local router as the next hop for routes advertised to the neighbor, commonly used in Route Reflector setups. | 
| Next-Hop-Self All | Extends Next-Hop-Self by applying the setting to all addresses, including multiple address families. | 
| Multi-Hop | Enables connections to eBGP neighbors across multiple hops; often required for peering over loopback addresses. | 
| Multi-Protocol | Enables multiprotocol BGP for support of additional address families like IPv6. | 
| Route Reflector Client | Marks the neighbor as a client for a route reflector; used to reduce the number of full BGP mesh connections. | 
| Soft reconfiguration inbound | Allows policy changes without resetting the session by storing inbound updates. | 
| BFD | Enable Bidirectional Forwarding Detection (BFD) for rapid link failure detection with the neighbor. | 
| Keepalive | Sets the interval (default 60 seconds) between keepalive messages to check the neighbor’s availability. | 
| BFD Strict Mode | Strict mode requires the BFD session to be established before allowing BGP to connect. If BFD goes down, BGP immediately closes the session. This prevents BGP sessions from remaining up when the underlying link is down. Requires BFD to be enabled. | 
| Hold Down Time | Time (default 180 seconds) before declaring a neighbor down if no keepalive messages are received. | 
| Connect Timer | Interval to attempt reconnecting with a neighbor after a disconnect. | 
| Send Defaultroute | Sends a default route to the neighbor, useful in small AS environments where a full routing table is not necessary. | 
| Enable AS-Override | Allows replacement of the neighbor’s AS with the local AS, common in BGP confederations. | 
| Remove Private AS | Configure how ASNs are propagated in outbound updates. Depending on selection, they can be removed or replaced. Remove Private AS: Remove Private ASNs in outbound updates. Remove All AS: Remove all ASNs in outbound updates. Replace Private AS: Replace private ASNs with our ASN in outbound updates. Replace All AS: Replace all ASNs with our ASN in outbound updates. | 
| Allow AS In | Accept incoming routes with AS path containing AS number with the same value as the current system AS. | 
| Disable Connected Check | Allows eBGP connections over loopback addresses by bypassing checks for direct connections. | 
| Attribute Unchanged | Keeps specified attributes (like MED, AS-Path, etc.) unchanged in updates to the neighbor. | 
| Enable BGP-Capabilities | Allows this neighbor to negotiate BGP Capabilities including link-local, extended nexthop, dynamic, and FQDN. | 
| Prefix-List In | Prefix list to filter inbound prefixes from this neighbor. | 
| Prefix-List Out | Prefix list to filter outbound prefixes sent to this neighbor. | 
| Route-Map In | Route-map to apply to routes received from this neighbor. | 
| Route-Map Out | Route-map to apply to routes advertised to this neighbor. | 
| Peer Group | Groups neighbors with similar configurations to simplify management. | 
Note
