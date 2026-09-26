---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-71-1
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-71.md
source_anchor: ""
source_lines: [1, 95]
sha256: 96796c587ae9109937ada887ee0a483000f4e4686c5ac89d1439bd7e3a20d5a8
---

# Overview

Secure Socket Tunneling Protocol (SSTP) transports a PPP tunnel over a TLS channel. The use of TLS over TCP port 443 allows SSTP to pass through virtually all firewalls and proxy servers.

# Introduction

Let's take a look at the SSTP connection mechanism:

1. A TCP connection is established from client to server (by default on port 443);
2. SSL validates the server certificate. If a certificate is valid, a connection is established otherwise the connection is turned down. (But see note below);
3. The client sends SSTP control packets within the HTTPS session which establishes the SSTP state machine on both sides;
4. PPP negotiation over SSTP. The client authenticates to the server and binds IP addresses to the SSTP interface;

SSTP tunnel is now established and packet encapsulation can begin;

Starting from v5.0beta2 SSTP does not require certificates to operate and can use any available authentication type. This feature will work only between two MikroTik routers, as it is not in accordance with Microsoft standards. Otherwise to establish secure tunnels **mschap** authentication and client/server certificates from the same chain should be used.

TLS SNI support has been added starting with 7.15beta10 version, Extension will be added to client hello packets if "Add SNI" checkbox is checked or set in CLI:

interface/sstp-client/set add-sni=yes

# SSTP Client

## Properties

| **authentication** (*chap, mschap1, mschap2, pap* ; Default:**"all"** ) | Allowed authentication methods, by default all methods are allowed. | 
| **disabled** (*yes \| no* ; Default:**yes** ) | Enables/disables tunnel. | 
| **add-default-route** (*yes \| no* ; Default:**no** ) | Whether to add L2TP remote address as a default route. | 
| **default-route-distance** (*byte* ; Default: ) | Sets distance value applied to auto created default route, if add-default-route is also selected. | 
| **mrru** (*integer: 512..65535\|disabled* ; Default:**disabled** ) | maximum packet size that can be received on the link. If a packet is bigger than tunnel MTU, it will be split into multiple packets, allowing full size IP or Ethernet packets to be sent over the tunnel. | 
| **proxy-port** (*string* ; Default:**443** ) | Sets proxy port. | 
| **add-sni** (*yes \| no* ; Default:**no** ) | Enables/disables service. | 
| **dial-on-demand** (*yes \| no* ; Default:**no** ) | Connects only when outbound traffic is generated. If selected, then route with gateway address from 10.112.112.0/24 network will be added while connection is not established. | 
| **name** (*string* ; Default: ) | Descriptive name of the interface. | 
| **tls-version**  (*any*  \| *only-1.2* ; Default:**any** ) | Specifies which TLS version to allow. | 
| **numbers** (*integer;* ) | Sets number for an tunnel in ROS. | 
| **user** (*string* ; Default: ) | User name used for authentication. | 
| **certificate** (*string*  \| *none* ; Default:**none** ) | Name of the client certificate | 
| **http-proxy** (*string* ; Default: ) | Proxy address field. | 
| **password** (*string* ; Default:**""** )*sensitive* | Password used for authentication. | 
| **verify-server-address-from-certificate** (*yes\|no* ; Default:**no** ) | SSTP client will verify server address in certificate. | 
| **verify-server-certificate** (*yes\|no* ; Default:**no** ) | SSTP client will verify server certificate. | 
| **ciphers** (aes256-gcm-sha384 \| aes256-sha; Default:**all** ) | Allowed ciphers. | 
| **keepalive-timeout** (*integer* ; Default:**60** ) | Sets keepalive timeout in seconds. | 
| **pfs** (*yes \| no \| required* ; Default:**no** ) | Specifies which TLS authentication to use. With pfs=yes, TLS will use ECDHE-RSA- and DHE-RSA-. For maximum security setting pfs=required will use only ECDHE. | 
| **comment** (*string* ; Default: ) | Short description of the tunnel. | 
| **max-mru** (*integer* ; Default:**1460** ) | Maximum Receive Unit. | 
| **max-mtu** (*integer* ; Default:**1460** ) | Maximum Transmission Unit. | 
| **port** (*integer* ; Default:**443** ) | Port to connect to. | 
| **connect-to** (*IP\|IPv6* ; Default: ) | Remote address of the SSTP server. | 
| **profile** (*name* ; Default:**default** ) | Specifies which PPP profile configuration will be used when establishing the tunnel. | 

