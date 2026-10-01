---
id: collect-261001-general-networking/general-networking/manual-how-tos-caddy-html-561e74b1-4
title: "Reverse Proxy Domain: \"531e7877-0b58-4f93-a9f0-54beee58bdea\""
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-how-tos-caddy-html-561e74b1.md
source_anchor: ""
source_lines: [367, 495]
sha256: f6d929cb394547e36f5a7fd540a47a36ded4f889ecfdc1196351baf242872e76
---

# Reverse Proxy Domain: "531e7877-0b58-4f93-a9f0-54beee58bdea"

- QUIC (with and without Client Hello evaluation)
- RDP
- SOCKSv4/v5
- SSH
- TLS (with and without Client Hello evaluation)
- Winbox
- Wireguard
- XMPP
Configuration Examples
SSH Multiplexing on HTTPS Port
SSH is a raw protocol matcher, it will match all traffic that looks like SSH in the scope of either the listener_wrapper, or a TCP port in global. Host Headers or SNI can not be evaluated since SSH does not send this information.
In this example, we want to allow SSH on the default HTTPS port. This will route the SSH traffic to a selected upstream and all unmatched traffic to the Reverse Proxy.
- Go to
- Press + to create a new Layer4 Route
| Options | Values | 
|---|---|
| Matchers: | SSH | 
| Upstream Domain: | 192.168.1.1 | 
| Upstream Port: | 22 | 
- Press Save and Apply
Now an SSH client can open a connection like ssh app1.example.com -p 443 and the SSH traffic will go through the same port as other HTTP/HTTPS traffic. Caddy becomes a protocol multiplexer.
Tip
If another route is added, e.g. with the RDP matcher, then SSH and RDP will be on the same port but can be proxied to different upstreams.
TLS (SNI) Multiplexing on HTTPS Port
There is an application with the hostname app1.example.com which should not be handled by the Handlers of the Reverse Proxy. The TLS traffic of this application should be routed directly to an upstream destination without TLS termination.
- Go to
- Press + to create a new Layer4 Route
| Options | Values | 
|---|---|
| Domain: | app1.example.com | 
| Matchers: | TLS (SNI) | 
| Upstream Domain: | 192.168.1.1 | 
| Upstream Port: | 8443 | 
- Press Save and Apply
Caddy listens on the default HTTP and HTTPS ports. All traffic it receives on these or any other listening ports, gets passed to the listener_wrapper. Inside this wrapper, the traffic can be inspected on Layer 7, and routing decisions can be made.
With the matcher TLS (SNI), the Client Hello of the TLS traffic is analyzed. When the Client Hello includes app1.example.com, the traffic will be matched by the new Layer4 Route. The raw TLS traffic will be streamed to the chosen upstream socket.
Any other traffic that is not matched by this Layer4 Route will be routed to the Handlers, where the configured Domains and Subdomains can receive and reverse proxy it.
Tip
If there should be TLS termination, configure a domain in with a certificate installed for the same SNI that should match in this route. Check Terminate TLS in the route, the certificate will be automatically matched.
Note
When Auto HTTPS is enabled, all clients will be permanently redirected to HTTPS automatically. If that should not happen, set it to Disable Redirects.
Inverted TLS (SNI) Multiplexing on HTTPS Port
Inverting the TLS (SNI) matcher can route all unmatched traffic, for example to a hosting panel where the domains are not under administrative control and can change at any time. The domains matched by SNI will be routed to the Reverse Proxy.
Attention
If you create additional routes, e.g., for SSH, make sure to use the sequence number to generate them before this route.
- Go to
- Press + to create a new Layer4 Route
- Enable the advanced mode toggle
| Options | Values | 
|---|---|
| Sequence: | 100 | 
| Routing Type: | listener_wrappers | 
| Protocol: | TCP | 
| Matchers: | TLS (SNI) | 
| Domain: | *.example.com*.opnsense.com | 
| Invert Matchers: | X | 
| Upstream Domain: | 192.168.1.1192.168.1.2 | 
| Upstream Port: | 443 | 
| Fail Duration: | 10 | 
- Press Save and Apply
With the inverted TLS (SNI) matcher, the Client Hello of the TLS traffic is analyzed. When the Client Hello includes either of *.example.com or *.opnsense.com, the traffic will be sent to the default Handlers, where the configured Domains and Subdomains can receive and reverse proxy it.
All other traffic will be streamed to the chosen socket of Upstream Domain and Upstream Port. Since we chose multiple upstreams and a health check, two servers can load balance all requests. The load balancing is just an example, and not necessary for this matcher to work.
Tip
If there are domains inside *.example.com that should be routed to a different upstream, just create an additional TLS (SNI) matcher for them. Set the sequence to a lower number to match it before the inverted route.
Tip
Caddy supports the HA Proxy Protocol. If the Protocol Header should be added to the upstream, set the Proxy Protocol version to v1 or v2.
Proxy TCP/UDP on Layer 4
We have an application that should receive all TCP/UDP traffic directed at port 5060.
- Go to
- Press + to create a new Layer4 Route
- Enable the advanced mode toggle
| Options | Values | 
|---|---|
| Routing Type: | global | 
| Protocol: | TCP | 
| Local Port: | 5060 | 
| Matchers: | ANY | 
| Upstream Domain: | 192.168.1.1 | 
| Upstream Port: | 5060 | 
- Press Save and + to create another Layer4 Route
| Options | Values | 
|---|---|
| Routing Type: | global | 
| Protocol: | UDP | 
| Local Port: | 5060 | 
| Matchers: | ANY | 
| Upstream Domain: | 192.168.1.1 | 
| Upstream Port: | 5060 | 
- Press Save and Apply
DNS and Wireguard Multiplexing
We have a DNS server that hosts one of our DNS zones. We want to allow Wireguard on the same port as DNS, but only from a certain remote ip range.
Note
The sequence is optional, but it can influence the processing order of created rules.
- Go to
- Press + to create a new Layer4 Route
- Enable the advanced mode toggle
| Options | Values | 
|---|---|
| Sequence: | 100 | 
| Routing Type: | global | 
| Protocol: | UDP | 
| Local Port: | 53 | 
| Matchers: | DNS | 
| Upstream Domain: | 192.168.1.1 | 
| Upstream Port: | 53 | 
- Press Save and + to create another Layer4 Route
| Options | Values | 
|---|---|
| Sequence: | 101 | 
| Routing Type: | global | 
| Protocol: | UDP | 
| Local Port: | 53 | 
| Matchers: | Wireguard | 
| Upstream Domain: | 172.16.1.1 | 
| Upstream Port: | 51820 | 
| Remote IP: | 203.0.113.0/24 | 
- Press Save and Apply
All of these Layer 7 routes will be automatically grouped under port UDP/53 in the chosen sequence order.
Caddy: Troubleshooting
FAQ
- Cloudflare is not required to get automatic certificates.
- You can use the os-acme-client plugin to generate wildcard certificates. Set up an automation in the ACME client that reloads Caddy (do not restart it).
- Destination NAT (Port Forward), NAT Reflection, Split Horizon DNS or DNS Overrides in Unbound are not required. Only create Firewall rules that allow traffic to the default ports of Caddy.
- Even though internal clients will use the external IP address to access the reverse proxied services, the traffic will not pass over the internet. It will stay inside the OPNsense. Only in rare cases where there is multi WAN, the traffic can be routed from one WAN interface to the other over the internet, due to reply-to settings.
- Firewall rules to allow Caddy to reach internal services are not required. OPNsense has a default rule that allows all traffic originating from itself to be allowed.
- ACME clients on reverse proxied upstream destinations will not be able to issue certificates. Caddy intercepts/.well-known/acme-challenge . This can be solved by using the HTTP-01 Challenge Redirection option in the advanced mode of domains. Please check the tutorial section for an example.
- When using Caddy with IPv6, the best choice is to have a GUA (Global Unicast Address) on the WAN interface, since otherwise the TLS-ALPN-01 challenge might fail.
- Let’s Encrypt or ZeroSSL can not be explicitly chosen. Caddy automatically issues one of these options, determined by speed and availability. These certificates can be found in/var/db/caddy/data/caddy/certificates .
- When an Upstream Destination only supports TLS connections, yet does not offer a valid certificate, enableTLS Insecure Skip Verify in a Handler to mitigate connection problems.
