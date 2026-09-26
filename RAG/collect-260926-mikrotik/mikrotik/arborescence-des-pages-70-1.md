---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-70-1
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-70.md
source_anchor: ""
source_lines: [1, 84]
sha256: 50787ce8afc4f073fedb043e9efa17f9dff3129462193e891ce84a44f4e19ec3
---

# Overview


Point to Point over Ethernet (PPPoE) is simply a method of encapsulating PPP packets into Ethernet frames. PPPoE is an extension of the standard Point to Point Protocol (PPP) and it the successor of PPPoA. PPPoE standard is defined in RFC 2516. The PPPoE client and server work over any Layer2 Ethernet level interface on the router, for example, Wireless, Ethernet, EoIP, etc. Generally speaking, PPPoE is used to hand out IP addresses to clients based on authentication by username (and also if required, by workstation) as opposed to workstation only authentication where static IP addresses or DHCP are used. It is advised not to use static IP addresses or DHCP on the same interfaces as PPPoE for obvious security reasons.

# Introduction

PPPoE provides the ability to connect a network of hosts over a simple bridging access device to a remote Access Concentrator.

Supported connections:

- MikroTik RouterOS PPPoE client to any PPPoE server;
- MikroTik RouterOS server (access concentrator) to multiple PPPoE clients (clients are available for almost all operating systems and most routers);

# PPPoE Operation

PPPoE has two distinct stages(phases):

1. Discovery phase;
2. Session phase;

## Discovery phase

There are four steps to the Discovery stage. When it completes, both peers know the PPPoE *SESSION_ID* and the peer's Ethernet address, which together define the PPPoE session uniquely:

1.  **PPPoE Active Discovery Initialization (PADI) -**  The PPPoE client sends out a*PADI* packet to the broadcast address. This packet can also populate the "service-name" field if a service name has been entered in the dial-up networking properties of the PPPoE client. If a service name has not been entered, this field is not populated
2. **PPPoE Active Discovery Offer (PADO) -**  The PPPoE server, or Access Concentrator, should respond to the*PADI* with a*PADO* if the Access Concentrator is able to service the "service-name" field that had been listed in the*PADI* packet. If no "service-name" field had been listed, the Access Concentrator will respond with a*PADO* packet that has the "service-name" field populated with the service names that the Access Concentrator can service. The*PADO* packet is sent to the unicast address of the PPPoE client
3. **PPPoE Active Discovery Request (PADR) -**  When a*PADO* packet is received, the PPPoE client responds with a*PADR* packet. This packet is sent to the unicast address of the Access Concentrator. The client may receive multiple*PADO* packets, but the client responds to the first valid*PADO* that the client received. If the initial*PADI* packet had a blank "service-name" field filed, the client populates the "service-name" field of the*PADR* packet with the first service name that had been returned in the*PADO* packet.
4. **PPPoE Active Discovery Session Confirmation (PADS) -**  When the*PADR* is received, the Access Concentrator generates a unique session identification (ID) for the Point-to-Point Protocol (PPP) session and returns this ID to the PPPoE client in the*PADS* packet. This packet is sent to the unicast address of the client.

PPPoE session termination:

- **PPPoE Active Discovery Terminate (PADT) -** Can be sent anytime after a session is established to indicate that a PPPoE session terminated. It can be sent by either server or client.

## Session phase

When the discovery stage is completed, both peers know *PPPoE Session ID* and other peer's *Ethernet (MAC) address* which together defines the PPPoE session. PPP frames are encapsulated in PPPoE session frames, which have Ethernet frame type **0x8864**. 

When a server sends confirmation and a client receives it, PPP Session is started that consists of the following stages:

1. **LCP negotiation** stage
2. **Authentication (CHAP/PAP)** stage
3. **IPCP negotiation** stage - where the client is assigned an IP address.

If any process fails, the LCP negotiation establishment phase is started again.


PPPoE server sends *Echo-Request* packets to the client to determine the state of the session, otherwise, the server will not be able to determine that session is terminated in cases when a client terminates session without sending *Terminate-Request* packet.

# MTU

Typically, the largest Ethernet frame that can be transmitted without fragmentation is 1500 bytes. PPPoE adds another 6 bytes of overhead and the PPP field adds two more bytes, leaving 1492 bytes for IP datagram. Therefore max PPPoE MRU and MTU values must not be larger than 1492.

TCP stacks try to avoid fragmentation, so they use an MSS (Maximum Segment Size). By default, MSS is chosen as MTU of the outgoing interface minus the usual size of the TCP and IP headers (40 bytes), which results in 1460 bytes for an Ethernet interface. Unfortunately, there may be intermediate links with lower MTU which will cause fragmentation. In such a case TCP stack performs path MTU discovery. Routers that cannot forward the datagram without fragmentation are supposed to drop the packet and send *ICMP-Fragmentation-Required* to originating host. When a host receives such an ICMP packet, it tries to lower the MTU. This should work in the ideal world, however in the real world many routers do not generate fragmentation-required datagrams, also many firewalls drop all ICMP datagrams.

The workaround for this problem is to adjust MSS if it is too big.

# PPPoE Client

## Properties

| Property | Description | 
|---|---|
| **ac-name** (*string* ; Default:**""** ) | Access Concentrator name, this may be left blank and the client will connect to any access concentrator on the broadcast domain | 
| **add-default-route** (*yes\|no* ; Default:**no** ) | Enable/Disable whether to add default route automatically | 
| **allow** (*mschap2\|mschap1\|chap\|pap* ; Default:**mschap2,mschap1,chap,pap** ) | allowed authentication methods, by default all methods are allowed | 
| **default-route-distance** (*byte [0..255]* ; Default:**1** ) | sets distance value applied to auto created default route, if add-default-route is also selected | 
| **dial-on-demand** (*yes\|no* ; Default:**no** ) | connects to AC only when outbound traffic is generated. If selected, then route with gateway address from 10.112.112.0/24 network will be added while connection is not established. | 
| **interface** (*string* ; Default: ) | interface name on which client will run | 
| **keepalive-timeout** (*integer* ; Default:**60** ) | Sets keepalive timeout in seconds. Keepalive-timeout=disabled option is added to disable echo packages on link, this will stop negotiation on mtu. At this point 1600 can be set. Check if server supports, most probably path mtu is blocking jumbo frames. If path of mtu allows jumbo frames , the mtu of 1600 will also work. | 
| **max-mru** (*integer* ; Default:**1460** ) | Maximum Receive Unit | 
| **max-mtu** (*integer* ; Default:**1460** ) | Maximum Transmission Unit | 
| **mrru** (*integer: 512..65535\|disabled* ; Default:**disabled** ) | maximum packet size that can be received on the link. If a packet is bigger than tunnel MTU, it will be split into multiple packets, allowing full size IP or Ethernet packets to be sent over the tunnel. | 
| **name** (*string* ; Default:**pppoe-out[i]** ) | name of the PPPoE interface, generated by RouterOS if not specified | 
| **password** (*string* ; Default: )*sensitive* | password used to authenticate | 
| **profile** (*string* ; Default:**default** ) | Specifies which PPP profile configuration will be used when establishing the tunnel. | 
| **service-name** (*string* ; Default:**""** ) | specifies the service name set on the access concentrator, can be left blank to connect to any PPPoE server | 
| **use-peer-dns** (*yes\|no* ; Default:**no** ) | enable/disable getting DNS settings from the peer | 
| **user** (*string* ; Default:**""** ) | username used for authentication | 

## Status

Command `/interface pppoe-client monitor` will display current PPPoE status.

