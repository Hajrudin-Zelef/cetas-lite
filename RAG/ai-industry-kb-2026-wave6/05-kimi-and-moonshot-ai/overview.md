---
id: ai-industry-kb-2026-wave6/05-kimi-and-moonshot-ai/overview
title: "§5. Kimi and Moonshot AI"
domain: kimi-and-moonshot-ai
role: deep-dive
task: actor-profile
actors: ["AWS", "Alibaba", "China", "DeepSeek", "Fireworks AI", "Moonshot", "OpenAI", "Z.ai"]
dates: ["2025-07", "2026-01-27", "2026-04", "2026-04-20", "2026-05-20", "2026-06-12", "2026-07-16", "2026-07-27", "2026-09"]
keywords: ["kimi", "agent", "agents", "attribution", "aws", "bedrock", "benchmark", "benchmarks", "context window", "cybersecurity", "deepseek", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2098, 2161]
section: "§5. Kimi and Moonshot AI"
delta_of: ai-industry-kb-2026
sha256: b4e231457ed58b3fe27a20e43be33da64b5ebe1475caeb08e17941aa3d810259
---

# §5. Kimi and Moonshot AI

Keywords: kimi k2, kimi k2.5, kimi k2.6, kimi k2.7 code, kimi k3, moonshot ai, agent swarm, modified mit, 100m mau, 20m monthly revenue, int4 quantization, mandatory thinking, native multimodal, 384 experts, 256k context, hosted deprecation, 01.ai exit, kai-fu lee, palantir of china, boss ai, 240 employees, deepseek qwen glm enterprise customization, coding arena 1679 elo, april 20 k2.6

## Summary

Moonshot AI's 2026 Kimi line moved from K2.5 through K2.7 Code while keeping one license constant: Modified MIT with an attribution clause at 100M MAU/$20M monthly revenue — commercial until you succeed. **Kimi K2.5 (2026-01-27)**: 1T total/32B active, 256K context, native multimodal, Agent Swarm (up to 100 parallel agents / 1,500 tool calls); its hosted API deprecation on **2026-05-20 did not remove the open weights**. **Kimi K2.6 (2026-04-20)** — the date is April 20, not April 13: 1T total/32B active across 384 experts, 256K, native multimodal, INT4 quantization, Agent Swarm up to 300 agents. **Kimi K2.7 Code (2026-06-12)**: post-trained from K2.6, native INT4, mandatory thinking mode, Modified MIT; vendor benchmarks only, with no independent SWE-bench or Terminal-Bench at launch. **Kimi K3** (API 2026-07-16, weights 2026-07-27, 2.8T) is covered in the main KB and is summarized here only. [VENDOR]

This section also records 01.AI's exit from foundation-model development: Kai-Fu Lee confirmed the company no longer builds foundation models and instead customizes DeepSeek/Qwen/GLM weights for enterprise infrastructure — described as the "Palantir of China" with a "Boss AI" product, ~240 employees, and roughly half of revenue from outside China. [SECONDARY]

## Key dated facts

### Kimi K2 (baseline, mid-July 2025)
- **Mid-July 2025** — Kimi K2: 1T total / 32B active, **15.5T training tokens**, 128K context, Modified MIT. [VENDOR]
- K2 is the pre-2026 baseline against which the 2026 K2.x releases are deltas. [DIRECTIONAL]

### Kimi K2.5 (2026-01-27)
- **2026-01-27** — Kimi K2.5: 1T total / 32B active, 256K context, native multimodal understanding, **Agent Swarm** — up to **100 parallel agents and 1,500 tool calls**. [VENDOR]
- The AWS Bedrock model card is the spec-verification source for K2.5's 1T/32B and 256K figures. [VENDOR]
- Context window was **200K at launch, expanded to 256K in April 2026** — the 256K figure is the steady-state spec. [VENDOR]
- Agent Swarm's 1,500-tool-call ceiling pairs with the 100-agent parallelism — both are vendor-stated maxima, not measured averages. [VENDOR]
- **2026-05-20** — Moonshot deprecated the **hosted** K2.5 API; the open weights were unaffected and remained available. Deprecation of the service ≠ retraction of the weights. [VENDOR]

