---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-22-3
title: "Configuration examples"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-22.md
source_anchor: ""
source_lines: [160, 287]
sha256: 90a1289672f4ca1cbdd32a05bd9e36df99502eb7f4806e86040bc2e1b1a8c387
---

# Configuration examples

| Property | Description | 
|---|---|
| **action** (*allow \| deny* ; Default:**allow** ) | Specifies the action to perform on matched packets:  | 
| **dst-address** (*Ip4[-Ip4 \| /0..32] \| Ip6/0..128* ; Default: ) | The destination address of the target server. | 
| **dst-host** (*string* ; Default: ) | IP address or DNS name used to make a connection to the target server (this is the string user wrote in a browser before specifying port and path to a particular web page | 
| **dst-port** (*integer[-integer[,integer[,...]]]: 0..65535* ; Default: ) | List or range of ports used by connection to the target server. | 
| **local-port** (*integer: 0..65535* ; Default: ) | Specifies the port of the web proxy via which the packet was received. This value should match one of the ports the web proxy is listening on. | 
| **method** (*any \| connect \| delete \| get \| head \| options \| post \| put \| trace* ; Default: ) | The HTTP method used in the request (see HTTP Methods section at the end of this document) | 
| **path** (*string* ; Default: ) | Name of the requested page within the target server (i.e. the name of a particular web page or document without the name of the server it resides on) | 
| **src-address** (*Ip4[-Ip4 \| /0..32] \| Ip6/0..128* ; Default: ) | The source address of the connection originator. | 


Read-only properties:

| Property | Description | 
|---|---|
| **hits** (*integer* ) | Count of requests that were matched by this rule | 

### Cache Management

The cache access list specifies, which requests (domains, servers, pages) have to be cached locally by web proxy, and which do not. This list is implemented exactly the same way as the web proxy access list. The default action is to cache an object (if no matching rule is found).

| Property | Description | 
|---|---|
| **action** (*allow \| deny* ; Default:**allow** ) | Specifies the action to perform on matched packets:  | 
| **dst-address** (*Ip4[-Ip4 \| /0..32] \| Ip6/0..128* ; Default: ) | The destination address of the target server | 
| **dst-host** (*string* ; Default: ) | IP address or DNS name used to make a connection to the target server (this is the string user wrote in a browser before specifying port and path to a particular web page | 
| **dst-port** (*integer[-integer[,integer[,...]]]: 0..65535* ; Default: ) | List or range of ports the packet is destined to. | 
| **local-port** (*integer: 0..65535* ; Default: ) | Specifies the port of the web proxy via which the packet was received. This value should match one of the ports the web proxy is listening on. | 
| **method** (*any \| connect \| delete \| get \| head \| options \| post \| put \| trace* ; Default: ) | The HTTP method used in the request (see HTTP Methods section at the end of this document) | 
| **path** (*string* ; Default: ) | Name of the requested page within the target server (i.e. the name of a particular web page or document without the name of the server it resides on) | 
| **src-address** (*Ip4[-Ip4 \| /0..32] \| Ip6/0..128* ; Default: ) | The source address of the connection originator | 

Read-only properties:

| Property | Description | 
|---|---|
| **hits** (*integer* ) | Count of requests that were matched by this rule | 

### Connections

This menu contains the list of current connections the proxy is serving.

Read-only properties:

| Property | Description | 
|---|---|
| **client** () |  | 
| **dst-address** (*Ip4 \| Ip6* ) | IPv4/Ipv6 destination address of the connection | 
| **protocol** (*string* ) | Protocol name | 
| **rx-bytes** (*integer* ) | The number of bytes received by the client | 
| **server** () |  | 
| **src-address** (*Ip4 \| Ip6* ) | Ipv4/ipv6 address of the connection originator | 
| **state** (*closing \| connecting \| converting \| hotspot \| idle \| resolving \| rx-header \| tx-body \| tx-eof \| tx-header \| waiting* ) | Connection state:  | 
| **tx-bytes** (*integer* ) | The number of bytes sent by the client | 

### Cache Inserts

This menu shows statistics on objects stored in a cache (cache inserts).

Read-only properties:

| Property | Description | 
|---|---|
| **denied** (*integer* ) | A number of inserts were denied by the caching list. | 
| **errors** (*integer* ) | Number of disk or other system-related errors | 
| **no-memory** (*integer* ) | Number of objects not stored because there was not enough memory | 
| **successes** (*integer* ) | A number of successful cache inserts. | 
| **too-large** (*integer* ) | Number of objects too large to store | 

### Cache Lookups

This menu shows statistics on objects read from cache (cache lookups).

Read-only properties:

| Property | Description | 
|---|---|
| **denied** (*integer* ) | Number of requests denied by the access list. | 
| **expired** (*integer* ) | Number of requests found in cache, but expired, and, thus, requested from an external server | 
| **no-expiration-info** (*integer* ) | Conditional request received for a page that does not have the information to compare the request with | 
| **non-cacheable** (*integer* ) | Number of requests requested from the external servers unconditionally (as their caching is denied by the cache access list) | 
| **not-found** (*integer* ) | Number of requests not found in the cache, and, thus, requested from an external server (or parent proxy if configured accordingly) | 
| **successes** (*integer* ) | Number of requests found in the cache. | 

### Cache Contents

This menu shows cached contents.

Read-only properties:

| Property | Description | 
|---|---|
| **file-size** (*integer* ) | Cached object size | 
| **last-accessed** (*time* ) |  | 
| **last-accessed-time** (*time* ) |  | 
| **last-modified** (*time* ) |  | 
| **last-modified-time** (*time* ) |  | 
| **uri** (*string* ) |  | 

# HTTP Methods

#### Options

This method is a request for information about the communication options available on the chain between the client and the server identified by the **Request-URI**. The method allows the client to determine the options and (or) the requirements associated with a resource without initiating any resource retrieval

#### GET

This method retrieves whatever information identified by the Request-URI. If the Request-URI refers to a data processing process then the response to the GET method should contain data produced by the process, not the source code of the process procedure(-s), unless the source is the result of the process.

The GET method can become a conditional GET if the request message includes an If-Modified-Since, If-Unmodified-Since, If-Match, If-None-Match, or If-Range header field. The conditional GET method is used to reduce the network traffic specifying that the transfer of the entity should occur only under circumstances described by conditional header field(-s).

The GET method can become a partial GET if the request message includes a Range header field. The partial GET method intends to reduce unnecessary network usage by requesting only parts of entities without transferring data already held by the client.

The response to a GET request is cacheable if and only if it meets the requirements for HTTP caching.

#### HEAD

This method shares all features of GET method except that the server must not return a message-body in the response. This retrieves the metainformation of the entity implied by the request which leads to its wide usage of it for testing hypertext links for validity, accessibility, and recent modification.

The response to a HEAD request may be cacheable in the way that the information contained in the response may be used to update the previously cached entity identified by that Request-URI.

#### POST

This method requests that the origin server accept the entity enclosed in the request as a new subordinate of the resource identified by the Request-URI.

The actual action performed by the POST method is determined by the origin server and usually is Request-URI dependent.

