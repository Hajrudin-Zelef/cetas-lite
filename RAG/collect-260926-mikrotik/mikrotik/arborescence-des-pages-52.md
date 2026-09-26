---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-52
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-52.md
source_anchor: ""
source_lines: [1, 27]
sha256: e78ab4e090174f4a70d9f32ce3a66603efa6cd956f076bb91947d1dae5d5c3b9
---

# Introduction

**Standards:** `RFC 2136, RFC 3007`

Dynamic DNS Update Tool gives a way to keep the domain name pointing to a dynamic IP address. It works by sending a domain name system update requests to the name server, which has a zone to be updated. Secure DNS updates are also supported.

The DNS update tool supports only one algorithm - **hmac-md5**. It's the only proposed algorithm for signing DNS messages.

DNS update tool works only with the BIND server, it will not work with DynDNS, EveryDNS, or any other similar service.

# Properties

| Property | Description | 
|---|---|
| **address** (*IP* ; Default: ) | Defines the IP address associated with the domain name. | 
| **dns-server** (*IP* ; Default: ) | DNS server to send updates to. | 
| **key** (*string* ; Default: ) | Authorization key to access the server. | 
| **key-name** (*string* ; Default: ) | Authorization key name (like a username) to access the server. | 
| **name** (*string* ; Default: ) | Name to attach with the IP address. | 
| **ttl** (*integer* ; Default: ) | Time to live for the item (in seconds). | 
| **zone** (*string* ; Default: ) | DNS zone where to update the domain name in. | 

The system clock time on your router can't differ from the DNS server's time by more than 5 minutes. Otherwise, the DNS server will ignore this request.

# Example

To tell 23.34.45.56 DNS server to (re)associate mydomain name in the myzone.com zone with 68.42.14.4 IP address specifying that the name of the key is dns-update-key and the actual key updates:
