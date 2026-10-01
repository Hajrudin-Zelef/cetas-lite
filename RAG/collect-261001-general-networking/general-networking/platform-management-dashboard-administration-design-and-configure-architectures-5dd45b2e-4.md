---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-5dd45b2e-4
title: "platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Microsoft"]
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e.md
source_anchor: ""
source_lines: [120, 181]
sha256: 0ffe3652c29e9db7c72837be9b46e93a9ab474cc486b110964c75d677da7a244
---

# platform-management-dashboard-administration-design-and-configure-architectures--5dd45b2e

| Multimedia Streaming | AF3n | 26, 28, 30 | Video AC (AC_VI) | 
| Signaling | CS3 | 24 | Video AC (AC_VI) | 
| Transactional Data | AF2n | 18, 20, 22 | Best Effort AC (AC_BE) | 
| OAM | CS2 | 16 | Best Effort AC (AC_BE) | 
| Bulk Data | AF1n | 10, 12, 14 | Background AC (AC_BK) | 
| Scavenger | CS1 | 8 | Background AC (AC_BK) | 
| Best Effort | DF | 0 | Best Effort AC (AC_BE) | 
* n as used in place for the drop indication of assured forwarding matches values 1-3.
For QoS prioritization to work end to end, ensure that upstream networking equipment supports QoS prioritization as well. The PCP and DSCP tags applied on the wireless access point should match the wired network configuration to ensure end-to-end QoS. For more information, please visit the article on Configuring MS Access Switch for Standard VoIP deployment article.
Custom Traffic Shaping
If your voice traffic does not match the built-in application signatures or is not listed, you can create your own signature for traffic shaping.
- Add the IP and ports used by your servers hosting Microsoft Lync / Skype for Business, Jabber, Spark, or other voice application.
    
  - In the Definition field click Add + and Custom expressions
  - In the text field, enter the IP address of each of your voice servers for example 172.16.1.123 or a range of server IPs 172.16.1.0/24
  - Also, add your servers as source addresses by using the CIDR notation localnet:172.16.1.123/32 for an individual server, or localnet:172.16.1.0/24 for a range of IPs.
  - Click the Add + button again when finished.
- If you have a dedicated voice SSID and dedicated voice VLAN add the local subnet of the client devices.
    
  - In the Definition field click Add + and Custom expressions
  - In the text field, enter localnet:192.168.0.1/16 indicating the source subnet of your client devices in CIDRnotation.
  - Click the Add + button again when finished.
- Set Per-client bandwidth limit to 'Ignore SSID per-client limit (unlimited)' and click Save changes.
Product Specific Recommendations
Cisco Meraki works closely with device manufacturers, for example Apple, to provide them with their own access points for interoperability testing. Meraki performs our own testing across the entire spectrum of devices and our customer support team handles and reports bugs quickly. This section will provide recommendations based on real-world deployments by Meraki customers combined with the best practices developed by Meraki and the vendors mentioned below.
Microsoft Lync / Skype for Business
This section will provide guidance on how to implement QoS for Microsoft Lync and Skype for Business. Microsoft Lync is a widely deployed enterprise collaboration application which connects users across many types of devices. This poses additional challenges because a separate SSID dedicated to the Lync application may not be practical. When you install Microsoft Lync Server / Skype for Business, Quality of Service (QoS) will not be enabled for any devices used in your organization that use an operating system other than Windows. For more guidance on deploying Lync over Wi-Fi, please read Microsoft's deployment guide, Delivering Lync 2013 Real-Time Communications over Wi-Fi.
Meraki's deep packet inspection can intelligently identify Lync calls made on your wireless network and apply traffic shaping policies to prioritize the Lync traffic - using the SIP Voice protocol. In addition to the Meraki built-in signatures for Skype and SIP, you should also identify each Lync server by IP and any custom ports used by your Lync clients or servers. Follow these steps to configure your traffic shaping rules for Lync / Skype.
- Go to Wireless > Configure > Firewall & traffic shaping and choose your SSID from the SSID drop down menu at the top of the screen.
- Click the drop down menu next to Shape traffic and choose Shape traffic on this SSID, then click Create a new rule.
- Consider setting a 'Per-client bandwidth limit' to 5 Mbps with 'Speed Burst'. This will apply to all non-voice application traffic
- Set Per-client bandwidth limit to unlimited
- Create a traffic shaping rule for All voice & video conferencing > 'Skype' and 'SIP (Voice)'
- Set the Per-client bandwidth limit to 'Ignore SSID per-client limit (unlimited)' from the drop down.
- Set PCP to '6'. The 802.1p parameter is no longer supported in Lync Server 2013. The parameter is still valid for backward compatibility with Microsoft Lync Server 2010; however, it has no effect on devices used with Lync Server 2013.
- Set DSCP to '46 (EF - Expedited Forwarding, Voice)'
- Add 'Custom expressions' for the IP and ports used by your servers hosting Microsoft Lync / Skype for Business
    
  - For Cloud-hosted Lync / Skype, add the domain names from the table below
  - Add the Port numbers from the table below or your own list of assigned port numbers
  - Add the IP address of each of your on-premise Lync servers
