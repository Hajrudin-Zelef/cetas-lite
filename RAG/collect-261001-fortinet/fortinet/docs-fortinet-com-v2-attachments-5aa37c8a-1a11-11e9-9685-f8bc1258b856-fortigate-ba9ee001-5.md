---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate-ba9ee001-5
title: "docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001.md
source_anchor: ""
source_lines: [656, 817]
sha256: 4886dfd47e51c632c10070befbe207cebbfc894e109c6eb437b61bf183b1c0d7
---

# docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001

Networking in Transparent Mode VLANs and Forwarding Domains
Example configuration
This example has three forwarding domains and VLANs configured. In this example, there are two VDOMs in
Transparent Mode: root and MGMT. Forwarding domain 0 is the default on the FortiGate or VDOM in
Transparent Mode.
l Root VDOM has:
l 3 forwarding domains, 0, 340, and 341.
l VLAN 340 configured on port1; packets will be tagged with ID 340
l VLAN 341 configured on port1; packets will be tagged with ID 341
l All other ports are untagged
l MGMT VDOM has got only the default forwarding domain 0
The expected behavior is the following:
l Packets untagged ingressing port1, port3 and port4 belong to the same broadcast domain in the root VDOM
l Packets tagged with VLAN 340 ingressing port1 and Packets untagged ingressing port2 belong to the same
broadcast domain in the root VDOM
l Packets tagged with VLAN 341 ingressing port1 and Packets untagged ingressing port5 belong to the same
broadcast domain in the root VDOM
l Packets untagged ingressing port6 belong to a different broadcast domain in the MGMT VDOM
CLI Syntax for forwarding domain 340
config system interface
edit "VLAN340"
set forward-domain 340
set interface "port1"
set vlanid 340
next
edit "port3"
set forward-domain 340
next
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
23

VLANs and Forwarding Domains Networking in Transparent Mode
end
VLANs vs Forwarding Domains
There are several differences between VLAN and a forwarding domain configured on a FortiGate in Transparent
Mode:
l A forwarding domain is used to create separated broadcast domains between VLANs and allow independent VLAN
learning - IVL (MAC addresses in the FDB). This would be equivalent to creating VLANs on a regular L2 switch.
When VLANs are used in the network, configuring different forwarding
domains is essential to avoid broadcast duplications. See also section
Default VLAN forwarding behavior for additional information.
l VLANs configured on interfaces are only used for tagging packets egressing the port and classifying packets at
ingress.
l The packets processed by the direct interface (or port) itself are always sent untagged and must be received
untagged.
VLAN Forwarding
VLAN forwarding allows you to forward all VLANs traffic of a trunk that was connecting two network devices and
where the FortiGate has been introduced, without having to perform any further configuration.
It is recommended to configure forwarding domains for each VLAN and disable this parameter in order to avoid
packet from looping into the trunk from one VLAN to another. By default, the parametervlanforward is
disabled on each physical interface of a FortiGate or VDOM in Transparent mode.
Unknown VLANs and VLAN Forwarding
When a FortiGate receives a tagged frame with an unknown VLAN ID, traffic can be handled one of two ways,
depending on whether VLAN forwarding is enabled or disabled.
By default, VLAN forwarding is disabled and any frames that are tagged with an unknown VLAN ID are dropped
by the FortiGate.
If you enable VLAN forwarding, frames tagged with an unknown VLAN ID are forwarded from the port that
received the frames to all other ports in the same forwarding domain(s). This allows you to insert the FortiGate
between two devices using trunk ports without any further configuration.
VLAN Trunking and MAC Address Learning
A FortiGate port becomes a trunk when 2 or more VLANs are configured on this port, in the same or different
forwarding domains.
When trunks are configured on a FortiGate, it is essential to create forwarddomains, in
order to avoid packets looping back on the VLANs of the trunk. This will confine all
broadcasts and multicast traffic between the interfaces belonging to a same forward
domain.
24 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

