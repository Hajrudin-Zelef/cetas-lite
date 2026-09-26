---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-36-4
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-36.md
source_anchor: ""
source_lines: [155, 264]
sha256: d90d97fdef6be458315277287884db9abecd30a8c8073a12dd860d36312729c4
---

# Introduction

| Parameters | Description | 
|---|---|
| **server** (read-only; name) | HotSpot server name client is logged in | 
| **user** (read-only; name) | name of the HotSpot user | 
| **domain** (read-only; text) | the domain of the user (if split from the username), a parameter is used only with RADIUS authentication | 
| **address** (read-only; IP address) | The IP address of the HotSpot user | 
| **mac-address** (read-only; MAC-address) | MAC-address of the HotSpot user | 
| **login-by** (read-only; multiple-choice: cookie**/** http-chap**/** http-pap**/** https**/** mac**/** mac-cookie**/** trial) | the authentication method used by the HotSpot client | 
| **uptime** (read-only; time) | current session time of the user, it is showing how long the user has been logged in | 
| **idle-time** (read-only; time) | the amount of time the user has been idle | 
| **session-time-left** (read-only; time) | the exact value of session-time, that is applied for the user. Value shows how long user is allowed to be online to be logged off automatically by **uptime** reached | 
| **idle-timeout** (read-only; time) | the exact value of the user's idle-timeout | 
| **keepalive-timeout** (read-only; time) | the exact value of the keepalive-timeout, that is applied for the user. Value shows how long the host can stay out of reach to be removed from the HotSpot | 
| **limit-bytes-in** (read-only; integer) | value shows how many bytes received from the client, an option is active when the appropriate parameter is configured for HotSpot user | 
| **limit-bytes-out** (read-only; integer) | value shows how many bytes send to the client, an option is active when the appropriate parameter is configured for HotSpot user | 
| **limit-bytes-total** (read-only; integer) | value shows how many bytes total were send/received from the client, an option is active when the appropriate parameter is configured for HotSpot user | 

# HotSpot Host

The host table lists all computers connected to the HotSpot server. The host table is informational and it is not possible to change any value there:

| Parameters | Description | 
|---|---|
| **mac-address** (read-only; MAC-address) | HotSpot user MAC-address | 
| **address** (read-only; IP address) | HotSpot client original IP address | 
| **to-address** (read-only; IP address) | The new client address assigned by HotSpot might be the same as the original **address** | 
| **server** (read-only; name) | HotSpot server name client is connected to | 
| **bridge-port** (read-only; name) | *"/interface bridge port"* the client is connected to, value is unknown when HotSpot is not configured on the bridge | 
| **uptime** (read-only; time) | value shows how long the user is online (connected to the HotSpot) | 
| **idle-time** (read-only; time) | time user has been idle | 
| **idle-timeout** (read-only; time) | value of the client idle-timeout (unauthorized client) | 
| **keepalive-timeout** (read-only; time) | keepalive-timeout value of the unauthorized client | 
| **bytes-in** (read-only; integer) | amount of bytes received from an unauthorized client | 
| **packet-in** (read-only; integer) | amount of packets received from an unauthorized client | 
| **bytes-out** (read-only; integer) | amount of bytes sent to an unauthorized client | 
| **packet-out** (read-only; integer) | amount of packets sent to an unauthorized client | 

# HotSpot walled-garden

Walled garden is a system which allows unauthorized use of some resources, but requires authorization to access other resources. This is useful, for example, to give access to some general information about HotSpot service provider or billing options

The menu only manager Walled Garden for HTTP and HTTPs protocols. Other protocols can also be include in Walled Garden, but that is configured elsewhere (in /ip hotspot walled-garden ip).

| Property | Description | 
|---|---|
| **action** (*allow \| deny* ; Default:**allow** ) | Action to perform, when packet matches the rule  | 
| **server** (*string* ; Default: ) | Name of the HotSpot server, rule is applied to. | 
| **src-address** (*IP* ; Default: ) | Source address of the user, usually IP address of the HotSpot client | 
| **method** (*string* ; Default: ) | HTTP method of the request | 
| **dst-host** (*string* ; Default: ) | Domain name of the destination web-server | 
| **dst-port** (*integer* ; Default: ) | TCP port number, client sends request to | 
| **path** (*string* ; Default: ) | The path of the request, path comes after '''http://dst_host' | 

**Read-only properties**

| Property | Description | 
|---|---|
| **dst-address** (*IP* ) |  | 
| **hits** (*integer* ) |  | 

Wildcard properties (dst-host and path) match a complete string (i.e., they will not match "example.com" if they are set to "example"). Available wildcards are '*' (match any number of any characters) and '?' (match any one character). Regular expressions are also accepted here, but if the property should be treated as a regular expression, it should start with a colon (':'). To show that no symbols are allowed before the given pattern, we use ^ symbol at the beginning of the pattern. To specify that no symbols are allowed after the given pattern, we use $ symbol at the end of the pattern.

## Example

To only permit bypassed access in walled garden to "www.example.com/test" but not to "www.example.com/test/test.php" :

/ip hotspot walled-garden
add dst-host=:^www.example.com path=":/test\$"

# HotSpot walled-garden ip

To bypass HotSpot authentication for other protocols and different src/dst addresses (or address-lists). Used for different services (Winbox, SSH, Telnet, SIP, etc.)

| Property | Description | 
|---|---|
| **action** (*accept \|drop\|reject* ; Default:**allow** ) | Action to perform, when packet matches the rule  | 
| **server** (*string* ; Default: ) | Name of the HotSpot server, rule is applied to. | 
| **src-address** (*IP* ; Default: ) | Source address of the user, usually IP address of the HotSpot client | 
| **dst-address** (*IP* ; Default: ) | Destination IP address, IP address of the WEB-server. Ignored if **dst-host** is already specified. | 
| **src-address-list** (*string* ; Default: ) | Source address list name | 
| **dst-address-list** (*string* ; Default: ) | Destination address list. Ignored if **dst-host** is already specified. | 
| **dst-host** (*string* ; Default: ) | Domain name of the destination web-server. When this parameter is specified dynamic entry is added to Walled Garden | 
| **dst-port** (*integer* ; Default: ) | TCP port number, client sends request to | 
| **protocol** (*integer \| string* ; Default: ) | IP protocol | 

# IP Binding

IP-Binding HotSpot menu allows to the setup of static One-to-One NAT translations, allows to bypass specific HotSpot clients without any authentication, and also allows to block specific hosts and subnets from the HotSpot network

| Property | Description | 
|---|---|
| **address** (*IP Range* ; Default:**""** ) | The original IP address of the client | 
| **mac-address** (*MAC* ; Default:**""** ) | MAC address of the client | 
| **server** (*string \| all* ; Default:**"all"** ) | Name of the HotSpot server.  | 
| **to-address** (*IP* ; Default:**""** ) | New IP address of the client, translation occurs on the router (client does not know anything about the translation) | 
| **type** (*blocked \| bypassed \| regular* ; Default:**""** ) | Type of the IP-binding action  | 

# Cookies

The menu contains all cookies sent to the HotSpot clients, which are authorized by cookie method, all the entries are read-only.

| Property | Description | 
|---|---|
| **domain** (*string* ) | The domain name (if split from the username) | 
| **expires-in** (*time* ) | How long the cookie is valid | 
| **mac-address** (*MAC* ) | Client's MAC-address | 
| **user** (*string* ) | HotSpot username | 

# MAC Cookie

