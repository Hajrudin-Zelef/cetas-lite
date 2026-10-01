---
id: collect-261001-general-networking/general-networking/questions-1000378-azure-vpn-gateway-s2s-disabling-replay-detection-fca95559
title: "questions-1000378-azure-vpn-gateway-s2s-disabling-replay-detection-fca95559"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-1000378-azure-vpn-gateway-s2s-disabling-replay-detection-fca95559.md
source_anchor: ""
source_lines: [1, 12]
sha256: 3d012cf68de02abbb7d6d3efe03139f22e2230f265b41a3a03a1b33d22f7d5ef
---

# questions-1000378-azure-vpn-gateway-s2s-disabling-replay-detection-fca95559

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm running an Azure VPN Gateway (VpnGw1, gen1, Route-based) and trying to connect a S2S connection to a Fortigate gateway. The connection is losing connectivity every so hours and I'm wondering if I can turn off Replay Detection as a possible solution.
To my knowledge this is enabled by default (I assume for security reasons) but I can't find a setting (via powershell) to turn this off.
If this is not possible on my gen1-gateway, is it possible on gen2, policy-based or other SKU's?
This answer has been awarded bounties worth 50 reputation by Community
Show activity on this post.
Replay detection is not a tunable parameter for Azure VPN Gateways at this time. This page provides a link for all elements that are tunable via an IPSec/IKE Policy:
