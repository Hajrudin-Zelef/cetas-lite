---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9-4
title: "blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9.md
source_anchor: ""
source_lines: [407, 598]
sha256: 384fb2084697f6fc9d0a767bd6c0130ed8ec3cede9d7fdc3dbbb0366eecfe98d
---

# blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9

| Members | VLAN10_MANAGEVLAN20_VPNVLAN30_CLEAR | 
Aliases
We define a few reusable aliases that help us condense our firewall rules. Some of them might become hard to maintain as they grow, in which case you might want to consider nesting aliases.
Navigate to Firewall → Aliases and create the following aliases.
Selective Routing Addresses
Services like banks might object to traffic originating from known VPN endpoints. We selectively route traffic from the VPN VLAN through the default WAN gateway.
| Name | SELECTIVE_ROUTING | 
| Type | Host(s) | 
| Description | External hosts reachable from IG_OUT_VPN networks through WAN | 
If you’re having issues with a service not working due to VPN, add the hostname
to this alias, e.g., netflix.com.
Admin / Anti-lockout Ports
| Name | PORTS_ANTI_LOCKOUT | 
| Type | Port(s) | 
| Content | 443 (Web GUI)22 (SSH) | 
| Description | OPNsense admin ports | 
Ports Allowed To Communicate Between VLANs
Allowed ports for intranet traffic. Amend the list depending on your needs.
| Name | PORTS_OUT_LAN | 
| Type | Port(s) | 
| Description | Ports allowed for intranet | 
Content:
- 53 DNS
- 5353:5354 mDNS
- 123 NTP
- 21 FTP
- 22 SSH
- 161 SNMP
- 80 HTTP
- 8080 : HTTP alt / UniFi device and application communication
- 443 HTTPS
- 8443 HTTPS alt / UniFi application GUI/API as seen in a web browser
- 8880 UniFi HTTP portal redirection
- 10001 UniFi device discovery
- 5001 iPerf
- 623 IPMI
- 5900 VNC
- 3389 RDP
- 49152:65535 ephemeral ports
Ports Allowed to Communicate with the Internet
Allow ports for egress internet traffic. Amend the list depending on your needs.
| Name | PORTS_OUT_WAN | 
| Type | Port(s) | 
| Description | Ports allowed for internet | 
Content:
- 21 FTP
- 22 SSH
- 80 HTTP
- 8080 HTTP alt
- 443 HTTPS
- 8443 HTTPS alt
- 465 SMTPS
- 587 : SMTPS
- 993 : IMAPS
- 49152:65535 ephemeral ports
A Fair Warning about Egress Filtering
As you add applications, you will be constantly amending the PORTS_OUT_WAN
list. Depending on the application, the required ports may be poorly documented,
so you’ll have to figure them out by inspecting the firewall logs. As other
users have mentioned in the comments, blocking all egress traffic for all VLANs
by default is probably not worth the hassle. Personally, I’ve given up on egress
filtering altogether because of the administrative overhead that comes with it.
It is useful for high-security VLANs connecting devices such as cash registers in a retail store. Another example is VLANs with many untrusted IoT devices that have noisy telemetry. Putting them into a VLAN with egress filtering prevents them from “calling home”.
If you create the alias ALL_PORTS = 1:65535 and add it to the Content
field of the PORTS_OUT_WAN alias, you can disable all egress filtering with
the option of re-enabling it again later.
NAT
Network Address Translation (NAT) is required to translate private to public IP addresses. We have the following requirements.
- Translate IG_OUT_WAN andIG_OUT_VPN network addresses to theWAN address
range. TranslatingIG_OUT_VPN toWAN allows selective routing.
- Translate IG_OUT_VPN network addresses to theWAN_VPN0 address range.
Navigate to Firewall → NAT → Outbound.
Select Manual outbound NAT rule generation and add the following rules.
IG_OUT_WAN to WAN
| Interface | WAN | 
| Source address | IG_OUT_WAN net | 
| Description | IG_OUT_WAN to WAN | 
IG_OUT_VPN to WAN
| Interface | WAN | 
| Source address | IG_OUT_VPN net | 
| Description | IG_OUT_VPN to WAN | 
IG_OUT_VPN to WAN_VPN0
| Interface | WAN_VPN0 | 
| Source address | IG_OUT_VPN net | 
| Description | IG_OUT_VPN to WAN_VPN0 | 
IG_OUT_VPN to WAN_VPN1
| Interface | WAN_VPN1 | 
| Source address | IG_OUT_VPN net | 
| Description | IG_OUT_VPN to WAN_VPN1 | 
Rules
Navigate to Firewall → Rules.
Anti-Lockout
Before adding any other rules, we add the anti-lockout ones on the
VLAN10_MANAGE and LAN networks, so we can’t lock ourselves out. 😅
Select Floating and add the following rule.
| Action | Pass | 
| Interface | LANVLAN10_MANAGE | 
| Protocol | TCP/UDP | 
| Source | any | 
| Destination | This Firewall | 
| Destination port range | PORTS_ANTI_LOCKOUT | 
| Description | Anti-lockout | 
This Firewall is a pre-defined alias representing all interface addresses of
OPNsense.
Allow Intranet Pings
We allow ICMP pings for the entire local network. Pings are maliciously abusable, so you may want to put stricter rules into place if required.
Select IG_LOCAL and add the following rule.
| Action | Pass | 
| Interface | IG_LOCAL | 
| TCP/IP Version | IPv4 | 
| Protocol | ICMP | 
| ICMP type | Echo Request | 
| Source | IG_LOCAL net | 
| Description | Allow intranet pings | 
Reject Intranet Traffic By Default
By default, we reject traffic on local interfaces instead of blocking it. Block drops packets silently. Reject returns a “friendly” response to the sender. To be able to override this rule, unchecking Quick is crucial! To use Firewall Logs to review blocked ports and amend our port list alias if necessary, we enable logging on this rule.
Select IG_LOCAL and add the following rule.
| Action | Reject | 
| Quick | unchecked | 
| Interface | IG_LOCAL | 
| TCP/IP Version | IPv4+IPv6 | 
| Protocol | any | 
| Source | IG_LOCAL net | 
| Destination | IG_LOCAL net | 
| Log | checked | 
| Description | Reject intranet traffic by default | 
Allow Intranet Traffic
We only allow intranet traffic on the ports defined in the PORTS_OUT_LAN
alias. We’ll override this rule for the VLAN40_GUEST network later, so we must
uncheck the Quick option again. For the Management network, you might want
to consider stricter rules, as well.
Select IG_LOCAL and add the following rule.
| Action | Pass | 
| Quick | unchecked | 
| Interface | IG_LOCAL | 
| Protocol | TCP/UDP | 
| Source | IG_LOCAL net | 
| Destination | IG_LOCAL net | 
| Destination port range | PORTS_OUT_LAN | 
| Description | Allow intranet traffic | 
Allow Internet Traffic
We allow internet traffic on PORTS_OUT_WAN for IG_OUT_WAN networks.
Select IG_OUT_WAN and add the following rule.
| Action | Pass | 
| Quick | unchecked | 
| Interface | IG_OUT_WAN | 
| Protocol | TCP/UDP | 
| Source | IG_OUT_WAN net | 
| Destination / Invert | checked | 
| Destination | IG_LOCAL net | 
| Destination port range | PORTS_OUT_WAN | 
| Description | Allow internet traffic through WAN | 
We later want to enable unrestricted internet access on the Guest network, so make sure to uncheck the Quick option!
Next, we allow internet traffic on PORTS_OUT_WAN for the IG_OUT_VPN
networks.
Select IG_OUT_VPN and add the following rules to configure selective routing.
| Action | Pass | 
| Interface | IG_OUT_VPN | 
| Protocol | TCP/UDP | 
| Source | IG_OUT_VPN net | 
| Destination | SELECTIVE_ROUTING | 
| Destination port range | PORTS_OUT_WAN | 
| Description | Allow selected internet traffic through WAN | 
| Action | Pass | 
| Protocol | TCP/UDP | 
| Source | IG_OUT_VPN net | 
| Destination / Invert | checked | 
| Destination | IG_LOCAL net | 
| Destination port range | PORTS_OUT_WAN | 
| Description | Allow internet traffic through WAN_VPN0 | 
| Gateway | WAN_VPN_GROUP | 
Restrict Guest Network
Select VLAN40_GUEST and add the following rules.
To block Web GUI and SSH access from the Guest network, we block traffic to any
OPNsense interface on the PORTS_ANTI_LOCKOUT ports. We enable logging for this
rule to be able to see if any guests try to access OPNsense.
| Action | Block | 
| Interface | VLAN40_GUEST | 
| Protocol | TCP/UDP | 
| Source | VLAN40_GUEST net | 
| Destination | This Firewall | 
| Destination port range | PORTS_ANTI_LOCKOUT | 
| Log | checked | 
| Description | Block admin ports | 
We block access to other local networks and also enable logging for the rule.
| Action | Block | 
| Interface | VLAN40_GUEST | 
| Protocol | TCP/UDP | 
| Source | VLAN40_GUEST net | 
| Destination | IG_LOCAL net | 
| Log | checked | 
| Description | Block traffic to local networks | 
