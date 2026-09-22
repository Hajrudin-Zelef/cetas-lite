---
id: ai-industry-kb-2026-wave6/05-kimi-and-moonshot-ai/main-actors
title: "Main actors"
domain: kimi-and-moonshot-ai
role: deep-dive
task: actor-profile
actors: ["AWS", "Alibaba", "Anthropic", "China", "DeepSeek", "LongCat", "Meituan", "Microsoft", "MiniMax", "Moonshot", "OpenRouter", "Z.ai", "vLLM"]
dates: ["2026-01-27", "2026-04-13", "2026-04-20", "2026-05-20", "2026-05-25", "2026-06-12", "2026-07-16", "2026-07-26", "2026-07-27", "2026-07-28", "2026-08-02", "2026-09-07", "2026-09-21"]
keywords: ["agent", "agentic", "agents", "attention", "aws", "bedrock", "benchmark", "benchmarks", "claude", "copilot", "cost", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2419, 2522]
section: "§5. Kimi and Moonshot AI"
sha256: 86ee07ab50b5e8824c90ef130f4ea292086abb0b80d7d037ecdc1c78b0afb52f
---

# Main actors

## Main actors

- **Moonshot AI** — vendor of the Kimi K2.x/K3 line; keeps the Modified-MIT license constant while scaling Agent Swarm (100 → 300 agents); deprecated hosted K2.5 without touching open weights. [VENDOR]
- **Kai-Fu Lee / 01.AI** — confirmed the pivot out of foundation models into enterprise customization of DeepSeek/Qwen/GLM weights ("Palantir of China," "Boss AI"). [SECONDARY]
- **AWS (Bedrock)** — publisher of the Kimi K2.5 model card used for spec verification. [VENDOR]
- **Secondary K2.7 coverage cluster** — devops.com, awesomeagents.ai, aifoss.dev, aimlapi.com (token-efficiency angle). [SECONDARY]

## Timeline and context

- **2025-07 (mid)** — Kimi K2 (1T/32B, 15.5T tokens, 128K, Modified MIT). [VENDOR]
- **2026-01-27** — Kimi K2.5 (1T/32B, 256K, Agent Swarm ≤100 agents/1,500 calls). [VENDOR]
- **2026-04-20** — Kimi K2.6 (1T/32B, 384 experts, 256K, INT4, Agent Swarm ≤300). [VENDOR]
- **2026-04-13 (pre)** — Community K2.6 preview sightings [COMMUNITY] — the chatter date, not the release date. [COMMUNITY]
- **2026-05-20** — K2.5 hosted API deprecated (open weights unaffected). [VENDOR]
- **2026-06-12** — Kimi K2.7 Code (from K2.6, INT4, mandatory thinking, Modified MIT; vendor benchmarks only). [VENDOR]
- Token efficiency is K2.7 Code's stated design target in secondary coverage (devops.com) — the efficiency arc, not a parameter jump. [SECONDARY]
- **2026-07-16 / 2026-07-27** — Kimi K3 API / weights (2.8T; main KB coverage). [SECONDARY]
- **2026-09** — K3: AA v4.3 43.8; LMArena coding #1 at 1,679 Elo. [SECONDARY]


### New verified timeline entries — expansion (continued)

- **2026-07** — model-router publishes the orchestrator/worker pattern with Kimi K3 at $3/$15 OpenRouter pricing [COMMUNITY].
- **2026-08-02** — generate-commit-extension verifies Kimi/Moonshot/MiniMax/GLM endpoints [COMMUNITY].
- **~2026-09** — kimi-openrouter-chat-provider adds K2.7 Code/K3 effort variants to Copilot Chat [COMMUNITY].


### New verified timeline entries — expansion (continued)

- **2026-07-16** — Kimi K3 released [SECONDARY] (theairankings.com).
- **2026-07-26** — K3 open weights shipped (largest open-weight release ever at 2.8T) [SECONDARY] (theairankings.com).
- **2026-07-27** — K3 model card begins publishing per-benchmark harness footnotes [SECONDARY] (emergent.sh).
- **2026-07-28** — Community llm-coding-benchmark wave 2: K3 scores 95 [COMMUNITY].
- **2026-09-07** — AA Intelligence Index v4.3 rebase; K3 (max) reads 44 [SECONDARY] (theairankings.com; morphllm.com).
- **2026-09-21** — Kimi Code Bench v2 public snapshot: K3 72.9% leads [SECONDARY] (benchlm.ai).


