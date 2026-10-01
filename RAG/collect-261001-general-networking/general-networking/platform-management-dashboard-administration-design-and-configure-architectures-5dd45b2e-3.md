---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-5dd45b2e-3
title: "platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e"
domain: general-networking
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e.md
source_anchor: ""
source_lines: [81, 119]
sha256: 7204431b1b568224dab5b5bc7efcbd51c1f7543858ae05bf624c6c51e9e1f3b8
---

# platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e

Bridge mode is recommended to improve roaming for voice over IP clients with seamless Layer 2 roaming. In bridge mode, the Meraki APs act as bridges, allowing wireless clients to obtain their IP addresses from an upstream DHCP server. Bridge mode works well in most circumstances, particularly for seamless roaming, and is the simplest option to put wireless clients on the LAN. To configure the client IP assignment modes please refer to our document on SSID Modes for Client IP Assignment.
When using Bridge mode, all APs on the same floor or area should support the same VLAN to allow devices to roam seamlessly between access points. Using Bridge mode will require a DHCP request when performing a Layer 3 roam between two subnets. For example, when a user on a VoIP call roams between APs on different VLANs without layer 3 roaming, the user's session will be interrupted as the external server must re-establish communication with the client's new IP address. During this time, a VoIP call will noticeably drop for several seconds, providing a degraded user experience.
Large wireless networks with multiple VLANS per floor may require IP session roaming at layer 3 to enable application and session persistence while a mobile client roams across multiple VLANs. With layer 3 roaming enabled, a client device will have a consistent IP address and subnet scope as it roams across multiple APs on different VLANs/subnets. If Layer 3 roaming is required on your network, please refer to our article on Layer 3 Roaming.
Note: It is strongly recommended to consult with a Cisco Meraki SE or Cisco Partner when considering layer 3 roaming options.
NAT mode is not recommended for Voice over IP: With NAT mode enabled, your devices will request a new DHCP IP address on each roam. Moving between APs in NAT mode will cause the connection to break when moving AP to AP. Applications requiring continuous traffic streams such as VoIP, VPN or media streams will be disrupted hen roaming between APs.
Segregate Traffic on a Voice VLAN
Voice traffic tends to come in large amounts of two-way UDP communication. Since there is no overhead on UDP traffic ensuring delivery, voice traffic is extremely susceptible to bandwidth limitations, clogged links, or even just non-voice traffic on the same line. Separating out your voice traffic allows it to function independently of other network traffic, and allows for more granular control over different types of traffic.
If a voice VLAN is specified on a Meraki MS switch, the port will accept tagged traffic on the voice VLAN. In addition, the port will send out LLDP and CDP advertisements recommending devices use that VLAN for voice traffic. The VLAN tagged on the wireless access point should match the Voice VLAN on your wired network. For more information, please visit the article on Configuring MS Access Switch for Standard VoIP deployment.
Band Selection
The 2.4 GHz band has only 3 channels that do not overlap, while the 5 GHz band has up to 19 individual channels in the US. A wireless network will provide the best quality of service for wireless voice when designed correctly to support 5 Ghz coverage for voice. This can be configured under Access Control > Wireless options > Band selection > '5 GHz band only'. After configuration, testing should be performed in all areas of your environment. If you do not have proper 5 GHz coverage after a post-install site survey, you can manually increase the power on the 5 GHz radio in Radio Settings > Channel Planning.
If you do not have a dedicated Voice SSID, enable 'Dual-band with band steering' to enable your voice devices to use both 2.4 GHz channels and 5 GHz. Devices will be steered to use the 5 GHz band. For more details refer to the Band Steering Overview article. With a dual-band network, client devices will be steered by the network. However, to improve voice quality, please follow our guide on configuring wireless band preference on client devices.
If you have devices that only support a 2.4 GHz network such as 802.11bgn devices, please contact Meraki Support to enable a 2.4 GHz only network.
Minimum Bitrate
For Voice networks, 12 Mbps is recommended as the minimum bitrate. Increasing this value requires proper coverage in the RF planning. An administrator can improve the performance of clients on the 2.4 GHz and 5Ghz band by disabling lower bitrates. Adjusting the bitrates can reduce the overhead on the wireless network and also in some cases improve roaming performance.
Bandwidth Limits
Consider placing a per-client bandwidth limit on the rest of your network traffic. Prioritizing voice will have a greater impact if all other applications are limited. For more details refer to the article Configuring Bandwidth Limitations and Enabling Speed Burst on Wireless Networks.
- Go to Wireless > Configure > Firewall & traffic shaping and choose your SSID from the SSID drop down menu at the top of the screen.
- Set a 'Per-client bandwidth limit' to 5 Mbps with 'Speed Burst'. This will apply to all non-voice application traffic. This step in the guide is optional.
- Set a 'Per-SSID bandwidth limit' to unlimited.
SpeedBurst enables a bursts of four times the allotted bandwidth limit for five seconds.
Traffic Shaping Rules
Use traffic shaping to offer voice traffic the necessary bandwidth. It is important to ensure that your voice traffic has enough bandwidth to operate. As such, traffic shaping rules can be implemented to allow voice traffic to use additional bandwidth, or limit other types of traffic to help prioritize voice traffic.
- Go to Wireless > Configure > Firewall & traffic shaping and choose your SSID from the SSID drop down menu at the top of the screen.
- Click the drop down menu next to Shape traffic and choose Shape traffic on this SSID, then click Create a new rule.
- Click Add + and select 'All VoIP & video conferencing'
- Set Per-client bandwidth limit to 'Ignore SSID per-client limit (unlimited)' and click Save changes.
PCP, DSCP, and WMM mapping
Many devices support Quality of Service (QoS) tags to maintain traffic priority across the network. Meraki MR access points support WMM to improve the performance of real-time data such as voice and video. WMM improves the reliability of applications in progress by preventing oversubscription of bandwidth. WMM accomplished this by enhancing the prioritization of traffic using the access categories: voice, video, best effort, and background data. WMM has mapped values which can be found in this article. To configure correct mapping for PCP and DSCP values, the following steps should be followed:
- Go to Wireless > Configure > Firewall & traffic shaping and choose your SSID from the SSID drop down menu at the top of the screen.
- Under the traffic shaping rules, make sure Shape Traffic for this SSID is selected and that there's a rule for All voice & video conferencing.
- Set PCP to '6' or the setting recommended by your device/application vendor (Note that PCP values can only be changed if the SSID has VLAN tagging enabled. This ensures there's a field to which the CoS value can be written).
- Set DSCP to 46 (EF - Expedited Forwarding, Voice) or the setting recommended by your device/application vendor.
Note: The DSCP tag, 46 (EF - Expedited Forwarding, Voice) maps to WMM Access Category AC-VO for Voice, Layer 2 CoS 5)
If no DSCP values are configured, the default DSCP to WMM mapping will be used. The access point does the mapping between the LAN's Layer 2 priority and the radio's WMM class. Below is table showing the mapping between common traffic types and their respective markings:
| RFC 4594-Based Model | 802.3 DSCP | 802.3 DSCP [Decimal] | IEEE 802.11 Model [802.11e WMM-AC] | 
| Voice + DSCP-Admit | EF + 44 | 46 | Voice AC (AC_VO) | 
| Broadcast Video | CS5 | 40 | Video AC (AC_VI) | 
| Multimedia Conferencing | AF4n | 34, 36, 38 | Video AC (AC_VI) | 
| Realtime Interactive | CS4 | 32 | Video AC (AC_VI) | 
