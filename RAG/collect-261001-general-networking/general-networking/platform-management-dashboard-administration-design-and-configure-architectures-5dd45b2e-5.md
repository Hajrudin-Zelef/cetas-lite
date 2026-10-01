---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-5dd45b2e-5
title: "platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["latency", "voice"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e.md
source_anchor: ""
source_lines: [182, 209]
sha256: 0763f49769f327307b4f7f6bddaedbebb9e28aaceed16aeeb9e01da71b301e42
---

# platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e

  - Set the Per-client bandwidth limit to Ignore network per-client limit
  - Set PCP to 6 (highest priority) and DSCP to 46 (EF - Expedited Forwarding, Voice).
- Create a second rule, and add All Video & music 
    
  - Set the Per-client bandwidth limit to Ignore network per-client limit
  - Set PCP to 5 and DSCP to 34 (AF41- Multimedia Conferencing, Low Drop).
To apply your new Apple device group policy, browse to Wireless > Access Control and enable Assign group policies by device type. Click Add group policy for a device type for each Apple device type (iPhone, iPad, iPod, and Mac OS X) and assign the Apple device group policy you created. Click save and your optimization is complete.
Vocera Badges
Vocera badges communicate to a Vocera server, and the server contains a mapping of AP MAC addresses to building areas. The server then sends an alert to security personnel for following up to that advertised location. Location accuracy requires a higher density of access points. In a high density deployment, you may need to reduce the transmit power of each AP manually to as low as 5 dB on all supported radios. Vocera provides additional documentation on deploying WLAN best practices to support Vocera badges. For more information download their document on Vocera WLAN Requirements and Best Practice
Some models of Vocera badges do not support 5 GHz or WPA2 AES encryption and require WPA1 TKIP. Please contact Cisco Meraki support to configure a WPA1 TKIP on your network.
Service Provider WiFi
Service providers are using WiFi to offload data from cellular networks to meet the ever-increasing demands of mobile device users. Two technologies enabling WiFi to meet this demand are WiFi calling and Hotspot 2.0.
WiFi Calling
Mobile network operators (MNOs) now allow their customers to place phone calls over Wi-Fi to save roaming costs and leverage WiFi coverage in buildings with poor cellular coverage. WiFi calling is now supported by majority of mobile devices and MSPs . Apple maintains a list of carriers that support WiFi Calling, and provides a user guide for this feature.
An Enterprise WiFi infrastructure handles more than just carrier voice traffic, and this limited spectrum is shared by other applications and services like Video streaming & Web Conferencing. Voice communication requires low latency and minimal jitter, so it's essential to design a network with end-to-end Quality of Service (QoS) and voice optimizations. This ensures WiFi calling packets are delivered efficiently, even when other applications are running.
Requirements
- 
    Your mobile carrier must support Wi-Fi Calling (see above)
- 
    Your phone must support Wi-Fi Calling
Note: Wi-Fi Calling requires an IPSec tunnel established between the phone and the carrier. Authentication occurs with a key exchange using details on the device SIM card, and a TLS certificate handshake.
You will need to ensure that connectivity toward TCP/443 (HTTPS/TLS), UDP 500 and UDP 4500 (ESP) is allowed across your network and upstream. If you need to open connections to specific external IPs, your MSP should be able to provide a list or FQDNs to match against.
Once the tunnel is established, you should see ESP traffic between the carrier and phone when examining a packet capture on the wired interface of the AP:
Hotspot 2.0
Hotspot 2.0 also known as Passpoint is a service provider feature that assists with carrier offloading. Part of the 802.11u amendment to the 802.11 standard additional information is included in Hotspot 2.0 configured SSIDs that Hotspot 2.0 client devices can analyze use to determine if it is able to join the network automatically.
Managed Service Providers (MSPs) can now take enable of Hotspot 2.0 options on the Cisco Meraki MR access points. Meraki allows MSPs to customize the Hotspot 2.0 SSID advertisements to allow their subscribers to easy roam between networks. Hotspot 2.0 options are only available to qualified Managed Service Providers. Please contact Cisco Meraki Support to check eligibility.
Troubleshooting VoIP
We have created a detailed article focused on troubleshooting VoIP on Meraki. Please visit the article: VoIP on Cisco Meraki: F.A.Q. and Troubleshooting Tips