### New verified timeline entries — expansion

- **2026-01** — Kimi Code CLI launched; K2.5-class open-weight visual agentic model released (marktechpost.com 2026-01-27 coverage) [SECONDARY].
- **2026-04-20** — K2.6 released: 1T-class open-weight MoE, 300-agent/4,000-step swarm, launch pricing $0.60/$2.50 [SECONDARY] (marktechpost.com).
- **2026-05-25** — `kimi-k2` API retirement; migrate to `kimi-k2.6` [COMMUNITY] (ccx docs).
- **Early 2026** — $700M round @ $10B valuation [SECONDARY] (clay.com).
- **2026-03** — $18B valuation reported (intermediate) [SECONDARY] (theagenttimes.com).
- **2026-05** — ~$2B round led by Meituan's Long-Z Investment @ $20B+; Tsinghua Capital, China Mobile, CPE Yuanfeng participate [SECONDARY] (mlq.ai).
- **End of 2026-08** — kimi-k2.5 and moonshot-v1 series sunset on hosted API [COMMUNITY] (dev.to).
- **2026-09** — `kimi-for-coding` model ID begins serving K2.8 Preview (community-observed) [COMMUNITY] (ccx docs).
- **2026** — Anthropic alleges distillation campaigns: 3.4M+ Claude exchanges attributed to Moonshot; 16M+/24K accounts overall (allegation, unverified) [SECONDARY] (computerworld.com; infoworld.com).
- Existing §5 anchors retained: K2.5 hosted retirement May 20 (weights-side), 01.AI pivot.

## Implications

1. Moonshot's Modified MIT is the steady commercial fact across the 2026 line; the 100M MAU / $20M threshold is the clause that turns a deployment review into a license review. [DIRECTIONAL]
2. The K2.5 hosted-deprecation episode establishes that "API deprecated" ≠ "weights removed" — a distinction worth re-checking for every 2026 host-deprecation headline. [DIRECTIONAL]
3. K2.6's date is April 20; community pre-sightings before April 13 are [COMMUNITY]-only. Treating preview-chatter as release dates repeats the Hunter-Alpha class of error. [DIRECTIONAL]
4. K2.7 Code's launch without independent benchmarks is the standing caution for vendor-only numbers at launch — hold the [VENDOR] tag until a third-party harness confirms. [DIRECTIONAL]
5. 01.AI's exit is the first confirmed foundation-model exit among the Chinese "tigers" — a strategic data point for the competitive section: the race is consolidating around labs with domestic silicon or platform leverage. [DIRECTIONAL]
6. Moonshot pairs the permissive-but-gated Modified MIT with a gateway free tier (K3 quota) and a paid coding product ($19/mo) — the license and the price card are two faces of one commercial strategy. [DIRECTIONAL]
7. K3's 1,679-Elo coding-arena #1 is the corpus's strongest open-weight human-preference datapoint; it sits alongside the AA v4.3 43.8 as the two-version-pinned K3 facts that must never be merged into one claim. [DIRECTIONAL]
8. The K2→K2.5→K2.6→K2.7 line holds total parameters at 1T while everything else moves (experts 384, swarm 100→300, INT4, mandatory thinking) — parameter count is the constant, efficiency and agenticity are the variables. [DIRECTIONAL]
9. The April-13 vs April-20 date pair is the corpus's standard "preview chatter vs formal release" pattern — always prefer the formal date and tag the chatter [COMMUNITY]. [DIRECTIONAL]
10. K2.7 Code shipped with vendor benchmarks only — the corpus's standing caution for launch-window numbers applies: hold the [VENDOR] tag until a third-party harness confirms. [VENDOR]
11. K3's only two version-pinned facts in this section (AA v4.3 43.8, LMArena 1,679 Elo) must never be merged into one claim; everything else about K3 points to the main KB. [DIRECTIONAL]


### New verified implications — expansion (continued)

- **The legal-framework note pairs with §7's H3 National Intelligence Law disclosure**: both labs' API surfaces carry the same jurisdictional caveat — it is now a standard diligence item for Chinese-lab APIs, not a Kimi-specific issue [DIRECTIONAL].


### New verified implications — expansion (continued)

