---
id: collect-261001-general-networking/general-networking/manual-captiveportal-html-486dacb4-2
title: "manual-captiveportal-html-486dacb4"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-captiveportal-html-486dacb4.md
source_anchor: ""
source_lines: [86, 168]
sha256: a472e28a2bb5bdaef3c4042a5ef3cedc0fa8517e4d252aeb3930b61582b13c42
---

# manual-captiveportal-html-486dacb4

| Disable firewall rules | If this option is set, no automatic firewall rules for portal redirection and traffic blocking will be generated. This option allows you to override the default portal behavior for advanced use cases, such as redirections for DNS on a non-standard port. See Captive Portal Firewall rules for an overview of required firewall rules. | 
| Authenticate using | Select an authenticator specified in | 
| Always send accounting requests | [RADIUS only] This will make the captive portal always send accounting requests, rather than just when there is a need for accounting (e.g. when there is a daily session limit). | 
| Enforce local group | Restrict access to users in the selected (local)group, to validate group membership, see | 
| Idle timeout (minutes) | Clients will be disconnected after this amount of inactivity. They may log in again immediately, though. | 
| Hard timeout (minutes) | Clients will be disconnected after this amount of time, regardless of activity. They may log in again immediately, though. | 
| Concurrent user logins | If this option is set, users can login on multiple machines at once. If disabled subsequent logins will cause machines previously logged in with the same username to be disconnected. | 
| SSL certificate | Certificate to use on the captive portal login system. Leave empty for HTTP only. | 
| Hostname | Hostname (of this machine) to redirect login page to, leave blank to use this interface IP address, otherwise make sure the client can access DNS to resolve this location. When using a SSL certificate, make sure both this name and the cert name are equal. | 
| Allowed addresses | Avoid authentication for addresses and subnets specified in this list | 
| Allowed MAC addresses | Avoid authentication for MAC addresses specified in this list | 
| Extended pre auth data | Offer extended data to the login template before authentication (mac addresses for upstream use). | 
| Custom template | Template to use for the login page, specified in the templates tab. | 
The file offered is a standard zip file, which can be unpacked locally and modified to your needs, the new contents can be saved into a new zip file and uploaded in a new template ()
Sessions
Basic real time reporting is integrated using the sessions menu, this shows the following information for each zone.
- Live top IP bandwidth usage
- Active Sessions
- Time left on Vouchers
Vouchers
Here you can create new vouchers for all voucher servers configured in
Examples
Migration notes & technical details
Important
Starting from OPNsense Community edition 25.1.4 or Business edition 25.10, the underlying captive portal implementation has moved from IPFW to PF. While in most cases this has no practical impact, a more detailed description of what this means, as well as any incompatibilities are described here.
Previously, our Captive Portal implementation was split up into two components, IPFW and PF. Packets would enter an interface, where IPFW was the first to handle them and redirect traffic to the portal. After authentication, a rule to allow all traffic to and from this client would be inserted. Any traffic passed by IPFW would then be handled by PF. The reverse was true for outbound traffic. This process has been simplified by moving the redirection and accounting logic to PF.
This has multiple benefits:
- The generated rules are now visible in the WebGUI and are logged by default, easing troubleshooting.
- The list of clients that have been authenticated is now an alias and is visible, as well as usable in rules. The alias is called __captiveportal_zone_<zoneid> .
- Unless custom shaper rules are used, IPFW does not need to be loaded anymore, significantly reducing implementation complexity.
The following are the only functional/behavioral changes:
- If you have forwarding rules defined on your captive portal zone that redirect a client for services other than HTTP/HTTPS, the “Filter Rule Association” option on this rule must be set to “Pass” so that this traffic is allowed after redirection; otherwise, this traffic will hit the default captive portal block rule. An example of such a scenario would be client DNS traffic redirected to a DNS service running on localhost.
- The “Allow Inbound” option has been dropped. This option only affected IPFW rules and controlled whether traffic from another network going to the captive portal zone would be allowed. This behavior is now determined by the ruleset of the network where the traffic is originating from. As an example, if some network on a non-captive interface is allowed everywhere according to the ruleset, this traffic is also allowed into the captive portal zone. If this is not desired, an explicit block rule must be configured on said interface.
- Unless you are overriding the (newly) automatically generated firewall rules, you don’t need an explicit pass rule for DNS (port 53) on the firewall, nor an allow rule for the captive portal zones (ports 8000-10000) anymore. These are now installed by default.
Captive Portal Firewall rules
When running a default Captive Portal zone, the necessary rules for redirection are automatically installed in the zone. These rules have a higher priority than any user-defined rules. Therefore, to allow for flexibility, these rules may be overridden using the “Disable firewall rules” option in the zone administration. The automatically generated rules are listed here so they may be recreated for proper portal functionality.
Redirect traffic to the zone webserver
All HTTP traffic going to port 80 is redirected to localhost port 9000 + <zone id>.
All HTTPS traffic going to the firewall on port 443 is redirected to localhost port 8000 + <zone id>.
| Type | Destination NAT (Port Forward) | 
| Interface | <Zone interface> | 
| Protocol | TCP | 
| Source Invert | Yes | 
| Source | __captiveportal_zone_<zone id> | 
| Destination Invert | Yes | 
| Destination | __captiveportal_zone_<zone id> | 
| Destination port range | 80 | 
| Redirect Target IP | 127.0.0.1 | 
| Redirect Target Port | 9000 + <zone id> | 
| NAT Reflection | Disable (advanced) | 
| Type | Destination NAT (Port Forward) | 
| Interface | <Zone interface> | 
| Protocol | TCP | 
| Source Invert | Yes | 
| Source | __captiveportal_zone_<zone id> | 
| Destination | This Firewall | 
| Destination port range | 443 | 
| Redirect Target IP | 127.0.0.1 | 
| Redirect Target Port | 8000 + <zone id> | 
| NAT Reflection | Disable | 
For IPv6, the same rules as above must be created, but the Destination must be set to the “<Zone interface> address”.
The destination for HTTP traffic is the inverted zone alias, so that traffic from unauthenticated clients going to authenticated or explicitly allowed clients/servers (allowed addresses in the zone administration) is not redirected. This is useful if unauthenticated clients should be able to access servers in the same zone.
Attention
If you use OIDC for authentication, the HTTPS requests would also be redirected before authentication is possible. To solve this, create an additional “No RDR (NOT)” rule before the other NAT rules with the identity provider IP addresses as destination.
| Type | Destination NAT (Port Forward) | 
| No RDR (NOT) | Yes | 
| Interface | <Zone interface> | 
| Protocol | TCP | 
| Source | any | 
| Destination | identity_provider_ip_addresses | 
| Destination port range | 443 | 
Allow DNS
In order to allow the client to resolve at least the OPNsense hostname, DNS must be allowed.
| Type | Firewall rule | 
| Action | Pass | 
| Interface | <Zone interface> | 
| Protocol | TCP/UDP | 
| Direction | In | 
| Source | <Zone net> | 
| Destination | This Firewall | 
| Destination port range | DNS/DNS | 
We define “This Firewall” as the destination since the default DNS service, Unbound, may return multiple IP addresses identifying the firewall.
Allow access to Captive Portal