### Kimi K2.6 (2026-04-20)
- **2026-04-20** — Kimi K2.6 (the date is April 20; pre-April-13 community sightings of a "K2.6 preview" are treated as [COMMUNITY] rumors). [VENDOR]
- Secondary reception: felloai framed K2.6 as "tying GPT-5-5 on coding" — a secondary reception claim, not a vendor statement; keep the framing where it belongs. [SECONDARY]
- 1T total / 32B active across **384 experts**; 256K context; native multimodal; **native INT4 quantization**; Modified MIT. [VENDOR]
- K2.6's two capability deltas over K2.5 are native multimodal understanding and native INT4 quantization — same 1T/32B footprint, different efficiency. [VENDOR]
- Agent Swarm scaled to **up to 300 agents**. [VENDOR]

### Kimi K2.7 Code (2026-06-12)
- **2026-06-12** — Kimi K2.7 Code: post-trained from K2.6, native INT4, **mandatory thinking mode**, Modified MIT. [VENDOR]
- "Mandatory thinking" means the thinking mode cannot be disabled — a deployment constraint for latency-sensitive integrations. [VENDOR]
- At launch there were **no independent SWE-bench or Terminal-Bench evaluations** — benchmark claims are vendor-only and must stay tagged [VENDOR]. [VENDOR]
- K2.7 Code is the corpus's example of the launch-pattern warning: vendor benchmark tables without third-party harness confirmation in the same window. [DIRECTIONAL]
- The K2.x arc is efficiency-first: K2.6's native INT4 and K2.7's mandatory thinking continue the line's token-efficiency direction rather than scaling raw parameters — the 1T total is unchanged since K2. [DIRECTIONAL]

### Kimi K3 (covered in main KB)
- K3: API **2026-07-16**, weights **2026-07-27**, 2.8T — covered in main KB § (Kimi K3); this section owns only the pointer and the K3 facts that appear in dated context below. [SECONDARY]
- K3 holds the **LMArena coding arena at 1,679 Elo** — the first open model to top a board outright (September 2026). [SECONDARY]
- K3 API pricing: **$3.00/$15.00** per M input/output (2026-07-16 price card) — the corpus's mid-tier open-weight API reference. [VENDOR]
- Kimi's commercial surface beyond weights: **free K3 quota** with a Kimi account (gateway-shaped free tier) and **Kimi Code paid from $19/mo**. [SECONDARY]

### Modified MIT (the constant license)
- All 2026 Kimi releases (K2.5, K2.6, K2.7 Code) ship under Moonshot's **Modified MIT**: above **100M MAU or $20M monthly revenue**, products must display the model attribution — "commercial until you succeed." [SECONDARY]
- Same clause structure as the Kimi K2 baseline; the attribution threshold is the commercial term to read before shipping. [DIRECTIONAL]

### 01.AI's exit from foundation models
- **Kai-Fu Lee confirmed 01.AI no longer builds foundation models**; the company now fine-tunes and customizes existing Chinese open-weight models (DeepSeek, Qwen, GLM) for enterprise AI infrastructure. [SECONDARY]
- The company is described as the **"Palantir of China"** with a **"Boss AI"** product; roughly **half of revenue comes from outside China**; headcount around **240 employees**. [SECONDARY]
- No convincing Yi model release was found in the Feb–Sep 2026 window; 01.AI is the first of the six Chinese "tigers" to exit the foundation-model race. [SECONDARY]
- 2026 "LingXi" hits in image-model results refer to a different company (Lingxi Qihang / Lingxi Technology), not 01.AI — do not conflate. [SECONDARY]
- Provenance note: the wave6 Part-1 source list does not attach a dedicated URL to the 01.AI pivot; the fact arrives as [SECONDARY] from the corpus's own secondary synthesis. It is included here for completeness, flagged accordingly. [SECONDARY]


### New verified facts — expansion (continued — K3 deployment and legal-context notes)

- **Fireworks hosts K3** at `app.fireworks.ai/models/fireworks/kimi-k3` (linked from the pi-configs research note) — third-party serving alongside the HF model card at `huggingface.co/moonshotai/Kimi-K3` [SECONDARY] (github.com/mattrobenolt/pi-configs).
- techtimes' pre-weights analysis flags the **legal framework for the hosted API**: China's **National Intelligence Law (2017), Cybersecurity Law (2017), and Data Security Law (2021)** apply to Moonshot as a Beijing-based entity — "no technical configuration of the hosted API changes that legal framework"; self-hosting addresses data at the inference layer only [SECONDARY] (techtimes.com).
- The same analysis notes the **K2 line's Modified MIT license** as the predecessor regime — consistent with the K2→K3 license-evolution reading in the previous block [SECONDARY] (techtimes.com).


### New verified facts — expansion (continued — K3 license resolution, weights, day-0 ecosystem, K3-vs-Flash economics)

