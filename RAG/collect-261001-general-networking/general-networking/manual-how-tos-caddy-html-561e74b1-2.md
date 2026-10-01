---
id: collect-261001-general-networking/general-networking/manual-how-tos-caddy-html-561e74b1-2
title: "Reverse Proxy Domain: \"531e7877-0b58-4f93-a9f0-54beee58bdea\""
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-how-tos-caddy-html-561e74b1.md
source_anchor: ""
source_lines: [153, 279]
sha256: 9b0bfa88b5d023f575eb1688e573b5869e967ca1812c6dc9f696292a53adb808
---

# Reverse Proxy Domain: "531e7877-0b58-4f93-a9f0-54beee58bdea"

For Cloudflare, set Trusted Proxies to the Cloudflare IP ranges and Client IP Headers to Cf-Connecting-Ip.
Reverse Proxy the OPNsense WebGUI
Tip
The same approach can be used for any upstream destination using TLS and a self-signed certificate.
Attention
- Open the OPNsense WebGUI in a browser (e.g. Chrome or Firefox). Inspect the certificate by clicking on the 🔒 in the address bar. Copy the SAN for later use. It can be a hostname, for exampleOPNsense.localdomain
- Save the certificate as.pem file. Open it up with a text editor, and copy the contents into a new entry in . Name the certificateopnsense-selfsigned
- Add a new Domain, for exampleopn.example.com
- Add a new Handler with the following options:
| Options | Values | 
|---|---|
| Frontend |  | 
| Domain: | opn.example.com | 
| Upstream |  | 
| Protocol | https:// | 
| Upstream Domain: | 127.0.0.1 | 
| Upstream Port: | 8443 - WebGUI Port | 
| TLS Trust Pool: | opnsense-selfsigned | 
| TLS Server Name: | OPNsense.localdomain | 
- Press Save and Apply
Go to
- Input opn.example.com in Alternate Hostnames to prevent the error: The HTTP_REFERER “https://opn.example.com/” does not match the predefined settings
- Press Save
Open https://opn.example.com and it should serve the reverse proxied OPNsense WebGUI. Check the log file for errors if it does not work, most of the time the TLS Server Name doesn’t match the SAN of the TLS Trust Pool. Caddy does not support certificates with only a CN Common Name.
Attention
Create an Access List to restrict access to the WebGUI.
Redirect ACME HTTP-01 Challenge
Sometimes an application behind Caddy uses its own ACME Client to get certificates, most likely with the HTTP-01 challenge. This plugin has a built in mechanism to redirect this challenge type easily to a destination behind it.
Make sure the chosen domain is externally resolvable. Create an A-Record on a public DNS server that points to the external IP Address of the OPNsense. In case of IPv6 availability, it is mandatory to create an AAAA-Record too, otherwise the TLS-ALPN-01 challenge might fail.
The configured Domain must use an empty port or 443 in the GUI, otherwise it can not use the TLS-ALPN-01 challenge for itself. The upstream destination must listen on Port 80 and serve /.well-known/acme-challenge/, for the same Domain that is configured in Caddy.
Go to
- Press ✎ and open an existing Domain or Subdomain and enable advanced mode
| Options | Values | 
|---|---|
| Domain: | foo.example.com | 
| HTTP-01 Challenge Redirection: | 192.168.10.1 | 
- Press Save and Apply
The HTTP-01 Challenge Redirection is active and the upstream destination located at 192.168.10.1 will be able to issue the certificate for the domain foo.example.com.
With this configuration, Caddy will choose the TLS-ALPN-01 challenge to get its own certificate for foo.example.com, and reverse proxy the HTTP-01 challenge to 192.168.10.1, where the upstream destination can listen on port 80 for foo.example.com. With TLS enabled in the Handler, an encrypted connection is automatically possible. The automatic HTTP to HTTPS redirection is also taken care of.
Filter by Domain
A large configuration can be challenging to navigate. To help, a filter functionality has been added to the top right corner of the Domains, Subdomains and Handlers tab, called Filter by Domain.
In Filter by Domain, one or multiple Domains and Subdomains can be selected. As filter result, only their corresponding configuration will be displayed in Domains, Subdomains and Handlers.
When creating a new Subdomain or Handler, the selection in the filter will be added automatically to the open form.
This filter is also used by the Add Handler and Search Handler buttons to scope the correct selection.
Advanced Configuration
Multiple Handlers for the same Domain
Handlers are not limited to one per domain or subdomain. If there are multiple different paths to handle (e.g. /foo/* and /bar/*), create a Handler for each of them.
When creating a Handler with an empty path, the templating logic will automatically place it last in the Caddyfile site block. This means, specific paths will always match before an empty path, regardless of their position in the configuration. This could be used to block specific paths with an Access List, route some paths to different upstreams, and then set an empty handle for all unmatched paths.
Different handling logics can be selected. E.g., handle_path to strip the path from all requests, or handle to preserve the path from all requests.
When using a mix of wildcard domains and subdomains, a Handler set exclusively on the wildcard domain will match after all subdomains. That way, all unmatched subdomains can be sent to a custom upstream.
Multiple domains with the same hostname and different ports can be created at the same time. E.g., opn.example.com:443 and opn.example.com:8443. Now the frontend can listen on multiple ports for the same domain. These domains will share the same certificate automatically if ACME manages them. Each of these sockets need their own Handler to proxy traffic.
An example Caddyfile could look like this:
# Reverse Proxy Domain: "531e7877-0b58-4f93-a9f0-54beee58bdea"
opn.example.com:443 {
        handle /private/* {
                @d72c1182-6f05-4c25-8d9f-6a226a9039ea {
                        not client_ip 192.168.0.0/16 172.16.0.0/12 10.0.0.0/8
                }
                handle @d72c1182-6f05-4c25-8d9f-6a226a9039ea {
                        abort
                }
                reverse_proxy 172.16.99.10:8443 {
                }
        }
        handle /different_upstream/* {
                reverse_proxy 192.168.1.33 {
                }
        }
        handle {
                reverse_proxy 172.16.99.10:8443 {
                }
        }
}
# Reverse Proxy Domain: "58760ae1-2409-4a6b-a6c4-d58b15706b55"
opn.example.com:8443 {
        handle_path /strip_this {
                reverse_proxy 10.10.10.10:8443 {
                }
        }
}
Tip
Access Lists and Basic Auth can match directly on Handlers for more complex access control scenarios.
Reverse Proxy a Webserver with Vhosts
Sometimes it is necessary to alter the host header in order to reverse proxy to another webserver with vhosts.
Since Caddy passes the original host header by default (e.g. app.external.example.com), if the upstream destination listens on a different hostname (e.g. app.internal.example.com), it would not be able to serve this request.
Go to
- Press + to create a new Domain
| Options | Values | 
|---|---|
| Domain: | app.external.example.com | 
- Press Save
Go to
- Press + to create a new HTTP Header
| Options | Values | 
|---|---|
| Header: | header_up | 
| Header Type: | Host | 
| Header Value: | {upstream_hostport} | 
- Press Save
Go to
- Press + to create a new Handler and open the Transport section.
| Options | Values | 
|---|---|
| Domain: | app.external.example.com | 
| Upstream Domain: | app.internal.example.com | 
| HTTP Headers: | header_up Host {upstream_hostport} | 
- Press Save and Apply
CrowdSec Integration
CrowdSec is a powerful alternative to a WAF. It uses logs to dynamically ban IP addresses of known bad actors. The Caddy plugin is prepared to emit the json logs for this integration.
Go to
- Enable Log HTTP Access in JSON Format
- Press Save
Go to
- Open each Domain that should be monitored by CrowdSec and open Access
- Enable HTTP Access Log
Now the HTTP access logs will appear in /var/log/caddy/access in json format, one file for each domain.
Next, connect to the OPNsense via SSH or console, go into the shell with Option 8.
Attention
This step requires the os-crowdsec plugin.
- Once in the shell, install the caddy collection from CrowdSec Hub. cscli collections install crowdsecurity/caddy
- Create the configuration file as /usr/local/etc/crowdsec/acquis.d/caddy.yaml with the following content:
filenames:
  - /var/log/caddy/access/*.log
force_inotify: true
poll_without_inotify: true
labels:
  type: caddy
