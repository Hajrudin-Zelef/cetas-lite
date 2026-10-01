---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-27384925962647-managing-broadcast-traffic-with-unifi-b3eb3117
title: "hc-en-us-articles-27384925962647-managing-broadcast-traffic-with-unifi-b3eb3117"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["latency", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-27384925962647-managing-broadcast-traffic-with-unifi-b3eb3117.md
source_anchor: ""
source_lines: [1, 47]
sha256: aa0a92606b099c3b4c9263e428f019c388123c6e7508c00bb333b35472bdfab2
---

# hc-en-us-articles-27384925962647-managing-broadcast-traffic-with-unifi-b3eb3117

Managing Broadcast Traffic with UniFi
Excessive broadcast and multicast traffic can significantly reduce network performance, particularly on wireless networks. This article explains what it is, its impact on UniFi deployments, and effective management strategies.
What Is Broadcast and Multicast Traffic?
Broadcast traffic:
- Sent to all devices on the local network (subnet).
- Common examples: ARP and DHCP discovery.
Multicast traffic:
- Sent to a specific group of devices.
- Used for:
  - IPTV and multimedia streaming
  - Device and service discovery (e.g., AirPlay, mDNS, Bonjour)
Unlike unicast (one-to-one), both broadcast and multicast consume more airtime, especially on wireless networks, as they are transmitted at the lowest supported data rate. More devices transmitting more data at slower rates can drastically reduce network performance.
Why It Matters: Airtime & Broadcast Storms
Broadcast and multicast traffic can:
- Consume excessive wireless airtime: Even small volumes of traffic (like mDNS) can use a majority of airtime. A few hundred Kbps of mDNS traffic can be the equivalent of a few hundreds Mbps of unicast traffic.
- Trigger broadcast storms: Incorrectly connected gateways, loops, or misconfigured devices can cause storms that crash the entire network.
- Cause subtle, harder to diagnose issues:
  - Intermittent AP connectivity
  - Low client throughput
  - High client latency or timeouts
Tip: If you're seeing poor WiFi performance, check your airtime utilization in Radios > AirView.
How to Manage Broadcast/Multicast Traffic
- Enable Multicast to Unicast
- Converts multicast to unicast for better performance in most networks.
- Best for: Small to medium networks.
- Avoid in: Large-scale deployments. Converting multicast to unicast may increase airtime if many clients are subscribed.
- How to enable: Settings > WiFi > Click on an SSID > Check ‘Multicast to Unicast’
- Enable VLAN Isolation for IoT & Media Devices
- Put devices like printers, AirPlay speakers, or IP cameras on a dedicated VLAN.
- Keeps discovery/multicast chatter localized.
- Use mDNS relaying only where needed.
- Example: Place smart speakers and casting devices on a "Media" VLAN and enable mDNS between that VLAN and required clients.
*You may also consider segmenting ProAV systems or SDVoE devices on their own VLANs using UniFi’s ProAV switch profiles to isolate high-bandwidth multicast traffic.
- Enable Multicast & Broadcast Control (SSID Setting)
-  
  - Restricts multicast/broadcast traffic to a list of approved MAC addresses.
  - Good for: Large/high-density environments.
  - Blocks: Cross-AP device discovery (e.g., Apple TVs won't work across APs).
  - Ensure: Allow your DHCP server's MAC address, or clients won't get IPs.
- Enable Device Isolation
- When multicast traffic still impacts airtime—especially in large or dense environments—isolating client devices from one another can help. This prevents direct device-to-device communication, including discovery and casting protocols. Use this as a last resort when other optimizations aren’t enough.
  - Apply at the SSID level:
Go to Settings > WiFi > Select your SSID > Enable 'Client Device Isolation'
  - Apply at the switch level (ACL):
Go to Settings > Networks > 'Device Isolation (ACL)' > Select your VLAN
- Apply at the SSID level:
Note: Switch-level Device Isolation (ACL) requires all UniFi switches to properly function. Additionally, Flex and Flex Mini switches do not support this functionality
