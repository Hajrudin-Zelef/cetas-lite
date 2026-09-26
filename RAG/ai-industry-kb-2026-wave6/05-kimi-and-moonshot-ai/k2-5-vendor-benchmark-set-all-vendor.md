---
id: ai-industry-kb-2026-wave6/05-kimi-and-moonshot-ai/k2-5-vendor-benchmark-set-all-vendor
title: "K2.5 vendor benchmark set (all [VENDOR])"
domain: kimi-and-moonshot-ai
role: deep-dive
task: benchmark
actors: ["Moonshot"]
dates: ["2026-01-27", "2026-04-20", "2026-05"]
keywords: ["benchmark", "agent", "agents", "arr", "copilot", "funding", "kimi", "mcp", "mxfp4", "research"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2400, 2418]
section: "§5. Kimi and Moonshot AI"
delta_of: ai-industry-kb-2026
sha256: 04a9fe00044839b8b68ec085b3a7bd9f7cdb434ff475bf92ed87abbcfa95655e
---

# K2.5 vendor benchmark set (all [VENDOR])

### K2.5 vendor benchmark set (all [VENDOR])
- SWE-bench Verified **76.8**; MMMU-Pro **78.5**; VideoMMMU **86.6**; HLE Full with tools **50.2**; BrowseComp **74.9** [SECONDARY reporting vendor] (marktechpost.com, 2026-01-27).
- Agent Swarm: **~4.5× faster on wide research tasks** [VENDOR] (aibase.com).
- Do not compare K2.5's SWE-bench Verified 76.8 with SWE-bench Pro figures for other models — different benchmark families [DIRECTIONAL].

### K2.6 swarm scale
- **300 sub-agents / 4,000 coordinated steps** maximum [SECONDARY] (marktechpost.com, 2026-04-20).

### K2.7 Code vendor benchmark table (all [VENDOR])
- Kimi Code Bench v2 **62.0**; Program Bench **53.6**; MLS Bench Lite **35.1**; Kimi Claw 24/7 Bench **46.9**; MCP Atlas **76.0**; MCP Mark Verified **81.1** [SECONDARY reporting vendor] (byteiota.com).
- **~30% fewer thinking tokens than K2.6** with mandatory/preserved thinking [VENDOR] (byteiota.com).

### K3 architecture figures
- 2.8T total / 104B active; 93 layers (69 KDA + 24 Gated MLA, 1 dense); 896 routed experts, 16/token, 2 shared; MoonViT-V2 ~401M; MXFP4/MXFP8; 1,048,576-token context [SECONDARY] (siml1169/kimi-copilot-provider_k3; ai-stack.ai; felloai.com; glows.ai).
- K3's 16-experts-per-token routing is 2× the K2 family's top-8 [DIRECTIONAL].

### Funding figures
- $700M @ $10B (early 2026); $18B (March, intermediate report); ~$2B led by Long-Z @ $20B+ (May 2026); $3.9B raised in six months; ARR >$200M (April); $30B = talks, not closed [SECONDARY] (mlq.ai; clay.com; theagenttimes.com; ainvest.com).

