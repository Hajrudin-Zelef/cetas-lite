---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59-1
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59.md
source_anchor: ""
source_lines: [1, 102]
sha256: 2fa668bb1c23dd9b438e9a91f6ba251c2c68230e1bc4c0308f4bedf6eb4402e9
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59

Cisco Nexus 9000 Series NX-OS Interfaces Configuration Guide, Release 10.3(x)
Bias-Free Language
Bias-Free Language
The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
A port channel is an aggregation of multiple physical interfaces that creates a logical interface. You can bundle up to 32
individual active links into a port channel to provide increased bandwidth and redundancy. Port channeling also load balances
traffic across these physical interfaces. The port channel stays operational as long as at least one physical interface within
the port channel is operational.
You can create a Layer 2 port channel by bundling compatible Layer 2 interfaces, or you can create Layer 3 port channels by
bundling compatible Layer 3 interfaces. You cannot combine Layer 2 and Layer 3 interfaces in the same port channel.
You can also change the port channel from Layer 3 to Layer 2. See the Configuring Layer 2 Interfaces chapter for information
about creating Layer 2 interfaces.
A Layer 2 port channel interface and it's member ports can have different STP parameters. Changing the STP parameters of the
port channel does not impact the STP parameters of the member ports because a port channel interface takes precedence if the
member ports are bundled.
Note
After a Layer 2 port becomes part of a port channel, all switchport configurations must be done on the port channel; you can
no longer apply switchport configurations to individual port-channel members. You cannot apply Layer 3 configurations to an
individual port-channel member either; you must apply the configuration to the entire port channel.
In releases prior to Cisco NX-OS Release 9.3(7), in a port-channel configuration with a member port operating as an individual
(I), you can define the STP port-type under the member port rather than the port-channel.
Beginning with Cisco NX-OS Release 9.3(7), in a port-channel configuration with a member port operating as an individual (I),
you can no longer define the STP port-type under the member port. It remains blocked by the STP. You must configure the STP
port-type under the port-channel.
You can use static port channels, with no associated aggregation protocol, for a simplified configuration.
For more flexibility, you can use the Link Aggregation Control Protocol (LACP), which is defined in IEEE 802.3ad. When you
use LACP, the link passes protocol packets. You cannot configure LACP on shared interfaces.
See the LACP Overview section for information about LACP.
Port Channels
A port channel bundles physical links into a channel group to create a single logical link that provides the aggregate bandwidth
of up to 32 physical links. If a member port within a port channel fails, the traffic previously carried over the failed link
switches to the remaining member ports within the port channel.
However, you can enable the LACP to use port channels more flexibly. Configuring port channels with LACP and static port channels
require a slightly different procedure (see the “Configuring Port Channels” section).
Note
The device does not support Port Aggregation Protocol (PAgP) for port channels.
Each port can be in only one port channel. All the ports in a port channel must be compatible; they must use the same speed
and duplex mode (see the “Compatibility Requirements” section). When you run static port channels with no aggregation protocol,
the physical links are all in the on channel mode; you cannot change this mode without enabling LACP (see the “Port-Channel
Modes” section).
You can create port channels directly by creating the port-channel interface, or you can create a channel group that acts
to aggregate individual ports into a bundle. When you associate an interface with a channel group, the software creates a
matching port channel automatically if the port channel does not already exist. In this instance, the port channel assumes
the Layer 2 or Layer 3 configuration of the first interface. You can also create the port channel first. In this instance,
the Cisco NX-OS software creates an empty channel group with the same channel number as the port channel and takes the default
Layer 2 or Layer 3 configuration, as well as the compatibility configuration (see the “Compatibility Requirements” section).
Note
The port channel is operationally up when at least one of the member ports is up and that port’s status is channeling. The
port channel is operationally down when all member ports are operationally down.
Port-Channel Interfaces
The following shows port-channel interfaces.
Figure 1. Port-Channel Interfaces
You can classify port-channel interfaces as Layer 2 or Layer 3 interfaces. In addition, you can configure Layer 2 port channels
in either access or trunk mode. Layer 3 port-channel interfaces have routed ports as channel members.
You can configure a Layer 3 port channel with a static MAC address. If you do not configure this value, the Layer 3 port channel
uses the router MAC of the first channel member to come up. See the Cisco Nexus 9000 Series NX-OS Layer 2 Switching Configuration Guide for information about configuring static MAC addresses on Layer 3 port channels.
See the "Configuring Layer 2 Interfaces" chapter for information about configuring Layer 2 ports in access or trunk mode and
the "Configuring Layer 3 Interfaces" chapter for information about configuring Layer 3 interfaces and subinterfaces.
Basic Settings
You can configure the following basic settings for the port-channel interface:
Bandwidth—Use this setting for informational purposes only; this setting is to be used by higher-level protocols.
Delay—Use this setting for informational purposes only; this setting is to be used by higher-level protocols.
Description
Duplex
IP addresses
Maximum Transmission Unit (MTU)
Shutdown
Speed
Compatibility Requirements
When you add an interface to a channel group, the software checks certain interface attributes to ensure that the interface
is compatible with the channel group. For example, you cannot add a Layer 3 interface to a Layer 2 channel group. The Cisco
NX-OS software also checks a number of operational attributes for an interface before allowing that interface to participate
in the port-channel aggregation.
The compatibility check includes the following operational attributes:
Network layer
(Link) speed capability
Speed configuration
Duplex capability
Duplex configuration
Port mode
Access VLAN
Trunk native VLAN
Tagged or untagged
Allowed VLAN list
MTU size
SPAN—Cannot be a SPAN source or a destination port
Storm control
Flow-control capability
Flow-control configuration
Media type, either copper or fiber
Use the show port-channel compatibility-parameters command to see the full list of compatibility checks that the Cisco NX-OS uses.
You can only add interfaces configured with the channel mode set to on to static port channels, and you can only add interfaces
configured with the channel mode as active or passive to port channels that are running LACP. You can configure these attributes
on an individual member port. If you configure a member port with an incompatible attribute, the software suspends that port
in the port channel.
Alternatively, you can force ports with incompatible parameters to join the port channel if the following parameters are the
same:
(Link) speed capability
Speed configuration
Duplex capability
Duplex configuration
Flow-control capability
