---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/3-5-zero-trust-in-the-dc
title: "3.5 Zero-trust in the DC"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [92, 101]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: c9589823c5489288b201de41f13862f2207380c1dbed508563b7d0f4f227fd7b
---

# 3.5 Zero-trust in the DC

### 3.5 Zero-trust in the DC
- Zero-trust (NIST SP 800-207) applied to DC east-west: never trust, always verify; microsegmentation is the network-layer implementation — identity/ workload-attached policy rather than perimeter zones `[secondary]` — ronutz/arsenal; NIST SP 800-207 is the reference architecture `[official]` (NIST).
- Cisco TrustSec/SGT: 16-bit Security Group Tags carried in-line (or via SXP) decouple policy from IP/VLAN; SGACLs enforce group-to-group policy in hardware — position it as fabric-native microsegmentation for Cisco shops `[vendor-reported]` — Cisco TrustSec docs (not deep-dived here; flagged for a future pass).
- Note: vendors' "zero trust" marketing bundles NAC, ZTNA, and segmentation; only the segmentation-relevant mechanisms are in scope here.

### Wave 3 verification
- Sources: 6. No conflicts. Gaps: ACL TCAM scale table; TrustSec/SGT deep dive deferred. Vendor marketing language normalized to mechanisms.

---

