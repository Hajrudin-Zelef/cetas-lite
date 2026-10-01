---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-1ff3fcb3-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-1ff3fcb3"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-1ff3fcb3.md
source_anchor: ""
source_lines: [1, 106]
sha256: f0aab0eb5cf5b20c0fd86bc89bfd9be2df72829b8b3c2dc2d34ada07b7db04aa
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-1ff3fcb3

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
One-to-One VLAN mapping can be configured only on trunk ports and not on dynamic trunk.
One-to-One VLAN mapping should be identical on both ports.
S-VLAN should be created and present in the allowed VLAN list of the trunk port where One-to-One VLAN mapping is configured.
Restrictions for VLAN Mapping
If VLAN mapping is enabled on an EtherChannel, the configuration does not apply to all member ports of the EtherChannel bundle
but applies only to the EtherChannel interface.
If VLAN mapping is enabled on an EtherChannel and a conflicting mapping translation is enabled on a member port, the port
is removed from the EtherChannel.
If a port belonging to an EtherChannel is configured with a VLAN mapping and the EtherChannel is configured with a conflicting
VLAN mapping, the port is removed from the EtherChannel.
Default native VLANs, user-configured native VLANs, and reserved VLANs cannot be used for VLAN mapping.
The S-VLAN used for VLAN mapping cannot be a part of any other Layer 3 configurations, EVPN, or LISP.
PVLAN support is not available when VLAN mapping is configured.
Restrictions for One to One VLAN Mapping
When One-to-One VLAN mapping is configured, multiple C-VLANs cannot be mapped to the same S-VLAN
Merging of C-VLAN and S-VLAN spanning-tree topology is not supported in case of one-to-one vlan mapping.
About VLAN Mapping
In a typical deployment of VLAN mapping, you want service provider to provide a transparent switching infrastructure that
includes customers’ switches at the remote location as a part of local site. This allows customers to use the same VLAN ID
space and run Layer 2 control protocols seamlessly across the provider network. In such scenarios, we recommend that service
providers do not impose their VLAN IDs on their customers.
One way to establish translated VLAN IDs (S-VLANs) is to map customer VLANs to VLANs (called VLAN ID translation) on trunk
ports that are connected to a customer network. Packets entering the port are mapped to service provider VLAN (S-VLAN) based
on the port number and the packet’s original customer VLAN-ID (C-VLAN).
Service providers’ internal assignments might conflict with a customer’s VLAN. To isolate customer traffic, a service provider
decides to map a specific VLAN into another one while the traffic is in its cloud.
Deployment Example
In the figure, the service provider provides Layer 2 VPN service to two different customers, A and B. The service provider separates the
data and control traffic between the two customers and from the providers’ own control traffic. The service provider network
must also be transparent to the customer edge devices.
All forwarding operations on Catalyst 9000 series switch are performed using S-VLAN and not C-VLAN information because the
VLAN ID is mapped to the S-VLAN on ingress.
Note
When you configure features on a port for VLAN mapping, you always use the S-VLAN rather than C-VLAN.
On an interface configured for VLAN mapping, the specified C-VLAN packets are mapped to the specified S-VLAN when they enter
the port. Symmetrical mapping to the customer C-VLAN occurs when packets exit the port.
The switch supports one-to-one VLAN mapping on trunk ports.
The switch supports these types of VLAN mapping on trunk ports:
One-to-one VLAN mapping.
Selective QinQ.
Figure shows a topology where a customer uses the same VLANs in multiple sites on different sides of a service-provider network.
The C-VLAN IDs is mapped to service-provider VLAN IDs for packet travel across the service-provider backbone. The C-VLAN IDs
are retrieved at the other side of the service-provider backbone for use in the other customer site. Configure the same set
of VLAN mappings at a customer-connected port on each side of the service-provider network.
One-to-One VLAN Mapping
One-to-one VLAN mapping occurs at the ingress and egress of the port and maps the customer C-VLAN ID in the 802.1Q tag to
the service-provider S-VLAN ID. You can also specify that packets with all other Vlan IDs are forwarded.
Selective Q-in-Q
Selective QinQ maps the specified customer VLANs entering the UNI to the specified S-VLAN ID. The S-VLAN ID is added to the
incoming unmodified C-VLAN and the packet travels the service provider network double-tagged. At the egress, the S-VLAN ID
is removed and the customer VLAN-ID is retained on the packet. By default, packets that do not match the specified customer
VLANs are dropped.
Configuration Guidelines for VLAN Mapping
Note
By default, no VLAN mapping is configured.
Maximum number of VLAN mapping configurations supported is 512 system wide.
Guidelines include the following:
If the VLAN mapping is enabled on an EtherChannel, the configuration does not apply to all member ports of the EtherChannel
bundle and applies only to the EtherChannel interface.
If a port belonging to an EtherChannel is configured with a VLAN mapping and the EtherChannel is configured with a conflicting
VLAN mapping, then the port is removed from the EtherChannel.
The member port of an EtherChannel is removed from the EtherChannel bundle if the mode of the port is changed to anything
other than ‘trunk’ mode.
To process control traffic consistently, either enable Layer 2 protocol tunneling (recommended), as follows:
Default native VLANs, user-configured native VLANs, and reserved VLANs (range 1002-1005) cannot be used for VLAN mapping.
The S-VLAN used for VLAN mapping cannot be a part of any other Layer 3 configurations like EVPN or LISP.
PVLAN support is not available when VLAN mapping is configured.
Configuration Guidelines for One-to-One VLAN Mapping
One-to-One VLAN mapping can be configured only on trunk ports and not on dynamic trunk.
One-to-One VLAN mapping should be identical on both ports.
S-VLAN should be created and present in the allowed VLAN list of the trunk port where One-to-One VLAN mapping is configured.
When One-to-One VLAN mapping is configured, multiple C-VLANs cannot be mapped to the same S-VLAN.
Merging of C-VLAN and S-VLAN spanning-tree topology is not supported in case of one-to-one VLAN mapping.
Configuration Guidelines for Selective Q-in-Q
S-VLAN should be created and present in the allowed VLAN list of the trunk port where Selective Q-in-Q is configured.
When Selective Q-in-Q is configured, the device supports Layer 2 protocol tunneling for CDP, STP, LLDP, and VTP. For emulated
point-to-point network topologies, it also supports PAgP, LACP, and UDLD protocols.
IP routing is not supported on Selective Q-in-Q enabled ports.
IPSG is not supported on Selective Q-in-Q enabled ports.
How to Configure VLAN Mapping
The following sections provide information about configuring VLAN mapping:
One-to-One VLAN Mapping
Note
VLAN Mapping is supported only with the network-advantage license level.
To configure one-to-one VLAN mapping to map a customer VLAN ID to a service-provider VLAN ID, perform this task:
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configure terminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
interface interface-id
Example:
Device(config)# interface gigabitethernet1/0/1
Enters interface configuration mode for the interface that is connected to the service-provider network. You can enter a physical
interface or an EtherChannel port channel.
