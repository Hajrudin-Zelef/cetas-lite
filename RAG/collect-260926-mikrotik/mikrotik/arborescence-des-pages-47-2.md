---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-47-2
title: "Common Actions and Associated properties"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-47.md
source_anchor: ""
source_lines: [129, 178]
sha256: 70f929b2d4203dc7a38daa66f7d34d495236aae5e725ab86ebf912a20da88e79
---

# Common Actions and Associated properties

| Property | Description | 
|---|---|
| **chain** (*name* ; Default: ) | Specifies to which chain rule will be added. If the input does not match the name of an already defined chain, a new chain will be created | 
| **comment** (*string* ; Default: ) | Descriptive comment for the rule | 
| **content** (*string* ; Default: ) | Match packets that contain specified text | 
| **dscp** (*integer: 0..63* ; Default: ) | Matches DSCP IP header field. | 
| **dst-address** (*IP/netmask \| IP range* ; Default: ) | Matches packets whose destination is equal to the specified IP or falls into the specified IP range. | 
| **dst-address-list** (*name* ; Default: ) | Matches the destination address of a packet against a user-defined address-list. Supports only one list! | 
| **dst-address-type** (*unicast \| local \| broadcast \| multicast* ) | Matches destination address type:  | 
| **dst-limit** (*integer[/time],integer,dst-address \| dst-port \| src-address[/time]* ; Default: ) | Matches packets until a given rate is exceeded. Rate is defined as packets per time interval. As opposed to the limit matcher, every flow has its own limit. Flow is defined by a mode parameter. Parameters are written in the following format: `rate[/time],burst,mode[/expire]` . | 
| **dst-port** (*integer[-integer]: 0..65535* ; Default: ) | List of destination port numbers or port number ranges | 
| **fragment** (*yes\|no* ; Default: ) | Matches fragmented packets. The first (starting) fragment does not count. If connection tracking is enabled there will be no fragments as the system automatically assembles every packet. **IPv4** only. | 
| **header** (*Type[:Mode]; Mode=contains\|exact; Type=hop\|dst\|route\|frag\|ah\|esp\|none\|proto* ) | Matches IPv6 next-header. Two types of header matching are possible controlled by "mode" parameter:  **IPv6** only**.** | 
| **hop-limit** (Mode:Value; Mode=equal \| greater-than \| less-than \| not-equal; Value=0..255) | Matches hop limit field in the IPv6 header. **IPv6** only. | 
| **hotspot** (*auth \| from-client \| http \| local-dst \| to-client* ; Default: ) | Matches packets received from HotSpot clients against various HotSpot matchers.  **IPv4** Only. | 
| **icmp-options** (*integer:integer* ; Default: ) | Matches ICMP type: code fields | 
| **in-bridge-port** (*name* ; Default: ) | Actual interface the packet has entered the router if the incoming interface is a bridge. Works only if `use-ip-firewall` is enabled in bridge settings. | 
| **in-bridge-port-list** (*name* ; Default: ) | Set of interfaces defined in interface list. Works the same as in-bridge-port | 
| **in-interface** (*name* ; Default: ) | Interface the packet has entered the router | 
| **in-interface-list** (*name* ; Default: ) | Set of interfaces defined in interface list. Works the same as in-interface | 
| **ingress-priority** (*integer: 0..63* ; Default: ) | Matches the priority of an ingress packet. Priority may be derived from VLAN, WMM, DSCP, or MPLS EXP bit. read more | 
| **ipsec-policy** (*in \| out, ipsec \| none* ; Default: ) | Matches the policy used by IPsec. Value is written in the following format: . The direction is Used to select whether to match the policy used for decapsulation or the policy that will be used for encapsulation.**direction, policy**  For example, if a router receives an IPsec encapsulated Gre packet, then rule `ipsec-policy=in,ipsec` will match Gre packet, but a rule`ipsec-policy=in,none` will match the ESP packet. | 
| **ipv4-options** (*any \| loose-source-routing \| no-record-route \| no-router-alert \| no-source-routing \| no-timestamp \| none \| record-route \| router-alert \| strict-source-routing \| timestamp* ; Default: ) | Matches IPv4 header options.  **IPv4** only. | 
| **limit** (*integer,time,integer* ; Default: ) | Matches packets up to a limited rate (packet rate or bit rate). A rule using this matcher will match until this limit is reached. Parameters are written in the following format: `rate[/time],burst:mode` . | 
| **nth** (*integer,integer* ; Default: ) | Matches every nth packet: `nth=2,1` rule will match every first packet of 2, hence, 50% of all the traffic that is matched by the rule | 
| **out-bridge-port** (*name* ; Default: ) | Actual interface the packet leaves the router if the outgoing interface is a bridge. Works only if is enabled in bridge settings.**use-ip-firewall** | 
| **out-bridge-port-list** (*name* ; Default: ) | Set of interfaces defined in interface list. Works the same as out-bridge-port | 
| **out-interface** (; Default: ) | Interface the packet is leaving the router | 
| **out-interface-list** (*name* ; Default: ) | Set of interfaces defined in interface list. Works the same as out-interface | 
| **packet-mark** (*no-mark \| string* ; Default: ) | Matches packets marked via mangle facility with particular packet mark. If is set, the rule will match any unmarked packet.**no-mark** | 
| **packet-size** (*integer[-integer]:0..65535* ; Default: ) | Matches packets of specified size or size range in bytes. | 
| **per-connection-classifier** (*ValuesToHash:Denominator/Remainder* ; Default: ) | PCC matcher ( or Per Stream Classifier) allows dividing traffic into equal streams with the ability to keep packets with a specific set of options in one particular stream. Streams are hashed based on selected values to hash:  Read more >> | 
| **port** (*integer[-integer]: 0..65535* ; Default: ) | Matches if any (source or destination) port matches the specified list of ports or port ranges. Applicable only if `protocol` is TCP or UDP | 
| **priority** (*integer: 0..63* ; Default:) | Matches the packet's priority after a new priority has been set. Priority may be derived from VLAN, WMM, DSCP, MPLS EXP bit, or from the priority that has been set using the set-priority action. Read more | 
| **protocol** (*name or protocol ID* ; Default:**tcp** ) | Matches particular IP protocol specified by protocol name or number | 
| **psd** (*integer,time,integer,integer* ; Default: ) | Attempts to detect TCP and UDP scans. Parameters are in the following format `WeightThreshold, DelayThreshold, LowPortWeight, HighPortWeight` **IPv4** only. | 
| **random** (*integer: 1..99* ; Default: ) | Matches packets randomly with a given probability | 
| **src-address** (*Ip/Netmask, Ip range* ; Default: ) | Matches packets whose source is equal to a specified IP or falls into a specified IP range | 
| **src-address-list** (*name* ; Default: ) | Matches the source address of a packet against a user-defined address list. Supports only one list! | 
| **src-address-type** (*unicast \| local \| broadcast \| multicast \| blackhole \| prohibit \| unreachable ;* Default: ) | mote{ta{tableMatches source address type:  | 
| **src-port** (*integer[-integer]: 0..65535* ; Default: ) | List of source ports and ranges of source ports. Applicable only if a protocol is TCP or UDP | 
| **src-mac-address** (*MAC address* ; Default: ) | Matches the source MAC address of the packet | 
| **tcp-flags** (*ack \| cwr \| ece \| fin \| psh \| rst \| syn \| urg* ; Default: ) | Matches specified TCP flags  | 
| **tcp-mss** (*integer[-integer]: 0..65535* ; Default: ) | Matches TCP MSS value of an IP packet | 
| **time** (*time-time,sat \| fri \| thu \| wed \| tue \| mon \| sun* ; Default: ) | Allows to create a filter based on the packets' arrival time and date or, for locally generated packets, departure time and date. The matcher takes into account the time and timezone configured on the router. | 
| **tls-host** (*string* ; Default: ) | Allows matching HTTPS traffic based on TLS SNI hostname. Accepts GLOB syntax for wildcard matching. Note that the matcher will not be able to match the hostname if the TLS handshake frame is fragmented into multiple TCP segments (packets). Watch our video about this value. | 
| **ttl** (*integer: 0..255* ; Default: ) | Matches packets TTL value. **IPv4** Only. | 

## Stateful Properties

