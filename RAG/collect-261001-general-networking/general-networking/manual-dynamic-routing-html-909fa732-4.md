---
id: collect-261001-general-networking/general-networking/manual-dynamic-routing-html-909fa732-4
title: "manual-dynamic-routing-html-909fa732"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-dynamic-routing-html-909fa732.md
source_anchor: ""
source_lines: [218, 298]
sha256: f580d9af35aecf29484eff502c8bcaec7ed77b24fb7e51dbff121dbc853aa60f
---

# manual-dynamic-routing-html-909fa732

Neighbors in BGP are external or internal peers with whom the router establishes BGP sessions. Neighbors exchange route information, and each neighbor can be configured individually or as part of a Peer Group.
Note
A Route Reflector (RR) minimizes the need for a full mesh of BGP connections in the same AS by reflecting routes from its clients to other clients. In typical iBGP setups, all routers must directly peer with each other to avoid routing loops, but route reflectors allow clients to peer only with the RR. Clusters of RRs and clients prevent loops, and redundancy can be achieved with multiple RRs. This setup is crucial for scalability in large networks. For a small network, do not enable any of the RR specific options.
| Options | Description | 
|---|---|
| Enabled | Enable/Disable | 
| Description | Optional description for the AS-Path list. | 
| Number | ACL rule number (0-4294967294). No sequence numbers; removing the ACL is required to insert between entries. | 
| Action | Set permit to match or deny to negate the rule. | 
| AS | AS pattern to match, with regex allowed (e.g., “.$” or “_1$”). | 
Note
An AS-Path List is used to filter routes based on their AS-Path attributes. By matching specific AS paths, you can control the acceptance or rejection of routes from particular AS sequences. This is useful for policy enforcement in multi-AS environments but not needed for small networks.
| Options | Description | 
|---|---|
| Enabled | Enable/Disable | 
| Description | Optional description for the Prefix-List. | 
| Name | Name of the Prefix-List, descriptive of its purpose. If there should be multiple entries for the same prefix list, give them all the same name. | 
| IP Version | IP version to use. | 
| Number | ACL sequence number (1-4294967294). | 
| Action | Set permit to match or deny to negate the rule. | 
| Network | Specifies a network pattern to match, with optional ge (greater than or equal) and le (less than or equal) attributes to control the prefix length range. For example, a pattern like 192.168.0.0/16 ge 24 le 28 matches any route within the 192.168.0.0/16 block with prefix lengths from /24 to /28. | 
Note
Prefix Lists are used to filter prefixes in BGP. They match prefixes and control the import/export of specific IP ranges, allowing for fine-grained network control. It is very important to filter routes, especially when peering between eBGP and iBGP. Leaking routes of RFC1918 addresses to eBGP peers or announcing wrong prefixes is bad practice and could result in peering bans from external providers.
| Options | Description | 
|---|---|
| Enabled | Enable/Disable | 
| Description | Optional description for the Community-List. | 
| Number | Community-List number (1-99 for standard, 100-500 for expanded). | 
| Sequence Number | ACL sequence number (10-99). | 
| Action | Set permit to match or deny to negate the rule. | 
| Community | Community pattern to match, with optional regex. | 
Note
Community Lists allow tagging and filtering routes based on community attributes. By assigning community tags, you can apply policies that influence route preference. Useful for large networks; not so much in small networks.
| Options | Description | 
|---|---|
| Enabled | Enable/Disable | 
| Description | Optional description for the route-map. | 
| Name | Name of the route-map, used in neighbor configuration. | 
| Action | Set permit to match or deny to negate the rule. | 
| ID | Route-map ID (1-65535). Sorting is managed automatically. | 
| AS-Path List | Select the AS-Path List. If multiples with the same name exist, selecting one is enough. | 
| Prefix List | Select the Prefix List. If multiples with the same name exist, selecting one is enough. | 
| Community List | Select the Community List. If multiples with the same name exist, selecting one is enough. | 
| Set | Free text field for setting attributes, e.g., “local-preference 300” or “community 1:1”. | 
Note
Route Maps act like conditional filters, allowing you to set and modify BGP route attributes based on match criteria. They can combine prefix lists, community lists, and AS-paths for detailed route manipulation.
| Options | Description | 
|---|---|
| Enabled | Enable/Disable | 
| Name | Name of the peer group. | 
| Remote AS mode | “Use Remote AS Number” will use the number specified in the “Remote AS” field, while “External” or “Internal” will ignore it in favor of the alternative “remote-as internal” and “remote-as external” settings. | 
| Remote AS | Remote AS for the peer group. | 
| Listen Ranges | Enter one or multiple IP networks in CIDR notation. Accept connections from any peers in the specified prefix. | 
| Update-Source Interface | Physical IPv4 interface facing the peer. | 
| Next-Hop-Self | Sets the local router as the next hop for routes advertised to the peer group, commonly used in Route Reflector setups. | 
| Send Defaultroute | Enable sending of default routes to the peer group. | 
| Prefix-List In | Prefix list to filter inbound prefixes from this peer group. | 
| Prefix-List Out | Prefix list to filter outbound prefixes sent to this peer group. | 
| Route-Map In | Route-map to apply to routes received from this peer group. | 
| Route-Map Out | Route-map to apply to routes advertised to this peer group. | 
Note
A Peer Group in BGP simplifies configurations by grouping neighbors with similar settings. Another possibility is defining Listen Ranges to accept connections from multiple peers without configuring each of them as neighbor individually. This approach reduces management complexity and ensures uniform settings across peers. Peer Groups are especially useful in larger networks where multiple BGP peers require identical policy; not so much in small networks.
Border Gateway Protocol (BGP) is an exterior gateway protocol used to exchange routing information between autonomous systems (AS) on the Internet. As a path-vector protocol, BGP makes routing decisions based on defined paths, network policies, or administrator-configured rules. BGP has two main types: iBGP, used for routing within a single AS (using private AS numbers from 64512 to 65534), and eBGP, which operates between different AS across the Internet (using public AS numbers 1 to 64511). BGP’s flexibility and scalability make it essential for global Internet routing and large network infrastructures.
Supplemental Protocols
BFD (Bidirectional Forward Detection)
| Options | Description | 
|---|---|
| Enable | This will activate the BFD service. | 
| Options | Description | 
|---|---|
| Enabled | Enable/Disable | 
| Description | Set an optional description for this neighbor. | 
| Peer-IP | Specify the IP of your neighbor. | 
| Multihop | Enables multi-hop mode, allowing BFD to expect packets with TTL less than 254 and listen on the multihop port (4784). Note: Echo mode is not supported in multi-hop (see RFC 5883, section 3). | 
| Local Address | Specifies the local IP address to bind the peer listener to and to use for sending packets. This option is mandatory for IPv6. | 
| Interface | Selects which interface to use for this BFD peer. | 
| Detect multiplier | Configures the detection multiplier to determine packet loss. The remote transmission interval will be multiplied by this value to determine the connection loss detection timer. The default value is 3. | 
| Receive interval | Configures the minimum interval that this system is capable of receiving control packets. Defaults to 300 ms. | 
| Transmit interval | The minimum transmission interval (less jitter) that this system wants to use to send BFD control packets. Defaults to 300ms. | 
Note
A BFD Neighbor refers to a neighboring device configured to use BFD with BGP, OSPF or other protocols. BFD neighbors exchange small, frequent packets to rapidly detect link failures, reducing convergence time in routing; usually a whole lot faster than typical keepalive or hello mechanisms.
