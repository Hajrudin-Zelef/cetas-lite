---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-14-2
title: "arborescence-des-pages-14"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-14.md
source_anchor: ""
source_lines: [57, 134]
sha256: 11a780dd7999c0728d7384dd16247c3770902aa995285dd18e2cd6392f3b14a7
---

# arborescence-des-pages-14

| Property | Description | 
|---|---|
| **area** (*name* ; mandatory) | The OSPF area to which the matching interface will be associated. | 
| **auth** (*simple \| md5 \| sha1 \| sha256 \| sha384 \| sha512* ) | Specifies authentication method for OSPF protocol messages.  If the parameter is unset, then authentication is not used. | 
| **auth-id** (*integer* ) | The key id is used to calculate message digest (used when MD5 or SHA authentication is enabled). The value should match all OSPF routers from the same region. | 
| **authentication-key** (*string) *sensitive** | The authentication key to be used, should match on all the neighbors of the network segment. | 
| **comment** (*string)* |  | 
| **cost** (*integer [0..65535])* | Interface cost expressed as link state metric. | 
| **dead-interval** (*time* ; Default:**40s** ) | Specifies the interval after which a neighbor is declared dead. This interval is advertised in hello packets. This value must be the same for all routers on a specific network, otherwise, adjacency between them will not form | 
| **disabled** (*yes \| no)* |  | 
| **hello-interval** (*time* ; Default:**10s** ) | The interval between **HELLO** packets that the router sends out this interface. The smaller this interval is, the faster topological changes will be detected, the tradeoff is more OSPF protocol traffic. This value must be the same for all the routers on a specific network, otherwise, adjacency between them will not form. | 
| **instance-id** (i*nteger [0..255]* ; Default:**0** ) |  | 
| **passive** () | If enabled, then do not send or receive OSPF traffic on the matching interfaces | 
| **prefix-list** (name) | Name of the address list containing networks that should be advertised to the v3 interface. | 
| **priority** (*integer: 0..255* ; Default:**128** ) | Router's priority. Used to determine the designated router in a broadcast network. The router with the highest priority value takes precedence. Priority value 0 means the router is not eligible to become a designated or backup designated router at all. ROS v7 default value is 128 (defined in RFC), and the default value in ROS v6 was 1, keep this in mind when if you had strict priorities set for DR/BDR election. | 
| **retransmit-interval** (*time* ; Default:**5s** ) | Time interval the lost link state advertisement will be resent. When a router sends a link state advertisement (LSA) to its neighbor, the LSA is kept until the acknowledgment is received. If the acknowledgment was not received in time (see transmit-delay), the router will try to retransmit the LSA. | 
| **transmit-delay** (*time* ; Default:**1s** ) | Link-state transmit delay is the estimated time it takes to transmit a link-state update packet on the interface. | 
| **type** (*broadcast \| nbma \| ptp \| ptmp \| ptp-unnumbered \| virtual-link* ; Default:**broadcast** ) | the OSPF network type on this interface. Note that if interface configuration does not exist, the default network type is 'point-to-point' on PtP interfaces and 'broadcast' on all other interfaces.  | 
| **vlink-neighbor-id** (*IP* ) | Specifies the **router-id** of the neighbor which should be connected over the virtual link. | 
| **vlink-transit-area** (*name* ) | A non-backbone area the two routers have in common over which the virtual link will be established. Virtual links can not be established through stub areas. | 

## `/routing/ospf/lsa`

List of all the LSAs currently in the LSA database.

| Read-only Property | Description | 
|---|---|
| **age** (*integer* ) | How long ago (in seconds) the last update occurred | 
| **area** (*string* ) | The area this LSA belongs to. | 
| **body** (*string* ) |  | 
| **checksum** (*string* ) | LSA checksum | 
| **dynamic** (*yes \| no* ) |  | 
| **flushing** (*yes \| no* ) |  | 
| **id** (*IP* ) | LSA record ID | 
| **instance** (*string* ) | The instance name this LSA belongs to. | 
| **link** (*string* ) |  | 
| **link-instance-id** (*IP* ) |  | 
| **originator** (*IP* ) | An originator of the LSA record. | 
| **self-originated** (*yes \| no* ) | Whether LSA originated from the router itself. | 
| **sequence**  (*string* ) | A number of times the LSA for a link has been updated. | 
| **type** (*string* ) |  | 
| **wraparound** (*string* ) |  | 

## `/routing/ospf/neighbor`

List of currently active OSPF neighbors.

| Read-only Property | Description | 
|---|---|
| **address** (*IP* ) | An IP address of the OSPF neighbor router | 
| **adjacency** (*time* ) | Elapsed time since adjacency was formed | 
| **area** (*string* ) |  | 
| **bdr** (*string* ) | An IP address of the Backup Designated Router | 
| **comment** (*string* ) |  | 
| **db-summaries** (*integer* ) |  | 
| **dr**  (*IP* ) | An IP address of the Designated Router | 
| **dynamic**  (*yes \| no* ) |  | 
| **inactive**  (*yes \| no* ) |  | 
| **instance** (*string* ) |  | 
| **ls-requests** (*integer* ) |  | 
| **ls-retransmits** (*integer* ) |  | 
| **priority** (*integer* ) | Priority configured on the neighbor | 
| **router-id** (*IP* ) | neighbor router's **RouterID** | 
| **state** (*down \| attempt \| init \| 2-way \| ExStart \| Exchange \| Loading \| full* ) |  | 
| **state-changes** (*integer* ) | Total count of OSPF state changes since neighbor identification | 

## `/routing/ospf/static-neighbor`

Static configuration of the OSPF neighbors. Required for non-broadcast multi-access networks.

| Read-only Property | Description | 
|---|---|
| **address** (*IP%iface* ; mandatory ) | The unicast IP address and an interface, that can be used to reach the IP of the neighbor. For example, `address=1.2.3.4%ether1` indicates that a neighbor with IP*1.2.3.4* is reachable on the*ether1* interface. | 
| **area** (*name* ; mandatory ) | Name of the area the neighbor belongs to. | 
| **comment** (*string)* |  | 
| **disabled** (*yes \| no)* |  | 
| **instance-id** (*integer [0..255]* ; Default: 0) |  | 
| **poll-interval** (*time* ; Default:**2m** ) | How often to send hello messages to the neighbors which are in a "down" state (i.e. there is no traffic from them) |
