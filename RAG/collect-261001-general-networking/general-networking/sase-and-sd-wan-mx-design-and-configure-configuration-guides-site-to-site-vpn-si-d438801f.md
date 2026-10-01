---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-si-d438801f
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-si-d438801f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters", "training"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-si-d438801f.md
source_anchor: ""
source_lines: [1, 34]
sha256: af1869d8adde89d43af1f866620aae5f479b9e3066925cb19487bf6d6615feb2
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-si-d438801f

Site-to-site VPN Firewall Rule Behavior
Click 日本語 for Japanese
Overview
Administrators have the ability to add firewall rules to restrict the traffic flow through the VPN tunnel for a Cisco Meraki MX Security Appliance. Similar to other Meraki firewall options, this firewall is stateful and will only block traffic if it does not match an existing flow.
These firewall rules will apply to all MX networks in the organization that participate in site-to-site VPN (both AutoVPN and IPsec VPN).
Learn more with these free online training courses on the Meraki Learning Hub:
Creating Firewall Rules
To create a firewall rule, follow the steps below.
- 
    Navigate to Security & SD-WAN > Configure > Site-to-site VPN.
- 
    In the Site-to-site outbound firewall section, select Add new next to the search bar.
- 
    Fill in the desired parameters for the rule
- 
    Select Save changes.
- FQDN, Wildcard FQDN, and Network Groups with FQDN objects cannot be applied to Site-to-Site VPN Outbound Firewall Rules.
Considerations for VPN Firewall Rules
When configuring VPN Firewall rules, it is important to remember that traffic should be stopped as close to the originating client device as possible. This cuts down on traffic over the VPN tunnel and will result in the best network performance. Because of this, site-to-site firewall rules are applied only to outgoing traffic. As such, the MX cannot block VPN traffic initiated by IPsec VPN peers.
Note: Outbound IPsec VPN peer tunnel and related BGP control traffic can also be blocked by these rules.
The image below demonstrates a misconfigured site-to-site firewall rule. Traffic initiated from "10.0.1.0/24 to 10.0.2.0/24" or vice-versa will be allowed.
Site-to-site firewall rules only apply to outbound traffic. This rule will never be applied as the source subnet is not a LAN subnet on the MX:
The following image demonstrates a site-to-site firewall rule that is applied correctly. Traffic from the 10.0.1.0/24 subnet will not be able to reach 10.0.2.0/24 subnet since the 10.0.1.0/24 subnet is a LAN subnet on the MX.
However, traffic from "10.0.2.0/24 to 10.0.1.0/24" (which is initiated from 3'rd party VPN side), will be allowed as the traffic will be considered an inbound traffic and site-to-site firewall rules cannot block it. Return traffic will also be allowed since MX is a stateful firewall.
When traffic passing through the MX matches a site-to-site VPN route, VPN firewall rules are applied in descending order (top-to-bottom). VPN traffic to both AutoVPN and IPsec VPN peers is only subject to the site-to-site firewall rules and is never subject to global Layer 3 firewall rules.
Layer 7 Firewall Rules
Unlike Layer 3 firewall rules, Layer 7 firewall rules configured on the Security & SD-WAN > Configure > Firewall page will still apply locally to client traffic destined across both AutoVPN and IPsec VPN peers.
Note - Site-to-Site Firewall Rules Behavior when Group Policy is Configured
- 
    If Site to Site Outbound Firewall Rule allows and Group Policy L3 denies, traffic will be denied.
- 
    If Site to Site Outbound Firewall Rule denies and Group Policy L3 allows, traffic will be denied.
- 
    If Site to Site Outbound Firewall Rule denies and Group Policy whitelisted preset is configured, traffic will be denied.
