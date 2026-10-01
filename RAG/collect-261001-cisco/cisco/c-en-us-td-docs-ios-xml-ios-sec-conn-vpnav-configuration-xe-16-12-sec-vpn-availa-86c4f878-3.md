---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-vpnav-configuration-xe-16-12-sec-vpn-availa-86c4f878-3
title: "c-en-us-td-docs-ios-xml-ios-sec-conn-vpnav-configuration-xe-16-12-sec-vpn-availa-86c4f878"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-vpnav-configuration-xe-16-12-sec-vpn-availa-86c4f878.md
source_anchor: ""
source_lines: [285, 309]
sha256: d1133ef04371b0040be7214bbc36bf1a9919ab5eef7429b0fb81a9d41413f7fe
---

# c-en-us-td-docs-ios-xml-ios-sec-conn-vpnav-configuration-xe-16-12-sec-vpn-availa-86c4f878

The standby name needs to be configured on all devices in the standby group, and the standby address needs to configured
on at least one member of the group. If the standby name is removed from the router, the IPsec SAs will be deleted. If the
standby name is added again, regardless of whether the same name or a different name is used, the crypto map (using the redundancy
option) will have to be reapplied to the interface.
The Cisco Support and Documentation website provides online resources to download documentation, software, and tools. Use
these resources to install and configure the software and to troubleshoot and resolve technical issues with Cisco products
and technologies. Access to most tools on the Cisco Support and Documentation website requires a Cisco.com user ID and password.
Feature Information for IPsec VPN High Availability Enhancements
The following table provides release information about the feature or features described in this module. This table lists
only the software release that introduced support for a given feature in a given software release train. Unless noted otherwise,
subsequent releases of that software release train also support that feature.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco
Feature Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
Table 1. Feature Information for IPsec VPN High Availability Enhancements
Feature Name
Releases
Feature Information
IPsec VPN High Availability Enhancements
Cisco IOS XE 3.1.0S
The IPsec VPN High Availability Enhancements feature consists of two features:Reverse Route Injection (RRI) and Hot Standby
Router Protocol (HSRP) with IPsec. When used together, these two features provide you with a simplified network design for
VPNs and reduced configuration complexity on remote peers when defining gateway lists.
The following commands were introduced or modified:
cryptomap (interface IPsec),
reverse-route.
