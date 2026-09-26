---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-50-3
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-50.md
source_anchor: ""
source_lines: [160, 221]
sha256: 5add2cc138c1f3cb4cdeab3b95ea8bc4b409ca8a2e54df2d7ccb2d9e1fdcf8c4
---

# Introduction

To see how many domain names are present and matched, you can run:

### Locally hosted adlist:

To create your adlist, you can create a Txt file with the domains. Example:

You can create the txt file on your PC, but it is also possible to create it in RouterOS, with following commands

"/file/add name=host.txt", and then you can run "file/edit host.txt contents" after adding entries, press "ctrl o" to save the entries.

To add file to adlist :

You can verify that file is formatted correctly with "/ip/dns/adlist/print" ,the results will show how many hostnames you have added, the hostname format must match the format given in previous example.

# Forwarders

DNS Forwarders allows a user to configure a named DNS forwarder that can be used for static FWD entries as *forward-to* value.

For each *Forwarder* is possible to configure multiple regular upstream and DoH servers. Configured *forwarder* servers will be used by round-robin algorithm - for each query, next server will be used to resolve DNS name.

## Forwarder configuration

In */ip/dns/forwaders* section, *forwarders* can added, modified or removed. 

| Property | Description | 
|---|---|
| **name** (*string;* Default: ) | Forwarder name. | 
| **dns-servers** (*string* ; Default: ) | An IP address or DNS name of a domain name server. Can contain multiple records, for example, *dns-servers=1.1.1.1,8.8.8.8,local.dns* | 
| **doh-servers** (*string;* Default: ) | A URL of DoH server. Can contain multiple records. | 
| **verify-doh-cert** (*yes* \|*no**;* Default:*yes* ) | Specifies whether to validate the DoH server, when one is being used. Will use the "/certificate" list in order to verify server validity. | 

## Configuration example

Configure/add a *forwarder*: 

Configure/add a statis DNS FWD entry:

Now each time when a router will receive request to resolve mikrotik.com, request using round-robin algorithm will be forwarded to *1.1.1.1*, *local.dns* or *Google DoH* server.

# mDNS

RouterOS supports Multicast DNS (mDNS) for local network service discovery. By default, mDNS operates within a single subnet. The mDNS repeater feature allows to extend mDNS functionality across different interfaces or VLANs using the "mdns-repeat-ifaces" property.

Impact of using mDNS Repeater:

- Cross-Subnet Service Discovery: Devices on different subnets or VLANs can discover each other, enhancing the ability to find services (e.g., printers, file sharing);
- Increased Network Traffic: mDNS repeater may increase multicast traffic, which could lead to congestion, especially in larger networks with many devices.

The mDNS repeater is commonly used with devices such as:

- Apple Ecosystem (AirPrint, AirPlay);
- Smart Home Devices (Thread, IoT);
- Chromecast and Media Streaming;
- Avahi (Linux/Unix).

To enable the mDNS repeater between interfaces, allowing devices connected to these interfaces to discover each other using mDNS, use the following command:

mDNS repeater requires multicast-capable interfaces (e.g., Ethernet, VLAN, bridge). Tunnel interfaces such as WireGuard are not supported.

Currently only IPv4 is supported.

MikroTik mDNS Repeater is a local service intercepting multicast packets to rebroadcast them, it requires an "input" rule. mDNS multicast traffic does not traverse the "forward" chain. In case you have strict firewall rules protecting your router from local subnets, you must explicitly allow mDNS traffic before any drop - allowed to receive UDP port 5353 traffic on the "input" chain.
