---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1-4
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2022-03"]
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1.md
source_anchor: ""
source_lines: [128, 191]
sha256: 5010e99e7abe91143f153f2454f4e8c0a8ca16e4095df2cda0c42a804d52ec73
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1

Number of Supported Sessions per MX Model
Below is the number of sessions allowed per MX model. When the limit is reached, new sessions will not be formed.
| Model | Max sessions | 
|---|---|
| MX450 | 1,500 | 
| MX250 | 1,000 | 
| MX105 | 750 | 
| MX100 | 250 | 
| MX95 | 500 | 
| MX85 | 250 | 
| MX84 | 100 | 
| MX75 | 250 | 
| MX67/68 | 100 | 
| MX64/65 | 50 | 
| Z3 | 5 | 
| vMX S/M/L | 100/250/500 | 
| vMX100 | 250 | 
| MX600 | 1000 | 
| MX400 | 750 | 
FAQ
- 
    Who signs the Meraki facilitated publicly trusted certificates? 
A publicly trusted Certificate Authority.
- 
    Can I use my own hostname or publicly trusted certificate on the MX as a server certificate? 
Yes, see Custom hostname certificates
- 
    How will AnyConnect be licensed on the Meraki MX? 
See AnyConnect licensing on the MX
- 
    Which MX/vMX models support AnyConnect? 
See caveats section
- 
    Can I use AnyConnect profiles? 
Yes, see the AnyConnect Profiles section. Only VPN profiles can be pushed via the MX. Others profiles, like Umbrella profiles, etc will not be pushed via the MX.
- 
    Can I configure different split-tunnel rules/VLANs/IP address pools for different sets of users? 
No, not at the moment. However, you can use group policies when authenticating with RADIUS to apply access policies to a user or groups of users on authentication.
- 
    Can I do certificate-based authentication? 
Yes, as a combination with username and password. See the certificate-based authentication section. Certificate-only authentication is currently in beta see Certificate-only authentication for more details.
- 
    Where can I download the AnyConnect client? 
On the AnyConnect Settings page on dashboard in the Client Connection section or on cisco.com
- 
    What are the current caveats/known issues with the AnyConnect feature & firmware? 
See caveats section
- 
    Which features are supported? Any plans to support Umbrella, posture scan, 802.1x, etc? 
VPN Only. Other AnyConnect modules that do not require additional server support can be used as well, for example, DART, Umbrella. This module must be deployed and configured separately as the MX does not support web launch, client software deployment, or update at this time. See AnyConnect on ASA vs. MX for more details.
- 
    Can I use IKEv2 on AnyConnect to connect to the MX Appliance? 
No, AnyConnect only supports TLS 1.2 ( TLS 1.3 on 19.1+ firmware) and DTLS 1.2 connections on the MX.
- 
    Can I run L2TP/IPsec Client VPN and AnyConnect VPN simultaneously on the MX? 
Yes.
- 
    Can I connect to the inside interface of the MX with AnyConnect? For example, connect to the MX from the LAN side? 
No, only inbound connections on the WAN side are supported at this time.
- 
    When will AnyConnect General Availability (GA)? 
AnyConnect GA'd on the MX 16.16+ firmware released in March 2022.
- Does Cisco AnyConnect custom server certificate support the use of a wildcard certificate? Example, XXX.mycompany.com vs wildcard certificate *.mycompany.com
Wildcard certificates are not supported.
