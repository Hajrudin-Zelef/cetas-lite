---
id: collect-261001-general-networking/general-networking/manual-captiveportal-html-486dacb4-1
title: "manual-captiveportal-html-486dacb4"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Microsoft"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-captiveportal-html-486dacb4.md
source_anchor: ""
source_lines: [1, 85]
sha256: 6cd2213cc49c84bf4f9d3f0e5101327b874ca03a0a4b587b23d6541a2ab38c61
---

# manual-captiveportal-html-486dacb4

Captive portal & GuestNET
A Captive Portal allows you to force authentication, or redirection to a click through page for network access. This is commonly used on hotspot networks, but is also widely used in corporate networks for an additional layer of security on wireless or Internet access.
Overview and terminology
Typical Applications
- Guest Network
- Hotel & Camping Wi-Fi Access
- Bring Your Own Device (BYOD)
Template Management
OPNsense’s unique template manager makes setting up your own login page an easy task. At the same time it offers additional functionalities, such as:
- URL redirection
- Option for your own Pop-up
- Custom Splash page
Zone Management
Different zones can be setup on each interface or multiple interfaces can share one zone setup. Each Zone can use a different Captive Portal Template or share it with another zone.
Authentication
Secure authentication via HTTPS or splash-only portal with URL redirection to a given page Different sources can be used to authenticate a user in a zone:
- LDAP [Microsoft Active Directory]
- Radius, including accounting updates
- Local user manager
- Vouchers / Tickets
- No authentication (Splash Screen Only)
- Multiple (a combination of above)
Voucher Manager
OPNsense’s Captive Portal has an easy voucher creation system that exports the vouchers to a csv file for use with your favorite application. The export allows you to print vouchers by merging them with your Microsoft Word or LibreOffice template and create a good looking handout with your logo and company style.
Timeouts & Welcome Back
Connection can be terminated after the user has been idle for a certain amount of time (idle timeout) and/or force a disconnect when a number of minutes have passed even if the user is still active (hard timeout). In case a user reconnects within the idle timeout and/or a hard timeout, no login is required and the user can resume its active session.
Bandwidth Management
The Built-in traffic shaper can be utilized to:
- Share bandwidth evenly
- Give priority to protocols port numbers and/or IP addresses
See also: Traffic Shaping
Portal bypass
MAC addresses and IP addresses/network ranges can be white listed to bypass the portal.
Platform Integration
Through the integrated REST API the captive portal application can be integrated with other services. See: Use the API
IPv6 support
The OPNsense Captive Portal fully supports IPv6-only and dual-stack networks. To facilitate this, the [Roaming] option is available and set by default in each zone. The IPv6 protocol commonly uses multiple IPv6 addresses on the same network interface of a client. These can be Link-local addressess, GUAs, ULAs, temporary/ privacy addresses or stable addresses.
Roaming allows the portal to register any IP alias a client is using, including IPv4 addresses. Once a client is connected, these IP addresses are collected in the background and are granted access.
Furthermore, the following criteria must be met for IPv6 to fully function:
- Hostwatch () must be enabled for the collection of IPv6 addresses to function.
- You must set a [Hostname] in the zone configuration and make sure a DNS record exists for this hostname pointing to the correct IPv6 address. If you’re using Unbound, DNS records for the configured interfaces can be synthesized with the [DNS64] option in .
Note
The background process collecting these IP addresses does this in a fixed interval. There may be a slight delay before all addresses are collected and granted access.
Modern Portal support
OPNsense implements the Captive Portal by redirecting all HTTP traffic to a local web server before authentication, hinting to the device that it is behind a portal. However, RFC 8910 introduces a new standardized method for networks to inform clients about the presence of a portal using DHCP. Furthermore, RFC 8908 describes an API standard implemented by the webserver pointed to by DHCP, where the client can fetch the current portal status. Apple has published a document going into more details.
Modern clients (especially iOS) moving towards this standardized API may experience redirection issues when connecting to a network only supporting forced redirection, which is often solved by utilizing this new standard instead.
To configure this, a few steps are required:
- You must install a valid, publicly trusted certificate on the Captive Portal zone. For example, you can use ACME client to automate this process. Doing so is best practice regardless of redirection method.
- The DHCPv4 server running in your Captive Portal zone must present option 114, of which the value must be set to the OPNsense webserver running the portal: https://<opnsense-hostname>/api/captiveportal/access/api .
Alternatively, a client can also be pointed to the redirected webserver directly:https://<opnsense-hostname>:<8000 + captive portal zone id>/api/captiveportal/access/api .
For example,https://opnsense.localdomain:8001/api/captiveportal/access/api for zone 1.
See the attention block below for more details.
To set this DHCP option, refer to the documentation for the respective DHCP server.
If a device in the captive portal zone supports this API, they will automatically use the DHCP option to determine that they are in a captive state. Keep in mind that forced redirection is still used for maximum compatibility. If you would like to excusively use this API standard instead, you can override the firewall rules for each zone and leave out the redirection rules, see rules.
Attention
Once a client has logged in through the portal, your firewall policies define what
this client can and cannot access. Unless you have configured the DHCP option to point
to the portal webserver directly by specifying the port, the
/api/captiveportal/access/api endpoint now points to the regular OPNsense WebGUI
on port 443, because authenticated clients are not redirected. This also means that
clients cannot determine captivity state anymore.
If you would like to restrict client access to this endpoint only, you must configure the proper forwarding rule as shown below. Note that this means clients cannot access the OPNsense WebGUI anymore, which is often desirable.
| Type | Destination NAT (Port Forward) | 
| Interface | <Zone interface> | 
| Version | IPv4+IPv6 | 
| Protocol | TCP | 
| Source | __captiveportal_zone_<zone id> | 
| Destination | This Firewall | 
| Destination port range | 443 | 
| Redirect Target IP | 127.0.0.1 | 
| Redirect Target Port | 8000 + <zone id> | 
| NAT Reflection | Disable (advanced) | 
| Firewall rule | Pass | 
Note
The OPNsense /api/captiveportal/access/api endpoint returns the following information:
- Whether a client is ‘captive’ at this point in time.
- The URL of the portal.
- If the client is authenticated and a hard timeout is set, how many seconds are remaining for this session.
Administration
The Administration menu offers access to zone configuration and template management.
When creating a zone, a couple of options are available which we will try to explain briefly in the grid below:
| Enabled | Enable the zone, which will install a network trap on the interfaces specified | 
| Zone number | Read-only sequence of the configured zone. This number is useful to determine the alias containing authenticated clients. For example, zone 0 will have an associated internal alias called __captiveportal_zone_0 . This alias can be inspected in . This zone id is also used if you are configuring the firewall rules yourself. | 
| Interfaces | Interfaces which should be guarded by this captive portal. | 
| Client Roaming | Allow a connecting client to use multiple IPs (bound to the same MAC) over the course of its session. This option is needed for maximum IPv6 compatibility and also affects IPv4 clients. | 
