---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-9fe5a217
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-9fe5a217"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-9fe5a217.md
source_anchor: ""
source_lines: [1, 51]
sha256: f09d4451d4b50a16454cdd30f29ccf7c485df438721bba923d2cf2107d6af340
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-9fe5a217

IPv6 Support on MX Security & SD-WAN Platforms - Security
Security
Firewall Rules
L3 Firewall
The MX, by default, allows Outbound IPv6 connection from the LAN to the Internet and its returning traffic. Inbound traffic originating from the Internet without an existing flow or a matching allowed Inbound firewall rule will be dropped. L3 Firewall rules can be leveraged in order to allow or block IPv6 traffic to and from desired source and destinations via the Inbound, Outbound and Site-to-site VPN firewalls rules.
Prerequisite: In order to view and configure IPv6 firewall rules, the organization must leverage the new Policy Object feature. Please review the Network Objects Configuration Guide KB for more details on enrolling into Policy Objects.
Note:
- 
    At this time, Network Objects and Object groups are not yet supported for IPv6. IPv6 Prefixes and IPs must be configured directly on the firewall rules.
- 
    It is not recommended to have an Allow Any Any rule on the Inbound firewall rule as this greatly reduces the security of the network. It is also strongly recommended that you specify source and destination addresses to limit the exposure when allowing traffic inbound from the Internet to the LAN.
- 
    Like with IPv4, you can configure multiple IPv6 Prefixes and IPv6 Addresses in the Source and Destination boxes.
- 
    You can configure Dual Stack (IPv4 and IPv6) within the same so long as the source and destination contain both IPv4 and IPv6 or Any.
Layer 3 Inbound Firewall Rules
By default, the MX will deny all IPv6 traffic sourced from the Internet without a matching firewall rule or existing flow to allow the traffic. Inbound rules can be used to block or allow access to traffic originating from the Internet destined to a device on the MX LAN.
Configuration:
- 
    Go to Security & SD-WAN and select the Firewall page
- 
    Click Add New button in the Inbound rules section and create the desired L3 rule
- 
    Click on the Finish editing button and Save the page to commit the changes
Layer 3 Outbound Firewall Rules
By default, the MX will allow all IPv6 traffic sourced from the LAN side between VLANs and out to the Internet. Outbound rules can be used to block or allow traffic from the LAN to the Internet or between different local VLANs.
Configuration:
- 
    Go to Security & SD-WAN and select the Firewall page
- 
    Click Add New button in the Outbound rules section and create the desired L3 rule
- 
    Move the rule to the desired order, click on the Finish editing button
- 
    Save the changes to commit them to the MX
Layer 3 Cellular Failover Rules and Inbound Cellular Failover Rules
Cellular failover rules are leveraged by the MX when the Primary and Secondary (If installed) WAN uplink connections fail or are unavailable. These are configured the same way the L3 Inbound and Outbound rules, however the Inbound Cellular failover rules only support IPv6.
Organization-wide Layer 3 Site-to-site VPN Outbound Firewall Rules
Similar to the L3 Outbound rules, the Site-to-site outbound firewall allows you to configure IPv6 and IPv4 L3 firewall rules to allow or block traffic between VPN MX.
Configuration:
- 
    Go to Security & SD-WAN and select the Site-to-site VPN page
- 
    Scroll to the Site-to-site outbound firewall rules section and configure the desired rule
- 
    Click on the Finish editing button and Save the rule to commit the changes
Advanced Malware Protection (AMP)
The MX Security Appliance will block HTTP-based file downloads based on the disposition received from the AMP cloud. For more details, refer to the AMP documentation.
Threat Grid
Threat Grid analyzes suspicious files against behavioral indicators and a malware knowledge base sourced from around the world to provide industry-leading accuracy and context-rich threat analytics. For more details, refer to the Threat Grid documentation.
Refer to the main KB: IPv6 Support on MX Security & SD-WAN Platforms [Core Fundamentals]
