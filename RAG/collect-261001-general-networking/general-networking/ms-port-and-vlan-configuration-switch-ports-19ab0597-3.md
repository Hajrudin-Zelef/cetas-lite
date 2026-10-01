---
id: collect-261001-general-networking/general-networking/ms-port-and-vlan-configuration-switch-ports-19ab0597-3
title: "ms-port-and-vlan-configuration-switch-ports-19ab0597"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/ms-port-and-vlan-configuration-switch-ports-19ab0597.md
source_anchor: ""
source_lines: [130, 140]
sha256: be5dea6170739d7caaa8b1f30f6ac91c084df89b289ce97bab7085dfff8693dd
---

# ms-port-and-vlan-configuration-switch-ports-19ab0597

In the virtual stack, select the ports to be aggregated. Once the ports have been selected, choose Aggregate at the top or bottom of the port list and accept the change notification.
Note: Link Aggregation is supported on ports sharing similar characteristics such as link speed and media-type (SFP/Copper).
Splitting Aggregated ports
To split an aggregated link, simply select the aggregated port and choose Split. This will revert the changes and split the group into its own separate ports.
*For more specific configuration and interoperability information, please reference our documentation.
Port Mirroring
It may be necessary to configure a mirrored port or range of ports. This is often useful for network devices that require monitoring of network traffic, such as a VoIP recording solution or an IDS (Intrusion Detection System).
MS switches support one-to-one or many-to-one mirror sessions. Intra-stack port mirroring is available on our stackable switches. Only one active destination port can be configured per switch/stack.
In order to enable and configure a mirrored port or range of ports, navigate to Switching > Monitor > Switch Ports. On this page select the ports that are intended for mirroring and hit the Mirror button:
Next, enter the destination port for the mirror session. If the ports are in a switch stack then also select the desired switch in the stack for the mirror destination.
Once the Mirror is configured it can be easily identified using the Mirror column in Dashboard:
