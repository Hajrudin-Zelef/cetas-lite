---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-746f678a-1
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-746f678a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "training"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-746f678a.md
source_anchor: ""
source_lines: [1, 74]
sha256: 9eb6003e5579449277cdde07485c74378cb00bbc18356b578c452c588e4fbb05
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-746f678a

MX Firewall Settings
Click 日本語 for Japanese
This article provides an overview of what traditional and next-generation firewalls are, in addition to the configuration and capabilities of the MX Security & SD-WAN Appliance.
The firewall settings page in the Meraki Dashboard is accessible via Security & SD-WAN > Configure > Firewall. On this page you can configure Layer 3 and Layer 7 outbound firewall rules, publicly available WAN appliance services, port forwarding, 1:1 NAT mappings, and 1:Many NAT mappings.
For information regarding firewall rules for communication between Meraki devices and the Meraki cloud, please refer to the article on Firewall Rules for Cloud Connectivity.
Learn more with these free online training courses on the Meraki Learning Hub:
Comparing Next Generation Firewalls to Traditional Firewalls
When selecting a firewall for your network you will encounter a myriad of potential options. Some will be next-generation, some will be traditional. Let's take a look at the key differences below.
Features
Firewall selection typically begins with a required feature set. Traditional firewalls have a limited feature set while next-generation firewalls have an extended feature set. More details on the difference in features available are provided below.
Performance
Enabling additional features such as detailed client tracking, intrusion detection/prevention, and stateful inspection will impact performance. These features require additional resource utilization (memory and CPU) to perform the underlying packet scanning that enables these features. Comparing the performance of a next-generation firewall to a traditional firewall is not typically possible given the default feature set enabled on each.
Management
A traditional firewall is managed via a CLI or GUI that is hosted on the device. Next-generation firewalls typically have Cloud or Controller management capabilities.
What is a traditional firewall?
A traditional firewall will provide basic access control list (ACL) capabilities at Layers 3 & 4 in addition to Layer 3 routing and flow state tracking capabilities. These features are typically enabled by default, though they may not be configured out-of-the-box. Additional configuration capabilities are fairly limited, as many of the foundational features that enable Software Defined Routing (SD-Routing) are not available on traditional firewalls.
Standard Feature List
A standard firewall's default feature set will include functionality such as:
- Stateful or stateless flow tracking
- Standard firewalling capabilities such as:
    
  - Access Control Lists (ACLs) based on Layer 3 & Layer 4 information
- Network Address Translation (NAT) and Port Address Translation (PAT)
Extended Feature List
- Traffic Shaping
    
  - Rate limiting / policing
- Quality of Service
    
  - DSCP & CoS Honoring / Marking
- Virtual Private Network (VPN)
What is a Next Generation Firewall?
A next-generation firewall (NGFW) is a network security device that provides capabilities beyond a traditional, stateful firewall. While a traditional firewall typically provides stateful inspection of incoming and outgoing network traffic, a next-generation firewall includes additional features like application awareness and control, integrated intrusion prevention, and cloud-delivered threat intelligence. NGFW appliances will also include Software Defined Networking (SDN) capabilities which traditional firewalls cannot achieve without orchestration and use of additional integrations.
The standard feature set of a NGFW is a 'generation ahead' of a traditional firewall. These standard features are the foundation upon which Software Defined Networking (SDN) and threat protection implementations are built.
Standard Feature List
A next-generation firewall's default feature set will include functionality such as:
- Stateful or stateless flow tracking
- Standard firewalling capabilities such as:
    
  - Access Control Lists (ACLs) based on Layer 3 & Layer 4 information
- Network Address Translation (NAT) and Port Address Translation (PAT)
- Deep packet inspection (DPI)
- URL Filtering
- Traffic Shaping
    
  - Rate limiting / policing
- Quality of Service (QoS)
    
  - DSCP & CoS Honoring / Marking
- Software Defined WAN (SD-WAN) and SD-Routing decision capabilities
- Client and host tracking (Layer 3 - Layer 7)
Extended Feature List
Below are a few features which we would typically expect to find
- Stateful packet inspection
- Virtual Private Network (VPN)
- Intrusion Detection and Prevention (IPS)
- Malware Protection
Outbound Rules
Outbound rules are Layer 3 firewall ACLs that permit or deny traffic between local VLANs or from the LAN to the Internet. Rules can match protocol, source IP address or subnet, source port, destination IP address, subnet, FQDN, domain name, and destination port.
These rules do not apply to Auto VPN traffic. To configure firewall rules between VPN peers, use the Site-to-site VPN Firewall Settings.
To add a rule, click Add a rule and configure the following fields:
- Policy: Select whether matching traffic is allowed or denied.
- Rule description: Add an optional description or comment for the rule.
- Protocol: Select TCP, UDP, ICMPv4/ICMPv6, or Any.
- Source: Enter an IP address, CIDR subnet, comma-separated list, or Any. Source IPs and subnets must match subnets configured on the MX Addressing & VLANs page.
- Destination: Enter an IP address, CIDR subnet, FQDN, domain name, comma-separated list, or Any.
- Src Port / Dst Port: Enter a port or port range. Multiple individual ports can be entered as a comma-separated list.
Considerations and Limitations
- Outbound connections are allowed by default. Add a default deny rule if the deployment requires explicit allow-list behavior.
- Firewall rules only apply to traffic passing through the MX. They do not apply to traffic sourced from or destined to the MX itself, such as LDAP binds or inbound Client VPN connections.
- MX VLAN interface IPs and MX WAN IPs are not evaluated as source or destination matches in allow or deny rules.
- Rules are flow-based. Existing flows may continue until they time out or the MX is rebooted. Rules apply immediately to new outbound Internet flows.
- In a hub-and-spoke VPN topology where the spoke uses the Hub as its default route, spoke traffic is evaluated against the Hub’s outbound Layer 3 firewall rules.
    
