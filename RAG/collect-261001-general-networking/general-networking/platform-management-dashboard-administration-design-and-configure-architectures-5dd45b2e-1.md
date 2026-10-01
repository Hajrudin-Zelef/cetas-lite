---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-5dd45b2e-1
title: "platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Microsoft"]
dates: []
keywords: ["ethernet", "voice"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e.md
source_anchor: ""
source_lines: [1, 30]
sha256: a6408666f43ebb922469fde86c1d10e1df225985e497aa202ee7b5390b3d5eb9
---

# platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e

Wireless VoIP QoS Best Practices
Click 日本語 for Japanese
This article is a guide to optimize Quality of Service (QoS) for wireless Voice over IP applications on Meraki MR wireless access points. Voice over IP (VoIP) has replaced telephones in enterprise networking with IP-based phones. While the majority of desk phones using VoIP require ethernet, there are many voice applications and wireless VoIP phones that operate over WiFi.
The Meraki MR series of WiFi access points have been tested by Cisco Meraki to provide the highest quality VoIP experience when using Cisco Jabber, Microsoft Lync, Microsoft Skype for Business, Broadsoft, Cisco 7900 Series phones, SpectraLINK phones, Ascom phones, Cisco phones, and Apple iPhones. This guide will provide recommendations for optimizing voice quality followed by product specific recommendations.
Measuring Voice Quality
By following this guide, you can significantly improve quality of service for the wireless voice applications and reduce or eliminate dropped calls, choppy speech, fuzzy speech, buzzing, echoing, long pauses, one-way audio, and issues while roaming between access points.
To develop this guide, we performed testing using Microsoft Lync's Pre-Call Diagnostics Tool. The endpoints used during testing were Macbook Pros running Office 365's Cloud-hosted Skype for Business Online, also known as Lync Online. All tests were performed while connected to an MR32 access point inside Meraki's headquarters in San Francisco, a high density corporate WiFi network. This tool measures 3 key metrics for voice quality:
- Network MOS - The Network Mean Opinion Score (MOS) is the network’s impact on the listening quality of the VoIP conversation. The score ranges from 1 to 5, with 1 being the poorest quality and 5 being the highest quality.
- Packet Loss Rate - The packet loss rate is the percent of packets that are lost during transmission.
- Interarrival Jitter - Interarrival jitter measures the variation in arrival times of packets being received in milliseconds (ms).
By combining this guide with best practices for configuring your client device configuration, application servers, WAN links, and wired network, you can measure and improve quality voice end-to-end. For more information on configuring your wired network to support Voice, please visit the article on Configuring MS Access Switch for Standard VoIP deployment article.
Voice Quality Before this Guide
With the default settings on the MR, we see the baseline for quality. Voice calls with Lync on this network would be acceptable to some users, but not acceptable to others. The results of the Lync testing show that the Network Mean Opinion Score (MOS) drops below 3.5. Values values dropping below 3.5 are termed unacceptable by many users. The packet loss jumps to 8 percent during a period of network congestion simulated by running a speed test. Jitter fluctuates from 12 milliseconds to over 36 milliseconds. Cisco recommends a target of 10 ms of jitter and no more than 50 ms of jitter. Jitter is handled using buffering in voice/video applications, adding a small delay. The human ear normally accepts up to about 140 milliseconds of delay without noticing it.
Voice Quality After this Guide
After making the changes exactly as described in this guide, we see a significant improvement in voice quality. The MOS score approaches 3.9, the packet loss is near zero, and the jitter is consistently below 6 milliseconds and always below 12 milliseconds. As the call starts, traffic shaping kicks in automatically with prioritization and QoS tagging. A speed test was performed with no impact to the voice traffic - packet loss increased to 0.5% during the congestion.
Wireless Voice Best Practices
A Cisco Meraki wireless network has the intelligence built-in with deep packet inspection to identify voice and video applications and prioritize the traffic using queuing and tagging to inform the rest of the network how to handle your voice traffic. Below is a summary of the best practices to provide the best voice quality over wireless.
- Perform a pre-install RF survey for overlapping 5 GHz voice-quality coverage with -67 dB signal strength in all areas.
- If possible, create a new SSID dedicated to your voice over IP devices.
    
  - Set Authentication type to 'Pre-shared key with WPA2'
  - Set WPA encryption mode to 'WPA2 only'
  - Enable '5 GHz band only'.
  - Set Minimum bitrate to '12 Mbps' or higher.
- Enable Bridge mode. If you cannot provide a VLAN across the entire floor, use Layer 3 roaming.
- Enable 'VLAN tagging' and assign a VLAN dedicated to wireless voice. If you cannot dedicate an SSID to voice, assign a VLAN dedicated to wireless.
- Set a 'Per-client bandwidth limit' to 5 Mbps with 'Speed Burst' to limit all non-voice traffic
- Set a 'Per-SSID bandwidth limit' to unlimited.
- Enable 'Traffic shaping' on the SSID to prioritize all voice traffic
    
