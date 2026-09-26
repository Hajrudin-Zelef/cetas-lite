---
id: ai-industry-kb-2026-wave6/09-meta-llama-muse-spark-and-muse-glimmer/contradictions-documented-not-silently-resolved
title: "Contradictions documented (not silently resolved)"
domain: meta-llama-muse-spark-and-muse-glimmer
role: deep-dive
task: reference
actors: ["Meta", "Microsoft", "OpenAI", "xAI"]
dates: ["2024-12-06", "2024-12-07", "2026-02", "2026-07-27", "2026-08-05", "2026-08-10", "2026-09-02"]
keywords: ["agent", "agents", "apache", "benchmarks", "disclosure", "distillation", "gpus", "grok", "grok 4", "license", "llama", "muse"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4499, 4524]
section: "§9. Meta: Llama, Muse Spark, and Muse Glimmer"
delta_of: ai-industry-kb-2026
sha256: d3439b7d56f9b7fd44cda6478678cc5e521b344abc5640df423560a9fe75b70c
---

# Contradictions documented (not silently resolved)

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

## Sources and URLs
- https://github.com/meta-llama/llama-models/blob/main/models/llama4/USE_POLICY.md
- https://theairankings.com/meta/muse-spark/
- https://agentpedia.codes/blog/muse-spark-1-3-complete-guide
- https://www.reuters.com/business/nvidia-forms-industry-alliance-open-ai-security-after-hugging-face-hack-2026-07-27/
- https://fourweekmba.com/ai-nvidia-meta-open-weights-coalition-distillation-policy/
- https://www.unite.ai/nvidia-and-microsoft-back-open-weight-ai-in-joint-letter/


### New sources — expansion

