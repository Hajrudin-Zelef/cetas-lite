---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-72-1
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-72.md
source_anchor: ""
source_lines: [1, 57]
sha256: a3c9b701683bf9a9b9bd9faad61d209cc7aaa5dd3c0b467bcad0cc2619a27511
---

# Overview

Layer Two Tunneling Protocol "L2TP" extends the PPP model by allowing the L2 and PPP endpoints to reside on different devices interconnected by a packet-switched network. L2TP includes PPP authentication and accounting for each L2TP connection. Full authentication and accounting of each connection may be done through a RADIUS client or locally. L2TP traffic uses UDP protocol for both control and data packets. UDP port 1701 is used only for link establishment, further traffic is using any available UDP port (which may or may not be 1701). This means that L2TP can be used with most firewalls and routers (even with NAT) by enabling UDP traffic to be routed through the firewall or router. L2TP standard is defined in RFC 2661. The L2TPv3 support added in 7.1 version. Support IPv4, IPv6.

# Introduction

It may be useful to use L2TP just as any other tunneling protocol with or without encryption. The L2TP standard says that the most secure way to encrypt data is using L2TP over IPsec (Note that it is the default mode for Microsoft L2TP client) as all L2TP control and data packets for a particular tunnel appear as homogeneous UDP/IP data packets to the IPsec system.

Multilink PPP (MP) is supported in order to provide MRRU (the ability to transmit full-sized 1500 and larger packets) and bridging over PPP links (using Bridge Control Protocol (BCP) that allows sending raw Ethernet frames over PPP links). This way it is possible to setup bridging without EoIP. The bridge should either have an administratively set MAC address or an Ethernet-like interface in it, as PPP links do not have MAC addresses.

L2TP does not provide encryption mechanisms for tunneled traffic. IPsec can be used for additional security layers.

# L2TP Client

## Properties

| Property | Description | 
|---|---|
| **add-default-route** (*yes \| no* ; Default:**no** ) | Whether to add L2TP remote address as a default route. | 
| **allow** (*mschap2 \| mschap1 \| chap \| pap* ; Default:**mschap2, mschap1, chap, pap** ) | Allowed authentication methods. | 
| **connect-to** (*IP\|IPv6* ; Default: ) | Remote address of L2TP server (if the address is in VRF table, VRF should be specified) | 
| **comment** (*string* ; Default: ) | Short description of the tunnel. | 
| **default-route-distance** (*byte* ; Default: ) | Since v6.2, sets distance value applied to auto created default route, if add-default-route is also selected | 
| **dial-on-demand** (*yes \| no* ; Default:**no** ) | connects only when outbound traffic is generated. If selected, then route with gateway address from 10.112.112.0/24 network will be added while connection is not established. | 
| **disabled** (*yes \| no* ; Default:**yes** ) | Enables/disables tunnel. | 
| **keepalive-timeout** (*integer [1..4294967295]* ; Default:**60s** ) | Since v6.0rc13, tunnel keepalive timeout in seconds. | 
| **max-mru** (*integer* ; Default:**1450** ) | Maximum Receive Unit. Max packet size that L2TP interface will be able to receive without packet fragmentation. | 
| **max-mtu** (*integer* ; Default:**1450** ) | Maximum Transmission Unit. Max packet size that L2TP interface will be able to send without packet fragmentation. | 
| **mrru** (*disabled \| integer* ; Default:**disabled** ) | Maximum packet size that can be received on the link. If a packet is bigger than tunnel MTU, it will be split into multiple packets, allowing full size IP or Ethernet packets to be sent over the tunnel. | 
| **name** (*string* ; Default: ) | Descriptive name of the interface. | 
| **password** (*string* ; Default:**""** )*sensitive* | Password used for authentication. | 
| **profile** (*name* ; Default:**default-encryption** ) | Specifies which PPP profile configuration will be used when establishing the tunnel. | 
| **user** (*string* ; Default: ) | User name used for authentication. | 
| **use-ipsec** (*yes \| no* ; Default:**no** ) | When this option is enabled, dynamic IPSec peer configuration and policy (transport mode) is added to encapsulate L2TP connection into IPSec tunnel. Multiple L2tp/ipsec clients behind the same NAT will not work in this mode. To achieve such scenario, disable use-ipsec and set static policies for clients with enabled tunnel=yes, level=unique settings. | 
| **allow-fast-path** (*yes \| no* ; Default: ) | Allow to forward packets without additional processing in the Linux kernel. | 
| **l2tp-proto-version** ( l2tpv2*\| l2tpv3-ip \| l2tpv3-udp \| l2tpv* ; Default:**l2tpv2** ) | Specify protocol version. | 
| **l2tpv3-cookie-length** ( 0*\| 4-bytes \| 8-bytes* ; Default:**0** ) | Configures an L2TPv3 pseudowire static session cookie. | 
| **l2tpv3-digest-hash** (*md5 \| none \| sha1* ; Default:**md5** ) | Specifies which hash function to be used. | 
| **use-peer-dns** (*yes \| no \| exclusively* ; Default:**no** ) | To use peer dns. | 
| **copy-from** | To copy created peer. | 
| **src-address** | Specify source address. | 
| **l2tpv3-circuit-id**  | Set the virtual circuit identifier to bind the one end of the L2TPv3 control channel. | 
| **ipsec-secret** (*string* ; Default: )*sensitive* | Preshared key used when use-ipsec is enabled. | 

# L2TP Server

An interface is created for each tunnel established to the given server. There are two types of interfaces in the L2TP server's configuration

- Static interfaces are added administratively if there is a need to reference the particular interface name (in firewall rules or elsewhere) created for the particular user;
- Dynamic interfaces are added to this list automatically whenever a user is connected and its username does not match any existing static entry (or in case the entry is active already, as there can not be two separate tunnel interfaces referenced by the same name);

Dynamic interfaces appear when a user connects and disappear once the user disconnects, so it is impossible to reference the tunnel created for that use in router configuration (for example, in firewall), so if you need persistent rules for that user, create a static entry for him/her. Otherwise, it is safe to use a dynamic configuration.

in both cases PPP users must be configured properly - static entries do not replace PPP configuration.

## Properties

