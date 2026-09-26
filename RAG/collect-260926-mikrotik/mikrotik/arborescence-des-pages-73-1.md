---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-73-1
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-73.md
source_anchor: ""
source_lines: [1, 79]
sha256: 4af066dffb9cb39a4a47bfd0ba8cf3279a2f4f831d5786c2c78be92e1423ee82
---

# Introduction

The OpenVPN security model is based on SSL, the industry standard for secure communications via the internet. OpenVPN implements OSI layer 2 or 3 secure network extensions using the SSL/TLS protocol. Support IPv4, IPv6.

# Introduction

OpenVPN has been ported to various platforms, including Linux and Windows, and its configuration is likewise on each of these systems, so it makes it easier to support and maintain. OpenVPN can run over User Datagram Protocol (UDP) or Transmission Control Protocol (TCP) transports, multiplexing created SSL tunnels on a single TCP/UDP port. OpenVPN is one of the few VPN protocols that can make use of a proxy, which might be handy sometimes.

# Limitations

ROS has its own ovpn implementation , not all ovpn features are supported and not all unsupported are listed. Currently, noteable unsupported OpenVPN features:

- LZO compression. **DEPRECATED** Compression is generally not recommended. VPN tunnels which use compression are susceptible to the VORALCE attack vector.
- NCP autonegotiation, cipher has to been specified in .ovpn file when connecting to an ROS ovpn server.

OpenVPN username is limited to 27 characters and the password to 233 characters. Password cap increased in 7.18_ab253 to 1000 characters.

# OVPN Client

| Property | Description | 
|---|---|
| **add-default-route** (*yes* \| *no* ; Default:**no** ) | Whether to add OVPN remote address as a default route. | 
| **auth** (*md5* \| *sha1* \| *null* \| *sha256* \| *sha512* ; Default:**sha1** ) | Allowed authentication methods. | 
| **certificate** (*string* \| *none* ; Default:**none** ) | Name of the client certificate | 
| **cipher** (*null* \|*aes128-cbc* \|*aes128-gcm* \|*aes192-cbc* \|*aes192-gcm* \|*aes256-cbc* \|*aes256-gcm* \|*blowfish128* ; Default:**blowfish128** ) | Allowed ciphers. In order to use GCM type ciphers, the "auth" parameter must be set to "null", because GCM cipher is also responsible for "auth", if used. | 
| **comment** (*string* ; Default: ) | Descriptive name of an item | 
| **connect-to** (*IP\|IPv6* ; Default: ) | Remote address of the OVPN server. | 
| **disabled** (*yes* \| *no* ; Default:**yes** ) | Whether the interface is disabled or not. By default it is disabled. | 
| **mac-address** (*MAC* ; Default: ) | Mac address of OVPN interface. Will be automatically generated if not specified. | 
| **max-mtu** (*integer* ; Default:**1500** ) | Maximum Transmission Unit. Max packet size that the OVPN interface will be able to send without packet fragmentation. | 
| **mode** (*ip* \| *ethernet* ; Default:**ip** ) | Layer3 or layer2 tunnel mode (alternatively tun, tap) | 
| **name** (*string* ; Default: ) | Descriptive name of the interface. | 
| **password** (*string* ; Default:**""** )*sensitive* | Password used for authentication. Value of password should not be longer than 1000 chars. | 
| **port** (*integer* ; Default:**1194** ) | Port to connect to. | 
| **profile** (*name* ; Default:**default** ) | Specifies which PPP profile configuration will be used when establishing the tunnel. | 
| **protocol** (*tcp*  \| *udp* ; Default:**tcp** ) | indicates the protocol to use when connecting with the remote endpoint. | 
| **verify-server-certificate**  (*yes* \| *no* ; Default: **no** ) | Checks the certificates CN or SAN against the "connect-to" parameter. The IP or hostname must be present in the server's certificate. | 
| **tls-version** (*any*  \| *only-1.2* ; Default:**any** ) | Specifies which TLS versions to allow | 
| **use-peer-dns**  (*yes* \| *no* ; Default:**no** ) | Whether to add DNS servers provided by the OVPN server to IP/DNS configuration. | 
| **route-nopull** (*yes* \| *no* ; Default:**no** ) | Specifies whether to allow the OVPN server to add routes to the OVPN client instance routing table. | 
| **user** (*string* ; Default: ) | User name used for authentication. | 

Also, it is possible to import the OVPN client configuration from a .ovpn configuration file. Such a file usually is provided from the OVPN server side and already includes configuration so you need to worry only about a few parameters.

OVPN client supports tls authentication. The configuration of tls-auth can be added only by importing .ovpn configuration file. Using tls-auth requires that you generate a shared-secret key, this key should be added to the client configuration file .ovpn.

Note* ROS client requires user name and password. Authentication is managed by server side, if its supports tls, then user name will be ignored.

# Tls-crypt, tls-crypt v2

To improve TLS auth, Tls-crypt is added in version 7.17rc3.

Tls-crypt, tls-crypt v2 is suppoorted only for ovpn client with following settings:

“auth SHA256” and no key-direction in server configuration,

“auth SHA256” and “key-direction 1” in client configuration is needed for authentication to work.

Example configuration files:

# OVPN Server

An interface is created for each tunnel established to the given server. There are two types of interfaces in the OVPN server's configuration

- Static interfaces are added administratively if there is a need to reference the particular interface name (in firewall rules or elsewhere) created for the particular user.
- Dynamic interfaces are added to this list automatically whenever a user is connected and its username does not match any existing static entry (or in case the entry is active already, as there can not be two separate tunnel interfaces referenced by the same name).

Dynamic interfaces appear when a user connects and disappear once the user disconnects, so it is impossible to reference the tunnel created for that use in router configuration (for example, in the firewall), so if you need a persistent rule for that user, create a static entry for him/her. Otherwise, it is safe to use dynamic configuration.

After upgrade to 7.17 version ovpn server will receive its configuration, due to multiple server support.

An disabled ovpn server with added mac will appear in configuration: 

/interface ovpn-server server add mac-address=99:99:99:99:99:99 name=ovpn-server1

In both cases PPP users must be configured properly - static entries do not replace PPP configuration.

## Server Configuration

### Properties

