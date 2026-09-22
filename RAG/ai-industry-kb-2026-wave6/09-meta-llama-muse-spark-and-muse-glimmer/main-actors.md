---
id: ai-industry-kb-2026-wave6/09-meta-llama-muse-spark-and-muse-glimmer/main-actors
title: "Main actors"
domain: meta-llama-muse-spark-and-muse-glimmer
role: deep-dive
task: actor-profile
actors: ["AWS", "China", "Cohere", "EU", "Google", "Hugging Face", "Meta", "Microsoft", "OpenAI", "United States", "xAI"]
dates: ["2024-04-18", "2024-07-23", "2024-09-25", "2024-10-28", "2024-12-06", "2024-12-07", "2025-04-05", "2025-04-29", "2026-01", "2026-01-26", "2026-02", "2026-02-05", "2026-03-31", "2026-05-20", "2026-07-24", "2026-08-05", "2026-08-07", "2026-08-10", "2026-09", "2026-09-02", "2026-09-07", "2026-09-20", "2026-09-22"]
keywords: ["agent", "agents", "apache", "aws", "bedrock", "benchmark", "benchmarks", "cohere", "disclosure", "distribution", "gpus", "grok"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4442, 4513]
section: "§9. Meta: Llama, Muse Spark, and Muse Glimmer"
sha256: a7c8fd431d890a16696368d5299689fb9a9f7207f7ecf3cf3938f6bec6d82dd3
---

# Main actors

## Main actors
- **Meta** — owner of Llama, Muse Spark, Muse Glimmer; runs a dual posture: gated Community License for Llama, straight Apache-2.0 for Glimmer 30B [SECONDARY].
- **Mark Zuckerberg** — published the Aug 10, 2026 open-source-AI essay arguing US policy should loosen AI training-data restrictions [SECONDARY].
- **Artificial Analysis** — independent benchmark operator; Spark 1.3 at 48 on Index v4.3 [SECONDARY].
- **OSI (Open Source Initiative)** — named Llama the reference confusing case for its Open Source AI Definition (Oct 2024) [SECONDARY].

