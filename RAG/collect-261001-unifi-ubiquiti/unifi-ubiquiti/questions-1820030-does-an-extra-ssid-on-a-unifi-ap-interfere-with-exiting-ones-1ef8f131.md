---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1820030-does-an-extra-ssid-on-a-unifi-ap-interfere-with-exiting-ones-1ef8f131
title: "questions-1820030-does-an-extra-ssid-on-a-unifi-ap-interfere-with-exiting-ones-1ef8f131"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1820030-does-an-extra-ssid-on-a-unifi-ap-interfere-with-exiting-ones-1ef8f131.md
source_anchor: ""
source_lines: [1, 10]
sha256: ea7ec35caea431eb58b011f2b40f790d448e6557c9dd52f44cf8db37f9c9cdec
---

# questions-1820030-does-an-extra-ssid-on-a-unifi-ap-interfere-with-exiting-ones-1ef8f131

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a UniFi access point on which I have historically created a few SSIDs. I now would like to add a new one and I was wondering to which point such numerous SSIDs can be a problem (from an interference perspective)?
I do not see any setting for the channel in the setup of the SSID so I guess the answer is no but I wanted to make sure.
UniFi APs have a limit of either 4 or 8 SSIDs per band, per AP group. Some older models like the AC-Lite only support up to 4 per band. Most models can have up to 8. This means you can have up to eight 2.4 GHz and up to eight 5 GHz networks, or eight dual-band SSIDs. The same applies to 6 GHz.
If your model is not too old, the maximum is then 8 SSIDs per band.