- You can test that DSCP markings are applied using the Meraki packet capture tool.
| Lync Online / Skype for Business Online servers | Lync & Skype port numbers | On Premise Lync / Skype for Business Servers | 
|  |  |  | 
The ports provided in the above table are the standard ports provided by Microsoft. Enabling QoS Configuration of the client device to modify the port ranges and assign the DSCP value 46. Microsoft's best practices include configuring the port ranges on both your servers and client devices. For details on enabling QoS, refer to Microsoft's article Managing Quality of Service (QoS) in Lync Server 2013.
Cisco 7925G Phones
Cisco 7925G, 7925G-EX, and 7926G VoIP phones require specific settings to inter-operate with Meraki MR access points configured with WPA2-PSK association requirements. For more in-depth information on integrating Cisco 792xG with MR Access Points, see the Cisco Unified Wireless IP Phone 792xG + Cisco Meraki Wireless LAN Deployment Guide.
- Cisco recommends to set band selection to 5 GHz only.
- Do not utilize the Dual band operation with Band Steering option. If the 2.4 GHz band needs to be used due to increased distance, select Dual band operation (2.4 GHz and 5 GHz) should be selected.
- Set the Minimum bitrate to 11 Mbps or higher.
- By default, Cisco Meraki access points currently tag voice frames marked with DSCP EF (46) as WMM UP 6 and call control frames marked with DSCP CS3 (24) as WMM UP 4.
Cisco Meraki access points will trust DSCP tags by default. Administrators should ensure that upstream QoS is in place and that the QoS markings outlined below are in place for the 7925 phones. To rewrite QoS tags for certain traffic types or source/destination, then create a traffic shaping rule as outlined in Custom Traffic Shaping above.
Below is the QoS and port information for voice and call control traffic used by the Cisco Unified Wireless IP Phone 7925G, 7925G-EX, and 7926G. For a full list of ports and protocols used by Cisco phones refer to the Cisco Unified Communications Manager TCP and UDP Port Usage Guide.
| Traffic Type | DSCP | PCP (802.1p) | WMM | Port Range | 
| Voice (RTP, STRP) | EF (46) | 5 | 6 | UDP 16384 - 32767 | 
| Call Control (SCCP, SCCPS) | CS3 (24) | 3 | 4 | TCP 2000, TCP 2443 | 
Cisco Unified Communications Manager only uses the port range 24576-32767 although other devices use the full range 16384 - 32767.
Apple iPhones
Apple and Cisco have created partnership to better support iOS business users by optimizing Cisco and Meraki networks for iOS devices and apps. For more information on this partnership, please see Apple's website. Meraki's group policies can be easily configured to optimize Apple devices on a Meraki network. First create a group policy you would like to apply to Apple devices.
- Browse to Network-wide > Configure > Group Policies, scroll down and click Add Group.
- Under traffic shaping, create a new rule, and add All VoIP & video conferencing
    
