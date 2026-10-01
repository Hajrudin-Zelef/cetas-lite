---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-29
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Meta", "Microsoft", "Nvidia", "OpenAI", "OpenRouter", "Oracle", "SpaceX", "xAI"]
dates: ["2025-07", "2025-12", "2026-05", "2026-05-15"]
keywords: ["agent", "agents", "agi", "antitrust", "aws", "bedrock", "benchmark", "benchmarks", "capex", "claude", "compute", "consumer"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1083, 1159]
sha256: 689b69a3920ba21f66e865ae336ba6afb63cd9122bb39ee7f7d5ce3d6b60fbcf
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

**Personnel:**
- **Rohit Prasad** (former AI chief) departed end of 2025; **Peter DeSantis** took over the consolidated AI group in December 2025 [secondary].
- **David Luan** left Amazon (~mid-2026) [secondary].
- **Pieter Abbeel** leads Frontier Model Research (FMR) — the new top AI priority for 2026 [secondary].

**Cross-cutting notes (2026 = the capex year):** Big Tech combined 2026 capex ≈ **$730B** (Amazon ~$220B, Google $195–205B, Microsoft ~$175B, Meta $130–145B), up ~78% vs ~$410B in 2025 [secondary]. **Circular AI deals:** equity and cloud-purchase commitments flow both ways between hyperscalers and frontier labs (Alphabet↔Anthropic, Amazon↔OpenAI, Amazon↔Anthropic, Microsoft/NVIDIA↔Anthropic). Treat headline figures as announced maxima, not cash paid. **Agents dominate workloads:** OpenRouter data (Sep 2026) puts agents at ~71% of token consumption; every lab shipped agent platforms (Gemini Enterprise Agent Platform, Bedrock AgentCore, Microsoft Foundry Agent Service).

### 7.7 AWS Bedrock + AgentCore (2026)

