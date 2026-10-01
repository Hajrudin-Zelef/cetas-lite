---
id: collect-261001-general-networking/general-networking/questions-50288-does-the-theoretical-limit-of-54-mbps-apply-to-any-wifi-protocol-f4244c34
title: "questions-50288-does-the-theoretical-limit-of-54-mbps-apply-to-any-wifi-protocol-f4244c34"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-50288-does-the-theoretical-limit-of-54-mbps-apply-to-any-wifi-protocol-f4244c34.md
source_anchor: ""
source_lines: [1, 11]
sha256: ed43b4c79ba36193348d21525fa88fa70c00a9d12b2cb2b4de0cb7d3d4e6159a
---

# questions-50288-does-the-theoretical-limit-of-54-mbps-apply-to-any-wifi-protocol-f4244c34

wireless - Does the theoretical limit of 54 Mbps apply to any WiFi protocol limited to 20 MHz wide channel, even on a 5 Ghz channel? - Network Engineering Stack Exchange
Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
3
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm wondering if I limit my channel width on a really nice (Meraki WAPs capable of 802.11ac) WAPs, to 20 MHz to stop them from overlapping, or if the Auto Width channel setting is making that change. if the limit of 54 Mbps will be reached. Or, is there something about the newer 802.11n protocol that allows for faster bitrate, even with that 20 MHz wide channel, and we've disabled 2.4 GHz radios.
For reference to the theoretical limits I speak of, this is one resource, but I'm not sure about what happens if you've turned off the 2.4 ghz radio use.
My reading of it is that by limiting an 802.11ac capable device, to 5 Ghz and to 20 MHz width, you're effectively downgrading 802.11ac to 802.11n 1X1, and the limit is actually about 72.2 Mbps, but I'm not sure I'm deducing that correctly.
Other than the addition of 256-QAM [MCS 8], and (optional) 5-8 streams, you are correct. 802.11ac's advantages are small when limited to 20MHz. In the 5GHz band, there's no real reason to limit it to 20MHz. There's plenty of channels and not a lot of overlap. (which is why 160MHz is even possible)
