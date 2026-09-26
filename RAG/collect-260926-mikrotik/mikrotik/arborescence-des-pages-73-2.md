---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-73-2
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-73.md
source_anchor: ""
source_lines: [80, 142]
sha256: a181d11f2c2de021cb7ee1d29aed4458768cdaee5f769df7ccae4ae62a8609b8
---

# Introduction

| Property | Description | 
|---|---|
| **auth** (*md5* \| *sha1* \| *null* \| *sha256* \| *sha512* ; Default:**sha1,md5,sha256,sha512** ) | Authentication methods that the server will accept. | 
| **certificate** (*name* \| *none* ; Default:**none** ) | Name of the certificate that the OVPN server will use. | 
| **cipher** (*null* \|*aes128-cbc* \|*aes128-gcm* \|*aes192-cbc* \|*aes192-gcm* \|*aes256-cbc* \|*aes256-gcm* \|*blowfish128* ; Default:**aes128-cbc,blowfish128** ) | Allowed ciphers. | 
| **default-profile** (*name* ; Default:**default** ) | Default profile to use. | 
| **disabled** (*yes* \| *no* ; Default:**yes** ) | Defines whether the OVPN server is enabled or not. | 
| **protocol (*tcp* \| *udp*; Default: tcp)** | Indicates the protocol to use when connecting with the remote endpoint. | 
| **keepalive-timeout** (*integer* \| *disabled* ; Default:**60** ) | Defines the time period (in seconds) after which the router is starting to send keepalive packets every second. If no traffic and no keepalive responses have come for that period of time (i.e. 2 * keepalive-timeout), not responding client is proclaimed disconnected | 
| **mac-address** (*MAC* ; Default: ) | Automatically generated MAC address of the server. | 
| **max-mtu** (*integer* ; Default:**1500** ) | Maximum Transmission Unit. Max packet size that the OVPN interface will be able to send without packet fragmentation. | 
| **mode** (*ip* \| *ethernet* ; Default:**ip** ) | Layer3 or layer2 tunnel mode (alternatively tun, tap) | 
| **name** *(string)*  | Name of the server | 
| **netmask** (*integer* ; Default:**24** ) | Subnet mask to be applied to the client. | 
| **port** (*integer* ; Default:**1194** ) | Port to run the server on. | 
| **require-client-certificate** (*yes* \| *no* ; Default:**no** ) | If set to yes, then the server checks whether the client's certificate belongs to the same certificate chain. | 
| **redirect-gateway** (*def1* \| *disabled* \| *ipv6;* Default:**disabled** ) | Specifies what kind of routes the OVPN client must add to the routing table. `def1` – Use this flag to override the default gateway by using 0.0.0.0/1 and 128.0.0.0/1 rather than 0.0.0.0/0. This has the benefit of overriding but not wiping out the original default gateway.`disabled` - Do not send redirect-gateway flags to the OVPN client.`ipv6` - Redirect IPv6 routing into the tunnel on the client side. This works similarly to the def1 flag, that is, more specific IPv6 routes are added (2000::/4 and 3000::/4), covering the whole IPv6 unicast space. | 
| **enable-tun-ipv6** (y*es* \| *no;* Default:**no** ) | Specifies if IPv6 IP tunneling mode should be possible with this OVPN server. | 
| **ipv6-prefix-len** (*integer;* Default:**64** ) | Length of IPv6 prefix for IPv6 address which will be used when generating OVPN interface on the server side. | 
| **reneg-sec**  (*integer;* Default:**3600)** | Key renegotiate seconds, the time the server periodically renegotiates the secret key for the data channel. | 
| **push-routes** (*string* ; Default: ) | Push route support are added in 7.14, the maximum of possible input is limited to 1400 characters or 37 pushed routes. IPv6 support added in 7.21_ab220. | 
| **tls-version** (any  \| *only-1.2 ;* Default:**any** ) | TLS protocol setting. | 
| **tun-server-ipv6** (*IPv6 prefix;* Default:**::** ) | IPv6 prefix address which will be used when generating the OVPN interface on the server side. | 
| **user-auth-method** (*mschap2 \| pap ; Default **pap*** ) | By the default pap authentication method is used, if preferred server authentication with chap challenge set mschap2 in server settings. | 
| **vrf** ()  | VRF in which listen for connection attempts | 

Also, it is possible to prepare a .ovpn file for the OVPN client which can be easily imported on the end device. **Server need to have option enabled - required client certificate to export work.** 

It is very important that the date on the router is within the range of the installed certificate's date of expiration. To overcome any certificate verification problems, enable **NTP** date synchronization on both the server and the client.

# Example

## Setup Overview

Assume that Office public IP address is 2.2.2.2 and we want two remote OVPN clients to have access to 10.5.8.20 and 192.168.55.0/24 networks behind the office gateway.

## Creating Certificates

All certificates can be created on the RouterOS server using the certificate manager. See example >>.

For the simplest setup, you need only an OVPN server certificate.

## Server Config

The first step is to create an IP pool from which client addresses will be assigned and some users.

Assume that the server certificate is already created and named "server"

## Client Config

Add manually which networks you want to access over the tunnel.

## Push Route

Push route support are added in 7.14, the maximum of possible input is limited to **1400** characters or 37 routes. IPv6 support added in 7.21_ab220.

example: route network/IP [netmask] [gateway] [metric].

## VRF support

Support starting from **7.17 version** is added, and couple changes introduced in configuration, if you use latest version, please refer to this example:

Server side configuration:
