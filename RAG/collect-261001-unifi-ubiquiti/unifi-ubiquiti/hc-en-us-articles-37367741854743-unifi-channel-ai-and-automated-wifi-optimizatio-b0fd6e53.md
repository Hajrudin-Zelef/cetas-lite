---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-37367741854743-unifi-channel-ai-and-automated-wifi-optimizatio-b0fd6e53
title: "hc-en-us-articles-37367741854743-unifi-channel-ai-and-automated-wifi-optimizatio-b0fd6e53"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-37367741854743-unifi-channel-ai-and-automated-wifi-optimizatio-b0fd6e53.md
source_anchor: ""
source_lines: [1, 27]
sha256: 138ef13c38336f6fbf0740d3d09f8cfc55ce9a7416c37db48c54e089e0761fab
---

# hc-en-us-articles-37367741854743-unifi-channel-ai-and-automated-wifi-optimizatio-b0fd6e53

UniFi Channel AI and Automated WiFi Optimization
Channel AI is UniFi’s built-in Radio Resource Management (RRM) tool that automatically analyzes your wireless environment and recommends optimal WiFi channels for all APs in your deployment.
Whether you’re setting up a new network or responding to a congested RF environment, Channel AI helps improve overall user experience, all without disrupting your clients.
How It Works
Channel AI scans the surrounding wireless environment using neighbor reports and automated RRM scans to detect nearby APs and sources of RF interference. It then uses UniFi’s proprietary AI/ML-driven algorithm, which continuously improves over time, to generate a recommended channel plan.
The result is an optimized Wi-Fi channel configuration—channel width and transmit power remain unchanged—applied non-disruptively, without impacting active clients.
Running Channel AI
We recommend using Channel AI throughout the lifecycle of your UniFi deployment—from initial setup to ongoing optimization in response to changing RF conditions or increased congestion.
To do so:
- Navigate to AirView > Radios > Channel AI View.
- Click Optimize.
- Channel AI will provide an optimization summary.
- Click Apply Changes.
- Repeat this process as needed for 2.4GHz, 5GHz, and 6GHz.
DFS Channel Inclusion
If you have dense deployments on 5GHz, and need to take advantage of DFS Channels, they can be included by following these steps:
- Navigate to Settings > WiFi > Default WiFi Speeds.
- Select Custom.
- Enable Extended 5 GHz Spectrum (DFS).
- Click Apply Changes.
If you are new to the concept of DFS Channels, we recommend familiarizing yourself using this document.
Channel Omission
If you would like to prevent Channel AI from selecting a particular channel, you can do so as follows:
- Navigate to AirView > Channel AI > Channel Plan.
- Select the Expand icon.
- Select the channel(s) you wish to omit. Bands can be changed above the Optimize button.
- Click Apply Changes.