- **The license question is settled enough to act on**: bespoke Kimi K3 License, MaaS>$20M/12mo needs a Moonshot agreement, branding above 100M MAU/$20M monthly revenue, internal use exempt — enterprises can plan around these terms [DIRECTIONAL].
- **The "40× on agentic workloads" finding** reframes K3's value proposition: benchmark parity with Fable 5 masks a serving-cost structure that punishes long tool-calling loops — exactly the workload K3 is marketed for [DIRECTIONAL].
- **KDA is a deployment tax**: custom attention buys 75% KV-cache savings but breaks prefix caching and needs vLLM updates — the efficiency story has a real ecosystem lag [DIRECTIONAL].
- **Moonshot's candor (admitting it trails Fable 5/Sol)** plus harness footnotes is becoming a trust moat — compare with the vendors in this wave that publish no footnotes [DIRECTIONAL].


### New verified implications — expansion (continued)

- **Two Kimi API surfaces coexist**: the subscription-based `api.kimi.com/coding` (Kimi Code) and the platform `api.moonshot.ai`/`api.moonshot.cn` (Anthropic-compatible) — community tools treat them as distinct lanes with distinct auth [DIRECTIONAL].
- **K3 effort tiers (Low/High/Max) are already user-facing in the wild** via community pickers, even though no official documentation of the tier definitions has surfaced in this research [DIRECTIONAL].
- **The orchestrator/worker pattern** (frontier brain + cheap Chinese open-weight hands) is exactly the cost structure the pricing sections of §5 predict — K3 at $3/$15 is the "hardest tasks" worker, not the budget one [DIRECTIONAL].
- **Cross-lab confirmation**: a single extension (generate-commit) defaults to `kimi-k2.6`, `glm-4.5-air`, and `MiniMax-M2.5-highspeed` — the Chinese-lab coding models are now a standard ensemble in developer tooling [DIRECTIONAL].


### New verified implications — expansion (continued)

- **Harness transparency as differentiator**: Moonshot's July 27 move to per-benchmark harness footnotes is the industry's best current practice — it makes the harness caveat checkable instead of rhetorical, and every vendor table in this knowledge base should be held to it [DIRECTIONAL].
- **Flat 1M-context pricing** (K3) vs token-hungry long-context economics (GLM-5.2's 43K tokens/task in §4): the two labs are pricing the same capability in opposite ways — watch which one the market rewards [DIRECTIONAL].
- **"End of super-cheap Chinese AI"**: K3 at ~3× its predecessor's price marks the moment Chinese labs started charging a capability premium instead of undercutting — the inverse of LongCat-2.0's §6 strategy, and a sign the segment is maturing [DIRECTIONAL].
- **Provider-fidelity tooling (Kimi-Vendor-Verifier)** acknowledges that third-party serving quality varies — the cross-provider throughput spread documented for GLM-5.3-Flash in §4 applies here too [DIRECTIONAL].


### New verified implications — expansion

- **License ambiguity as adoption tax**: with "custom Kimi K3 License" and "modified MIT + $20M MaaS trigger" both in circulation, enterprises cannot clear K3 on headlines — the license file on the exact checkpoint must be read, the same discipline the aifoss.dev licensing survey prescribes for all 2026 open-weight releases [DIRECTIONAL].
- **Infrastructure open-sourcing as moat**: releasing MoonEP, FlashKDA, and AgentEnv alongside K3 weights turns the model into a serving-stack play, mirroring MiniMax's MSA kernel strategy in §7 — the moat is the stack, not just the checkpoint [DIRECTIONAL].
- **MuonClip is a training-stability story with product consequences**: 15.5T stable tokens with no loss spike underwrites the family's aggressive post-training cadence (K2.5 → K2.6 → K2.7 in months) [DIRECTIONAL].
- **Price ladder discipline**: the K2.5→K2.6→K2.7→K3 list-price ladder ($0.60 → $0.95 → $0.95 → $3.00 input) shows Moonshot monetizing capability steps rather than undercutting — contrast with LongCat-2.0's undercut strategy in §6 [DIRECTIONAL].
- **Two-provider reality**: Moonshot API vs Kimi Code are separate providers with non-interchangeable keys — integration guides that conflate them will break [DIRECTIONAL] (gemmaclaw docs).
- **Deprecation velocity**: three API generations (kimi-k2, kimi-k2.5, moonshot-v1) sunset within ~3 months — pinning model strings without a list-models check is a production risk [DIRECTIONAL].
- **Funding structure**: Meituan's Long-Z leading Moonshot's round while Meituan builds LongCat in-house is a hedge — exposure to the open-weights frontier whether it comes from inside or outside [DIRECTIONAL].

