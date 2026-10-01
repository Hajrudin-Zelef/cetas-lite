---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-5dd45b2e-2
title: "platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["latency", "voice"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e.md
source_anchor: ""
source_lines: [31, 80]
sha256: d66ae367a9a8e0c164d6c69996176eced2909e5ac33ad90414b0d6ec5e8fcd5a
---

# platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e

  - Create a traffic shaping rule for 'All voice & video conferencing'
  - Add 'Custom expressions' for the IP and ports used by your servers hosting Microsoft Lync / Skype for Business, Jabber, or Spark
  - Set the Per-client bandwidth limit to 'Ignore SSID per-client limit (unlimited)' from the drop down.
  - Set PCP to '6'
  - Set DSCP to '46 (EF - Expedited Forwarding, Voice)'
- Verify the Voice VLAN is tagged correctly
- Verify your uplinks and Switches have Quality of Service defined with the maximum PCP and DSCP values
- Verify DSCP trust is enabled on switch ports to APs and uplinks
- Verify your Windows Group Policy to ensure your devices are tagging application traffic with DSCP (not on by default)
- Verify your voice server configuration to ensure Microsoft Lync / Skype and Call Manager have DSCP enabled (not on by default)
Summary of 802.11 Standards
All Meraki MR series access points support the most recent 802.11 standards implemented to assist devices to roam between access points and ensure voice calls maintain a quality user experience.
- 802.11r: Fast BSS transition to permit fast and secure hand-offs from one access point to the other in a seamless manner
- 802.11i: Enabling client devices authenticated via 802.1X to authenticate with decreased latency whilst roaming
- 802.11k: assisted roaming allows clients to request neighbor reports for intelligent roaming across access points.
- 802.11e: Wireless Multimedia Extensions (WMM) traffic prioritization ensures wireless VoIP phones receive higher priority.
- WMM Power Save: maximizes power conservation and battery life on devices without sacrificing Quality of Service.
- 802.11u: Hotspot 2.0 also known as Passpoint is a service provider feature that assists with carrier offload.
Pre-Install Survey
The design and layout of access points is critical to the quality of voice over WiFi. Configuration changes cannot overcome a flawed AP deployment. In a network designed for Voice, the wireless access points are grouped closer together and have more overlapping coverage, because voice clients should roam between access points before dropping a call. Designing with smaller cells and lower power settings on the access point are key elements to ensure the overlapping coverage from neighboring APs/cells. Set a clear requirement based on the device type when performing a survey.
Pre-site surveys are useful for identifying and characterizing certain challenging areas and potential sources for interference, such as existing WiFi networks, rogues, and non-802.11 interference from sources such as microwave ovens and many cordless telephones. Post-site surveys should be performed at least 48 hours after installation to allow the network to settle on channel and power settings.
- Prefer 5 GHz coverage for voice applications due to the lower noise floor compared to 2.4 GHz
- Verify an AP can be seen from the phone at -67 dBm or better in all areas to be covered
- Verify that the AP sees the phone at -67 dBm or better in all areas as well
- Signal to Noise Ratio should always 25 dB or more in all areas to provide coverage for Voice applications
- Channel utilization should be under 50%
For more guidelines on site surveys read our article on Performing a Wireless Site Survey. For more detailed guidelines on designing RF specifically for Cisco Voice over WiFi please read Cisco's Voice over WLAN Guide.
Network Configuration
Making the changes described in this section will provide a significant improvement in voice quality and user satisfaction by following the best practices for configuring your SSIDs, IP assignment, Radio Settings, and traffic shaping rules.
Add a Dedicated Voice SSID
Voice optimization typically requires a different configuration including access control and traffic shaping to address device specific recommendations. You should create a separate Voice SSID for devices dedicated to voice applications. While this is not a requirement, we recommend to create a separate network to follow this guide. In networks with VoIP handsets from two different manufacturers, it is common to create two voice SSIDs.
If you plan to deploy more than 4 SSIDs please read our guide on the Consequences of Multiple SSIDs.
Authentication Type
Voice over WiFi devices are often mobile and moving between access points while passing voice traffic. The quality of the voice call is impacted by roaming between access points. Roaming is impacted by the authentication type. The authentication type depends on the device and it's supported auth types. It's best to choose the auth type that is the fastest and supported by the device. If your devices do not support fast roaming, Pre-shared key with WPA2 is recommended. WPA2-Enterprise without fast roaming can introduce delay during roaming due to its requirement for full re-authentication. When fast roaming is utilized with WPA2-Enterprise, roaming times can be reduced from 400-500 ms to less than 100 ms, and the transition time from one access point to another will not be audible to the user. The following list of auth types is in order of fastest to slowest.
- Open (no encryption)
- Pre-shared key with WPA2 and Fast roaming
- WPA2-Enterprise with Fast roaming
- Pre-shared key with WPA2
- WPA2-Enterprise
WPA2 only for Encryption Mode
Voice devices can benefit from having a single type of encryption used. By default, SSIDs on Cisco Meraki access points that are configured as WPA2 will utilize a combination of both WPA1 TKIP and WPA2 AES encryption. WPA2 (AES) is recommended and required in order to utilize caching or fast roaming. The WPA encryption setting is SSID specific, and can be found on the Wireless > Configure > Access control page.
- If all Voice devices support WPA2, the 'WPA2 only' option is recommended for Voice over IP devices.
- If the device does not support AES, it is also possible to force TKIP only. Please contact Cisco Meraki support to configure this option.
For step-by-step instructions on changing the WPA encryption mode, see our document on Setting a WPA Encryption Mode.
802.11r fast roaming
Enabling 802.11r is recommended to improve voice quality while roaming, especially when 802.1X is used for authentication. While PSK can benefit from 802.11r, there is typically less latency during the roam, as we are not waiting for a RADIUS response, and it is not always needed. The 802.11r standard was designed to improve VoIP and voice applications on mobile devices connected to Wi-Fi, in addition to or instead of cellular networks. When mobile devices roam from one area to another, they disassociate from one access point and reassociate to the next access point. Enabling 802.11r benefits VoIP devices by reducing the roaming time spent changing between access points. Some client devices are not compatible with Fast BSS Transition (802.11r). You may wish to check your devices for compatibility.
This feature can be enabled from the Configure > Access control page under Network access > 802.11r. If this option does not appear, a firmware update may be required. For more details on 802.11r, refer to our guide on Fast Roaming Technologies
It has been determined that configuring an SSID with WPA2-PSK and 802.11r fast roaming could pose a security risk due to a vulnerability. The vulnerability allows potential attackers the ability to obtain the PSK for the SSID when a client fast roams to another AP.
This vulnerability has been resolved in r25.7 and all modern firmware releases for the MR and CW platforms.
Layer 2 and Layer 3 Roaming
