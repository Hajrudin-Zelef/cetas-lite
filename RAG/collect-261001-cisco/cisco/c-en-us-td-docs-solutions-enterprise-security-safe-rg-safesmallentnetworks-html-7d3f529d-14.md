---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-14
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [445, 472]
sha256: 6bbeddd00c5f4a68e01326184fce3572c97878d37a130835e4bf9d04ee6454e7
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

Although many devices increasingly support 802.1X, there will always be devices that require network connectivity but do not and/or cannot support 802.1X. Examples of such devices include network printers, badge readers, legacy servers, and PXE boot machines. Some provision must be made for these devices.
Cisco provides two features to accommodate non-802.1X devices, MAC Authentication Bypass (MAB) and the Guest VLAN. These features provide fallback mechanisms when there is no 802.1X supplicant. After 802.1X times out on a port, the port can move to an open state if MAB succeeds or if the Guest VLAN is configured. Judicious application of either or both of these features will be required for a successful 802.1X deployment.
Note Network-specific testing will be required to determine the optimal values for 802.1X timers to accommodate the various non-802.1X-capable devices on your network.
802.1X in Small Enterprise Networks
As mentioned above, one of the requirements for 802.1X authentication is the requirement for a supplicant. This has typically been a challenge in enterprise environments with a wide range of devices and limited or no management of many of these devices. In many enterprises this is still the case and this makes a company-wide 802.1X very challenging. At the same time there are pockets of an enterprise network where 802.1X may be a good choice.
For example, 802.1X-protected ports may be a good choice for the network ports in the company's headquarters or main site, as these locations are more likely to have managed PCs.
Other locations in the enterprise network still need protection, but user network access may be better served by a NAC Appliance solution. For networks requiring role-based access control using posture assessments to ensure security compliance, Cisco NAC Appliance should be considered. In addition, network access ports in open areas such as lobbies and meeting rooms may use 802.1X or Cisco Clean Access NAC to protect these ports.
When considering the 802.1X deployment, there are four main 802.1X authentication options to consider.
•Basic 802.1X Authentication—An 802.1X controlled port with an 802.1X client directly connected
•IP Phone Ports—An IP Phone and an 802.1X controlled port with an 802.1X client connected to the phone
•MAC Auth By-Pass—Using the MAC address of the client to provide authentication and bypass the 802.1X authentication process. Printer and legacy device support are typical applications.
•Web Auth—Allowing a user to authenticate by entering username and passwords in a Web page. Legacy device support and guest access are typical deployment applications.
For more information on the Cisco IBNS 802.1X network access solution, see: http://www.cisco.com/go/ibns.
NAC 802.1X and CISF in Combination
The three key access security features discussed above have been discussed in isolation, but can be combined. In particular, the CISF features should be considered "baseline" features that are applied on all access ports and either NAC or 802.1X maybe overlaid on top of the CISF configuration.
The Cisco Clean Access and 802.1X configuration are also compatible (although they are not often combined in wired networks), the key consideration in combining the two is how to give the appearance of a SSO for the end user. Both 802.1X and NAC require authentication, as 802.1X authenticates the client initially, a mechanism of communicating the 802.1X authentication result to the Cisco Clean Access system is required.
If the authenticating clients join an Windows Active Directory network, the Cisco Clean Access Active Directory SSO feature allows the clients to authenticate to active directory once they have performed their 802.1X authentication. The CAM, when a client is detected, checks Active Directory to see if the client has authenticated, which allows a SSO experience for client devices that are using 802.1X and NAC.
Secure Mobility
Today workers use laptops, smartphones, and other smart mobile devices to access information and applications at anytime and from anywhere there is an Internet connection. While embracing a mobile workforce clearly boosts productivity and makes the small enterprise more competitive, there are a number of challenges that arise from the use of mobile technologies. To start with, workers often use the same devices to access both business and personal information. Devices used outside the enterprise onsite controls may potentially introduce viruses, worms, spyware, and other type of malware as mobile workers connect back to the corporate network. Confidential and proprietary information may also be lost or stolen while mobile users connect outside the company premises.
On the other hand, the great variety in hardware types, operating systems, and applications represents a clear challenge to the enforcement of security controls and policies. In order to continue to foster innovation, enable productivity, and meet the needs of the mobile workforce, companies must adapt to the changing trends in mobility. A viable solution is one that enables access for mobile workers while ensuring that the corporate data, assets, and network remain secure. On the other hand, users want the flexibility of choosing how, when, and where to access both personal and professional information to be productive, without being inconvenienced by security checks.
The small enterprise network design provides persistent and secure mobile access by implementing the Cisco AnyConnect Secure Mobility solution. This solution delivers secure, persistent connectivity to all mobile employees independently from the type of mobile device used. The Cisco AnyConnect Secure Mobility solution also ensures a consistent enforcement of the network security policies to all users, no matter when, where, and how they connect to the network.
The Cisco AnyConnect Secure Mobility is a collection of features across multiple Cisco products that extends control and security into borderless networks. The products that work together to provide AnyConnect Secure Mobility are:
•Cisco IronPort Web Security appliance (WSA)
•Cisco ASA 5500 series adaptive security appliance (ASA)
Cisco AnyConnect Secure Mobility addresses the challenges of a mobile workforce by offering the following features:
•Secure, persistent connectivity—Cisco AnyConnect (with the Cisco ASA at the headend) provides the remote access connectivity portion of AnyConnect Secure Mobility. The connection is secure because both the user and device must be authenticated and validated prior to being provided access to the network. The connection is persistent because Cisco AnyConnect is typically configured to be always-on, even when roaming between networks. Although Cisco AnyConnect is always-on, it is also flexible enough to apply different policies based on location, allowing users access to the Internet in a "captive portal" situation, when users must accept terms of agreement before accessing the Internet.
•Persistent security and policy enforcement—The Web Security appliance applies context-aware policies, including enforcing acceptable use policies and protection from malware for all users, including mobile (remote) users. The WSA also accepts user authentication information from the AnyConnect client, providing an automatic authentication step for the user to access Web content.
Figure 22 Cisco AnyConnect Secure Mobility Solution
