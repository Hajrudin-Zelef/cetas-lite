---
id: collect-261001-general-networking/general-networking/2026-05-setup-iot-vlans-and-firewall-rules-with-html-c183dce4
title: "2026-05-setup-iot-vlans-and-firewall-rules-with-html-c183dce4"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["latency", "throughput"]
source: docs/RAG/collect-261001-general-networking/2026-05-setup-iot-vlans-and-firewall-rules-with-html-c183dce4.md
source_anchor: ""
source_lines: [1, 21]
sha256: 5f4b59bf443be14128bf8df003e80a95169a8f6618d390e94d42cab2638e6c75
---

# 2026-05-setup-iot-vlans-and-firewall-rules-with-html-c183dce4

VLAN Configuration and Network Segmentation
Defining a dedicated VLAN for Internet of Things (IoT) devices is the foundational step in securing a modern smart home. By isolating these often insecure devices from the primary data network, you minimize the lateral movement surface area for potential threats. Within the UniFi Controller, navigate to Settings > Networks and create a new network. Assign a descriptive name such as "IoT_VLAN" and set a unique VLAN ID (e.g., 20). Configure the Gateway IP/Subnet to a non-overlapping range, such as 10.0.20.1/24. Ensure that IGMP Snooping is enabled to manage multicast traffic effectively, which is critical for discovery protocols used by smart speakers and streaming devices.
Wireless Infrastructure and SSID Provisioning
Once the network is defined, a corresponding SSID must be created. Navigate to Settings > WiFi and add a new wireless network. Link this SSID directly to the IoT VLAN created in the previous step. It is recommended to use WPA2-PSK for maximum compatibility with older IoT sensors, though WPA3 should be tested if your hardware supports it. Important Warning: Note that SSID overrides per Access Point are no longer available in controller versions 6.0.23 and higher. Instead, UniFi utilizes AP Groups to manage which SSIDs are broadcast on specific hardware. Ensure you configure your AP Groups correctly before pushing changes to avoid disconnecting remote sensors.
Firewall Rule Implementation Strategy
The security of the IoT segment relies entirely on Firewall Rules configured under Settings > Firewall & Security > Firewall Rules. The logic follows a "deny by default" philosophy for inter-VLAN traffic. Apply the following rules in the LAN IN section to ensure proper isolation while maintaining functionality:
1. Allow Established and Related: This rule allows traffic from the IoT VLAN back to the Trusted LAN only if the connection was initiated by a trusted device. Set Action to "Accept," Source to "Any," and Destination to "Any," checking the "Established" and "Related" boxes.
2. Drop Invalid State: This rule discards packets that do not belong to a known connection. Set Action to "Drop," State to "Invalid."
3. Drop IoT to Trusted LAN: This is the primary isolation rule. Set Action to "Drop," Source to "IoT_Network," and Destination to "Trusted_LAN_Network." This prevents IoT devices from initiating communication with your PCs or NAS.
4. Drop IoT to Gateway: To prevent devices from accessing the UniFi OS management interface, create a LAN LOCAL rule. Set Action to "Drop," Source to "IoT_Network," and Destination to "Gateway IP Addresses" on ports 80, 443, and 22.
Lab Performance Metrics
The following table illustrates the impact of firewall inspection and VLAN routing on network performance within a UniFi Dream Machine Pro environment.
| Metric | Trusted LAN (Inter-VLAN) | IoT VLAN (Isolated) | Overhead Impact | 
|---|---|---|---|
| Max Throughput | 940 Mbps | 910 Mbps | ~3.2% | 
| Latency (ms) | 0.24 ms | 0.48 ms | +0.24 ms | 
| Jitter | 0.08 ms | 0.15 ms | Minimal | 
| CPU Usage (Peak) | 12% | 18% | Stateful Inspection Load | 
Advanced Multicast and mDNS
For features like AirPlay or Google Cast to work across the Trusted LAN and IoT VLAN, mDNS (Multicast DNS) must be enabled. In the UniFi interface, go to Settings > Services > mDNS and toggle the switch to "On." This allows the gateway to repeat discovery packets across the VLAN boundaries without compromising the firewall rules that block direct TCP/UDP access to sensitive data. If specific devices require more granular access, create "Allow" rules above the "Drop" rules using IP Groups (Port Groups) to permit only the necessary ports, such as 8008-8009 for Chromecast functionality.
[guides]
