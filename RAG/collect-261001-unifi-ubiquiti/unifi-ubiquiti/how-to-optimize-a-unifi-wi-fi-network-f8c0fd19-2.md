---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/how-to-optimize-a-unifi-wi-fi-network-f8c0fd19-2
title: "how-to-optimize-a-unifi-wi-fi-network-f8c0fd19"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/how-to-optimize-a-unifi-wi-fi-network-f8c0fd19.md
source_anchor: ""
source_lines: [57, 79]
sha256: 06ddffe8c9cc78b5207dd6ab15c7e514844856d5d1065a9a29ac35759f42b738
---

# how-to-optimize-a-unifi-wi-fi-network-f8c0fd19

Roaming Challenges
Client devices (not access points) make roaming decisions. The only thing we can really do is nudge the client device to move to a different access point. Without this, clients often stick to further access points rather than roaming quickly. There are a few settings we can optimize to make this process better (depending on your devices).
Fast Roaming (802.11r)
Fast roaming dramatically reduces the time needed for handoffs by pre-authenticating clients with other APs. This speeds up the process, but requires a device that supports 802.11r…which isn’t new, but you’d be surprised at how many devices still don’t support it.
This is why a segregated network is so important, because if you have a VLAN with newer devices, this is a great feature to turn on. This has drastically improved my Wi-Fi roaming experience. If you don’t, be careful, as this can cause connection issues.
BSS Transition
This is on by default, and improves transitions as well. This is generally safer (in my experience) to leave on than Fast Roaming, but if you experience connectivity issues, you might want to disable this.
Troubleshooting Common Issues
Here are some common issues you’ll experience when optimizing your UniFi Wi-Fi Network, and how you can fix them.
Poor Roaming Performance
If devices aren’t transitioning properly between APs:
- Reduce transmit power on APs
- Verify fast roaming is enabled for the SSID (if client devices support it)
Interference Problems
If you notice sudden performance degradation on one or more access points:
- Run a new RF environment scan (spectral or airtime)
- Look for new sources of interference
- If you manually set the channel, adjust it to a channel with less interference. If you are using auto, run the channelization optimization (settings > WiFi)
- Adjust channel width if necessary (dropping from 80 MHz to 40 MHz, etc)
Connectivity Issues After Configuration Changes
If devices struggle to connect after making changes, check the last thing you did. We aren’t adjusting that much here that can cause massive issues, but some of these settings can impact performance or connectivity. Make each to change, test, and validate that it works as expected, then move on to the next change. If you make all of the changes at the same time, you might not be sure what is causing issues.
Conclusion on UniFi WiFi Optimization
These are guidelines you can follow that should provide a solid foundation and optimize your UniFi Wi-Fi network. As always, you’ll have to monitor and adjust these settings based on your experiences, and ongoing maintenance in this area may be required, but if you carefully implement these settings, you should see a big improvement.
