---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-32-3
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-32.md
source_anchor: ""
source_lines: [229, 370]
sha256: dea463f59695b27b8b3f3873f706c0b3782059cce874acfc8852c5d7086b1738
---

# Overview

1. ipv4 - wait X seconds, if no changes, then use the IPv4 LDP session and distribute IPv4 labels
2. ipv4+ds6 - wait for IPv6 hello, dual-stack element indicates that there should be IPv6
3. ipv6 - wait X seconds, if no changes, then use the IPv6 LDP session and distribute IPv6 labels
4. ipv6+ds6 - use IPv6 LDP session and distribute IPv6 labels
5. ipv4,ipv6 - use IPv6 LDP session and distribute IPv4 and IPv6 labels
6. ipv4,ipv6+ds6 - use IPv6 LDP session and distribute IPv4 and IPv6 labels

# Property Reference

## LDP Instance

**Sub-menu:** `/mpls`

**Properties**

| Property | Description | 
|---|---|
| **afi (***ip \| ipv6***;** Default:**)** | Determines supported address families by the instance. | 
| **comments** (*string* ; Default: ) | Short description of the entry | 
| **disabled** (*yes \| no* ; Default:**no** ) |  | 
| **distribute-for-default** (*yes \| no* ; Default: no) | Defines whether to map label for the default route. | 
| **hop-limit** (*integer[0..255]* ; Default: ) | Max hop limit used for loop detection. Works in combination with the **loop-detect** property. | 
| **loop-detect** (*yes \| no* ; Default: ) | Defines whether to run LSP loop detection. Will not work correctly if not enabled on all LSRs. Should be used only on non-TTL networks such as ATMs. | 
| **lsr-id** (*IP* ; Default: ) | Unique label switching router's ID. | 
| **path-vector-limit** (*IP* ; Default: ) | Max path vector limit used for loop detection. Works in combination with the **loop-detect** property. | 
| **preferred-afi** (ip \| ipv6; Default:**ipv6** ) | Determining which address family connection is preferred. Value is also set in dual-stack element (if used). | 
| **transport-addresses** (*IP* ; Default: ) | Specifies LDP session connections origin addresses and also advertises these addresses as transport addresses to LDP neighbors. | 
| **use-explicit-null** (*yes \| no* ; Default: no) | Whether to distribute explicit-null label bindings. | 
| **vrf (***name* ; Default: main**)** | Name of the VRF table this instance will operate on. | 

## Interface

**Sub-menu:** `/mpls ldp interface`

| Property | Description | 
|---|---|
| **afi (***ip \| ipv6***;** Default:**)** | Determines interface address family. Only AFIs that are configured as supported by the instance is taken into account. If the value is not explicitly specified then it is considered to be equal to the instance-supported AFIs. | 
| **accept-dynamic-neighbors** (*yes \| no* ; Default:) | Defines whether to discover neighbors dynamically or use only statically configured in LDP neighbors menu | 
| **comments** (*string* ; Default: ) | Short description of the entry | 
| **disabled** (*yes \| no* ; Default:**no** ) |  | 
| **hello-interval**  (*string* ; Default: ) | The interval between hello packets that the router sends out on specified interface/s. The default value is 5s. | 
| **hold-time**  (*string* ; Default: ) | Specifies the interval after which a neighbor discovered on the interface is declared as not reachable. The default value is 15s. | 
| **interface**  (*string* ; Default: ) | Name of the interface or interface list where LDP will be listening. | 
| **transport-addresses**  (List of*IPs* ; Default: ) | Used transport addresses if differs from LDP Instance settings. | 

## Neighbors

**Sub-menu:** `/mpls ldp neighbor`

List of discovered and statically configured LDP neighbors.

**Properties**

| Property | Description | 
|---|---|
| **comments** (*string* ; Default: ) | Short description of the entry | 
| **disabled** (*yes \| no* ; Default:**no** ) |  | 
| **send-targeted**  (*yes \| no* ; Default: ) | Specifies whether to try to send targeted hellos, used for targeted (not directly connected) LDP sessions. | 
| **transport**  (*IP* ; Default: ) | Remote transport address. | 

**Read-only Properties**

