---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-221029967-unifi-troubleshooting-connectivity-issues-dc138363
title: "hc-en-us-articles-221029967-unifi-troubleshooting-connectivity-issues-dc138363"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["latency", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-221029967-unifi-troubleshooting-connectivity-issues-dc138363.md
source_anchor: ""
source_lines: [1, 78]
sha256: 12c8f9446d31382546cc72060895b15cb44412103b984fee58f3e37390ac9652
---

# hc-en-us-articles-221029967-unifi-troubleshooting-connectivity-issues-dc138363

Optimizing WiFi Connectivity and Reducing Latency
WiFi optimization varies depending on your environment—what works best in a low-density residential setting may differ from the approach needed in a high-density office environment. This guide explains best practices for minimizing wireless latency (e.g., video buffering, slow load times) and preventing disconnections, whatever your environment may be.
If you're looking to maximize wireless speeds, refer to this guide.
If your devices are unable to connect to WiFi, see this troubleshooting article.
Optimizing for Low-Density Residential Environments
For home networks with few competing devices and minimal interference, follow these steps to maximize performance:
- Use Channel AI to automatically optimize your UniFi AP channelization and reduce interference. See Channel AI and Automated WiFi Optimization for more information.
  - To learn more about how low interference and airtime utilization is a key part of network planning, see our article on Resolving WiFi Performance Issues.
- 
Use High Transmission Power on APs, unless they are very close together.
  - Configure this globally in Radios > Radios, select all radios and then adjust their Settings.
  - Configure this per AP in UniFi Devices; click on an AP and adjust its Settings.
- 
Set channel width (see steps above) to:
  - 20 MHz for 2.4 GHz to reduce interference.
  - 
80 MHz for 5 GHz for maximum throughput.
    - While some users prefer 40 MHz for denser homes or 160 MHz for top performance, we recommend 80 MHz for most use cases.
  - 160-320 MHz for 6 GHz (note that not all devices support 320 MHz, but those that don’t will still connect at 160 MHz).
- 
Ensure proper device behavior in a mixed-band environment.
  - While 6 GHz offers excellent performance, some devices may "stick" to it even if a 5 GHz signal would provide better coverage.
  - If a device struggles to maintain a stable connection, first check if it supports fast roaming and ensure optimal AP placement.
  - As a last resort, you can disable 6 GHz on the device to encourage it to use a stronger 5 GHz signal.
- 
Monitor RSSI (signal strength) in Radios > Radios.
  - A signal strength of -70 dBm or better is required for stable connections.
  - If signal strength is low despite high transmission power, consider additional APs.
Optimizing for High-Density Environments (e.g., Offices)
For environments with many users and overlapping networks, follow these best practices:
- Use Channel AI to automatically optimize your UniFi AP channelization and reduce interference. See Channel AI and Automated WiFi Optimization for more information.
  - To learn more about how low interference and airtime utilization is a key part of network planning, see our article on Resolving WiFi Performance Issues.
- 
Use medium transmission power on APs to prevent them from interfering with one another.
  - Configure this globally in Radios > Radios, select all radios and then adjust their Settings.
  - Configure this per AP in UniFi Devices; click on an AP and adjust its Settings.
- 
Set channel width to:
  - 20 MHz for 2.4 GHz to reduce interference.
  - 40 MHz for 5 GHz to balance throughput and stability.
  - 160-320 MHz for 6 GHz (note that not all devices support 320 MHz, but those that don’t will still connect at 160 MHz).
- 
Watch for devices defaulting to 6 GHz.
  - Some devices may prefer 6 GHz over 5 GHz, even if the signal is weaker. If connectivity issues arise, disable 6 GHz on the device.
- 
Monitor RSSI (signal strength) in Radios > Radios.
  - A signal strength of -70 dBm or better is required for stable connections.
  - If signal strength is low despite high transmission power, consider additional APs.
Other Settings to Consider
These settings can further optimize WiFi performance in both residential and high-density environments. They are all located in Settings > WiFi; click on an SSID and find them under the Advanced section of its settings:
- Fast Roaming: Keep enabled to allow seamless transitions between APs.
- Band Steering: Keep enabled to prioritize 5 GHz over 2.4 GHz on compatible clients.
- Multicast Enhancement: Converts multicast traffic (e.g., streaming to multiple devices) into unicast to improve reliability. Some smart home devices may require it to be disabled.
Advanced Troubleshooting
If you are still experiencing WiFi issues, follow these steps:
- 
Identify affected devices.
  - If connectivity issues are unique to one specific device or model, contact Ubiquiti Support.
- 
Check the signal strength of wireless clients in UniFi Network > Radios > Coverage.
  - Each client should have a minimum signal strength of -70 dBm, though -65 dBm or better is recommended.
  - If signal strength is too low, consider:
    - Moving clients closer to the nearest AP.
    - Adding additional APs to expand WiFi coverage.
    - Setting Transmission Power to Auto in Radios > Select All > Edit Radios or per AP in UniFi Devices > select an AP > Settings.
- 
Check if affected devices are connecting to the 2.4 GHz band in UniFi Network > Radios > Coverage.
  - If a device supports 6 GHz or 5 GHz but still connects to 2.4 GHz, move it closer to an AP or add an additional AP.
  - If a device does not support 5 GHz, consider upgrading to one that does, as 2.4 GHz is more prone to congestion and poor performance.
- 
Minimize the use of wireless meshing whenever possible.
  - Wireless meshing increases interference, and each “wireless hop” reduces throughput by at least 50%.
  - If you must use meshing, ensure APs are connected at a signal strength of -60 dBm or better.
  - For optimal WiFi mesh performance, refer to this guide.
- 
Conduct a WiFi site survey to ensure optimal coverage.
  - If you have a UniFi Gateway or UniFi Cloud Gateway, use the WiFiman Mobile app (iOS / Android) to measure signal strength and latency.
  - If you do not have a UniFi Gateway, are unsure how to perform a site survey, or need help interpreting results, consider consulting a WiFi surveyor.