- **OpenAI on Bedrock (Apr 28, 2026):** GPT-5.5, Codex, and OpenAI agents available on Amazon Bedrock [secondary].
- **AWS as exclusive third-party cloud for OpenAI Frontier** (OpenAI's enterprise agent platform) — term of the Feb 27, 2026 $50B partnership [secondary].
- **Stateful Runtime Environment:** joint AWS–OpenAI development, integrated with **Bedrock AgentCore**; expected to launch "within the next few months" of Feb 2026 [secondary].
- **Bedrock AgentCore** is AWS's agent-runtime counterpart to Microsoft Foundry / Gemini Enterprise Agent Platform plays [secondary].
- **Nova lifecycle (official Bedrock schedule):** Nova Premier EOL **Sep 14, 2026**; Nova Canvas + two Nova Reel versions EOL **Sep 30, 2026**; no automatic migration; new customers cut off [secondary, citing official schedule].
- **Trainium commitment:** OpenAI committed to **2 GW of Trainium capacity** for training workloads (Feb 2026 deal) [secondary].

## §8 — XAI / GROK

> **2026 at a glance — xAI/SpaceXAI:** merged into SpaceX ($1.25T, Feb 2), then the $75B SPCX IPO (Jun 12, largest ever); five Grok generations (4.20 → 4.3 → 4.5 → 4.6 → 4.7 on Sep 21, 2.1T params); Colossus toward ~2 GW / 555k GPUs with a Clean Air Act lawsuit (Apr 14, unresolved); Grok everywhere — Copilot, Bedrock, Foundry; founder exodus (6 of 12).

### 8.1 Grok model releases (Feb → Sep 2026)
- **Grok 4** (July 2025, pre-window context) — xAI's flagship reasoning model; baseline for 2026 iterations. [secondary]
- **Grok 4.1 / 4.x iterations** through early-mid 2026 (Track C).
- **Grok 4.6** — cited **Artificial Analysis Intelligence Index 61** (measured on the index revision current at the time). [independent: Artificial Analysis]
- **Grok 4.7** — cited **Artificial Analysis Intelligence Index 46** — **measured on a different index revision than Grok 4.6's 61; the two scores are NOT directly comparable.** This is the canonical example of the index-revision comparability problem (see conventions). [independent: Artificial Analysis]
- **Grok Code / Grok coding variants** — xAI's developer/coding offerings expanded through 2026 (Track C).
- **Grok Imagine** (image/video generation) continued iterations. [secondary]

### 8.2 The xAI–SpaceX merger — treated as closed (Track C)
- Track C treats the **xAI–SpaceX merger as closed** during the window. Combined entity ("SpaceXAI" in some filings — cf. the 18 Sept 2026 antitrust suit naming "SpaceXAI"). [secondary]
- Strategic logic: combined compute (Colossus + Starlink), talent, and capital; Starlink bandwidth for distributed training/inference. [secondary]
- **AI slowdown antitrust class action (18 Sept 2026)** names SpaceXAI alongside Anthropic, OpenAI, Google. [independent: AP] https://www.cnbctv18.com/technology/anthropic-openai-and-google-sued-over-alleged-deal-to-slow-ai-development-19994529.htm

### 8.3 Infrastructure — Colossus
- **Colossus** (Memphis) — xAI's flagship supercomputer; continued expansion through 2026 toward multi-hundred-thousand-GPU scale. [secondary: Track C]
- **Colossus 2** planning/expansion noted (Track C).

### 8.4 Distribution & partnerships
- **Grok in X (Twitter)** — continued default integration; premium tiers. [secondary]
- **Grok in Tesla vehicles** — rollout continued. [secondary]
- **Microsoft Azure / Oracle Cloud** hosting discussions (Track C).
- **DoD classified-network AI (1 May 2026)** — xAI among the eight companies (see §2.4). [secondary]

### 8.5 Controversies
- Grok content-moderation incidents continued through 2026 (Track C — specifics to verify).
- Regulatory scrutiny of X/xAI data practices (Track C).

### 8.6 xAI benchmark snapshot (as reported)

| Model | Artificial Analysis Intelligence Index | Comparability note |
|---|---|---|
| Grok 4.6 | 61 [independent] | Different index revision from 4.7 |
| Grok 4.7 | 46 [independent] | Different index revision from 4.6 — NOT comparable |

---

### 8.7 Grok 4.x — full release table (from Track C; all verified as shipped)

| Model | Release date | Context | API price (in/out per 1M) | Notes |
|---|---|---|---|---|
| Grok 4 / Grok 4 Heavy | Jul 9–10, 2025 | 256k | $3 / $15 | Single-agent + multi-agent (Heavy) tiers; SuperGrok Heavy $300/mo introduced [secondary] |
| Grok 4.1 / 4.1 Thinking | Nov 17, 2025 | n/a | $0.20 / $0.50 ($0.05 cached) | Silent rollout Nov 1–14, 2025; focus on EQ, creative writing, ~3× hallucination cut [secondary] |
| Grok 4.20 (= "Grok 4.2") | Feb 17, 2026 | 256k | $1.25 / $2.50 | Public beta; "rapid-learning" architecture, 4-agent parallel reasoning, medical-document feature. Branding "4.2" vs "4.20" used inconsistently across sources [secondary] |
| Grok 4.3 | Beta Apr 17, 2026 (SuperGrok Heavy); API Apr 30, 2026 | 1M | $1.25 / $2.50 ($0.20 cached) | Video input (up to 5 min), native PDF/XLSX/PPTX output, reasoning always-on; 8 legacy models retired May 15, 2026 [secondary] |
| Grok 4.5 | Jul 8, 2026 | 500k | $2 / $6 | 1.5T-param MoE "V9" architecture; co-trained with Cursor session data [vendor-reported] |
| Grok 4.6 | Aug 12, 2026 | 500k | $2 / $6 | Long-running agents focus; distributed to Copilot, Bedrock, Foundry, Gemini Enterprise Agent Platform later in August [secondary] |
| Grok 4.7 | **Sep 21, 2026** | 500k | $2 / $6 | 2.1T params; supplemental SpaceX engineering-data training; released by post-merger "SpaceXAI" [secondary] |
| Grok 5 | Not released | n/a | n/a | Training on full SpaceX historical corpus targeted before year-end [unverified — ByteVyte, single source] |

**Version check:** Grok 4, 4.1, 4.2/4.20, 4.3, 4.5, 4.6, 4.7 are all real shipped versions. Numbering skips 4.4 (no public release under that name found — likely internal). "4.20" and "4.2" refer to the same Feb 2026 release.

### 8.8 Grok benchmarks by model (vendor-reported vs independent)

**Grok 4 (Jul 2025):** Vendor-reported (xAI livestream) [vendor-reported]: HLE 25.4% (Grok 4, no tools), 44.4% (Grok 4 Heavy, multi-agent + tools); GPQA 87.5–88.9%; AIME Heavy 100%; ARC-AGI-2 16.2% (Heavy) / 15.9%. Independent [independent]: Artificial Analysis Intelligence Index 73, ahead of OpenAI o3 (70), Gemini 2.5 Pro (70), Claude Opus 4 (64) at the time.

**Grok 4.1 (Nov 2025):** Vendor claims [vendor-reported]: LMArena Text Arena #1 at 1,483 Elo (31 pts ahead of nearest rival); EQ-Bench3 1,586 (vs 1,206 for Grok 4); hallucination rate ~12.09% → ~4.22% on real-world queries; FActScore errors 9.89% → <3%. One day after the consumer launch, xAI added API access at $0.20/1M input ($0.05 cached) / $0.50/1M output — notably cheap, likely an enticement/price-cut play [secondary].

