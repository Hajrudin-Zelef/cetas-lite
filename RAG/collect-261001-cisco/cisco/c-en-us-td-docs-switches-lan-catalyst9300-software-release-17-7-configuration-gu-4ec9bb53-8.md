---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53-8
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53.md
source_anchor: ""
source_lines: [341, 418]
sha256: 7b59bc9580eca75da87d85bdb5ec4cfeee0d6fca5683f6ec549205c8a04ca5ed
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53

Device# show device-tracking policy glean_only_DHCP               
Device-tracking policy glean_only_DHCP configuration: 
  security-level guard
  device-role node
  NOT gleaning from Neighbor Discovery
  gleaning from DHCP6
  NOT gleaning from ARP
  NOT gleaning from DHCP4
  NOT gleaning from protocol unkn
Policy glean_only_DHCP is applied on the following targets: 
Target               Type  Policy               Feature        Target range
Gi1/0/1              PORT  glean_only_DHCP      Device-tracking vlan all
IEEE 802.1x authentication is enabled.
This means only authenticated hosts are allowed to request addresses from the DHCP server and attach themselves to the network.
| Note | The following 802.1x configuration is for example purposes only. | 
<output truncated>
 interface GigabitEthernet 1/0/1
 description 802.1x+MAB+IPT
 authentication control-direction in
 authentication event server dead action authorize vlan <vlan id>
 authentication event no-response action authorize vlan <vlan id>
 authentication event server alive action reinitialize
 authentication host-mode multi-domain
 authentication port-control auto
 authentication periodic
 authentication timer reauthenticate server
 authentication violation protect
 mab
 trust device cisco-phone
 dot1x pae authenticator
 dot1x timeout quiet-period 30
 dot1x timeout server-timeout 5
 dot1x timeout tx-period 1
 dot1x max-req 1
 dot1x max-reauth-req 1
<output truncated>
Events that cause a change in the configuration occur in any typical network. For example, a host may be unplugged from one port and then plugged back into another port, or an interface may flap, or you may have configured the shutdown, followed by the no shutdown interface configuration commands. For the duration that the host is not connected, or the interface is down, the host or interface is considered "unauthenticated". Because of this absence of host or interface authentication, the corresponding binding table entry is removed from the binding table.
When such a host connects back to the network or when such an interface is restored, the client does not reinstantiate the DHCP sequence until the DHCP lease time expires. Until the DHCP sequence is reinstantiated, a valid address fails to be stored in the binding table. If the entry is not in the binding table, the IPv6 Source Guard’s filter function drops all packets initiated by that host.
In order to prevent such a situation, configure the data-glean recovery function.
To configure data-glean recovery, create a custom SISF-based device-tracking policy, configure the data-glean policy parameter to recover binding information from DHCP Server, and attach it to the necessary targets.
| Note | When configuring data-glean recovery from DHCP, for binding information retrieval to work as expected, the DHCPv6 Leasequery configuration (as in RFC 5007), is required. Ensure that the leasequery configuration is enabled on the DHCP Server. | 
glean_only_DHCP), to recover binding information. It remains attached to the same target as the IPv6 Source Guard policy, that is, Gigabit
                                 Ethernet 1/0/1:
Device# configure terminal
Device(config)# device-tracking policy glean_only_DHCP
Device(config-device-tracking)#  data-glean recovery dhcp
Device(config-device-tracking)# exit
Device# show device-tracking policy glean_only_DHCP               
Device-tracking policy glean_only_DHCP configuration: 
  security-level guard
  device-role node
  data-glean recovery dhcp                       <<< Recovery of binding information is configured.
  NOT gleaning from Neighbor Discovery
  gleaning from DHCP6
  NOT gleaning from ARP
  NOT gleaning from DHCP4
  NOT gleaning from protocol unkn
Policy glean_only_DHCP is applied on the following targets: 
Target               Type  Policy               Feature        Target range
Gi1/0/1              PORT  glean_only_DHCP      Device-tracking vlan all 
Device# show device-tracking policies interface Gi1/0/1
Target               Type  Policy               Feature        Target range
Gi1/0/1              PORT  glean_only_DHCP      Device-tracking vlan all
Gi1/0/1              PORT  src-guard-policy     Source guard   vlan all
With this additional configuration, valid entries are automatically restored in the binding table if they are removed prematurely.
| Related Topic | Document Title | 
|---|---|
| SISF | Configuring SISF-Based Device Tracking chapter of the Security Configuration Guide | 
| Description | Link | 
|---|---|
| The Cisco Support website provides extensive online resources, including documentation and tools for troubleshooting and resolving technical issues with Cisco products and technologies. To receive security and technical information about your products, you can subscribe to various services, such as the Product Alert Tool (accessed from Field Notices), the Cisco Technical Services Newsletter, and Really Simple Syndication (RSS) Feeds. Access to most tools on the Cisco Support website requires a Cisco.com user ID and password. | http://www.cisco.com/support | 
This table provides release and related information for the features explained in this module.
These features are available in all the releases subsequent to the one they were introduced in, unless noted otherwise.
| Release | Feature | Feature Information | 
|---|---|---|
| Cisco IOS XE Everest 16.5.1a | IPv6 First Hop Security | First Hop Security in IPv6 is a set of IPv6 security features, the policies of which can be attached to a physical interface, an EtherChannel interface, or a VLAN. An IPv6 software policy database service stores and accesses these policies. When a policy is configured or modified, the attributes of the policy are stored or updated in the software policy database, then applied as was specified. The IPv6 Snooping Policy feature has been deprecated. Although the commands are visible on the CLI and you can configure them, we recommend that you use the Switch Integrated Security Feature (SISF)-based Device Tracking feature instead. | 
| Cisco IOS XE Amsterdam 17.1.1 | IPv6 ND Inspection | Starting with this release, the IPv6 ND Inspection feature is deprecated and the SISF- based device tracking feature replaces it and offers the same capabilities. While the IPv6 ND Inspection commands are still available on the CLI and the existing configuration continues to be supported, the commands will be removed from the CLI in a later release. For more information about the replacement feature, see the Configuring SISF-Based Device Tracking chapter in this guide. | 
Use the Cisco Feature Navigator to find information about platform and software image support. To access Cisco Feature Navigator, go to Cisco Feature Navigator.
