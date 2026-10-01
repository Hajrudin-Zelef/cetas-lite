---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-4f4bd2bc-3
title: "c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-4f4bd2bc"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-4f4bd2bc.md
source_anchor: ""
source_lines: [68, 124]
sha256: 13093e9b630ffd8009642462920334c78583ff0b1a581520bd8628f80d65ed1f
---

# c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-4f4bd2bc

You can configure access rules that control management traffic destined to the ASA. Access control rules for to-the-box management traffic (defined by such commands as http, ssh, or telnet) have higher precedence than an management access rule applied with the control-plane option. Therefore, such permitted management traffic will be allowed to come in even if explicitly denied by the to-the-box ACL.
Information About EtherType Rules
This section describes EtherType rules and includes the following topics:
Supported EtherTypes and Other Traffic
An EtherType rule controls the following:
- EtherType identified by a 16-bit hexadecimal number, including common types IPX and MPLS unicast or multicast.
- Ethernet V2 frames.
- BPDUs, which are permitted by default. BPDUs are SNAP-encapsulated, and the ASA is designed to specifically handle BPDUs.
- Trunk port (Cisco proprietary) BPDUs. Trunk BPDUs have VLAN information inside the payload, so the ASA modifies the payload with the outgoing VLAN if you allow BPDUs.
- IS-IS.
Access Rules for Returning Traffic
Because EtherTypes are connectionless, you need to apply the rule to both interfaces if you want traffic to pass in both directions.
Allowing MPLS
If you allow MPLS, ensure that Label Distribution Protocol and Tag Distribution Protocol TCP connections are established through the ASA by configuring both MPLS routers connected to the ASA to use the IP address on the ASA interface as the router-id for LDP or TDP sessions. (LDP and TDP allow MPLS routers to negotiate the labels (addresses) used to forward packets.)
On Cisco IOS routers, enter the appropriate command for your protocol, LDP or TDP. The interface is the interface connected to the ASA.
Licensing Requirements for Access Rules
Prerequisites
Before you can create an access rule, create the ACL. See the general operations configuration guide for more information.
Guidelines and Limitations
This section includes the guidelines and limitations for this feature.
Supported in single and multiple context mode.
Supported in routed and transparent firewall modes.
Supports IPv6. The source and destination addresses can include any mix of IPv4 and IPv6 addresses.
- The per-user ACL uses the value in the timeout uauth command, but it can be overridden by the AAA per-user session timeout value.
- If traffic is denied because of a per-user ACL, syslog message 109025 is logged. If traffic is permitted, no syslog message is generated. The log option in the per-user ACL has no effect.
Default Settings
See the “Implicit Permits” section.
Configuring Access Rules
Detailed Steps
|  |  | 
|---|---|
| access-group access_list {{ in \| out } interface interface_name [ per-user-override \| control-plane ] \| global } ciscoasa(config)# access-group outside_access in interface outside | Binds an ACL to an interface or applies it globally. Specify the extended or EtherType ACL name. You can configure one access-group command per ACL type per interface. You cannot reference empty ACLs or ACLs that contain only a remark. For an interface-specific rule:  By default, VPN remote access traffic is not matched against interface ACLs. However, if you use the no sysopt connection permit-vpn command to turn off this bypass, the behavior depends on whether there is a vpn-filter applied in the group policy and whether you set the per-user-override option: – No per-user-override, no vpn-filter —Traffic is matched against the interface ACL. – No per-user-override, vpn-filter —Traffic is matched first against the interface ACL, then against the VPN filter. – per-user-override, vpn-filter —Traffic is matched against the VPN filter only. See Per-User ACL Guidelines. For a global rule, specify the global keyword to apply the ACL to the inbound direction of all interfaces. | 
Examples
The following example shows how to use the access-group command:
The access-list command lets any host access the global address using port 80. The access-group command specifies that the access-list command applies to traffic entering the outside interface.
Monitoring Access Rules
To monitor network access, enter the following command:
Configuration Examples for Permitting or Denying Network Access
This section includes typical configuration examples for permitting or denying network access.
The following example adds a network object for inside server 1, performs static NAT for the server, and enables access to from the outside for inside server 1.
The following example allows all hosts to communicate between the inside and hr networks but only specific hosts to access the outside network:
For example, the following sample ACL allows common EtherTypes originating on the inside interface:
The following example allows some EtherTypes through the ASA, but it denies all others:
The following example denies traffic with EtherType 0x1256 but allows all others on both interfaces:
The following example uses object groups to permit specific traffic on the inside interface:
Feature History for Access Rules
Table 6-2 lists each feature change and the platform release in which it was implemented.
|  |  |  | 
|---|---|---|
| Interface access rules | 7.0(1) | Controlling network access through the ASA using ACLs. We introduced the following command: access-group. | 
| Global access rules | 8.3(1) | Global access rules were introduced. We modified the following command: access-group. | 
| Support for Identity Firewall | 8.4(2) | You can now use identity firewall users and groups for the source and destination. You can use an identity firewall ACL with access rules, AAA rules, and for VPN authentication. We modified the following commands: access-list extended. | 
| EtherType ACL support for IS-IS traffic | 8.4(5), 9.1(2) | In transparent firewall mode, the ASA can now pass IS-IS traffic using an EtherType ACL. We modified the following command: access-list ethertype { permit \| deny } is-is. | 
| Support for TrustSec | 9.0(1) | You can now use TrustSec security groups for the source and destination. You can use an identity firewall ACL with access rules. We modified the following commands: access-list extended. | 
| Unified ACL for IPv4 and IPv6 | 9.0(1) | ACLs now support IPv4 and IPv6 addresses. You can even specify a mix of IPv4 and IPv6 addresses for the source and destination. The any keyword was changed to represent IPv4 and IPv6 traffic. The any4 and any6 keywords were added to represent IPv4-only and IPv6-only traffic, respectively. The IPv6-specific ACLs are deprecated. Existing IPv6 ACLs are migrated to extended ACLs. See the release notes for more information about migration. We modified the following commands: access-list extended, access-list webtype. We removed the following commands: ipv6 access-list, ipv6 access-list webtype, ipv6-vpn-filter | 
| Extended ACLand object enhancement to filter ICMP traffic by ICMP code | 9.0(1) | ICMP traffic can now be permitted/denied based on ICMP code. We introduced or modified the following commands: access-list extended, service-object, service. | 
| Transactional Commit Model on Rule Engine for Access groups | 9.1(5) | When enabled, a rule update is applied after the rule compilation is completed; without affecting the rule matching performance. We introduced the following commands: asp rule-engine transactional-commit, show running-config asp rule-engine transactional-commit, clear configure asp rule-engine transactional-commit. |
