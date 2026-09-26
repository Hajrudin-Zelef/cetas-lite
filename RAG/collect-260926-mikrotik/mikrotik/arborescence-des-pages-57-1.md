---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-57-1
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-57.md
source_anchor: ""
source_lines: [1, 76]
sha256: 67e25521256b25c8c9c1110d144af66cc2e7f8c775d1bca75dbdfa38c0cd3b14
---

# Introduction

WireGuard<sup>®</sup> is an extremely simple yet fast and modern VPN that utilizes state-of-the-art cryptography. It aims to be faster, simpler, leaner, and more useful than IPsec while avoiding massive headaches. It intends to be considerably more performant than OpenVPN. WireGuard is designed as a general-purpose VPN for running on embedded interfaces and super computers alike, fit for many different circumstances. Initially released for the Linux kernel, it is now cross-platform (Windows, macOS, BSD, iOS, Android) and widely deployable.

# Properties

| Property | Description | 
|---|---|
| **comment** (*string* ; Default: ) | Short description of the tunnel. | 
| **disabled** (*yes \| no* ; Default:**no** ) | Enables/disables the tunnel. | 
| **listen-port** (*integer; Default: 13231* ) | Port for WireGuard service to listen on for incoming sessions. | 
| **mtu** (*integer [0..65536]* ; Default:**1420** ) | Layer3 Maximum transmission unit. | 
| **name** (*string* ; Default: ) | Name of the tunnel. | 
| **vrf** (string; Default:**main** ) | Specify vrf for wireguard socket, more details bellow | 
| **private-key** (*string* ; Default: )*sensitive* | A base64 private key. If not specified, it will be automatically generated upon interface creation. Each network interface has a private key and a list of peers. | 

## `The` **vrf** parameter in WireGuard context

**vrf**

The `vrf` parameter does **not** apply to the WireGuard interface itself (e.g., `wg0`, `wg1`), but rather to the **UDP socket** used for transporting encrypted packets.

There are two distinct layers in WireGuard operation:

1. **WireGuard interface** — this is the virtual network device through which*plain* (unencrypted) packets flow. These packets are encrypted and sent out through the UDP socket, or decrypted when they come in from it.
2. **UDP sockets** — these handle the*encrypted* traffic: receiving encrypted packets from the network and sending encrypted packets out.

The `vrf` parameter is relevant to **the UDP socket layer** (case 2). It specifies **which routing table (VRF)** the socket should use to determine how encrypted packets are sent or received.

This ensures that encrypted traffic follows the correct routing path and uses the proper source IP, preventing issues where packets might otherwise go out via the wrong interface or route.**Example:**

Suppose interface `eth1` belongs to VRF `foo`.

If you want WireGuard to send and receive encrypted packets via `eth1`, you should configure WireGuard with `vrf=foo`.

It’s perfectly normal for the WireGuard interface (`wg0`) itself to be in a different VRF than the one specified by the `vrf` parameter used by the internal UDP socket.


## Read-only properties

| Property | Description | 
|---|---|
| **public-key** (*string* ) | A base64 public key is calculated from the private key. Each peer has a public key. Public keys are used by peers to authenticate each other. They can be passed around for use in configuration files. | 
| **running** (*yes \| no* ) | Whether the interface is running. | 

# Peers

| Property | Description | 
|---|---|
| **allowed-address** (*IP/IPv6 prefix* ;*Default* : ) | List of IP (v4 or v6) addresses with CIDR masks from which incoming traffic for this peer is allowed and to which outgoing traffic for this peer is directed. This IP address has to be in the same subnet as WireGuard interface set on ROS. If WireGuard interface is at 192.168.99.1/24, You have to input 192.168.99.2 to the client. By adding this IP under 'Allowed Address', you are saying that only this specific client (phone for example) is permitted to connect to this peer configuration. Allowed-address range cannot overlap on one interface, so you need to set own range for each peer. | 
| **comment***(string; Default: )* | Short description of the peer. | 
| **disabled***(yes \| no; Default: **no**)* | Enables/disables the peer. | 
| **endpoint-address***(IP/Hostname; Default: )* | The IP address or hostname. It is used by WireGuard to establish a secure connection between two peers. | 
| **endpoint-port** (*integer:0..65535**; Default:* ) | The Endpoint port is the UDP port on which a WireGuard peer listens for incoming traffic. | 
| **interface** (*string; Default:* ) | Name of the WireGuard interface the peer belongs to. | 
| **persistent-keepalive** (*integer:0..65535; Default: 0* ) | A seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated empty packet to the peer for the purpose of keeping a stateful firewall or NAT mapping valid persistently. For example, if the interface very rarely sends traffic, but it might at anytime receive traffic from a peer, and it is behind NAT, the interface might benefit from having a persistent keepalive interval of 25 seconds. | 
| **preshared-key** (*string; Default:* )*sensitive* | A base64 preshared key. Optional, and may be omitted. This option adds an additional layer of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance. Also can be generated automatically or entered manually, when the key is provided by the system administrator. | 
| **private-key***(auto/none; Default: **none**) sensitive* | A base64 private key. | 
| **public-key** (*string; Default:* ) | A base64 public key is calculated from the private key. Each peer has a public key. Public keys are used by peers to authenticate each other. They can be passed around for use in configuration files. | 
| **show-client-config**  **sensitive* |  *Will show already created Peer configuration and generate a QR code for easier peer setup on a client device. Does not affect the WireGuard Server. To view QR code, show-sensitive is required from 7.21_ab548.* | 
| Used for the client-server setup scenario, when the configuration is imported using a qr code for a client, configuration details on tab with qrcode will appear once it has been set in the fields: |  | 
| **client-address** *(IP/IPv6 prefix; Default: )* | When imported using a qr code for a client (for example, a phone), then this address for the wg interface is set on that device. | 
| **client-dns** *(IP/IPv6 prefix; Default: )* | Specify when using WireGuard Server as a VPN gateway for peer traffic. | 
| **client-endpoint** *(IP/IPv6 prefix; Default: )* | The IP address and port number of the WireGuard Server. | 
| **client-keepalive** (*integer:0..65535; Default: 0* ) | Same as **persistent-keepalive** but from peer side. | 
| **client-listen-port** (*integer:0..65535**; Default:* ) | The local port upon which this WireGuard tunnel will listen for incoming traffic from peers, and the port from which it will source outgoing packets. | 
| **client-allowed-address *(****IP/IPv6 prefix; Default: ::/0**)* | Starting from 7.21 Allowed IPs can be configured. | 
| **name** (*string; Default:*  ) | Allows adding name to a peer. Name will be used as a reference for a peer in WireGuard logs. (Available from RouterOS version 7.15) | 
| **responder** *(yes \| no; Default:* *no**)* | Specifies if peer is intended to be connection initiator or only responder. Should be used on WireGuard devices that are used as "servers" for other devices as clients to connect to. Otherwise router will all repeatedly try to connect "endpoint-address" or "current-endpoint-address". | 

The "**AllowedIPs**" configuration provided to the client through the WireGuard peer export (configuration file or QR code) cannot be changed and will currently be set to "0.0.0.0/0, ::/0". If it is necessary to modify these values on the remote end, this must be done through the remote peer software used for the WireGuard connection. Support has been added from 7.21 versions.

# Importing, Exporting Wireguard

Configuration can be done in various ways, here is simple wg import file example: export

