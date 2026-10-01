---
id: collect-261001-huawei/huawei/guides-huawei-huawei-qinq-configuration-0c985c03
title: "guides-huawei-huawei-qinq-configuration-0c985c03"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/guides-huawei-huawei-qinq-configuration-0c985c03.md
source_anchor: ""
source_lines: [1, 14]
sha256: 8f3ede8c19b8b25deb88fb6d1db1e5262d43590815cf9005f0ccfeed992e7af0
---

# guides-huawei-huawei-qinq-configuration-0c985c03

QinQ is an open standard (IEEE 802.1ad) which allows a layer-2 frame to be provided with an additional VLAN tag. As an ISP you can provide a layer-2 extension over your network independently of the customer's VLAN concept.
In the following example, the customer wants to establish a layer-2 connection in VLAN20 between his CUST-1 and CUST-2 switches. His ISP can offer him such a connection, but has already used VLAN20 in its network. As the customer cannot/will not change the VLAN20 either, the ISP is using the QinQ solution. In its ISP backbone, it has reserved VLAN100 for this customer, which is used for transmission between the two switches. This is done by attaching a second VLAN tag 100 to the front of each frame from the customer. This VLAN tag 100 is only used within the ISP network and is removed again on the other side. This means that the customer never sees this VLAN tag, for him it looks as if both switches (Cust-A and Cust-B) are directly connected to each other.
The drawing shows the setup with the QinQ configuration between the two ISP switches (shown in red) and the corresponding configurations on the switches.
Huawei QinQ configuration
On the customer switches it is a standard configuration with an access port for the client and a trunk port in the direction of the ISP.
CUST-1 and CUST-2:
interface GigabitEthernet0/0/1description Trunk to ISPport link-type trunkport trunk allow-pass vlan 20!interface GigabitEthernet0/0/2description Client Portport link-type accessport default vlan 20stp edged-port enable
On the ISP switches, a dot1q tunnel in VLAN100 is configured on the uplink to the customer switches and the QinQ configuration on the link to the other ISP switch.
ISP1- and ISP-2:
interface GigabitEthernet0/0/1 port link-type dot1q-tunnel port default vlan 100!interface GigabitEthernet0/0/2 port link-type trunk port trunk allow-pass vlan 100 qinq protocol 9100
After these configurations on the switches, the two clients can reach each other, which can be checked with a ping.
Huawei QinQ Wireshark capture
If a Wireshark recording is made on the link between the two ISP switches, the ping packets between the clients can be analyzed. It is easy to see that two VLAN tags are present. The outer VLAN100 is the VLAN that the ISP adds and removes on the other side. The inner VLAN20 is the original VLAN tag from the client:
Leave a Comment
