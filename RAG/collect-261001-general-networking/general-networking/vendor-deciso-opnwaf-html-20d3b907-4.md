---
id: collect-261001-general-networking/general-networking/vendor-deciso-opnwaf-html-20d3b907-4
title: "vendor-deciso-opnwaf-html-20d3b907"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "attention", "memory", "parameters"]
source: docs/RAG/collect-261001-general-networking/vendor-deciso-opnwaf-html-20d3b907.md
source_anchor: ""
source_lines: [215, 321]
sha256: 96cc27cb9df954b11e127780b118d848bbf571284754f766babd0236e56c9559
---

# vendor-deciso-opnwaf-html-20d3b907

| OutlookAnywhere | mail.example.com/rpc - ExternalClientAuthenticationMethod set to Negotiate | 
| ClientAccessService | autodiscover.example.com/Autodiscover/Autodiscover.xml | 
When using a self-signed certificate, the authority for the certificate must be imported into .
The certificate must include mail.example.com and autodiscover.example.com in its SAN.
Without trust established between the OPNsense and the Exchange Server, the connection will fail since only encrypted
connections are allowed to an Exchange Server.
Attention
Exchange Servers use a feature called “Extended Protection” which requires the same certificate to be used by the web application firewall and the Exchange Server itself.
Attention
Outlook passthrough via either RCP/HTTP or MAPI/HTTP is getting more challenging due to the deprecation of the NTML protocol. If authentication popups happen, it can help to set the “Multi Processing Modules” to “mpm-prefork”. Mitigating all authentication issues might not be possible anymore, since the required NTLM replacement protocol OAuth is only available for Exchange Online.
Setup
Create two virtual servers with the hostnames of the Exchange Server, e.g., autodiscover.example.com and
mail.example.com. Select Enable ACME or use your own certificate, set Header Security to Off / compatibility mode,
set Web Protection to Detection Only. Adjust these later once the Exchange Server works correctly through the reverse proxy.
Create a Location with the Type Exchange Server for each of these virtual servers. As Remote destinations use the internal IP address
of the Exchange Server, e.g., https://192.168.10.10. If the virtual servers use the same hostnames as the Exchange Server,
trust is automatically established with host header passthrough.
These new Locations will create all virtual directories the Exchange Server requires automatically.
With the options Restrict Exchange Paths and Access control, access to specific paths can be restricted. This is recommended for the /ecp path.
The finished configuration should look like this:
Virtual Servers
Virtual Server
| Option | Description | 
|---|---|
| Enabled | X | 
| ServerName | mail.example.com | 
| Trust |  | 
| Enable ACME | X | 
| SSL Proxy check peer | X | 
| Security |  | 
| Header Security | Off / compatibility mode | 
| TLS Security profile | Intermediate | 
| Web Protection | Detection Only | 
Location
| Option | Description | 
|---|---|
| Enabled | X | 
| VirtualServer | mail.example.com | 
| Type | Exchange Server | 
| Remote destinations | https://192.168.10.10 | 
| Restrict Exchange Paths | /ecp | 
| Access control | 192.168.0.0/16 172.16.0.0/12 10.0.0.0/8 | 
Virtual Server
| Option | Description | 
|---|---|
| Enabled | X | 
| ServerName | autodiscover.example.com | 
| Trust |  | 
| Enable ACME | X | 
| SSL Proxy check peer | X | 
| Security |  | 
| Header Security | Off / compatibility mode | 
| TLS Security profile | Intermediate | 
| Web Protection | Detection Only | 
Location
| Option | Description | 
|---|---|
| Enabled | X | 
| VirtualServer | autodiscover.example.com | 
| Type | Exchange Server | 
| Remote destinations | https://192.168.10.10 | 
| Restrict Exchange Paths | /ecp | 
| Access control | 192.168.0.0/16 172.16.0.0/12 10.0.0.0/8 | 
Note
In case an internal hostname is used in Remote destinations, ensure this name is in the SAN and common name of the self-signed certificate of the Exchange Server. This hostname must be resolvable from the OPNsense. Do not use the same hostname for Virtual servers and Remote destinations to avoid creating a reverse proxy loop.
Test web protection
When web protection was enabled, we always advise to test if it’s actually functional. Luckily this is quite easy to test using a webbrowser. For this example we will try to inject some sql code in the url, which should be blocked when properly configured:
https://your.example.domain/?id=100 or 'x'='y'
This should show a page similar to the one below:
When deploying web protection for virtual servers, start with the Detection Only setting that can be set per virtual server. This way, you can evaluate the Web Security log file, and look for rules that match.
This will reveal if the web application might be outdated and needs patching, because several web protection rules match and would block connections.
If they are false positives, the rule IDs can be set as exemptions with the option Disable Security Rules by ID. Search the rules in the dropdown, and select multiple ones you want to exclude.
After this configuration, set the Web Protection to On (default) to enable it. The web application should now be configured for production. If there are still errors, repeat the above steps.
Attention
Do not exclude too many rules. These matches could be a potential misconfiguration of the web application behind the WAF. Only exclude rules that totally break the functionality of the web application.
Secure WebDav and HTTP File Servers
These servers have specific requirements to work through a WAF. They need an extended set of HTTP Verbs, and higher thresholds for the Request and Response Body.
A popular example for a WebDAV Server is Nextcloud or Owncloud.
Go to the Web Protection Settings, and set the Allowed HTTP Verbs to:
COPY, DELETE, GET, HEAD, LOCK, MKCOL, MOVE, OPTIONS, POST, PROPFIND, PROPPATCH, PUT, TRACE, UNLOCK.
To allow large file uploads, set Request Body Limit Action to Process Partial. If you want to process as much content of the file as possible, enable the advanced mode and set custom values for the Request Body and Response Body limits.
If the file is larger than the configured limits, it will only be processed partially. This means, the whole file will be uploaded, but only a portion of the file is analyzed by the web application parser. Rejecting can improve security, yet will make large files fail completely if they exceed the configured hard limits.
Note
Increasing the Body limits will increase the log file sizes, and will eventually use the disk of the OPNsense to write files upon inspection. For this, the Request Body in Memory Limit can be increased to 1GB to focus on RAM usage. If you want to use the least resources, logging and disk I/O, leave all settings on default, and set Request Body Limit Action to Process Partial.
Tip
If many different file extensions are hosted on the WebDAV server, some of these will be blocked by default rules. In that case,
disable the rule: 920440 (URL file extension is restricted by policy)
Request Headers
In some cases it is a requirement to manipulate request headers. The Request Header Directive can add, merge, change or remove HTTP request headers.
In our example, we unset the Accept-Encoding header to potentially prevent BREACH attacks.
Go to and create a new header:
| Option | Description | 
|---|---|
| Type | Unset | 
| Header | Accept-Encoding | 
| Value | (leave this empty) | 
Afterwards, go to an existing location in and select it in (Proxy Options) Request Headers.
After applying the configuration, the header will be unset from all requests of this location to the Remote destinations.
Tip
More information about the available request header types can be found here: https://httpd.apache.org/docs/current/mod/mod_headers.html#requestheader
Protect a local server with certificates
In the above virtual host configuration are a couple of parameters related to client authentication. The advantage of using these is that you can prevent unauthorized access to services using certificates signed by a (local) certificate authority.
To use this functionality, first make sure you have a certificate authority defined in which you are going to use to create certificates for your clients.
Next step is to add a VirtualServer which contains at least the following information:
| Option | Description | 
|---|---|
| ServerName | The fully qualified domain name this host listens to | 
