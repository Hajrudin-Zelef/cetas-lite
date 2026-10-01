---
id: collect-261001-general-networking/general-networking/ms-port-and-vlan-configuration-determining-the-rstp-2f-2f-2f-2fstp-root-bridge-o-14367d85
title: "ms-port-and-vlan-configuration-determining-the-rstp-2f-2f-2f-2fstp-root-bridge-o-14367d85"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "training"]
source: docs/RAG/collect-261001-general-networking/ms-port-and-vlan-configuration-determining-the-rstp-2f-2f-2f-2fstp-root-bridge-o-14367d85.md
source_anchor: ""
source_lines: [1, 10]
sha256: 366771457b60662ddfd6e24c79441e098ec1e5d60c1573ea8752b3fd854bfcb1
---

# ms-port-and-vlan-configuration-determining-the-rstp-2f-2f-2f-2fstp-root-bridge-o-14367d85

Determining the RSTP/STP Root Bridge on an MS Switch Network
The Cisco Meraki MS series switch is compatible with industry standard STP (IEEE 802.1D) and RSTP (IEEE 802.1w). RSTP is used by default. Information on how to configure RSTP on a Cisco Meraki Switch can be found in the Configuring Spanning Tree on Meraki Switches (MS) article.
Learn more with this free online training course on the Meraki Learning Hub:
Overview
Confirming the correct root bridge election in a switched network is important for increasing the stability of a switched network. The root bridge is a central hub that can interconnects other legs of the network and can be exposed to heavy switching traffic. If the elected root bridge doesn't have the sufficient resources to handle the task,network performance may be affected. The root bridge is selected by manually configuring its bridge priority to a low value. 32768 is the default value out of a range from 0 to 61440. If all switches in a single spanning tree have the same bridge priority, the switch with the lowest MAC address will become the root bridge.
The diagram below is an illustration of a scenario in which the elected root bridge might not be the optimal choice:
In this diagram, Switches D and B comprise the core of the network, Switch A and C are access layer switches. Assuming all switches with their default bridge priority, because switch A has the lowest MAC address, it will be elected as a root bridge in this network . If the link between switch D and B is 1 gigabit and links between A and D as well as A and D are 100 Mbps, the Gigabit Ethernet link that connects switch D and swith B will be blocked, as a result of the election of Switch A as a root. This block causes all the data to flow via a 100-Mbps link across the access layer resulting in a non optimal layer 2 network.
Determination of the Root Bridge on MS Switches
To avoid the scenario above, it is important to determine what device in the switch fabric is acting as the Spanning-tree root bridge. In order to identify which MS switch is the Root Bridge, navigate to Switch > Monitor > Switches and select the desired switch. In the switch status area, take note of the field called RSTP Root as it will identify which switch is the Root Bridge.
Information on which switch is elected as the root bridge is accessible on the dashboard. Best practice dictates the root bridge to be the core of the network, this can be accomplished by manually setting the priority of the switch selected to be the root bridge. Information on enabling RSTP and setting the STP root bridge on a Cisco Meraki MS switch can be found in our Configuring Spanning Tree on Meraki Switches article.
