---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-12568927589143-content-and-domain-filtering-in-unifi-6e24d21e
title: "hc-en-us-articles-12568927589143-content-and-domain-filtering-in-unifi-6e24d21e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["cyber"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-12568927589143-content-and-domain-filtering-in-unifi-6e24d21e.md
source_anchor: ""
source_lines: [1, 65]
sha256: 9718c8b468c45b351bc6546802972588105b4be2a4162ecd0276d6a9f9b0c5b1
---

# hc-en-us-articles-12568927589143-content-and-domain-filtering-in-unifi-6e24d21e

Content and Domain Filtering in UniFi
UniFi Gateways offer content filtering to block access to malicious, explicit, and inappropriate websites across your network. This is a DNS-level feature that can be flexibly applied to specific VLANs or client devices.
For a full overview of UniFi’s Traffic and Policy Management capabilities, see here.
For a full overview of UniFi's Network and Cyber Security capabilities, see here.
Configuring Content Filtering
- Navigate to Content Filtering: Settings > CyberSecure > Content Filter.
- Create Name: Assign a name to the policy.
- Select Source: Apply the policy to specific Networks and/or client Devices.
- (Optional) Enable Ad Blocking: Uses UniFi's internal database to block domains known to be associated with ads.
- (Optional) Enable Safe Search: Take advantage of the Safe Search feature integrated with Google, Bing, and/or YouTube, designed to block explicit content.
- 
Select Filtering Policy
  - Off
  - Basic: A pre-defined configuration used to block content associated with malicious, explicit, or adult domains.
  - Enhanced: An advanced filtering solution powered by Cloudflare, offering 100+ category-based filters managed by their industry-leading security team. Ideal for professional and enterprise environments requiring strict policy enforcement or compliance-driven controls. For more information, see CyberSecure Enhanced by Proofpoint and Cloudflare.
- 
(Optional) Create Allowlist and Blocklist
  - Allowlist: Use this to whitelist a domain that would otherwise be blocked by an above filtering policy.
  - Blocklist: Use this to block additional domains that are not already blocked by an above filtering policy.
- (Optional) Set a Schedule: Set up a Schedule for when the policy should be applied.
How It Works
Content filtering works by redirecting DNS traffic to the UniFi Gateway for inspection. Requests are compared against internal blocklists (or enhanced category filters), and access is denied before the browser can connect to the site. This mechanism allows filtering without any client-side configuration and is updated automatically.
Note: DNS Shield cannot be used with Content Filtering on UniFi Gateways running version 4.2 and earlier.
Support for Local DNS Resolution
Because content filtering relies on DNS redirection, requests to internal or manually configured DNS servers (such as those used by Active Directory) may fail unless routing is explicitly handled.
The table below summarizes DNS and SNAT requirements based on your topology:
| Scenario | Requires DNS Forwarding | Requires SNAT | 
| DC on a local VLAN | ✅ | ❌ | 
| DC via Site Magic SD-WAN | ✅ | ❌ | 
| DC via Route-Based IPsec VPN (with VTI address) | ✅ | ❌ | 
| DC via Route-Based IPsec VPN (no VTI address) | ✅ | ✅ | 
| DC via Policy-Based IPsec VPN | ✅ | ✅ | 
To restore proper resolution of internal DNS zones, follow the steps below.
Creating a DNS Forwarding Rule
- Navigate to DNS forwarding rules.
  - Network 9.3: Settings > Policy Engine > DNS.
  - Network 9.4: Settings > Policy Table.
- Click:
  - In Network 9.3: Click Create DNS.
  - In Network 9.4: Click Create New Policy > Select "DNS".
- Add a new domain entry:
  - Domain: Your internal domain (e.g., corp.local)
  - Target IP: IP address of your domain controller
Note: Enter only the base domain (e.g., example.com). Do not add SRV records or specific hostnames such as _ldap._tcp.dc.example.com or dc1.example.com. Any DNS query matching this domain (including subdomains) will be forwarded.
Creating a Source NAT (SNAT) Rule
Required only for policy-based IPsec or route-based VPNs without VTI addresses.
- Navigate to NAT section.
  - Network 9.3: Settings > Policy Engine > NAT.
  - Network 9.4: Settings > Policy Table.
- Click:
  - In Network 9.3: Click Create New > Src. NAT
  - In Network 9.4 Create New Policy > Select "NAT" > Src. NAT
- Configure the rule:
  - Rule Type: Source NAT
  - Description: e.g., SNAT for DNS to remote DC
  - Protocol: UDP and TCP
  - Source IP: Leave blank
  - Destination IP: IP of the domain controller
  - Destination Port: Port 53 (for TCP/UDP)
  - Translated IP: IP of the Gateway’s LAN interface
  - 
Interface:
    - For Policy based VPN - use the WAN interface where IPsec is configured
    - For Route-based VPN - use the Site-to-Site VPN interface
This ensures that return packets from the DNS server can properly reach the Gateway.
