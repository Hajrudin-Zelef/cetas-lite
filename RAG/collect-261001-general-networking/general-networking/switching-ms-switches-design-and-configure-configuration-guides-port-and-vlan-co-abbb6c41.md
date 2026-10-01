---
id: collect-261001-general-networking/general-networking/switching-ms-switches-design-and-configure-configuration-guides-port-and-vlan-co-abbb6c41
title: "switching-ms-switches-design-and-configure-configuration-guides-port-and-vlan-co-abbb6c41"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "voice"]
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-design-and-configure-configuration-guides-port-and-vlan-co-abbb6c41.md
source_anchor: ""
source_lines: [1, 15]
sha256: 8f6fc6e8bc8ee686c42d5a93e90e06f89bbe5ad64476ec7b9f698fd0be0173ec
---

# switching-ms-switches-design-and-configure-configuration-guides-port-and-vlan-co-abbb6c41

Configuring the MS Access Switch for Standard VoIP Deployments
Standard Data/VoIP deployments commonly utilize a three port switch built in to the VoIP phone to connect a workstation and phone to the same switch port:
For these deployments the MS Access Switch should be configured with Voice VLANs and QoS (Quality of Service) to separate voice traffic into its own broadcast domain and tag it for optimal transfer and prioritization. The MS Access Switch utilizes LLDP (Link Layer Discovery Protocol) to recognize the VoIP phone connected and retrieve various properties of the phone. QoS on the MS Switch tags specific traffic that is incoming on the desired switch ports and prioritizes traffic which is important for voice traffic due to its sensitivity to latency. This article describes how to configure the MS Access Switch using these features to optimize VoIP performance and minimize poor call quality:
- Configure the desired switch port(s) as access port(s)
- Configure data and voice VLANs
- Configure QoS settings
Configuring Access Ports
- Navigate to Configure > Switch ports and choose the specific port to be modified.
- Using the Type drop down, change the port to access:
Configure Data and Voice VLANs
Once the port has been changed to an access port, enter the data and voice VLANs:
Configure QoS settings
QoS settings can be enabled on dashboard under Configure > Switch Settings, in the Quality of service section:
This rule tells the MS Access Switch to place a class tag on the specific voice packets that enter it. See this knowledge base article for more information on QoS. This tag will follow the voice packets to other switches. Below is an example of a configured rule under the QoS section. VLAN 20 is the voice VLAN for this network:
These configurations ensure the switch will correctly prioritize voice traffic as it traverses the network and reduce the chance of voice data being forwarded incorrectly.
