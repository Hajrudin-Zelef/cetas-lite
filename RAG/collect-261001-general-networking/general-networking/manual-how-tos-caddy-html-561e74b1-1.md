---
id: collect-261001-general-networking/general-networking/manual-how-tos-caddy-html-561e74b1-1
title: "Reverse Proxy Domain: \"531e7877-0b58-4f93-a9f0-54beee58bdea\""
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-how-tos-caddy-html-561e74b1.md
source_anchor: ""
source_lines: [1, 152]
sha256: 3d708fc08f366834681d388db9cefc591e99722e27d8d6f5e8a16ce3bbe72445
---

# Reverse Proxy Domain: "531e7877-0b58-4f93-a9f0-54beee58bdea"

Caddy: Reverse Proxy
Features
Fast and extensible multi-platform HTTP/1-2-3 web server with automatic HTTPS
By default, Caddy automatically obtains and renews TLS certificates (Let’s Encrypt and ZeroSSL) for all your sites.
- Reverse Proxy HTTP, HTTPS and WebSockets
- Route UDP/TCP traffic with the included Layer4 module: https://github.com/mholt/caddy-l4
- Dynamic DNS module included: https://github.com/mholt/caddy-dynamicdns
- Cloudflare DNS Provider included: https://github.com/caddy-dns/cloudflare
Installation
- Install “os-caddy” from the OPNsense Plugins.
Prepare OPNsense for Caddy After Installation
Attention
Caddy uses port 80 and 443. So the OPNsense WebGUI or other plugins can’t bind to these ports.
Go to
- Change the TCP Port to 8443 (example), do not forget to adjust the firewall rules to allow access to the WebGUI. On LAN there is a hidden anti-lockout rule that takes care of this automatically. On other interfaces, make sure to add explicit rules.
- Enable the checkbox for HTTP Redirect - Disable web GUI redirect rule.
Go to
- Create Firewall rules that allow HTTP andHTTPS to destinationThis Firewall onWAN
| Option | Values | 
|---|---|
| Interface | WAN | 
| TCP/IP Version | IPv4+IPv6 | 
| Protocol | TCP | 
| Source | Any | 
| Destination | This Firewall | 
| Destination port range | from: HTTP to:HTTP | 
| Description | Caddy Reverse Proxy HTTP | 
| Option | Values | 
|---|---|
| Interface | WAN | 
| TCP/IP Version | IPv4+IPv6 | 
| Protocol | TCP/UDP | 
| Source | Any | 
| Destination | This Firewall | 
| Destination port range | from: HTTPS to:HTTPS | 
| Description | Caddy Reverse Proxy HTTPS | 
Go to and create the same rules for the LAN interface. Now external and internal clients can connect to Caddy, and Let’s Encrypt or ZeroSSL certificates will be issued automatically. If using a VPN to connect remote clients to the OPNsense, additional firewall rules could be needed.
Note
If you disable QUIC by removing HTTP/3 in , the Caddy Reverse Proxy HTTPS rule only needs TCP as protocol.
Standard Configuration
Note
The tutorial section implies that Prepare OPNsense for Caddy after installation has been followed.
Creating a Simple Reverse Proxy
Attention
The domain has to be externally resolvable. Create an A-Record on a public DNS server that points your domain to the external IP address of your OPNsense.
Go to
- Check Enabled to enable Caddy
- Input a valid email address into the Acme Email field. This is mandatory to receive automatic Let’s Encrypt and ZeroSSL certificates
- Auto HTTPS should be set toOn (default)
- Press Apply
Go to
- Press + to add a Domain as frontend.
| Options | Values | 
|---|---|
| Frontend |  | 
| Protocol: | https:// | 
| Domain: | foo.example.com | 
| Port: | Leave empty | 
| Certificate: | Auto HTTPS | 
- Press Save
- Go to
- Press + to add a Handler that routes the traffic from the frontend to a target upstream service.
| Options | Values | 
|---|---|
| Frontend |  | 
| Domain: | https://foo.example.com | 
| Upstream |  | 
| Protocol: | http:// orhttps:// - depending on your upstream webserver | 
| Upstream Domain: | 192.168.10.1 | 
| Upstream Port: | 80 - or set the port required by your upstream webserver | 
| TLS Insecure Skip Verify | X - if https:// was chosen | 
- Press Save and Apply
The automatic certificate will be installed. Check the Logfile if there are errors. Now the frontend domain foo.example.com:80/443 receives all requests, and reverse proxies them to the upstream destination 192.168.10.1:80 (or custom port).
Tip
Issued certificates can be verified in and in the associated dashboard widget.
Note
TLS Insecure Skip Verify can be used in private networks. If the upstream destination is in an insecure network consider using proper certificate handling.
Restrict Access to Internal IPs
Since the reverse proxy will accept all connections, restricting access with a firewall rule would impact all domains. Access Lists can restrict access per domain. In this example, they are used to restrict access to only internal IPv4 networks, refusing connections from the internet.
Go to
- Press + to create a new Access List
| Options | Values | 
|---|---|
| Access List Name: | private_ipv4 | 
| Client IP Addresses: | 192.168.0.0/16172.16.0.0/1210.0.0.0/8 | 
| Description: | Allow access from private IPv4 ranges | 
- Press Save
Go to
- Edit an existing Domain or Subdomain and expand the Access Tab.
| Options | Values | 
|---|---|
| Access List: | private_ipv4 | 
- Press Save and Apply
Now, all connections without a private IPv4 address will be blocked. Some applications might demand a HTTP Error code instead of having their connection blocked, an example are monitoring systems. For these a custom HTTP Response Code can be set in the advanced mode.
Note
Access Lists can be set on Domains, Subdomains and Handlers. Setting them on Domains or Subdomains is recommended for simplicity.
Restrict Access with Basic Auth
Since the reverse proxy will accept all connections, restricting access with a firewall rule would impact all domains. Basic Auth will restrict access to one or multiple users.
Go to
- Press + to create a new User
| Options | Values | 
|---|---|
| User: | John | 
| Password: | RandomPassword | 
- Press Save and create additional Users if needed, e.g. Sarah .
Go to
- Edit an existing Domain or Subdomain and expand the Access Tab.
| Options | Values | 
|---|---|
| Basic Auth: | John ,Sarah | 
- Press Save and Apply
Now, all anonymous connections have to authenticate with Basic Auth before accessing the reverse proxied service.
Note
Basic Auth can be set on Domains, Subdomains and Handlers. Setting it on Domains or Subdomains is recommended for simplicity.
Tip
For even higher security demands, configure Client Auth (mTLS) on a domain.
Dynamic DNS
Go to
- Select Cloudflare from the list
- Input the API Key
- Choose if DynDns IP Version should include IPv4 and/or IPv6.
- Press Save
Go to
- Edit a domain or subdomain and enable the Dynamic DNS checkbox.
- Press Save and Apply
Check the Logfile for the DynDNS updates. Set it to Informational and search for the chosen domain.
Note
Enabling the Dynamic DNS checkboxes can have different results:
- Base Domain: example.com @
- Wildcard Domain: example.com *
- Subdomain: example.com opn
Use subdomains if you see errors in the log like:
failed setting DNS record(s) with new IP address(es)”,”zone”:”opn.example.com”,”error”:”expected 1 zone, got 0
This means the zone opn.example.com @ does not exist, and the provider expects example.com opn for the update. You can see the current configuration in .
Wildcard Domain with Subdomains
Tip
For Cloudflare, this is the recommended setup.
Note
If you use Dynamic DNS, subdomains are needed due to the way the API updates the DNS Records in hosted zones.
Go to
- Select Cloudflare from the list
- Input the API Key
- Set Resolvers to 1.1.1.1
Go to
- Create*.example.com as domain and activate the DNS-01 Challenge checkbox. Alternatively, use a certificate imported or generated in . It has to be a wildcard certificate. You could generate one with the os-acme-client plugin.
- Create all subdomains in relation to the*.example.com domain, for examplefoo.example.com andbar.example.com .
- Check Dynamic DNS for the new subdomains, if needed.
Go to
- Create a Handler with *.example.com as domain andfoo.example.com as subdomain. Most of the same configuration as with base domains are possible. The subdomain dropdown only shows when a wildcard domain has been configured.
Note
The certificate of a wildcard domain will only contain *.example.com, not a SAN for example.com. If there is a service that should match example.com exactly, create an additional domain for example.com with an additional Handler for its upstream destination. Subdomains do not support setting ports, they will always track the ports of their assigned parent wildcard domain.
Tip
