---
id: collect-261001-general-networking/general-networking/vendor-deciso-opnwaf-html-20d3b907-3
title: "vendor-deciso-opnwaf-html-20d3b907"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["apache", "attention"]
source: docs/RAG/collect-261001-general-networking/vendor-deciso-opnwaf-html-20d3b907.md
source_anchor: ""
source_lines: [122, 214]
sha256: 87d6fa16037a566c3e711abd1d7f096b5eedddbabdc37cef7f12b01a9be72552
---

# vendor-deciso-opnwaf-html-20d3b907

| Overlay error pages | Overlay common error pages with the ones specified in the virtual server. | 
| Description | User friendly description for this location | 
| Proxy Options |  | 
| OIDC Auth Required | Require OpenID Connect authentication for this location if a provider has been selected in the virtual server. | 
| OIDC Claims | Select claims that must match for authorization. Multiple claims will be evaluated via OR operator. The default “valid-user” will allow access for any authenticated user in your OIDC scope. | 
| Proxy Options |  | 
| TLS header passthrough | Select which headers to passthrough to the client, all headers will be prefixed with X- to distinct them more easily from the applications perspective. The original headers use underscores (_) these will be replaced for minus (-) signs to prevent applications dropping them. | 
| Request Headers | Select how headers should be processed in the request from this location to the destination | 
| Preserve Host | When enabled, this option will pass the Host: line from the incoming request to the proxied host, instead of the hostname specified in the location. This option should normally be turned Off. It is mostly useful in special configurations like proxied mass name-based virtual hosting, where the original Host header needs to be evaluated by the backend server. | 
| Connection timeout | Connect timeout in seconds. The number of seconds the server waits for the creation of a connection to the backend to complete. | 
| timeout | Socket timeout in seconds. The number of seconds the server waits for data sent by / to the backend. | 
| Response field size | Adjust the size of the proxy response field buffer. The buffer size should be at least the size of the largest expected header size from a proxied response. | 
| nocanon | Normally, mod_proxy will canonicalise ProxyPassed URLs. But this may be incompatible with some backends, particularly those that make use of PATH_INFO. The optional nocanon keyword suppresses this and passes the URL path raw to the backend. Note that this keyword may affect the security of your backend, as it removes the normal limited protection against URL-based attacks provided by the proxy. | 
The options here are quite simple, first you define a path on your end (/ in our example), next you define one or more
destinations this path should map to (for example you could point to a public server here, like https://opnsense.org).
Note
When more than one destination is provided, the load will be balanced automatically.
Tip
Constraining access to allow only specific networks or hosts can be arranged using the Access control input.
Proxy Pass Match
The Proxy Pass Match type is the advanced alternative to Proxy Pass.
Choosing it will turn the Local path field into Location Match, and the new Remote path field into Proxy Pass Match.
These types allow you to match requests based on a regular expression pattern instead of just a literal path.
The match is entered into Local path and the substitution groups can be set in Remote path.
Here is an example how this can look like:
| Option | Description | 
|---|---|
| Local path | ^/manual/(.*)$ | 
| Remote path | /$1 | 
Tip
- ^ : Match start of the URL path
- /manual/ :     Match the literal string /manual/
- (.*) : Capture any characters (zero or more) after /manual/ — this is group 1
- $ : Match end of the string
- $1 : Reference the captured group from the local path. In this example it strips /manual/ from the URL path internally.
Attention
This is an advanced feature for edge cases like stripping paths from requests to form a new base path, or anchoring a path precisely. It can also be used to prevent trailing slashes being attached which break some URL parameter schemes. In most cases using the plain Proxy Pass will give you the desired result automatically.
Redirect
| Option | Description | 
|---|---|
| Enabled | Enable this location | 
| VirtualServer | The server this location belongs to | 
| Local path | Local path of the HTTP request to match (e.g. / for all paths). You can also create multiple location entries, each with their own specific path (e.g./docs ). They will be processed in the order of their creation. | 
| Type | Redirect | 
| HTTP redirection message | Choose the HTTP redirection message. The default is 307, but others like 301 and 308 are also available. | 
| Remote destinations | Locations to redirect requests to, only one is allowed per location per redirect | 
| Access control | List of networks allowed to access this path (empty means any) | 
| Description | User friendly description for this location | 
When setting up a redirect, it will also match HTTP if Redirect HTTP to HTTPS in General Settings has been enabled. If not, only HTTPS is matched.
Note
When a / location with a Redirect has been created, there can’t be any additional ProxyPass locations that match
the same / location, nor a more specific /docs location. The redirect will match first, since it will catch and
redirect all traffic of the virtual server location. What is possible though, is that there is a /docs location that
redirects, and an additional /html location that proxies traffic, in the scope of the same virtual server.
Redirect Match
The Redirect Match type is the advanced alternative to Redirect.
Choosing it will turn the Local path field into Location Match, and the Remote destinations field into Redirect Match.
Wrapping Redirect into a Location behaves slightly differently than doing the same with Proxy Pass, as regular expression groups cannot be passed into the location.
The intended way is using environment variables. See mod_alias for more information.
Here is an example how this can look like:
| Option | Description | 
|---|---|
| Local path | /error/(?<NUMBER>[0-9]+) | 
| Remote destinations | http://example.com/errors/%{env:MATCH_NUMBER}.html | 
Tip
When using the normal Redirect, a common trap is redirects that are infinite due to the apache trailing slash issue.
This can be solved via Redirect Match by setting Local path as ^/?$ which force a match from the start of the
first found slash.
Attention
In most cases using the plain Redirect will give you the desired result automatically.
Exchange Server
| Option | Description | 
|---|---|
| Enabled | Enable this location | 
| VirtualServer | The server this location belongs to | 
| Type | Exchange Server | 
| Remote destinations | Locations to redirect requests to, only one is allowed per location per redirect | 
| Restrict Exchange Paths | Restrict Exchange Server specific paths to networks provided in the Access control field. If paths are selected, exactly these paths will have the Access control attached. Access to path / is filtered per default with a redirect to /owa. All non-selected paths will be allowed from all networks. | 
| Access control | Constrain access to networks provided in this list, when not provided no constraints apply. When type is Exchange Server, it will restrict access to paths selected in Restrict Exchange Paths. | 
| Description | User friendly description for this location | 
Prerequisites
To successfully reverse proxy an Exchange Server, a few conditions must be met:
- The communication between Apache and the Exchange Server must happen via HTTPS.
- The Exchange Server must have its internal and external URLs set correctly, preferably to the same hostnames that will be set as virtual servers.
Common hostname/path combinations are:
| VirtualDirectory | Internal and external URL of Exchange Server | 
|---|---|
| OwaVirtualDirectory | mail.example.com/owa | 
| EcpVirtualDirectory | mail.example.com/ecp | 
| WebServicesVirtualDirectory | mail.example.com/EWS/Exchange.asmx | 
| ActiveSyncVirtualDirectory | mail.example.com/Microsoft-Server-ActiveSync | 
| OabVirtualDirectory | mail.example.com/OAB | 
| MapiVirtualDirectory | mail.example.com/mapi | 
