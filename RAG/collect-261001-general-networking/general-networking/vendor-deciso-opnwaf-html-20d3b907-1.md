---
id: collect-261001-general-networking/general-networking/vendor-deciso-opnwaf-html-20d3b907-1
title: "vendor-deciso-opnwaf-html-20d3b907"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/vendor-deciso-opnwaf-html-20d3b907.md
source_anchor: ""
source_lines: [1, 60]
sha256: d4a9996f652fc6bee509ddece69089a1948618bb94d994ee77b893c8d5cd28ca
---

# vendor-deciso-opnwaf-html-20d3b907

Web Application Firewall
As part of the OPNsense Business Edition, Deciso offers a plugin to easily protect webservices against all sort of injection attacks and provides encryption for traffic to and from the outside world.
Our Web Application Firewall plugin offers some functionality which can also be found in community plugins available, but in a more user friendly manner. It combines the features most commonly used in reverse proxies, such as TLS offloading and load balancing.
To ease maintenance the OPNWAF plugin offers usage of both internal certificates or newly generated
using the ACME protocol via Let’s Encrypt with a single click.
Prerequisites
Before using this plugin in combination with Let’s Encrypt, make sure port 443 isn’t being used for the web gui of this firewall ().
Note
When using Let’s Encrypt, The Web Application Firewall uses the tls-alpn-01 challenge type for easy domain verification, this requires the virtual server to listen on port 443. Make sure the firewall allows incoming HTTPS connections on port 443. If the client connects via a custom port, you can forward these requests to port 443, and configure the virtual server to forward these requests to the correct internal port.
Installation
To install this plugin, go to and search for os-OPNWAF, the [+] button downloads and installs the software.
Next go to to enable it.
General
Before deep diving into the settings pages, we will explain the most important terminology used in this module.
Virtual servers
A virtual server (also known as a virtual host) is a a concept which allows the use of multiple domains on a single webserver using the same port. In our case it offers the possibility to host various webservers inside your network and forward traffic to them in a secure fashion.
Locations
Locations reside in virtual servers and describe on a path level how requests are being handled, if for example one would
like to forward only a subdirectory (like /api) to a server in the network, the location is where to configure this.
Web protection
The web protection options offer easy access to the OWASP ModSecurity ruleset , which offers a set of generic attack detection rules against a wide range attacks including the OWASP Top Ten.
Setup
Before configuring virtual servers, let’s take a look at the general settings pages (). After installation, the module itself should be enabled by default.
In order to use the integrated ACME client (for Let’s Encrypt), the ACME enable checkbox needs to be set, the certificate agreement needs to be accepted (next checkbox) and contact email needs to be specified.
Optionally a permanent redirect from HTTP to HTTPS can be enabled for all virtual servers. The HTTP port can be customized if necessary by enabling the advanced mode. Do not forget to create an additional firewall rule to allow access to the HTTP Port. When it is non standard, a port forward is necessary (e.g., listen on 8080, forward 80 to 8080).
Web protection is not enabled by default, but you can enable it in the Web protection tab. This is also the place to configure the module and settings which apply for all virtual hosts.
To optionally configure a default catch-all virtual server, select a certificate for Strict SNI Check in .
The certificate can be self-signed and match the CN of the the default server name. Any invalid SNI will now be served an error document.
If a wildcard A-Record is defined for a base domain (e.g., *.example.com in A 203.0.113.1), each subdomain would be answered
by the first configured virtual server otherwise, even if the SNI does not match it.
General Settings
| Option | Description | 
|---|---|
| General Settings |  | 
| Enabled | Enable the gateway webserver. | 
| ACME Settings |  | 
| Enable ACME | Enable the ACME protocol to automatically provision certificates using Let’s Encrypt. This will need the virtual server to be accessible on the standard HTTPS port (443 | 
| ACME Certificate Agreement | When you use mod_md to obtain a certificate, you become a customer of the CA (e.g. Let’s Encrypt). That means you need to read and agree to their Terms of Service, so that you understand what they offer and what they might exclude or require from you. mod_md cannot, by itself, agree to such a thing. | 
| ACME Contact Email | The ACME protocol requires you to give a contact url when you sign up. Currently, Let’s Encrypt wants an email address (and it will use it to inform you about renewals or changed terms of service). | 
| Server Settings |  | 
| HTTP Version | Select the maximum allowed HTTP version this server can use for any connection. If you want to proxy to a location via HTTP/2, use h2:// or h2c:// inside a virtual server. | 
| Redirect HTTP to HTTPS | Enables a permanent redirect (301 Moved Permanently) from HTTP to HTTPS. This will bind the default HTTP port additionally for all virtual hosts. Make sure this port is not bound to a different service, like the default WebGUI redirect rule. | 
| HTTP Port | When enabling the HTTP to HTTPS redirect, this port will be bound for HTTP (default 80). | 
| Multi Processing Modules | Select the processing module the server should use. The default mpm_event is the most scalable module. For older software and state sensitive protocols, mpm_prefork can be chosen. | 
| Strict SNI Check | Respond with an error when a client requests an SNI that does not exist as configured virtual server. This requires a certificate for a default catch-all virtual server; once added the strict SNI check is active. The certificate can be a self-signed one, it is only used as a placeholder for the TLS engine. | 
| Server Name | The default name of this server. It is also used for the default catch-all virtual server if “Strict SNI Check” is enabled. | 
| Error Document | Choose error documents to use for common issues, like page not found. This will be added to the default catch-all virtual server if “Strict SNI Check” is enabled. | 
| Server Limits |  | 
| Limit Internal Recursion | The limit prevents the server from crashing when entering an infinite loop of internal redirects or subrequests. Such loops are usually caused by misconfigurations. | 
| Limit Request Body | The limit (in bytes) on the allowed size of an HTTP request message body. If the client request exceeds that limit, the server will return an error response instead of servicing the request. | 
| Limit Request Fields | The limit on the number of request header fields allowed in an HTTP request. A server needs this value to be larger than the number of fields that a normal client request might include. | 
| Limit Request Size | The limit (in bytes) on the allowed size of an HTTP request header field. A server needs this value to be large enough to hold any one header field from a normal client request. | 
| Limit Request Line | The limit (in bytes) on the allowed size of a client’s HTTP request-line. Since the request-line consists of the HTTP method, URI, and protocol version, the LimitRequestLine directive places a restriction on the length of a request-URI allowed for a request on the server. | 
| Limit XML Request Body | The limit (in bytes) on the maximum size of an XML-based request body. | 
Web Protection
It is recommended to enable the web protection, but not change the default values of individual options without reason. The defaults are tuned for high security while also maintaining good performance and compatibility.
| Option | Description | 
|---|---|
| Enable | Enable content protection to prevent against different types of (injection) attacks | 
| Paranoia Level | A paranoia level of 1 is default. In this level, most core rules are enabled. PL1 is advised for beginners, installations covering many different sites and applications, and for setups with standard security requirements. A higher level will increase sensitivity at the cost of more false positives. | 
