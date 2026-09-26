---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-50-1
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-50.md
source_anchor: ""
source_lines: [1, 65]
sha256: ae7dcc742438810632a41d329c3e628aefc5ad18d80b7098d61070938f43cae2
---

# Introduction

Domain Name System (DNS) usually refers to the Phonebook of the Internet. In other words, DNS is a database that links strings (known as hostnames), such as www.mikrotik.com to a specific IP address, such as 159.148.147.196.

A MikroTik router with a DNS feature enabled can be set as a DNS cache for any DNS-compliant client. Moreover, the MikroTik router can be specified as a primary DNS server under its DHCP server settings. When the remote requests are enabled, the MikroTik router responds to TCP and UDP DNS requests on port 53.

When both static and dynamic servers are set, static server entries are preferred, however, it does not indicate that a static server will always be used (for example, previously query was received from a dynamic server, but static was added later, then a dynamic entry will be preferred).

When DNS server *allow-remote-requests* are used make sure that you limit access to your server over TCP and UDP protocol port 53 only for known hosts.

There are several options on how you can manage DNS functionality on your LAN - use public DNS, use the router as a cache, or do not interfere with DNS configuration. Let us take as an example the following setup: Internet service provider (ISP) → Gateway (GW) → Local area network (LAN). The GW is RouterOS based device with the default configuration:

- You do not configure any DNS servers on the "GW" DHCP server network configuration - the device will forward the DNS server IP address configuration received from `ISP` to `LAN` devices;
- You configure DNS servers on the "GW" DHCP server network configuration - the device will give configured DNS servers to `LAN` devices (also "/ip dns set allow-remote-requests=yes*" must* be enabled);
- "dns-none" configured under DNS servers on "GW" DHCP server network configuration - the device will not forward any of the **dynamic** DNS servers to `LAN` devices;

## DNS configuration

DNS facility is used to provide domain name resolution for the router itself as well as for the clients connected to it.

| Property | Description | 
|---|---|
| **allow-remote-requests** (*yes* \| *no* ; Default:**no** ) | Specifies whether to allow router usage as a DNS cache for remote clients. Otherwise, only the router itself will use DNS configuration. | 
| **address-list-extra-time** *(time; Default: **0s**)* | Extra time added to TTL when creating address list entry. | 
| **cache-max-ttl** (*time* ; Default:**1w** ) | Maximum time-to-live for cache records. In other words, cache records will expire unconditionally after cache-max-TTL time. Shorter TTLs received from DNS servers are respected. | 
| **cache-size** (*integer[64..4294967295]* ; Default:**2048** ) | Specifies the size of the DNS cache in KiB. | 
| **max-concurrent-queries** (*integer* ; Default:**100** ) | Specifies how many concurrent queries are allowed. | 
| **max-concurrent-tcp-sessions** (*integer* ; Default:**20** ) | Specifies how many concurrent TCP sessions are allowed. | 
| **max-udp-packet-size** (*integer [50..65507]* ; Default:**4096** ) | Maximum size of allowed UDP packet. | 
| **mdns-repeat-ifaces** (*list of interfaces* ; Default: ) | Once an interface in this list receives an mDNS packet, it will forward it to all other interfaces in this list. Only supports IPv4. | 
| **query-server-timeout** (*time* ; Default:**2s** ) | Specifies how long to wait for a query response from a server. | 
| **query-total-timeout** (*time* ; Default:**10s** ) | Specifies how long to wait for query response in total. Note that this setting must be configured taking into account "query-server-timeout" and the number of used DNS servers. | 
| **servers** (*list of IPv4/IPv6 addresses@vrf* ; Default: ) | List of DNS server IPv4/IPv6 addresses | 
| **cache-used** (*integer* ) | Shows the currently used cache size in KiB | 
| **dynamic-server** (*IPv4/IPv6 list* ) | List of dynamically added DNS servers from different services, for example, DHCP. | 
| **doh-max-concurrent-queries** (*integer* ; Default:**50** ) | Specifies how many DoH concurrent queries are allowed. | 
| **doh-max-server-connections** (*integer* ; Default:**5** ) | Specifies how many concurrent connections to the DoH server are allowed. | 
| **doh-timeout** (*time* ; Default:**5s** ) | Specifies how long to wait for query response from the DoH server. | 
| **use-doh-server** (*string; Default: )* | Specified which DoH server must be used for DNS queries. DoH functionality overrides " *servers* " usage if specified. The server must be specified with an "https://" prefix. Supports only one DoH server. | 
| **verify-doh-cert**  (*yes* \| *no* ; Default:**no** ) | Specifies whether to validate the DoH server, when one is being used. Will use the "/certificate" list in order to verify server validity. | 
| **vrf** (vrf; Default: main) | Specifies the VRF that should use the DNS resolver. The DNS resolver processes only requests originating from the designated VRF or from the resolver itself. | 

Dynamic DNS servers are obtained from different facilities available in RouterOS, for example, DHCP client, VPN client, IPv6 Router Advertisements, etc.

Servers are processed in a queue order - static servers as an ordered list, dynamic servers as an ordered list. When DNS cache has to send a request to the server, it tries servers one by one until one of them responds. After that this server is used for all types of DNS requests. Same server is used for any types of DNS requests, for example, A and AAAA types. If you use only dynamic servers, then the DNS returned results can change after reboot, because servers can be loaded into IP/DNS settings in a different order due to a different speeds on how they are received from facilities mentioned above.

If at some point the server which was being used becomes unavailable and can not provide DNS answers, then the DNS cache restarts the DNS server lookup process and goes through the list of specified servers once more.

## DNS Cache

This menu provides two lists with DNS records stored on the server:

- *"* /ip dns cache*"* : this menu provides a list with cache DNS entries that RouterOS cache can reply with to client requests ;
- *"* /ip dns cache all*"* : This menu provides a complete list with all cached DNS records stored including also, for example, PTR records.

You can empty the DNS cache with the command: "/ip dns cache flush"*.*

## DNS Static

The MikroTik RouterOS DNS cache has an additional embedded DNS server feature that allows you to configure multiple types of DNS entries that can be used by the DNS clients using the router as their DNS server. This feature can also be used to provide false DNS information to your network clients. For example, resolving any DNS request for a certain set of domains (or for the whole Internet) to your own page.

The server is also capable of resolving DNS requests based on basic regular expressions so that multiple requests can be matched with the same entry. In case an entry does not conform with DNS naming standards, it is considered a regular expression. The list is ordered and checked from top to bottom. Regular expressions are checked first, then the plain records.

Use regex to match DNS requests:

