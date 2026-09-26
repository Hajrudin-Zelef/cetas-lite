---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-22-2
title: "Configuration examples"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-22.md
source_anchor: ""
source_lines: [99, 159]
sha256: 7028271da6ee88ad7895cf8cae550cc83d12582b2e2d6e12eab0656397f913cc
---

# Configuration examples

| Property | Description | 
|---|---|
| **always-from-cache** (*yes \| no* ; Default:**no** ) | ignore client refresh requests if the content is considered fresh | 
| **anonymous** (*yes \| no* ; Default:**no** ) | If not set, the IP address of the client would be passed X-Forwarded-For header (could be accessed using HTTP_X_FORWARDED_FOR environment variable in remote servers) | 
| **cache-administrator** (*string* ; Default:**webmaster** ) | Administrator's e-mail displayed on proxy error page | 
| **cache-hit-dscp** (*integer: 0..63* ; Default:**4** ) | Automatically mark cache hit with the provided DSCP value | 
| **cache-on-disk** (*yes \| no* ; Default:**no** ) | Whether to store cache on disk | 
| **cache-path** (*string* ; Default:**web-proxy** ) | A path where the cache will be stored, when cache-on-disk is enabled. | 
| **max-cache-object-size** (*integer: 0..4294967295[KiB]* ; Default:**2048KiB** ) | Specifies the maximal cache object size, measured in kilobytes | 
| **max-cache-size** (*none \| unlimited \| integer: 0..4294967295[KiB]* ; Default:**unlimited** ) | Specifies the maximal cache size, measured in kilobytes | 
| **max-client-connections** (*integer: Dynamic*  ; Default:**600** ) | Maximal number of connections accepted from clients (any further connections will be rejected) | 
| **max-fresh-time** (*time* ; Default:**3d** ) | Maximal time to store a cached object. The validity period of an object is usually defined by the object itself, but in case it is set too high, you can override the maximal value | 
| **max-server-connections** (*integer: Dynamic*  ; Default:**600** ) | Maximal number of connections made to servers (any further connections from clients will be put on hold until some server connections will terminate) | 
| **parent-proxy** (*Ip4 \| ip6* ; Default:**0.0.0.0** ) | IP address and port of another HTTP proxy to redirect all requests to. If set to **0.0.0.0** parent proxy is not used. | 
| **parent-proxy-port** (*integer: 0..65535* ; Default:**0** ) | Port that parent proxy is listening on. | 
| **port** (*integer: 0..65535* ; Default:**8080** ) | TCP port the proxy server will be listening on. This port has to be specified on all clients that want to use the server as an HTTP proxy. A transparent (with zero configuration for clients) proxy setup can be made by redirecting HTTP requests to this port in the IP firewall using the destination NAT feature | 
| **serialize-connections** (*yes \| no* ; Default:**no** ) | Do not make multiple connections to the server for multiple client connections, if possible (i.e. server supports persistent HTTP connections). Clients will be served on the FIFO principle; the next client is processed when the response transfer to the previous one is completed. If a client is idle for too long (max 5 seconds by default), it will give up waiting and open another connection to the server | 
| **src-address** (*Ip4 \| Ip6* ; Default:**0.0.0.0** ) | A proxy will use a specified address when connecting to the parent proxy or website. If set to **0.0.0.0** then the appropriate IP address will be taken from the routing table. | 

### Access List

An access list is configured like regular firewall rules. Rules are processed from the top to the bottom. The first matching rule specifies the decision of what to do with this connection. There is a total of 6 classifiers that specify matching constraints. If none of these classifiers is specified, the particular rule will match every connection.

If a connection is matched by a rule, the action property of this rule specifies whether a connection will be allowed or not. If the particular connection does not match any rule, it will be allowed.

| Property | Description | 
|---|---|
| **action** (*allow \| deny* ; Default:**allow** ) | Specifies whether to pass or deny matched packets | 
| **dst-address** (*Ip4[-Ip4 \| /0..32] \| Ip6/0..128* ; Default: ) | The destination address of the target server. | 
| **dst-host** (*string* ; Default: ) | IP address or DNS name used to make a connection to the target server (this is the string user wrote in a browser before specifying the port and path to a particular web page | 
| **dst-port** (*integer[-integer[,integer[,...]]]: 0..65535* ; Default: ) | List or range of ports the packet is destined to | 
| **local-port** (*integer: 0..65535* ; Default: ) | Specifies the port of the web proxy via which the packet was received. This value should match one of the ports the web proxy is listening on. | 
| **method** (*any \| connect \| delete \| get \| head \| options \| post \| put \| trace* ; Default: ) | The HTTP method used in the request (see HTTP Methods section at the end of this document) | 
| **path** (*string* ; Default: ) | Name of the requested page within the target server (i.e. the name of a particular web page or document without the name of the server it resides on) | 
| **redirect-to** (*string* ; Default: ) | In case of access is denied by this rule, the user shall be redirected to the URL specified here | 
| **src-address** (*Ip4[-Ip4 \| /0..32] \| Ip6/0..128* ; Default: ) | The source address of the connection originator. | 


Read-only properties:

| Property | Description | 
|---|---|
| **hits** (*integer* ) | Count of requests that were matched by this rule | 


Wildcard properties (dst-host and dst-path) match a complete string (i.e., they will not match "example.com" if they are set to "example"). Available wildcards are '*' (match any number of any characters) and '?' (match any one character). Regular expressions are also accepted here, but if the property should be treated as a regular expression, it should start with a colon (':').

Small hints in using regular expressions:

- \\ symbol sequence is used to enter \ character in the console;
- \. pattern means. only (in regular expressions single dot in a pattern means any symbol);
- to show that no symbols are allowed before the given pattern, we use the ^ symbol at the beginning of the pattern;
- to specify that no symbols are allowed after the given pattern, we use the $ symbol at the end of the pattern;
- to enter [ or ] symbols, you should escape them with backslash "\.";

It is strongly recommended to deny all IP addresses except those behind the router as the proxy still may be used to access your internal-use-only (intranet) web servers. Also, consult examples in Firewall Manual on how to protect your router.

### Direct Access

If a **parent-proxy** property is specified, it is possible to tell the proxy server whether to try to pass the request to the parent proxy or to resolve it by connecting to the requested server directly. The direct Access List is managed just like the Proxy Access List described in the previous chapter except for the action argument. Unlike the access list, the direct proxy access list has a default action equal to deny. It takes place when no rules are specified or a particular request did not match any rule.