Networking in Transparent Mode Inter-VDOM links between NAT/Route and Transparent VDOMs
In the case where a trunk port is configured with a VLAN in a different forwarding domains, the MAC address of
the network device connected to this port learns the FDB of each forwarding domain. This is Independent VLAN
Learning (IVL).
VLAN Translation
The same forwarding domain can include several different VLANs. Therefore, a frame ingressing an interface
with a certain VLAN ID can be forwarded to another port with another VLAN ID. This is sometimes referred as
VLAN translation.
Inter-VDOM links between NAT/Route and Transparent VDOMs
Inter-VDOM links between NAT/Route and Transparent mode VDOMS can be useful for configurations where the
NAT/Route VDOMs that share a common Internet service route, which can be routed through a Transparent
VDOM that provides additional functionality, like common Security inspection, WAN optimization, explicit
proxying and so on.
Other examples include:
l Performing SSL offloading in the Transparent mode VDOM and providing Internet access through a NAT/Route
mode VDOM.
l Applying WAN optimization in a Transparent mode VDOM and other security features in the NAT/Route mode
VDOM.
l Using a dedicated Transparent mode VDOM for the explicit web proxy in front of a NAT/Route mode VDOM that
applies other security features.
l An ISP configuration with multiple per-tenant NAT/Route mode VDOMs all sharing a single Internet connection but
where the ISP only presents a single routed subnet. Each tenant can then be assigned an IP from the subnet for
their respective VDOM link interface while using a single physical port to connect to the ISP router.
For more information about inter-VDOM links, please refer to the Virtual Domains handbook.
Replay Traffic Scenario
Situations can arise where an identical TCP packet enters twice the FortiGate via 2 different ports. This can be
due to a firewall or other network device redirecting packets out on the same port it has received it.
The FortiGate will in this condition detect a replay packet and drop it.
If the network topology or culprit devices cannot be changed to avoid this, the workaround on the FortiGate can
be to disable TCP replay verification packets.
config system global
set anti-replay | loose | strict | disable |
end
The debug flow diagnosis output hereafter shows the message indicating this condition:
id=20085 trace_id=179 msg="vd-VDOM_VLAN1 received a packet(proto=6, 10.10.253.9:10709
>10.10.248.5:25) from TO_EXTERNAL ."
id=20085 trace_id=179 msg="Find an existing session, id-00041475, original direction"
id=20085 trace_id=179 msg="replay packet, drop "
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
25

Packet Forwarding using Cisco Protocols Networking in Transparent Mode
For additional diagnosis and troubleshooting procedures,go to http://kb.fortinet.com .
Packet Forwarding using Cisco Protocols
In order to pass Cisco Discover Protocol (CDP) or Cisco VLAN Trunk Protocol (VTP) packets through a FortiGate
in Transparent mode, the parameter stpforward must be applied on the port configuration. VTP and CDP
packets are sent to the destination MAC address 01-00-0C-CC-CC-CC.
A Cisco NATIVE VLAN carries CDP/VTP frames. The frames of this VLAN must be
received on the FortiGate physical interfaces (not VLAN sub-interface). Physical
interfaces are the only ones that can send/accept non-tagged packets.
The example below will allow CDP and VTP packets to be sent from port3 up to the Remote unit, through two
VDOMs, via one physical port and three port aggregations.
Port and Port aggregation configuration:
config system interface
edit "port1"
set vdom "VD1"
next
edit "port2"
set vdom "VD1"
next
edit "port3"
set vdom "VD1"
set stpforward enable
next
edit "port5"
set vdom "VD3"
next
edit "port6"
set vdom "VD3"
next
edit "port17"
set vdom "VD2"
next
edit "port18"
set vdom "VD2"
next
edit "port19"
set vdom "VD2"
next
edit "port20"
set vdom "VD2"
next
edit "LACP_VD2_IN"
set vdom "VD2"
set stpforward enable
set type aggregate
set member "port17" "port18"
next
26 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

