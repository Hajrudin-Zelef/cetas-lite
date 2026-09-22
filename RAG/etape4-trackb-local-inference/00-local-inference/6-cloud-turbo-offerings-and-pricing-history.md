---
id: etape4-trackb-local-inference/00-local-inference/6-cloud-turbo-offerings-and-pricing-history
title: "6. Cloud / Turbo offerings and pricing history"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: pricing
actors: ["DeepSeek", "United States"]
dates: []
keywords: ["pricing", "agent", "cost", "deepseek", "gpu", "inference", "research"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [492, 546]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 17110d5613718b9a121b4c099a35aec189759f8ddebfb97429590254a5a21b56
---

# 6. Cloud / Turbo offerings and pricing history

## 6. Cloud / Turbo offerings and pricing history

> The task asked about "Ollama Cloud / Turbo". Research found **"Ollama Cloud"** as the current product name with `:cloud` models. A distinct "Turbo" branded tier was **not confirmed in sources** — [unverified]; earlier (pre-2026) reporting referenced cloud access more loosely. Flagging rather than asserting.

### Pricing eras
1. **GPU-time era (through at least Jul 2026):** cloud tiers from free to $100/month; "tracks usage based on GPU time, not token limits" [independent: TechCrunch, 9 Jul 2026].
2. **Transparent per-token era (by Sep 2026):** "transparent pricing update replaces vague cloud access with a metered model: paid plans include monthly usage credits… The unit of cost is no longer the seat or the session, it is the actual prompt and output volume" [secondary: Tech To Heart, ~6 Sep 2026].

### Current plans (per-token era) [secondary: mindstone-agent doc, verified against ollama.com/pricing 9 Sep 2026; corroborated by Tech To Heart]
| Plan | Price | Included | Concurrency |
|---|---|---|---|
| Free | $0 | starter usage credits + starter models | 1 concurrent request |
| Pro | $20/mo | $60 monthly usage credits, larger models | 3 concurrent requests (50x free usage) |
| Max | $100/mo | $300 monthly usage credits | 10 concurrent requests (250x free) |
| Team | $500/mo | $1,000 shared usage | shared billing |
- Credits can be added to any plan for pay-as-you-go access to all models [secondary].
- Local model usage is unlimited on all plans [secondary].
- Sources: https://github.com/mindstone-agent/mindstone-agent/blob/HEAD/OLLAMA_CLOUD_QWEN3.5.md ; https://techtoheart.com/ollama-clouds-off-peak-pricing-turns-scheduling-into-a-cost-lever/

### Per-model per-token rates (USD per 1M tokens; verified 8 Sep 2026 against ollama.com/pricing) [secondary: wobblebot implementation doc]
| Model | Input | Cached input | Output |
|---|---|---|---|
| `gpt-oss:120b` | 0.15 | 0.014 | 0.60 |
| `gpt-oss:20b` | 0.07 | 0.035 | 0.30 |
| `qwen3.5:397b` | 0.60 | full input rate | 3.60 |
| `gemma4:31b` | 0.14 | 0.05 | 0.40 |
- Source: https://github.com/carldog/wobblebot/blob/HEAD/docs/implementation/ollama-cloud.md

### Off-peak pricing (DeepSeek V4) [secondary: Tech To Heart, ~6 Sep 2026]
- Off-peak: **DeepSeek V4 Flash** $0.22 input / $0.007 cached / $0.66 output per 1M; **DeepSeek V4 Pro** $0.66 / $0.022 / $1.98.
- Peak (12:00–18:00 UTC, Mon–Fri): **exactly 2x** off-peak rates.
- Off-peak applies all day Saturday and Sunday.
- Source: https://techtoheart.com/ollama-clouds-off-peak-pricing-turns-scheduling-into-a-cost-lever/

### Usage mechanics (as reported by third-party operators, Sep 2026)
- Four usage levels (light→extra-heavy models); usage measured by weighted input/cached/output tokens against plan allowance [secondary: coding-plan skill doc].
- Session limits reset every 5 hours; weekly limits every 7 days (no published numbers) [secondary].
- Extra-usage regime: after included quota, billing draws from prepaid balance at per-token rates; API returns 200 in both regimes with no quota headers — only signal is 429 at full exhaustion [secondary: merchant-routing-engine doc]. *Treat operator observations as [secondary], not official.*

---

## 7. Enterprise features

- **Team plan**: conflicting reports —
  - (a) **$500/mo with $1,000 shared usage**, shared billing [secondary: Sep 2026 docs above];
  - (b) **$25/seat/mo, 5-seat minimum ($125/mo min)**, usage included per seat, overage from shared team balance billed pay-as-you-go, US + Europe model access, zero data retention, shared billing/admin, priority support [secondary: coding-plan skill doc, sign-up state observed **28 Aug 2026**].
  - **Flag:** the two figures likely reflect different page states/dates (Aug 28 vs Sep 9, 2026) or different plan presentations. Both are [secondary]; neither confirmed from the live pricing page by this research.
- **Enterprise (custom)**: custom terms and support for larger orgs; everything in Team + volume pricing, security/procurement support, deployment planning with Ollama; no published price [secondary].
- **SSO / access controls / MDM**: listed as **"Coming soon: SSO, model access controls, MDM installer for Windows and macOS"** (as of 28 Aug 2026) — i.e. **not shipped** at cut-off [secondary].
- **Team sign-up state**: "Join waitlist", not an open purchase button (28 Aug 2026) [secondary].
- The self-hosted core remains free/open-source (MIT); local inference has no usage costs beyond hardware/electricity [independent/secondary].
- Positioning gap noted by third parties: Ollama has **no built-in auth, RBAC, audit logging, or multi-tenant controls** on the local server — "175,000+ Ollama instances are exposed without auth" [secondary: solon GTM doc]; production multi-user setups pair it with gateways (LiteLLM etc.) [secondary: markaicode, Aug 2026].

---