| Property | Description | 
|---|---|
| **active-connect** (*yes \| no* ) | Indicates that active role have been selected and the router is trying to establish the session. | 
| **addresses** (*list of IPs* ) | List of discovered addresses on the neighbor | 
| **inactive** (*yes \| no* ) | Whether binding is active and can be selected as a candidate for forwarding. | 
| **dynamic** (*yes \| no* ) | Whether entry was dynamically added | 
| **local-transport** (*IP* ) | Selected local transport address. | 
| **on-demand** (*yes \| no* ) | Downstream On Demand label distribution | 
| **operational** (*yes \| no* ) | Indicates whether the peer is operational. | 
| **passive** (*yes \| no* ) | Indicates whether the peer is in a passive role. | 
| **passive-wait** (*yes \| no* ) | Indicates whether the peer is in a passive role and currently is waiting for the session to be initialized. | 
| **path-vector-limit** (*integer* ) |  | 
| **peer** (*IP:integer* ) | LSR-ID and label space of the neighbor | 
| **sending-targeted-hello** (*yes \| no* ) | Whether targeted hellos are being sent to the neighbor. | 
| **throttled** (*yes \| no* ) | Indicates whether session is in throttled state. Session is throttled after initialization failure, max throttle time 120s. | 
| **used-afi** (*yes \| no* ) | Used transport AFI | 
| **vpls** (*yes \| no* ) | Whether neighbor is used by VPLS tunnel | 

## Accept Filter

**Sub-menu:** `/mpls ldp accept-filter`

List of label bindings that should be accepted from LDP neighbors.

| Property | Description | 
|---|---|
| **accept** (*yes \| no* ; Default:**no** ) | Whether to accept label bindings from the neighbors for the specified prefix. | 
| **comments** (*string* ; Default: ) | Short description of the entry | 
| **disabled** (*yes \| no* ; Default:**no** ) |  | 
| **neighbor** (*string* ; Default: ) | Neighbor to which this filter applies. | 
| **prefix**  (*IP/mask* ; Default: ) | Prefix to match. | 
| **vrf**  (name; Default: ) |  | 

## Advertise Filter

**Sub-menu:** `/mpls ldp advertise-filter`

List of label bindings that should be advertised to LDP neighbors.

| Property | Description | 
|---|---|
| **advertise** (*yes \| no* ; Default:**no** ) | Whether to advertise label bindings to the neighbors for the specified prefix. | 
| **comments** (*string* ; Default: ) | Short description of the entry | 
| **disabled** (*yes \| no* ; Default:**no** ) |  | 
| **neighbor** (*string* ; Default: ) | Neighbor to which this filter applies. | 
| **prefix**  (*IP/mask* ; Default: ) | Prefix to match. | 
| **vrf**  (name; Default: ) |  | 

## Local Mapping

**Sub-menu:** `/mpls local-mapping`

This sub-menu shows labels bound to the routes locally in the router. In this menu also static mappings can be configured if there is no intention to use LDP dynamically.

**Properties**

| Property | Description | 
|---|---|
| **comments** (*string* ; Default: ) | Short description of the entry | 
| **disabled** (*yes \| no* ; Default:**no** ) |  | 
| **dst-address** (*IP/Mask* ; Default: ) | Destination prefix the label is assigned to. | 
| **label** (*integer[0..1048576] \| alert \| expl-null \| expl-null6 \| impl-null \| none* ; Default: ) | Label number assigned to destination. | 
| **vrf (***name* ; Default: main**)** | Name of the VRF table this mapping belongs to. | 

**Read-only Properties**

| Property | Description | 
|---|---|
| **adv-path** () |  | 
| **inactive** (*yes \| no* ) | Whether binding is active and can be selected as a candidate for forwarding. | 
| **dynamic** (*yes \| no* ) | Whether entry was dynamically added | 
| **egress** (*yes \| no* ) |  | 
| **gateway** (*yes \| no* ) | Whether the destination is reachable through the gateway. | 
| **local** (*yes \| no* ) | Whether the destination is locally reachable on the router | 
| **peers** (*IP:label_space* ) | IP address and label space of the peer to which this entry was advertised. | 

## Remote Mapping

**Sub-menu:** `/mpls remote-mapping`

