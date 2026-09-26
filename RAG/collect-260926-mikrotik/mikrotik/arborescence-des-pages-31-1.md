---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-31-1
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-31.md
source_anchor: ""
source_lines: [1, 128]
sha256: 501a12cfc447ef632d90bacc52b4dd98957dce298601e5763c32135f64ba2aba
---

# Summary

The Border Gateway Protocol (BGP) allows setting up an inter-domain dynamic routing system that automatically updates routing tables of devices running BGP in case of network topology changes.

BGP is an inter-autonomous system routing protocol based on the distance-vector algorithm. It is used to exchange routing information across the Internet and is the only protocol that is designed to deal with a network of the Internet's size and the only protocol that can deal well with having multiple connections to unrelated routing domains.

BGP is designed to allow for sophisticated administrative routing policies to be implemented. It does not exchange information about network topology but rather reachability information. As such, BGP is better suited to inter-AS environments and special cases like informational feeds. If you just need to enable dynamic routing in your network, consider OSPF instead.

The feature is not supported on SMIPS devices (hAP lite, hAP lite TC, and hAP mini).


Standards and Technologies:

- RFC 4271 Border Gateway Protocol 4
- RFC 4456 BGP Route Reflection
- RFC 5065 Autonomous System Confederations for BGP
- RFC 1997 BGP Communities Attribute
- RFC 8092 BGP Large Communities
- RFC 4360, 5668 BGP Extended Communities
- RFC 2385 TCP MD5 Authentication for BGPv4
- RFC 5492 Capabilities Advertisement with BGP-4
- RFC 2918 Route Refresh Capability
- RFC 4760 Multiprotocol Extensions for BGP-4
- RFC 2545 Use of BGP-4 Multiprotocol Extensions for IPv6 Inter-Domain Routing
- RFC 4893 BGP Support for Four-octet AS Number Space
- RFC 4364 BGP/MPLS IP Virtual Private Networks (VPNs)
- RFC 4761 Virtual Private LAN Service (VPLS) Using BGP for Auto-Discovery and Signalling
- RFC 6286 - AS-wide Unique BGP Identifier for BGP-4
- RFC 4273 - SNMP peer table monitoring (OID 1.3.6.1.2.1.15.3.1) (IPv4 only)
- RFC 6793 - 4-byte ASN support and Aggregator attribute.

# BGP Terminology

- AS - Autonomous System
- ASN - Autonomous System Number
- NLRI - Network Layer Reachability Information is what is being exchanged between BGP peers and represents how to reach the prefixes.
- IGP - Interior Gateway Protocol
- EGP - Exterior Gateway protocol
- RR - Route reflector is the router in the BGP network that reflects advertisements to all the neighbors, avoiding the requirement for full BGP mesh.
- Route server - is the BGP router that does not participate in traffic forwarding. Routes are typically not even installed in the FIB.
- loopback address - a /32 address configured on a dummy bridge interface, that can act as a loopback.

# BGP Basics

BGP routers exchange reachability information by means of a transport protocol, which in the case of BGP is TCP (port 179). Upon forming a TCP connection these routers exchange **OPEN** messages to negotiate and confirm supported capabilities.

After agreeing on capabilities to use, the session is considered to be established and peers can start to exchange NLRIs via **UPDATE** messages. This information contains an indication of what sequence of full paths (BGP AS numbers) the route should take in order to reach the destination network (NLRI prefix).

The peers initially exchange their full routing tables and after the initial exchange, incremental updates are sent as the routing tables change. Thus, BGP does not require a periodic refresh of the entire BGP routing table.

BGP maintains the routing table version number which must be the same between any two given peers for the duration of the connection.

**KEEPALIVE** messages are sent periodically to ensure that the connection is up and running, if **KEEPALIVE** messages are not received within the **Hold Time** interval, the connection will be closed.

To respond to errors or special conditions, **NOTIFICATION** messages can be generated and sent to the remote peer, notification message type also indicates whether the connection should be immediately closed.

There can be two types of BGP connections:

- **iBGP** - is an "internal" link connecting peers from the same AS
- **eBGP** - is an "external" link connecting peers belonging to two different AS-es

A particular AS might have multiple BGP speakers and provide transit service to other AS-es. This implies that BGP speakers must maintain a consistent view of routing within the AS. A consistent view of the routes exterior to the AS is provided by having all BGP routers within the AS establish direct iBGP connections with each other (full mesh) or by utilizing a Router Reflector setup.

Using a set of administrative policies BGP speakers within the AS come to an agreement as to which entry/exit point to use for a particular destination. This information is communicated to the interior routers of the AS using the interior routing protocol (IGP), for example, OSPF, RIP, or static routing. In certain setups, iBGP can take the IGP protocol role as well.

For certain BGP attributes handling behavior may change depending on what type of connection is set up, for example, the LOCAL-PREF attribute is not advertised to eBGP peers.

RouterOS divides configuration and session monitoring into four menus:

- instance menu (`/routing/bgp/instance` )

- connection menu (`/routing/bgp/connection` )
- sessions menu(`/routing/bgp/session` )
- template menu (`/routing/bgp/template` )

## Instance Menu

Starting from ROSv7.20, instead of auto detecting instance based on router-ids, BGP routing instances are now explicitly defined in instance menu.

BGP routing instance is necessary for best path route selection and other instance dependant features like VPN, EVPN and so on.

## Connection Menu

Let's look at a very basic eBGP configuration example assuming, that Router1 IP is 192.168.1.1, AS 65531 and Router2 IP 192.168.1.2, AS 65532:

The BGP connection menu defines BGP outgoing connections as well as acts as a template matcher for incoming BGP connections.

 `local.role` parameter is used to indicate that this connection will be the eBGP. Also, notice that the connection does not require a remote AS number to be specified, RouterOS can determine a remote AS number dynamically from the first received **OPEN** message.

The parameter equivalent to other vendors and older RouterOS "update-source" is "`local.address`". In most cases, it can be left unconfigured, and let the router determine the address.

When a local address is not specified, BGP will try to guess the local address depending on the current setup:

- if the peer is iBGP
  - if loopback available 
    - pick the highest loopback address
  - if loopback is not available
    - pick any highest IP address on the router
- if loopback available 
- if the peer is eBGP
  - if a remote peer's IP is not from a directly connected network:
    - and multihop is not set, then throw an error
    - and multihop is enabled:
      - if loopback available
        - pick the highest loopback address
      - if loopback is not available
        - pick any highest IP address on the router
    - if loopback available
  - if a remote peer's IP is from a directly connected network:
    - and multihop is not set:
      - pick the local routers IP address from that connected network
    - and multihop is set:
      - if loopback available
        - pick the highest loopback address
      - if loopback is not available
        - pick any highest IP address on the router
    - if loopback available
  - and multihop is not set:
- if a remote peer's IP is not from a directly connected network:

In addition to connection-specific parameters, template-specific parameters are also directly exposed in this menu, for easier configuration in simple scenarios (when templates are not necessary).

Listening on subnets should not be enabled in unsafe environments, denial of service is possible with such configuration. Firewall must be configured to protect the router.

See "listen" parameter for more details.

## Session Menu