## Timeline and context
- **2024-10-28** — OSI names Llama the reference confusing case for open-source-AI definition work [SECONDARY].
- **2025-08** — OpenAI gpt-oss released under Apache 2.0 — the Western permissive precedent Glimmer follows a year later [SECONDARY].
- **2026-03-31** — Google Gemma 4 family released, the first Gemma off the custom Gemma Terms (also Apache 2.0) [SECONDARY].
- **2026-05-20** — Cohere Command A+ released, Cohere's first Apache-2.0 model — Glimmer (Aug 10) is Meta's entry in the same 2026 "Western Apache-2.0 firsts" wave [SECONDARY].
- **2026-07-24** — Meta signs the "Open Weights and American AI Leadership" letter (25 signatories) [SECONDARY].
- **2026-08-05** — Muse Spark 1.2 API launched (closed, 1M context); Muse Code terminal agent in beta [SECONDARY].
- **2026-08-10** — Muse Glimmer 30B released under Apache 2.0 (Meta's first straight Apache-2.0 model); Zuckerberg's open-source-AI essay published the same day [SECONDARY].
- **2026-09-02** — Muse Spark 1.3 shipped [SECONDARY].
- **2026-09-07** — AA Index v4.3: Spark 1.3 at 48 [SECONDARY].
- **2026-09-20** — "The Open Weights Illusion" essay restates the open-washing critique with Llama as reference case [SECONDARY].
- September 2026 — Spark 1.2 open weights still unshipped despite the "soon" promise of Aug 10 [SECONDARY].


### New verified timeline entries — expansion

- 2024-04-18 (anchor): Llama 3 released with 8B/70B variants, 8K context, ~15T tokens. [SECONDARY, S1][SECONDARY, S3]
- 2024-07-23 (anchor): Llama 3.1 released with 8B/70B/405B variants and 128K context; 405B defined the open-frontier category. [SECONDARY, S1][SECONDARY, S3]
- 2024-09-25: Llama 3.2 released — 1B/3B edge text models and 11B/90B vision models. [SECONDARY, S1]
- 2024-12-06 (dated 12-07 in some timezones): Llama 3.3 70B released as instruction-only, text-only. [SECONDARY, S2][SECONDARY, S4]
- 2025-04-05: Llama 4 Scout, Maverick, and Behemoth announced; Scout and Maverick released same-day on Hugging Face, AWS Bedrock, and Azure. [SECONDARY, S5][SECONDARY, S7][SECONDARY, S8]
- 2025-04-29/30: LlamaCon 2025 — Meta AI app launches, Llama API enters limited preview, 1.2B cumulative downloads disclosed. [SECONDARY, S26][SECONDARY, S27][SECONDARY, S28][SECONDARY, S29]
- 2025-04 → 2026 (rolling): Behemoth internal targets slip April → June → fall 2025 or later; model never releases. [SECONDARY, S9][SECONDARY, S10][SECONDARY, S11]
- 2026-01-26/27: Meta announces premium subscription testing for Instagram, Facebook, and WhatsApp. [SECONDARY, S32]
- 2026-02-05: Vibes standalone AI video app begins testing in Brazil and Mexico. [SECONDARY, S30][SECONDARY, S31]
- 2026-08-05: Muse Spark 1.2 released (coding-focused update). [SECONDARY, S13]
- 2026-08-07: Public-board check finds no Spark 1.2 entries for the vendor's TB 2.1 / DeepSWE v1.1 claims. [SECONDARY, S12]
- 2026-08-10: Muse Glimmer 30B released under Apache 2.0. [SECONDARY, S17][SECONDARY, S18]
- 2026-09-02: Muse Spark 1.3 released with full vendor benchmark card and `max`/`xhigh` reasoning modes. [SECONDARY, S22][SECONDARY, S23]
- 2026-09-22 (research date): No LlamaCon 2026 confirmed; Behemoth still unreleased; Vibes freemium unpriced. [SECONDARY, S11][SECONDARY, S30]

## Implications
1. Meta runs the most internally divided open strategy in the corpus: a genuinely permissive Glimmer 30B (Apache 2.0) alongside the gated Llama Community License — read the LICENSE file per repo, per release, not per lab [DIRECTIONAL].
2. The Llama 4 EU multimodal exclusion is a hard legal blocker for EU vision deployments (Scout/Maverick) — procurement must treat it as a compliance question, not just marketing nuance, given the EU AI Act's open-source exemption [SECONDARY/DIRECTIONAL].
3. Spark 1.2's "open weights soon" (Aug 10) that never shipped by September 2026 is a data point in the open-washing ledger — announcements are not releases [SECONDARY/DIRECTIONAL].
4. Zuckerberg's essay asks for deregulation of training data while Meta keeps its own weights gated — both camps invoke "openness" for opposite regulatory ends; evaluate policy claims against actual licenses [DIRECTIONAL].
5. Glimmer 30B belongs to the 2026 "Western Apache-2.0 firsts" cohort (Gemma 4 in March, Command A+ in May) — the Western-labs answer to the Chinese permissive wave, one model at a time [SECONDARY/DIRECTIONAL].


### New verified implications — expansion

- Llama 3.3's positioning — a 70B model matching the 405B model on MMLU (86.0 vs 88.6) and HumanEval (88.4 vs 89.0) — was Meta's clearest signal that inference efficiency, not parameter scale, had become the product strategy for the Llama line well before the MoE pivot of Llama 4. [VENDOR, S2]
- The Llama 4 LMArena episode (`Llama-4-Maverick-03-26-Experimental` submitted instead of the public checkpoint) is a concrete instance of benchmark-submission vs public-checkpoint divergence; it justifies treating all of Meta's leaderboard-adjacent claims with checkpoint-pinned skepticism. [SECONDARY, S8]
- Behemoth's non-release over reported "insufficient capability gains" suggests Meta's internal bar for flagship releases moved from training-scale milestones to marginal-gain thresholds — a discipline signal, though the underlying reports are single-sourced. [SECONDARY, S9]
- Muse Spark's two-tier data-use pricing (Contributor tier permits training-data use; Standard does not) is the most explicit price-for-data trade in any major lab's 2026 pricing, and it reframes "cheap API" narratives: part of the discount is paid in data rights. [SECONDARY, S16][SECONDARY, S13]
- Glimmer 30B's Apache 2.0 weights with closed training data/code is the now-standard "open-weight, not open-source" pattern; the expansion's license fields deliberately separate the two to avoid the ambiguity that weaker coverage collapses. [SECONDARY, S17][SECONDARY, S20]
- Spark 1.3's `max` vs `xhigh` score gaps (OSWorld 66.9 vs 57.2; GDPval 1754 vs 1709; JobBench 64.9 vs 61.2) show reasoning-effort settings moving benchmark results by ~5–10 points — any Spark 1.3 comparison that omits the mode is underspecified. [VENDOR, S22]
- Meta's 1.2B cumulative Llama download figure (LlamaCon 2025) remains the largest disclosed open-model distribution footprint in the industry, and it underwrites Meta's platform strategy: the models are loss leaders for the Meta AI app, Vibes, and the contributor-data flywheel. [VENDOR, S26][SECONDARY, S29]
- The Vibes standalone test (Brazil/Mexico first, freemium unpriced) and the January 2026 premium-subscription tests indicate Meta is still searching for a direct-monetization layer on top of its AI distribution rather than relying solely on ads. [SECONDARY, S30][SECONDARY, S32]


- The Muse line's three releases in 28 days (Spark 1.2 on 2026-08-05, Glimmer on 2026-08-10, Spark 1.3 on 2026-09-02) show Meta shipping agent products at a cadence closer to a startup's than a platform incumbent's — consistent with the pressure narrative around open-weight competitors in the same window. [SECONDARY, S13][SECONDARY, S17][SECONDARY, S22]
- Glimmer's local-first positioning (quantized <20GB, 24–32GB GPUs) is Meta's answer to the local-agent wave; its Apache 2.0 weights make it the most permissively licensed agent model in Meta's portfolio, in contrast to the Llama Community License terms on the flagship models. [SECONDARY, S17][SECONDARY, S18]
- Spark 1.3's output-to-context ratio (~90%) is a quiet differentiator for long-horizon coding agents that need to emit large diffs and plans, not just read large contexts. [SECONDARY, S22]
- Meta's January–February 2026 monetization tests (premium subscriptions across three apps, Vibes freemium) suggest the company expects AI features to carry direct revenue, not just engagement — a strategic shift worth tracking against the contributor-pricing data flywheel. [SECONDARY, S30][SECONDARY, S32]

- Meta's 2026 agent releases cluster in August–September, immediately after the July–August open-weights wave (Grok 4.5/4.6, MiMo V2.5, GPT-5.x class) — the timing is consistent with competitive response, though no source states this causally.
- The 1.2B download figure predates the Llama 4 generation; cumulative downloads as of late 2026 were not disclosed in any source found. [DIRECTIONAL — absence of disclosure, not a sourced fact]
### Contradictions documented (not silently resolved)
- Scout training provenance: some summaries say Scout and Maverick were both partially distilled from Behemoth; another says Maverick was co-distilled while Scout trained from scratch. Both retained; neither is independently confirmed. [SECONDARY, S5][SECONDARY, S6]
- Llama 3.3 release date: 2024-12-06 vs 2024-12-07 depending on timezone/source; retained as a range. [SECONDARY, S2][SECONDARY, S4]
- Spark 1.3 Artificial Analysis scores: the base §9 records 48 (AA v4.3); newer coverage reports 61/62 without a safely matching methodology version. The two are not combined; the 48 stands as the v4.3-pinned figure. [methodology-version rule]
- Behemoth status: confirmed delayed/unreleased; "shelved in favor of Llama 5" is speculation and is not stated as fact. [SECONDARY, S9][UNVERIFIED]
- Spark 1.2's vendor benchmarks (TB 2.1 82.9%, DeepSWE v1.1 59.3%) had no public-board entries two days post-launch; carried as vendor-only claims. [VENDOR, S13][SECONDARY, S12]
- The existing §9 treats Spark 1.2 as a model launch; the expansion corrects this to a coding-focused update of the Spark line, per multiple secondary sources. [SECONDARY, S13][SECONDARY, S15]

