---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1219439-installing-ubiquiti-unifi-ap-ac-pro-results-in-ubiquiti-unifi-df957043
title: "questions-1219439-installing-ubiquiti-unifi-ap-ac-pro-results-in-ubiquiti-unifi--df957043"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1219439-installing-ubiquiti-unifi-ap-ac-pro-results-in-ubiquiti-unifi--df957043.md
source_anchor: ""
source_lines: [1, 13]
sha256: df99f172dac89c160b57f5ed36112b5b75f31f5fee62f9e92cccaef3a1c297fe
---

# questions-1219439-installing-ubiquiti-unifi-ap-ac-pro-results-in-ubiquiti-unifi--df957043

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have previously had three Unifi AP Long Range installed that has worked like a charm. We have then extended the house and for this I installed a new UniFi AP AC PRO. However after installing the Pro the Unifi AP LRs started to become unstable, sometimes we lost connection all together.
They are not near each other so they should not disturb each others signal by working in the same frequency. Where the UniFi AP AC PRO was installed we had no WiFi signal to begin with, only from neighbors. I then tried to replace a Unifi AP LR with a Pro and this one worked much better. We have never had any problem with the AP LRs before this and I don't think it's hardware related since all became unstable at the same time.
Is there some compatibility issue that I have missed here?
Your most likely problem is that you have gone from 3 devices that can sit on 1,6, and 11 to more than 3 devices in the 2.4 Band. So you have interference; select "APs" in the tabs and then "Performance" should become available as a view, and will show some of that. The 5 GHz band that the AC pro's add should be clearer, but is shorter range.
Your second likely problem (even though you say there was no wifi there) is almost certainly too much power - You have LRs and Pros and likely are running them full blast, which is likely not ideal. This adds to your likely interference problems.
A contributing factor may be the outdated firmware you are running and/or your outdated controller. The current release versions are quite stable in my experience. The AC-Pros will have a more detailed "RF Scan" function under the new control when looking at the device - I don't know if that feature is available on the old controller you are running, I haven't used it in years.
Finally, if you are using auto channel in the 2.4 band (default) it's probably time to move to a manual channel setting scheme, as auto channel on UniFi is an "at boot time" choice (it's not re-evaluated later) and can result in poor channel choices with dense deployments.
