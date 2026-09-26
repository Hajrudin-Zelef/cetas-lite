---
id: collect-260926-mikrotik/mikrotik/index-php-stc-article-download-369-398-fc120320-2
title: "QoS aware traffic shaping for TikTok live streaming over congested LAN using MikroTik queue tree"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/index-php-stc-article-download-369-398-fc120320.md
source_anchor: ""
source_lines: [95, 105]
sha256: a0ae5866ef6fe2ded95aacc0381248aa1698bf0c9a956f63744c9274ea0a3eb7
---

# QoS aware traffic shaping for TikTok live streaming over congested LAN using MikroTik queue tree

The shaping scheme does not substantially increase absolute throughput — traffic shaping reorders and prioritizes existing traffic rather than augmenting physical capacity. However, on user-relevant dimensions (latency, jitter, packet loss), the experimental configuration delivers a noticeably better profile: delay −49%, jitter slightly improved, packet loss cut nearly in half. Queue tree based prioritization is effective for protecting delay-sensitive flows against background traffic on constrained access links.

### 3.4 Practical Implications and Limitations

Suitable for environments where a few high-priority real-time flows coexist with best-effort services over a limited shared link. Limitations: single topology, fixed bandwidth, one target application — not directly generalizable to YouTube Live or higher access rates without further validation.

## 4. Conclusion

Average throughput rose only from 70.67 Kbps to 80 Kbps (both "Poor"), confirming queue-based QoS redistributes existing capacity rather than extending bandwidth. Temporal metrics improved markedly: delay 88.65 ms → 45.16 ms, jitter 95.61 ms → 89.80 ms, packet loss 7.66% → 3.78%. Queue-tree-based prioritization is an effective, practical approach to protecting TikTok Live traffic on bandwidth-constrained access links without immediate bandwidth upgrades.

Future work: multiple bandwidth levels and congestion intensities, alternative schedulers (HTB, HFSC, PCC/PCQ), adaptive policies, and QoE indicators (buffering frequency, playback smoothness) to complement network-level QoS metrics.
