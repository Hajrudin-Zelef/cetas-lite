---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066-1
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "license"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066.md
source_anchor: ""
source_lines: [1, 102]
sha256: 700e4834b7a1b744683151a85050a79e2ad25bce7015728cc1814938c461160d
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
Both switches in the Cisco StackWise Virtual pair must be directly connected to each other.
Both switches in the Cisco StackWise Virtual pair must be of the same switch model.
Both switches in the Cisco StackWise Virtual pair must be running the same license level.
Both switches in the Cisco StackWise Virtual pair must be running the same software version.
Both switches in the Cisco StackWise Virtual pair must be running the same SDM template.
All the ports used for configuring a StackWise Virtual Link (SVL) must share the same speed. For example, you cannot configure
a 10G or a 40G port to form an SVL, simultaneously.
Restrictions for Cisco StackWise Virtual
Common Restrictions
The following are the restrictions common to all the switches:
If a dual-rate optic is used as SVL and/or DAD link, it automatically links up to the highest speed supported by the dual-rate
optic (example, 10/25G dual optic will link up at 25G, and 40/100G dual optic will link up at 100G), and lower speeds cannot
be configured on the SVL and/or DAD links.
When deploying Cisco StackWise Virtual, ensure that VLAN ID 4094 is not used anywhere on the network. All inter-chassis system
control communication between stack members is carried over the reserved VLAN ID 4094 from the global range. This does not
apply to Cisco Catalyst 9500X Series Switches.
Dual-Active Detection (DAD) and SVL configuration must be performed manually and the devices should be rebooted for the configuration
changes to take effect. This does not apply to Cisco Catalyst 9500X Series Switches.
Cisco StackWise Virtual configuration commands will be recognised only on a switch running Network Advantage license. The
configuration commands will not be recognised on a Network Essentials license.
Only Cisco Transceiver Modules are supported.
Configuring SVL using 1G interfaces is not supported.
The interface VLAN MAC address that is assigned by default, can be overridden using the mac-address command. If this command is configured on a single SVI or router port that requires Layer 3 injected packets, all other SVIs
or routed ports on the device also must be configured with the same first four most significant bytes (4MSB) of the MAC address.
For example, if you set the MAC address of any SVI to xxxx.yyyy.zzzz, set the MAC address of all other SVIs to start with
xxxx.yyyy. If Layer 3 injected packets are not used, this restriction does not apply.
Note
This applies to all Layer 3 ports, SVIs, and routed ports. This does not apply to GigabitEthernet0/0 port.
Secure StackWise Virtual is supported only on two node front-side stacking.
Do not configure Secure Stackwise Virtual and Federal Information Processing Standards (FIPS) at the same time as they are
mutually exclusive features that cannot co-exist.
Configuring both at the same time is redundant as Secure StackWise Virtual is FIPS 140-2 compliant. Secure StackWise Virtual
will encrypt control packets as well. Therefore, enabling FIPS is not required. This does not apply to Cisco Catalyst 9500X
Series Switches.
Switches operating in SVL models are FIPS 140-2 compliant. FIPS keys must be configured on both switch members individually
to bring up SVL with FIPS mode.
Only 128-bit authorization key is supported.
Secure StackWise Virtual is not supported on DAD Links.
Broadcast, Unknown Unicast and Multicast (BUM) Traffic Optimization is not applicable to VLANs with standalone or physical
ports.
Restrictions for Cisco Catalyst 9500X Series Switches
Configuring SVL using 1G interfaces is not supported.
SVL and DAD links are not supported on breakout interfaces when operating in SVL mode. Breakout interfaces are only supported
for regular data and control traffic carrying ports (not SVL and DAD links) when operating in SVL mode.
Cisco StackWise Virtual is not supported on No Payload Encryption (NPE) images.
Restrictions for Cisco Catalyst 9500 Series Switches
SVL and DAD links are not supported on breakout interfaces when operating in SVL mode. Breakout interfaces are only supported
for regular data and control traffic carrying ports (not SVL and DAD links) when operating in SVL mode.
When configuring SVLs on Cisco Catalyst 9500 Series Switches with C9500-NM-2Q (2x40G), you cannot use a combination of fixed
downlink and modular uplink ports. SVLs should have the same speed on each member. The 40G ports on a C9500-NM-2Q cannot be
combined with the downlink ports on a switch as they have different speeds.
In a Cisco StackWise Virtual solution, ports that support 4X10G breakout cables and QSA can be used only as data-only ports
and cannot be used for configuring SVLs or DAD links.
Restrictions for Cisco Catalyst 9500 Series High Performance Switches
On C9500-32C switches, you can configure SVL and DAD only on interfaces numbered 1-16 on the front panel of the switch.
On C9500-32QC, you can configure SVL and DAD only on native 100G and 40G interfaces (default configuration ports). You cannot
configure SVL and DAD on converted 100G and 40G interfaces.
SVL and DAD ports are not supported on sub-interfaces.
In a Cisco StackWise Virtual solution, ports that support 4X10G breakout cables can be used only as data-only ports and cannot
be used for configuring SVLs or DAD links.
Information About Cisco StackWise Virtual
Cisco StackWise Virtual on Cisco Catalyst 9500 Series Switches
This section describes the Cisco StackWise Virtual features specific to Cisco Catalyst 9500 Series Switches and Cisco Catalyst
9500 Series High Performance Switches.
The following switch models support Cisco StackWise Virtual:
Table 1. Switches Supporting StackWise Virtual
Cisco Catalyst 9500 Series Switches
Cisco Catalyst 9500 Series High Performance Switches
Cisco Catalyst 9500X Series Switches
C9500-24Q
C9500-32C
C9500X-28C8D
C9500-12Q
C9500-32QC
C9500X-60L4D
C9500-40X
C9500-24Y4C
C9500-16X
C9500-48Y4C
On C9500-40X and C9500-16X models of the Cisco Catalyst 9500 Series Switches, you can configure SVLs and DAD links on any
of the network modules. C9500-40X and C9500-16X support the following network modules.
Table 2. Supported Network Modules
Cisco Catalyst 9500 Series Switch Model
Network Modules
C9500-40X
C9500-NM-8X
C9500-NM-2Q
C9500-16X
You can establish SVLs and DAD links on Cisco Catalyst 9500 Series Switches using a combination of modular uplink and fixed
downlink ports.
You can establish SVLs using 40G or 10G Ethernet connections on Cisco Catalyst 9500 Series Switches and 100G, 40G, 25G and
10G Ethernet connections on Cisco Catalyst 9500 Series High Performance Switches.
Note
Ensure that the cables and/or transceivers on all the SVL and DAD links are not disturbed during SVL bring up.
You can configure up to 8 SVLs in a Cisco StackWise Virtual solution using Cisco Catalyst 9500 Series Switches or Cisco Catalyst
9500 Series High Performance Switches.
Overview of Cisco StackWise Virtual
Cisco StackWise Virtual is a network system virtualization technology that pairs two directly connected switches into one
virtual switch. The switches in a Cisco StackWise Virtual solution increase operational efficiency by using single control
and management plane, scale system bandwidth with distributed forwarding plane, and help in building resilient networks using
the recommended network design. Cisco StackWise Virtual allows two directly connected physical switches to operate as a single
