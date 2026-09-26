---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-16
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-16.md
source_anchor: ""
source_lines: [1, 123]
sha256: d9c3e0dfc5e30f97226e747a05ded476ea8b62a7953ddb0cd159564732ab82a5
---

# Summary

IP/Services lists the protocols and ports used by various MikroTik RouterOS services and containers, including those for incoming connections.

It helps to determine which MikroTik services (or containers) are listening on specific ports, and what needs to be blocked or allowed if you want to restrict or permit access to certain services.

The default services that can be configured from IP/Services section:

| Property | Description | 
|---|---|
| **telnet** | Telnet service | 
| **ftp** | FTP service | 
| **www** | WebFig HTTP service | 
| **ssh** | SSH service | 
| **www-ssl** | WebFig HTTPS service | 
| **api** | API service | 
| **winbox** | Responsible for WinBox tool access, as well as MikroTik smartphone app and Dude | 
| **api-ssl** | API over SSL service | 
| **reverse-proxy** | Reverse Proxy service | 

# Properties

Note that it is not possible to add new services, only existing service modifications are allowed.

**Sub-menu:** `/ip service`

| Property | Description | 
|---|---|
| **address** (*IP address/netmask \| IPv6/0..128* ; Default: ) | List of IP/IPv6 prefixes from which the service is accessible. When this parameter is set, packets are not dropped at the network level, but access to the service is denied for sources not matching the specified addresses. This option is best suited for restricting access within trusted networks. To block access from external or untrusted networks, we recommend using a Firewall instead. | 
| **certificate** (*name* ; Default:**none** ) | The name of the certificate used by a particular service. Applicable only for services that depend on certificates ( *www-ssl, api-ssl* ) | 
| **name** (*name* ; Default:**none** ) | Service name | 
| **max-sessions**   (*integer: 1..1000* ; Default: 20) | Max simultaneous session count for service | 
| **port** (*integer: 1..65535* ; Default: ) | The port particular service listens on | 
| ***tls-version*** (*any* \|*only-1.2* ; Default:**any** ) | Specifies which TLS versions to allow by a particular service | 
| **vrf** (*name* ; Default:**main** ) | Specify which VRF instance to use by a particular service | 

## Read-only properties

| Property | Description | 
|---|---|
| **Container** | Name of the container listening on the port | 
| **Local** | Router local address used for the connection | 
| **Remote** | Remote address that established the connection to the service | 

## Example

For example, allow API only from a specific IP/IPv6 address range

Example that shows dynamic services that listens or has establish connections to router services

# Protocols and ports

The table below shows the list of protocols and ports used by RouterOS.

| Proto/Port | Description | 
|---|---|
| **20/tcp** | FTP data connection | 
| **21/tcp** | FTP control connection | 
| **22/tcp** | Secure Shell (SSH) remote login protocol | 
| **23/tcp** | Telnet protocol | 
| **53/tcp 53/udp** | DNS | 
| **67/udp** | Bootstrap protocol or DHCP Server | 
| **68/udp** | Bootstrap protocol or DHCP Client | 
| **80/tcp** | World Wide Web HTTP | 
| **123/udp** | Network Time Protocol (NTP) | 
| **161/udp** | Simple Network Management Protocol (SNMP) | 
| **179/tcp** | Border Gateway Protocol (BGP) | 
| **443/tcp** | Secure Socket Layer (SSL) encrypted HTTP | 
| **500/udp** | Internet Key Exchange (IKE) protocol | 
| **520/udp 521/udp** | RIP routing protocol | 
| **546/udp** | DHCPv6 Client message | 
| **547/udp** | DHCPv6 Server message | 
| **646/tcp** | LDP transport session | 
| **646/udp** | LDP hello protocol | 
| **1080/tcp** | SOCKS proxy protocol | 
| **1698/udp 1699/udp** | RSVP TE Tunnels | 
| **1701/udp** | Layer 2 Tunnel Protocol (L2TP) | 
| **1723/tcp** | Point-To-Point Tunneling Protocol (PPTP) | 
| **1900/udp 2828/tcp** | Universal Plug and Play (uPnP) | 
| **1966/udp** | MME originator message traffic | 
| **1966/tcp** | MME gateway protocol | 
| **2000/tcp** | Bandwidth test server | 
| **5246,5247/udp** | CAPsMAN | 
| **5350/udp** | NAT-PMP client | 
| **5351/udp** | NAT-PMP server | 
| **5678/udp** | Mikrotik Neighbor Discovery Protocol | 
| **6343/tcp** | Default OpenFlow port | 
| **8080/tcp** | HTTP Web Proxy | 
| **8291/tcp** | Winbox | 
| **8728/tcp** | API | 
| **8729/tcp** | API-SSL | 
| **20561/udp** | MAC winbox | 
| **/1** | ICMP | 
| **/2** | Multicast \| IGMP | 
| **/4** | IPIP encapsulation | 
| **/41** | IPv6 (encapsulation) | 
| **/46** | RSVP TE tunnels | 
| **/47** | General Routing Encapsulation (GRE) - used for PPTP and EoIP tunnels | 
| **/50** | Encapsulating Security Payload for IPv4 (ESP) | 
| **/51** | Authentication Header for IPv4 (AH) | 
| **/89** | OSPF routing protocol | 
| **/103** | Multicast \| PIM | 
| **/112** | VRRP | 

# Web server

The table below shows the list of properties that can be enabled/disabled for web services. All of properties are enabled by default and can be disabled if desired.

In this table "plain" refers to HTTP connections and "secure" to HTTPS connections.

| Property | Description | 
|---|---|
|  index-plain: (Default: **yes** ) | Home page/login page (Can be disabled when webfig-plain and graphs-plain are disabled) | 
| webfig-plain: (Default: **yes** ) | WebFig interface | 
| graphs-plain: (Default: **yes** ) | Graph page | 
| rest-plain: (Default: **yes** ) | REST API support | 
| crl-plain: (Default: **yes** ) | CRL(Certificate Revocation List) | 
| scep-plain: (Default: **yes** ) | SCEP(Simple Certificate Enrollment Protocol) | 
| acme-plain: (Default: **yes** ) | ACME Challenge | 
| index-secure: (Default: **yes** ) | Home page/login page (Can be disabled when webfig-secure and graphs-secure are disabled) | 
| webfig-secure: (Default: **yes** ) | WebFig interface | 
| graphs-secure: (Default: **yes** ) | Graph page | 
| rest-secure: (Default: **yes** ) | REST API support |
