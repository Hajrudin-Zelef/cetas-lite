---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-72-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-72.md
source_anchor: ""
source_lines: [58, 131]
sha256: 838bd1af0b9db370e2d6024368fc5d53bec42a41ca35ffedf9f2bdc7b21b729d
---

# Overview

| Property | Description | 
|---|---|
| **authentication** (*pap \| chap \| mschap1 \| mschap2* ; Default:**mschap1,mschap2** ) | Authentication methods that server will accept. | 
| **default-profile** (*name* ; Default:**default-encryption** ) | default profile to use | 
| **enabled** (*yes \| no* ; Default:**no** ) | Defines whether L2TP server is enabled or not. | 
| **max-mru** (*integer* ; Default:**1450** ) | Maximum Receive Unit. Max packet size that L2TP interface will be able to receive without packet fragmentation. | 
| **keepalive-timeout** (*integer* ; Default:**30** ) | If server during keepalive-timeout period does not receive any packets, it will send keepalive packets every second, five times. If the server still does not receive any response from the client, then the client will be disconnected after 5 seconds. Logs will show 5x "LCP missed echo reply" messages and then disconnect. | 
| **max-mtu** (*integer* ; Default:**1450** ) | Maximum Transmission Unit. Max packet size that L2TP interface will be able to send without packet fragmentation. | 
| **use-ipsec** (*no \| yes \| require* ; Default:**no** ) | When this option is enabled, dynamic IPSec peer configuration is added to suite most of the L2TP road-warrior setups. When require is selected server will accept only those L2TP connection attempts that were encapsulated in the IPSec tunnel. | 
| **ipsec-secret** (*string* ; Default: )*sensitive* | Preshared key used when use-ipsec is enabled | 
| **accept-proto-version** ( all*\|*  l2tpv2*\|*  l2tpv3; Default:**all** ) cli-only | Specify protocol version. | 
| **accept-pseudowire-type** ( all*\|* ether*\|* ppp; Default:**all** ) | Set the pseudowire signaling protocol for specific pseudowire type. | 
| **allow-fast-path** (*no \| yes* ; Default:**no** ) | To forward packets without additional processing in the Linux kernel. | 
| **caller-id-type** ( ip-address*\|* number; Default:**ip-address** ) | If same source IP address is used for multiple clients set id type to number. | 
| **max-sessions** ( unlimited / number; Default:**unlimited** ) | Set number of needed sessions. | 
| **one-session-per-host**   (*no \| yes \|*  ; Default:**no** ) | To allow one session per host. | 
| **l2tpv3-circuit-id** (Default: ) | Set the virtual circuit identifier to bind the one end of the L2TPv3 control channel. | 
| **l2tpv3-cookie-length** (0*\|* 4-bytes*\|* 8-bytes; Default:**0** ) | Configures an L2TP pseudowire static session cookie. | 
| **l2tpv3-digest-hash** ( md5*\|* none*\|* sha1; Default:**md5** ) | Specifies which hash function to be used. | 
| **l2tpv3-ether-interface-list** (Default: ) | Set your interface list for example the default ones- all, dynamic, none, static. | 
| **mrru** (*disabled \| integer* ; Default:**disabled** ) | Maximum packet size that can be received on the link. If a packet is bigger than tunnel MTU, it will be split into multiple packets, allowing full size IP or Ethernet packets to be sent over the tunnel. | 

# Quick Example

## L2TP Server

On the servers side we will enable L2TP-server and create a PPP profile for a particular user:

## L2TP Client

L2TP client setup in the RouterOS is very simple. In the following example, we already have a preconfigured 3 unit setup. We will take a look more detailed on how to set up L2TP client with username "MT-User", password "StrongPass" and server 192.168.51.3:

# L2TP Ether

# Overview

Layer 2 Tunnel Protocol Version 3 (L2TPv3) is a draft from the Internet Engineering Task Force (IETF) working group. It introduces various improvements to the original L2TP, allowing the encapsulation of Layer 2 (L2) payloads within L2TP. More precisely, L2TPv3 outlines the protocol for tunnelling Layer 2 payloads across an IP core network using L2 virtual private networks (VPNs).

To establish an **L2TP Ether** tunnel, the **L2TP Ether interface** must be created on the **client side**, while the **L2TP server** must be enabled on the **remote (server) side**. Once both sides are configured correctly, a **dynamic interface** is automatically created between them, forming a transparent Layer 2 connection across the IP network.

##### **Server Side (L2TP Server)**

`/interface l2tp-server server set enabled=yes`

##### **Client Side (L2TP Ether Interface)**

/interface l2tp-ether add connect-to=1.1.1.1 disabled=no

## Properties

| Property | Description | 
|---|---|
| **connect-to** (*IP* ; Default: ) | Remote address of L2TP server. | 
| **comment** ( st*ring* ; Default: ) | Short description of the tunnel. | 
| **disabled** (*yes \| no* ; Default:**yes** ) | Enables/disables tunnel. | 
| **mac-address** ( string; Default:**auto** ) | Set desired mac address of interface. | 
| **unmanaged-mode** (*yes \| no* ; Default:**no** ) | Set unmanaged mode active, the configuration for additional settings will be possible, such as: **peer-cookie, send-cookie, local-tunnel-id, local-session-id, remote-tunnel-id, remote-session-id, local-address.** | 
| **local-tunnel-id ( string; Default: disabled)** | Set value for local-tunnel-id, an integer required. | 
| **local-session-id ( string; Default: disabled)** | Set value for local-session-id, an integer required. | 
| **remote-tunnel-id (** string; Default:**disabled)** | Set value for remote-tunnel-id, an integer required. | 
| **remote-session-id (** string; Default: **disabled)** | Set value for remote-session-id, an integer required. | 
| **peer-cookie (** string; Default:**disabled)** | Sets optional peer cookie. To enable cookie enter remote cookie value (8 or 16 character hex string value expected) to disable leave empty. | 
| **send-cookie (** string; Default: **disabled)** | Sets optional cookie. To enable cookie enter remote cookie value (8 or 16 character hex string value expected) to disable leave empty. | 
| **mtu** (*auto* ; Default:**1420** ) | Maximum Transmission Unit. Max packet size that L2TP interface will be able to send without packet fragmentation. | 
| **name** (*string* ; Default: ) | Descriptive name of the interface. | 
| **local-address** (*IP address* ; Default: ) | Set local address for **unmanaged** mode. | 
| **use-ipsec** (*yes \| no* ; Default:**no** ) | When this option is enabled, dynamic IPSec peer configuration and policy is added to encapsulate L2TP connection into IPSec tunnel. | 
| **allow-fast-path** (*yes \| no* ; Default:**no** ) | Allow to forward packets without additional processing in the Linux kernel. | 
| **l2tp-proto-version** (*l2tpv3-ip \| l2tpv3-udp \|* ; Default:***l2tpv3-udp***  ) | Specify protocol version. | 
| **cookie-length** ( 0*\| 4-bytes \| 8-bytes* ; Default:**0** ) | Configures an L2TPv3 pseudowire static session cookie. | 
| **digest-hash** (*md5 \| none \| sha1* ; Default:**md5** ) | Specifies which hash function to be used. | 
| **use-l2-specific-sublayer** (*yes \| no* ; Default:**no** ) | Specify source address. | 
| **circuit-id**  | Set the virtual circuit identifier to bind the one end of the L2TPv3 control channel, this works as identifier for each redundant pseudowire. | 
| **ipsec-secret** (*string* ; Default: )*sensitive* | Preshared key used when use-ipsec is enabled. |
