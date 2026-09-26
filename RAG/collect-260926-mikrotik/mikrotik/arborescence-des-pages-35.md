---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-35
title: "arborescence-des-pages-35"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-35.md
source_anchor: ""
source_lines: [1, 66]
sha256: c3678c96b6d2d589809f63d7d653dc271c27d5f30afd1882b2583a08f4bab471
---

# arborescence-des-pages-35

## General Properties

**Sub-menu:** `/mpls`

| Property | Description | 
|---|---|
| **dynamic-label-range** (*range of integer[16..1048575]* ; Default:**16-1048575** ) | Range of Label numbers used for dynamic allocation. First 16 labels are reserved for special purposes (as defined in RFC). If you intend to configure labels statically then adjust dynamic default range not to include numbers that will be used in static configuration. | 
| **propagate-ttl** (*yes \| no* ; Default:**yes** ) | Whether to copy TTL values from IP header to MPLS header. If this option is set to **no** then hops inside MPLS cloud will be invisible from traceroutes. | 

## Forwarding Table

**Sub-menu:** `/mpls forwarding-table`


Entries in this sub-menu shows label bindings for specific routes that will be used in MPLS label switching. Properties in this menu are read-only

| Property | Description | 
|---|---|
| **prefix** (*IP/Mask* ) | Destination prefix for which labels are assigned | 
| **label** (*integer* ) | Ingress MPLS label | 
| **ldp** (*yes \| no* ) | Whether labels are LDP signaled | 
| **nexthops** () | Array of the nexthops,each entry in the array represents one ECMP nexthop. Array entry can contain several parameters:  | 
| **out-label** (*integer* ) | Label number which is added or switched to for outgoing packet. | 
| **packets** (*integer* ) | Number of packets matched by this entry | 
| **te-sender** |  | 
| **te-session**  |  | 
| **traffic-eng**  | Shows whether entry is signaled by RSVP-TE (Traffic Engineering) | 
| **type** *(string)* | Type of the entry, for example, "vpls", etc. | 
| **vpls** (*yes \| no* ) | Shows whether entry is used for VPLS tunnels. | 
| **vpn** |  | 
| **vrf** | Name of the VRF table this entry belongs to | 


For example we have forwarding table as shown below.

## Interface

**Sub-menu:** `/mpls interface`


This menu allows to configure maximum allowed MPLS MTUs (path MTU + MPLS tag size). Configuration of MPLS MTU is useful in cases when there are large variety of possible MTUs along the path. Configuring MPLS MTU to a minimum value that can pass all the hops will ensure that MPLS packet will not be silently dropped on the devices that do not support big enough MTU.

Listed entries are ordered, first entry (iterating from the top to the bottom) that matches the interface will be used.

Order of the entries is important due to possibility that different interface lists can contain the same interface and in addition that interface can be referenced directly.

Selection of the MPLS MTU happens in the following manner:

- If interface matched the entry from this table, then try to use configured MPLS MTU value
- If interface does not match any entry then consider MPLS MTU equal to L2MTU
- If interface does not support L2MTU, then consider MPLS MTU equal to L3 MTU

On MPLS ingress path MTU is chosen by min(MPLS MTU - tagsize, l3mtu). Which means that on interfaces that do not support L2MTU and default L3 MTU is set to 1500, max path MTU will be 1500 - tag size (interface will not be able to pass full IP frame without fragmentation). In such scenarios L3MTU must be increased by max observed tag size.

**Properties**

| Property | Description | 
|---|---|
| **comment** (*string* ; Default: ) | Short description of the interface | 
| **disabled** (*yes \| no* ; Default:**no** ) | If set to **yes** then this configuration is ignored. | 
| **interface** (*name* ; Default:) | Name of the interface or interface-list to match. | 
| **input**  (*yes \| no* ; Default:**yes** ) | Whether to allow MPLS input on the interace | 
| **mpls-mtu** (*integer [512..65535]* ; Default:**1508** ) | Option represents how big packets can be carried over the interface with added MPLS labels. `Read More >>` | 


Example configuration:
