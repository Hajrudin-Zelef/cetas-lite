---
id: collect-261001-general-networking/general-networking/vendor-deciso-opnwaf-html-20d3b907-2
title: "vendor-deciso-opnwaf-html-20d3b907"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/vendor-deciso-opnwaf-html-20d3b907.md
source_anchor: ""
source_lines: [61, 121]
sha256: 18db161c495635a9076b8eea422fbba3c42268a7ff1594604a77c800741ff043
---

# vendor-deciso-opnwaf-html-20d3b907

| Allowed HTTP Verbs | Leave empty to use the default HTTP Verbs. Choosing one or more items here will overwrite the default globally. If there are WebDAV servers, adjusting these Verbs can be mandatory with methods like PROPPATCH or PROPFIND. | 
| Request Body Limit Action | What to do if the request body size is above the configured limit. The default is Reject. Keep in mind that this setting will automatically be set to ProcessPartial when using DetectionOnly mode. | 
| Request Body Limit | Maximum request body size in bytes we will accept for buffering. If you support file uploads then the value given has to be as large as the largest file you are willing to accept. The hard limit is 1GB. | 
| Request Body No Files Limit | Maximum request body size in bytes we will accept for buffering with files excluded. You want to keep this value as low as practical. The hard limit is 1GB. | 
| Request Body In Memory Limit | Store the request body data in memory. When the multipartparser reaches this limit, it will start using your hard disk for storage. That is slow, but unavoidable. The hard limit is 1GB. | 
| Response Body Limit Action | What happens when we encounter a response body larger than the configured limit? By default, we process what we have and let the rest through with the Process Partial option. That is somewhat less secure, but does not break any legitimate pages. | 
| Response Body Limit | Maximum response body size (in bytes) we will accept for buffering. The hard limit is 1GB. | 
| Regex Match Limit | Maximum regex matching length in security rules. If set too high, could cause performance issues or DoS. | 
| Regex Match Limit Recursion | Maximum regex recursion matching length in security rules. If set too high, could cause performance issues or DoS. | 
Configure virtual servers
With the general settings in place, we can start adding virtual servers to offload traffic to machines in our network. First go to and click on the [+] in the top section of the screen, which defines the virtual servers.
| Option | Description | 
|---|---|
| Enabled | Enable this virtual server. | 
| LogLevel | (advanced mode) Log verbosity level | 
| ServerName | Fully qualified hostname for this server. | 
| Port | Port number this vhost will listen on, can easily be combined with firewall nat rules to map traffic to non standard ports when origination from remote destinations. (e.g., listen on 8443, forward 443 to 8443). | 
| Error Document | Choose error documents to use for common issues, like page not found. | 
| Description | User friendly description for this vhost (optional). | 
| Trust |  | 
| Enable ACME | Enable the ACME protocol to automatically provision certificates using Let’s Encrypt, when set will ignore the selected certificate (and enable SSL on this virtual server). | 
| Certificate | When using a certificate available in the system trust store, select it here. | 
| SSL Proxy check peer | This directive configures host name checking for server certificates when mod_ssl is acting as an SSL client. The check will succeed if the host name from the request URI matches one of the CN attribute(s) of the certificate’s subject, or matches the subjectAltName extension. If the check fails, the SSL request is aborted and a 502 status code (Bad Gateway) is returned. | 
| Client Auth |  | 
| CA for client auth | Require a client certificate signed by the provided authority before allowing a connection. | 
| CRL for client auth | Attach the (first) found certificate revocation list for the selected CA to this virtual host. Please note when no CRL is offered all clients are rejected. | 
| Verify depth for client auth | The depth actually is the maximum number of intermediate certificate issuers, i.e. the number of CA certificates which are max allowed to be followed while verifying the client certificate. | 
| OpenID Connect |  | 
| OIDC Provider | Select an OpenID Connect Provider for authentication created in “System - Access - OpenID Connect”. Afterwards, select the claim in the individual locations of this virtual server. | 
| OIDC Redirect URI | The redirect_uri for this OpenID Connect client; this is a vanity URL that must ONLY point to a path on your server protected by this module but it must NOT point to any actual content that needs to be served. Leave empty to use the provided default. | 
| OIDC HTTP Timeout Short | Timeout in seconds for short duration HTTP calls. This defines the maximum duration that a request may take to complete and is used for Client Registration and OP Discovery requests. | 
| OIDC HTTP Timeout Long | Timeout in seconds for long duration HTTP calls. This defines the maximum duration that a request make take to complete and is used for most requests to remote endpoints. | 
| OIDC Pass Claims As | Select how claims should be passed from the virtual server to the location. The default sends them as headers. | 
| Security |  | 
| Header Security | Header security, by default several privacy and security related headers are set, in some cases (old applications for example) you might want to disable sending default headers to clients. HSTS can be disabled here if necessary. | 
| TLS Security profile | TLS security profile as documented by Mozilla | 
| Disable Security Rules by ID | Select one or multiple Web Protection rules to disable via their IDs. This can help to selectively disable rules that cause false positives, without disabling the Web Protection completely. | 
| Web Protection | When Web Protection is enabled for the host you may disable it for specific destinations here, or set it to detection only for logging purposes. | 
The section above defines the port the virtual server will listen on. Remember, in order to use ACME (Let’s encrypt) this should either be 443 or the traffic should be forwarded from port 443 to the port defined here.
Note
Port numbers can be reused. Multiple virtual servers can share the same port. Hostnames must be unique. They are used to identify the virtual server via SNI (Server Name Indication).
Warning
The ALPN protocol (the challenge type used by Let’s Encrypt) will resolve the FQDNs specified in the virtual host entry to the IP address of the firewall. If your DNS records point to both IPv4 and IPv6 addresses, IPv6 will be preferred by the challenge, so make sure your firewall is reachable via IPv6 as well if this is the case.
When supplying a certificate manually via the system trust store you can assign it in this dialog as well.
Configure locations
The virtual server itself doesn’t provide much content to the user other than offering a page telling access is prohibited, so the next step is to map directories to external locations. These can be defined in the Locations grid underneath the Virtual servers.
There are different types of locations:
- Proxy Pass, which Reverse Proxies the HTTP traffic
- Proxy Pass Match, which Reverse Proxies the HTTP traffic but has regex support
- Redirect, which creates a HTTP redirect
- Redirect Match, which creates a HTTP redirect but has regex support
- Exchange Server, a template for Microsoft Exchange servers
Proxy Pass
| Option | Description | 
|---|---|
| Enabled | Enable this location | 
| VirtualServer | The server this location belongs to | 
| Local path | Local path of the HTTP request to match (e.g. / for all paths). You can also create multiple location entries, each with their own specific path (e.g./docs ). They will be processed in the order of their creation. | 
| Type | ProxyPass | 
| Remote destinations | Locations to forward requests to, when more than one is provided, requests will be loadbalanced in a round robin fashion. Supports http ,https ,wswss ,h2 andh2c destinations. When your webapp uses websockets and https requests, usewss:// | 
| Access control | List of networks allowed to access this path (empty means any) | 
