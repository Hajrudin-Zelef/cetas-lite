---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-72927-unifi-network-controller-wlan-channel-display-meaning-be2aa39c
title: "questions-72927-unifi-network-controller-wlan-channel-display-meaning-be2aa39c"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-72927-unifi-network-controller-wlan-channel-display-meaning-be2aa39c.md
source_anchor: ""
source_lines: [1, 10]
sha256: cf44470081893bf2167907d2f1ce052336e1265af1497c64d708f1074619c45c
---

# questions-72927-unifi-network-controller-wlan-channel-display-meaning-be2aa39c

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I don't quite understand the WLAN channel display. Elsewhere in the UniFi Network Controller, the channel is set to 1. But what does the number 3 in the display below mean?
To clearify the situation i added three more screenshots with different WLAN settings. Because the behavoir also exist with channels in mid i dont think it has to do with borders, but the channel width seem to have an effect.
In the 2.4 GHz band, there are three non-overlapping bands for 802.11b/g/n from which you can choose for the 22/20 MHz bandwidth (the EE term, not the networking term). Those are channels 1, 6. and 11 as the center of the band. Choosing a different channel can be problematic if there are neighboring devices using the non-overlapping channels because you will overlap your band with the neighbor bands, causing interference and noise in your and the neighbor bands,
With 802.11n using 40 MHz bandwidth, You must use channel 3 instead if channel 1 as the center of the band because channel 1 would have the band extend below the frequencies allowed.
