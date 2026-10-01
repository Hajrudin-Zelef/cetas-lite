---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-115002262328-unifi-feature-guide-wireless-uplink-17e01450
title: "hc-en-us-articles-115002262328-unifi-feature-guide-wireless-uplink-17e01450"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-115002262328-unifi-feature-guide-wireless-uplink-17e01450.md
source_anchor: ""
source_lines: [1, 26]
sha256: cd14adabb40da8fc533e68235944cc0afa52f81243737f27df3d67f8f3326a3a
---

# hc-en-us-articles-115002262328-unifi-feature-guide-wireless-uplink-17e01450

Considerations for Optimal Wireless Mesh Networks
What is a Mesh Network?
A mesh network is a flexible and powerful way to extend wireless coverage across areas where running cable isn’t feasible. It consists of Access Points (APs) that are wirelessly connected to each other, with client devices connecting through these APs. This setup allows you to eliminate dead zones and create seamless wireless coverage, especially in environments where wired installations are impractical.
While hardwiring APs is the ideal method to provide the highest performance and stability, wireless meshing is an alternative—especially in homes or spaces that aren’t wired for Ethernet. With proper setup, meshed networks can deliver reliable coverage and performance throughout your environment.
UniFi offers devices like the UniFi Express lineup which is ideal to serve as both a wired gateway and a wireless meshing AP, making them perfect for hybrid deployments.
Note: To use wireless meshing, make sure it’s enabled in your UniFi Network application settings.
Setting Up Wireless Meshing
To get started:
- Enable wireless meshing by navigating to Settings>WiFi>Wireless Meshing.
- On your APs:
 
  - Enable Mesh Connect on those that will uplink wirelessly. UniFi Devices>[Select AP]>Settings>Mesh Connect
  - Enable Mesh Parent on those that will serve as the uplink source. UniFi Devices>[Select AP]>Settings>Mesh Parent
- Enable Mesh Connect on those that will uplink wirelessly. UniFi Devices>[Select AP]>Settings>Mesh Connect
Important: Do not enable Mesh Connect on APs that already have a wired uplink, as this may cause unnecessary routing complexity or loops.
Tips for the Best Wireless Mesh Performance
1. Minimize Wireless 'Hops'
Every wireless 'hop' between APs reduces available throughput. To preserve performance, we recommend limiting to a maximum of two wireless hops. Whenever possible, place a meshed AP closer to a wired AP or to another AP with fewer hops from the uplink.
2. Avoid Overloading a Single Parent AP
While multiple APs can mesh to the same parent, it’s best to limit this. Too many downstream connections can increase interference and load, reducing performance. Spread out your APs to balance the wireless backhaul.
3. Ensure Strong Signal Between Meshed APs
A strong signal ensures reliable and high-speed communication between APs. Aim for at least -60 dBm signal strength between your wireless AP and its parent. This may require repositioning APs or removing obstacles like thick walls or furniture to maintain line-of-sight.
4. Optimize the RF Environment
Mesh connections in UniFi primarily use the 5 GHz band. A crowded or noisy RF environment can degrade wireless meshing performance. Consider the following:
- Use DFS channels, which are often less congested.
- Refer to our Optimizing Wireless Client Connectivity guide for more RF best practices.
