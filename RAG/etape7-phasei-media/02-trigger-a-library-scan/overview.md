---
id: etape7-phasei-media/02-trigger-a-library-scan/overview
title: "Trigger a library scan"
domain: trigger-a-library-scan
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [560, 561]
section: "Trigger a library scan"
sha256: f1d3a71ec781f5da3a8a50aead3149494a7218ac7ff30bc7d215966bf1f2081c
---

# Trigger a library scan
curl -s -X POST -H "X-Emby-Token: $K" "$J/Library/Refresh"
