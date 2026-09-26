---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-36-1
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-36.md
source_anchor: ""
source_lines: [1, 63]
sha256: 4335d5632e8a70ff3205b8b930622f2d85c3270f7ba6350906b1d1ccd3f18762
---

# Introduction

The MikroTik HotSpot Gateway provides authentication for clients before access to public networks.

**Hotspot (captive portal)** - uses web-proxy and it is capable of using only the default routing table, at the moment. Making the PCC(per connection-classifier) not a valid method, due to the, multiple routing tables used.

Important

## **HotSpot Gateway features:**

- different authentication methods of clients, using a local client database on the router, or remote RADIUS server;
- users accounting in a local database on the router, or on remote RADIUS server;
- a walled-garden system, access to some web pages without authorization;
- login page modification, where you can put information about the company;
- automatic and transparent change any IP address of a client to a valid address;
- HotSpot can inform DHCP clients that they are behind a captive portal (RFC7710);

A hotspot can work reliably only when IPv4 is used. Hotspot relies on Firewall NAT rules which currently are not supported for IPv6.

# Example

Verify HotSpot configuration:

## **Parameters asked during the setup process**

| Parameter | Description | 
|---|---|
| **hotspot interface** (*string* ; Default:**allow** ) | Interface name on which to run HotSpot. To run HotSpot on a bridge interface, make sure public interfaces are not included in the bridge ports. | 
| **local address of network** (*IP* ; Default:**10.5.50.1/24** ) | HotSpot gateway address | 
| **masquerade network** (*yes \| no* ; Default:**yes** ) | Whether to masquerade HotSpot network, when **yes** rule is added to*/ip firewall nat* with*action=masquerade* | 
| **address pool of network** (*string* ; Default:**yes** ) | Address pool for HotSpot network, which is used to change user IP address to a valid address. Useful if providing network access to mobile clients that are not willing to change their networking settings. | 
| **select certificate** (*none \| import-other-certificate* ; Default: ) | Choose SSL certificate, when HTTPS authorization method is required. | 
| **ip address of smtp server** (*IP* ; Default:**0.0.0.0** ) | The IP address of the SMTP server, where to redirect HotSpot's network SMTP requests (25 TCP port) | 
| **dns servers** (*IP* ; Default:**0.0.0.0** ) | DNS server addresses used for HotSpot clients, configuration taken from */ip dns* menu of the HotSpot gateway | 
| **dns name** (*string* ; Default:**""** ) | the domain name of the HotSpot server, a full qualified domain name is required, for example, www.example.com | 
| **name of local hotspot user** (*string* ; Default:**"admin"** ) | username of one automatically created HotSpot user, added to */ip hotspot user* | 
| **password for the user'** (*string* ; Default: ) | Password for automatically created HotSpot user | 

# HotSpot

The menu is designed to manage the HotSpot servers of the router. It is possible to run HotSpot on Ethernet, wireless, VLAN, and bridge interfaces. One HotSpot server is allowed per interface. When HotSpot is configured on the bridge interface, set HotSpot interface as bridge interface, not as bridge port, do not add public interfaces to bridge ports. You can add HotSpot servers manually to the */ip/hotspot* menu, but it is advised to run */ip/hotspot/setup*, which adds all necessary settings.

| Parameters | Description | 
|---|---|
| **name** (text) | HotSpot server's name or identifier | 
| **address-pool** (name/none; default:*none* ) | address space used to change HotSpot client *any* IP address to a valid address. Useful for providing public network access to mobile clients that are not willing to change their networking settings | 
| **idle-timeout** (time/none; default:*5m* ) | period of inactivity for unauthorized clients. When there is no traffic from this client (literally client computer should be switched off), once the timeout is reached, a user is dropped from the HotSpot host list, its used address becomes available | 
| **keepalive-timeout** (time/none; default:*none* ) | Value of how long host can stay out of reach to be removed from the HotSpot | 
| **login-timeout** (time/none; default:*none* ) | Period of time after which if a host hasn't been authorized itself with a system the host entry gets deleted from host table. Loop repeats until the host logs in the system. Enable if there are situations where a host cannot log in after being too long in the host table unauthorized. | 
| **interface** (name of an interface) | Interface to run HotSpot on | 
| **addresses-per-mac** (integer**/** unlimited; default: 2) | Number of IP addresses allowed to be bind with the MAC address, when multiple HotSpot clients connected with one MAC-address | 
| **profile** (name; default:***default*)** | HotSpot server default HotSpot profile, which is located in */ip/hotspot/profile* | 

Read-only

| Parameters | Description | 
|---|---|
| keepalive-timeout (read-only; time) | The exact value of the keepalive-timeout, that is applied to the user. Value shows how long the host can stay out of reach to be removed from the HotSpot | 

# HotSpot Profile

This submenu contains list of Hotspot server profiles. There may be various different HotSpot systems, defined as Server Profiles, on the same gateway machine. One or more interfaces can be grouped into one server profile. There are very few settings for the servers on particular interfaces - most of the configuration is set in the server profiles. For example, it is possible to make completely different set of servlet pages for each server profile, and define different RADIUS servers for authentication.

