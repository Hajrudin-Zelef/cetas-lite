---
id: ai-industry-kb-2026-wave6/06-longcat-and-meituan/main-actors
title: "Main actors"
domain: longcat-and-meituan
role: deep-dive
task: reference
actors: ["Apple", "China", "DeepSeek", "Huawei", "LongCat", "Meituan", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "United States", "Z.ai"]
dates: ["2025-08-29", "2025-09-05", "2025-09-22", "2025-12-22", "2026-01-14", "2026-02-05", "2026-03-11", "2026-03-12", "2026-04-20", "2026-05-29", "2026-06-12", "2026-06-29", "2026-06-30", "2026-07-01", "2026-07-05", "2026-08-28", "2026-09", "2026-09-02"]
keywords: ["agent", "agentic", "apache", "ascend", "attention", "benchmark", "benchmarks", "claude", "consumer", "cost", "deepseek", "distribution"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2911, 3073]
section: "§6. LongCat and Meituan"
sha256: 3465f8dbaf8adc663b429494669681526efe885379f9526428adc1385ddde854
---

# Main actors

## Main actors

- **Meituan (LongCat team)** — vendor of the LongCat line; ran the Owl Alpha stealth episode; claims domestic-only training hardware for LongCat-2.0. [VENDOR]
- **Tencent (Hunyuan)** — vendor of Hy4 Preview (Apache 2.0) with the Hy3-Preview/Hy3 license split. [VENDOR]
- **Huawei (Noah's Ark Lab)** — vendor of openPangu 2.0; official license unverified. [SECONDARY]

## Timeline and context

- **2025-09-22/23** — LongCat-Flash-Thinking original release. [VENDOR]
- **~2025-11** — LongCat-Flash-Omni (560B/27B, 128K, MIT). [VENDOR]
- **2026-01-14** — LongCat-Flash-Thinking-2601 effective (Meituan changelog; reporting also cites Jan 16). [VENDOR]
- **2026-05-29** — Six-model LongCat Flash service sunset. [VENDOR]
- **2026-06-12** — Huawei announces openPangu 2.0 at HDC. [SECONDARY]
- **2026-06-29** — Meituan reveals the Owl Alpha stealth episode (~10.1 trillion tokens/month). [SECONDARY]
- **2026-06-30** — LongCat-2.0 revealed; openPangu 2.0 staged release begins on GitCode. [VENDOR]
- **2026-07-01** — Secondary coverage frames Owl Alpha as a stealth model that had been "quietly topping OpenRouter all along" (bitcoinlfg) — reception framing, not vendor claim. [SECONDARY]
- **2026-07-05** — LongCat-2.0 weights + inference code, MIT. [VENDOR]
- **2026-08-28** — Tencent Hy4 Preview released and open-sourced (Apache 2.0). [VENDOR]


### New verified timeline entries — expansion (continued — platform history)

- **2025-08-29** — LongCat-Flash-Chat release [VENDOR].
- **2025-09-05** — LongCat API Platform launch [VENDOR].
- **2025-09-22** — LongCat-Flash-Thinking release [VENDOR].
- **2025-12-22** — Flash-Chat upgraded (256K ctx, 9 languages) [VENDOR].
- **2026-01-14** — LongCat-Flash-Thinking-2601 release [VENDOR].
- **2026-02-05** — LongCat-Flash-Lite release (68.5B/3B, N-gram table) [VENDOR].
- **2026-03-11** — LongCat-Flash-Omni-2603 release [VENDOR].
- **2026-03-12** — Flash-Thinking auto-routed to 2601 [VENDOR].
- **2026-04-20** — LongCat-2.0-Preview (5M tokens/day quota) [VENDOR].
- **2026-05-29** — Six-model sunset (exact list above) [VENDOR].
- **2026-06-30** — LongCat-2.0 release + billing (Token Pack 30-day, Pay-As-You-Go) [VENDOR].
- **2026-09-02 / 09-10** — Enterprise services (verification, invoicing, tiered discounts) [VENDOR].


### New verified timeline entries — expansion (continued)

- **2026-01** — LongCat-Flash-Thinking-2601 technical report (arXiv 2601.16725); model card with full vendor benchmark table [VENDOR/SECONDARY].
- **2026** — LongCat-Flash-Thinking-ZigZag released (sparse-attention variant) [VENDOR].
- **2026** — LongCat-Flash-Thinking GitHub repo documents DORA RL framework and two-phase pipeline [VENDOR].


### New verified timeline entries — expansion

- **2025-09** — LongCat-Flash technical report (arXiv 2509.01322): 560B ScMoE, shortcut-connected MoE, zero-computation experts [SECONDARY] (arxiv.org).
- **2026-05-29** — Six-model platform sunset (existing §6); exact retirement list still unrecovered — open gap [DIRECTIONAL].
- **2026-06-30** — LongCat-2.0 unveiled [SECONDARY] (cryptobriefing.com).
- **2026-07-01** — LongCat-2.0 weights released: 1.6T MoE, 1M context, MIT, meituan-longcat org; already leading OpenRouter pre-release [SECONDARY] (cryptobriefing.com; venturebeat.com).
- Existing §6 anchors retained: LongCat-2.0 release coverage, 10.1T tokens/month platform volume, Owl Alpha.

## Implications

1. Thinking-2601's canonical date is the changelog's Jan 14; the Jan-16 reporting variant is a secondary-citation artifact — vendor changelogs beat coverage dates. [DIRECTIONAL]
2. The Owl-Alpha trillion/billion unit error is a live example of how secondary writeups distort magnitudes; always re-check the source unit on token-volume claims. [DIRECTIONAL]
3. LongCat-2.0's domestic-hardware claim joins DeepSeek V4 and GLM-5.1 as 2026's vendor-reported domestic-silicon set — three independent claims, zero independent verifications, all tagged accordingly. [DIRECTIONAL]
4. June 30, 2026 was a double event day: LongCat-2.0's reveal and openPangu 2.0's staged release — Chinese open weights hit the 500B+ tier from two labs on the same date. [DIRECTIONAL]
5. The Meituan changelog (Jan 14) vs coverage (Jan 16) gap is a two-day secondary-citation lag — prefer primary sources for effective dates across all six sections, not just this one. [DIRECTIONAL]
6. The 5-day reveal-to-weights gap (June 30 → July 5) is the fast end of the corpus's reveal/release spectrum — compare with API-only releases whose weights are promised months out. [DIRECTIONAL]


### New verified implications — expansion (continued)

- **Three price surfaces coexist for the same model**: $2 token packs (retail), $5–10/month aggregator subscriptions (OpenCode), and direct API pay-as-you-go — LongCat is being sold like a consumer good, a subscription, and a utility simultaneously [DIRECTIONAL].


### New verified implications — expansion (continued)

- **LSA-as-evolution-of-DSA** puts LongCat in direct architectural dialogue with DeepSeek — the Chinese labs are iterating on each other's sparse-attention work, not just scaling [DIRECTIONAL].
- **Dynamic activation breaks cost accounting**: per-token pricing for a model whose active parameters vary 33B→56B by token complexity is a pricing fiction — the "cheap tokens" story assumes average-case complexity [DIRECTIONAL].
- **The customized-Muon-at-scale note** (mbrukman) independently corroborates the Muon-family optimizer trend from §5's MuonClip — Chinese labs are converging on Muon variants for trillion-parameter stability [DIRECTIONAL].
- **Token packs "like mobile games"** is the retail packaging of the pay-as-you-go model — the same Token Pack mechanism in the official ChangeLog, now with a $2 entry point [DIRECTIONAL].


### New verified implications — expansion (continued)

- **The sunset list resolves the open question**: May 29 retired Chat, Thinking, Thinking-2601, Omni-2603, Lite, and Chat-2602-Exp — meaning the Flash family was wound down as a *line*, not just selected checkpoints, to make room for 2.0 [DIRECTIONAL].
- **API lifecycles are measured in months**: Omni-2603 (March 11) was sunset by May 29 — developers pinning to dated LongCat checkpoints should expect ~2–3-month support windows [DIRECTIONAL].
- **Quota-for-feedback (5M → 120M tokens/day)** is a clever RLHF flywheel: free inference in exchange for structured model feedback [DIRECTIONAL].
- **Enterprise billing features (invoicing, tiered discounts) landed in September 2026** — the platform is maturing from developer playground to procurement-ready in the same quarter as the 2.0 launch [DIRECTIONAL].
- **In-place vs dated upgrades coexist**: Flash-Chat was upgraded silently under one name, while Thinking got dated checkpoints (2601) and auto-routing — two versioning philosophies in one platform [DIRECTIONAL].


### New verified implications — expansion (continued)

- **Mid-training on synthesized agentic trajectories** is the recipe innovation: rather than hoping RL discovers tool-use patterns, the team injects structured agentic interaction data at mid-training because such patterns are "extremely scarce" in natural corpora — a data-side answer to the agentic cold-start problem [DIRECTIONAL].
- **ZigZag vs LongCat-2.0's LongCat Sparse Attention**: two different answers to the same long-context cost problem (layer-level SSA replacement + YaRN vs Streaming/Cross-Layer/Hierarchical indexing) — the lab is running parallel sparse-attention experiments across model lines [DIRECTIONAL].
- **Noise-injected training data** for robustness against API failures and missing data shows the training regime is explicitly modeling the messiness of real tool environments, not just clean benchmarks [DIRECTIONAL].
- **Automated task-synthesis evaluation** (keyword-generated random tasks) is a generalization check that static benchmarks can't provide — worth watching as an eval methodology [DIRECTIONAL].
- **MLX community quants at 362GB on Apple silicon** prove the 560B-class models are within reach of high-end local hardware at reduced precision — the self-hosting conversation for this tier is about RAM, not data centers [DIRECTIONAL].


### New verified implications — expansion

- **Domestic-chip training at 1.6T scale is the story**: 50,000 domestic ASICs with no NVIDIA, "first of its scale," plus a first-class NPU deployment path (SGLang-FluentLLM) — the model is a proof point for China's semiconductor self-sufficiency narrative under US export controls since late 2022 [DIRECTIONAL] (cryptobriefing.com).
- **Undercut pricing as distribution strategy**: $0.75/$2.95 vs flagship Western rates, paired with pre-release OpenRouter traction — LongCat is buying developer share with price, the inverse of Moonshot's capability-step ladder in §5 [DIRECTIONAL].
- **N-gram embedding module as architectural differentiation**: a 135B-parameter embedding-side module (not attention, not MoE experts) doing ~100× embedding-space expansion is unusual — it attacks memory I/O bottlenecks in large-batch inference rather than raw per-token quality [DIRECTIONAL] (venturebeat.com).
- **MOPD vs monolithic RLHF**: segregating post-training into Agent/Reasoning/Interaction expert clusters instead of one blended reward function is a structural bet that agentic coding needs specialized, not averaged, optimization [DIRECTIONAL] (venturebeat.com).
- **Prover variant extends the franchise**: a Lean4-specialized 560B open-weight prover with anti-reward-hacking checks (theorem consistency + legality) shows the LongCat program is not a single coding model but a family strategy [DIRECTIONAL].
- **Benchmark-provenance discipline**: the vendor README's `*` convention (in-house harness vs cited rival reports) is exactly the transparency pattern to demand everywhere — asymmetric provenance (59.5 in-house vs 58.6* cited) should never be flattened into a "beats GPT-5.5" headline without the footnote [DIRECTIONAL].

## Sources and URLs

- https://longcat.chat/platform/docs/ChangeLog.html
- https://www.marktechpost.com/2025/11/02/longcat-flash-omni-a-sota-open-source-omni-modal-model-with-560b-parameters-with-27b-activated-excelling-at-real-time-audio-visual-interaction/
- https://venturebeat.com/ai/chinese-food-delivery-firm-meituans-open-source-ai-model-longcat-flash
- https://www.testingcatalog.com/meituan-launches-longcat-2-0-1-6t-parameter-model-on-apis/
- https://www.wellfunded.news/articles/meituan-open-sources-longcat-2-agentic-coding-model-chinese-chips
- https://felloai.com/longcat-2-0/
- https://groundtruth.day/news/longcat-2-trillion-param-open-model-domestic-chips.html
- https://en.wikipedia.org/wiki/Tencent_Hy
- https://techbullion.com/tencent-hunyuan-open-sources-hy4-preview-built-through-deep-model-product-co-design/
- https://insideai.news/news/press-release/tencent-hunyuan-open-sources-hy4-preview-marking-its-third-major-release-in-six-months/9274/
- https://medium.com/@webstercameron171/hunyuan-hy4-preview-just-went-open-source-what-builders-need-to-know-d029254bce40
- https://earlyterms.com/term/hy4-preview
- https://www.ai-all.info/en/ai-models/openpangu-2-0
- https://www.huaweicentral.com/huawei-introduced-openpangu-2-0-pro-model/
- https://github.com/bespokeontology/openpangu-2.0-flash-cuda-rocm/blob/HEAD/README.md


### New sources — expansion (continued)

- https://www.marktechpost.com/2026/07/05/meituan-releases-longcat-2-0-a-1-6t-parameter-open-moe-model-with-native-1m-context-and-longcat-sparse-attention/
- https://cryptobriefing.com/meituan-longcat-2-undercuts-gpt-claude-pricing/
- https://cryptobriefing.com/china-meituan-frontier-ai-domestic-chips/
- https://medium.com/tamago-labs/review-longcat-2-0-the-ai-model-that-sells-tokens-like-mobile-games-a3e8194daddd
- https://github.com/mbrukman/longcat-2.0
- https://ayinedjimi-consultants.fr/static/pdf/longcat-2-0-llm-open-source-hardware-chinois-ascend.pdf


### New sources — expansion (continued)

- https://longcat.chat/platform/docs/ChangeLog.html
- https://github.com/meituan-longcat/LongCat-Flash-Thinking-2601


### New sources — expansion (continued)

- https://huggingface.co/meituan-longcat/LongCat-Flash-Thinking-2601
- https://huggingface.co/meituan-longcat/LongCat-Flash-Thinking-ZigZag
- https://arxiv.org/pdf/2601.16725v1.pdf
- https://huggingface.co/inferencerlabs/LongCat-Flash-Thinking-2601-MLX-5.5bit
- https://www.aibase.com/news/24677
- https://github.com/meituan-longcat/LongCat-Flash-Thinking


### New sources — expansion

- https://huggingface.co/docs/transformers/model_doc/longcat_flash
- https://github.com/brndngln/longcat-flash-prover
- https://github.com/akihikowatanabe/paper_notes/issues/4995
- https://cryptobriefing.com/meituan-longcat-2-undercuts-gpt-claude-pricing/
- https://cryptobriefing.com/meituan-longcat-2-coding-model/
- https://www.marktechpost.com/2026/07/05/meituan-releases-longcat-2-0-a-1-6t-parameter-open-moe-model-with-native-1m-context-and-longcat-sparse-attention/
- https://venturebeat.com/technology/meituan-open-sources-longcat-2-0-the-1-6t-near-frontier-agentic-coding-model-thats-been-leading-openrouter-trained-entirely-on-chinese-chips
- https://medium.com/tamago-labs/review-longcat-2-0-the-ai-model-that-sells-tokens-like-mobile-games-a3e8194daddd
- https://longcat.chat/platform/docs/ChangeLog.html
- https://github.com/meituan-longcat/LongCat-2.0
- https://arxiv.org/pdf/2509.01322
- https://ayinedjimi-consultants.fr/static/pdf/longcat-2-0-llm-open-source-hardware-chinois-ascend.pdf

*Contradictions recorded in this file: (1) Hardware identity: "domestic ASICs/chips" (multiple stronger sources) vs "Huawei Ascend 910B" (single French PDF) — keep the specific chip claim [UNVERIFIED]; (2) Terminal-Bench version RESOLVED as 2.1 via vendor README — normalize all "Terminal Bench 2" phrasings to 2.1; (3) 35T+ vs "over 30T" training tokens — variance recorded; (4) secondary benchmark table (MMLU 88.4 etc.) has unattributed methodology — do not merge with vendor README figures; (5) SWE-bench Pro vs Verified vs Multilingual are different suites — never interchange; (6) May 29 six-model sunset list now RESOLVED (see platform changelog block): Flash-Chat, Flash-Thinking, Flash-Thinking-2601, Flash-Omni-2603, Flash-Lite, Flash-Chat-2602-Exp.*

