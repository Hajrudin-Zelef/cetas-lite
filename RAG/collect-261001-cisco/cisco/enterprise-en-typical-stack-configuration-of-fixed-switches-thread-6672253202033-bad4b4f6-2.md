---
id: collect-261001-cisco/cisco/enterprise-en-typical-stack-configuration-of-fixed-switches-thread-6672253202033-bad4b4f6-2
title: "enterprise-en-typical-stack-configuration-of-fixed-switches-thread-6672253202033-bad4b4f6"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/enterprise-en-typical-stack-configuration-of-fixed-switches-thread-6672253202033-bad4b4f6.md
source_anchor: ""
source_lines: [55, 75]
sha256: dc813da1a7f3910a90b27fd3f16a6e9d19c981d778c92626c29117300b6266e4
---

# enterprise-en-typical-stack-configuration-of-fixed-switches-thread-6672253202033-bad4b4f6

In V200R008C00 and earlier versions, you can configure a maximum of 64 Eth-Trunks on a relay agent to provide the relay function for multiple stacks. This restriction does not apply to versions later than V200R008C00.
After multiple switches form a stack, the following features cannot be configured in the stack:
Y.1731 one- and two-way frame delay measurement
N:1 VLAN Mapping
IPv6 over IPv4 tunnel
IPv4 over IPv6 tunnel
E-Trunk
When you establish a stack on the switches that support both stack card connection and service port connection, such as S5720-C-EI, note the following:
All member switches must use the same stack connection mode.
When a member switch has stack cards installed and the service port stack configuration, the switch uses the service port connection mode to establish a stack. It does not use the stack card connection mode even though a stack fails to be established in service port connection mode and stack cards are connected correctly.
A switch uses the stack card connection mode to establish a stack only when it has no service port stack configuration.
If a switch is currently using the stack card connection mode, perform the service port stack configuration on the switch before changing the stack connection mode to service port connection. After the service port stack configuration is complete, the switch uses the service port connection mode when restarting.
If a switch using the stack card connection mode has service port configuration, a smooth upgrade cannot be performed on the switch.
If a switch is currently using the service port connection mode, correctly connect stack cards and stack cables and clear the existing service port stack configuration before changing the stack connection mode to stack card connection. You can use the reset stack-port configurationcommand to clear the existing service port stack configuration.
When changing service port connection to stack card connection, you are advised to remove the cables connected to service ports to prevent loops.
Connect a stack to other network devices using an Eth-Trunk and add one port of each member switch to the Eth-Trunk.
When a stack connects to access devices, configure ports directly connected to terminals as STP edge ports to prevent STP re-calculation when the ports alternate between Up and Down states. This configuration ensures normal traffic forwarding.
If storm control needs to be configured on many ports, replace storm control with traffic suppression to save CPU resources.
If port security needs to be configured on many ports, replace port security with MAC address learning limiting to save CPU resources.
Loops may occur on a network to which a stack connects. Run the mac-address flapping action error-down command to set an interface to the error-down state when MAC address flapping is detected on the interface. This improves system processing performance and allows the peer device to detect that the interface becomes Down. Additionally, if the peer device has redundant links, traffic can be rapidly switched to a normal link. Hope to help you.
See more please click