# SSTP Server

## Properties

| **authentication** (*chap, mschap1, mschap2, pap* ; Default:**"all"** ) | Allowed authentication methods, by default all methods are allowed. | 
| **keepalive-timeout** (*integer* ; Default:**60** ) | Sets keepalive timeout in seconds. | 
| **port** (*string* ; Default:**443** ) | Sets port used. | 
| **certificate** (*string*  \| *none* ; Default:**none** ) | Name of the certificate in use. | 
| **max-mru** (*integer* ; Default:**1460** ) | Maximum Receive Unit. | 
| **max-mtu** (*integer* ; Default:**1460** ) | Maximum Transmission Unit. | 
| **tls-version**  (*any*  \| *only-1.2* ; Default:**any** ) | Specifies which TLS version to allow. | 
| **ciphers** (aes256-gcm-sha384 \| aes256-sha; Default:**all** ) | Allowed ciphers. | 
| **verify-client-certificate** (*yes\|no* ; Default:**no** ) | SSTP server will verify client certificate. | 
| **mrru** (*integer: 512..65535\|disabled* ; Default:**disabled** ) | maximum packet size that can be received on the link. If a packet is bigger than tunnel MTU, it will be split into multiple packets, allowing full size IP or Ethernet packets to be sent over the tunnel. | 
| **default-profile** (*name* ; Default:**default** ) | Specifies which PPP profile configuration will be used when establishing the tunnel. | 
| **enabled** (*yes \| no* ; Default:**no** ) | Enables/disables service. | 
| **pfs** (*yes \| no \| required* ; Default:**no** ) | Specifies which TLS authentication to use. With pfs=yes, TLS will use ECDHE-RSA- and DHE-RSA-. For maximum security setting pfs=required will use only ECDHE. | 

# Certificates

To set up a secure SSTP tunnel, certificates are required. On the server, authentication is done only by *username* and *password,* but on the client - the server is authenticated using a server certificate. It is also used by the client to cryptographically bind SSL and PPP authentication, meaning - the clients send a special value over SSTP connection to the server, this value is derived from the key data that is generated during PPP authentication and server certificate, this allows the server to check if both channels are secure.

If SSTP clients are on Windows PCs then the only way to set up a secure SSTP tunnel when using a self-signed certificate is by importing the "server" certificate on the SSTP server and on the Windows PC adding a CA certificate in the trusted root.

If your server certificate is issued by a CA which is already known by Windows, then the Windows client will work without any additional certificate imports to a trusted root.

RSA key length must be at least 472 bits if a certificate is used by SSTP. Shorter keys are considered as security threats.

A similar configuration on RouterOS client would be to import the CA certificate and enabling the verify-server-certificate option. In this scenario, Man-in-the-Middle attacks are not possible.

Between two Mikrotik routers, it is also possible to set up an insecure tunnel by not using certificates at all. In this case, data going through the SSTP tunnel is using anonymous DH and Man-in-the-Middle attacks are easily accomplished. This scenario is not compatible with Windows clients.

It is also possible to make a secure SSTP tunnel by adding additional authorization with a client certificate. Configuration requirements are:

- certificates on both server and client
- verification options enabled on server and client

This scenario is also not possible with Windows clients, because there is no way to set up a client certificate on Windows.

### Certificate Error Messages

When SSL handshake fails, you will see one of the following certificate errors:

